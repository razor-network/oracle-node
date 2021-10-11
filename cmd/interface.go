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

type utilsInterface interface {
	GetOptions(bool, string, string) bind.CallOpts
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	WaitForBlockCompletion(*ethclient.Client, string) int
	WaitForCommitState(*ethclient.Client, string, string) (uint32, error)
	AssignPassword(*pflag.FlagSet) string
	GetDefaultPath() (string, error)
	GetAmountInDecimal(*big.Int) *big.Float
	ConnectToClient(string) *ethclient.Client
	GetStakerId(*ethclient.Client, string) (uint32, error)
	GetStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetUpdatedStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetConfigData() (types.Configurations, error)
	ParseBool(str string) (bool, error)
}

type tokenManagerInterface interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type assetManagerInterface interface {
	CreateJob(*ethclient.Client, *bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error)
}

type stakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	Delegate(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)
	SetDelegationAcceptance(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error)
	SetCommission(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)
	DecreaseCommission(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)
}

type accountInterface interface {
	CreateAccount(path string, password string) accounts.Account
}

type flagSetInterface interface {
	GetStringAddress(*pflag.FlagSet) (string, error)
	GetStringName(*pflag.FlagSet) (string, error)
	GetStringUrl(*pflag.FlagSet) (string, error)
	GetStringSelector(*pflag.FlagSet) (string, error)
	GetInt8Power(*pflag.FlagSet) (int8, error)
	GetStringStatus(*pflag.FlagSet) (string, error)
	GetUint8Commission(*pflag.FlagSet) (uint8, error)
}

type utilsCmdInterface interface {
	SetCommission(*ethclient.Client, uint32, *bind.TransactOpts, uint8, utilsInterface, stakeManagerInterface, transactionInterface) error
	DecreaseCommission(*ethclient.Client, uint32, *bind.TransactOpts, uint8, utilsInterface, stakeManagerInterface, transactionInterface, utilsCmdInterface) error
	DecreaseCommissionPrompt() bool
}
