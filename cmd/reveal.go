package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

var voteManagerUtils voteManagerInterface
func HandleRevealState(client *ethclient.Client, address string, staker bindings.StructsStaker, epoch uint32, razorUtils utilsInterface) error {
	epochLastCommitted, err := razorUtils.GetEpochLastCommitted(client, address, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

func Reveal(client *ethclient.Client, committedData []*big.Int, secret []byte, account types.Account, commitAccount string, config types.Configurations, razorUtils utilsInterface, voteManagerUtils voteManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	epoch, err := razorUtils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	commitments, err := razorUtils.GetCommitments(client, account.Address)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	if razorUtils.AllZero(commitments) {
		log.Error("Did not commit")
		return core.NilHash, err
	}

	tree, err := razorUtils.GetMerkleTree(committedData)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	root := [32]byte{}
	originalRoot := tree.RootV1()
	copy(root[:], originalRoot)

	secretBytes32 := [32]byte{}
	copy(secretBytes32[:], secret)
	log.Debugf("Revealing vote for epoch: %d  votes: %s  root: %s  secret: %s  commitAccount: %s",
		epoch,
		committedData,
		"0x"+common.Bytes2Hex(originalRoot),
		"0x"+common.Bytes2Hex(secret),
		commitAccount,
	)
	log.Info("Revealing votes...")
	txn, err := voteManagerUtils.Reveal(client, txnOpts, epoch, committedData, secretBytes32)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}
