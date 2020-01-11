package transactions

import (
	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"strconv"
	"time"
)

var logTransactionDetails = logging.MustGetLogger("TransactionDetails")

func init() {
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
}

const (
	Date = int(qtCore.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursTraspassed
	HoursBurned
	TransactionID
	Addresses
	Inputs
	Outputs
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
	_ *qtCore.QDateTime    `property:"date"`
	_ int                  `property:"status"`
	_ int                  `property:"type"`
	_ string               `property:"amount"`
	_ string               `property:"hoursTraspassed"`
	_ string               `property:"hoursBurned"`
	_ string               `property:"transactionID"`
	_ *address.AddressList `property:"addresses"`
	_ *address.AddressList `property:"inputs"`
	_ *address.AddressList `property:"outputs"`
}

func NewTransactionDetailFromCoreTransaction(transaction core.Transaction, txType int) (*TransactionDetails, error) {

	txnDetails := NewTransactionDetails(nil)
	t := time.Unix(int64(transaction.GetTimestamp()), 0)

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

	txnDetails.SetType(txType)

	addresses := address.NewAddressList(nil)
	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)

	txnIns := transaction.GetInputs()
	var totalAmount float64
	for e := range txnIns {

		qIn := address.NewAddressDetails(nil)
		qIn.SetAddress(txnIns[e].GetSpentOutput().GetAddress().String())

		skyUint64, err := txnIns[e].GetCoins(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins balance")
			continue
		}
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins quotient")
			continue
		}
		skyFloat := float64(skyUint64) / float64(accuracy)
		qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
		totalAmount += skyFloat

		chUint64, err := txnIns[e].GetCoins(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours balance")
			continue
		}
		accuracy, err = util.AltcoinQuotient(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours quotient")
			continue
		}
		qIn.SetAddressCoinHours(util.FormatCoins(chUint64, accuracy))
		inputs.AddAddress(qIn)
		addresses.AddAddress(qIn)
	}
	accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
	if err != nil {
		logTransactionDetails.WithError(err).Warn("Couldn't get Sky accuracy")
	}

	txnDetails.SetAmount(util.FormatCoins(uint64(totalAmount), accuracy))
	txnDetails.SetInputs(inputs)
	var finalHours uint64
	for _, out := range transaction.GetOutputs() {
		sky, err := out.GetCoins(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins balance")
			continue
		}
		qOu := address.NewAddressDetails(nil)
		qOu.SetAddress(out.GetAddress().String())
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins quotient")
			continue
		}
		qOu.SetAddressSky(util.FormatCoins(sky, accuracy))
		val, err := out.GetCoins(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours balance")
			continue
		}
		accuracy, err = util.AltcoinQuotient(coin.CoinHour)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours quotient")
			continue
		}
		finalHours += val
		qOu.SetAddressCoinHours(util.FormatCoins(val, accuracy))
		outputs.AddAddress(qOu)
		addresses.AddAddress(qOu)
	}
	accuracy, err = util.AltcoinQuotient(coin.CoinHoursTicker)
	if err != nil {
		logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours accuracy")
	}
	txnDetails.SetHoursTraspassed(util.FormatCoins(finalHours, accuracy))

	txnDetails.SetAddresses(addresses)
	txnDetails.SetOutputs(outputs)

	return txnDetails, nil
}
