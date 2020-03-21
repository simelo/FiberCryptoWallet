package transactions

import (
	"context"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sort"
)

func init() {
}

type TransactionList struct {
	core.QAbstractListModel
	Ctx     context.Context
	cancel  context.CancelFunc
	txnFind map[string]struct{}
	_       map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`
	_ func() `destructor:"destroyer"`

	_ func(transaction *TransactionDetails) `slot:"addTransaction,auto"`
	_ func(index int)                       `slot:"removeTransaction,auto"`
	_ func(txns []*TransactionDetails)      `slot:"addMultipleTransactions"`
	_ func()                                `slot:"clear"`

	_ []*TransactionDetails `property:"transactions"`
}

func (txnList *TransactionList) init() {
	txnList.SetRoles(map[int]*core.QByteArray{
		Date:          core.NewQByteArray2("date", -1),
		Status:        core.NewQByteArray2("status", -1),
		Type:          core.NewQByteArray2("type", -1),
		Amount:        core.NewQByteArray2("amount", -1),
		TransactionID: core.NewQByteArray2("transactionID", -1),
		BlockHeight:   core.NewQByteArray2("blockHeight", -1),
		Addresses:     core.NewQByteArray2("addresses", -1),
		Inputs:        core.NewQByteArray2("inputs", -1),
		Outputs:       core.NewQByteArray2("outputs", -1),
		CoinOptions:   core.NewQByteArray2("coinOptions", -1),
	})

	txnList.ConnectDestroyTransactionList(txnList.destroy)
	txnList.ConnectRowCount(txnList.rowCount)
	txnList.ConnectData(txnList.data)
	txnList.ConnectRoleNames(txnList.roleNames)
	txnList.ConnectAddMultipleTransactions(txnList.addMultipleTransactions)
	txnList.ConnectClear(txnList.clear)
	txnList.Ctx, txnList.cancel = context.WithCancel(context.Background())
	txnList.ConnectAddTransaction(txnList.addTransaction)
	txnList.ConnectRemoveTransaction(txnList.removeTransaction)
	txnList.txnFind = make(map[string]struct{})
}

func (txnList *TransactionList) rowCount(*core.QModelIndex) int {
	return len(txnList.Transactions())
}

func (txnList *TransactionList) roleNames() map[int]*core.QByteArray {
	return txnList.Roles()
}

func (txnList *TransactionList) addTransaction(transaction *TransactionDetails) {
	logTransactionDetails.Info("Adding transaction for history")
	txnList.BeginInsertRows(core.NewQModelIndex(), len(txnList.Transactions()), len(txnList.Transactions()))
	qml.QQmlEngine_SetObjectOwnership(transaction, qml.QQmlEngine__CppOwnership)
	txnList.SetTransactions(append(txnList.Transactions(), transaction))
	txnList.EndInsertRows()
}

func (txnList *TransactionList) removeTransaction(index int) {
	logTransactionDetails.Info("Removing transaction for history")
	txnList.BeginRemoveRows(core.NewQModelIndex(), index, index)
	txnList.SetTransactions(append(txnList.Transactions()[:index], txnList.Transactions()[index+1:]...))
	txnList.EndRemoveRows()
}

func (txnList *TransactionList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(txnList.Transactions()) {
		return core.NewQVariant()
	}

	transaction := txnList.Transactions()[index.Row()]

	switch role {
	case Date:
		{
			return core.NewQVariant1(transaction.Date())
		}
	case Status:
		{
			return core.NewQVariant1(transaction.Status())
		}
	case Type:
		{
			return core.NewQVariant1(transaction.Type())
		}
	case Amount:
		{
			return core.NewQVariant1(transaction.Amount())
		}
	case TransactionID:
		{
			return core.NewQVariant1(transaction.TransactionID())
		}
	case BlockHeight:
		{
			return core.NewQVariant1(transaction.BlockHeight())
		}
	case Addresses:
		{
			return core.NewQVariant1(transaction.Addresses())
		}

	case Inputs:
		{
			return core.NewQVariant1(transaction.Inputs())
		}
	case Outputs:
		{
			return core.NewQVariant1(transaction.Outputs())
		}
	case CoinOptions:
		{
			return core.NewQVariant1(transaction.CoinOptions())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (txnList *TransactionList) addMultipleTransactions(txns []*TransactionDetails) {
	logTransactionDetails.Info("adding multiple transactions")

	var newTxnList = make(map[string]struct{})
	for _, txn := range txns {
		newTxnList[txn.TransactionID()] = struct{}{}
		if _, ok := txnList.txnFind[txn.TransactionID()]; !ok {
			txnList.addTransaction(txn)
			txnList.txnFind[txn.TransactionID()] = struct{}{}
		}
	}
	var posToRemove = make([]int, 0)
	for index, qTxn := range txnList.Transactions() {
		if _, ok := newTxnList[qTxn.TransactionID()]; !ok {
			posToRemove = append(posToRemove, index)
			delete(txnList.txnFind, qTxn.TransactionID())
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(posToRemove)))
	for e := range posToRemove {
		txnList.removeTransaction(posToRemove[e])
	}
}

func (txnList *TransactionList) clear() {
	txnList.BeginRemoveRows(core.NewQModelIndex(), 0, len(txnList.Transactions())-1)
	txnList.SetTransactions(make([]*TransactionDetails, 0))
	txnList.EndRemoveRows()
}

func (txnList *TransactionList) destroy() {
	txnList.cancel()
}
