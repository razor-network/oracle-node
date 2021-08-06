package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

func HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch *big.Int) {
	numberOfProposedBlocks, err := utils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}
	for i := 0; i < int(numberOfProposedBlocks.Int64()); i++ {
		proposedBlock, err := utils.GetProposedBlock(client, account.Address, epoch, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Values in the block")
		log.Infof("Medians: %s", proposedBlock.BlockMedians)
		medians, err := MakeBlock(client, account.Address, epoch, false)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Locally calculated data:")
		log.Infof("Medians: %s\n", medians)
		isEqual, assetId := utils.IsEqual(proposedBlock.BlockMedians, medians)
		if !isEqual {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Info("Block Values: ", proposedBlock.BlockMedians)
			log.Info("Local Calculations: ", medians)
			err := Dispute(client, config, account, epoch, big.NewInt(int64(i)), assetId)
			if err != nil {
				log.Error("Error in disputing...", err)
				continue
			}
		} else {
			log.Info("Proposed median matches with local calculations. Will not open dispute.")
			break
		}
	}
}

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch *big.Int, blockId *big.Int, assetId int) error {
	blockManager := utils.GetBlockManager(client)
	_, sortedVotes, err := getSortedVotes(client, account.Address, assetId, epoch)
	if err != nil {
		return err
	}
	log.Infof("Epoch: %s, Sorted Votes: %s", epoch, sortedVotes)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	txn, err := blockManager.GiveSorted(txnOpts, epoch, big.NewInt(int64(assetId - 1)), sortedVotes)
	if err != nil {
		log.Error(err)
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	finalizeTxn, err := blockManager.FinalizeDispute(finalizeDisputeTxnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", finalizeTxn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
	return nil
}
