package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/models/addresses"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/blockchain"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"
)

func init() {
	wallets.WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	wallets.QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	addresses.AddressesModel_QmlRegisterType2("WalletsManager", 1, 0, "AddressModel")
	addresses.QAddress_QmlRegisterType2("WalletsManager", 1, 0, "QAddress")
	wallets.WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
	transactions.TransactionModel_QmlRegisterType2("TransactionManager", 1, 0, "QTransactionModel")
	blockchain.BlockchainStatusModel_QmlRegisterType2("BlockchainModels", 1, 0, "BlockchainStatusModel")
}
