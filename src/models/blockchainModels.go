package models

import (
	"fmt"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models" // callable as skycoin
	"github.com/fibercrypto/fibercryptowallet/src/core"
	modelUtil "github.com/fibercrypto/fibercryptowallet/src/models/util"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	qtcore "github.com/therecipe/qt/core"
)

var logBlockchain = logging.MustGetLogger("modelsBlockchain")

// BlockchainStatusModel Contains info about the blockchain to be show.
type BlockchainStatusModel struct {
	qtcore.QObject
	_ func() `constructor:"init"`

	infoRequester core.BlockchainStatus

	_ string            `property:"numberOfBlocks"`
	_ bool              `property:"loading"`
	_ *qtcore.QDateTime `property:"timestampLastBlock"`
	_ string            `property:"hashLastBlock"`
	_ *modelUtil.Map    `property:"coinOpts"`

	_ func(ticket string) `signal:"update,auto"`
}

func (blockchainStatus *BlockchainStatusModel) init() {
	// block details
	blockchainStatus.SetNumberOfBlocksDefault("0")
	blockchainStatus.SetTimestampLastBlockDefault(qtcore.NewQDateTime3(qtcore.NewQDate3(2000, 1, 1), qtcore.NewQTime3(0, 0, 0, 0), qtcore.Qt__LocalTime))
	blockchainStatus.SetHashLastBlockDefault("")

	blockchainStatus.SetLoading(true)
}

func (blockchainStatus *BlockchainStatusModel) update(ticket string) {
	logBlockchain.Info("Updating blockchain")
	var err error
	blockchainStatus.infoRequester, err = util.LoadBlockchainStatus(ticket)
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't can load the blockchain status")
		return
	}

	if err := blockchainStatus.updateInfo(); err != nil {
		logBlockchain.WithError(err).Warn("Couldn't update blockchain Info")
		return
	}
	return
}

// updateInfo request the needed information
func (blockchainStatus *BlockchainStatusModel) updateInfo() error {
	logBlockchain.Info("Updating Blockchain Status")
	blockchainStatus.SetLoading(true)
	block, err := blockchainStatus.infoRequester.GetLastBlock()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get last block")
		return err
	}

	lastBlockHash, err := block.GetHash()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get the hash of the last block")
		return err
	}
	numberOfBlocks, err := blockchainStatus.infoRequester.GetNumberOfBlocks()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get the number of blocks")
		return err
	}
	timestamp, err := block.GetTime()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get block time")
		return err
	}
	year, month, day, h, m, s := util.ParseDate(int64(timestamp))

	coinOpts := modelUtil.NewMap(nil)
	for _, asset := range blockchainStatus.infoRequester.SupportedAssets() {
		currentSupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinCurrentSupply, skycoin.Sky)
		if err != nil {
			logBlockchain.WithError(err).Warnf("Couldn't get current coin supply of %s", asset)
			return err
		}
		accuracy, err := util.AltcoinQuotient(asset)
		if err != nil {
			logBlockchain.WithError(err).Warn("Couldn't get " + asset + " coins quotient")
		}

		coinOpts.SetValue(fmt.Sprintf("Current %s supply:", asset), util.FormatCoins(currentSupply, accuracy)+" "+asset)

		totalSupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinTotalSupply, skycoin.Sky)
		if err != nil {
			logBlockchain.WithError(err).Warnf("Couldn't get total coin supply of %s", asset)
			return err
		}
		coinOpts.SetValue(fmt.Sprintf("Total %s supply:", asset), util.FormatCoins(totalSupply, accuracy)+" "+asset)
	}

	blockchainStatus.SetCoinOpts(coinOpts)

	// block details
	blockchainStatus.SetNumberOfBlocks(util.FormatCoins(numberOfBlocks, 1))
	blockchainStatus.SetTimestampLastBlock(qtcore.NewQDateTime3(qtcore.NewQDate3(year, month, day), qtcore.NewQTime3(h, m, s, 0), qtcore.Qt__LocalTime))
	blockchainStatus.SetHashLastBlock(string(lastBlockHash))
	blockchainStatus.SetLoading(false)

	return nil
}
