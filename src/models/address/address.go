package address

import (
	"github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

var logAddressModel = logging.MustGetLogger("Address Model")

func init() {
	ModelAddress_QmlRegisterType2("Address", 1, 0, "QAddressList")
	QAddress_QmlRegisterType2("Address", 1, 0, "QAddress")
}

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	// AddressSky
	// AddressCoinHours
	CoinOptions
	WalletId
)

type QAddress struct {
	core.QObject
	_ string    `property:"address"`
	_ *util.Map `property:"coinOptions"`
	// _ string    `property:"addressSky"`
	// _ string    `property:"addressCoinHours"`
	_ string `property:"walletId"`
}

type ModelAddress struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`

	_ func(transaction *QAddress) `signal:"addAddress"`
	_ func(index int)             `signal:"removeAddress,auto"`
	_ func([]*QAddress)           `slot:"loadModel"`

	_ []*QAddress `property:"addresses"`
}

func (al *ModelAddress) init() {
	al.SetRoles(map[int]*core.QByteArray{
		Address: core.NewQByteArray2("address", -1),
		// AddressSky:       core.NewQByteArray2("addressSky", -1),
		// AddressCoinHours: core.NewQByteArray2("addressCoinHours", -1),
		CoinOptions: core.NewQByteArray2("coinOptions", -1),
		WalletId:    core.NewQByteArray2("walletId", -1),
	})

	al.ConnectRowCount(al.rowCount)
	al.ConnectData(al.data)
	al.ConnectRoleNames(al.roleNames)
	al.ConnectLoadModel(al.loadModel)
	al.ConnectAddAddress(al.addAddress)

}

func (al *ModelAddress) rowCount(*core.QModelIndex) int {
	return len(al.Addresses())
}

func (al *ModelAddress) roleNames() map[int]*core.QByteArray {
	return al.Roles()
}

func (al *ModelAddress) addAddress(address *QAddress) {
	logAddressModel.Info("Adding Address")
	al.BeginInsertRows(core.NewQModelIndex(), len(al.Addresses()), len(al.Addresses()))
	al.SetAddresses(append(al.Addresses(), address))
	al.EndInsertRows()
}

func (al *ModelAddress) removeAddress(index int) {
	logAddressModel.Info("Removing Address")
	if index >= len(al.Addresses()) {
		return
	}
	al.BeginRemoveRows(core.NewQModelIndex(), index, index)
	al.SetAddresses(append(al.Addresses()[:index], al.Addresses()[index+1:]...))
	al.EndRemoveRows()
}

func (al *ModelAddress) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(al.Addresses()) {
		return core.NewQVariant()
	}

	address := al.Addresses()[index.Row()]

	switch role {
	case Address:
		{
			return core.NewQVariant1(address.Address())
		}
	// case AddressCoinHours:
	// 	{
	// 		return core.NewQVariant1(address.AddressCoinHours())
	// 	}
	// case AddressSky:
	// 	{
	// 		return core.NewQVariant1(address.AddressSky())
	// 	}
	case CoinOptions:
		{
			return core.NewQVariant1(address.CoinOptions())
		}
	case WalletId:
		{
			return core.NewQVariant1(address.WalletId())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (al *ModelAddress) loadModel(Qaddresses []*QAddress) {
	for _, addr := range Qaddresses {
		qml.QQmlEngine_SetObjectOwnership(addr, qml.QQmlEngine__CppOwnership)
	}

	addresses := make([]*QAddress, 0)
	address := NewQAddress(nil)
	address.SetAddress("--------------------------")
	// address.SetAddressSky("0")
	// address.SetAddressCoinHours("0")
	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	addresses = append(addresses, address)
	addresses = append(addresses, Qaddresses...)

	al.BeginResetModel()
	al.SetAddresses(addresses)
	al.EndResetModel()

}

func CompareModelAddress(modelAddrs1, modelAddrs2 *ModelAddress) bool {
	// var isEqual = true
	if len(modelAddrs1.Addresses()) != len(modelAddrs2.Addresses()) {
		return false
	}

	for i := 0; i < len(modelAddrs1.Addresses()); i++ {
		if CompareQAddress(modelAddrs1.Addresses()[i], modelAddrs2.Addresses()[i]) {
			return false
		}
	}
	return true
}

func CompareQAddress(a, b *QAddress) bool {
	return a.Address() == b.Address() && util.CompareMaps(a.CoinOptions(), b.CoinOptions()) &&
		a.WalletId() == b.WalletId()
}
