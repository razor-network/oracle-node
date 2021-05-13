package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Job struct {
	Id        *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Fulfilled bool
	Result    *big.Int
}
