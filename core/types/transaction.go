package types

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TransactionOptions struct {
	Client          *ethclient.Client
	EtherValue      *big.Int
	Amount          *big.Int
	ChainId         *big.Int
	Config          Configurations
	ContractAddress string
	MethodName      string
	Parameters      []interface{}
	ABI             string
	Account         Account
}
