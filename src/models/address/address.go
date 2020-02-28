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
	CoinOptions
	WalletId
)

type QAddress struct {
	core.QObject
	_ string    `property:"address"`
	_ *util.Map `property:"coinOptions"`
	_ string    `property:"walletId"`
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

func (modelAddrs *ModelAddress) init() {
	modelAddrs.SetRoles(map[int]*core.QByteArray{
		Address:     core.NewQByteArray2("address", -1),
		CoinOptions: core.NewQByteArray2("coinOptions", -1),
		WalletId:    core.NewQByteArray2("walletId", -1),
	})

	modelAddrs.ConnectRowCount(modelAddrs.rowCount)
	modelAddrs.ConnectData(modelAddrs.data)
	modelAddrs.ConnectRoleNames(modelAddrs.roleNames)
	modelAddrs.ConnectLoadModel(modelAddrs.loadModel)
	modelAddrs.ConnectAddAddress(modelAddrs.addAddress)

}

func (modelAddrs *ModelAddress) rowCount(*core.QModelIndex) int {
	return len(modelAddrs.Addresses())
}

func (modelAddrs *ModelAddress) roleNames() map[int]*core.QByteArray {
	return modelAddrs.Roles()
}

func (modelAddrs *ModelAddress) addAddress(address *QAddress) {
	logAddressModel.Info("Adding Address")
	modelAddrs.BeginInsertRows(core.NewQModelIndex(), len(modelAddrs.Addresses()), len(modelAddrs.Addresses()))
	modelAddrs.SetAddresses(append(modelAddrs.Addresses(), address))
	modelAddrs.EndInsertRows()
}

func (modelAddrs *ModelAddress) editAddress() {

}

func (modelAddrs *ModelAddress) removeAddress(index int) {
	logAddressModel.Info("Removing Address")
	if index >= len(modelAddrs.Addresses()) {
		return
	}
	modelAddrs.BeginRemoveRows(core.NewQModelIndex(), index, index)
	modelAddrs.SetAddresses(append(modelAddrs.Addresses()[:index], modelAddrs.Addresses()[index+1:]...))
	modelAddrs.EndRemoveRows()
}

func (modelAddrs *ModelAddress) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(modelAddrs.Addresses()) {
		return core.NewQVariant()
	}

	address := modelAddrs.Addresses()[index.Row()]

	switch role {
	case Address:
		{
			return core.NewQVariant1(address.Address())
		}
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

func (modelAddrs *ModelAddress) loadModel(Qaddresses []*QAddress) {
	for _, addr := range Qaddresses {
		qml.QQmlEngine_SetObjectOwnership(addr, qml.QQmlEngine__CppOwnership)
	}

	addresses := make([]*QAddress, 0)
	address := NewQAddress(nil)
	address.SetAddress("--------------------------")

	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	addresses = append(addresses, address)
	addresses = append(addresses, Qaddresses...)

	modelAddrs.BeginResetModel()
	modelAddrs.SetAddresses(addresses)
	modelAddrs.EndResetModel()

}

func CompareModelAddress(modelAddrs1, modelAddrs2 *ModelAddress) bool {
	if len(modelAddrs1.Addresses()) != len(modelAddrs2.Addresses()) {
		return false
	}

	for i := 0; i < len(modelAddrs1.Addresses()); i++ {
		if !CompareQAddress(modelAddrs1.Addresses()[i], modelAddrs2.Addresses()[i]) {
			return false
		}
	}
	return true
}

func CompareQAddress(a, b *QAddress) bool {
	return a.Address() == b.Address() && util.CompareMaps(a.CoinOptions(), b.CoinOptions()) &&
		a.WalletId() == b.WalletId()

}
