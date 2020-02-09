package ethereum

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/ethereum/config"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

const (
	EtherName    = "ether"
	EtherTicker  = "eth"
	EthereFamily = "ethereum"
)

type EthereumPlugin struct {
}

func (p *EthereumPlugin) ListSupportedAltcoins() []core.AltcoinMetadata {
	return []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:     EtherName,
			Ticker:   EtherTicker,
			Family:   EthereFamily,
			HasBip44: false,
			Accuracy: 18,
		},
	}
}

func (p *EthereumPlugin) ListSupportedFamilies() []string {
	return []string{EthereFamily}
}

func (p *EthereumPlugin) RegisterTo(manager core.AltcoinManager) {
	for _, info := range p.ListSupportedAltcoins() {
		manager.RegisterAltcoin(info, p)
	}
}

func (p *EthereumPlugin) GetName() string {
	return "Ethereum"
}

func (p *EthereumPlugin) GetDescription() string {
	return "FiberCrypto wallet connector for Ethereum"
}

func (p *EthereumPlugin) LoadWalletEnvs() []core.WalletEnv {
	wltSources, err := config.GetWalletSources()
	if err != nil {
		return nil
	}
	wltEnvs := make([]core.WalletEnv, 0)

	for _, wltS := range wltSources {
		tp := wltS.Tp
		source := wltS.Source
		if tp == string(config.LocalWallet) {
			wltEnvs = append(wltEnv, NewWalletDirectory(source))
		}
	}
	return wltEnvs
}

//TODO waiting for PEX implementation in ethereum plugin
func (p *EthereumPlugin) LoadPEX(netType string) (core.PEX, error) {
	return nil, nil
}

//TODO waiting for Address implementation in ethereum plugin
func (p *EthereumPlugin) AddressFromString(addrStr string) (core.Address, error) {
	return nil, nil
}

//TODO waiting for PubKey implementation in ethereum plugin
func (p *EthereumPlugin) PubKeyFromBytes(b []byte) (core.PubKey, error) {
	return nil, nil
}

//TODO waiting for SecKey implementation in ethereum plugin
func (p *EthereumPlugin) SecKeyFromBytes(b []byte) (core.SecKey, error) {
	return nil, nil
}

func NewEthereumPlugin() core.AltcoinPlugin {
	return &EthereumPlugin{}
}
