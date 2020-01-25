package ethtypes

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumApi interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error)
	ProtocolVersion(ctx context.Context) (*big.Int, error)
}
