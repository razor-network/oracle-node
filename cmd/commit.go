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
)

func HandleCommitState(client *ethclient.Client, address string) []*big.Int {
	data, err := utils.GetActiveAssetsData(client, address)
	if err != nil {
		log.Error("Error in getting active assets: ", err)
		return nil
	}
	log.Debug("Data: ", data)
	return data
}

func Commit(client *ethclient.Client, data []*big.Int, secret []byte, account types.Account, config types.Configurations) error {
	if state, err := utils.GetDelayedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return err
	}

	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		return err
	}

	commitment := solsha3.SoliditySHA3([]string{"uint32", "uint256[]", "bytes32"}, []interface{}{epoch, data, "0x" + hex.EncodeToString(secret)})
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

	log.Debugf("Committing: epoch: %d, commitment: %s, secret: %s, account: %s", epoch, "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(secret), account.Address)

	log.Info("Commitment sent...")
	txn, err := voteManager.Commit(txnOpts, epoch, commitmentToSend)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", txn.Hash())
	if utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash())) == 0 {
		return errors.New("block not mined")
	}
	return nil
}
