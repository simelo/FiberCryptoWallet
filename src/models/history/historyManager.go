package history

import (
	"errors"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
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
	walletEnv core.WalletEnv
	filters   map[string]struct{}
	update    chan bool
	_         func()                                   `constructor:"init"`
	_         func(list *transactions.TransactionList) `slot:"loadTransactionAsync"`
	_         func(string)                             `slot:"addFilter"`
	_         func(string)                             `slot:"removeFilter"`
	_         func()                                   `signal:"hasChanged"`
}

func (historyManager *HistoryManager) init() {
	logHistoryManager.Debug("Starting HistoryManager")
	historyManager.ConnectLoadTransactionAsync(historyManager.loadTransactionsAsync)
	historyManager.ConnectAddFilter(historyManager.addFilter)
	historyManager.ConnectRemoveFilter(historyManager.removeFilter)
	historyManager.walletEnv = models.GetWalletEnv()
	historyManager.update = make(chan bool)
	historyManager.filters = make(map[string]struct{}, 0)
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

func (historyManager *HistoryManager) addFilter(addr string) {
	logHistoryManager.Info("Add filter")
	if _, ok := historyManager.filters[addr]; !ok {
		historyManager.filters[addr] = struct{}{}
	}
}

func (historyManager *HistoryManager) removeFilter(addr string) {
	logHistoryManager.Info("Remove filter")
	delete(historyManager.filters, addr)
}

func (historyManager *HistoryManager) loadTransactionsAsync(list *transactions.TransactionList) {
	logHistoryManager.Info("Load transactions asynchronously")
	var withFilter = false
	var getTxns = func() []*transactions.TransactionDetails {
		txnList, err := getTxnList(historyManager.walletEnv.GetWalletSet().ListWallets(), withFilter, historyManager.filters)
		if err != nil {
			// TODO
			return []*transactions.TransactionDetails{}
		}
		return txnList
	}
	go func() {
		list.AddMultipleTransactions(getTxns())
		for {
			select {
			case <-time.After(time.Duration(config.GetDataUpdateTime()) * time.Second):
				list.AddMultipleTransactions(getTxns())
			case withFilter = <-historyManager.update:
				list.AddMultipleTransactions(getTxns())
			case <-list.Ctx.Done():
				return
			}
		}

	}()
}

func getAddrMapByWalletIter(wltIter core.WalletIterator) (map[string]core.Address, error) {
	var addrList = make(map[string]core.Address, 0)

	for wltIter.Next() {
		addrIter, err := wltIter.Value().GetLoadedAddresses()
		if err != nil {
			logHistoryManager.WithError(err).Warnf("Couldn't get address for wallet: %s",
				wltIter.Value().GetLabel())
			return nil, err
		}

		for addrIter.Next() {
			addrList[addrIter.Value().String()] = addrIter.Value()
		}
	}
	return addrList, nil
}

func getTxnAmount(txn core.Transaction, txnType int, addrMap map[string]core.Address) (string, error) {
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
		for _, input := range txn.GetInputs() {
			inpAddr, err := getAddressFromInput(input)
			if err != nil {
				logHistoryManager.WithError(err).Warnf(
					"Couldn't get string address from input %s", input.GetId())
				return "N/A", err
			}
			if _, ok := addrMap[inpAddr.String()]; ok {
				inpBalance, err := input.GetCoins(mainAsset)
				if err != nil {

					logHistoryManager.WithError(err).Warnf(
						"Could't get main balance of input %s", input.GetId())
					return "N/A", err
				}
				amount += inpBalance
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

func getTxnType(txn core.Transaction, addrsMap map[string]core.Address) int {
	var isInput, someOutputs, allOutputs = false, false, true
	for _, input := range txn.GetInputs() {
		inpAddr, err := getAddressFromInput(input)
		if err != nil {
			logHistoryManager.WithError(err).Warnf("Couldn't get address from input %s", input.GetId())
			continue
		}

		if _, ok := addrsMap[inpAddr.String()]; ok {
			isInput = true
			break
		}
	}

	for _, output := range txn.GetOutputs() {
		outAddr, err := output.GetAddress()
		if err != nil {
			logHistoryManager.WithError(err).Warnf("Couldn't get address from output %s", output.GetId())
			continue
		}

		if _, ok := addrsMap[outAddr.String()]; ok {
			someOutputs = true
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

func getTxnList(walletIter core.WalletIterator, withFilter bool, filterAddress map[string]struct{}) ([]*transactions.TransactionDetails, error) {
	logHistoryManager.Info("Getting list of transactions")
	if walletIter == nil {
		return nil, errors.New("wallet iterator is empty")
	}
	addrsMap, err := getAddrMapByWalletIter(walletIter)
	if err != nil {
		logHistoryManager.Warn(err)
		return nil, err
	}

	var txnList = make([]*transactions.TransactionDetails, 0)
	var atomicTxn = make(map[string]struct{})

	for strAddr, corAddr := range addrsMap {
		if _, ok := filterAddress[strAddr]; !ok && withFilter {
			continue
		}
		txnIter := corAddr.GetCryptoAccount().ListTransactions()
		if txnIter == nil {
			logHistoryManager.Warn("transaction iterator of address %s is empty", strAddr)
			continue
		}

		for txnIter.Next() {
			if _, ok := atomicTxn[txnIter.Value().GetId()]; !ok {
				txnType := getTxnType(txnIter.Value(), addrsMap)
				txnDetails, err := transactions.NewTransactionDetailFromCoreTransaction(txnIter.Value(), txnType)
				if err != nil {
					logHistoryManager.WithError(err).Warnf("error obtaining transaction"+
						" details from core transaction: %s", txnIter.Value().GetId())
					return nil, err
				}

				amount, err := getTxnAmount(txnIter.Value(), txnType, addrsMap)
				if err != nil {
					return nil, err
				}
				txnDetails.SetAmount(amount)
				txnList = append(txnList, txnDetails)

				atomicTxn[txnIter.Value().GetId()] = struct{}{}
			}
		}
	}
	sort.Sort(ByDate(txnList))
	return txnList, nil
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
