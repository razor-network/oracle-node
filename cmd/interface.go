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
)

type utilsInterface interface {
	GetOptions(bool, string, string) bind.CallOpts
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	WaitForBlockCompletion(*ethclient.Client, string) int
	WaitForCommitState(*ethclient.Client, string, string) (uint32, error)
	AssignPassword(*pflag.FlagSet) string
	GetDefaultPath() (string, error)
}

type tokenManagerInterface interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type stakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
}

type accountInterface interface {
	CreateAccount(path string, password string) accounts.Account
}
