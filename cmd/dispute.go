package cmd

import (
	"errors"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var giveSortedAssetIds []int

func (*UtilsStruct) HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32) error {
	sortedProposedBlockIds, err := razorUtils.GetSortedProposedBlockIds(client, epoch)
	if err != nil {
		return err
	}
	log.Debug("SortedProposedBlockIds: ", sortedProposedBlockIds)

	biggestStake, biggestStakerId, err := cmdUtils.GetBiggestStakeAndId(client, account.Address, epoch)
	if err != nil {
		return err
	}
	log.Debug("Biggest Stake: ", biggestStake)

	medians, err := cmdUtils.MakeBlock(client, account.Address, types.Rogue{IsRogue: false})
	if err != nil {
		return err
	}
	log.Debug("Locally calculated data:")
	log.Debugf("Medians: %d", medians)

	randomSortedProposedBlockIds := rand.Perm(len(sortedProposedBlockIds)) //returns random permutation of integers from 0 to n-1
	for _, i := range randomSortedProposedBlockIds {
		blockId := sortedProposedBlockIds[i]
		proposedBlock, err := razorUtils.GetProposedBlock(client, epoch, blockId)
		if err != nil {
			log.Error(err)
			continue
		}
		if proposedBlock.BiggestStake.Cmp(biggestStake) != 0 && proposedBlock.Valid {
			log.Debug("Biggest Stake in proposed block: ", proposedBlock.BiggestStake)
			log.Warn("PROPOSED BIGGEST STAKE DOES NOT MATCH WITH ACTUAL BIGGEST STAKE")
			log.Info("Disputing BiggestStakeProposed...")
			txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
				Client:         client,
				Password:       account.Password,
				AccountAddress: account.Address,
				ChainId:        core.ChainId,
				Config:         config,
			})
			DisputeBiggestStakeProposedTxn, err := blockManagerUtils.DisputeBiggestStakeProposed(client, txnOpts, epoch, uint8(i), biggestStakerId)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Info("Txn Hash: ", transactionUtils.Hash(DisputeBiggestStakeProposedTxn))
			status := razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(DisputeBiggestStakeProposedTxn).String())
			if status == 1 {
				continue
			}
		}

		log.Debug("Values in the block")
		log.Debugf("Medians: %d", proposedBlock.Medians)

		isEqual, j := utils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			activeAssetIds, _ := razorUtils.GetActiveAssetIds(client)
			assetId := int(activeAssetIds[j])
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
			if proposedBlock.Valid {
				err := cmdUtils.Dispute(client, config, account, epoch, uint8(i), assetId)
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

func (*UtilsStruct) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int) error {
	blockManager := razorUtils.GetBlockManager(client)
	numOfStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		return err
	}

	var sortedStakers []uint32

	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := razorUtils.GetVotes(client, uint32(i))
		if err != nil {
			return err
		}
		if votes.Epoch == epoch {
			sortedStakers = append(sortedStakers, uint32(i))
		}
	}

	log.Debugf("Epoch: %d, StakerId's who voted: %d", epoch, sortedStakers)
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !utils.Contains(giveSortedAssetIds, assetId) {
		cmdUtils.GiveSorted(client, blockManager, txnOpts, epoch, uint16(assetId), sortedStakers)
	}

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	finalizeTxn, err := blockManagerUtils.FinalizeDispute(client, finalizeDisputeTxnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(finalizeTxn))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(finalizeTxn).String())
	return nil
}

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint16, sortedStakers []uint32) {
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
