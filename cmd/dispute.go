//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

var (
	giveSortedLeafIds []int
)

//blockId is id of the block

//This function handles the dispute and if there is any error it returns the error
func (*UtilsStruct) HandleDispute(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue, backupNodeActionsToIgnore []string) error {

	sortedProposedBlockIds, err := razorUtils.GetSortedProposedBlockIds(rpcParameters, epoch)
	if err != nil {
		log.Error("Error in fetching sorted proposed block id: ", err)
		return err
	}
	log.Debug("HandleDispute: SortedProposedBlockIds: ", sortedProposedBlockIds)

	biggestStake, biggestStakerId, err := cmdUtils.GetBiggestStakeAndId(rpcParameters, epoch)
	if err != nil {
		return err
	}
	log.Debugf("HandleDispute: Biggest stake: %s, Biggest staker Id: %d", biggestStake, biggestStakerId)

	log.Debugf("HandleDispute: Calling GetLocalMediansData() with arguments epoch = %d, blockNumber = %d, rogueData = %+v", epoch, blockNumber, rogueData)
	locallyCalculatedData, err := cmdUtils.GetLocalMediansData(rpcParameters, account, epoch, blockNumber, rogueData)
	if err != nil {
		return err
	}
	medians := locallyCalculatedData.MediansData
	revealedCollectionIds := locallyCalculatedData.RevealedCollectionIds
	revealedDataMaps := locallyCalculatedData.RevealedDataMaps

	log.Debug("HandleDispute: Local Medians data:", medians)
	log.Debug("HandleDispute: Revealed collection ids:", revealedCollectionIds)
	log.Debug("HandleDispute: Local revealed data maps:", revealedDataMaps)

	randomSortedProposedBlockIds := utils.Shuffle(sortedProposedBlockIds) //shuffles the sortedProposedBlockIds array
	transactionOptions := types.TransactionOptions{
		ChainId: core.ChainId,
		Config:  config,
		Account: account,
	}
	log.Debug("HandleDispute: Shuffled sorted proposed blocks: ", randomSortedProposedBlockIds)

	log.Debug("Iterating over random sorted proposed blocks to check dispute...")
	for _, blockId := range randomSortedProposedBlockIds {
		proposedBlock, err := razorUtils.GetProposedBlock(rpcParameters, epoch, blockId)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("HandleDispute: Block ID: ", blockId)
		log.Debug("HandleDispute: Proposed block ", proposedBlock)

		//blockIndex is index of blockId in sortedProposedBlock
		blockIndex := utils.IndexOf(sortedProposedBlockIds, blockId)
		if blockIndex == -1 {
			log.Error("Block is not present in SortedProposedBlockIds array")
			continue
		}
		log.Debug("HandleDispute: Block Index: ", blockIndex)
		// Biggest staker dispute
		if proposedBlock.BiggestStake.Cmp(biggestStake) != 0 && proposedBlock.Valid {
			log.Debug("Biggest Stake in proposed block: ", proposedBlock.BiggestStake)
			log.Warn("PROPOSED BIGGEST STAKE DOES NOT MATCH WITH ACTUAL BIGGEST STAKE")
			log.Info("Disputing BiggestStakeProposed...")
			txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, types.TransactionOptions{
				ChainId: core.ChainId,
				Config:  config,
				Account: account,
			})
			if err != nil {
				log.Error(err)
				continue
			}

			client, err := rpcParameters.RPCManager.GetBestRPCClient()
			if err != nil {
				log.Error(err)
				continue
			}

			log.Debugf("Executing DisputeBiggestStakeProposed transaction with arguments epoch = %d, blockIndex = %d, biggest staker Id = %d", epoch, blockIndex, biggestStakerId)
			disputeBiggestStakeProposedTxn, err := blockManagerUtils.DisputeBiggestStakeProposed(client, txnOpts, epoch, uint8(blockIndex), biggestStakerId)
			if err != nil {
				log.Error(err)
				continue
			}
			disputeBiggestStakeProposedTxnHash := transactionUtils.Hash(disputeBiggestStakeProposedTxn)
			log.Info("Txn Hash: ", disputeBiggestStakeProposedTxnHash.Hex())
			WaitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(rpcParameters, disputeBiggestStakeProposedTxnHash.Hex())

			//If dispute happens, then storing the bountyId into disputeData file
			if WaitForBlockCompletionErr == nil {
				log.Debug("Storing bounty Id in dispute data file....")
				err = cmdUtils.StoreBountyId(rpcParameters, account)
				if err != nil {
					log.Error(err)
				}
				continue
			}
		}

		// Ids Dispute
		log.Debug("HandleDispute: Locally revealed collection ids: ", revealedCollectionIds)
		log.Debug("HandleDispute: Revealed collection ids in the block ", proposedBlock.Ids)

		log.Debugf("HandleDispute: Calling CheckDisputeForIds() with arguments epoch = %d, blockIndex = %d, proposed revealed Ids = %v, locally calculated revealed Ids = %v", epoch, blockIndex, proposedBlock.Ids, revealedCollectionIds)
		idDisputeTxn, err := cmdUtils.CheckDisputeForIds(rpcParameters, transactionOptions, epoch, uint8(blockIndex), proposedBlock.Ids, revealedCollectionIds)
		if err != nil {
			log.Error("Error in disputing: ", err)
			continue
		}
		if idDisputeTxn != nil {
			idDisputeTxnHash := transactionUtils.Hash(idDisputeTxn)
			log.Debugf("Txn Hash: %s", idDisputeTxnHash.Hex())
			WaitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(rpcParameters, idDisputeTxnHash.Hex())

			//If dispute happens, then storing the bountyId into disputeData file
			if WaitForBlockCompletionErr == nil {
				log.Debug("Storing bounty Id in dispute data file...")
				err = cmdUtils.StoreBountyId(rpcParameters, account)
				if err != nil {
					log.Error(err)
				}
				continue
			}
		}

		// Median Value dispute
		isEqual, mismatchIndex := utils.IsBigIntArrayEqual(proposedBlock.Medians, medians)
		if !isEqual && !utils.Contains(backupNodeActionsToIgnore, "disputeMedians") {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Debug("HandleDispute: Block Values: ", proposedBlock.Medians)
			log.Debug("HandleDispute: Local Calculations: ", medians)
			if proposedBlock.Valid && len(proposedBlock.Ids) != 0 && len(proposedBlock.Medians) != 0 {
				// median locally calculated: [100, 200, 300, 500]   median proposed: [100, 230, 300, 500]
				// ids [1, 2, 3, 4]
				// Sorted revealed values would be the vote values for the wrong median, here 230
				log.Debug("HandleDispute: Mismatch index while iterating: ", mismatchIndex)
				if mismatchIndex < 0 || mismatchIndex >= len(proposedBlock.Ids) {
					log.Error("Mismatch index is out of bounds, cannot continue dispute for medians")
					continue
				}

				collectionIdOfWrongMedian := proposedBlock.Ids[mismatchIndex]
				log.Debug("HandleDispute: Collection Id of wrong median: ", collectionIdOfWrongMedian)

				//collectionId starts from 1 and in SortedRevealedValues, the keys start from 0 which are collectionId-1 mapping to respective revealed data for that collectionId.
				//e.g. collectionId = [1,2,3,4] & Sorted Reveal Votes: map[0:[100] 1:[200 202] 2:[300]]
				//Here 0th key in map represents collectionId 1.

				sortedValues := revealedDataMaps.SortedRevealedValues[collectionIdOfWrongMedian-1]
				log.Debug("HandleDispute: Sorted values: ", sortedValues)
				leafId, err := razorUtils.GetLeafIdOfACollection(rpcParameters, collectionIdOfWrongMedian)
				if err != nil {
					log.Error("Error in leaf id: ", err)
					continue
				}
				log.Debug("HandleDispute: Leaf Id: ", leafId)
				log.Debugf("Calling Dispute() with arguments epoch = %d, blockIndex = %d, proposed block = %+v, leafId = %d, sortedValues = %s", epoch, uint8(blockIndex), proposedBlock, leafId, sortedValues)
				disputeErr := cmdUtils.Dispute(rpcParameters, config, account, epoch, uint8(blockIndex), proposedBlock, leafId, sortedValues)
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

//This function returns the local median data
func (*UtilsStruct) GetLocalMediansData(rpcParameters rpc.RPCParameters, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) (types.ProposeFileData, error) {
	if rogueData.IsRogue {
		// As the staker has proposed with incorrect medians in rogue mode so those values needs to be compared with the correct calculated medians
		log.Debug("Staker proposed in rogue mode, now calculating medians correctly...")
		return calculateMedian(rpcParameters, account, epoch, blockNumber)
	}

	// Fetching the data from file only if the node is not in rogue mode and
	// the proposed data in memory is nil or epoch in propose data from memory doesn't match with current epoch
	nilProposedData := globalProposedDataStruct.MediansData == nil || globalProposedDataStruct.RevealedDataMaps == nil || globalProposedDataStruct.RevealedCollectionIds == nil
	if nilProposedData || epoch != globalProposedDataStruct.Epoch {
		log.Debug("Global propose data struct is not updated, getting the proposed data from file...")
		fileName, err := pathUtils.GetProposeDataFileName(account.Address)
		if err != nil {
			log.Error("Error in getting file name to read median data: ", err)
			return calculateMedian(rpcParameters, account, epoch, blockNumber)
		}
		log.Debug("GetLocalMediansData: Propose data file path: ", fileName)
		proposedData, err := fileUtils.ReadFromProposeJsonFile(fileName)
		if err != nil {
			log.Errorf("Error in getting propose data from file %s: %v", fileName, err)
			return calculateMedian(rpcParameters, account, epoch, blockNumber)
		}
		log.Debugf("GetLocalMediansData: Proposed data from file: %+v", proposedData)
		if proposedData.Epoch != epoch {
			log.Errorf("File %s doesn't contain latest median data", fileName)
			return calculateMedian(rpcParameters, account, epoch, blockNumber)
		}
		return proposedData, err
	}

	return globalProposedDataStruct, nil
}

func calculateMedian(rpcParameters rpc.RPCParameters, account types.Account, epoch uint32, blockNumber *big.Int) (types.ProposeFileData, error) {
	stakerId, err := razorUtils.GetStakerId(rpcParameters, account.Address)
	if err != nil {
		log.Error("Error in getting stakerId: ", err)
		return types.ProposeFileData{}, err
	}
	log.Debug("GetLocalMediansData: Staker Id: ", stakerId)

	log.Debug("Calculating the medians data again...")
	log.Debugf("GetLocalMediansData: Calling MakeBlock() with arguments blockNumber = %s, epoch = %d, rogueData = %+v", blockNumber, epoch, types.Rogue{IsRogue: false})
	medians, revealedCollectionIds, revealedDataMaps, err := cmdUtils.MakeBlock(rpcParameters, blockNumber, epoch, types.Rogue{IsRogue: false})
	if err != nil {
		log.Error("Error in calculating block medians: ", err)
		return types.ProposeFileData{}, err
	}
	calculatedProposedData := types.ProposeFileData{
		MediansData:           medians,
		RevealedCollectionIds: revealedCollectionIds,
		RevealedDataMaps:      revealedDataMaps,
		Epoch:                 epoch,
	}

	log.Debugf("Locally calculated data, Medians: %s", calculatedProposedData.MediansData)
	return calculatedProposedData, nil
}

//This function check for the dispute in different type of Id's
func (*UtilsStruct) CheckDisputeForIds(rpcParameters rpc.RPCParameters, transactionOpts types.TransactionOptions, epoch uint32, blockIndex uint8, idsInProposedBlock []uint16, revealedCollectionIds []uint16) (*Types.Transaction, error) {
	//checking for hashing whether there is any dispute or not
	hashIdsInProposedBlock := solsha3.SoliditySHA3([]string{"uint16[]"}, []interface{}{idsInProposedBlock})
	log.Debug("CheckDisputeForIds: Hash of reveal Ids in proposed block: ", hashIdsInProposedBlock)
	hashRevealedCollectionIds := solsha3.SoliditySHA3([]string{"uint16[]"}, []interface{}{revealedCollectionIds})
	log.Debug("CheckDisputeForIds: Hash of locally calculated reveal collection Ids: ", hashRevealedCollectionIds)

	isEqual, _ := utils.IsByteArrayEqual(hashIdsInProposedBlock, hashRevealedCollectionIds)

	if isEqual {
		log.Debug("Revealed collection Ids are correct, no need to dispute for ids")
		return nil, nil
	}

	// Check if the error is in sorted ids
	isSorted, index0, index1 := utils.IsSorted(idsInProposedBlock)
	if !isSorted {
		transactionOpts.ABI = bindings.BlockManagerMetaData.ABI
		transactionOpts.MethodName = "disputeOnOrderOfIds"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, index0, index1}
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, transactionOpts)
		if err != nil {
			return nil, err
		}

		client, err := rpcParameters.RPCManager.GetBestRPCClient()
		if err != nil {
			return nil, err
		}

		log.Debug("Disputing sorted order of ids!")
		log.Debugf("CheckDisputeForIds: Executing DisputeOnOrderOfIds transaction with arguments epoch: %d, blockIndex: %d, index0: %d, index1: %d", epoch, blockIndex, index0, index1)
		return blockManagerUtils.DisputeOnOrderOfIds(client, txnOpts, epoch, blockIndex, big.NewInt(int64(index0)), big.NewInt(int64(index1)))
	}

	// Check if the error is collectionIdShouldBePresent
	isValueMissing, _, missingCollectionId := utils.CheckValueMissingInArray(revealedCollectionIds, idsInProposedBlock)
	if isValueMissing {
		transactionOpts.ABI = bindings.BlockManagerMetaData.ABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBePresent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, missingCollectionId}
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, transactionOpts)
		if err != nil {
			return nil, err
		}
		gasLimit := txnOpts.GasLimit
		incrementedGasLimit, err := gasUtils.IncreaseGasLimitValue(rpcParameters, gasLimit, core.DisputeGasMultiplier)
		if err != nil {
			return nil, err
		}
		txnOpts.GasLimit = incrementedGasLimit

		client, err := rpcParameters.RPCManager.GetBestRPCClient()
		if err != nil {
			return nil, err
		}

		log.Debug("Disputing collection id should be present!")
		log.Debugf("CheckDisputeForIds: Executing DisputeCollectionIdShouldBePresent transaction with arguments epoch: %d, blockIndex: %d, missingCollectionId: %d", epoch, blockIndex, missingCollectionId)
		return blockManagerUtils.DisputeCollectionIdShouldBePresent(client, txnOpts, epoch, blockIndex, missingCollectionId)
	}

	// Check if the error is collectionIdShouldBeAbsent
	isValuePresent, positionOfPresentValue, presentCollectionId := utils.CheckValueMissingInArray(idsInProposedBlock, revealedCollectionIds)
	if isValuePresent {
		transactionOpts.ABI = bindings.BlockManagerMetaData.ABI
		transactionOpts.MethodName = "disputeCollectionIdShouldBeAbsent"
		transactionOpts.Parameters = []interface{}{epoch, blockIndex, presentCollectionId, big.NewInt(int64(positionOfPresentValue))}
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, transactionOpts)
		if err != nil {
			return nil, err
		}
		gasLimit := txnOpts.GasLimit
		incrementedGasLimit, err := gasUtils.IncreaseGasLimitValue(rpcParameters, gasLimit, core.DisputeGasMultiplier)
		if err != nil {
			return nil, err
		}
		txnOpts.GasLimit = incrementedGasLimit

		client, err := rpcParameters.RPCManager.GetBestRPCClient()
		if err != nil {
			return nil, err
		}

		log.Debug("Disputing collection id should be absent!")
		log.Debugf("CheckDisputeForIds: Executing DisputeCollectionIdShouldBeAbsent transaction with arguments epoch: %d, blockIndex: %d, presentCollectionId: %d, positionOfPresentValue: %d", epoch, blockIndex, presentCollectionId, positionOfPresentValue)
		return blockManagerUtils.DisputeCollectionIdShouldBeAbsent(client, txnOpts, epoch, blockIndex, presentCollectionId, big.NewInt(int64(positionOfPresentValue)))
	}

	return nil, nil
}

