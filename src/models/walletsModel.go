package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

const (
	Name = int(core.Qt__UserRole) + iota + 1
	EncryptionEnabled
	FileName
	Expand
	Addresses
	Currency
	CoinOptions
)

var logWalletsModel = logging.MustGetLogger("Wallets Model")

type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet               `property:"wallets"`

	_ func(*QWallet)                                                                   `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, address *address.AddressList) `slot:"editWallet"`
	_ func(row int)                                                                    `slot:"removeWallet"`
	_ func([]*QWallet)                                                                 `slot:"loadModel"`
	_ func([]*QWallet)                                                                 `slot:"updateModel"`
	_ func([]*QWallet, string)                                                         `slot:"filterWalletByCurrency"`
	_ int                                                                              `property:"count"`
}

type QWallet struct {
	core.QObject
	_ string               `property:"name"`
	_ int                  `property:"encryptionEnabled"`
	_ string               `property:"sky"`
	_ string               `property:"fileName"`
	_ string               `property:"currency"`
	_ *address.AddressList `property:"addresses"`
	_ *modelUtil.Map       `property:"coinOptions"`
	_ bool                 `property:"expand"`
}

func (walletModel *WalletModel) init() {
	logWalletsModel.Info("Initialize Wallet model")
	walletModel.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
		Expand:            core.NewQByteArray2("expand", -1),
		Addresses:         core.NewQByteArray2("addresses", -1),
		Currency:          core.NewQByteArray2("currency", -1),
		CoinOptions:       core.NewQByteArray2("coinOpts", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(walletModel, qml.QQmlEngine__CppOwnership)
	walletModel.ConnectData(walletModel.data)
	walletModel.ConnectSetData(walletModel.setData)
	walletModel.ConnectRowCount(walletModel.rowCount)
	walletModel.ConnectColumnCount(walletModel.columnCount)
	walletModel.ConnectRoleNames(walletModel.roleNames)

	walletModel.ConnectAddWallet(walletModel.addWallet)
	walletModel.ConnectFilterWalletByCurrency(walletModel.filterWalletByCurrency)
	walletModel.ConnectEditWallet(walletModel.editWallet)
	walletModel.ConnectRemoveWallet(walletModel.removeWallet)
	walletModel.ConnectLoadModel(walletModel.loadModel)
	walletModel.ConnectUpdateModel(walletModel.updateModel)
}

func (walletModel *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	logWalletsModel.Info("Loading data for index")

	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(walletModel.Wallets()) {
		return core.NewQVariant()
	}

	var w = walletModel.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant1(w.Name())
		}

	case EncryptionEnabled:
		{
			return core.NewQVariant1(w.EncryptionEnabled())
		}

	case FileName:
		{
			return core.NewQVariant1(w.FileName())
		}
	case Expand:
		{
			return core.NewQVariant1(w.IsExpand())
		}
	case Addresses:
		{
			return core.NewQVariant1(w.Addresses())
		}
	case Currency:
		{
			return core.NewQVariant1(w.Currency())
		}
	case CoinOptions:
		{
			return core.NewQVariant1(w.CoinOptions())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (walletModel *WalletModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {

	if !index.IsValid() {
		return false
	}

	if index.Row() >= len(walletModel.Wallets()) {
		return false
	}

	var w = walletModel.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			w.SetName(value.ToString())
		}
	case EncryptionEnabled:
		{
			w.SetEncryptionEnabled(value.ToInt(nil))
		}
	case FileName:
		{
			w.SetFileName(value.ToString())
		}
	case Expand:
		{
			w.SetExpand(value.ToBool())
		}
	default:
		{
			return false
		}
	}

	walletModel.DataChanged(index, index, []int{role})
	return true
}

func (walletModel *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(walletModel.Wallets())
}

func (walletModel *WalletModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (walletModel *WalletModel) roleNames() map[int]*core.QByteArray {
	return walletModel.Roles()
}

func (walletModel *WalletModel) addWallet(w *QWallet) {
	logWalletsModel.Info("Add Wallet")
	if w.Pointer() == nil {
		return
	}
	walletModel.BeginInsertRows(core.NewQModelIndex(), len(walletModel.Wallets()), len(walletModel.Wallets()))
	qml.QQmlEngine_SetObjectOwnership(w, qml.QQmlEngine__CppOwnership)
	walletModel.SetWallets(append(walletModel.Wallets(), w))
	walletModel.SetCount(walletModel.Count() + 1)
	walletModel.EndInsertRows()
}

func (walletModel *WalletModel) editWallet(row int, name string, encrypted bool, address *address.AddressList) {
	logWalletsModel.Info("Edit Wallet")
	pIndex := walletModel.Index(row, 0, core.NewQModelIndex())
	walletModel.setData(pIndex, core.NewQVariant1(name), Name)
	w := walletModel.Wallets()[row]
	w.SetAddresses(address)
	walletModel.DataChanged(pIndex, pIndex, []int{Addresses})

	if encrypted {
		walletModel.setData(pIndex, core.NewQVariant1(1), EncryptionEnabled)
	} else {
		walletModel.setData(pIndex, core.NewQVariant1(0), EncryptionEnabled)
	}
}

func (walletModel *WalletModel) removeWallet(row int) {
	logWalletsModel.Info("Remove wallets for index")
	walletModel.BeginRemoveRows(core.NewQModelIndex(), row, row)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], walletModel.Wallets()[row+1:]...))
	walletModel.SetCount(walletModel.Count() - 1)
	walletModel.EndRemoveRows()
}

func (walletModel *WalletModel) updateModel(wallets []*QWallet) {
	go func() {
		for i, wlt := range wallets {
			if len(wlt.Addresses().Addresses()) != len(walletModel.Wallets()[i].Addresses().Addresses()) ||
				wlt.Name() != walletModel.Wallets()[i].Name() ||
				wlt.EncryptionEnabled() != walletModel.Wallets()[i].EncryptionEnabled() {
				walletModel.editWallet(i, wlt.Name(), wlt.EncryptionEnabled() == 1, wlt.Addresses())
			}
		}
	}()
}

func (walletModel *WalletModel) loadModel(wallets []*QWallet) {
	logWalletsModel.Info("Loading wallets")

	for _, wlt := range wallets {
		qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
	}
	walletModel.BeginResetModel()
	walletModel.SetWallets(wallets)

	walletModel.EndResetModel()
	walletModel.SetCount(len(walletModel.Wallets()))
}

func (walletModel *WalletModel) filterWalletByCurrency(wallets []*QWallet, currency string) {
	logWalletsModel.Info("Loading wallets")
	qWalletList := make([]*QWallet, 0)
	for _, wlt := range wallets {
		if wlt.Currency() == currency {
			qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
			qWalletList = append(qWalletList, wlt)
		}
	}
	walletModel.BeginResetModel()
	walletModel.SetWallets(qWalletList)

	walletModel.EndResetModel()
	walletModel.SetCount(len(walletModel.Wallets()))
}
