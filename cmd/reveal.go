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

func (*UtilsStruct) Reveal(client *ethclient.Client, committedData []*big.Int, secret []byte, account types.Account, commitAccount string, config types.Configurations) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	epoch, err := razorUtils.GetEpoch(client)
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
		return core.NilHash, nil
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
		Parameters:      []interface{}{epoch, committedData, secretBytes32},
	})

	log.Debugf("Revealing vote for epoch: %d  votes: %s secret: %s  commitAccount: %s",
		epoch,
		committedData,
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
