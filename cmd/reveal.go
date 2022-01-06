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

func (*UtilsStructMockery) HandleRevealState(client *ethclient.Client, staker bindings.StructsStaker, epoch uint32) error {
	epochLastCommitted, err := razorUtilsMockery.GetEpochLastCommitted(client, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

func (*UtilsStructMockery) Reveal(client *ethclient.Client, committedData []*big.Int, secret []byte, account types.Account, commitAccount string, config types.Configurations) (common.Hash, error) {
	if state, err := razorUtilsMockery.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	epoch, err := razorUtilsMockery.GetEpoch(client)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	commitments, err := razorUtilsMockery.GetCommitments(client, account.Address)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	if razorUtilsMockery.AllZero(commitments) {
		log.Error("Did not commit")
		return core.NilHash, nil
	}

	secretBytes32 := [32]byte{}
	copy(secretBytes32[:], secret)

	txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerABI,
		MethodName:      "reveal",
		Parameters:      []interface{}{epoch, committedData, secretBytes32},
	})

	log.Debugf("Revealing vote for epoch: %d  votes: %s secret: %s  commitAccount: %s",
		epoch,
		committedData,
		"0x"+common.Bytes2Hex(secret),
		commitAccount,
	)
	log.Info("Revealing votes...")
	txn, err := voteManagerUtilsMockery.Reveal(client, txnOpts, epoch, committedData, secretBytes32)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}
