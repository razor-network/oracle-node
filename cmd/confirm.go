package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
)

var blockManagerUtils blockManagerInterface

func (utilsStruct UtilsStruct) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	log.Info("Claiming block reward...")
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(options)
	txn, err := utilsStruct.blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn).Hex())
	return utilsStruct.transactionUtils.Hash(txn), nil
}
