package hardware

import (
	"github.com/chebyrash/promise"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
)

type SkyWalletHelper struct {
	di hardware_wallet.DeviceInteraction
}

func NewSkyWalletHelper() hardware_wallet.DeviceHelper{
	return &SkyWalletHelper{di: NewSkyWalletInteraction()}
}

func (dev *SkyWalletHelper) FirstAddress(walletType string) *promise.Promise {
	prm := dev.di.AddressGen(1, 0, false, walletType)
	return prm.Then(func(data interface{}) interface{} {
		addresses := data.([]string)
		return addresses[0]
	}).Catch(func(err error) error {
		return err
	})
}

func (dev *SkyWalletHelper) DeviceMatch(wlt core.Wallet) *promise.Promise {
	return dev.FirstAddress(skyWallet.WalletTypeDeterministic).Then(func(data interface{}) interface{} {
		checkForDerivation := func(dt string) bool {
			firstAddr := data.(string)
			addrs := wlt.GenAddresses(core.AccountAddress, 0, 1, nil)
			if addrs.Next() {
				addr := addrs.Value()
				return addr.String() == firstAddr
			}
			return false
		}
		if checkForDerivation(skyWallet.WalletTypeDeterministic) {
			return true
		}
		return checkForDerivation(skyWallet.WalletTypeBip44)
	}).Catch(func(err error) error {
		return err
	})
}
