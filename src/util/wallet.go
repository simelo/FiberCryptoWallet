package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// SimpleWalletOutput put together transacion output with originating wallet
type SimpleWalletOutput struct {
	Wallet core.Wallet
	UxOut  core.TransactionOutput
}

// GetWallet return wallet
func (wo *SimpleWalletOutput) GetWallet() core.Wallet {
	return wo.Wallet
}

// GetOutput return transaction output.
func (wo *SimpleWalletOutput) GetOutput() core.TransactionOutput {
	return wo.UxOut
}

// SimpleWalletAddress put together address with owner wallet
type SimpleWalletAddress struct {
	Wallet core.Wallet
	UxOut  core.Address
}

// GetWallet return wallet
func (wa *SimpleWalletAddress) GetWallet() core.Wallet {
	return wa.Wallet
}

// GetAddress return address
func (wa *SimpleWalletAddress) GetAddress() core.Address {
	return wa.UxOut
}

// Type assertions
var (
	_ core.WalletOutput  = &SimpleWalletOutput{}
	_ core.WalletAddress = &SimpleWalletAddress{}
)

func IsValidWalletType(walletType string, walletSet core.WalletSet) bool {
	for _, validWalletType := range walletSet.SupportedWalletTypes() {
		if walletType == validWalletType {
			return true
		}
	}
	return false
}
