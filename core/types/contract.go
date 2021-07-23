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
	Active    bool
	Creator   common.Address
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

type Locks struct {
	Amount        *big.Int
	WithdrawAfter *big.Int
}
