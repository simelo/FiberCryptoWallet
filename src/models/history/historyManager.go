package history

import (
	"errors"
	"sort"
	"time"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/fibercrypto/fibercryptowallet/src/models"
	"github.com/fibercrypto/fibercryptowallet/src/models/transactions"

	"github.com/fibercrypto/fibercryptowallet/src/util"

	qtCore "github.com/therecipe/qt/core"
)

var (
	logHistoryManager = logging.MustGetLogger("modelsHistoryManager")
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
	addresses       map[string]string
	walletsIterator core.WalletIterator
	end             chan bool
	_               func()                                    `constructor:"init"`
	_               func()                                    `signal:"newTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactions"`
	_               func(list *transactions.TransactionList)  `slot:"loadTransactionAsync"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactionsWithFilters"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactionsWithFilters"`
	_               func(string)                              `slot:"addFilter"`
	_               func(string)                              `slot:"removeFilter"`
	_               func()                                    `slot:"update"`
}

func (historyManager *HistoryManager) init() {
	logHistoryManager.Debug("Starting HistoryManager")
	historyManager.ConnectGetTransactions(historyManager.getTransactions)
	historyManager.ConnectLoadTransactionAsync(historyManager.loadTransactionsAsync)
	historyManager.ConnectGetTransactionsWithFilters(historyManager.getTransansactionsWithFilters)
	historyManager.ConnectGetNewTransactions(historyManager.getNewTransactions)
	historyManager.ConnectGetNewTransactionsWithFilters(historyManager.getNewTransactionsWithFilters)
	historyManager.ConnectAddFilter(historyManager.addFilter)
	historyManager.ConnectRemoveFilter(historyManager.removeFilter)
	historyManager.ConnectUpdate(historyManager.updateTxns)
	historyManager.walletEnv = models.GetWalletEnv()

	historyManager.txnForAddresses = make(map[string][]core.Transaction, 0)
	historyManager.newTxn = make(map[string][]core.Transaction, 0)
	historyManager = historyManager
	historyManager.txnFinded = make(map[string]struct{}, 0)
}

