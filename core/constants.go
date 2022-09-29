package core

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength int64 = 1200
var NumberOfStates int64 = 5
var ChainId = big.NewInt(0x109B4597)
var StateLength = uint64(EpochLength / NumberOfStates)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

//TODO: Check what should be the timeout
var RPCTimeout = time.Second * 10
