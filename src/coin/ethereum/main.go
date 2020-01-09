package ethereum

import (
	eth "github.com/fibercrypto/fibercryptowallet/src/coin/ethereum/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logEthereum = logging.MustGetLogger("Ethereum Altcoin")

func init() {

	err := core.GetMultiPool().CreateSection(eth.PoolSection, eth.NewEthereumConnectionFactory("defaul"))
	if err != nil {
		logEthereum.Warn("Couldn't create section for Ethereum")
	}

}
