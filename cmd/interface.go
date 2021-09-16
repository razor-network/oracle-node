package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

type utilsInterface interface {
	GetTokenManager(*ethclient.Client) *bindings.RAZOR
	GetOptions(bool, string, string) bind.CallOpts
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	WaitForBlockCompletion(*ethclient.Client, string) int
	WaitForCommitState(*ethclient.Client, string, string) (uint32, error)
}

type tokenManagerInterface interface {
	Allowance(*bind.CallOpts, common.Address, common.Address, *ethclient.Client) (*big.Int, error)
	Approve(*bind.TransactOpts, common.Address, *big.Int, *ethclient.Client) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type stakeManagerInterface interface {
	Stake(*bind.TransactOpts, uint32, *big.Int, *ethclient.Client) (*Types.Transaction, error)
}
