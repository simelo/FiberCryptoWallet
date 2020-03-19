package outputs

import (
	"context"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sync"
	"time"
)

func init() {
	ModelOutputs_QmlRegisterType2("Outputs", 1, 0, "QOutputs")
}

var logModelOutputs = logging.MustGetLogger("modelOutputs")

const (
	OutputID = int(qtCore.Qt__UserRole) + iota + 1
	CoinFeature
	AddressOwner
	WalletOwner
)

type ModelOutputs struct {
	qtCore.QAbstractListModel

	Ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	_      func() `constructor:"init"`
	_      func() `destructor:"destroy"`

	_ map[int]*qtCore.QByteArray `property:"roles"`
	_ []*QOutput                 `property:"outputs"`
	_ string                     `property:"address"`
	_ bool                       `property:"loading"`
	_ func([]*QOutput)           `slot:"addOutputs"`
	_ func(row int) *QOutput     `slot:"get"`
	_ func()                     `slot:"loadModelAsync"`
	_ func(row int)              `slot:"removeOutputs"`
}

type QOutput struct {
	qtCore.QObject
	coreOut core.TransactionOutput
	_       string         `property:"outputID"`
	_       string         `property:"addressOwner"`
	_       string         `property:"walletOwner"`
	_       *modelUtil.Map `property:"coinFtr"`
}

func (modelOutputs *ModelOutputs) init() {
	modelOutputs.SetRoles(map[int]*qtCore.QByteArray{
		OutputID:     qtCore.NewQByteArray2("outputID", -1),
		CoinFeature:  qtCore.NewQByteArray2("coinFtr", -1),
		AddressOwner: qtCore.NewQByteArray2("addressOwner", -1),
		WalletOwner:  qtCore.NewQByteArray2("walletOwner", -1),
	})

	modelOutputs.ConnectDestroyModelOutputs(modelOutputs.destroy)

	modelOutputs.ConnectRowCount(modelOutputs.rowCount)
	modelOutputs.ConnectRoleNames(modelOutputs.roleNames)
	modelOutputs.ConnectData(modelOutputs.data)
	modelOutputs.ConnectRemoveOutputs(modelOutputs.removeOutputs)
	modelOutputs.ConnectGet(modelOutputs.get)
	modelOutputs.ConnectLoadModelAsync(modelOutputs.loadModelAsync)
	modelOutputs.Ctx, modelOutputs.cancel = context.WithCancel(context.Background())
	// modelOutputs.walletsByOutputs = make(map[string]string)
	modelOutputs.wg = sync.WaitGroup{}
	modelOutputs.SetLoading(true)
}

func (modelOutputs *ModelOutputs) destroy() {
	logModelOutputs.Info("Destroy ModelOutputs")
	modelOutputs.cancel()
	modelOutputs.wg.Wait()

}

func (modelOutputs *ModelOutputs) rowCount(*qtCore.QModelIndex) int {
	return len(modelOutputs.Outputs())
}

func (modelOutputs *ModelOutputs) roleNames() map[int]*qtCore.QByteArray {
	return modelOutputs.Roles()
}

func (modelOutputs *ModelOutputs) data(index *qtCore.QModelIndex, role int) *qtCore.QVariant {
	if !index.IsValid() {
		return qtCore.NewQVariant()
	}

	if index.Row() >= len(modelOutputs.Outputs()) {
		return qtCore.NewQVariant()
	}

	qo := modelOutputs.Outputs()[index.Row()]

	switch role {
	case OutputID:
		{
			return qtCore.NewQVariant1(qo.OutputID())
		}
	case CoinFeature:
		{
			return qtCore.NewQVariant1(qo.CoinFtr())
		}
	case AddressOwner:
		{
			return qtCore.NewQVariant1(qo.AddressOwner())
		}
	case WalletOwner:
		{
			return qtCore.NewQVariant1(qo.WalletOwner())
		}
	default:
		{
			return qtCore.NewQVariant()
		}
	}
}

func (modelOutputs *ModelOutputs) get(row int) *QOutput {
	if row >= len(modelOutputs.Outputs()) {
		logModelOutputs.Errorf("Unable access to index %d: Index out of range", row)
		return NewQOutput(nil)
	}
	return modelOutputs.Outputs()[row]
}

func (modelOutputs *ModelOutputs) loadModelAsync() {
	logModelOutputs.Info("Load model Async")

	isLoadingChan := make(chan bool)
	outToAdd := make(chan *QOutput)
	m := sync.Mutex{}
	OutListToRemove := make(chan []string)
	go loadCoreOutputsAsync(isLoadingChan, outToAdd, OutListToRemove, modelOutputs.Ctx)

	// controller function
	go func() {
		for {
			select {
			case loading := <-isLoadingChan:
				modelUtil.Helper.RunInMain(func() { modelOutputs.SetLoading(loading) })
				break
			case qOut := <-outToAdd:
				modelUtil.Helper.RunInMain(func() {
					m.Lock()
					defer m.Unlock()
					modelOutputs.addOutputs(qOut)

				})
				break
			case outIdList := <-OutListToRemove:
				modelUtil.Helper.RunInMain(func() {
					m.Lock()
					defer m.Unlock()
					for e := range outIdList {
						modelOutputs.removeOutpByOutpId(outIdList[e])
					}
				})
			case <-modelOutputs.Ctx.Done():
				return
			}
		}
	}()
}

