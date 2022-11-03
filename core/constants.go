// This file contains constants, DO NOT MODIFY.

package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var EpochLength uint64 = 1200
var NumberOfStates uint64 = 5
var ChainId = big.NewInt(132333505628089)
var StateLength = EpochLength / NumberOfStates
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

var DefaultProvider = "http://127.0.0.1:8545"
var DefaultGasMultiplier = 1.0
var DefaultBufferPercent = 20
var DefaultGasPrice = 1
var DefaultWaitTime = 1
var DefaultGasLimit = 2
var DefaultRPCTimeout = 10
var DefaultLogLevel = ""

//DisputeGasMultiplier is a constant gasLimitMultiplier to increase gas Limit for function `disputeCollectionIdShouldBeAbsent` and `disputeCollectionIdShouldBePresent`
var DisputeGasMultiplier float32 = 5.5

// Following are the constants which will be used to derive different file paths

var DataFileDirectory = "data_files"
var CommitDataFile = "_commitData.json"
var ProposeDataFile = "_proposeData.json"
var DisputeDataFile = "_disputeData.json"
var AssetsDataFile = "assets.json"
var ConfigFile = "razor.yaml"
var LogFileDirectory = "logs"
var DefaultPathName = ".razor"

//LoggerTimeout is threshold number of seconds after which logger will time out from fetching blockNumber
var LoggerTimeout = 10

// LoggerTimeoutErr is the custom error message that would be displayed as a field in logs if logger times out
var LoggerTimeoutErr = "Logger Timeout, error in fetching block number"
