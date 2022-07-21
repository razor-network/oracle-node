//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

//This function approves the transaction if the user has sufficient balance otherwise it fails the transaction
func (*UtilsStruct) Approve(txnArgs types.TransactionOptions) (common.Hash, error) {
	opts := razorUtils.GetOptions()
	allowance, err := tokenManagerUtils.Allowance(txnArgs.Client, &opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
	if err != nil {
		return core.NilHash, err
	}
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Debug("Sufficient allowance, no need to increase")
		return core.NilHash, nil
	} else {
		log.Info("Sending Approve transaction...")
		txnArgs.ContractAddress = core.RAZORAddress
		txnArgs.MethodName = "approve"
		txnArgs.ABI = bindings.RAZORMetaData.ABI
		txnArgs.Parameters = []interface{}{common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount}
		txnOpts := razorUtils.GetTxnOpts(txnArgs)
		txn, err := tokenManagerUtils.Approve(txnArgs.Client, txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
		txnHash := transactionUtils.Hash(txn)
		if err != nil {
			return core.NilHash, err
		}
		log.Info("Txn Hash: ", txnHash.Hex())
		return txnHash, nil
	}
}
