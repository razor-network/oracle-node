package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"razor/accounts"
	"razor/core/types"
)

//go:generate mockery --name OptionUtils --output ./mocks/ --case=underscore
//go:generate mockery --name ABIUtils --output ./mocks/ --case=underscore
//go:generate mockery --name Path --output ./mocks/ --case=underscore
//go:generate mockery --name Account --output ./mocks/ --case=underscore
//go:generate mockery --name Bind --output ./mocks/ --case=underscore
//go:generate mockery --name Utils --output ./mocks/ --case=underscore

var Options OptionUtils
var AbiUtils ABIUtils
var PathUtils Path
var AccountUtils Account
var BindUtils Bind
var UtilsInterface Utils

type OptionUtils interface {
	SuggestGasPriceWithRetry(*ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address) (uint64, error)
}

type ABIUtils interface {
	Parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
}

type Utils interface {
	GetGasPrice(*ethclient.Client, types.Configurations) *big.Int
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	GetGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error)
}

type Path interface {
	GetDefaultPath() (string, error)
}

type Account interface {
	GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey
}

type Bind interface {
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
}

type OptionUtilsStruct struct{}
type ABIUtilsStruct struct{}
type PathStruct struct{}
type AccountStruct struct{}
type BindStruct struct{}

type utils struct{}

type OptionsPackageStruct struct {
	Options        OptionUtils
	AbiUtils       ABIUtils
	PathUtils      Path
	AccountUtils   Account
	BindUtils      Bind
	UtilsInterface Utils
}
