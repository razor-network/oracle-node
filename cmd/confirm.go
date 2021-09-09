package cmd

import (
	log "github.com/sirupsen/logrus"
	"razor/core/types"
	"razor/utils"
)

func ClaimBlockReward(options types.TransactionOptions) {
	blockManager := utils.GetBlockManager(options.Client)
	txn, err := blockManager.ClaimBlockReward(utils.GetTxnOpts(options))
	if err != nil {
		log.Error("Error in claiming block reward: ", err)
		return
	}
	log.Info("Claiming block reward...")
	log.Info("Txn Hash: ", txn.Hash().Hex())
	utils.WaitForBlockCompletion(options.Client, txn.Hash().Hex())
}
