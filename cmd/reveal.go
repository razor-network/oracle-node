package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

func (*UtilsStruct) HandleRevealState(client *ethclient.Client, staker bindings.StructsStaker, epoch uint32) error {
	epochLastCommitted, err := razorUtils.GetEpochLastCommitted(client, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

func (*UtilsStruct) Reveal(client *ethclient.Client, config types.Configurations, account types.Account, commitData types.CommitData, secret []byte) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	var (
		values []bindings.StructsAssignedAsset
		proofs [][][32]byte
	)

	for i := 0; i < len(commitData.SeqAllottedCollections); i++ {
		value := bindings.StructsAssignedAsset{
			MedianIndex: uint16(commitData.SeqAllottedCollections[i].Uint64()),
			Value:       uint32(commitData.Leaves[i].Uint64()),
		}
		proof := utils.MerkleInterface.GetProofPath(merkleTree, value.MedianIndex)
		values = append(values, value)
		proofs = append(proofs, proof)
	}

	treeRevealData := bindings.StructsMerkleTree{
		Values: values,
		Proofs: proofs,
		Root:   utils.MerkleInterface.GetMerkleRoot(merkleTree),
	}

	secretBytes32 := [32]byte{}
	copy(secretBytes32[:], secret)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerABI,
		MethodName:      "reveal",
		Parameters:      []interface{}{epoch, treeRevealData, secretBytes32},
	})

	log.Debugf("Revealing vote for epoch: %d secret: %s  commitAccount: %s",
		epoch,
		"0x"+common.Bytes2Hex(secret),
		account.Address)
	log.Info("Revealing votes...")
	txn, err := voteManagerUtils.Reveal(client, txnOpts, epoch, treeRevealData, secretBytes32)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}
