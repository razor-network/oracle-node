package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

type UtilsMock struct{}

type TokenManagerMock struct{}

type TransactionMock struct{}

type AssetManagerMock struct{}

type StakeManagerMock struct{}

type AccountMock struct{}

type FlagSetMock struct{}

type UtilsCmdMock struct{}

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(*ethclient.Client, string, string) (uint32, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var GetDefaultPathMock func() (string, error)

var GetAmountInDecimalMock func(amountInWei *big.Int) *big.Float

var ConnectToClientMock func(string) *ethclient.Client

var GetStakerIdMock func(*ethclient.Client, string) (uint32, error)

var GetStakerMock func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)

var GetUpdatedStakerMock func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)

var GetConfigDataMock func() (types.Configurations, error)

var ParseBoolMock func(string) (bool, error)

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

var DelegateMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)

var SetDelegationAcceptanceMock func(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error)

var SetCommissionContractMock func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)

var DecreaseCommissionContractMock func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)

var CreateAccountMock func(string, string) accounts.Account

var CreateJobMock func(*bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error)

var GetStringAddressMock func(*pflag.FlagSet) (string, error)

var GetStringNameMock func(*pflag.FlagSet) (string, error)

var GetStringUrlMock func(*pflag.FlagSet) (string, error)

var GetStringSelectorMock func(*pflag.FlagSet) (string, error)

var GetInt8PowerMock func(*pflag.FlagSet) (int8, error)

var GetStringStatusMock func(*pflag.FlagSet) (string, error)

var GetUint8CommissionMock func(*pflag.FlagSet) (uint8, error)

var SetCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, utilsInterface, stakeManagerInterface, transactionInterface) error

var DecreaseCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, utilsInterface, stakeManagerInterface, transactionInterface) error

func (u UtilsMock) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return GetOptionsMock(pending, from, blockNumber)
}

func (u UtilsMock) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return GetTxnOptsMock(transactionData)
}

func (u UtilsMock) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return WaitForBlockCompletionMock(client, hashToRead)
}

func (u UtilsMock) WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitStateMock(client, accountAddress, action)
}

func (u UtilsMock) AssignPassword(flagSet *pflag.FlagSet) string {
	return AssignPasswordMock(flagSet)
}

func (u UtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (u UtilsMock) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return GetAmountInDecimalMock(amountInWei)
}

func (u UtilsMock) ConnectToClient(provider string) *ethclient.Client {
	return ConnectToClientMock(provider)
}

func (u UtilsMock) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return GetStakerIdMock(client, address)
}

func (u UtilsMock) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return GetStakerMock(client, address, stakerId)
}

func (u UtilsMock) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return GetUpdatedStakerMock(client, address, stakerId)
}

func (u UtilsMock) GetConfigData() (types.Configurations, error) {
	return GetConfigDataMock()
}

func (u UtilsMock) ParseBool(str string) (bool, error) {
	return ParseBoolMock(str)
}

func (tokenManagerMock TokenManagerMock) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return AllowanceMock(client, opts, owner, spender)
}

func (tokenManagerMock TokenManagerMock) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	return ApproveMock(client, opts, spender, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (assetManagerMock AssetManagerMock) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, power int8, name string, selector string, url string) (*Types.Transaction, error) {
	return CreateJobMock(opts, power, name, selector, url)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}

func (stakeManagerMock StakeManagerMock) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	return DelegateMock(client, opts, epoch, stakerId, amount)
}

func (stakeManagerMock StakeManagerMock) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	return SetDelegationAcceptanceMock(client, opts, status)
}

func (stakeManagerMock StakeManagerMock) SetCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	return SetCommissionContractMock(client, opts, commission)
}

func (stakeManagerMock StakeManagerMock) DecreaseCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	return DecreaseCommissionContractMock(client, opts, commission)
}

func (account AccountMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}

func (flagSetMock FlagSetMock) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return GetStringNameMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return GetStringAddressMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return GetStringUrlMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return GetStringSelectorMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return GetInt8PowerMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return GetStringStatusMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8CommissionMock(flagSet)
}

func (utilsCmdMock UtilsCmdMock) SetCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	return SetCommissionMock(client, stakerId, opts, commission, razorUtils, stakeManagerUtils, transactionUtils)
}

func (utilsCmdMock UtilsCmdMock) DecreaseCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	return DecreaseCommissionMock(client, stakerId, opts, commission, razorUtils, stakeManagerUtils, transactionUtils)
}
