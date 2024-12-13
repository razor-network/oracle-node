//Package cmd provides all functions related to command line
package cmd

import (
	"encoding/hex"
	"errors"
	Types "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

/*
GetSalt calculates the salt on the basis of previous epoch and the medians of the previous epoch.
If the previous epoch doesn't contain any medians, then the value is fetched from the smart contract.
*/
func (*UtilsStruct) GetSalt(rpcParameters rpc.RPCParameters, epoch uint32) ([32]byte, error) {
	previousEpoch := epoch - 1
	log.Debug("GetSalt: Previous epoch: ", previousEpoch)
	numProposedBlock, err := razorUtils.GetNumberOfProposedBlocks(rpcParameters, previousEpoch)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Number of proposed blocks: ", numProposedBlock)
	blockIndexToBeConfirmed, err := razorUtils.GetBlockIndexToBeConfirmed(rpcParameters)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Block Index to be confirmed: ", blockIndexToBeConfirmed)
	if numProposedBlock == 0 || (numProposedBlock > 0 && blockIndexToBeConfirmed < 0) {
		return razorUtils.GetSaltFromBlockchain(rpcParameters)
	}
	blockId, err := razorUtils.GetSortedProposedBlockId(rpcParameters, previousEpoch, big.NewInt(int64(blockIndexToBeConfirmed)))
	if err != nil {
		return [32]byte{}, errors.New("Error in getting blockId: " + err.Error())
	}
	log.Debug("GetSalt: Block Id: ", blockId)
	previousBlock, err := razorUtils.GetProposedBlock(rpcParameters, previousEpoch, blockId)
	if err != nil {
		return [32]byte{}, errors.New("Error in getting previous block: " + err.Error())
	}
	log.Debug("GetSalt: PreviousBlock: ", previousBlock)
	log.Debugf("GetSalt: Calling CalculateSalt() with arguments previous epoch = %d, previous block medians = %s", previousEpoch, previousBlock.Medians)
	return razorUtils.CalculateSalt(previousEpoch, previousBlock.Medians), nil
}

/*
HandleCommitState fetches the collections assigned to the staker and creates the leaves required for the merkle tree generation.
Values for only the collections assigned to the staker is fetched for others, 0 is added to the leaves of tree.
*/
func (*UtilsStruct) HandleCommitState(rpcParameters rpc.RPCParameters, epoch uint32, seed []byte, commitParams *types.CommitParams, rogueData types.Rogue) (types.CommitData, error) {
	numActiveCollections, err := razorUtils.GetNumActiveCollections(rpcParameters)
	if err != nil {
		return types.CommitData{}, err
	}
	log.Debug("HandleCommitState: Number of active collections: ", numActiveCollections)
	log.Debugf("HandleCommitState: Calling GetAssignedCollections() with arguments number of active collections = %d", numActiveCollections)
	assignedCollections, seqAllottedCollections, err := razorUtils.GetAssignedCollections(rpcParameters, numActiveCollections, seed)
	if err != nil {
		return types.CommitData{}, err
	}

	leavesOfTree := make([]*big.Int, numActiveCollections)
	results := make(chan types.CollectionResult, numActiveCollections)
	errChan := make(chan error, numActiveCollections)

	defer close(results)
	defer close(errChan)

	var wg sync.WaitGroup

	// Clean up any expired API results cache data before performing the commit
	commitParams.LocalCache.Cleanup()

	log.Debug("Iterating over all the collections...")
	for i := 0; i < int(numActiveCollections); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var leaf *big.Int

			log.Debugf("HandleCommitState: Is the collection at iterating index %v assigned: %v ", i, assignedCollections[i])
			if assignedCollections[i] {
				collectionId, err := razorUtils.GetCollectionIdFromIndex(rpcParameters, uint16(i))
				if err != nil {
					log.Error("Error in getting collection ID: ", err)
					errChan <- err
					return
				}
				collectionData, err := razorUtils.GetAggregatedDataOfCollection(rpcParameters, collectionId, epoch, commitParams)
				if err != nil {
					log.Error("Error in getting aggregated data of collection: ", err)
					errChan <- err
					return
				}
				if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "commit") {
					log.Warn("YOU ARE COMMITTING VALUES IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
					collectionData = razorUtils.GetRogueRandomValue(100000)
					log.Debug("HandleCommitState: Collection data in rogue mode: ", collectionData)
				}
				log.Debugf("HandleCommitState: Data of collection %d: %s", collectionId, collectionData)
				leaf = collectionData
			} else {
				leaf = big.NewInt(0)
			}
			log.Debugf("Sending index: %v,  leaf data: %v to results channel", i, leaf)
			results <- types.CollectionResult{Index: i, Leaf: leaf}
		}(i)
	}

	wg.Wait()

	for i := 0; i < int(numActiveCollections); i++ {
		select {
		case result := <-results:
			log.Infof("Received from results: Index: %d, Leaf: %v", result.Index, result.Leaf)
			leavesOfTree[result.Index] = result.Leaf
		case err := <-errChan:
			if err != nil {
				// Returning the first error from the error channel
				log.Error("Error in getting collection data: ", err)
				return types.CommitData{}, err
			}
		}
	}

	log.Debug("HandleCommitState: Assigned Collections: ", assignedCollections)
	log.Debug("HandleCommitState: SeqAllottedCollections: ", seqAllottedCollections)
	log.Debug("HandleCommitState: Leaves: ", leavesOfTree)

	return types.CommitData{
		AssignedCollections:    assignedCollections,
		SeqAllottedCollections: seqAllottedCollections,
		Leaves:                 leavesOfTree,
	}, nil
}

