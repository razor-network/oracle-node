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

type StructsJob struct {
	Id           uint8  `json:"id"`
	SelectorType uint8  `json:"selectorType"`
	Weight       uint8  `json:"weight"`
	Power        int8   `json:"power"`
	Name         string `json:"name"`
	Selector     string `json:"selector"`
	Url          string `json:"url"`
}
