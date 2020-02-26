package wallets

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sync"
)

var cosaTest []core.CoinTrait

const (
	Name = int(qtCore.Qt__UserRole) + iota + 1
	EncryptionEnabled
	FileName
	Addresses
	Currency
	CoinOptions
)

func init() {
	WalletModel_QmlRegisterType2("QWallets", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("QWallets", 1, 0, "QWallet")
	cosaTest = append(cosaTest, util.NewCoinTrait("SKY", "12.0"), util.NewCoinTrait("SCH", "50.5"))
}

var logWalletsModel = logging.MustGetLogger("Wallets Model")

type WalletModel struct {
	qtCore.QAbstractListModel

	mutex sync.Mutex
	_     func()                     `constructor:"init"`
	_     map[int]*qtCore.QByteArray `property:"roles"`
	_     []*QWallet                 `property:"wallets"`

	_ func(*QWallet)                 `slot:"addWallet"`
	_ func(row int, wallet *QWallet) `slot:"editWallet"`
	_ func(row int)                  `slot:"removeWallet"`
	_ func([]*QWallet)               `slot:"loadModelAsync"`
	_ func([]*QWallet)               `slot:"loadModel"`
	_ func([]*QWallet)               `slot:"updateModel"`
	_ func([]*QWallet, string)       `slot:"filterWalletByCurrency"`
	_ int                            `property:"count"`
	// receivChannel chan *models.updateWalletInfo
	walletByName map[string]*QWallet
}

type QWallet struct {
	qtCore.QObject
	_ string                `property:"name"`
	_ int                   `property:"encryptionEnabled"`
	_ string                `property:"fileName"`
	_ string                `property:"currency"`
	_ *address.ModelAddress `property:"addresses"`
	_ *modelUtil.Map        `property:"coinOptions"`
}

func (walletModel *WalletModel) init() {
	logWalletsModel.Info("Initialize Wallet model")
	walletModel.SetRoles(map[int]*qtCore.QByteArray{
		Name:              qtCore.NewQByteArray2("name", -1),
		EncryptionEnabled: qtCore.NewQByteArray2("encryptionEnabled", -1),
		FileName:          qtCore.NewQByteArray2("fileName", -1),
		Addresses:         qtCore.NewQByteArray2("addresses", -1),
		Currency:          qtCore.NewQByteArray2("currency", -1),
		CoinOptions:       qtCore.NewQByteArray2("coinOpts", -1),
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
	// walletModel.receivChannel = models.walletManager.suscribe()
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

func (walletModel *WalletModel) data(index *qtCore.QModelIndex, role int) *qtCore.QVariant {
	logWalletsModel.Info("Loading data for index")

	if !index.IsValid() {
		return qtCore.NewQVariant()
	}

	if index.Row() >= len(walletModel.Wallets()) {
		return qtCore.NewQVariant()
	}

	var w = walletModel.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			return qtCore.NewQVariant1(w.Name())
		}

	case EncryptionEnabled:
		{
			return qtCore.NewQVariant1(w.EncryptionEnabled())
		}

	case FileName:
		{
			return qtCore.NewQVariant1(w.FileName())
		}
	case Addresses:
		{
			return qtCore.NewQVariant1(w.Addresses())
		}
	case Currency:
		{
			return qtCore.NewQVariant1(w.Currency())
		}
	case CoinOptions:
		{
			return qtCore.NewQVariant1(w.CoinOptions())
		}
	default:
		{
			return qtCore.NewQVariant()
		}
	}
}

func (walletModel *WalletModel) setData(index *qtCore.QModelIndex, value *qtCore.QVariant, role int) bool {

	return true
}

func (walletModel *WalletModel) rowCount(parent *qtCore.QModelIndex) int {
	return len(walletModel.Wallets())
}

func (walletModel *WalletModel) columnCount(parent *qtCore.QModelIndex) int {
	return 1
}

func (walletModel *WalletModel) roleNames() map[int]*qtCore.QByteArray {
	return walletModel.Roles()
}

func (walletModel *WalletModel) loadModelAsync(qWalletList []*QWallet) {
	logWalletsModel.Info("Load Wallets Async")
	getPos := func(list []*QWallet, obj *QWallet) int {
		for e := range list {
			if list[e].FileName() == obj.FileName() {
				return e
			}
		}
		return -1
	}

	logWalletsModel.Info(len(qWalletList))
	for e := range qWalletList {
		if _, ok := walletModel.walletByName[qWalletList[e].FileName()]; !ok {
			walletModel.addWallet(qWalletList[e])
		} else {
			// if !CompareWallet(qWalletList[e], walletModel.walletByName[qWalletList[e].FileName()]) {
			walletModel.editWallet(getPos(walletModel.Wallets(), qWalletList[e]), qWalletList[e])
			// }
		}
	}
}

func (walletModel *WalletModel) addWallet(w *QWallet) {
	logWalletsModel.Info("Add Wallet")
	// if w.Pointer() == nil {
	// 	return
	// }

	var row = len(walletModel.Wallets())
	for pos, wlt := range walletModel.Wallets() {
		if w.Name() < wlt.Name() {
			row = pos
			break
		}
	}

	walletModel.walletByName[w.FileName()] = w

	walletModel.BeginInsertRows(qtCore.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(w, qml.QQmlEngine__CppOwnership)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], append([]*QWallet{w}, walletModel.Wallets()[row:]...)...))
	walletModel.SetCount(walletModel.Count() + 1)
	walletModel.EndInsertRows()
	logWalletsModel.Info("End Add Wallet")
}

