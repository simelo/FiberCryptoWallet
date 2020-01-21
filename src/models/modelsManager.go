package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	"github.com/therecipe/qt/qml"

	qtcore "github.com/therecipe/qt/core"
)

type ModelManager struct {
	qtcore.QObject
	wltManager     WalletManager
	addressesModel map[string]*address.AddressList
	_              func()                            `constructor:"init"`
	_              func(*WalletManager)              `slot:"setWalletManager"`
	_              func(string) *address.AddressList `slot:"getAddressModel"`
}

func (mm *ModelManager) init() {
	mm.ConnectSetWalletManager(mm.setWalletManager)
	mm.ConnectGetAddressModel(mm.getAddressModel)
	qml.QQmlEngine_SetObjectOwnership(mm, qml.QQmlEngine__CppOwnership)
	mm.addressesModel = make(map[string]*address.AddressList, 0)
}

func (mm *ModelManager) setWalletManager(wm *WalletManager) {
	mm.wltManager = *wm
}

func (mm *ModelManager) getAddressModel(wltName string) *address.AddressList {
	addrModel, ok := mm.addressesModel[wltName]
	if !ok {
		addrModel = address.NewAddressList(nil)
		qml.QQmlEngine_SetObjectOwnership(addrModel, qml.QQmlEngine__CppOwnership)
		addrModel.SetAddresses(mm.wltManager.getAddresses(wltName))
		if len(addrModel.Addresses()) == 0 {
			addr := address.NewAddressDetails(nil)
			addr.SetAddress("--------------------------")
			addr.SetCoinOptions(nil)
			addrModel.SetAddresses([]*address.AddressDetails{addr})
		}
		addrModel.RemoveAddress(0)
		mm.addressesModel[wltName] = addrModel
	}

	return addrModel
}
