package core

// WalletIterator iterates over sequences of wallets
type WalletIterator interface {
	// Value of wallet at iterator pointer position
	Value() FullWallet
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// WalletSet allows for creating wallets and listing those in a set
type WalletSet interface {
	// ListWallets returns an iterator over wallets in the set
	ListWallets() WalletIterator
	// GetWallet to lookup wallet by ID
	GetWallet(id string) FullWallet
	// CreateWallet instantiates a new wallet given account seed
	CreateWallet(name string, seed string, walletType string, isEncryptrd bool, pwd PasswordReader, scanAddressesN int) (FullWallet, error)
	// DefaultWalletType default wallet type
	DefaultWalletType() string
	// SupportedWalletTypes list supported wallet type names
	SupportedWalletTypes() []string
}

// WalletStorage provides access to the underlying wallets data store
type WalletStorage interface {
	// Encrypt protects wallet data using cryptography
	Encrypt(walletName string, password PasswordReader)
	// Decrypt unlocks wallet for accessing internal data
	Decrypt(walletName string, password PasswordReader)
	// IsEncrypted queries whether wallet data is encrypted or not
	IsEncrypted(walletName string) (bool, error)
}

// AddressType differentiates between BIP44 public external and internal change addresses
type AddressType uint32

const (
	// AccountAddress refers to public external address for sharing with peers
	AccountAddress AddressType = iota
	// ChangeAddress refers to internal change address
	ChangeAddress
)

// Wallet defines the minimal contract implemented by every wallet
type Wallet interface {
	// WalletId returns wallet local identifier
	WalletId() string
	// WalletLabel provides a human-readable name for this wallet to be shown in GUI controls
	WalletLabel() string
	// SetLabel establishes a label for this wallet
	SetWalletLabel(wltName string)
	// WalletCryptoAccount instantiate object to determine wallet balance and transaction history
	WalletCryptoAccount() CryptoAccount
}

// WatchWallet defines the contract for distributing-only wallets
type WatchWallet interface {
	// GenAddresses discover new addresses based on default hierarchically deterministic derivation sequences
	GenAddresses(accountIndex uint32, addrType AddressType, startIndex, count uint32, pwd PasswordReader) (AddressIterator, error)
	// GetLoadedAddressesForAccount iterates over wallet addresses discovered and known to have previous history and coins
	GetLoadedAddressesForAccount(accountIndex uint32, addrType AddressType) (AddressIterator, error)
	// GetAllLoadedAddresses iterates over wallet addresses discovered and known to have previous history and coins
	GetAllLoadedAddresses() (AddressIterator, error)
}

// FullWallet defines the contract that must be satisfied by altcoin crypto wallets
type FullWallet interface {
	Wallet
	BlockchainVisor
	WatchWallet
	TxnSigner
}

// SeedGenerator establishes the contract for generating BIP39-compatible mnemonics
type SeedGenerator interface {
	// GenerateMnemonic generates a valid BIP-39 mnemonic phrase
	GenerateMnemonic(wordCount int) (string, error)
	// VerifyMnemonic shall determine whether a given mnemonic phrase complies to BIP39
	VerifyMnemonic(seed string) (bool, error)
}

// WalletEnv is the entry point to manage wallets
type WalletEnv interface {
	// GetStorage provides access to wallet data store
	GetStorage() WalletStorage
	// GetWalletSet loads wallets in this environment
	GetWalletSet() WalletSet
}

// WalletObject represents the contract implemented by objects managed by wallets
type WalletObject interface {
	GetWallet() Wallet
}
