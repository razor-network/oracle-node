package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var StateLength uint64 = 180
var EpochLength int64 = 900

var NumberOfStates int64 = 5
var ChainId = big.NewInt(80001)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
