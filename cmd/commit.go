package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
)

func HandleCommitState(client *ethclient.Client, address string) []*big.Int {
	jobs, collections, err := utils.GetActiveAssets(client, address)
	if err != nil {
		log.Error("Error in getting active jobs: ", err)
		return nil
	}

	dataFromJobs := utils.GetDataToCommitFromJobs(jobs)
	var dataFromCollections []*big.Int
	for _, collection := range collections {
		data, err := utils.Aggregate(client, address, collection)
		if err != nil {
			log.Error(err)
		}
		dataFromCollections = append(dataFromCollections, data)
	}
	data := append(dataFromJobs, dataFromCollections...)
	log.Info("Data", data)

	return data
}

func Commit(client *ethclient.Client, data []*big.Int, secret []byte, account types.Account, config types.Configurations) error {
	if state, err := utils.GetDelayedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return err
	}

	root, err := utils.GetMerkleTreeRoot(data)
	if err != nil {
		return err
	}

	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		return err
	}

	// Required if 2 or more instances of same staker is running and one of them has already committed in the current epoch
	commitments, err := utils.GetCommitments(client, account.Address, epoch)
	if err != nil {
		return err
	}
	if !utils.AllZero(commitments) {
		return errors.New("already committed")
	}

	commitment := solsha3.SoliditySHA3([]string{"uint256", "bytes32", "bytes32"}, []interface{}{epoch.String(), "0x" + hex.EncodeToString(root), "0x" + hex.EncodeToString(secret)})

	voteManager := utils.GetVoteManager(client)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)

	log.Infof("Committing: epoch: %s, root: %s, commitment: %s, secret: %s, account: %s", epoch, "0x"+hex.EncodeToString(root), "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(secret), account.Address)

	txn, err := voteManager.Commit(txnOpts, epoch, commitmentToSend)
	if err != nil {
		return err
	}

	log.Info("Commitment sent...")
	log.Info("Txn Hash: ", txn.Hash())
	if utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash())) == 0 {
		log.Error("Commit failed....")
	}
	return nil
}
