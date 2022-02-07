package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var StateLength uint64 = 360
var EpochLength int64 = 1800

var NumberOfStates int64 = 5
var ChainId = big.NewInt(0x17ac300421d1b)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
