//Package types include the different user defined items of possible different types in a single type
package types

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TransactionOptions struct {
	Client          *ethclient.Client
	Password        string
	EtherValue      *big.Int
	Amount          *big.Int
	AccountAddress  string
	ChainId         *big.Int
	Config          Configurations
	ContractAddress string
	MethodName      string
	Parameters      []interface{}
	ABI             string
}
