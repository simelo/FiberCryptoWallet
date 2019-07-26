package transactions

import "github.com/therecipe/qt/core"

//Wallet duplicate!!!
type Wallet struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"encryptionEnabled"`
	_ int    `property:"sky"`
	_ int    `property:"coinHours"`
	_ string `property:"fileName"`
}
