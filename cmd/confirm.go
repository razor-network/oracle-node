package cmd

import (
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/common"
)

func (*UtilsStructMockery) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	log.Info("Claiming block reward...")
	txnOpts := razorUtilsMockery.GetTxnOpts(options)
	txn, err := blockManagerUtilsMockery.ClaimBlockReward(options.Client, txnOpts)
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn).Hex())
	return transactionUtilsMockery.Hash(txn), nil
}
