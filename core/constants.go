// This file contains constants, DO NOT MODIFY.

package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EpochLength    uint64 = 300
	NumberOfStates uint64 = 5
	StateLength           = EpochLength / NumberOfStates
)

// ChainId corresponds to the SKALE chain
var ChainId = big.NewInt(0x561bf78b)

const MaxRetries uint = 8

var NilHash = common.Hash{0x00}

const BlockCompletionTimeout = 30

//Following are the default config values for all the config parameters
const (
	DefaultGasMultiplier    float32 = 1.0
	DefaultBufferPercent    int32   = 20
	DefaultGasPrice         int32   = 0
	DefaultWaitTime         int32   = 5
	DefaultGasLimit         float32 = 2
	DefaultGasLimitOverride uint64  = 30000000
	DefaultRPCTimeout       int64   = 10
	DefaultHTTPTimeout      int64   = 10
	DefaultLogLevel                 = ""
)

//BufferStateSleepTime is the sleeping time whenever buffer state hits
const BufferStateSleepTime int32 = 2

//Following are the default logFile parameters in config
const (
	DefaultLogFileMaxSize    = 200
	DefaultLogFileMaxBackups = 10
	DefaultLogFileMaxAge     = 60
)

//DisputeGasMultiplier is a constant gasLimitMultiplier to increase gas Limit for function `disputeCollectionIdShouldBeAbsent` and `disputeCollectionIdShouldBePresent`
const DisputeGasMultiplier float32 = 5.5

// Following are the constants which will be used to derive different file paths

const (
	DataFileDirectory = "data_files"
	CommitDataFile    = "_commitData.json"
	ProposeDataFile   = "_proposeData.json"
	DisputeDataFile   = "_disputeData.json"
	AssetsDataFile    = "assets.json"
	ConfigFile        = "razor.yaml"
	LogFileDirectory  = "logs"
	DefaultPathName   = ".razor"
)

//BlockNumberInterval is the interval in seconds after which blockNumber needs to be calculated again
const BlockNumberInterval = 5

//APIKeyRegex will be used as a regular expression to be matched in job Urls
const APIKeyRegex = `\$\{(.+?)\}`

// Following are the constants which defines retry attempts and retry delay if there is an error in processing request
const (
	ProcessRequestRetryAttempts uint  = 2
	ProcessRequestRetryDelay    int64 = 2
)

//SwitchClientDuration is the time after which alternate client from secondary RPC will be switched back to client from primary RPC
const SwitchClientDuration = 5 * EpochLength

const (
	// HexReturnType is the ReturnType for a job if that job returns a hex value
	HexReturnType = "hex"

	// HexArrayReturnType is the ReturnType for a job if that job returns a hex array value
	HexArrayReturnType = "^hexArray\\[\\d+\\]$"

	// HexArrayExtractIndexRegex will be used as a regular expression to extract index from hexArray return type
	HexArrayExtractIndexRegex = `^hexArray\[(\d+)\]$`
)

// Following are the constants which helps in calculating iteration for a staker
const (
	BatchSize     = 1000
	NumRoutines   = 10
	MaxIterations = 10000000
)

// Following are the constants used in custom http.Transport configuration for the common HTTP client that we use for all the requests
const (
	HTTPClientMaxIdleConns        = 15
	HTTPClientMaxIdleConnsPerHost = 5
)

const GetStakeSnapshotMethod = "getStakeSnapshot"

// Following are the event names that nodes will listen to in order to update the jobs/collections in the cache
const (
	JobCreatedEvent               = "JobCreated"
	CollectionCreatedEvent        = "CollectionCreated"
	JobUpdatedEvent               = "JobUpdated"
	CollectionUpdatedEvent        = "CollectionUpdated"
	CollectionActivityStatusEvent = "CollectionActivityStatus"
)
