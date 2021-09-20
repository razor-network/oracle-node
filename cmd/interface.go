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

type utilsInterface interface {
	GetTokenManager(*ethclient.Client) *bindings.RAZOR
	GetOptions(bool, string, string) bind.CallOpts
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	WaitForBlockCompletion(*ethclient.Client, string) int
	AssignPassword(*pflag.FlagSet) string
 ConnectToClient( string) *ethclient.Client
 FetchBalance( *ethclient.Client,  string) (*big.Int, error)
 AssignAmountInWei( *pflag.FlagSet) *big.Int
 CheckAmountAndBalance( *big.Int,  *big.Int) *big.Int
 GetAmountInDecimal( *big.Int) *big.Float
}

type tokenManagerInterface interface {
	Allowance(*bind.CallOpts, common.Address, common.Address, *ethclient.Client) (*big.Int, error)
	Approve(*bind.TransactOpts, common.Address, *big.Int, *ethclient.Client) (*Types.Transaction, error)
	Transfer( *ethclient.Client, *bind.TransactOpts,  common.Address,  *big.Int) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}
