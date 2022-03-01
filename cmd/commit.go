package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

//TODO: rogue mode

/*HandleCommitState fetches the collections assigned to the staker and creates the leaves required for the merkle tree generation.
Values for only the collections assigned to the staker is fetched for others, 0 is added to the leaves of tree.*/
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
			collectionData, err := utils.UtilsInterface.GetAggregatedDataOfCollection(client, uint16(i), epoch)
			if err != nil {
				return types.CommitData{}, err
			}
			log.Debugf("Data of collection %d:%s", i, collectionData)
			leavesOfTree = append(leavesOfTree, collectionData)
		} else {
			leavesOfTree = append(leavesOfTree, big.NewInt(0))
		}
	}

	return types.CommitData{
		AssignedCollections:    assignedCollections,
		SeqAllottedCollections: seqAllottedCollections,
		Leaves:                 leavesOfTree,
	}, nil
}

/*Commit finally commits the data to the smart contract. It calculates the commitment to send using the merkle tree root and the seed.*/
func (*UtilsStruct) Commit(client *ethclient.Client, seed []byte, root [32]byte, epoch uint32, account types.Account, config types.Configurations) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 0 {
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
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
	})

	log.Debugf("Committing: epoch: %d, commitment: %s, seed: %s, account: %s", epoch, "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(seed), account.Address)

	log.Info("Commitment sent...")
	txn, err := voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}
