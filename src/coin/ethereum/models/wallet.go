package ethereum

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func NewWalletDirectory(path string) *WalletsDirectory {
	return &WalletsDirectory{
		path: path,
	}
}

// Implements WalletEnv, WalletStorage and WalletSet interfaces
type WalletsDirectory struct {
	path string
}

func (walletDir *WalletsDirectory) GetStorage() core.WalletStorage {
	return walletDir
}

func (walletDir *WalletsDirectory) GetWalletSet() core.WalletSet {
	return walletDir
}