//This function finalizes the dispute and return the error if there is any
func (*UtilsStruct) Dispute(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, blockIndex uint8, proposedBlock bindings.StructsBlock, leafId uint16, sortedValues []*big.Int) error {
	txnArgs := types.TransactionOptions{
		ChainId: core.ChainId,
		Config:  config,
		Account: account,
	}

	if !utils.Contains(giveSortedLeafIds, leafId) {
		var (
			start int
			end   int
		)
		lenOfSortedValues := len(sortedValues)
		for {
			if start >= end && start != 0 && end != 0 {
				break
			}
			if end == 0 {
				end = lenOfSortedValues
			}
			log.Debugf("Dispute: Calling GiveSorted with arguments epoch = %d, leafId = %d, sortedValues = %s", epoch, leafId, sortedValues[start:end])
			err := cmdUtils.GiveSorted(rpcParameters, txnArgs, epoch, leafId, sortedValues[start:end])
			if err != nil {
				if err.Error() == errors.New("gas limit reached").Error() {
					end = end / 2
				} else {
					log.Error("Error in GiveSorted: ", err)
					txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
					if err != nil {
						return err
					}
					log.Debugf("Dispute: Calling CheckToDoResetDispute with arguments epoch = %d, sortedValues = %s", epoch, sortedValues)
					cmdUtils.CheckToDoResetDispute(rpcParameters, txnOpts, epoch, sortedValues)
					return err
				}
			} else {
				threshold := end - start
				start = end
				if end != lenOfSortedValues {
					end = end + threshold
					if end > lenOfSortedValues {
						end = lenOfSortedValues
					}
				}
			}
		}
		// Adding leafId to giveSortedLeafIds as giveSorted is done for that leafId
		giveSortedLeafIds = append(giveSortedLeafIds, int(leafId))
	}
	log.Debugf("Dispute: Calling GetCollectionIdPositionInBlock with arguments leafId = %d, proposed block = %+v", leafId, proposedBlock)
	positionOfCollectionInBlock := cmdUtils.GetCollectionIdPositionInBlock(rpcParameters, leafId, proposedBlock)
	log.Debug("Dispute: Position of collection id in block: ", positionOfCollectionInBlock)

	log.Info("Finalizing dispute...")
	finalizeDisputeTxnArgs := txnArgs
	finalizeDisputeTxnArgs.ContractAddress = core.BlockManagerAddress
	finalizeDisputeTxnArgs.MethodName = "finalizeDispute"
	finalizeDisputeTxnArgs.ABI = bindings.BlockManagerMetaData.ABI
	finalizeDisputeTxnArgs.Parameters = []interface{}{epoch, blockIndex, positionOfCollectionInBlock}
	finalizeDisputeTxnOpts, err := razorUtils.GetTxnOpts(rpcParameters, finalizeDisputeTxnArgs)
	if err != nil {
		return err
	}

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return err
	}

	log.Debugf("Executing FinalizeDispute transaction with arguments epoch = %d, blockIndex = %d, positionOfCollectionInBlock = %d", epoch, blockIndex, positionOfCollectionInBlock)
	finalizeTxn, err := blockManagerUtils.FinalizeDispute(client, finalizeDisputeTxnOpts, epoch, blockIndex, positionOfCollectionInBlock)
	if err != nil {
		log.Error("Error in FinalizeDispute: ", err)
	}

	var nilTransaction *Types.Transaction

	if finalizeTxn != nilTransaction {
		finalizeTxnHash := transactionUtils.Hash(finalizeTxn)
		log.Info("Txn Hash: ", finalizeTxnHash.Hex())
		log.Debug("Dispute: Checking mining status of FinalizeDispute transaction...")
		WaitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(rpcParameters, finalizeTxnHash.Hex())
		//If dispute happens, then storing the bountyId into disputeData file
		if WaitForBlockCompletionErr == nil {
			log.Debug("Storing bounty Id in dispute data file...")
			err = cmdUtils.StoreBountyId(rpcParameters, account)
			if err != nil {
				return err
			}
		} else {
			log.Error("Error in WaitForBlockCompletion for FinalizeDispute: ", WaitForBlockCompletionErr)
		}
	}

	//Resetting dispute irrespective of FinalizeDispute transaction status
	log.Debug("Resetting dispute ...")

	resetDisputeTxnArgs := txnArgs
	resetDisputeTxnArgs.ContractAddress = core.BlockManagerAddress
	resetDisputeTxnArgs.MethodName = "resetDispute"
	resetDisputeTxnArgs.ABI = bindings.BlockManagerMetaData.ABI
	resetDisputeTxnArgs.Parameters = []interface{}{epoch}
	resetDisputeTxnOpts, err := razorUtils.GetTxnOpts(rpcParameters, resetDisputeTxnArgs)
	if err != nil {
		return err
	}

	cmdUtils.ResetDispute(rpcParameters, resetDisputeTxnOpts, epoch)

	return nil
}

