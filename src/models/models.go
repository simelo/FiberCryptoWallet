package models

func init() {

	BlockchainStatusModel_QmlRegisterType2("BlockchainModels", 1, 0, "BlockchainStatusModel")
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
	ConfigManager_QmlRegisterType2("Config", 1, 0, "ConfigManager")
	KeyValueStorage_QmlRegisterType2("Config", 1, 0, "Options")
	ModelManager_QmlRegisterType2("WalletsManager", 1, 0, "ModelManager")
	ModelWallets_QmlRegisterType2("OutputsModels", 1, 0, "QWallets")
	ModelAddresses_QmlRegisterType2("OutputsModels", 1, 0, "QAddresses")
	ModelOutput_QmlRegisterType2("OutputsModels", 1, 0, "QOutputs")
	QBridge_QmlRegisterType2("Utils", 1, 0, "QBridge")
}
