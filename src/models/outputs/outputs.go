package outputs

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	qtCore "github.com/therecipe/qt/core"
)

func init() {
	ModelOutputs_QmlRegisterType2("Outputs", 1, 0, "QOutputs")
}

const (
	OutputID = int(qtCore.Qt__UserRole) + iota + 1
	// AddressSky
	// AddressCoinHours

	// Address
	CoinOpts
	AddressOwner
	WalletOwner
)

type ModelOutputs struct {
	qtCore.QAbstractListModel
	_ func() `constructor:"init"`

	_ map[int]*qtCore.QByteArray `property:"roles"`
	_ []*QOutput                 `property:"outputs"`
	_ string                     `property:"address"`

	_ func([]*QOutput) `slot:"addOutputs"`
	_ func([]*QOutput) `slot:"insertOutputs"`
	_ func([]*QOutput) `slot:"loadModel"`
	_ func()           `slot:"cleanModel"`
	_ func(string)     `slot:"removeOutputsFromAddress"`
	_ func(string)     `slot:"removeOutputsFromWallet"`
}

type QOutput struct {
	qtCore.QObject
	core.TransactionOutput

	_ string `property:"outputID"`
	// _ string `property:"addressSky"`
	// _ string `property:"addressCoinHours"`
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

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddOutputs(m.addOutputs)
	m.ConnectInsertOutputs(m.insertOutputs)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectCleanModel(m.cleanModel)
	m.ConnectRemoveOutputsFromAddress(m.removeOutputsFromAddress)
	m.ConnectRemoveOutputsFromWallet(m.removeOutputsFromWallet)
}

func (m *ModelOutputs) removeOutputsFromAddress(addr string) {
	old := m.Outputs()
	newModel := make([]*QOutput, 0)
	for _, out := range old {
		if out.AddressOwner() != addr {
			newModel = append(newModel, out)
		}
	}
	m.loadModel(newModel)
}

func (m *ModelOutputs) removeOutputsFromWallet(wltId string) {
	old := m.Outputs()
	newModel := make([]*QOutput, 0)
	for _, out := range old {
		if out.WalletOwner() != wltId {
			newModel = append(newModel, out)
		}
	}
	m.loadModel(newModel)
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

func (m *ModelOutputs) insertRows(row int, count int) bool {
	m.BeginInsertRows(qtCore.NewQModelIndex(), row, row+count)
	m.EndInsertRows()
	return true
}

func (m *ModelOutputs) addOutputs(mo []*QOutput) {
	m.SetOutputs(mo)
	m.insertRows(len(m.Outputs()), len(mo))
}

func contains(outputs []*QOutput, output *QOutput) bool {
	for _, out := range outputs {
		x := output.OutputID()
		y := out.OutputID()
		if y == x {
			return true
		}
	}
	return false
}

func (m *ModelOutputs) insertOutputs(mo []*QOutput) {
	toInsert := m.Outputs()
	for _, outputToInsert := range mo {
		if !contains(toInsert, outputToInsert) {
			toInsert = append(toInsert, outputToInsert)
		}
	}
	m.loadModel(toInsert)
}

func (m *ModelOutputs) loadModel(mo []*QOutput) {
	m.BeginResetModel()
	m.SetOutputs(mo)
	m.EndResetModel()
}

func (m *ModelOutputs) cleanModel() {
	m.BeginResetModel()
	m.SetOutputs(make([]*QOutput, 0))
	m.EndResetModel()
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
