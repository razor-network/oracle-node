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

func HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, razorUtils utilsInterface, proposeUtils proposeUtilsInterface, cmdUtils utilsCmdInterface, blockManagerUtils blockManagerInterface, transactionUtils transactionInterface) error {
	numberOfProposedBlocks, err := razorUtils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := 0; i < int(numberOfProposedBlocks); i++ {
		proposedBlock, err := razorUtils.GetProposedBlock(client, account.Address, epoch, uint8(i))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("Values in the block")
		log.Debugf("Medians: %d", proposedBlock.Medians)
		medians, err := proposeUtils.MakeBlock(client, account.Address, false, razorUtils, proposeUtils)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("Locally calculated data:")
		log.Debugf("Medians: %d\n", medians)
		activeAssetIds, _ := razorUtils.GetActiveAssetIds(client, account.Address, epoch)

		isEqual, j := razorUtils.IsEqual(proposedBlock.Medians, medians)
		if !isEqual {
			assetId := int(activeAssetIds[j])
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
			if proposedBlock.Valid {
				err := cmdUtils.Dispute(client, config, account, epoch, uint8(i), assetId, razorUtils, cmdUtils, blockManagerUtils, transactionUtils)
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
			break
		}
	}
	giveSortedAssetIds = []int{}
	return nil
}

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, razorUtils utilsInterface, cmdUtils utilsCmdInterface, blockManagerUtils blockManagerInterface, transactionUtils transactionInterface) error {
	blockManager := razorUtils.GetBlockManager(client)
	numOfStakers, err := razorUtils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		return err
	}

	var sortedStakers []uint32

	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := razorUtils.GetVotes(client, account.Address, uint32(i))
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

	if !razorUtils.Contains(giveSortedAssetIds, assetId) {
		cmdUtils.GiveSorted(client, blockManager, txnOpts, epoch, uint8(assetId), sortedStakers)
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
