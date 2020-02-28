package transactions

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"time"
)

var logTransactionDetails = logging.MustGetLogger("TransactionDetails")

func init() {
	TransactionDetails_QmlRegisterType2("Transaction", 1, 0, "QTransactionDetail")
	TransactionList_QmlRegisterType2("Transaction", 1, 0, "QTransactionList")

}

const (
	Date = int(qtCore.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	TransactionID
	BlockHeight
	Addresses
	Inputs
	Outputs
	CoinOptions
)

const (
	TransactionStatusConfirmed = iota
	TransactionStatusPending
	TransactionStatusPreview
)

const (
	TransactionTypeSend = iota
	TransactionTypeReceive
	TransactionTypeInternal
	TransactionTypeGeneric
)

type TransactionDetails struct {
	qtCore.QObject
	Txn core.Transaction
	_   *qtCore.QDateTime     `property:"date"`
	_   int                   `property:"status"`
	_   int                   `property:"type"`
	_   uint64                `property:"blockHeight"`
	_   string                `property:"amount"`
	_   string                `property:"transactionID"`
	_   *address.ModelAddress `property:"addresses"`
	_   *address.ModelAddress `property:"inputs"`
	_   *address.ModelAddress `property:"outputs"`
	_   *modelUtil.Map        `property:"coinOptions"`
}

func NewTransactionDetailFromCoreTransaction(transaction core.Transaction, txType int) (*TransactionDetails, error) {
	txnDetails := NewTransactionDetails(nil)
	t := time.Unix(int64(transaction.GetTimestamp()), 0)
	txnDetails.Txn = transaction
	txnDetails.SetDate(qtCore.NewQDateTime3(qtCore.NewQDate3(t.Year(), int(t.Month()), t.Day()),
		qtCore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtCore.Qt__LocalTime))

	switch transaction.GetStatus() {
	case core.TXN_STATUS_CONFIRMED:
		txnDetails.SetStatus(TransactionStatusConfirmed)
	case core.TXN_STATUS_PENDING:
		txnDetails.SetStatus(TransactionStatusPending)
	default:
		txnDetails.SetStatus(TransactionStatusPreview)
	}

	txnDetails.SetTransactionID(transaction.GetId())

	txnDetails.SetBlockHeight(transaction.GetBlockHeight())

	txnDetails.SetType(txType)

	addresses := address.NewModelAddress(nil)
	inputList := address.NewModelAddress(nil)
	outputsList := address.NewModelAddress(nil)

	var containsAddress = func(addr string) bool {
		for _, addrDetail := range addresses.Addresses() {
			if addrDetail.Address() == addr {
				return true
			}
		}
		return false
	}

	for _, input := range transaction.GetInputs() {

		qIn := address.NewQAddress(nil)
		unspentOutput, err := input.GetSpentOutput()
		if err != nil {
			logTransactionDetails.WithError(err).Errorf(
				"Couldn't get unspentOutput from input %s in transaction %s", input.GetId(), transaction.GetId())
			continue
		}
		inputAddrs, err := unspentOutput.GetAddress()
		if err != nil {
			logTransactionDetails.WithError(err).Errorf(
				"Couldn't get address from input %s in transaction %s", input.GetId(), transaction.GetId())
			continue
		}

		qIn.SetAddress(inputAddrs.String())
		inputCoinOptions := modelUtil.NewMap(nil)

		for _, v := range input.GetCoinTraits() {
			inputCoinOptions.SetValue(v.GetTrait(), v.GetValue())
		}

		qIn.SetCoinOptions(inputCoinOptions)
		inputList.AddAddress(qIn)

		if !containsAddress(qIn.Address()) {
			addresses.AddAddress(qIn)
		}
	}

	txnDetails.SetInputs(inputList)

	for _, out := range transaction.GetOutputs() {

		qOu := address.NewQAddress(nil)
		outAddrs, err := out.GetAddress()
		if err != nil {
			logTransactionDetails.WithError(err).Errorf("Couldn't get output address for %s in transaction %s",
				out.GetId(), transaction.GetId())
			continue
		}
		qOu.SetAddress(outAddrs.String())
		outputCoinOptions := modelUtil.NewMap(nil)

		for _, v := range out.GetCoinTraits() {
			outputCoinOptions.SetValue(v.GetTrait(), v.GetValue())
		}

		qOu.SetCoinOptions(outputCoinOptions)

		outputsList.AddAddress(qOu)

		if !containsAddress(qOu.Address()) {
			addresses.AddAddress(qOu)
		}
	}

	txnCoinOptions := modelUtil.NewMap(nil)
	for _, v := range transaction.GetCoinTraits() {
		txnCoinOptions.SetValue(v.GetTrait(), v.GetValue())
	}

	if len(txnCoinOptions.GetKeys()) > 0 {
		txnDetails.SetAmount(txnCoinOptions.GetValue(txnCoinOptions.GetKeys()[0]))
	} else {
		txnDetails.SetAmount("")
	}

	txnDetails.SetCoinOptions(txnCoinOptions)
	txnDetails.SetOutputs(outputsList)
	txnDetails.SetAddresses(addresses)

	return txnDetails, nil
}
