package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
)

var blockManagerUtils blockManagerInterface

func ClaimBlockReward(options types.TransactionOptions, razorUtils utilsInterface, blockManagerUtils blockManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	log.Info("Claiming block reward...")
	txnOpts := razorUtils.GetTxnOpts(options)
	txnOpts.GasLimit = 250000
	txn, err := blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn).Hex())
	return transactionUtils.Hash(txn), nil
}
