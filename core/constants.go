package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var EpochLength int64 = 1200
var NumberOfStates int64 = 5
var ChainId = big.NewInt(132333505628089)
var StateLength = uint64(EpochLength / NumberOfStates)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
