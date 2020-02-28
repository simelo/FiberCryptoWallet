package outputs

import (
	"context"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sort"
)

func init() {
	ModelOutputs_QmlRegisterType2("Outputs", 1, 0, "QOutputs")
}

var logModelOutputs = logging.MustGetLogger("modelOutputs")

const (
	OutputID = int(qtCore.Qt__UserRole) + iota + 1
	CoinOpts
	AddressOwner
	WalletOwner
)

type ModelOutputs struct {
	qtCore.QAbstractListModel

	Ctx    context.Context
	cancel context.CancelFunc

	_ func() `constructor:"init"`
	_ func() `destructor:"destroy"`

	_ map[int]*qtCore.QByteArray `property:"roles"`
	_ []*QOutput                 `property:"outputs"`
	_ string                     `property:"address"`
	_ bool                       `property:"loading"`
	_ func([]*QOutput)           `slot:"addOutputs"`
	_ func(row int) *QOutput     `slot:"get"`
	_ func(output []*QOutput)    `slot:"loadModelAsync"`
	_ func(row int)              `slot:"removeOutputs"`
}

type QOutput struct {
	qtCore.QObject

	_ string         `property:"outputID"`
	_ string         `property:"addressOwner"`
	_ string         `property:"walletOwner"`
	_ *modelUtil.Map `property:"coinOpt"`
}

func (modelOutputs *ModelOutputs) init() {
	modelOutputs.SetRoles(map[int]*qtCore.QByteArray{
		OutputID:     qtCore.NewQByteArray2("outputID", -1),
		CoinOpts:     qtCore.NewQByteArray2("coinOpts", -1),
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
	modelOutputs.SetLoading(true)
}

func (modelOutputs *ModelOutputs) destroy() {
	logModelOutputs.Info("Destroy ModelOutputs")
	modelOutputs.cancel()
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
	case CoinOpts:
		{
			return qtCore.NewQVariant1(qo.CoinOpt())
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

func (modelOutputs *ModelOutputs) loadModelAsync(qOutList []*QOutput) {

	if len(modelOutputs.Outputs()) == 0 {
		for e := range qOutList {
			modelOutputs.addOutputs(qOutList[e])
		}
		return
	}

	removePosList := make([]int, 0)

	for k, v := range modelOutputs.Outputs() {
		if !contains(qOutList, v) {
			removePosList = append(removePosList, k)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(removePosList)))

	for e := range removePosList {
		modelOutputs.RemoveOutputs(removePosList[e])
	}

	for e := range qOutList {
		if !contains(modelOutputs.Outputs(), qOutList[e]) {
			modelOutputs.addOutputs(qOutList[e])
		}
	}
}

func (modelOutputs *ModelOutputs) addOutputs(output *QOutput) {
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

func (modelOutputs *ModelOutputs) removeOutputs(row int) {
	logModelOutputs.Info("Remove outputs for index")

	modelOutputs.BeginRemoveRows(qtCore.NewQModelIndex(), row, row)
	modelOutputs.SetOutputs(append(modelOutputs.Outputs()[:row], modelOutputs.Outputs()[row+1:]...))
	modelOutputs.EndRemoveRows()
}

func contains(outputs []*QOutput, output *QOutput) bool {
	for _, out := range outputs {
		if out.OutputID() == output.OutputID() {
			return true
		}
	}
	return false
}

func FromOutputsToQOutputs(output core.TransactionOutput, wltName string) *QOutput {
	qOutput := NewQOutput(nil)
	qOutput.SetOutputID(output.GetId())
	qOutput.SetWalletOwner(wltName)
	outAddrs, err := output.GetAddress()
	if err != nil {
		logModelOutputs.WithError(err).Errorf("Couldn't get address for output %s", output.GetId())
		return qOutput
	}
	qOutput.SetAddressOwner(outAddrs.String())
	coinOpts := modelUtil.NewMap(nil)

	for _, coinTrait := range output.GetCoinTraits() {
		coinOpts.SetValue(coinTrait.GetTrait(), coinTrait.GetValue())
	}

	qOutput.SetCoinOpt(coinOpts)

	return qOutput
}
