package address

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"sync"
	"time"
)

var logAddressModel = logging.MustGetLogger("Address Model")

func init() {
	ModelAddress_QmlRegisterType2("Address", 1, 0, "QAddressList")
	QAddress_QmlRegisterType2("Address", 1, 0, "QAddress")
}

const (
	Address = int(qtCore.Qt__UserRole) + 1<<iota
	CoinOptions
	WalletId
)

type updateAddrModel struct {
	addressList []*QAddress
	toReplace   bool
}

type updateQAddr struct {
	strAddr string
	coinFtr *modelUtil.Map
}

type QAddress struct {
	qtCore.QObject
	corAddr core.Address
	_       string         `property:"address"`
	_       *modelUtil.Map `property:"coinOptions"`
	_       string         `property:"walletId"`
}

type ModelAddress struct {
	qtCore.QAbstractListModel
	wltByAddr   map[string]string
	wltList     []string
	stop        chan struct{}
	updtModelCh chan updateAddrModel
	wg          sync.WaitGroup

	_ map[int]*qtCore.QByteArray `property:"roles"`

	_ func() `constructor:"init"`
	_ func() `destructor:"destroy"`

	_ func(address *QAddress)                      `signal:"addAddress"`
	_ func(index int)                              `signal:"removeAddress,auto"`
	_ func(index int, coinFtr *modelUtil.Map)      `signal:"editAddress"`
	_ func(strAddr string, coinFtr *modelUtil.Map) `signal:"editAddressByStrAddr"`
	_ func(addressList []*QAddress)                `slot:"loadModelAsync"`
	_ func(addressList []*QAddress)                `slot:"loadModel"`
	_ func()                                       `slot:"cleanModel"`
	_ func()                                       `slot:"start"`

	_ []*QAddress `property:"addresses"`
}

func (modelAddrs *ModelAddress) init() {
	modelAddrs.SetRoles(map[int]*qtCore.QByteArray{
		Address:     qtCore.NewQByteArray2("address", -1),
		CoinOptions: qtCore.NewQByteArray2("coinOptions", -1),
		WalletId:    qtCore.NewQByteArray2("walletId", -1),
	})

	modelAddrs.ConnectRowCount(modelAddrs.rowCount)
	modelAddrs.ConnectData(modelAddrs.data)
	modelAddrs.ConnectRoleNames(modelAddrs.roleNames)
	modelAddrs.ConnectLoadModelAsync(modelAddrs.loadModelAsync)
	modelAddrs.ConnectLoadModel(modelAddrs.loadModel)
	modelAddrs.ConnectAddAddress(modelAddrs.addAddress)
	modelAddrs.ConnectEditAddress(modelAddrs.editAddress)
	modelAddrs.ConnectEditAddressByStrAddr(modelAddrs.editAddressByStrAddr)
	modelAddrs.ConnectStart(modelAddrs.start)
	modelAddrs.ConnectCleanModel(modelAddrs.cleanModel)
	modelAddrs.ConnectDestroyModelAddress(modelAddrs.destroy)

	modelAddrs.wg = sync.WaitGroup{}
	modelAddrs.stop = make(chan struct{}, 1)
	modelAddrs.wltByAddr = make(map[string]string)
	modelAddrs.wltList = make([]string, 0)
	modelAddrs.updtModelCh = make(chan updateAddrModel)
}

func (modelAddrs *ModelAddress) rowCount(*qtCore.QModelIndex) int {
	return len(modelAddrs.Addresses())
}

func (modelAddrs *ModelAddress) roleNames() map[int]*qtCore.QByteArray {
	return modelAddrs.Roles()
}

func (modelAddrs *ModelAddress) addAddress(address *QAddress) {
	logAddressModel.Infof("Adding Address %s", address.corAddr.String())
	var pos = 0
	for index, qAddr := range modelAddrs.Addresses() {
		if qAddr.WalletId() > address.WalletId() {
			pos = index
			break
		}
	}
	modelAddrs.BeginInsertRows(qtCore.NewQModelIndex(), pos, pos)
	modelAddrs.SetAddresses(append(modelAddrs.Addresses()[:pos], append([]*QAddress{address}, modelAddrs.Addresses()[pos:]...)...))
	modelAddrs.EndInsertRows()
	modelAddrs.wltByAddr[address.Address()] = address.WalletId()
}

