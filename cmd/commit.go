//Package cmd provides all functions related to command line
package cmd

import (
	"encoding/hex"
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

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
	numProposedBlocks, err := utils.UtilsInterface.GetNumberOfProposedBlocks(client, previousEpoch)
	if err != nil {
		return [32]byte{}, err
	}
	blockIndexToBeConfirmed, err := utils.UtilsInterface.GetBlockIndexToBeConfirmed(client)
	if err != nil {
		return [32]byte{}, err
	}
	if numProposedBlocks == 0 || (numProposedBlocks > 0 && blockIndexToBeConfirmed < 0) {
		return utils.VoteManagerInterface.GetSaltFromBlockchain(client)
	}
	blockId, err := utils.UtilsInterface.GetSortedProposedBlockId(client, previousEpoch, big.NewInt(int64(blockIndexToBeConfirmed)))
	if err != nil {
		return [32]byte{}, errors.New("Error in getting blockId: " + err.Error())
	}
	previousBlock, err := utils.UtilsInterface.GetProposedBlock(client, previousEpoch, blockId)
	if err != nil {
		return [32]byte{}, errors.New("Error in getting previous block: " + err.Error())
	}
	return utils.UtilsInterface.CalculateSalt(previousEpoch, previousBlock.Medians), nil
}

/*
HandleCommitState fetches the collections assigned to the staker and creates the leaves required for the merkle tree generation.
Values for only the collections assigned to the staker is fetched for others, 0 is added to the leaves of tree.
*/
func (*UtilsStruct) HandleCommitState(client *ethclient.Client, epoch uint32, seed []byte, rogueData types.Rogue) (types.CommitData, error) {
	numActiveCollections, err := utils.UtilsInterface.GetNumActiveCollections(client)
	if err != nil {
		return types.CommitData{}, err
	}

	assignedCollections, seqAllottedCollections, err := utils.UtilsInterface.GetAssignedCollections(client, numActiveCollections, seed)
	if err != nil {
		return types.CommitData{}, err
	}

	var leavesOfTree []*big.Int
	for i := 0; i < int(numActiveCollections); i++ {
		if assignedCollections[i] {
			collectionId, err := utils.UtilsInterface.GetCollectionIdFromIndex(client, uint16(i))
			if err != nil {
				return types.CommitData{}, err
			}
			collectionData, err := utils.UtilsInterface.GetAggregatedDataOfCollection(client, collectionId, epoch)
			if err != nil {
				return types.CommitData{}, err
			}
			if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "commit") {
				collectionData = razorUtils.GetRogueRandomValue(100000)
			}
			log.Debugf("Data of collection %d:%s", collectionId, collectionData)
			leavesOfTree = append(leavesOfTree, collectionData)
		} else {
			leavesOfTree = append(leavesOfTree, big.NewInt(0))
		}
	}
	log.Debug("Assigned Collections: ", assignedCollections)
	log.Debug("SeqAllottedCollections: ", seqAllottedCollections)
	log.Debug("Leaves: ", leavesOfTree)
	return types.CommitData{
		AssignedCollections:    assignedCollections,
		SeqAllottedCollections: seqAllottedCollections,
		Leaves:                 leavesOfTree,
	}, nil
}

/*
Commit finally commits the data to the smart contract. It calculates the commitment to send using the merkle tree root and the seed.
*/
func (*UtilsStruct) Commit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, seed []byte, root [32]byte) (common.Hash, error) {
	if state, err := razorUtils.GetBufferedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	commitment := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(root[:]), "0x" + hex.EncodeToString(seed)})
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         big.NewInt(config.ChainId),
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
	})

	log.Debugf("Committing: epoch: %d, commitment: %s, seed: %s, account: %s", epoch, "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(seed), account.Address)

	log.Info("Commitment sent...")
	txn, err := voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	txnHash := transactionUtils.Hash(txn)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}
