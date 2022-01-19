package cmd

import (
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/common"
)

func (*UtilsStruct) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	log.Info("Claiming block reward...")
	txnOpts := razorUtils.GetTxnOpts(options)
	txn, err := blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn).Hex())
	return transactionUtils.Hash(txn), nil
}
