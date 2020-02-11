package core

// CoinValueMetric enumerates all possible values of blockchain metrics
type CoinValueMetric uint32

const (
	// CoinCurrentSupply to retrieve amount of coins distributed to tenants
	CoinCurrentSupply CoinValueMetric = iota
	// CoinTotalSupply to retrieve total amount of coins
	CoinTotalSupply
)

// BlockchainStatus measure blockchain metrics in real time
type BlockchainStatus interface {
	// GeCoinValue retrieves value of a blockchain metric
	GetCoinValue(coinvalue CoinValueMetric, ticker string) (uint64, error)
	// GetLastBlock retrieves block at the tip of he block chain
	GetLastBlock() (Block, error)
	// GetNumberOfBlocks determine number of blocks in the blockchain
	GetNumberOfBlocks() (uint64, error)
}

// BlockchainVisor abstract interface for transactions management and utility functions for specific blockchain.
// The service should use the blockchain node to implement given interface.
type BlockchainVisor interface {
	// Transfer instantiates unsigned transaction to send funds from any wallet address to single destination
	Transfer(to TransactionOutput, options KeyValueStore) (Transaction, error)
	// SendFromAddress instantiates unsigned transaction to send funds from specific source addresses
	// to multiple destination addresses
	SendFromAddress(from []Address, to []TransactionOutput, change Address, options KeyValueStore) (Transaction, error)
	// Spend instantiate unsigned transaction spending specific outputs to send to multiple destination addresses
	Spend(unspent, new []TransactionOutput, change Address, options KeyValueStore) (Transaction, error)
	// ScanOutputs scan the blockchain looking for outputs
	ScanOutputs(unspentOnly bool) (TransactionOutputIterator, error)
}
