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

var giveSortedLeafIds []int

//blockId is id of the block

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

	medians, revealedCollectionIds, revealedDataMaps, err := cmdUtils.GetLocalMediansData(client, account, epoch, blockNumber, rogueData)
	if err != nil {
		return err
	}

	randomSortedProposedBlockIds := rand.Perm(len(sortedProposedBlockIds)) //returns random permutation of integers from 0 to n-1
	transactionOptions := types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	}

	for _, blockId := range randomSortedProposedBlockIds {
		proposedBlock, err := razorUtils.GetProposedBlock(client, epoch, uint32(blockId))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("Proposed block ", blockId, proposedBlock)

		//blockIndex is index of blockId in sortedProposedBlock
		blockIndex := utils.IndexOf(sortedProposedBlockIds, uint32(blockId))

		// Biggest staker dispute
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
			disputeBiggestStakeProposedTxn, err := blockManagerUtils.DisputeBiggestStakeProposed(client, txnOpts, epoch, uint8(blockIndex), biggestStakerId)
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

		// Ids Dispute
		log.Debug("Locally revealed collection ids: ", revealedCollectionIds)
		log.Debug("Revealed collection ids in the block ", proposedBlock.Ids)

		idDisputeTxn, err := cmdUtils.CheckDisputeForIds(client, transactionOptions, epoch, uint8(blockIndex), proposedBlock.Ids, revealedCollectionIds)
		if err != nil {
			log.Error("Error in disputing: ", err)
		}
		if idDisputeTxn != nil {
			log.Debugf("Txn Hash: %s", transactionUtils.Hash(idDisputeTxn).String())
			status := razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(idDisputeTxn).String())
			if status == 1 {
				continue
			}
		}

		// Median Value dispute
		isEqual, mismatchIndex := utils.IsEqualUint32(proposedBlock.Medians, medians)
		if !isEqual {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("Block Values: ", proposedBlock.Medians)
			log.Debug("Local Calculations: ", medians)
			if proposedBlock.Valid {
				// median locally calculated: [100, 200, 300, 500]   median proposed: [100, 230, 300, 500]
				// ids [1, 2, 3, 4]
				// Sorted revealed values would be the vote values for the wrong median, here 230
				collectionIdOfWrongMedian := proposedBlock.Ids[mismatchIndex]
				sortedValues := revealedDataMaps.SortedRevealedValues[collectionIdOfWrongMedian]
				leafId, err := utils.UtilsInterface.GetLeafIdOfACollection(client, collectionIdOfWrongMedian)
				if err != nil {
					log.Error("Error in leaf id: ", err)
					continue
				}
				disputeErr := cmdUtils.Dispute(client, config, account, epoch, uint8(blockIndex), proposedBlock, leafId, sortedValues)
				if disputeErr != nil {
					log.Error("Error in disputing...", disputeErr)
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
	giveSortedLeafIds = []int{}
	return nil
}

func (*UtilsStruct) GetLocalMediansData(client *ethclient.Client, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) ([]uint32, []uint16, *types.RevealedDataMaps, error) {
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
	if _mediansData == nil || _revealedCollectionIds == nil || _revealedDataMaps == nil {
		medians, revealedCollectionIds, revealedDataMaps, err := cmdUtils.MakeBlock(client, blockNumber, epoch, types.Rogue{IsRogue: false})
		if err != nil {
			log.Error("Error in calculating block medians")
			return nil, nil, nil, err
		}
		_mediansData = razorUtils.ConvertUint32ArrayToBigIntArray(medians)
		_revealedCollectionIds = revealedCollectionIds
		_revealedDataMaps = revealedDataMaps
	}

	mediansInUint32 := razorUtils.ConvertBigIntArrayToUint32Array(_mediansData)
	log.Debug("Locally calculated data:")
	log.Debugf("Medians: %d", mediansInUint32)
	return mediansInUint32, _revealedCollectionIds, _revealedDataMaps, nil
}

func (*UtilsStruct) CheckDisputeForIds(client *ethclient.Client, transactionOpts types.TransactionOptions, epoch uint32, blockIndex uint8, idsInProposedBlock []uint16, revealedCollectionIds []uint16) (*types2.Transaction, error) {
	// Check if the error is in sorted ids
	isSorted, index0, index1 := utils.IsSorted(idsInProposedBlock)
	if !isSorted {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeOnOrderOfIds"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, index0, index1}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		log.Debug("Disputing sorted order of ids!")
		log.Debugf("Epoch: %d, blockIndex: %d, index0: %d, index1: %d", epoch, blockIndex, index0, index1)
		return blockManagerUtils.DisputeOnOrderOfIds(client, txnOpts, epoch, blockIndex, big.NewInt(int64(index0)), big.NewInt(int64(index1)))
	}

	// Check if the error is collectionIdShouldBePresent
	isMissing, _, missingCollectionId := utils.IsMissing(revealedCollectionIds, idsInProposedBlock)
	if isMissing {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBePresent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, missingCollectionId}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		log.Debug("Disputing collection id should be present!")
		log.Debugf("Epoch: %d, blockIndex: %d, missingCollectionId: %d", epoch, blockIndex, missingCollectionId)
		return blockManagerUtils.DisputeCollectionIdShouldBePresent(client, txnOpts, epoch, blockIndex, missingCollectionId)
	}

	// Check if the error is collectionIdShouldBeAbsent
	isPresent, positionOfPresentValue, presentCollectionId := utils.IsMissing(idsInProposedBlock, revealedCollectionIds)
	if isPresent {
		transactionOpts.ABI = bindings.BlockManagerABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBeAbsent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, presentCollectionId, big.NewInt(int64(positionOfPresentValue))}
		txnOpts := razorUtils.GetTxnOpts(transactionOpts)
		log.Debug("Disputing collection id should be absent!")
		log.Debugf("Epoch: %d, blockIndex: %d, presentCollectionId: %d, positionOfPresentValue: %d", epoch, blockIndex, presentCollectionId, positionOfPresentValue)
		return blockManagerUtils.DisputeCollectionIdShouldBeAbsent(client, txnOpts, epoch, blockIndex, presentCollectionId, big.NewInt(int64(positionOfPresentValue)))
	}

	return nil, nil
}

