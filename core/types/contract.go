package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
	AssetType *big.Int
}

type Collection struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Result            *big.Int
}