func (modelOutputs *ModelOutputs) addOutputs(output *QOutput) {
	logModelOutputs.Infof("Adding output with id: %s", output.OutputID())
	modelOutputs.wg.Add(1)
	defer modelOutputs.wg.Done()

	var row = len(modelOutputs.Outputs())
	for k, v := range modelOutputs.Outputs() {
		if output.WalletOwner() < v.WalletOwner() {
			row = k
			break
		}
	}
	modelOutputs.BeginInsertRows(qtCore.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(output, qml.QQmlEngine__CppOwnership)
	modelOutputs.SetOutputs(append(modelOutputs.Outputs()[:row], append([]*QOutput{output}, modelOutputs.Outputs()[row:]...)...))
	modelOutputs.EndInsertRows()
}

func (modelOutputs *ModelOutputs) removeOutpByOutpId(outId string) {
	logModelOutputs.Infof("Removing output with id: %s", outId)
	for index, output := range modelOutputs.Outputs() {
		if output.coreOut.GetId() == outId {
			modelOutputs.removeOutputs(index)
			return
		}
	}
}

func (modelOutputs *ModelOutputs) removeOutputs(row int) {
	logModelOutputs.Infof("Removing outputs with index: %d", row)
	modelOutputs.wg.Add(1)
	defer modelOutputs.wg.Done()
	modelOutputs.BeginRemoveRows(qtCore.NewQModelIndex(), row, row)
	modelOutputs.SetOutputs(append(modelOutputs.Outputs()[:row], modelOutputs.Outputs()[row+1:]...))
	modelOutputs.EndRemoveRows()
}

func FromOutputsToQOutputs(output core.TransactionOutput, wltName string) *QOutput {
	qOutput := NewQOutput(nil)
	qOutput.coreOut = output
	qOutput.SetOutputID(output.GetId())
	qOutput.SetWalletOwner(wltName)
	outAddrs, err := output.GetAddress()
	if err != nil {
		logModelOutputs.WithError(err).Errorf("Couldn't get address for output %s", output.GetId())
		return qOutput
	}
	qOutput.SetAddressOwner(outAddrs.String())
	coinFtr := modelUtil.NewMap(nil)
	// wg := sync.WaitGroup{}
	coinTrait := output.GetCoinTraits()
	for e := range coinTrait {
		coinFtr.SetValue(coinTrait[e].GetTrait(), coinTrait[e].GetValue())
	}
	logModelOutputs.Info(coinFtr.GetKeys())
	qOutput.SetCoinFtr(coinFtr)

	return qOutput
}

func loadCoinFeatures(output core.TransactionOutput) *modelUtil.Map {
	coinFtr := modelUtil.NewMap(nil)
	for _, trait := range output.GetCoinTraits() {
		models.Helper.RunInMain(func() {
			coinFtr.SetValue(trait.GetTrait(), trait.GetValue())
		})
	}
	return coinFtr
}

func loadCoreOutputsAsync(isLoading chan<- bool, send chan<- *QOutput, outToRemove chan<- []string, ctx context.Context) {
	walletByOutputs := make(map[string]string)
	for {
		select {
		case <-time.After(time.Duration(config.GetDataUpdateTime()) * time.Second):
			walletIter := models.GetWalletEnv().GetWalletSet().ListWallets()
			if walletIter == nil {
				// TODO error
				continue
			}
			for walletIter.Next() {
				outputIter, err := walletIter.Value().GetCryptoAccount().ScanUnspentOutputs()
				if err != nil {
					logModelOutputs.WithError(err).Errorf(
						"Couldn't get output iterator from wallet: %s", walletIter.Value().GetLabel())
					isLoading <- true
					continue
				}
				isLoading <- false
				// Outputs list for the current wallet
				isPersistent := make(map[string]struct{})
				for outputIter.Next() {
					isPersistent[outputIter.Value().GetId()] = struct{}{}
					if _, ok := walletByOutputs[outputIter.Value().GetId()]; !ok {
						walletByOutputs[outputIter.Value().GetId()] = walletIter.Value().GetLabel()
						// <- OutToAdd
						send <- FromOutputsToQOutputs(outputIter.Value(), walletIter.Value().GetLabel())
					}
				}
				removeOutList := make([]string, 0)
				for outputId, walletId := range walletByOutputs {
					if walletId == walletIter.Value().GetLabel() {
						if _, ok := isPersistent[outputId]; !ok {

							removeOutList = append(removeOutList, outputId)
						}
					}
				}

				for e := range removeOutList {
					delete(walletByOutputs, removeOutList[e])
				}
				outToRemove <- removeOutList
				// remove no persistent outputs
				// <- removeOutList
			}
			break
		case <-ctx.Done():
			return
		}
	}
}
