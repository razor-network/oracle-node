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
	Amount      *big.Int
	UnlockAfter *big.Int
}

type StructsJob struct {
	Id           uint16 `json:"id"`
	SelectorType uint8  `json:"selectorType"`
	Weight       uint8  `json:"weight"`
	Power        int8   `json:"power"`
	Name         string `json:"name"`
	Selector     string `json:"selector"`
	Url          string `json:"url"`
}

type AssignedAsset struct {
	LeafId uint16   `json:"leafId"`
	Value  *big.Int `json:"value"`
}

type MedianResult struct {
	LeafId uint16
	Median *big.Int
}

type CustomJob struct {
	URL      string `json:"URL"`
	Name     string `json:"name"`
	Selector string `json:"selector"`
	Power    int8   `json:"power"`
	Weight   uint8  `json:"weight"`
}

type DataSourceURL struct {
	Type       string                 `json:"type"`
	URL        string                 `json:"url"`
	Body       map[string]interface{} `json:"body"`
	Header     map[string]string      `json:"header"`
	ReturnType string                 `json:"returnType"`
}
