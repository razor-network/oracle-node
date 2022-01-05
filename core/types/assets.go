package types

import (
	"math/big"
	"razor/pkg/bindings"
)

type Job struct {
	Id           uint8
	SelectorType uint8
	Weight       uint8
	Power        uint8
	Name         string
	Selector     string
	Url          string
}

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
