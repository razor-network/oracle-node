package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HandleRevealState(client *ethclient.Client, address string, staker bindings.StructsStaker, epoch uint32) error {
	epochLastCommitted, err := utils.GetEpochLastCommitted(client, address, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

func Reveal(client *ethclient.Client, committedData []*big.Int, secret []byte, account types.Account, commitAccount string, config types.Configurations) {
	if state, err := utils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return
	}

	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error(err)
		return
	}
	commitments, err := utils.GetCommitments(client, account.Address)
	if err != nil {
		log.Error(err)
		return
	}
	if utils.AllZero(commitments) {
		log.Error("Did not commit")
		return
	}

	tree, err := utils.GetMerkleTree(committedData)
	if err != nil {
		log.Error(err)
		return
	}

	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	voteManager := utils.GetVoteManager(client)

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
	txn, err := voteManager.Reveal(txnOpts, epoch, committedData, secretBytes32)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Txn Hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}
