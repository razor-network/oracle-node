package cmd

import (
	"errors"
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
		activeAssetIds, _ := utils.GetActiveAssetIds(client, account.Address, epoch)

		isEqual, i := utils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			assetId := int(activeAssetIds[i])
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
	numOfStakers, err := utils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		return err
	}

	var sortedStakers []uint32

	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := utils.GetVotes(client, account.Address, uint32(i))
		if err != nil {
			return err
		}
		if votes.Epoch == epoch {
			sortedStakers = append(sortedStakers, uint32(i))
		}
	}

	log.Debugf("Epoch: %d, StakerId's who voted: %d", epoch, sortedStakers)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !utils.Contains(giveSortedAssetIds, assetId) {
		GiveSorted(client, blockManager, txnOpts, epoch, uint8(assetId), sortedStakers)
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

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	txn, err := blockManager.GiveSorted(txnOpts, epoch, assetId, sortedStakers)
	if err != nil && err.Error() != errors.New("gas limit reached").Error() {
		return
	}
	if err != nil && err.Error() == errors.New("gas limit reached").Error() {
		log.Error("Error in calling GiveSorted: ", err)
		mid := len(sortedStakers) / 2
		GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers[:mid])
		GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers[mid:])
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", txn.Hash())
	giveSortedAssetIds = append(giveSortedAssetIds, int(assetId))
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}
