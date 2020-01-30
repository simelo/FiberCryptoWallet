package util

import (
	"fmt"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
)

func init() {
	ModelsUtil_QmlRegisterType2("ModelUtils", 1, 0, "Utils")
}

var logModelUtils = logging.MustGetLogger("ModelUtils")

type ModelsUtil struct {
	qtCore.QObject
	_ func()                       `constructor:"init"`
	_ func(currency string) string `slot:"getTicketOfCurrency"`
}

func (mu *ModelsUtil) init() {
	mu.ConnectGetTicketOfCurrency(mu.getTicketOfCurrency)
}

func (mu *ModelsUtil) getTicketOfCurrency(currency string) string {
	logModelUtils.Info("Getting ticket for currency: ", currency)
	fmt.Println(currency)
	switch currency {
	case "skycoin":
		return "SKY"
	default:
		return ""
	}
}
