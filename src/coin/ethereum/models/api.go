package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

type EthereumApiMock struct {
	mock.Mock
}

func (m *EthereumApiMock) Create() (interface{}, error) {
	return m, nil
}

func (m *EthereumApiMock) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	args := m.Called(ctx, hash)
	return args.Get(0).(*types.Block), args.Error(1)
}

func (m *EthereumApiMock) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	args := m.Called(ctx, number)
	return args.Get(0).(*types.Block), args.Error(1)
}

func (m *EthereumApiMock) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	args := m.Called(ctx, hash)
	return args.Get(0).(*types.Receipt), args.Error(1)
}

func (m *EthereumApiMock) ProtocolVersion(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}
