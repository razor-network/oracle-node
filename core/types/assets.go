package types

import (
	"math/big"
	"razor/pkg/bindings"
)

type Collection struct {
	Active            bool
	Id                uint8
	AssetIndex        uint8
	Power             int8
	AggregationMethod uint32
	Name              string
}

type Asset struct {
	Job        bindings.StructsJob
	Collection bindings.StructsCollection
}

type Locks struct {
	Amount        *big.Int
	Commission    *big.Int
	WithdrawAfter *big.Int
}
