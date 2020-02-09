package ethereum

import (
	"encoding/json"

	"github.com/fibercrypto/fibercryptowallet/src/util"

	"github.com/fibercrypto/fibercryptowallet/src/coin/ethereum/config"
	eth "github.com/fibercrypto/fibercryptowallet/src/coin/ethereum/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logEthereum = logging.MustGetLogger("Ethereum Altcoin")

func init() {

	err := config.RegisterConfig()
	if err != nil {
		logEthereum.WithError(err).Warn("Couldn't register Ethereum configuration")
	}

	nodeStr, err := config.GetOption(config.SettingPathToNode)
	if err != nil {
		logEthereum.WithError(err).Warn("Couldn't get node option")
	}

	node := make(map[string]string)
	err = json.Unmarshal([]byte(nodeStr), &node)
	if err != nil {
		logEthereum.WithError(err).Warn("Couldn't unmarshal option values")
	}

	err = core.GetMultiPool().CreateSection(eth.PoolSection, eth.NewEthereumConnectionFactory(node["address"]))
	if err != nil {
		logEthereum.Warn("Couldn't create section for Ethereum")
	}

	util.RegisterAltcoin(eth.NewEthereumPlugin())

}
