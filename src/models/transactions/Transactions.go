package transactions

import "github.com/therecipe/qt/core"

const (
	Date = int(core.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursReceived
	HoursBurned
	TransactionID
	SentAddress
	ReceivedAddress
)

const (
	transactionStatusConfirmed = iota
	transactionStatusPending
	transactionStatusPreview
)

const (
	transactionTypeSend = iota
	transactionTypeReceive
)

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	AddressSky
	AddressCoinHours
)

func init() {
	TransactionDetails_QmlRegisterType2("TransactionModels", 1, 0, "TransactionDetails")
	// AddressDetails_QmlRegisterType2("TransactionModels", 1, 0, "AddressDetails")
	AddressList_QmlRegisterType2("TransactionModels", 1, 0, "AddressList")
}
