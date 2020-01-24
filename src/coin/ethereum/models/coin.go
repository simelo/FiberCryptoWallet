package ethereum

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logCoin = logging.MustGetLogger("Ethereum coin")

const (
	eth = "eth"
)

type EthereumTransaction struct {
	txn       *types.Transaction
	fee       uint64
	timestamp core.Timestamp
	status    core.TransactionStatus
}

func (ethTxn *EthereumTransaction) SupportedAssets() []string {
	logCoin.Info("Gettgin supported assets")
	return []string{eth}

}

func (ethTxn *EthereumTransaction) GetTimestamp() core.Timestamp {
	logCoin.Info("Getting timestamp")
	return ethTxn.timestamp
}

func (ethTxn *EthereumTransaction) GetStatus() core.TransactionStatus {
	logCoin.Info("Getting status")
	return ethTxn.status
}

// TODO
func (ethTxn *EthereumTransaction) GetInputs() []core.TransactionInput {
	logCoin.Info("Getting inputs")
	return nil
}

// TODO
func (ethTxn *EthereumTransaction) GetOutputs() []core.TransactionOutput {
	logCoin.Info("Getting outputs")
	return nil
}

func (ethTxn *EthereumTransaction) GetId() string {
	logCoin.Info("Getting id")
	return ethTxn.txn.Hash().String()
}

func (ethTxn *EthereumTransaction) ComputeFee(ticker string) (uint64, error) {
	logCoin.Info("Computing fee")
	if ticker != eth {
		return 0, errors.ErrInvalidAltcoinTicker
	}
	return ethTxn.fee, nil
}

//TODO
func (ethTxn *EthereumTransaction) VerifyUnsigned() error {
	return nil
}

//TODO
func (ethTxn *EthereumTransaction) VerifySigned() error {
	return nil
}

//TODO
func (ethTxn *EthereumTransaction) IsFullySigned() (bool, error) {
	return false, nil
}