//This function sorts the Id's recursively
func GiveSorted(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions, epoch uint32, leafId uint16, sortedValues []*big.Int) error {
	if len(sortedValues) == 0 {
		return errors.New("length of sortedValues is 0")
	}
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		log.Error("Error in getting txnOpts: ", err)
		return err
	}
	disputesMapping, err := razorUtils.Disputes(rpcParameters, epoch, common.HexToAddress(txnArgs.Account.Address))
	if err != nil {
		log.Error("Error in getting disputes mapping: ", err)
		return err
	}
	log.Debugf("GiveSorted: Disputes mapping: %+v", disputesMapping)

	if disputesMapping.LastVisitedValue.Cmp(sortedValues[len(sortedValues)-1]) == 0 {
		return errors.New("giveSorted already done")
	}

	isGiveSortedInitiated := disputesMapping.LastVisitedValue.Cmp(big.NewInt(0)) > 0 && disputesMapping.AccWeight.Cmp(big.NewInt(0)) > 0
	if isGiveSortedInitiated && disputesMapping.LeafId != leafId {
		log.Error("Give sorted is in progress for another leafId")
		return errors.New("another giveSorted in progress")
	}
	log.Debug("GiveSorted: Is give sorted initiated: ", isGiveSortedInitiated)

	log.Info("Calling GiveSorted...")

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return err
	}

	log.Debugf("Executing GiveSorted transaction with arguments epoch = %d, leafId = %d , sortedValues = %s", epoch, leafId, sortedValues)
	txn, err := blockManagerUtils.GiveSorted(client, txnOpts, epoch, leafId, sortedValues)
	if err != nil {
		return err
	}

	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	giveSortedLeafIds = append(giveSortedLeafIds, int(leafId))
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txnHash.Hex())
	if err != nil {
		log.Error("Error in WaitForBlockCompletion for giveSorted: ", err)
		return err
	}
	return nil
}

