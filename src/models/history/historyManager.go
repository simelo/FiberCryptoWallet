package history

import (
	"context"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"sort"
	"time"

	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/fibercrypto/fibercryptowallet/src/core"

	// local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/models"
	"github.com/fibercrypto/fibercryptowallet/src/models/transactions"
	qtCore "github.com/therecipe/qt/core"
)

var logHistoryManager = logging.MustGetLogger("modelsHistoryManager")

func init() {
	HistoryManager_QmlRegisterType2("HistoryModels", 1, 0, "HistoryManager")
}

const (
	dateTimeFormatForGo  = "2006-01-02T15:04:05"
	dateTimeFormatForQML = "yyyy-MM-ddThh:mm:ss"
)

/*
	HistoryManager
	Represent the controller of history page and all the actions over this page
*/

type HistoryManager struct {
	qtCore.QObject
	// filtersNew map[string]bool // new filter (^-^)
	filters    []string                                  // Old filter
	_          func()                                    `constructor:"init"`
	_          func() []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
	_          func() []*transactions.TransactionDetails `slot:"loadHistory"`
	_          func(string)                              `slot:"addFilter"`
	_          func(string)                              `slot:"removeFilter"`
	ctx        context.Context
	cancelFunc context.CancelFunc
	ownTxns    []core.Transaction
	walletEnv  core.WalletEnv
}

func (historyMang *HistoryManager) init() {
	historyMang.ConnectLoadHistoryWithFilters(historyMang.loadHistoryWithFilters)
	historyMang.ConnectLoadHistory(historyMang.loadHistory)
	historyMang.ConnectAddFilter(historyMang.addFilter)
	historyMang.ConnectRemoveFilter(historyMang.removeFilter)
	historyMang.walletEnv = models.GetWalletEnv()
	historyMang.ctx, historyMang.cancelFunc = context.WithCancel(context.Background())
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
func (historyMang *HistoryManager) getTransactionsOfAddresses(filterAddresses []string) []*transactions.TransactionDetails {
	logHistoryManager.Info("Getting transactions of Addresses")
	addresses := historyMang.getAddressesWithWallets()

	find := make(map[string]struct{}, len(filterAddresses))
	for _, addr := range filterAddresses {
		find[addr] = struct{}{}
	}

	txnFind := make(map[string]struct{})
	txnsList := make([]core.Transaction, 0)

	wltIterator := historyMang.walletEnv.GetWalletSet().ListWallets()

	if wltIterator == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't get transactions of Addresses")
		return make([]*transactions.TransactionDetails, 0)
	}

	for wltIterator.Next() {
		addressIterator, err := wltIterator.Value().GetLoadedAddresses()
		if err != nil {
			logHistoryManager.Warn("Couldn't get address iterator")
			continue
		}
		for addressIterator.Next() {
			_, ok := find[addressIterator.Value().String()]
			if ok {
				txnsIterator := addressIterator.Value().GetCryptoAccount().ListTransactions()
				if txnsIterator == nil {
					logHistoryManager.Warn("Couldn't get transaction iterator")
					continue
				}
				for txnsIterator.Next() {
					_, ok2 := txnFind[txnsIterator.Value().GetId()]
					if !ok2 {
						txnsList = append(txnsList, txnsIterator.Value())
						txnFind[txnsIterator.Value().GetId()] = struct{}{}
					}
				}
			}
		}
	}

	txnsDetailsList := make([]*transactions.TransactionDetails, 0)

	var getTxnType = func(txn core.Transaction) int {
		var txnType = transactions.TransactionTypeGeneric
		for _, input := range txn.GetInputs() {
			logHistoryManager.Info(input.GetSpentOutput().GetAddress().String())
			if _, ok := addresses[input.GetSpentOutput().GetAddress().String()]; ok {
				txnType = transactions.TransactionTypeSend
			}
		}

		for _, output := range txn.GetOutputs() {
			if _, ok := addresses[output.GetAddress().String()]; ok {

				if txnType == transactions.TransactionTypeSend {
					logHistoryManager.Info("internal\n\n\n\n")
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
				if _, ok := addresses[output.GetAddress().String()]; ok {
					coins, err := output.GetCoins(mainAsset)
					if err != nil {
						logHistoryManager.WithError(err).Error(
							"Could't get main balance of output with address %s", output.GetAddress().String())
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
					if input.GetSpentOutput().GetAddress().String() == addr {
						return true
					}
				}
				return false
			}
			for _, output := range txn.GetOutputs() {
				if !isInput(output.GetAddress().String()) {
					coins, err := output.GetCoins(mainAsset)
					if err != nil {
						logHistoryManager.WithError(err).Error(
							"Could't get main balance of output with address %s", output.GetAddress().String())
					}
					amount += coins
				}
			}
		}
		return util.FormatCoins(amount, accuracy)
	}

	for e := range txnsList {
		txnDetail, err := transactions.NewTransactionDetailFromCoreTransaction(txnsList[e], getTxnType(txnsList[e]))
		if err != nil {
			logHistoryManager.WithError(err).Errorf(
				"Could't get a Transaction Details from the transaction: %s", txnsList[e].GetId())
		}
		txnDetail.SetAmount(getAmountFromTxn(txnsList[e], getTxnType(txnsList[e])))

		txnsDetailsList = append(txnsDetailsList, txnDetail)
	}

	sort.Sort(ByDate(txnsDetailsList))
	return txnsDetailsList
}
func (historyMang *HistoryManager) loadHistoryWithFilters() []*transactions.TransactionDetails {
	logHistoryManager.Info("Loading history with some filters")
	filterAddresses := historyMang.filters
	return historyMang.getTransactionsOfAddresses(filterAddresses)

}

func (historyMang *HistoryManager) loadHistory() []*transactions.TransactionDetails {
	logHistoryManager.Info("Loading history")
	addresses := historyMang.getAddressesWithWallets()

	filterAddresses := make([]string, 0)
	for addr := range addresses {
		filterAddresses = append(filterAddresses, addr)
	}

	return historyMang.getTransactionsOfAddresses(filterAddresses)

}

func (historyMang *HistoryManager) addFilter(addr string) {
	logHistoryManager.Info("Add filter")
	alreadyIs := false
	for _, filter := range historyMang.filters {
		if filter == addr {
			alreadyIs = true
			break
		}
	}
	if !alreadyIs {
		historyMang.filters = append(historyMang.filters, addr)
	}

}

func (historyMang *HistoryManager) removeFilter(addr string) {
	logHistoryManager.Info("Remove filter")

	for i := 0; i < len(historyMang.filters); i++ {
		if historyMang.filters[i] == addr {
			historyMang.filters = append(historyMang.filters[0:i], historyMang.filters[i+1:]...)
			break
		}
	}

}
func (historyMang *HistoryManager) getAddressesWithWallets() map[string]string {
	logHistoryManager.Info("Get Addresses with wallets")
	response := make(map[string]string, 0)
	it := historyMang.walletEnv.GetWalletSet().ListWallets()
	if it == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't load addresses")
		return response
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
