package transactions

import (
	"github.com/therecipe/qt/core"
)

const (
	Name = int(core.Qt__UserRole) + 1<<iota
	EncryptionEnabled
	Sky
	CoinHours
	FileName
)

// WalletModel duplicated!!!
type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Wallet                `property:"wallets"`

	_ func(*Wallet)                                                          `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky, coinHours int) `slot:"editWallet"`
	_ func(row int)                                                          `slot:"removeWallet"`
	_ func()                                                                 `slot:"loadModel"`
}

func (m *WalletModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		Sky:               core.NewQByteArray2("sky", -1),
		CoinHours:         core.NewQByteArray2("coinHours", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddWallet(m.addWallet)
	m.ConnectEditWallet(m.editWallet)
	m.ConnectRemoveWallet(m.removeWallet)
	m.ConnectLoadModel(m.loadModel)

	m.loadModel()

}

func (m *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Wallets()) {
		return core.NewQVariant()
	}

	var w = m.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant1(w.Name())
		}

	case EncryptionEnabled:
		{
			return core.NewQVariant1(w.EncryptionEnabled())
		}

	case Sky:
		{
			return core.NewQVariant1(w.Sky())
		}

	case CoinHours:
		{
			return core.NewQVariant1(w.CoinHours())
		}
	case FileName:
		{
			return core.NewQVariant1(w.FileName())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Wallets())
}

func (m *WalletModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *WalletModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *WalletModel) addWallet(w *Wallet) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Wallets()), len(m.Wallets()))
	m.SetWallets(append(m.Wallets(), w))
	m.EndInsertRows()

}

func (m *WalletModel) editWallet(row int, name string, encrypted bool, sky, coinHours int) {
	w := m.Wallets()[row]
	w.SetName(name)
	w.SetEncryptionEnabled(0)
	if encrypted {
		w.SetEncryptionEnabled(1)
	}
	w.SetSky(sky)
	w.SetCoinHours(coinHours)
	w.SetFileName(w.FileName())
	pIndex := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, EncryptionEnabled, Sky, CoinHours, FileName})
}

func (m *WalletModel) removeWallet(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetWallets(append(m.Wallets()[:row], m.Wallets()[row+1:]...))
	m.EndRemoveRows()
}

// func getWalletsModel() ([]*Wallet, error) {
// 	c := util.NewClient()
// 	wallets, err := c.Wallets()
// 	if err != nil {
// 		return nil, err
// 	}
// 	walletsModels := make([]*Wallet, 0)
// 	for _, w := range wallets {
// 		wlt, err := walletResponseToWallet(&w)
// 		if err != nil {
// 			return nil, err
// 		}
// 		walletsModels = append(walletsModels, wlt)
// 	}
// 	return walletsModels, nil

// }

func (m *WalletModel) loadModel() {

	wlt := NewWallet(nil)
	wlt.SetName("Wallet1")
	wlt.SetEncryptionEnabled(0)
	wlt.SetSky(12)
	wlt.SetCoinHours(6)
	wlt.SetFileName("fdfdfdfdfdfdfdfd")
	m.addWallet(wlt)

	wlt1 := NewWallet(nil)
	wlt1.SetName("Wallet2")
	wlt1.SetEncryptionEnabled(1)
	wlt1.SetSky(123)
	wlt1.SetCoinHours(61)
	wlt1.SetFileName("ddddfdfdfdfdfdfd")
	m.addWallet(wlt1)

}
