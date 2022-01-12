package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
	"razor/utils"
)

type Utils struct{}
type TransactionUtils struct{}
type StakeManagerUtils struct{}

type UtilsStruct struct {
	razorUtils        utilsInterface
	stakeManagerUtils stakeManagerInterface
	transactionUtils  transactionInterface
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	utilsInterface := utils.StartRazor(utils.OptionsPackageStruct{
		Options:        utils.Options,
		UtilsInterface: utils.UtilsInterface,
	})
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	return utilsInterface.GetTxnOpts(transactionData)
}

func (u Utils) GetEpoch(client *ethclient.Client) (uint32, error) {
	return utils.GetEpoch(client)
}

func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}
