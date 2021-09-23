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

type StakeManagerMock struct{}

type AccountMock struct{}

var GetTokenManagerMock func(*ethclient.Client) *bindings.RAZOR

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(*ethclient.Client, string, string) (uint32, error)

var GetDefaultPathMock func() (string, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var ConnectToClientMock func(string) *ethclient.Client

var FetchBalanceMock func(*ethclient.Client, string) (*big.Int, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet) *big.Int

var CheckAmountAndBalanceMock func(amountInWei *big.Int, balance *big.Int) *big.Int

var GetAmountInDecimalMock func(*big.Int) *big.Float

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var TransferMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

var CreateAccountMock func(string, string) accounts.Account

func (u UtilsMock) GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	return GetTokenManagerMock(client)
}

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

func (u UtilsMock) ConnectToClient(provider string) *ethclient.Client {
	return ConnectToClientMock(provider)
}

func (u UtilsMock) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return FetchBalanceMock(client, accountAddress)
}

func (u UtilsMock) AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	return AssignAmountInWeiMock(flagSet)
}

func (u UtilsMock) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return CheckAmountAndBalanceMock(amountInWei, balance)
}

func (u UtilsMock) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return GetAmountInDecimalMock(amountInWei)
}

func (u UtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (tokenManagerMock TokenManagerMock) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return AllowanceMock(client, opts, owner, spender)
}

func (tokenManagerMock TokenManagerMock) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	return ApproveMock(client, opts, spender, amount)
}

func (tokenManagerMock TokenManagerMock) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	return TransferMock(client, opts, recipient, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}

func (account AccountMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}
