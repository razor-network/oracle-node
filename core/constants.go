package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var EpochLength int64 = 1200
var NumberOfStates int64 = 5
var ChainId = big.NewInt(0x109B4597)
var StateLength = uint64(EpochLength / NumberOfStates)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

//Following are the default config values for all the config parameters

var DefaultProvider = "http://127.0.0.1:8545"
var DefaultGasMultiplier = 1.0
var DefaultBufferPercent = 20
var DefaultGasPrice = 1
var DefaultWaitTime = 1
var DefaultGasLimit = 2
var DefaultRPCTimeout = 10
var DefaultLogLevel = ""

//BlockNumberInterval is the interval in seconds after which blockNumber needs to be calculated again
var BlockNumberInterval = 5
