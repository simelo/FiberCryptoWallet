package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/util"
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
	_ func() []string              `slot:"getSupportedCurrencies"`
}

func (mu *ModelsUtil) init() {
	mu.ConnectGetTicketOfCurrency(mu.getTicketOfCurrency)
	mu.ConnectGetSupportedCurrencies(mu.getSupportedCurrencies)
}

func (mu *ModelsUtil) getTicketOfCurrency(currency string) string {
	logModelUtils.Info("Getting ticket for currency: ", currency)
	ticket, err := util.GetTicketOfCurrncies(currency)
	if err != nil {
		logModelUtils.WithError(err)
		return ""
	}
	return ticket
}

func (mu *ModelsUtil) getSupportedCurrencies() []string {
	return util.LoadAllExistentCurrencies()
}
