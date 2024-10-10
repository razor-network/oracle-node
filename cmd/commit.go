//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"encoding/hex"
	"errors"
	Types "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

/*
GetSalt calculates the salt on the basis of previous epoch and the medians of the previous epoch.
If the previous epoch doesn't contain any medians, then the value is fetched from the smart contract.
*/
func (*UtilsStruct) GetSalt(ctx context.Context, client *ethclient.Client, epoch uint32) ([32]byte, error) {
	previousEpoch := epoch - 1
	log.Debug("GetSalt: Previous epoch: ", previousEpoch)
	numProposedBlock, err := razorUtils.GetNumberOfProposedBlocks(ctx, client, previousEpoch)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Number of proposed blocks: ", numProposedBlock)
	blockIndexToBeConfirmed, err := razorUtils.GetBlockIndexToBeConfirmed(ctx, client)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Block Index to be confirmed: ", blockIndexToBeConfirmed)
	if numProposedBlock == 0 || (numProposedBlock > 0 && blockIndexToBeConfirmed < 0) {
		return utils.VoteManagerInterface.GetSaltFromBlockchain(client)
	}
	blockId, err := razorUtils.GetSortedProposedBlockId(ctx, client, previousEpoch, big.NewInt(int64(blockIndexToBeConfirmed)))
	if err != nil {
		return [32]byte{}, errors.New("Error in getting blockId: " + err.Error())
	}
	log.Debug("GetSalt: Block Id: ", blockId)
	previousBlock, err := razorUtils.GetProposedBlock(ctx, client, previousEpoch, blockId)
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
func (*UtilsStruct) HandleCommitState(ctx context.Context, client *ethclient.Client, epoch uint32, seed []byte, commitParams *types.CommitParams, rogueData types.Rogue) (types.CommitData, error) {
	numActiveCollections, err := razorUtils.GetNumActiveCollections(ctx, client)
	if err != nil {
		return types.CommitData{}, err
	}
	log.Debug("HandleCommitState: Number of active collections: ", numActiveCollections)
	log.Debugf("HandleCommitState: Calling GetAssignedCollections() with arguments number of active collections = %d", numActiveCollections)
	assignedCollections, seqAllottedCollections, err := razorUtils.GetAssignedCollections(ctx, client, numActiveCollections, seed)
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
				collectionId, err := razorUtils.GetCollectionIdFromIndex(ctx, client, uint16(i))
				if err != nil {
					log.Error("Error in getting collection ID: ", err)
					errChan <- err
					return
				}
				collectionData, err := razorUtils.GetAggregatedDataOfCollection(ctx, client, collectionId, epoch, commitParams)
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
func (*UtilsStruct) Commit(ctx context.Context, client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, seed []byte, values []*big.Int) (common.Hash, error) {
	if state, err := razorUtils.GetBufferedState(latestHeader, stateBuffer, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	commitmentToSend, err := CalculateCommitment(seed, values)
	if err != nil {
		log.Error("Error in getting commitment: ", err)
		return core.NilHash, err
	}

	txnOpts := razorUtils.GetTxnOpts(ctx, types.TransactionOptions{
		Client:          client,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
		Account:         account,
	})

	log.Info("Commitment sent...")
	log.Debugf("Executing Commit transaction with epoch = %d, commitmentToSend = %v", epoch, commitmentToSend)
	txn, err := voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func CalculateSeed(ctx context.Context, client *ethclient.Client, account types.Account, keystorePath string, epoch uint32) ([]byte, error) {
	log.Debugf("CalculateSeed: Calling CalculateSecret() with arguments epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
	_, secret, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return nil, err
	}
	log.Debugf("CalculateSeed: Getting Salt for current epoch %d...", epoch)
	salt, err := cmdUtils.GetSalt(ctx, client, epoch)
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

func VerifyCommitment(ctx context.Context, client *ethclient.Client, account types.Account, keystorePath string, epoch uint32, values []*big.Int) (bool, error) {
	commitmentStruct, err := razorUtils.GetCommitment(ctx, client, account.Address)
	if err != nil {
		log.Error("Error in getting commitments: ", err)
		return false, err
	}
	log.Debugf("VerifyCommitment: CommitmentStruct: %+v", commitmentStruct)

	seed, err := CalculateSeed(ctx, client, account, keystorePath, epoch)
	if err != nil {
		log.Error("Error in calculating seed: ", err)
		return false, err
	}

	calculatedCommitment, err := CalculateCommitment(seed, values)
	if err != nil {
		log.Error("Error in calculating commitment for given committed values: ", err)
		return false, err
	}
	log.Debug("VerifyCommitment: Calculated commitment: ", calculatedCommitment)

	if calculatedCommitment == commitmentStruct.CommitmentHash {
		log.Debug("VerifyCommitment: Calculated commitment for given values is EQUAL to commitment of the epoch")
		return true, nil
	}
	log.Debug("VerifyCommitment: Calculated commitment for given values DOES NOT MATCH with commitment in the epoch")
	return false, nil
}
