package transactions

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/models/addresses"
	"github.com/therecipe/qt/core"
)

// TransactionModel -
type TransactionModel struct {
	core.QObject

	_ *core.QDateTime           `property:"date"`
	_ int                       `property:"status"`
	_ int                       `property:"type"`
	_ int                       `property:"amount"`
	_ int                       `property:"hoursReceived"`
	_ int                       `property:"hoursBurned"`
	_ string                    `property:"transactionID"`
	_ string                    `property:"sentAddress"`
	_ string                    `property:"receivedAddress"`
	_ *addresses.AddressesModel `property:"inputs"`
	_ *addresses.AddressesModel `property:"outputs"`
}