func (*UtilsStruct) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockIndex uint8, proposedBlock bindings.StructsBlock, leafId uint16, sortedValues []uint32) error {
	blockManager := razorUtils.GetBlockManager(client)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if !utils.Contains(giveSortedLeafIds, leafId) {
		cmdUtils.GiveSorted(client, blockManager, txnOpts, epoch, leafId, sortedValues)
	}

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	positionOfCollectionInBlock := cmdUtils.GetCollectionIdPositionInBlock(client, leafId, proposedBlock)
	finalizeTxn, err := blockManagerUtils.FinalizeDispute(client, finalizeDisputeTxnOpts, epoch, blockIndex, positionOfCollectionInBlock)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(finalizeTxn))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(finalizeTxn).String())
	return nil
}

func GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []uint32) {
	if len(sortedValues) == 0 {
		return
	}
	txn, err := blockManagerUtils.GiveSorted(blockManager, txnOpts, epoch, leafId, sortedValues)
	if err != nil {
		if err.Error() == errors.New("gas limit reached").Error() {
			log.Error("Error in calling GiveSorted: ", err)
			mid := len(sortedValues) / 2
			GiveSorted(client, blockManager, txnOpts, epoch, leafId, sortedValues[:mid])
			GiveSorted(client, blockManager, txnOpts, epoch, leafId, sortedValues[mid:])
		} else {
			return
		}
	}
	log.Info("Calling GiveSorted...")
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	giveSortedLeafIds = append(giveSortedLeafIds, int(leafId))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(txn).String())
}

func (*UtilsStruct) GetCollectionIdPositionInBlock(client *ethclient.Client, leafId uint16, proposedBlock bindings.StructsBlock) *big.Int {
	ids := proposedBlock.Ids
	idToBeDisputed, err := utils.UtilsInterface.GetCollectionIdFromLeafId(client, leafId)
	if err != nil {
		log.Error("Error in fetching collection id from leaf id")
		return nil
	}
	for i := 0; i < len(ids); i++ {
		if ids[i] == idToBeDisputed {
			return big.NewInt(int64(i))
		}
	}
	return nil
}
