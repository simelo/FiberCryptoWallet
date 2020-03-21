package history

import (
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"time"

	"sync"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/fibercrypto/fibercryptowallet/src/models"
	"github.com/fibercrypto/fibercryptowallet/src/models/transactions"

	qtCore "github.com/therecipe/qt/core"
)

var (
	logHistoryManager = logging.MustGetLogger("modelsHistoryManager")
	historyManager    *HistoryManager
)

const (
	dateTimeFormatForGo  = "2006-01-02T15:04:05"
	dateTimeFormatForQML = "yyyy-MM-ddThh:mm:ss"
)

/*
	HistoryManager
	Represent the controller of history page and all the actions over this page
*/
func init() {
	HistoryManager_QmlRegisterType2("HistoryModels", 1, 0, "HistoryManager")
}

type HistoryManager struct {
	qtCore.QObject
	walletEnv       core.WalletEnv
	newTxn          map[string][]core.Transaction
	txnFinded       map[string]struct{}
	filters         []string
	txnForAddresses map[string][]core.Transaction
	mutexForNew     sync.Mutex
	mutexForAll     sync.Mutex
	mutexForUpdate  sync.Mutex
	addresses       map[string]string
	walletsIterator core.WalletIterator
	end             chan bool
	_               func()                                    `constructor:"init"`
	_               func()                                    `signal:"newTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactionsWithFilters"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactionsWithFilters"`
	_               func(string)                              `slot:"addFilter"`
	_               func(string)                              `slot:"removeFilter"`
	_               func()                                    `slot:"update"`
}

func (hm *HistoryManager) init() {
	hm.ConnectGetTransactions(hm.getTransactions)
	hm.ConnectGetTransactionsWithFilters(hm.getTransansactionsWithFilters)
	hm.ConnectGetNewTransactions(hm.getNewTransactions)
	hm.ConnectGetNewTransactionsWithFilters(hm.getNewTransactionsWithFilters)
	hm.ConnectAddFilter(hm.addFilter)
	hm.ConnectRemoveFilter(hm.removeFilter)
	hm.ConnectUpdate(hm.updateTxns)
	hm.walletEnv = models.GetWalletEnv()

	hm.txnForAddresses = make(map[string][]core.Transaction, 0)
	hm.newTxn = make(map[string][]core.Transaction, 0)
	uptimeTicker := time.NewTicker(10 * time.Second)
	historyManager = hm
	hm.txnFinded = make(map[string]struct{}, 0)
	go func() {
		for {
			select {
			case <-uptimeTicker.C:
				logHistoryManager.Debug("Updating history")
				hm.mutexForUpdate.Lock()
				go hm.updateTxns()
			}
			historyManager = hm
		}
	}()
}

func (hm *HistoryManager) reviewForNew() {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	for _, txns := range hm.newTxn {
		for _, _ = range txns {
			hm.NewTransactions()
			return
		}
	}
}

type ByDate []*transactions.TransactionDetails

func (a ByDate) Len() int {
	return len(a)
}
func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	d1, _ := time.Parse(dateTimeFormatForGo, a[i].Date().ToString(dateTimeFormatForQML))
	d2, _ := time.Parse(dateTimeFormatForGo, a[j].Date().ToString(dateTimeFormatForQML))
	return d1.After(d2)
}

func (hm *HistoryManager) updateTxns() {
	defer hm.mutexForUpdate.Unlock()
	logHistoryManager.Info("Getting transactions of Addresses")
	hm.addresses = hm.getAddressesWithWallets()
	wltIterator := hm.walletEnv.GetWalletSet().ListWallets()
	if wltIterator == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't get transactions of Addresses")
		return
	}
	for wltIterator.Next() {
		logHistoryManager.Debug("Getting addresses history for wallet ", wltIterator.Value().GetId())
		addressIterator, err := wltIterator.Value().GetLoadedAddresses()
		if err != nil {
			logHistoryManager.Warn("Couldn't get address iterator")
			continue
		}

		for addressIterator.Next() {
			txnsIterator := addressIterator.Value().GetCryptoAccount().ListTransactions()
			if txnsIterator == nil {
				logHistoryManager.Warn("Couldn't get transaction iterator")
				continue
			}
			for txnsIterator.Next() {
				corrupted := false
				if _, exist := hm.txnFinded[txnsIterator.Value().GetId()]; !exist {
					for _, in := range txnsIterator.Value().GetInputs() {
						out, err := in.GetSpentOutput()
						if err != nil {
							corrupted = true
							break
						}
						outAddr, err := out.GetAddress()
						if err != nil {
							corrupted = true
							break
						}
						if _, exist := hm.addresses[outAddr.String()]; exist {
							hm.mutexForNew.Lock()
							_, exist2 := hm.newTxn[outAddr.String()]
							if exist2 {
								hm.newTxn[outAddr.String()] = append(hm.newTxn[outAddr.String()], txnsIterator.Value())
							} else {
								hm.newTxn[outAddr.String()] = []core.Transaction{txnsIterator.Value()}
							}
							hm.mutexForNew.Unlock()
						}
					}
					if corrupted {
						continue
					}
					for _, out := range txnsIterator.Value().GetOutputs() {
						outAddr, err := out.GetAddress()
						if err != nil {
							logHistoryManager.WithError(err).Warn("Couldn't get address")
							corrupted = true
							break
						}
						if _, exist := hm.addresses[outAddr.String()]; exist {
							hm.mutexForNew.Lock()
							_, exist2 := hm.newTxn[outAddr.String()]
							if exist2 {
								hm.newTxn[outAddr.String()] = append(hm.newTxn[outAddr.String()], txnsIterator.Value())
							} else {
								hm.newTxn[outAddr.String()] = []core.Transaction{txnsIterator.Value()}
							}
							hm.mutexForNew.Unlock()
						}
					}
					if corrupted {
						continue
					}
					hm.NewTransactions()
					hm.txnFinded[txnsIterator.Value().GetId()] = struct{}{}
				}
			}

		}
	}
}

