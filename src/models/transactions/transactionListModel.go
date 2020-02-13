package transactions

import (
	"github.com/therecipe/qt/core"
)

func init() {
}

type TransactionList struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`

	_ func(transaction *TransactionDetails) `signal:"addTransaction,auto"`
	_ func(index int)                       `signal:"removeTransaction,auto"`
	_ func(txns []*TransactionDetails)      `slot:"addMultipleTransactions"`
	_ func()                                `slot:"clear"`

	_ []*TransactionDetails `property:"transactions"`
}

func (txnList *TransactionList) init() {
	logTransactionDetails.Info("Initialize Transaction list for History")
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

	txnList.ConnectRowCount(txnList.rowCount)
	txnList.ConnectData(txnList.data)
	txnList.ConnectRoleNames(txnList.roleNames)
	txnList.ConnectAddMultipleTransactions(txnList.addMultipleTransactions)
	txnList.ConnectClear(txnList.clear)

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

	for _, txn := range txns {
		txnList.addTransaction(txn)
	}
}

func (txnList *TransactionList) clear() {
	txnList.BeginRemoveRows(core.NewQModelIndex(), 0, len(txnList.Transactions())-1)
	txnList.SetTransactions(make([]*TransactionDetails, 0))
	txnList.EndRemoveRows()
}