func (walletModel *WalletModel) editWallet(row int, wallet *QWallet) {
	logWalletsModel.Info("Edit Wallet")
	if row == -1 {
		return
	}
	walletModel.removeWallet(row)
	walletModel.addWallet(wallet)
	// lIndex := walletModel.Index(len(walletModel.Wallets())-1, 0, qtCore.NewQModelIndex())
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
	walletModel.BeginRemoveRows(qtCore.NewQModelIndex(), row, row)
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

func FromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {
	qWallet := NewQWallet(nil)
	qWallet.SetName(wlt.GetLabel())
	qWallet.SetAddresses(address.NewModelAddress(nil))
	qWallet.SetFileName(wlt.GetId())
	qWallet.SetCurrency(wlt.GetCoinType())
	qWallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qWallet.SetEncryptionEnabled(1)
	}

	coinOpts := modelUtil.NewMap(nil)
	logWalletsModel.Info("Test:Assets list: ", wlt.GetCryptoAccount().ListAssets())

	// Iterate on all asset (the first asset need be the main asset ) and obtains the balance for that.
	for _, asset := range wlt.GetCryptoAccount().ListAssets() {
		// 	logWalletsModel.Info(asset)
		bl, err := wlt.GetCryptoAccount().GetBalance(asset)
		if err != nil {
			logWalletsModel.WithError(err).Warnf("Couldn't get %s balance", util.AltcoinCaption(asset))
		} else {
			accuracy, err := util.AltcoinQuotient(asset)
			if err != nil {
				logWalletsModel.WithError(err).Warnf("Couldn't get %s Altcoin quotient", util.AltcoinCaption(asset))
			}
			logWalletsModel.Info("Test:balance ", util.FormatCoins(bl, accuracy)+" "+asset)
			coinOpts.SetValueAsync(util.AltcoinCaption(asset), util.FormatCoins(bl, accuracy)+" "+asset)
		}
	}

	qWallet.SetCoinOptions(coinOpts)

	addrIt, err := wlt.GetLoadedAddresses()

	if err != nil {
		logWalletsModel.WithError(err).Error("Couldn't get address iterator ")
	}

	var addressList = address.NewModelAddress(nil)
	var addressDetailList = make([]*address.QAddress, 0)
	for addrIt.Next() {
		var addrsDetail = address.NewQAddress(nil)
		addrsDetail.SetWalletId(wlt.GetId())
		addrsDetail.SetAddress(addrIt.Value().String())

		var addressCoinOption = modelUtil.NewMap(nil)
		for _, asset := range addrIt.Value().GetCryptoAccount().ListAssets() {
			balance, err := addrIt.Value().GetCryptoAccount().GetBalance(asset)
			if err != nil {
				logWalletsModel.WithError(err).Warnf("Couldn't get balance for %s asset", asset)
			}

			accuracy, err := util.AltcoinQuotient(asset)
			if err != nil {
				logWalletsModel.WithError(err).Warnf("Couldn't get accuracy for %s asset", asset)
			}

			addressCoinOption.SetValueAsync(asset, util.FormatCoins(balance, accuracy))
		}

		addrsDetail.SetCoinOptions(addressCoinOption)
		addressDetailList = append(addressDetailList, addrsDetail)
	}

	addressList.SetAddresses(addressDetailList)
	qWallet.SetAddresses(addressList)

	return qWallet
}
