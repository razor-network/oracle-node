//Package core contains the most important values of project
package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength int64 = 1800
var NumberOfStates int64 = 5
var StateBuffer uint64 = 5
var ChainId = big.NewInt(0x785B4B9847B9)
var StateLength = uint64(EpochLength / NumberOfStates)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
var BatchSize = 1000
var NumRoutines = 10
var MaxIterations = 10000000