func (modelAddrs *ModelAddress) editAddress(index int, coinFtr *modelUtil.Map) {
	logAddressModel.Info("edit address in index :", index)
	modelAddrs.Addresses()[index].SetCoinOptions(coinFtr)
	pindex := modelAddrs.Index(0, 0, qtCore.NewQModelIndex())
	lindex := modelAddrs.Index(len(modelAddrs.Addresses())-1, 0, qtCore.NewQModelIndex())
	modelAddrs.DataChanged(pindex, lindex, []int{CoinOptions})
}

func (modelAddrs *ModelAddress) removeAddress(index int) {
	logAddressModel.Info("Removing Address")
	if index >= len(modelAddrs.Addresses()) {
		return
	}
	strAddr := modelAddrs.Addresses()[index].Address()
	modelAddrs.BeginRemoveRows(qtCore.NewQModelIndex(), index, index)
	modelAddrs.SetAddresses(append(modelAddrs.Addresses()[:index], modelAddrs.Addresses()[index+1:]...))
	modelAddrs.EndRemoveRows()
	delete(modelAddrs.wltByAddr, strAddr)
}

func (modelAddrs *ModelAddress) data(index *qtCore.QModelIndex, role int) *qtCore.QVariant {
	if !index.IsValid() || index.Row() >= len(modelAddrs.Addresses()) {
		return qtCore.NewQVariant()
	}

	address := modelAddrs.Addresses()[index.Row()]

	switch role {
	case Address:
		{
			return qtCore.NewQVariant1(address.Address())
		}
	case CoinOptions:
		{
			return qtCore.NewQVariant1(address.CoinOptions())
		}
	case WalletId:
		{
			return qtCore.NewQVariant1(address.WalletId())
		}
	default:
		{
			return qtCore.NewQVariant()
		}
	}
}

func (modelAddrs *ModelAddress) loadModel(addrList []*QAddress) {
	// if len(modelAddrs.Addresses()) == 0 {
	// 	for e := range addrList {
	// 		modelAddrs.AddAddress(addrList[e])
	// 	}
	// 	return
	// }
	logAddressModel.Info(modelAddrs.wltByAddr)
	if len(modelAddrs.Addresses()) < len(addrList) {
		for e := range addrList {
			if _, ok := modelAddrs.wltByAddr[addrList[e].corAddr.String()]; !ok {
				modelAddrs.addAddress(addrList[e])
			}
		}
	}
}

func (modelAddrs *ModelAddress) loadModelAsync(addrList []*QAddress) {
	if len(modelAddrs.wltList) == 0 {
		modelAddrs.wltList = []string{addrList[0].WalletId()}
		modelAddrs.updtModelCh <- updateAddrModel{
			addressList: addrList,
			toReplace:   true,
		}
		return
	}

	if modelAddrs.wltList[0] != addrList[0].WalletId() {
		modelAddrs.wltList = []string{addrList[0].WalletId()}
		modelAddrs.updtModelCh <- updateAddrModel{
			addressList: addrList,
			toReplace:   true,
		}
	} else if len(modelAddrs.Addresses()) < len(addrList) {
		addrsToUpdate := make([]*QAddress, 0)
		for e := range addrList {
			if _, ok := modelAddrs.wltByAddr[addrList[e].Address()]; !ok {
				addrsToUpdate = append(addrsToUpdate, addrList[e])
			}
		}

		modelAddrs.updtModelCh <- updateAddrModel{
			addressList: addrsToUpdate,
			toReplace:   false,
		}
	}
}

