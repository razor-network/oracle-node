package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
)

type UtilsMock struct{}

type TokenManagerMock struct{}

type TransactionMock struct{}

type StakeManagerMock struct{}

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(client *ethclient.Client, accountAddress string, action string) (uint32, error)

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

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

func (tokenManagerMock TokenManagerMock) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return AllowanceMock(client, opts, owner, spender)
}

func (tokenManagerMock TokenManagerMock) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	return ApproveMock(client, opts, spender, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}
