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

type AccountMock struct{}

var GetTokenManagerMock func(*ethclient.Client) *bindings.RAZOR

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(client *ethclient.Client, accountAddress string, action string) (uint32, error)

var AssignPasswordMock func(flagSet *pflag.FlagSet) string

var GetDefaultPathMock func() (string, error)

var AllowanceMock func(*bind.CallOpts, common.Address, common.Address, *ethclient.Client) (*big.Int, error)

var ApproveMock func(*bind.TransactOpts, common.Address, *big.Int, *ethclient.Client) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*bind.TransactOpts, uint32, *big.Int, *ethclient.Client) (*Types.Transaction, error)

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

func (u UtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (tokenManagerMock TokenManagerMock) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, client *ethclient.Client) (*big.Int, error) {
	return AllowanceMock(opts, owner, spender, client)
}

func (tokenManagerMock TokenManagerMock) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int, client *ethclient.Client) (*Types.Transaction, error) {
	return ApproveMock(opts, spender, amount, client)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (account AccountMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}