func (modelAddrs *ModelAddress) start() {
	updateAddressList := func(updtQAddrCh chan updateQAddr) {
		for _, qAddr := range modelAddrs.Addresses() {
			go func(addr core.Address, updateQAddrsCh chan updateQAddr) {
				LoadCoinFtrAsync(addr, updateQAddrsCh)
			}(qAddr.corAddr, updtQAddrCh)
		}
	}
	go func() {
		mutex := sync.Mutex{}
		modelAddrs.wg.Add(1)
		updateQAddrsCh := make(chan updateQAddr, 100)
		for {
			select {
			case <-time.After(time.Duration(config.GetDataUpdateTime()) * time.Second):
				updateAddressList(updateQAddrsCh)
				break

			case update := <-updateQAddrsCh:
				modelUtil.Helper.RunInMain(func() {
					mutex.Lock()
					defer mutex.Unlock()
					modelAddrs.editAddressByStrAddr(update.strAddr, update.coinFtr)
				})
				break
			case updateModel := <-modelAddrs.updtModelCh:
				if updateModel.toReplace {
					modelUtil.Helper.RunInMain(func() {
						mutex.Lock()
						defer mutex.Unlock()
						modelAddrs.cleanModel()
					})
				}

				modelUtil.Helper.RunInMain(func() {
					mutex.Lock()
					defer mutex.Unlock()
					for _, newAddr := range updateModel.addressList {
						modelAddrs.addAddress(newAddr)
					}
					updateAddressList(updateQAddrsCh)
				})
				break
			case <-modelAddrs.stop:
				close(updateQAddrsCh)
				close(modelAddrs.stop)
				close(modelAddrs.updtModelCh)
				modelAddrs.wg.Done()
				return
			}
		}
	}()
}

func (modelAddrs *ModelAddress) cleanModel() {
	modelAddrs.BeginResetModel()
	modelAddrs.SetAddresses(make([]*QAddress, 0))
	modelAddrs.EndResetModel()
	modelAddrs.wltByAddr = make(map[string]string)
}

func (modelAddrs *ModelAddress) destroy() {
	modelAddrs.stop <- struct{}{}
	modelAddrs.wg.Wait()
}

func (modelAddrs *ModelAddress) editAddressByStrAddr(strAddr string, coinFtr *modelUtil.Map) {
	logAddressModel.Infof("edit address %s", strAddr)
	if _, ok := modelAddrs.wltByAddr[strAddr]; ok {
		for index, address := range modelAddrs.Addresses() {
			if address.corAddr.String() == strAddr {
				modelAddrs.editAddress(index, coinFtr)
				return
			}
		}
	}
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
	return a.Address() == b.Address() && modelUtil.CompareMaps(a.CoinOptions(), b.CoinOptions()) &&
		a.WalletId() == b.WalletId()
}

func FromCorAddrToQAddr(address core.Address, wltName string) *QAddress {
	qAddr := NewQAddress(nil)
	qAddr.corAddr = address
	qAddr.SetAddress(address.String())
	qAddr.SetWalletId(wltName)
	coinFtr := modelUtil.NewMap(nil)
	for _, asset := range address.GetCryptoAccount().ListAssets() {
		coinFtr.SetValue(asset, "N/A")
	}

	qAddr.SetCoinOptions(coinFtr)
	return qAddr
}

func LoadCoinFtrAsync(address core.Address, update chan updateQAddr) {
	newCF, err := LoadCoinFtr(address)
	if err != nil {
		return
	}

	update <- updateQAddr{
		strAddr: address.String(),
		coinFtr: newCF,
	}
}

func LoadCoinFtr(address core.Address) (*modelUtil.Map, error) {
	coinFtr := modelUtil.NewMap(nil)
	for _, asset := range address.GetCryptoAccount().ListAssets() {
		balance, err := address.GetCryptoAccount().GetBalance(asset)
		if err != nil {
			logAddressModel.WithError(err).Warnf(
				"Couldn't get balance for asset: %s in address %s", asset, address.String())
			return nil, err
		}

		accuracy, err := util.AltcoinQuotient(asset)
		if err != nil {
			logAddressModel.WithError(err).Warnf(
				"Couldn't get accuracy for asset: %s in address %s", asset, address.String())
			return nil, err

		}

		coinFtr.SetValueAsync(asset, util.FormatCoins(balance, accuracy))
	}
	return coinFtr, nil
}
