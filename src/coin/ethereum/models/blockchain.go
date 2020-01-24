package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

type EthereumBlockchainInfo struct {
}

type EthereumBlockchain struct { //Implements BlockchainStatus interface
	lastTimeStatusRequested uint64
	CacheTime               uint64
	cachedStatus            *EthereumBlockchainInfo
}

func (eb *EthereumBlockchain) GetCoinValue(coinvalue core.CoinValueMetric, ticker string) (uint64, error) {
	return 0, nil
}

func (eb *EthereumBlockchain) GetLastBlock() (core.Block, error) {
	clt, err := NewEthereumApiClient("default")
	nCtx := context.Background()
	ethBlk, err := clt.BlockByNumber(nCtx, nil)
	if err != nil {
		return nil, err
	}

}

func (eb *EthereumBlockchain) GetNumberOfBlocks() (uint64, error) {
	return 0, nil
}

func (eb *EthereumBlockchain) GetBlockByHash(hash string) (core.Block, error) {
	return nil, nil
}

func (eb *EthereumBlockchain) GetRangeBlocks(start, end uint64) ([]core.Block, error) {
	return nil, nil
}

type EthereumBlock struct {
	ethBlock *types.Block
	version  int
}

func (eb *EthereumBlock) GetHash() ([]byte, error) {
	return eb.ethBlock.Hash().Bytes(), nil
}

func (eb *EthereumBlock) GetPrevHash() ([]byte, error) {
	return eb.ethBlock.ParentHash().Bytes(), nil
}

//TODO look for right values, there is no version field in ethereum blocks
func (eb *EthereumBlock) GetVersion() (uint32, error) {
	return 0, nil
}

func (eb *EthereumBlock) GetTime() (core.Timestamp, error) {
	tm := core.Timestamp(eb.ethBlock.Time())
	return tm, nil
}

func (eb *EthereumBlock) GetHeight() (uint64, error) {
	return eb.ethBlock.NumberU64(), nil
}

//TODO learn how the fee is calculated in eth
func (eb *EthereumBlock) GetFee(ticker string) (uint64, error) {
	return 0, nil
}

func (eb *EthereumBlock) GetSize() (uint64, error) {
	size := uint64(eb.ethBlock.Size())
	return size, nil
}

func (eb *EthereumBlock) IsGenesisBlock() (bool, error) {
	return eb.ethBlock.NumberU64() == 0, nil
}

func (eb *EthereumBlock) GetTotalAmount() (uint64, error) {
	var total uint64 = 0
	for _, txn := range eb.ethBlock.Transactions() {
		total += txn.Value().Uint64()
	}
	return total, nil
}

func (eb *EthereumBlock) GetTransactions() ([]Transaction, error)
