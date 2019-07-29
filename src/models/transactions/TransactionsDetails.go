package transactions

import (
	"github.com/simelo/FiberCryptoWallet/src/models"
	"github.com/therecipe/qt/core"
)

// TransactionDetails -
type TransactionDetails struct {
	core.QObject

	_ *core.QDateTime      `property:"date"`
	_ int                  `property:"status"`
	_ int                  `property:"type"`
	_ int                  `property:"amount"`
	_ int                  `property:"hoursReceived"`
	_ int                  `property:"hoursBurned"`
	_ string               `property:"transactionID"`
	_ string               `property:"sentAddress"`
	_ string               `property:"receivedAddress"`
	_ *models.AddressModel `property:"inputs"`
	_ *models.AddressModel `property:"outputs"`
}
