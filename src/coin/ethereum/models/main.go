package ethereum

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

const (
	EtherName = "ether"
	EtherTicker = "eth"
	EthereFamily = "ethereum"
)

type EthereumPlugin struct{

}
func (p *EthereumPlugin) ListSupportedAltcoins() []core.AltcoinMetadata{
	return []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:EtherName,
			Ticker:EtherTicker,
			Family:EthereFamily,
			HasBip44:false,
			Accuracy: 18,
		}
	}
}

func (p *EthereumPlugin) ListSupportedFamilies() []string{
	return []string{EthereFamily}
}

func (p *EthereumPlugin) RegisterTo(manager core.AltcoinManager){
	for _, info := range p.ListSupportedAltcoins(){
		manager.RegisterAltcoin(info, p)
	}
}

func (p *EthereumPlugin) GetName() string{
	return "Ethereum"
}

func (p *EthereumPlugin) GetDescription() string{
	return "FiberCrypto wallet connector for Ethereum"
}

func (p *EthereumPlugin) LoadWalletEnvs []core.WalletEnv{
	
}