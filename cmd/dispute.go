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

func (utilsStruct UtilsStruct) HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32) error {
	sortedProposedBlockIds, err := utilsStruct.razorUtils.GetSortedProposedBlockIds(client, account.Address, epoch)
	if err != nil {
		return err
	}
	log.Debug("SortedProposedBlockIds: ", sortedProposedBlockIds)

	biggestInfluence, biggestInfluenceId, err := utilsStruct.proposeUtils.getBiggestInfluenceAndId(client, account.Address, epoch, utilsStruct)
	if err != nil {
		return err
	}
	log.Debug("Biggest Influence: ", biggestInfluence)

	medians, err := utilsStruct.proposeUtils.MakeBlock(client, account.Address, false, utilsStruct)
	if err != nil {
		return err
	}
	log.Debug("Locally calculated data:")
	log.Debugf("Medians: %d", medians)

	for i := 0; i < len(sortedProposedBlockIds); i++ {
		blockId := sortedProposedBlockIds[i]
		proposedBlock, err := utilsStruct.razorUtils.GetProposedBlock(client, account.Address, epoch, blockId)
		if err != nil {
			log.Error(err)
			continue
		}
		if proposedBlock.BiggestInfluence.Cmp(biggestInfluence) != 0 && proposedBlock.Valid {
			log.Debug("Biggest Influence in proposed block: ", proposedBlock.BiggestInfluence)
			log.Warn("PROPOSED BIGGEST INFLUENCE DOES NOT MATCH WITH ACTUAL BIGGEST INFLUENCE")
			log.Info("Disputing BiggestInfluenceProposed...")
			txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
				Client:         client,
				Password:       account.Password,
				AccountAddress: account.Address,
				ChainId:        core.ChainId,
				Config:         config,
			})
			DisputeBiggestInfluenceProposedTxn, err := utilsStruct.blockManagerUtils.DisputeBiggestInfluenceProposed(client, txnOpts, epoch, uint8(i), biggestInfluenceId)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(DisputeBiggestInfluenceProposedTxn))
			status := utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(DisputeBiggestInfluenceProposedTxn).String())
			if status == 1 {
				continue
			}
		}

		log.Debug("Values in the block")
		log.Debugf("Medians: %d", proposedBlock.Medians)

		isEqual, j := utilsStruct.razorUtils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			activeAssetIds, _ := utilsStruct.razorUtils.GetActiveAssetIds(client, account.Address)
			assetId := int(activeAssetIds[j])
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
			if proposedBlock.Valid {
				err := utilsStruct.cmdUtils.Dispute(client, config, account, epoch, uint8(i), assetId, utilsStruct)
				if err != nil {
					log.Error("Error in disputing...", err)
					continue
				}
			} else {
				log.Info("Block already disputed")
				continue
			}
		} else {
			log.Info("Proposed median matches with local calculations. Will not open dispute.")
			continue
		}
	}
	giveSortedAssetIds = []int{}
	return nil
}

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, utilsStruct UtilsStruct) error {
	blockManager := utilsStruct.razorUtils.GetBlockManager(client)
	numOfStakers, err := utilsStruct.razorUtils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		return err
	}

	var sortedStakers []uint32

	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := utilsStruct.razorUtils.GetVotes(client, uint32(i))
		if err != nil {
			return err
		}
		if votes.Epoch == epoch {
			sortedStakers = append(sortedStakers, uint32(i))
		}
	}

	log.Debugf("Epoch: %d, StakerId's who voted: %d", epoch, sortedStakers)
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !razorUtils.Contains(giveSortedAssetIds, assetId) {
		utilsStruct.cmdUtils.GiveSorted(client, blockManager, txnOpts, epoch, uint8(assetId), sortedStakers)
	}

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	finalizeTxn, err := utilsStruct.blockManagerUtils.FinalizeDispute(client, finalizeDisputeTxnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(finalizeTxn))
	utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(finalizeTxn).String())
	return nil
}

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	txn, err := blockManager.GiveSorted(txnOpts, epoch, assetId, sortedStakers)
	if err != nil {
		if err.Error() == errors.New("gas limit reached").Error() {
			log.Error("Error in calling GiveSorted: ", err)
			mid := len(sortedStakers) / 2
			GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers[:mid])
			GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers[mid:])
		} else {
			return
		}
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", txn.Hash())
	giveSortedAssetIds = append(giveSortedAssetIds, int(assetId))
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}
