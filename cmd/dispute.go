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

func (*UtilsStructMockery) HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32) error {
	sortedProposedBlockIds, err := razorUtilsMockery.GetSortedProposedBlockIds(client, epoch)
	if err != nil {
		return err
	}
	log.Debug("SortedProposedBlockIds: ", sortedProposedBlockIds)

	biggestInfluence, biggestInfluenceId, err := cmdUtilsMockery.GetBiggestInfluenceAndId(client, account.Address, epoch)
	if err != nil {
		return err
	}
	log.Debug("Biggest Influence: ", biggestInfluence)

	medians, err := cmdUtilsMockery.MakeBlock(client, account.Address, types.Rogue{IsRogue: false})
	if err != nil {
		return err
	}
	log.Debug("Locally calculated data:")
	log.Debugf("Medians: %d", medians)

	for i := 0; i < len(sortedProposedBlockIds); i++ {
		blockId := sortedProposedBlockIds[i]
		proposedBlock, err := razorUtilsMockery.GetProposedBlock(client, epoch, blockId)
		if err != nil {
			log.Error(err)
			continue
		}
		if proposedBlock.BiggestInfluence.Cmp(biggestInfluence) != 0 && proposedBlock.Valid {
			log.Debug("Biggest Influence in proposed block: ", proposedBlock.BiggestInfluence)
			log.Warn("PROPOSED BIGGEST INFLUENCE DOES NOT MATCH WITH ACTUAL BIGGEST INFLUENCE")
			log.Info("Disputing BiggestInfluenceProposed...")
			txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
				Client:         client,
				Password:       account.Password,
				AccountAddress: account.Address,
				ChainId:        core.ChainId,
				Config:         config,
			})
			DisputeBiggestInfluenceProposedTxn, err := blockManagerUtilsMockery.DisputeBiggestInfluenceProposed(client, txnOpts, epoch, uint8(i), biggestInfluenceId)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Info("Txn Hash: ", transactionUtilsMockery.Hash(DisputeBiggestInfluenceProposedTxn))
			status := razorUtilsMockery.WaitForBlockCompletion(client, transactionUtilsMockery.Hash(DisputeBiggestInfluenceProposedTxn).String())
			if status == 1 {
				continue
			}
		}

		log.Debug("Values in the block")
		log.Debugf("Medians: %d", proposedBlock.Medians)

		isEqual, j := utils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			activeAssetIds, _ := razorUtilsMockery.GetActiveAssetIds(client)
			assetId := int(activeAssetIds[j])
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
			if proposedBlock.Valid {
				err := cmdUtilsMockery.Dispute(client, config, account, epoch, uint8(i), assetId)
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

func (*UtilsStructMockery) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int) error {
	blockManager := razorUtilsMockery.GetBlockManager(client)
	numOfStakers, err := razorUtilsMockery.GetNumberOfStakers(client, account.Address)
	if err != nil {
		return err
	}

	var sortedStakers []uint32

	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := razorUtilsMockery.GetVotes(client, uint32(i))
		if err != nil {
			return err
		}
		if votes.Epoch == epoch {
			sortedStakers = append(sortedStakers, uint32(i))
		}
	}

	log.Debugf("Epoch: %d, StakerId's who voted: %d", epoch, sortedStakers)
	txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !utils.Contains(giveSortedAssetIds, assetId) {
		cmdUtilsMockery.GiveSorted(client, blockManager, txnOpts, epoch, uint16(assetId), sortedStakers)
	}

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	finalizeTxn, err := blockManagerUtilsMockery.FinalizeDispute(client, finalizeDisputeTxnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(finalizeTxn))
	razorUtilsMockery.WaitForBlockCompletion(client, transactionUtilsMockery.Hash(finalizeTxn).String())
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
