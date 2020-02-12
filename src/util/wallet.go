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

// GetId provides transaction output ID
func (wo *SimpleWalletOutput) GetId() string {
	return wo.UxOut.GetId()
}

// IsSpent determines whether there exists a confirmed transaction with an input spending this output
func (wo *SimpleWalletOutput) IsSpent() bool {
	return wo.UxOut.IsSpent()
}

// GetAddress returns the address of the party receiving funds
func (wo *SimpleWalletOutput) GetAddress() core.Address {
	return wo.UxOut.GetAddress()
}

// GetCoins looks up coins for asset represented by ticker that have been transferred in this output
func (wo *SimpleWalletOutput) GetCoins(ticker string) (uint64, error) {
	return wo.UxOut.GetCoins(ticker)
}

// SupportedAssets enumerates tickers of crypto assets supported by this output
func (wo *SimpleWalletOutput) SupportedAssets() []string {
	return wo.UxOut.SupportedAssets()
}

// SimpleWalletAddress put together address with owner wallet
type SimpleWalletAddress struct {
	Wallet  core.Wallet
	Address core.Address
}

// GetWallet return wallet
func (wa *SimpleWalletAddress) GetWallet() core.Wallet {
	return wa.Wallet
}

// IsBip32 flag shall be set if address generation complies to BIP 32
func (wa *SimpleWalletAddress) IsBip32() bool {
	return wa.Address.IsBip32()
}

// String return human-readable representation of this address
func (wa *SimpleWalletAddress) String() string {
	return wa.Address.String()
}

// GetCryptoAccount provides access to address transaction history
func (wa *SimpleWalletAddress) GetCryptoAccount() core.CryptoAccount {
	return wa.Address.GetCryptoAccount()
}

// Bytes binary representation for address
func (wa *SimpleWalletAddress) Bytes() []byte {
	return wa.Address.Bytes()
}

// Checksum computes address consistency token
func (wa *SimpleWalletAddress) Checksum() core.Checksum {
	return wa.Address.Checksum()
}

// Verify checks that the address appears valid for the public key
func (wa *SimpleWalletAddress) Verify(pubkey core.PubKey) error {
	return wa.Address.Verify(pubkey)
}

// Null returns true if the address is null
func (wa *SimpleWalletAddress) Null() bool {
	return wa.Address.Null()
}

// Type assertions
var (
	_ core.TransactionOutput = &SimpleWalletOutput{}
	_ core.WalletObject      = &SimpleWalletOutput{}
	_ core.Address           = &SimpleWalletAddress{}
	_ core.WalletObject      = &SimpleWalletOutput{}
)
