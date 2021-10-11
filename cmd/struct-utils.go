package cmd

import (
	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"
)

type Utils struct{}
type TokenManagerUtils struct{}
type TransactionUtils struct{}
type StakeManagerUtils struct{}
type AssetManagerUtils struct{}
type AccountUtils struct{}
type FlagSetUtils struct{}
type UtilsCmd struct{}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return utils.GetOptions(pending, from, blockNumber)
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return utils.GetTxnOpts(transactionData)
}

func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (u Utils) WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitState(client, accountAddress, action)
}

func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u Utils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u Utils) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

func (u Utils) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetStakerId(client, address)
}

func (u Utils) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetConfigData() (types.Configurations, error) {
	return GetConfigData()
}

func (u Utils) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func (tokenManagerUtils TokenManagerUtils) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Allowance(opts, owner, spender)
}

func (tokenManagerUtils TokenManagerUtils) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Approve(opts, spender, amount)
}

func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}

func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Delegate(opts, epoch, stakerId, amount)
}

func (stakeManagerUtils StakeManagerUtils) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetDelegationAcceptance(opts, status)
}

func (stakeManagerUtils StakeManagerUtils) SetCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetCommission(opts, commission)
}

func (stakeManagerUtils StakeManagerUtils) DecreaseCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.DecreaseCommission(opts, commission)
}

func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, power int8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateJob(opts, power, name, selector, url)
}

func (account AccountUtils) CreateAccount(path string, password string) ethAccounts.Account {
	return accounts.CreateAccount(path, password)
}

func (flagSetUtils FlagSetUtils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (flagSetUtils FlagSetUtils) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

func (flagSetUtils FlagSetUtils) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

func (flagSetUtils FlagSetUtils) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

func (flagSetUtils FlagSetUtils) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

func (flagSetUtils FlagSetUtils) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

func (flagSetUtils FlagSetUtils) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("commission")
}

func (cmdUtils UtilsCmd) SetCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	return SetCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
}

func (cmdUtils UtilsCmd) DecreaseCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	return DecreaseCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
}