func (historyManager *HistoryManager) reviewForNew() {
	// historyManager.mutexForNew.Lock()
	// defer historyManager.mutexForNew.Unlock()
	for _, txns := range historyManager.newTxn {
		for _, _ = range txns {
			historyManager.NewTransactions()
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

func (historyManager *HistoryManager) updateTxns() {
	// defer historyManager.mutexForUpdate.Unlock()
	logHistoryManager.Info("Getting transactions of Addresses")
	historyManager.addresses = historyManager.getAddressesWithWallets()
	wltIterator := historyManager.walletEnv.GetWalletSet().ListWallets()
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
		var newTxnsFinded bool
		for addressIterator.Next() {
			newTxnsFinded = false
			txnsIterator := addressIterator.Value().GetCryptoAccount().ListTransactions()
			if txnsIterator == nil {
				logHistoryManager.Warn("Couldn't get transaction iterator")
				continue
			}
			for txnsIterator.Next() {
				corrupted := false
				if _, exist := historyManager.txnFinded[txnsIterator.Value().GetId()]; !exist {
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
						if _, exist := historyManager.addresses[outAddr.String()]; exist {
							// historyManager.mutexForNew.Lock()
							_, exist2 := historyManager.newTxn[outAddr.String()]
							if exist2 {
								historyManager.newTxn[outAddr.String()] = append(historyManager.newTxn[outAddr.String()], txnsIterator.Value())
							} else {
								historyManager.newTxn[outAddr.String()] = []core.Transaction{txnsIterator.Value()}
							}
							// historyManager.mutexForNew.Unlock()
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
						if _, exist := historyManager.addresses[outAddr.String()]; exist {
							// historyManager.mutexForNew.Lock()
							_, exist2 := historyManager.newTxn[outAddr.String()]
							if exist2 {
								historyManager.newTxn[outAddr.String()] = append(historyManager.newTxn[outAddr.String()], txnsIterator.Value())
							} else {
								historyManager.newTxn[outAddr.String()] = []core.Transaction{txnsIterator.Value()}
							}
							// historyManager.mutexForNew.Unlock()
						}
					}
					if corrupted {
						continue
					}
					newTxnsFinded = true
					historyManager.txnFinded[txnsIterator.Value().GetId()] = struct{}{}
				}
			}
			if newTxnsFinded {
				models.Helper.RunInMain(func() {
					historyManager.NewTransactions()
				})
			}

		}
	}
}

func (historyManager *HistoryManager) getTransactions() []*transactions.TransactionDetails {
	// historyManager.mutexForAll.Lock()

	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, txns := range historyManager.txnForAddresses {
		for _, txn := range txns {
			if _, exist := added[txn.GetId()]; !exist {
				// txnDetail, err := TransactionDetailsFromCoreTxn(txn, historyManager.addresses)
				// if err != nil {
				// 	logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				// }
				// txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	// historyManager.mutexForAll.Unlock()
	newTxns := historyManager.getNewTransactions()
	txnsForReturn = append(txnsForReturn, newTxns...)
	return txnsForReturn
}

func (historyManager *HistoryManager) getTransansactionsWithFilters() []*transactions.TransactionDetails {
	// historyManager.mutexForAll.Lock()

	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range historyManager.filters {
		for _, txn := range historyManager.txnForAddresses[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				// txnDetail, err := TransactionDetailsFromCoreTxn(txn, historyManager.addresses)
				// if err != nil {
				// 	logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				// }
				// txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	// defer historyManager.mutexForAll.Unlock()
	newTxns := historyManager.getNewTransactionsWithFilters()
	txnsForReturn = append(txnsForReturn, newTxns...)
	return txnsForReturn
}

func (historyManager *HistoryManager) getNewTransactions() []*transactions.TransactionDetails {
	// historyManager.mutexForNew.Lock()
	// defer historyManager.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for addr, _ := range historyManager.newTxn {
		for _, txn := range historyManager.newTxn[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				// txnDetail, err := TransactionDetailsFromCoreTxn(txn, historyManager.addresses)
				// if err != nil {
				// 	logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				// }
				// added[txn.GetId()] = struct{}{}
				// txnsForReturn = append(txnsForReturn, txnDetail)
			}
			// historyManager.mutexForAll.Lock()
			_, exist := historyManager.txnForAddresses[addr]
			if exist {
				historyManager.txnForAddresses[addr] = append(historyManager.txnForAddresses[addr], txn)
			} else {
				historyManager.txnForAddresses[addr] = []core.Transaction{txn}
			}
			// historyManager.mutexForAll.Unlock()
		}
		historyManager.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (historyManager *HistoryManager) getNewTransactionsWithFilters() []*transactions.TransactionDetails {
	// historyManager.mutexForNew.Lock()
	// defer historyManager.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range historyManager.filters {
		for _, txn := range historyManager.newTxn[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				// txnDetail, err := TransactionDetailsFromCoreTxn(txn, historyManager.addresses)
				// if err != nil {
				// 	logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				// }
				// txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}

			// historyManager.mutexForAll.Lock()
			_, exist := historyManager.txnForAddresses[addr]
			if exist {
				historyManager.txnForAddresses[addr] = append(historyManager.txnForAddresses[addr], txn)
			} else {
				historyManager.txnForAddresses[addr] = []core.Transaction{txn}
			}
			// historyManager.mutexForAll.Unlock()
		}
		historyManager.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (historyManager *HistoryManager) addFilter(addr string) {
	logHistoryManager.Info("Add filter")
	alreadyIs := false
	for _, filter := range historyManager.filters {
		if filter == addr {
			alreadyIs = true
			break
		}
	}
	if !alreadyIs {
		historyManager.filters = append(historyManager.filters, addr)
	}
}

func (historyManager *HistoryManager) removeFilter(addr string) {
	logHistoryManager.Info("Remove filter")

	for i := 0; i < len(historyManager.filters); i++ {
		if historyManager.filters[i] == addr {
			historyManager.filters = append(historyManager.filters[0:i], historyManager.filters[i+1:]...)
			break
		}
	}
}

func (historyManager *HistoryManager) getAddressesWithWallets() map[string]string {
	logHistoryManager.Info("Get Addresses with wallets")
	response := make(map[string]string, 0)
	it := historyManager.walletEnv.GetWalletSet().ListWallets()
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

func (historyManager *HistoryManager) loadTransactionsAsync(list *transactions.TransactionList) {
	logHistoryManager.Info("Load transactions asynchronously")
	txnList, err := historyManager.getTxn()
	if err != nil {
		logHistoryManager.Error(err)
	}
	list.AddMultipleTransactions(txnList)
}

func (historyManager *HistoryManager) getTxn() ([]*transactions.TransactionDetails, error) {
	logHistoryManager.Info("Getting transactions of Addresses")
	addresses := historyManager.getAddressesWithWallets()

	wltIterator := historyManager.walletEnv.GetWalletSet().ListWallets()
	if wltIterator == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't get transactions of Addresses")
		return make([]*transactions.TransactionDetails, 0), errors.New("couldn't get transactions of Addresses")
	}

	txnFind := make(map[string][]string)

	txnList := make([]core.Transaction, 0)
	txnDeetlList := make([]*transactions.TransactionDetails, 0)
	for wltIterator.Next() {
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
				_, ok := txnFind[txnsIterator.Value().GetId()]
				if !ok {
					txD, err := transactions.NewTransactionDetailFromCoreTransaction(txnsIterator.Value(),
						1, []string{}, []string{})

					if err != nil {
						panic("Do something") // TODO
					}
					txnDeetlList = append(txnDeetlList, txD)

					txnList = append(txnList, txnsIterator.Value())
				}

				txnFind[txnsIterator.Value().GetId()] = append(txnFind[txnsIterator.Value().GetId()],
					addressIterator.Value().String())

			}
		}
	}

	var getTxnType = func(txn core.Transaction) int {
		var txnType = transactions.TransactionTypeGeneric
		for _, input := range txn.GetInputs() {
			spendtOut, err := input.GetSpentOutput()
			if err != nil {
				logHistoryManager.Warnf("Couldn't get spendOutput from input %s", input.GetId())
				continue
			}
			addr, err := spendtOut.GetAddress()
			if err != nil {
				logHistoryManager.Warnf("Couldn't get spendOutput from spendtOut %s", spendtOut.GetId())
				continue
			}
			if _, ok := addresses[addr.String()]; ok {
				txnType = transactions.TransactionTypeSend
			}
		}

		for _, output := range txn.GetOutputs() {
			addr, err := output.GetAddress()
			if err != nil {
				logHistoryManager.Warnf("Couldn't get spendOutput from output %s", output.GetId())
				continue
			}
			if _, ok := addresses[addr.String()]; ok {

				if txnType == transactions.TransactionTypeSend {
					return transactions.TransactionTypeInternal
				} else {
					return transactions.TransactionTypeReceive
				}
			}
		}
		return txnType
	}

	var getAmountFromTxn = func(txn core.Transaction, txnType int) string {
		var mainAsset = txn.SupportedAssets()[0]
		var amount, accuracy uint64

		getAmountOfOutputs := func() uint64 {
			var outputAmount uint64
			for _, output := range txn.GetOutputs() {
				addr, err := output.GetAddress()
				if err != nil {
					logHistoryManager.Warnf("Couldn't get spendOutput from output %s", output.GetId())
					continue
				}
				if _, ok := addresses[addr.String()]; ok {
					coins, err := output.GetCoins(mainAsset)
					if err != nil {
						logHistoryManager.WithError(err).Error(
							"Could't get main balance of output with address %s", addr.String())
					}
					outputAmount += coins
				}
			}
			return outputAmount
		}

		accuracy, err := util.AltcoinQuotient(mainAsset)
		if err != nil {
			logHistoryManager.WithError(err).Error(
				"Could't get accuracy for %s asset", mainAsset)
			return "N/A"
		}

		switch txnType {
		case transactions.TransactionTypeReceive:
			amount += getAmountOfOutputs()
			break

		case transactions.TransactionTypeSend:
			for _, input := range txn.GetInputs() {
				coins, err := input.GetCoins(mainAsset)
				if err != nil {
					logHistoryManager.WithError(err).Error(
						"Could't get main balance of input with address %s", input.GetId())
				}
				amount += coins
			}
			amount -= getAmountOfOutputs()
			break
		case transactions.TransactionTypeInternal:
			isInput := func(addr string) bool {
				for _, input := range txn.GetInputs() {
					spendtOut, err := input.GetSpentOutput()
					if err != nil {
						logHistoryManager.Warnf("Couldn't get spendOutput from input %s", input.GetId())
						continue
					}
					inpAddr, err := spendtOut.GetAddress()
					if err != nil {
						logHistoryManager.Warnf("Couldn't get address from spendOutput %s", input.GetId())
						continue
					}
					if inpAddr.String() == addr {
						return true
					}
				}
				return false
			}
			for _, output := range txn.GetOutputs() {
				addr, err := output.GetAddress()
				if err != nil {
					logHistoryManager.Warnf("Couldn't get address of output %s", output.GetId())
					continue
				}
				if !isInput(addr.String()) {
					coins, err := output.GetCoins(mainAsset)
					if err != nil {
						logHistoryManager.WithError(err).Error(
							"Could't get main balance of output with address %s", addr)
					}
					amount += coins
				}
			}
		}
		return util.FormatCoins(amount, accuracy)
	}

	txnDetailsList := make([]*transactions.TransactionDetails, 0)
	for e := range txnList {
		wltList := make([]string, 0)
		findWlt := make(map[string]struct{})
		for _, addr := range txnFind[txnList[e].GetId()] {
			if _, ok := findWlt[addresses[addr]]; !ok {
				wltList = append(wltList, addresses[addr])
				findWlt[addresses[addr]] = struct{}{}
			}
		}

		txnDetail, err := transactions.NewTransactionDetailFromCoreTransaction(txnList[e], getTxnType(txnList[e]),
			txnFind[txnList[e].GetId()], wltList)

		if err != nil {
			logHistoryManager.WithError(err).Errorf(
				"Could't get a Transaction Details from the transaction: %s", txnList[e].GetId())
		}
		txnDetail.SetAmount(getAmountFromTxn(txnList[e], getTxnType(txnList[e])))

		txnDetailsList = append(txnDetailsList, txnDetail)
	}

	sort.Sort(ByDate(txnDetailsList))
	return txnDetailsList, nil
}
