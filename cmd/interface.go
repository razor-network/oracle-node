package cmd

import (
	"crypto/ecdsa"
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
	ConvertUintArrayToUint8Array(uintArr []uint) []uint8
	WaitForDisputeOrConfirmState(client *ethclient.Client, accountAddress string, action string) (uint32, error)
	PrivateKeyPrompt() string
	PasswordPrompt() string
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	AssignAmountInWei(*pflag.FlagSet) *big.Int
	CheckAmountAndBalance(*big.Int, *big.Int) *big.Int
	GetAmountInDecimal(*big.Int) *big.Float
	WaitForCommitState(*ethclient.Client, string, string) (uint32, error)
	GetDefaultPath() (string, error)
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
	CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint8, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error)
	AddJobToCollection(*ethclient.Client, *bind.TransactOpts, uint8, uint8) (*Types.Transaction, error)
	UpdateJob(*ethclient.Client, *bind.TransactOpts, uint8, int8, string, string) (*Types.Transaction, error)
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
	ImportECDSA(string, *ecdsa.PrivateKey, string) (accounts.Account, error)
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
	GetUintSliceJobIds(*pflag.FlagSet) ([]uint, error)
	GetUint32Aggregation(set *pflag.FlagSet) (uint32, error)
	GetUint8JobId(*pflag.FlagSet) (uint8, error)
	GetUint8CollectionId(*pflag.FlagSet) (uint8, error)
}

type cryptoInterface interface {
	HexToECDSA(string) (*ecdsa.PrivateKey, error)
}
