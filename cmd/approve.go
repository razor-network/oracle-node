package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

func (utilsStruct UtilsStruct) approve(txnArgs types.TransactionOptions) (common.Hash, error) {
	opts := utilsStruct.razorUtils.GetOptions()
	allowance, err := utilsStruct.tokenManagerUtils.Allowance(txnArgs.Client, &opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
	if err != nil {
		return common.Hash{0x00}, err
	}
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Debug("Sufficient allowance, no need to increase")
		return common.Hash{0x00}, nil
	} else {
		log.Info("Sending Approve transaction...")
		txnArgs.ContractAddress = core.RAZORAddress
		txnArgs.MethodName = "approve"
		txnArgs.ABI = bindings.RAZORABI
		txnArgs.Parameters = []interface{}{common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount}
		txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)
		txn, err := utilsStruct.tokenManagerUtils.Approve(txnArgs.Client, txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
		if err != nil {
			return common.Hash{0x00}, err
		}
		log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))
		return utilsStruct.transactionUtils.Hash(txn), nil
	}
}
