package wallets

import (
	"context"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

const (
	Name = int(qtCore.Qt__UserRole) + iota + 1
	EncryptionEnabled
	FileName
	Addresses
	Currency
	CoinOptions
	Loading
)

func init() {
	WalletModel_QmlRegisterType2("QWallets", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("QWallets", 1, 0, "QWallet")
}

var logWalletsModel = logging.MustGetLogger("Wallets Model")

type WalletModel struct {
	qtCore.QAbstractListModel
	Ctx    context.Context
	cancel context.CancelFunc
	_      func() `constructor:"init"`
	_      func() `destructor:"destroy"`

	_ map[int]*qtCore.QByteArray `property:"roles"`
	_ []*QWallet                 `property:"wallets"`
	_ int                        `property:"count"`

	_ func(*QWallet)                                     `slot:"addWallet"`
	_ func(row int, wallet *QWallet, changedRoles []int) `slot:"editWallet"`
	_ func(row int)                                      `slot:"removeWallet"`
	_ func([]*QWallet)                                   `slot:"loadModelAsync"`
	_ func([]*QWallet)                                   `slot:"loadModel"`
	_ func([]*QWallet)                                   `slot:"updateModel"`
	_ func([]*QWallet, string)                           `slot:"filterWalletByCurrency"`

	_ func() `signal:"changeDetails"`

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
	_ bool                  `property:"loading"`
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
		Loading:           qtCore.NewQByteArray2("loading", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(walletModel, qml.QQmlEngine__CppOwnership)
	walletModel.ConnectDestroyWalletModel(walletModel.destroy)
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
	walletModel.walletByName = make(map[string]*QWallet, 0)

	walletModel.Ctx, walletModel.cancel = context.WithCancel(context.Background())
}

func (walletModel *WalletModel) changeExpanded(id string) {
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
	case Loading:
		{
			return qtCore.NewQVariant1(w.IsLoading())
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
	// logWalletsModel.Infof("Name: %d",Name)
	// logWalletsModel.Infof("EncryptionEnabled: %d",EncryptionEnabled)
	// logWalletsModel.Infof("FileName: %d",FileName)
	// logWalletsModel.Infof("Addresses: %d",Addresses)
	// logWalletsModel.Infof("Currency: %d",Currency)
	// logWalletsModel.Infof("CoinOptions: %d",CoinOptions)
	// logWalletsModel.Infof("Loading: %d",Loading)

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
			if changedRoles := CompareWallet(qWalletList[e],
				walletModel.walletByName[qWalletList[e].FileName()]); len(changedRoles) != 0 {
				// edit wallet changing all properties that match with the roles changed.
				walletModel.editWallet(getPos(walletModel.Wallets(), qWalletList[e]), qWalletList[e], changedRoles)
			}
		}
	}
}

func (walletModel *WalletModel) addWallet(w *QWallet) {
	logWalletsModel.Info("Add Wallet")

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

func (walletModel *WalletModel) editWallet(row int, wallet *QWallet, changedRoles []int) {
	logWalletsModel.Info("Edit Wallet")
	if row == -1 {
		return
	}

	wlt := walletModel.Wallets()[row]
	logWalletsModel.Info(changedRoles)
	for e := range changedRoles {
		switch {
		case changedRoles[e] == Name:
			wlt.SetName(wallet.Name())
			break

		case changedRoles[e] == EncryptionEnabled:
			wlt.SetEncryptionEnabled(wallet.EncryptionEnabled())
			break
		case changedRoles[e] == Addresses:
			wlt.SetAddresses(wallet.Addresses())
			break

		case changedRoles[e] == CoinOptions:
			wlt.SetCoinOptions(wallet.CoinOptions())
			break
		case changedRoles[e] == Loading:
			wlt.SetLoading(wallet.IsLoading())
			break
		}
	}
	pindex := walletModel.Index(0, 0, qtCore.NewQModelIndex())
	lindex := walletModel.Index(len(walletModel.Wallets())-1, 0, qtCore.NewQModelIndex())
	walletModel.DataChanged(pindex, lindex, changedRoles)
	walletModel.ChangeDetails()
}

func (walletModel *WalletModel) removeWallet(row int) {
	logWalletsModel.Info("Remove wallets for index")
	walletModel.BeginRemoveRows(qtCore.NewQModelIndex(), row, row)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], walletModel.Wallets()[row+1:]...))
	walletModel.SetCount(walletModel.Count() - 1)
	walletModel.EndRemoveRows()
}

func (walletModel *WalletModel) updateModel(wallets []*QWallet) {
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

func (walletModel *WalletModel) destroy() {
	logWalletsModel.Info("Destroying WalletModel")
	walletModel.cancel()
}

func CompareWallet(a, b *QWallet) []int {
	var changedRoles = make([]int, 0)
	if !address.CompareModelAddress(a.Addresses(), b.Addresses()) {
		changedRoles = append(changedRoles, Addresses)
	}
	if !modelUtil.CompareMaps(a.CoinOptions(), b.CoinOptions()) {
		changedRoles = append(changedRoles, CoinOptions)
	}
	if a.EncryptionEnabled() != b.EncryptionEnabled() {
		changedRoles = append(changedRoles, EncryptionEnabled)
	}
	if a.FileName() != b.FileName() {
		changedRoles = append(changedRoles, FileName)
	}
	if a.Currency() != b.Currency() {
		changedRoles = append(changedRoles, Currency)
	}
	if a.Name() != b.Name() {
		changedRoles = append(changedRoles, Name)
	}
	if a.IsLoading() != b.IsLoading() {
		changedRoles = append(changedRoles, Loading)
	}
	return changedRoles
}

func FromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {
	qWallet := NewQWallet(nil)
	qWallet.SetName(wlt.GetLabel())
	qWallet.SetAddresses(address.NewModelAddress(nil))
	qWallet.SetFileName(wlt.GetId())
	qWallet.SetCurrency(wlt.GetCoinType())
	qWallet.SetEncryptionEnabled(0)
	qWallet.SetLoading(false)
	if isEncrypted {
		qWallet.SetEncryptionEnabled(1)
	}

	coinOpts := modelUtil.NewMap(nil)

	// Iterate on all asset (the first asset need be the main asset ) and obtains the balance for that.
	for _, asset := range wlt.GetCryptoAccount().ListAssets() {
		bl, err := wlt.GetCryptoAccount().GetBalance(asset)
		if err != nil {
			logWalletsModel.WithError(err).Warnf("Couldn't get %s balance", util.AltcoinCaption(asset))
			qWallet.SetLoading(true)
		} else {
			accuracy, err := util.AltcoinQuotient(asset)
			if err != nil {
				logWalletsModel.WithError(err).Warnf("Couldn't get %s Altcoin quotient", util.AltcoinCaption(asset))
				qWallet.SetLoading(true)
			}
			coinOpts.SetValueAsync(util.AltcoinCaption(asset), util.FormatCoins(bl, accuracy)+" "+asset)
		}
	}

	qWallet.SetCoinOptions(coinOpts)

	addrIt, err := wlt.GetLoadedAddresses()

	if err != nil {
		logWalletsModel.WithError(err).Error("Couldn't get address iterator ")
		qWallet.SetLoading(true)
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
				qWallet.SetLoading(true)
			}
			accuracy, err := util.AltcoinQuotient(asset)
			if err != nil {
				logWalletsModel.WithError(err).Warnf("Couldn't get accuracy for %s asset", asset)
				qWallet.SetLoading(true)
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
