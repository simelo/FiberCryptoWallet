package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func mockEthApiBlockByNumber(mock *EthereumApiMock, ctx context.Context, number *big.Int, blk *types.Block, err error) {
	mock.On("BlockByNumber", ctx, number).Return(blk, err)
}

func mockEthApiProtocolVersion(mock *EthereumApiMock, ctx context.Context, version *big.Int, err error) {
	mock.On("ProtocolVersion", ctx).Return(version, err)
}
