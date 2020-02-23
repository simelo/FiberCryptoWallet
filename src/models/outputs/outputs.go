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

func (m *ModelOutputs) init() {
	m.SetRoles(map[int]*qtCore.QByteArray{
		OutputID:     qtCore.NewQByteArray2("outputID", -1),
		CoinOpts:     qtCore.NewQByteArray2("coinOpts", -1),
		AddressOwner: qtCore.NewQByteArray2("addressOwner", -1),
		WalletOwner:  qtCore.NewQByteArray2("walletOwner", -1),
	})

	m.ConnectDestroyModelOutputs(m.destroy)
	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectRemoveOutputs(m.removeOutputs)
	m.ConnectGet(m.get)
	m.ConnectLoadModelAsync(m.loadModelAsync)
	m.Ctx, m.cancel = context.WithCancel(context.Background())
	m.SetLoading(true)
}

func (m *ModelOutputs) destroy() {
	logModelOutputs.Info("Destroy")
	m.cancel()
}

func (m *ModelOutputs) rowCount(*qtCore.QModelIndex) int {
	return len(m.Outputs())
}

func (m *ModelOutputs) roleNames() map[int]*qtCore.QByteArray {
	return m.Roles()
}

func (m *ModelOutputs) data(index *qtCore.QModelIndex, role int) *qtCore.QVariant {
	if !index.IsValid() {
		return qtCore.NewQVariant()
	}

	if index.Row() >= len(m.Outputs()) {
		return qtCore.NewQVariant()
	}

	qo := m.Outputs()[index.Row()]

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

func (m *ModelOutputs) get(row int) *QOutput {
	if row >= len(m.Outputs()) {
		logModelOutputs.Errorf("Unable access to index %d: Index out of range", row)
		return NewQOutput(nil)
	}
	return m.Outputs()[row]
}

func (m *ModelOutputs) loadModelAsync(qOutList []*QOutput) {

	if len(m.Outputs()) == 0 {
		for e := range qOutList {
			m.addOutputs(qOutList[e])
		}
		return
	}

	removePosList := make([]int, 0)

	for k, v := range m.Outputs() {
		if !contains(qOutList, v) {
			removePosList = append(removePosList, k)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(removePosList)))

	for e := range removePosList {
		m.RemoveOutputs(removePosList[e])
	}

	for e := range qOutList {
		if !contains(m.Outputs(), qOutList[e]) {
			m.addOutputs(qOutList[e])
		}
	}
}

func (m *ModelOutputs) addOutputs(output *QOutput) {
	var row = len(m.Outputs())
	for k, v := range m.Outputs() {
		if output.WalletOwner() < v.WalletOwner() {
			row = k
			break
		}
	}
	m.BeginInsertRows(qtCore.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(output, qml.QQmlEngine__CppOwnership)
	m.SetOutputs(append(m.Outputs()[:row], append([]*QOutput{output}, m.Outputs()[row:]...)...))
	m.EndInsertRows()
}

func (m *ModelOutputs) removeOutputs(row int) {
	logModelOutputs.Info("Remove outputs for index")

	m.BeginRemoveRows(qtCore.NewQModelIndex(), row, row)
	m.SetOutputs(append(m.Outputs()[:row], m.Outputs()[row+1:]...))
	m.EndRemoveRows()
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
	qOutput.SetAddressOwner(output.GetAddress().String())
	coinOpts := modelUtil.NewMap(nil)
	for _, coinTrait := range output.GetCoinTraits() {
		coinOpts.SetValue(coinTrait.GetTrait(), coinTrait.GetValue())
	}
	qOutput.SetCoinOpt(coinOpts)

	return qOutput
}
