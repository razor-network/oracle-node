package cmd

import (
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

var GetTokenManagerMock func(*ethclient.Client) *bindings.RAZOR

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(client *ethclient.Client, accountAddress string, action string) (uint32, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var ConnectToClientMock func(string) *ethclient.Client

var FetchBalanceMock func(*ethclient.Client, string) (*big.Int, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet) *big.Int

var CheckAmountAndBalanceMock func(amountInWei *big.Int, balance *big.Int) *big.Int

var GetAmountInDecimalMock func(*big.Int) *big.Float

var AllowanceMock func(*bind.CallOpts, common.Address, common.Address, *ethclient.Client) (*big.Int, error)

var ApproveMock func(*bind.TransactOpts, common.Address, *big.Int, *ethclient.Client) (*Types.Transaction, error)

var TransferMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*bind.TransactOpts, uint32, *big.Int, *ethclient.Client) (*Types.Transaction, error)

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
func (tokenManagerMock TokenManagerMock) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, client *ethclient.Client) (*big.Int, error) {
	return AllowanceMock(opts, owner, spender, client)
}

func (tokenManagerMock TokenManagerMock) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int, client *ethclient.Client) (*Types.Transaction, error) {
	return ApproveMock(opts, spender, amount, client)
}

func (tokenManagerMock TokenManagerMock) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	return TransferMock(client, opts, recipient, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}
