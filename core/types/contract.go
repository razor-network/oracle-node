package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Job struct {
	Active    bool
	Id        uint8
	AssetType uint8
	Power     int8
	Epoch     uint32
	Creator   common.Address
	Name      string
	Selector  string
	Url       string
}

type Collection struct {
	Id                uint8
	Name              string
	AggregationMethod uint32
	JobIDs            []uint8
	Power             int8
}

type Locks struct {
	Amount        *big.Int
	WithdrawAfter *big.Int
}
