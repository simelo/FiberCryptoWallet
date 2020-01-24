package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logBlockchain = logging.MustGetLogger("Skycoin Blockchain")

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
	version  uint32
	fee      uint64
}

func (eb *EthereumBlock) GetHash() ([]byte, error) {
	logBlockchain.Info("Getting hash")
	return eb.ethBlock.Hash().Bytes(), nil
}

func (eb *EthereumBlock) GetPrevHash() ([]byte, error) {
	logBlockchain.Info("Getting previous block hash")
	return eb.ethBlock.ParentHash().Bytes(), nil
}

func (eb *EthereumBlock) GetVersion() (uint32, error) {
	logBlockchain.Info("Getting version")
	return eb.version, nil
}

func (eb *EthereumBlock) GetTime() (core.Timestamp, error) {
	logBlockchain.Info("Getting time")
	tm := core.Timestamp(eb.ethBlock.Time())
	return tm, nil
}

func (eb *EthereumBlock) GetHeight() (uint64, error) {
	logBlockchain.Info("Getting heigth")
	return eb.ethBlock.NumberU64(), nil
}

func (eb *EthereumBlock) GetFee(ticker string) (uint64, error) {
	logBlockchain.Info("Getting fee")
	return eb.fee, nil
}

func (eb *EthereumBlock) GetSize() (uint64, error) {
	logBlockchain.Info("Getting size")
	size := uint64(eb.ethBlock.Size())
	return size, nil
}

func (eb *EthereumBlock) IsGenesisBlock() (bool, error) {
	logBlockchain.Info("Getting if is genesis block")
	return eb.ethBlock.NumberU64() == 0, nil
}

func (eb *EthereumBlock) GetTotalAmount() (uint64, error) {
	logBlockchain.Info("Getting total amount")
	var total uint64 = 0
	for _, txn := range eb.ethBlock.Transactions() {
		total += txn.Value().Uint64()
	}
	return total, nil
}

func (eb *EthereumBlock) GetTransactions() ([]Transaction, error) {
	logBlockchain.Info("Getting transactions")
}
