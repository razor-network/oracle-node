package types

import (
	"math/big"
)

type TransactionOptions struct {
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
