package ethereum

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/ethereum/ethtypes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logNetwork = logging.MustGetLogger("Ethereum network")

const (
	PoolSection = "ethereum"
)

type EthereumApiClient struct {
	ethtypes.EthereumApi
	pool core.MultiPoolSection
}

func NewEthereumApiClient(section string) (ethtypes.EthereumApi, error) {
	logNetwork.Info("Creating Ethereum api client")
	mpool := core.GetMultiPool()
	pool, err := mpool.GetSection(section)
	if err != nil {
		return nil, err
	}

	obj := pool.Get()
	if err != nil {
		for _, ok := err.(core.NotAvailableObjectsError); ok; _, ok = err.(core.NotAvailableObjectsError) {
			if err == nil {
				break
			}
		}
		return nil, err
	}

	ethApi, ok := obj.(ethtypes.EthereumApi)
	if !ok {
		logNetwork.Errorf("There is no proper client in %s pool", section)
		return nil, errors.ErrInvalidPoolObjectType
	}
	return &EthereumApiClient{
		EthereumApi: ethApi,
		pool:        pool,
	}, nil
}

func ReturnEthereumApiClient(obj ethtypes.EthereumApi) {
	poolObj, ok := obj.(*EthereumApiClient)
	if !ok {
		return
	}
	poolObj.pool.Put(poolObj.EthereumApi)
}
