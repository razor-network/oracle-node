package types

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TransactionOptions struct {
	Client            *ethclient.Client
	Password          string
	EtherValue        *big.Int
	Amount            *big.Int
	AccountAddress    string
	ChainId           *big.Int
	GasMultiplier     float32
}
