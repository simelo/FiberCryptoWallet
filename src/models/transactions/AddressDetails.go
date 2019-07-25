package transactions

import (
	"github.com/therecipe/qt/core"
)

//AddressDetails duplicated!!!
type AddressDetails struct {
	core.QObject

	_ string  `property:"address"`
	_ float32 `property:"addressSky"`
	_ int     `property:"addressCoinHours"`
}
