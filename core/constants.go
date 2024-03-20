// This file contains constants, DO NOT MODIFY.

package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength uint64 = 300
var NumberOfStates uint64 = 5
var StateLength = EpochLength / NumberOfStates

// ChainId corresponds to the SKALE chain
var ChainId = big.NewInt(0x561bf78b)

var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

//Following are the default config values for all the config parameters

var DefaultGasMultiplier float32 = 1.0
var DefaultBufferPercent int32 = 0
var DefaultGasPrice int32 = 0
var DefaultWaitTime int32 = 1
var DefaultGasLimit float32 = 2
var DefaultGasLimitOverride uint64 = 30000000
var DefaultRPCTimeout int64 = 10
var DefaultHTTPTimeout int64 = 10
var DefaultLogLevel = ""

//BufferStateSleepTime is the sleeping time whenever buffer state hits
var BufferStateSleepTime int32 = 2

//Following are the default logFile parameters in config

var DefaultLogFileMaxSize = 200
var DefaultLogFileMaxBackups = 52
var DefaultLogFileMaxAge = 365

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

//BlockNumberInterval is the interval in seconds after which blockNumber needs to be calculated again
var BlockNumberInterval = 5

//APIKeyRegex will be used as a regular expression to be matched in job Urls
var APIKeyRegex = `\$\{(.+?)\}`

// Following are the constants which defines retry attempts and retry delay if there is an error in processing request

var ProcessRequestRetryAttempts uint = 2
var ProcessRequestRetryDelay = 2

//SwitchClientDuration is the time after which alternate client from secondary RPC will be switched back to client from primary RPC
var SwitchClientDuration = 5 * EpochLength

// HexReturnType is the ReturnType for a job if that job returns a hex value
var HexReturnType = "hex"

// HexArrayReturnType is the ReturnType for a job if that job returns a hex array value
var HexArrayReturnType = "^hexArray\\[\\d+\\]$"

// HexArrayExtractIndexRegex will be used as a regular expression to extract index from hexArray return type
var HexArrayExtractIndexRegex = `^hexArray\[(\d+)\]$`

// Following are the constants which helps in calculating iteration for a staker

var BatchSize = 1000
var NumRoutines = 10
var MaxIterations = 10000000

// Following are the constants used in custom http.Transport configuration for the common HTTP client that we use for all the requests

var HTTPClientMaxIdleConns = 15
var HTTPClientMaxIdleConnsPerHost = 5
