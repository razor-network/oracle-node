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
	AssignPassword(*pflag.FlagSet) string
	ConnectToClient(string) *ethclient.Client
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	AssignAmountInWei(*pflag.FlagSet) *big.Int
	CheckAmountAndBalance(*big.Int, *big.Int) *big.Int
	GetAmountInDecimal(*big.Int) *big.Float
	WaitForCommitState(*ethclient.Client, string, string) (uint32, error)
	GetDefaultPath() (string, error)
	PasswordPrompt() string
	ParseBool(str string) (bool, error)
}

type tokenManagerInterface interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
	Transfer(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type assetManagerInterface interface {
	CreateJob(*ethclient.Client, *bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error)
	SetAssetStatus(*ethclient.Client, *bind.TransactOpts, bool, uint8) (*Types.Transaction, error)
	GetActiveStatus(*ethclient.Client, *bind.CallOpts, uint8) (bool, error)
}

type stakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	ResetLock(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	Delegate(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)
}

type accountInterface interface {
	CreateAccount(path string, password string) accounts.Account
}

type keystoreInterface interface {
	Accounts(string) []accounts.Account
}

type flagSetInterface interface {
	GetStringFrom(*pflag.FlagSet) (string, error)
	GetStringTo(*pflag.FlagSet) (string, error)
	GetStringAddress(*pflag.FlagSet) (string, error)
	GetUint32StakerId(*pflag.FlagSet) (uint32, error)
	GetStringName(*pflag.FlagSet) (string, error)
	GetStringUrl(*pflag.FlagSet) (string, error)
	GetStringSelector(*pflag.FlagSet) (string, error)
	GetInt8Power(*pflag.FlagSet) (int8, error)
	GetUint8AssetId(*pflag.FlagSet) (uint8, error)
	GetStringStatus(*pflag.FlagSet) (string, error)
}

type cmdUtilsInterface interface {
	CheckCurrentStatus(*ethclient.Client, string, uint8, utilsInterface, assetManagerInterface) (bool, error)
}