/*
Commit finally commits the data to the smart contract. It calculates the commitment to send using the merkle tree root and the seed.
*/
func (*UtilsStruct) Commit(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, commitmentToSend [32]byte) (common.Hash, error) {
	if state, err := razorUtils.GetBufferedState(latestHeader, stateBuffer, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
		Account:         account,
	})
	if err != nil {
		return core.NilHash, err
	}

	log.Info("Commitment sent...")
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	log.Debugf("Executing Commit transaction with epoch = %d, commitmentToSend = %v", epoch, commitmentToSend)
	txn, err := voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func CalculateSeed(rpcParameters rpc.RPCParameters, account types.Account, keystorePath string, epoch uint32) ([]byte, error) {
	log.Debugf("CalculateSeed: Calling CalculateSecret() with arguments epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
	_, secret, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return nil, err
	}
	log.Debugf("CalculateSeed: Getting Salt for current epoch %d...", epoch)
	salt, err := cmdUtils.GetSalt(rpcParameters, epoch)
	if err != nil {
		log.Error("Error in getting salt: ", err)
		return nil, err
	}
	seed := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(salt[:]), "0x" + hex.EncodeToString(secret)})
	return seed, nil
}

func CalculateCommitment(seed []byte, values []*big.Int) ([32]byte, error) {
	log.Debug("CalculateCommitment: Calling CreateMerkle() with argument Leaves = ", values)
	merkleTree, err := merkleUtils.CreateMerkle(values)
	if err != nil {
		return [32]byte{}, errors.New("Error in getting merkle tree: " + err.Error())
	}
	log.Debug("CalculateCommitment: Merkle Tree: ", merkleTree)
	log.Debug("CalculateCommitment: Calling GetMerkleRoot() for the merkle tree...")
	merkleRoot, err := merkleUtils.GetMerkleRoot(merkleTree)
	if err != nil {
		return [32]byte{}, errors.New("Error in getting root: " + err.Error())
	}
	commitment := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(merkleRoot[:]), "0x" + hex.EncodeToString(seed)})
	log.Debug("CalculateCommitment: Commitment: ", hex.EncodeToString(commitment))
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)
	return commitmentToSend, nil
}

func VerifyCommitment(rpcParameters rpc.RPCParameters, account types.Account, commitmentFetched [32]byte) (bool, error) {
	commitmentStruct, err := razorUtils.GetCommitment(rpcParameters, account.Address)
	if err != nil {
		log.Error("Error in getting commitments: ", err)
		return false, err
	}
	log.Debugf("VerifyCommitment: CommitmentStruct: %+v", commitmentStruct)

	if commitmentFetched == commitmentStruct.CommitmentHash {
		log.Debug("VerifyCommitment: Commitment fetched from memory/file system for given values is EQUAL to commitment of the epoch")
		return true, nil
	}
	log.Debug("VerifyCommitment: Commitment fetched from memory/file system for given values DOES NOT MATCH with commitment in the epoch")
	return false, nil
}

func GetCommittedDataForEpoch(rpcParameters rpc.RPCParameters, account types.Account, epoch uint32, rogueData types.Rogue) (types.CommitFileData, error) {
	// Attempt to fetch global commit data from memory if epoch matches
	if globalCommitDataStruct.Epoch == epoch {
		log.Debugf("Epoch in global commit data is equal to current epoch %v. Fetching commit data from memory!", epoch)
	} else {
		// Fetch from file if memory data is outdated
		log.Debugf("GetCommittedDataForEpoch: Global commit data epoch %v doesn't match current epoch %v. Fetching from file!", globalCommitDataStruct.Epoch, epoch)
		log.Info("Getting the commit data from file...")
		fileName, err := pathUtils.GetCommitDataFileName(account.Address)
		if err != nil {
			return types.CommitFileData{}, err
		}

		log.Debug("GetCommittedDataForEpoch: Commit data file path: ", fileName)
		commitDataFromFile, err := fileUtils.ReadFromCommitJsonFile(fileName)
		if err != nil {
			return types.CommitFileData{}, err
		}

		log.Debug("GetCommittedDataForEpoch: Committed data from file: ", commitDataFromFile)
		if commitDataFromFile.Epoch != epoch {
			log.Errorf("File %s doesn't contain latest committed data", fileName)
			return types.CommitFileData{}, errors.New("commit data file doesn't contain latest committed data")
		}

		// Update global commit data struct since the file data is valid
		updateGlobalCommitDataStruct(types.CommitData{
			Leaves:                 commitDataFromFile.Leaves,
			SeqAllottedCollections: commitDataFromFile.SeqAllottedCollections,
			AssignedCollections:    commitDataFromFile.AssignedCollections,
		}, commitDataFromFile.Commitment, epoch)
	}

	// Verify the final selected commit data
	log.Debugf("Verifying commit data for epoch %v...", epoch)
	isValid, err := VerifyCommitment(rpcParameters, account, globalCommitDataStruct.Commitment)
	if err != nil {
		return types.CommitFileData{}, err
	}
	if !isValid {
		return types.CommitFileData{}, errors.New("commitment verification failed for selected commit data")
	}

	// If rogue mode is enabled, alter the commitment data
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "reveal") {
		log.Warn("YOU ARE REVEALING VALUES IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		globalCommitDataStruct.Leaves = generateRogueCommittedData(len(globalCommitDataStruct.Leaves))
		log.Debugf("Global Commit data struct in rogue mode: %+v", globalCommitDataStruct)
	}

	return globalCommitDataStruct, nil
}

func generateRogueCommittedData(length int) []*big.Int {
	var rogueCommittedData []*big.Int
	for i := 0; i < length; i++ {
		rogueCommittedData = append(rogueCommittedData, razorUtils.GetRogueRandomValue(10000000))
	}
	return rogueCommittedData
}
