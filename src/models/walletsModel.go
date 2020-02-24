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

	_             func(*QWallet)                 `slot:"addWallet"`
	_             func(row int, wallet *QWallet) `slot:"editWallet"`
	_             func(row int)                  `slot:"removeWallet"`
	_             func([]*QWallet)               `slot:"loadModelAsync"`
	_             func([]*QWallet)               `slot:"loadModel"`
	_             func([]*QWallet)               `slot:"updateModel"`
	_             func([]*QWallet, string)       `slot:"filterWalletByCurrency"`
	_             int                            `property:"count"`
	receivChannel chan *updateWalletInfo
	walletByName  map[string]*QWallet
}

type QWallet struct {
	core.QObject
	_ string                `property:"name"`
	_ int                   `property:"encryptionEnabled"`
	_ string                `property:"fileName"`
	_ string                `property:"currency"`
	_ *address.ModelAddress `property:"addresses"`
	_ *modelUtil.Map        `property:"coinOptions"`
}

func (walletModel *WalletModel) init() {
	logWalletsModel.Info("Initialize Wallet model")
	walletModel.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
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
	walletModel.ConnectLoadModelAsync(walletModel.loadModelAsync)
	walletModel.ConnectAddWallet(walletModel.addWallet)
	walletModel.ConnectFilterWalletByCurrency(walletModel.filterWalletByCurrency)
	walletModel.ConnectEditWallet(walletModel.editWallet)
	walletModel.ConnectRemoveWallet(walletModel.removeWallet)
	walletModel.ConnectLoadModel(walletModel.loadModel)
	walletModel.ConnectUpdateModel(walletModel.updateModel)
	walletModel.receivChannel = walletManager.suscribe()
	walletModel.walletByName = make(map[string]*QWallet, 0)
	// go func() {
	// 	for {
	// 		wi := <-walletModel.receivChannel
	// 		if wi.isNew {
	// 			walletModel.addWallet(wi.wallet)
	// 		} else {
	// 			encrypted := false
	// if wi.wallet.EncryptionEnabled() == 1 {
	// 	encrypted = true
	// }
	// walletModel.editWallet(wi.row, wi.wallet.Name(), encrypted, wi.wallet.Addresses())

	// }
	// }
	// }()

}

func (walletModel *WalletModel) changeExpanded(id string) {
	// w := walletModel.walletByName[id]
	// w.SetExpand(!w.IsExpand())

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

func (walletModel *WalletModel) loadModelAsync(qWalletList []*QWallet) {
	logWalletModel.Info("Load Wallets Async")
	getPos := func(list []*QWallet, obj *QWallet) int {
		for e := range list {
			if list[e].FileName() == obj.FileName() {
				return e
			}
		}
		return -1
	}
	logWalletModel.Info(len(qWalletList))
	for e := range qWalletList {
		if _, ok := walletModel.walletByName[qWalletList[e].FileName()]; !ok {
			walletModel.addWallet(qWalletList[e])
		} else {
			if !CompareWallet(qWalletList[e], walletModel.walletByName[qWalletList[e].FileName()]) {
				walletModel.editWallet(getPos(walletModel.Wallets(), qWalletList[e]), qWalletList[e])
			}
		}
	}
}

func (walletModel *WalletModel) addWallet(w *QWallet) {
	logWalletsModel.Info("Add Wallet")
	if w.Pointer() == nil {
		return
	}

	var row = len(walletModel.Wallets())
	for pos, wlt := range walletModel.Wallets() {
		if w.Name() < wlt.Name() {
			row = pos
			break
		}
	}

	walletModel.walletByName[w.FileName()] = w

	walletModel.BeginInsertRows(core.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(w, qml.QQmlEngine__CppOwnership)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], append([]*QWallet{w}, walletModel.Wallets()[row:]...)...))
	walletModel.SetCount(walletModel.Count() + 1)
	walletModel.EndInsertRows()
	logWalletModel.Info("End Add Wallet")
}

func (walletModel *WalletModel) editWallet(row int, wallet *QWallet) {
	logWalletsModel.Info("Edit Wallet")
	if row == -1 {
		return
	}
	walletModel.removeWallet(row)
	walletModel.addWallet(wallet)
	// lIndex := walletModel.Index(len(walletModel.Wallets())-1, 0, core.NewQModelIndex())
	// w := walletModel.Wallets()[row]
	// w.SetName(name)
	// w.SetAddresses(address)
	// walletModel.DataChanged(row, row, []int{Addresses})

	// if encrypted {
	// 	w.SetEncryptionEnabled(1)
	// } else {
	// 	w.SetEncryptionEnabled(0)
	// }
	// walletModel.DataChanged(pIndex, lIndex, []int{Name, EncryptionEnabled})
}

func (walletModel *WalletModel) removeWallet(row int) {
	logWalletsModel.Info("Remove wallets for index")
	walletModel.BeginRemoveRows(core.NewQModelIndex(), row, row)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], walletModel.Wallets()[row+1:]...))
	walletModel.SetCount(walletModel.Count() - 1)
	walletModel.EndRemoveRows()
}

func (walletModel *WalletModel) updateModel(wallets []*QWallet) {
	// for i, wlt := range wallets {
	// 	walletModel.editWallet(i, wlt.Name(), wlt.EncryptionEnabled() == 1, wlt.Addresses())
	// }
}

func (walletModel *WalletModel) loadModel(wallets []*QWallet) {
	logWalletsModel.Info("Loading wallets")

	for _, wlt := range wallets {
		qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
	}
	for _, w := range wallets {
		walletModel.walletByName[w.FileName()] = w
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

func CompareWallet(a, b *QWallet) bool {
	return address.CompareModelAddress(a.Addresses(), b.Addresses()) &&
		modelUtil.CompareMaps(a.CoinOptions(), b.CoinOptions()) &&
		a.EncryptionEnabled() == b.EncryptionEnabled() &&
		a.FileName() == b.FileName() &&
		a.Currency() == b.Currency() &&
		a.Name() == b.Name()
}