//This function returns the collection Id position in block
func (*UtilsStruct) GetCollectionIdPositionInBlock(rpcParameters rpc.RPCParameters, leafId uint16, proposedBlock bindings.StructsBlock) *big.Int {
	ids := proposedBlock.Ids
	idToBeDisputed, err := razorUtils.GetCollectionIdFromLeafId(rpcParameters, leafId)
	if err != nil {
		log.Error("Error in fetching collection id from leaf id")
		return nil
	}
	log.Debug("GetCollectionIdPositionInBlock: Id to be disputed: ", idToBeDisputed)
	for i := 0; i < len(ids); i++ {
		if ids[i] == idToBeDisputed {
			log.Debug("GetCollectionIdPositionInBlock: Position of collectionId in block: ", i)
			return big.NewInt(int64(i))
		}
	}
	return nil
}

//This function saves the bountyId in disputeData file and return the error if there is any
func (*UtilsStruct) StoreBountyId(rpcParameters rpc.RPCParameters, account types.Account) error {
	disputeFilePath, err := pathUtils.GetDisputeDataFileName(account.Address)
	if err != nil {
		return err
	}
	log.Debug("StoreBountyId: Dispute data file path: ", disputeFilePath)

	var latestBountyId uint32

	latestHeader, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return err
	}
	log.Debug("StoreBountyId: Latest header: ", latestHeader)

	log.Debugf("StoreBountyId: Calling GetBountyIdFromEvents with arguments blockNumber = %d, address = %s", latestHeader.Number, account.Address)
	latestBountyId, err = cmdUtils.GetBountyIdFromEvents(rpcParameters, latestHeader.Number, account.Address)
	if err != nil {
		return err
	}

	if _, err := path.OSUtilsInterface.Stat(disputeFilePath); !errors.Is(err, os.ErrNotExist) {
		log.Debug("Fetching dispute data from dispute data file...")
		disputeData, err = fileUtils.ReadFromDisputeJsonFile(disputeFilePath)
		if err != nil {
			return err
		}
		log.Debugf("StoreBountyId: Dispute data: %+v", disputeData)
	}

	if latestBountyId != 0 {
		//prepending the latestBountyId to the queue
		disputeData.BountyIdQueue = append([]uint32{latestBountyId}, disputeData.BountyIdQueue...)
	}

	//saving the updated bountyIds to disputeData file
	log.Debug("Saving the updated bountyIds to dispute data file...")
	err = fileUtils.SaveDataToDisputeJsonFile(disputeFilePath, disputeData.BountyIdQueue)
	if err != nil {
		return err
	}
	return nil
}

