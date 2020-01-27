package ethereum

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logBlockchain = logging.MustGetLogger("Skycoin Blockchain")

func NewEthereumBlockchain(invalidCacheTime uint64) *EthereumBlockchain {
	return &EthereumBlockchain{
		lastTimeStatusRequestedLastBlock:     0,
		lastTimeStatusRequestedNumberOfBlock: 0,
		CacheTime:                            invalidCacheTime,
		lastBlock:                            nil,
		numberOfBlocks:                       0,
	}
}

type EthereumBlockchain struct { //Implements BlockchainStatus interface
	lastTimeStatusRequestedNumberOfBlock uint64
	lastTimeStatusRequestedLastBlock     uint64
	CacheTime                            uint64
	lastBlock                            core.Block
	numberOfBlocks                       uint64
}

func (eb *EthereumBlockchain) numberOfBlocksCachedIsValid() bool {
	elapsed := uint64(time.Now().UTC().UnixNano()) - eb.lastTimeStatusRequestedNumberOfBlock
	return elapsed < eb.CacheTime
}

func (eb *EthereumBlockchain) lastBlockCachedIsValid() bool {
	elapsed := uint64(time.Now().UTC().UnixNano()) - eb.lastTimeStatusRequestedLastBlock
	return elapsed < eb.CacheTime
}

// TODO
func (eb *EthereumBlockchain) GetCoinValue(coinvalue core.CoinValueMetric, ticker string) (uint64, error) {
	return 0, nil
}

func (eb *EthereumBlockchain) GetLastBlock() (core.Block, error) {
	logBlockchain.Info("Getting last block")
	if eb.lastBlockCachedIsValid() {
		return eb.lastBlock, nil

	}
	clt, err := NewEthereumApiClient(PoolSection)
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting last block")
		return nil, err
	}
	nCtx := context.Background()
	ethBlk, err := clt.BlockByNumber(nCtx, nil)
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting last block")
		return nil, err
	}
	//version, err := clt.ProtocolVersion(nCtx)
	//if err != nil {
	//	logBlockchain.WithError(err).Error("Error getting last block")
	//	return nil, err
	//}
	eb.lastTimeStatusRequestedLastBlock = uint64(time.Now().UTC().UnixNano())
	eb.lastBlock = NewEthereumBlock(ethBlk, 0 /*uint32(version.Uint64())*/)
	return eb.lastBlock, nil
}

func (eb *EthereumBlockchain) GetNumberOfBlocks() (uint64, error) {
	logBlockchain.Infof("Getting number of blocks")
	if eb.numberOfBlocksCachedIsValid() {
		return eb.numberOfBlocks, nil
	}

	lastBlock, err := eb.GetLastBlock()
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting number of blocks")
		return 0, err
	}
	heigth, err := lastBlock.GetHeight()
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting number of blocks")
		return 0, nil
	}
	eb.lastTimeStatusRequestedNumberOfBlock = uint64(time.Now().UTC().UnixNano())
	eb.numberOfBlocks = (heigth + 1)
	return eb.numberOfBlocks, nil
}

func (eb *EthereumBlockchain) GetBlockByHash(hash string) (core.Block, error) {
	logBlockchain.Info("Getting block by hash")
	clt, err := NewEthereumApiClient(PoolSection)
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting block by hash")
		return nil, err
	}
	nCtx := context.Background()
	blk, err := clt.BlockByHash(nCtx, common.HexToHash(hash))
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting block by hash")
		return nil, err
	}

	//version, err := clt.ProtocolVersion(nCtx)
	//if err != nil {
	//	logBlockchain.WithError(err).Error("Error getting block by hash")
	//	return nil, err
	//}
	return NewEthereumBlock(blk, 0 /*uint32(version.Uint64())*/), nil
}

func (eb *EthereumBlockchain) GetRangeBlocks(start, end uint64) ([]core.Block, error) {
	logBlockchain.Info("Getting block range")
	if end < start {
		logBlockchain.WithError(errors.ErrInvalidRange).Error("Error getting block range")
		return nil, errors.ErrInvalidRange
	}
	clt, err := NewEthereumApiClient(PoolSection)
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting block range")
		return nil, err
	}
	nCtx := context.Background()
	//version, err := clt.ProtocolVersion(nCtx)
	//if err != nil {
	//	return nil, err
	//}
	answ := make([]core.Block, 0)
	for i := start; i <= end; i++ {
		nCtx = context.Background()
		blk, err := clt.BlockByNumber(nCtx, new(big.Int).SetUint64(i))
		if err != nil {
			logBlockchain.WithError(err).Error("Error getting block")
			continue
		}
		answ = append(answ, NewEthereumBlock(blk, 0 /*uint32(version.Uint64())*/))
	}
	return answ, nil
}

func NewEthereumBlock(blk *types.Block, version uint32) *EthereumBlock {
	return &EthereumBlock{
		version:  version,
		ethBlock: blk,
	}
}

type EthereumBlock struct {
	ethBlock *types.Block
	version  uint32
	ethTxns  []core.Transaction
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
	if ticker != eth {
		return 0, errors.ErrInvalidAltcoinTicker
	}
	txns, err := eb.GetTransactions()
	if err != nil {
		logBlockchain.WithError(err).Error("Error getting fee")
		return 0, err
	}
	var fee uint64 = 0
	for _, txn := range txns {
		txnFee, err := txn.ComputeFee(eth)
		if err != nil {
			logBlockchain.WithError(err).Error("Error getting fee")
			return 0, nil
		}
		fee += txnFee
	}
	return fee, nil
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

func (eb *EthereumBlock) GetTransactions() ([]core.Transaction, error) {
	logBlockchain.Info("Getting transactions")
	if eb.ethTxns == nil {
		clt, err := NewEthereumApiClient(PoolSection)
		if err != nil {
			logBlockchain.WithError(err).Error("Error getting ethereum api client")
			return nil, err
		}
		defer ReturnEthereumApiClient(clt)
		txns := make([]core.Transaction, 0)
		var errorOcurred bool
		for _, txn := range eb.ethBlock.Transactions() {
			nCtx := context.Background()
			rcp, err := clt.TransactionReceipt(nCtx, txn.Hash())
			if err != nil {
				logCoin.WithError(err).Error("Error getting transaction receipt2")
				continue
			}
			fee := (new(big.Int).Mul(txn.GasPrice(), new(big.Int).SetUint64(rcp.GasUsed))).Uint64()

			nTxn := &EthereumTransaction{
				fee:       fee,
				status:    core.TXN_STATUS_CONFIRMED,
				timestamp: core.Timestamp(eb.ethBlock.Time()),
				txn:       txn,
			}
			txns = append(txns, nTxn)
		}
		if !errorOcurred {
			eb.ethTxns = txns
		}
		return txns, nil
	}
	return eb.ethTxns, nil
}
