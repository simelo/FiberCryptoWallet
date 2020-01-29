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
}

const (
	Date = int(qtCore.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursTraspassed
	HoursBurned
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
	_   *qtCore.QDateTime    `property:"date"`
	_   int                  `property:"status"`
	_   int                  `property:"type"`
	_   uint64               `property:"blockHeight"`
	_   string               `property:"amount"`
	_   string               `property:"hoursTraspassed"`
	_   string               `property:"hoursBurned"`
	_   string               `property:"transactionID"`
	_   *address.AddressList `property:"addresses"`
	_   *address.AddressList `property:"inputs"`
	_   *address.AddressList `property:"outputs"`
	_   *modelUtil.Map       `property:"coinOptions"`
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

	addresses := address.NewAddressList(nil)
	inputList := address.NewAddressList(nil)
	outputsList := address.NewAddressList(nil)

	for _, input := range transaction.GetInputs() {

		qIn := address.NewAddressDetails(nil)
		qIn.SetAddress(input.GetSpentOutput().GetAddress().String())
		inputCoinOptions := modelUtil.NewMap(nil)

		for _, v := range input.GetCoinTraits() {
			inputCoinOptions.SetValue(v.GetTrait(), v.GetValue())
		}

		qIn.SetCoinOptions(inputCoinOptions)
		inputList.AddAddress(qIn)
		addresses.AddAddress(qIn)
	}

	txnDetails.SetInputs(inputList)

	for _, out := range transaction.GetOutputs() {

		qOu := address.NewAddressDetails(nil)
		qOu.SetAddress(out.GetAddress().String())
		outputCoinOptions := modelUtil.NewMap(nil)

		for _, v := range out.GetCoinTraits() {
			outputCoinOptions.SetValue(v.GetTrait(), v.GetValue())
		}

		qOu.SetCoinOptions(outputCoinOptions)

		outputsList.AddAddress(qOu)
		addresses.AddAddress(qOu)
	}

	txnCoinOptions := modelUtil.NewMap(nil)
	for _, v := range transaction.GetCoinTraits() {
		txnCoinOptions.SetValue(v.GetTrait(), v.GetValue())
	}

	txnDetails.SetCoinOptions(txnCoinOptions)
	txnDetails.SetOutputs(outputsList)
	txnDetails.SetAddresses(addresses)

	return txnDetails, nil
}