//This function resets the dispute
func (*UtilsStruct) ResetDispute(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, epoch uint32) {
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return
	}
	log.Debug("Executing ResetDispute transaction with arguments epoch = ", epoch)
	txn, err := blockManagerUtils.ResetDispute(client, txnOpts, epoch)
	if err != nil {
		log.Error("error in resetting dispute", err)
		return
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txnHash.Hex())
	if err != nil {
		log.Error("Error in WaitForBlockCompletion for resetDispute: ", err)
		return
	}
	log.Info("Dispute has been reset")
}

//This function returns the bountyId from events
func (*UtilsStruct) GetBountyIdFromEvents(rpcParameters rpc.RPCParameters, blockNumber *big.Int, bountyHunter string) (uint32, error) {
	fromBlock, err := razorUtils.EstimateBlockNumberAtEpochBeginning(rpcParameters, blockNumber)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	log.Debug("GetBountyIdFromEvents: Checking for events from block: ", fromBlock)
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.StakeManagerAddress),
		},
	}
	log.Debugf("GetBountyIdFromEvents: Query to send in filter logs: %+v", query)
	logs, err := clientUtils.FilterLogsWithRetry(rpcParameters, query)
	if err != nil {
		return 0, err
	}
	contractAbi, err := utils.ABIInterface.Parse(strings.NewReader(bindings.StakeManagerMetaData.ABI))
	if err != nil {
		return 0, err
	}
	bountyId := uint32(0)
	for _, vLog := range logs {
		data, unpackErr := abiUtils.Unpack(contractAbi, "Slashed", vLog.Data)
		if unpackErr != nil {
			log.Debug(unpackErr)
			continue
		}
		topics := vLog.Topics
		// topics[1] gives bounty hunter address in data type common.Hash
		// Converting address to common.Hash to compare with bounty hunter address from topics
		addressFromEvents := topics[1]
		bountyHunterInHash := common.HexToHash(bountyHunter)
		if bountyHunterInHash == addressFromEvents {
			bountyId = data[0].(uint32)
		}
	}
	log.Debug("GetBountyIdFromEvents: Bounty Id: ", bountyId)
	return bountyId, nil
}

func (*UtilsStruct) CheckToDoResetDispute(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, epoch uint32, sortedValues []*big.Int) {
	// Fetch updated dispute mapping
	disputesMapping, err := razorUtils.Disputes(rpcParameters, epoch, txnOpts.From)
	if err != nil {
		log.Error("Error in getting disputes mapping: ", disputesMapping)
		return
	}
	log.Debugf("CheckToDoResetDispute: Disputes mapping: %+v", disputesMapping)
	log.Debug("CheckToDoResetDispute: Updated Last visited value: ", disputesMapping.LastVisitedValue)
	//Checking whether LVV is equal to maximum value in sortedValues, if not equal resetting dispute
	if disputesMapping.LastVisitedValue.Cmp(big.NewInt(0)) != 0 && disputesMapping.LastVisitedValue.Cmp(sortedValues[len(sortedValues)-1]) != 0 {
		log.Debug("CheckToDoResetDispute: Calling Reset Dispute with arguments epoch = ", epoch)
		cmdUtils.ResetDispute(rpcParameters, txnOpts, epoch)
	}
}
