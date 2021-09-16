package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
)

func approve(txnArgs types.TransactionOptions, razorUtils utilsInterface, tokenManagerUtils tokenManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	opts := razorUtils.GetOptions(false, txnArgs.AccountAddress, "")
	allowance, err := tokenManagerUtils.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress), txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Debug("Sufficient allowance, no need to increase")
		return common.Hash{0x00}, nil
	} else {
		log.Info("Sending Approve transaction...")
		txnOpts := razorUtils.GetTxnOpts(txnArgs)
		txn, err := tokenManagerUtils.Approve(txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount, txnArgs.Client)
		if err != nil {
			return common.Hash{0x00}, err
		}
		log.Info("Txn Hash: ", transactionUtils.Hash(txn))
		return transactionUtils.Hash(txn), nil
	}
}
