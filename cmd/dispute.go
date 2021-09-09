package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32) {
	numberOfProposedBlocks, err := utils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}
	for i := 0; i < int(numberOfProposedBlocks); i++ {
		proposedBlock, err := utils.GetProposedBlock(client, account.Address, epoch, uint8(i))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Values in the block")
		log.Infof("Medians: %d", proposedBlock.BlockMedians)
		medians, err := MakeBlock(client, account.Address, false)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Locally calculated data:")
		log.Infof("Medians: %d\n", medians)
		isEqual, assetId := utils.IsEqual(proposedBlock.BlockMedians, medians)
		if !isEqual {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Info("Block Values: ", proposedBlock.BlockMedians)
			log.Info("Local Calculations: ", medians)
			err := Dispute(client, config, account, epoch, uint8(i), assetId)
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

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int) error {
	blockManager := utils.GetBlockManager(client)
	sortedVotes, err := getSortedVotes(client, account.Address)
	if err != nil {
		return err
	}

	log.Infof("Epoch: %d, Sorted Votes: %s", epoch, sortedVotes)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	GiveSorted(client, blockManager, txnOpts, epoch, uint8(assetId-1), utils.ConvertBigIntArrayToUint32Array(sortedVotes))

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	finalizeTxn, err := blockManager.FinalizeDispute(finalizeDisputeTxnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", finalizeTxn.Hash())
	utils.WaitForBlockCompletion(client, finalizeTxn.Hash().String())
	return nil
}

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedVotes []uint32) {
	txn, err := blockManager.GiveSorted(txnOpts, epoch, assetId, sortedVotes)
	if err != nil {
		log.Error("Error in calling GiveSorted: ", err)
		mid := len(sortedVotes) / 2
		GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedVotes[:mid])
		GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedVotes[mid:])
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}
