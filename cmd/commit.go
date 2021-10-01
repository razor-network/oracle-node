package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"razor/core"
	"razor/core/types"
)

var voteManagerUtils voteManagerInterface

func HandleCommitState(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	data, err := razorUtils.GetActiveAssetsData(client, address, epoch)
	if err != nil {
		return nil, err
	}
	log.Debug("Data: ", data)
	return data, nil
}

func Commit(client *ethclient.Client, data []*big.Int, secret []byte, account types.Account, config types.Configurations, razorUtils utilsInterface, voteManagerUtils voteManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	epoch, err := razorUtils.GetEpoch(client, account.Address)
	if err != nil {
		return core.NilHash, err
	}

	commitment := solsha3.SoliditySHA3([]string{"uint32", "uint256[]", "bytes32"}, []interface{}{epoch, data, "0x" + hex.EncodeToString(secret)})
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
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
	txn, err := voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}
