package pending

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models" // callable as skycoin
	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/models/transactions"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
)

var logPendingTxn = logging.MustGetLogger("modelsPendingTransaction")

type PendingTransactionList struct {
	qtCore.QObject
	PEX       core.PEX
	WalletEnv core.WalletEnv

	_ func() `constructor:"init"`

	_ []*transactions.TransactionDetails            `property:"transactions"`
	_ bool                                          `property:"loading"`
	_ func(bool) []*transactions.TransactionDetails `slot:"recoverTransactions"`
	_ func()                                        `slot:"getAll"`
	_ func()                                        `slot:"cleanPendingTxns"`
	_ func()                                        `slot:"getMine"`
}

func (model *PendingTransactionList) init() {
	model.ConnectGetAll(model.getAll)
	model.ConnectGetMine(model.getMine)
	model.ConnectRecoverTransactions(model.recoverTransactions)
	model.SetLoading(true)
	model.ConnectCleanPendingTxns(model.cleanPendingTxns)

	altManager := local.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}
	model.PEX = &skycoin.SkycoinPEX{}
	model.WalletEnv = walletsEnvs[0]

}

func (model *PendingTransactionList) cleanPendingTxns() {
	model.SetTransactions(make([]*transactions.TransactionDetails, 0))
}

func (model *PendingTransactionList) recoverTransactions(mine bool) []*transactions.TransactionDetails {
	model.SetLoading(true)
	if mine {
		model.getMine()
	} else {
		model.getAll()
	}
	return model.Transactions()
}

func (model *PendingTransactionList) getAll() {
	logPendingTxn.Info("Getting txn details")

	txns, err := model.PEX.GetTxnPool()
	if err != nil {
		logPendingTxn.WithError(err).Warn("Couldn't get txn pool")
		return
	}
	ptModels := make([]*transactions.TransactionDetails, 0)
	for txns.Next() {
		ptModel, err := transactions.NewTransactionDetailFromCoreTransaction(txns.Value(),
			transactions.TransactionTypeGeneric)

		if err != nil {
			logPendingTxn.WithError(err).Warn("Couldn't get txn pool")
			return
		}

		ptModels = append(ptModels, ptModel)
	}
	model.SetLoading(false)
	model.SetTransactions(ptModels)

}

func (model *PendingTransactionList) getMine() {
	logPendingTxn.Info("Getting txn details")

	wallets := model.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		logPendingTxn.WithError(nil).Warn("Couldn't load list of wallets")
		return
	}
	ptModels := make([]*transactions.TransactionDetails, 0)
	for wallets.Next() {
		wlt := wallets.Value()
		account := wlt.GetCryptoAccount()
		txns, err := account.ListPendingTransactions()
		if err != nil {
			// display an error in qml app when Mine is selected
			logPendingTxn.WithError(err).Warn("Couldn't list pending transactions")
			continue
		}
		for txns.Next() {
			txn := txns.Value()
			if txn.GetStatus() == core.TXN_STATUS_PENDING {
				ptModel, err := transactions.NewTransactionDetailFromCoreTransaction(txn,
					transactions.TransactionTypeGeneric)

				if err != nil {
					logPendingTxn.WithError(err).Warn("Couldn't list pending transactions")
				}
				ptModels = append(ptModels, ptModel)
			}
		}
	}
	model.SetLoading(false)
	model.SetTransactions(ptModels)
}
