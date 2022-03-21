package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

var giveSortedAssetIds []int

func (*UtilsStruct) HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) error {

	sortedProposedBlockIds, err := razorUtils.GetSortedProposedBlockIds(client, epoch)
	if err != nil {
		log.Error("Error in fetching sorted proposed block id: ", err)
		return err
	}
	log.Debug("SortedProposedBlockIds: ", sortedProposedBlockIds)

	biggestStake, biggestStakerId, err := cmdUtils.GetBiggestStakeAndId(client, account.Address, epoch)
	if err != nil {
		return err
	}
	log.Debug("Biggest Stake: ", biggestStake)

	medians, revealedCollectionIds, err := cmdUtils.GetLocalMediansData(client, account, epoch, blockNumber, rogueData)
	if err != nil {
		return err
	}

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
			disputeBiggestStakeProposedTxn, err := blockManagerUtils.DisputeBiggestStakeProposed(client, txnOpts, epoch, uint8(i), biggestStakerId)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Info("Txn Hash: ", transactionUtils.Hash(disputeBiggestStakeProposedTxn))
			status := razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(disputeBiggestStakeProposedTxn).String())
			if status == 1 {
				continue
			}
		}
		log.Debug("Locally revealed collection ids: ", revealedCollectionIds)
		log.Debug("Revealed collection ids in the block ", proposedBlock.Ids)

		isRevealedIdsMatched, idsMismatchIndex := cmdUtils.CheckDisputeTypeForIds(proposedBlock.Ids, revealedCollectionIds)

		// Check if the error is collectionIdShouldBeAbsent

		log.Debug("Locally calculated medians: ", medians)
		log.Debug("Medians in the block: ", proposedBlock.Medians)

		isEqual, j := utils.IsEqualUint32(proposedBlock.Medians, medians)
		if !isEqual {
			activeAssetIds, _ := razorUtils.GetActiveCollections(client)
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

func (*UtilsStruct) GetLocalMediansData(client *ethclient.Client, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) ([]uint32, []uint16, error) {
	//TODO: Implement proper file reading and writing

	//	if _mediansData == nil && !rogueData.IsRogue {
	//		fileName, err := cmdUtils.GetMedianDataFileName(account.Address)
	//		if err != nil {
	//			log.Error("Error in getting file name to read median data: ", err)
	//			goto CalculateMedian
	//		}
	//		epochInFile, medianDataFromFile, err := razorUtils.ReadDataFromFile(fileName)
	//		if err != nil {
	//			log.Errorf("Error in getting median data from file %s: %t", fileName, err)
	//			goto CalculateMedian
	//		}
	//		if epochInFile != epoch {
	//			log.Errorf("File %s doesn't contain latest median data: %t", fileName, err)
	//			goto CalculateMedian
	//		}
	//		_mediansData = medianDataFromFile
	//	}
	//CalculateMedian:
	if _mediansData == nil || _revealedCollectionIds == nil {
		medians, revealedCollectionIds, err := cmdUtils.MakeBlock(client, blockNumber, epoch, types.Rogue{IsRogue: false})
		if err != nil {
			log.Error("Error in calculating block medians")
			return nil, nil, err
		}
		_mediansData = razorUtils.ConvertUint32ArrayToBigIntArray(medians)
		_revealedCollectionIds = revealedCollectionIds
	}

	mediansInUint32 := razorUtils.ConvertBigIntArrayToUint32Array(_mediansData)
	log.Debug("Locally calculated data:")
	log.Debugf("Medians: %d", mediansInUint32)
	return mediansInUint32, _revealedCollectionIds, nil
}

func (*UtilsStruct) CheckDisputeTypeForIds(client *ethclient.Client, transactionOpts types.TransactionOptions, epoch uint32, blockIndex uint8, idsInProposedBlock []uint16, revealedCollectionIds []uint16) (*types2.Transaction, error) {
	// Check if the error is in sorted ids
	isSorted, index0, index1 := utils.IsSorted(idsInProposedBlock)
	if !isSorted {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeOnOrderOfIds"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, index0, index1}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		return blockManagerUtils.DisputeOnOrderOfIds(client, txnOpts, epoch, blockIndex, big.NewInt(int64(index0)), big.NewInt(int64(index1)))
	}

	// Check if the error is collectionIdShouldBePresent
	isMissing, _, missingValue := utils.IsMissing(revealedCollectionIds, idsInProposedBlock)
	if isMissing {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBePresent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, missingValue}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		return blockManagerUtils.DisputeCollectionIdShouldBePresent(client, txnOpts, epoch, blockIndex, missingValue)
	}

	// Check if the error is collectionIdShouldBeAbsent
	isPresent, positionOfPresentValue, presentValue := utils.IsMissing(idsInProposedBlock, revealedCollectionIds)
	if isPresent {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBeAbsent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, presentValue, big.NewInt(int64(positionOfPresentValue))}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		return blockManagerUtils.DisputeCollectionIdShouldBeAbsent(client, txnOpts, epoch, blockIndex, presentValue, big.NewInt(int64(positionOfPresentValue)))
	}

	return nil, nil
}

func (*UtilsStruct) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int) error {
	blockManager := razorUtils.GetBlockManager(client)
	numOfStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		return err
	}

	var sortedStakers []uint32
	//TODO: How to get sorted values
	for i := 1; i <= int(numOfStakers); i++ {
		votes, err := razorUtils.GetVoteValue(client, epoch, uint32(i))
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

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint16, sortedValues []uint32) {
	if len(sortedValues) == 0 {
		return
	}
	txn, err := blockManagerUtils.GiveSorted(blockManager, txnOpts, epoch, assetId, sortedValues)
	if err != nil {
		if err.Error() == errors.New("gas limit reached").Error() {
			log.Error("Error in calling GiveSorted: ", err)
			mid := len(sortedValues) / 2
			GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedValues[:mid])
			GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedValues[mid:])
		} else {
			return
		}
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	giveSortedAssetIds = append(giveSortedAssetIds, int(assetId))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(txn).String())
}
