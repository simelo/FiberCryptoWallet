package transactions

import (
<<<<<<< HEAD
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	qtcore "github.com/therecipe/qt/core"
=======
	"strconv"
	"time"

	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
>>>>>>> c6fca981... [eth][config] Add config package for register ethereum plugin settings
)

func init() {
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
}

const (
	Date = int(qtcore.Qt__UserRole) + 1<<iota
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
)

type TransactionDetails struct {
	qtcore.QObject
	_ *qtcore.QDateTime    `property:"date"`
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
