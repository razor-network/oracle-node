package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var giveSortedAssetIds []int

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
		log.Debug("Values in the block")
		log.Debugf("Medians: %d", proposedBlock.Medians)
		medians, err := MakeBlock(client, account.Address, false)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("Locally calculated data:")
		log.Debugf("Medians: %d\n", medians)
		isEqual, assetId := utils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
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
	giveSortedAssetIds = []int{}
}

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int) error {
	blockManager := utils.GetBlockManager(client)
	sortedVotes, err := getSortedVotes(client, account.Address, uint8(assetId), epoch)
	if err != nil {
		return err
	}

	log.Debugf("Epoch: %d, Sorted Votes: %s", epoch, sortedVotes)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !utils.Contains(giveSortedAssetIds, assetId) {
		GiveSorted(client, blockManager, txnOpts, epoch, uint8(assetId), utils.ConvertBigIntArrayToUint32Array(sortedVotes))
	}

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
	giveSortedAssetIds = append(giveSortedAssetIds, int(assetId))
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}
