package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

	obj, err := pool.Get()

	if err != nil {
		for err == errors.ErrObjectPoolUndeflow {
			obj, err = pool.Get()
		}
		if err != nil {
			return nil, err
		}
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

type EthereumConnectionFactory struct {
	url string
}

func (cf *EthereumConnectionFactory) Create() (interface{}, error) {
	return NewEthereumClientExtended(cf.url)
}

type ethereumClientExtended struct {
	*ethclient.Client
	c *rpc.Client
}

func (clt *ethereumClientExtended) ProtocolVersion(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := clt.c.CallContext(ctx, &result, "eth_protocolVersion")
	if err != nil {
		return nil, err
	}

	return (*big.Int)(&result), err
}

func NewEthereumClientExtended(url string) (*ethereumClientExtended, error) {
	clt, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	c, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &ethereumClientExtended{
		Client: clt,
		c:      c,
	}, nil

}

func NewEthereumConnectionFactory(url string) *EthereumConnectionFactory {

	return &EthereumConnectionFactory{
		url: url,
	}
}
