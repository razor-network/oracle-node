//Package cmd provides all functions related to command line
package cmd

import (
	"encoding/hex"
	"errors"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

/*
GetSalt calculates the salt on the basis of previous epoch and the medians of the previous epoch.
If the previous epoch doesn't contain any medians, then the value is fetched from the smart contract.
*/
func (*UtilsStruct) GetSalt(client *ethclient.Client, epoch uint32) ([32]byte, error) {
	previousEpoch := epoch - 1
	log.Debug("GetSalt: Previous epoch: ", previousEpoch)
	numProposedBlock, err := razorUtils.GetNumberOfProposedBlocks(client, previousEpoch)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Number of proposed blocks: ", numProposedBlock)
	blockIndexToBeConfirmed, err := razorUtils.GetBlockIndexToBeConfirmed(client)
	if err != nil {
		return [32]byte{}, err
	}
	log.Debug("GetSalt: Block Index to be confirmed: ", blockIndexToBeConfirmed)
	if numProposedBlock == 0 || (numProposedBlock > 0 && blockIndexToBeConfirmed < 0) {
		return utils.VoteManagerInterface.GetSaltFromBlockchain(client)
	}
	blockId, err := razorUtils.GetSortedProposedBlockId(client, previousEpoch, big.NewInt(int64(blockIndexToBeConfirmed)))
	if err != nil {
		return [32]byte{}, errors.New("Error in getting blockId: " + err.Error())
	}
	log.Debug("GetSalt: Block Id: ", blockId)
	previousBlock, err := razorUtils.GetProposedBlock(client, previousEpoch, blockId)
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
func (*UtilsStruct) HandleCommitState(client *ethclient.Client, epoch uint32, seed []byte, rogueData types.Rogue) (types.CommitData, error) {
	numActiveCollections, err := razorUtils.GetNumActiveCollections(client)
	if err != nil {
		return types.CommitData{}, err
	}
	log.Debug("HandleCommitState: Number of active collections: ", numActiveCollections)
	log.Debugf("HandleCommitState: Calling GetAssignedCollections() with arguments number of active collections = %d", numActiveCollections)
	assignedCollections, seqAllottedCollections, err := razorUtils.GetAssignedCollections(client, numActiveCollections, seed)
	if err != nil {
		return types.CommitData{}, err
	}

	var leavesOfTree []*big.Int

	log.Debug("Creating a local cache which will store API result and expire at the end of commit state")
	localCache := cache.NewLocalCache(time.Second * time.Duration(core.StateLength))

	log.Debug("Iterating over all the collections...")
	for i := 0; i < int(numActiveCollections); i++ {
		log.Debug("HandleCommitState: Iterating index: ", i)
		log.Debug("HandleCommitState: Is the collection assigned: ", assignedCollections[i])
		if assignedCollections[i] {
			collectionId, err := razorUtils.GetCollectionIdFromIndex(client, uint16(i))
			if err != nil {
				return types.CommitData{}, err
			}
			collectionData, err := razorUtils.GetAggregatedDataOfCollection(client, collectionId, epoch, localCache)
			if err != nil {
				return types.CommitData{}, err
			}
			if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "commit") {
				log.Warn("YOU ARE COMMITTING VALUES IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
				collectionData = razorUtils.GetRogueRandomValue(100000)
				log.Debug("HandleCommitState: Collection data in rogue mode: ", collectionData)
			}
			log.Debugf("HandleCommitState: Data of collection %d: %s", collectionId, collectionData)
			leavesOfTree = append(leavesOfTree, collectionData)
		} else {
			leavesOfTree = append(leavesOfTree, big.NewInt(0))
		}
	}
	log.Debug("HandleCommitState: Assigned Collections: ", assignedCollections)
	log.Debug("HandleCommitState: SeqAllottedCollections: ", seqAllottedCollections)
	log.Debug("HandleCommitState: Leaves: ", leavesOfTree)

	localCache.StopCleanup()

	return types.CommitData{
		AssignedCollections:    assignedCollections,
		SeqAllottedCollections: seqAllottedCollections,
		Leaves:                 leavesOfTree,
	}, nil
}

/*
Commit finally commits the data to the smart contract. It calculates the commitment to send using the merkle tree root and the seed.
*/
func (*UtilsStruct) Commit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, seed []byte, values []*big.Int) (common.Hash, error) {
	if state, err := razorUtils.GetBufferedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	commitmentToSend, err := CalculateCommitment(seed, values)
	if err != nil {
		log.Error("Error in getting commitment: ", err)
		return core.NilHash, err
	}

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
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

func CalculateSeed(client *ethclient.Client, account types.Account, keystorePath string, epoch uint32) ([]byte, error) {
	log.Debugf("CalculateSeed: Calling CalculateSecret() with arguments epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
	_, secret, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return nil, err
	}
	log.Debugf("CalculateSeed: Getting Salt for current epoch %d...", epoch)
	salt, err := cmdUtils.GetSalt(client, epoch)
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

func VerifyCommitment(client *ethclient.Client, account types.Account, keystorePath string, epoch uint32, values []*big.Int) (bool, error) {
	commitmentStruct, err := razorUtils.GetCommitment(client, account.Address)
	if err != nil {
		log.Error("Error in getting commitments: ", err)
		return false, err
	}
	log.Debugf("VerifyCommitment: CommitmentStruct: %+v", commitmentStruct)

	seed, err := CalculateSeed(client, account, keystorePath, epoch)
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
