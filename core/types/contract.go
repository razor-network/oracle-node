package types

import (
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