func (hm *HistoryManager) getTransactions() []*transactions.TransactionDetails {
	hm.mutexForAll.Lock()

	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, txns := range hm.txnForAddresses {
		for _, txn := range txns {
			if _, exist := added[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	hm.mutexForAll.Unlock()
	newTxns := hm.getNewTransactions()
	txnsForReturn = append(txnsForReturn, newTxns...)
	return txnsForReturn
}

func (hm *HistoryManager) getTransansactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForAll.Lock()

	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range hm.filters {
		for _, txn := range hm.txnForAddresses[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	defer hm.mutexForAll.Unlock()
	newTxns := hm.getNewTransactionsWithFilters()
	txnsForReturn = append(txnsForReturn, newTxns...)
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactions() []*transactions.TransactionDetails {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for addr, _ := range hm.newTxn {
		for _, txn := range hm.newTxn[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				added[txn.GetId()] = struct{}{}
				txnsForReturn = append(txnsForReturn, txnDetail)
			}
			hm.mutexForAll.Lock()
			_, exist := hm.txnForAddresses[addr]
			if exist {
				hm.txnForAddresses[addr] = append(hm.txnForAddresses[addr], txn)
			} else {
				hm.txnForAddresses[addr] = []core.Transaction{txn}
			}
			hm.mutexForAll.Unlock()
		}
		hm.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range hm.filters {
		for _, txn := range hm.newTxn[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}

			hm.mutexForAll.Lock()
			_, exist := hm.txnForAddresses[addr]
			if exist {
				hm.txnForAddresses[addr] = append(hm.txnForAddresses[addr], txn)
			} else {
				hm.txnForAddresses[addr] = []core.Transaction{txn}
			}
			hm.mutexForAll.Unlock()
		}
		hm.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (hm *HistoryManager) addFilter(addr string) {
	logHistoryManager.Info("Add filter")
	alreadyIs := false
	for _, filter := range hm.filters {
		if filter == addr {
			alreadyIs = true
			break
		}
	}
	if !alreadyIs {
		hm.filters = append(hm.filters, addr)
	}
}

func (hm *HistoryManager) removeFilter(addr string) {
	logHistoryManager.Info("Remove filter")

	for i := 0; i < len(hm.filters); i++ {
		if hm.filters[i] == addr {
			hm.filters = append(hm.filters[0:i], hm.filters[i+1:]...)
			break
		}
	}
}
func (hm *HistoryManager) getAddressesWithWallets() map[string]string {
	logHistoryManager.Info("Get Addresses with wallets")
	response := make(map[string]string, 0)
	it := hm.walletEnv.GetWalletSet().ListWallets()
	if it == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't load addresses")
		return nil
	}
	for it.Next() {
		wlt := it.Value()
		addresses, err := wlt.GetLoadedAddresses()
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get loaded addresses")
			continue
		}
		for addresses.Next() {
			response[addresses.Value().String()] = wlt.GetId()
		}

	}
	return response
}

func getTxnType(txn core.Transaction, addrsMap map[string]string) int {
	var isInput, someOutputs, allOutputs = false, false, true
	var wltLblList = make(map[string]struct{}, 0)
	for _, input := range txn.GetInputs() {
		inpAddr, err := getAddressFromInput(input)
		if err != nil {
			logHistoryManager.WithError(err).Warnf("Couldn't get address from input %s", input.GetId())
			continue
		}

		if wltLbl, ok := addrsMap[inpAddr.String()]; ok {
			wltLblList[wltLbl] = struct{}{}
			isInput = true
		}
	}

	for _, output := range txn.GetOutputs() {
		outAddr, err := output.GetAddress()
		if err != nil {
			logHistoryManager.WithError(err).Warnf("Couldn't get address from output %s", output.GetId())
			continue
		}

		if wlt, ok := addrsMap[outAddr.String()]; ok {
			someOutputs = true
			if _, ok2 := wltLblList[wlt]; !ok2 {
				allOutputs = false
			}
		} else {
			allOutputs = false
		}
	}

	switch {
	case isInput && allOutputs:
		return transactions.TransactionTypeInternal
	case isInput && !allOutputs:
		return transactions.TransactionTypeSend
	case !isInput && someOutputs:
		return transactions.TransactionTypeReceive
	default:
		return transactions.TransactionTypeGeneric
	}
}

func getTxnAmount(txn core.Transaction, txnType int, addrMap map[string]string) (string, error) {
	var mainAsset = txn.SupportedAssets()[0]
	var amount uint64

	switch txnType {
	case transactions.TransactionTypeInternal:
		var inputContainAddrs = make(map[string]struct{})
		for _, input := range txn.GetInputs() {
			inpBalance, err := input.GetCoins(mainAsset)
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Could't get main balance of input %s", input.GetId())
				return "N/A", err
			}
			amount += inpBalance
			inpAddr, err := getAddressFromInput(input)
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get string address from input %s", input.GetId())
				return "N/A", err
			}
			inputContainAddrs[inpAddr.String()] = struct{}{}
		}

		for _, output := range txn.GetOutputs() {
			addrs, err := output.GetAddress()
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get address from output %s", output.GetId())
				return "N/A", err
			}

			if _, ok := inputContainAddrs[addrs.String()]; ok {
				outBalance, err := output.GetCoins(mainAsset)
				if err != nil {
					logHistoryManager.WithError(err).Warnf(
						"Couldn't get main balance from output %s", output.GetId())
					return "N/A", err
				}
				amount -= outBalance
			}
		}
		break
	case transactions.TransactionTypeSend:
		var wltInp = make(map[string]struct{})
		for _, input := range txn.GetInputs() {
			inpAddr, err := getAddressFromInput(input)
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get string address from input %s", input.GetId())
				return "N/A", err
			}
			if wlt, ok := addrMap[inpAddr.String()]; ok {
				wltInp[wlt] = struct{}{}
			}
		}
		for _, output := range txn.GetOutputs() {
			addrs, err := output.GetAddress()
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get address from output %s", output.GetId())
				return "N/A", err
			}

			if wlt, ok := addrMap[addrs.String()]; ok {
				if _, ok2 := wltInp[wlt]; ok2 {
					continue
				}
				outBalance, err := output.GetCoins(mainAsset)
				if err != nil {
					logHistoryManager.WithError(err).Warnf(
						"Couldn't get main balance from output %s", output.GetId())
					return "N/A", err
				}
				amount += outBalance
			} else {
				outBalance, err := output.GetCoins(mainAsset)
				if err != nil {
					logHistoryManager.WithError(err).Warnf(
						"Couldn't get main balance from output %s", output.GetId())
					return "N/A", err
				}
				amount += outBalance
			}
		}
		break
	case transactions.TransactionTypeReceive:
		for _, output := range txn.GetOutputs() {
			addrs, err := output.GetAddress()
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get address from output %s", output.GetId())
				return "N/A", err
			}
			if _, ok := addrMap[addrs.String()]; ok {
				outBalance, err := output.GetCoins(mainAsset)
				if err != nil {
					logHistoryManager.WithError(err).Warnf(
						"Couldn't get main balance from output %s", output.GetId())
					return "N/A", err
				}
				amount += outBalance
			}
		}
		break
	}
	accuracy, err := util.AltcoinQuotient(mainAsset)
	if err != nil {
		logHistoryManager.WithError(err).Warnf(
			"Couldn't get accuracy from asset %s", mainAsset)

		return "N/A", err
	}

	return util.FormatCoins(amount, accuracy), nil
}

func TransactionDetailsFromCoreTxn(txn core.Transaction, addresses map[string]string) (*transactions.TransactionDetails, error) {
	logHistoryManager.Info("Getting list of transactions")

	txnType := getTxnType(txn, addresses)
	txnDetails, err := transactions.NewTransactionDetailFromCoreTransaction(txn, txnType)
	if err != nil {
		logHistoryManager.WithError(err).Warnf("error obtaining transaction"+
			" details from core transaction: %s", txn.GetId())
		return nil, err
	}

	amount, err := getTxnAmount(txn, txnType, addresses)
	if err != nil {
		return nil, err
	}

	txnDetails.SetAmount(amount)

	return txnDetails, nil
}

func getAddressFromInput(input core.TransactionInput) (core.Address, error) {
	spendOutput, err := input.GetSpentOutput()
	if err != nil {
		logHistoryManager.WithError(err).Warnf(
			"Couldn't get spend output from input %s", input.GetId())
		return nil, err
	}
	return spendOutput.GetAddress()
}
