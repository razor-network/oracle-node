//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
)

//This function approves the transaction if the user has sufficient balance otherwise it fails the transaction
func (*UtilsStruct) Approve(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions) (common.Hash, error) {
	allowance, err := razorUtils.Allowance(rpcParameters, common.HexToAddress(txnArgs.Account.Address), common.HexToAddress(core.StakeManagerAddress))
	if err != nil {
		return core.NilHash, err
	}
	log.Debugf("Approve: Allowance amount: %d, Amount to get approved: %d", allowance, txnArgs.Amount)
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Debug("Sufficient allowance, no need to increase")
		return core.NilHash, nil
	} else {
		log.Info("Sending Approve transaction...")
		txnArgs.ContractAddress = core.RAZORAddress
		txnArgs.MethodName = "approve"
		txnArgs.ABI = bindings.RAZORMetaData.ABI
		txnArgs.Parameters = []interface{}{common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount}
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
		if err != nil {
			return core.NilHash, err
		}

		client, err := rpcParameters.RPCManager.GetBestRPCClient()
		if err != nil {
			return core.NilHash, err
		}

		log.Debug("Executing Approve transaction with amount: ", txnArgs.Amount)
		txn, err := tokenManagerUtils.Approve(client, txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
		if err != nil {
			return core.NilHash, err
		}
		txnHash := transactionUtils.Hash(txn)
		log.Info("Txn Hash: ", txnHash.Hex())
		return txnHash, nil
	}
}
