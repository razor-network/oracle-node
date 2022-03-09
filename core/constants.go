package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength int64 = 300
var NumberOfStates int64 = 5
var StateLength = uint64(EpochLength / NumberOfStates)

var ChainId = big.NewInt(31337)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
