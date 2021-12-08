package cmd

import (
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/common"
)

var blockManagerUtils blockManagerInterface

func (utilsStruct UtilsStruct) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	log.Info("Claiming block reward...")
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(options, utilsStruct.packageUtils)
	txn, err := utilsStruct.blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn).Hex())
	return utilsStruct.transactionUtils.Hash(txn), nil
}
