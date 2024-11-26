// This file contains constants, DO NOT MODIFY.

package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EpochLength    uint64 = 450
	NumberOfStates uint64 = 5
	StateLength           = EpochLength / NumberOfStates
)

// ChainId corresponds to the EUROPA chain
var ChainId = big.NewInt(0x79f99296)

const (
	MaxRetries         uint  = 3
	RetryDelayDuration int64 = 1
)

var NilHash = common.Hash{0x00}

const (
	BlockCompletionAttempts          = 4
	BlockCompletionAttemptRetryDelay = 2
	BlockCompletionTimeout           = 15
)

//Following are the default config values for all the config parameters
const (
	DefaultGasMultiplier    float32 = 1.0
	DefaultBufferPercent    int32   = 5
	DefaultGasPrice         int32   = 1
	DefaultWaitTime         int32   = 1
	DefaultGasLimit         float32 = 2
	DefaultGasLimitOverride uint64  = 30000000
	DefaultRPCTimeout       int64   = 5
	DefaultHTTPTimeout      int64   = 5
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

const (
	//BlockNumberInterval is the interval in seconds after which blockNumber needs to be calculated again by blockMonitor
	BlockNumberInterval = 5

	// StaleBlockNumberCheckInterval specifies the duration in seconds after which the BlockMonitor
	// switches to an alternate endpoint if the block number remains unchanged, indicating a potential stale endpoint.
	StaleBlockNumberCheckInterval = 15
)

//EndpointsContextTimeout defines the maximum duration in seconds to wait for establishing a connection for an endpoint
const EndpointsContextTimeout = 5

//APIKeyRegex will be used as a regular expression to be matched in job Urls
const APIKeyRegex = `\$\{(.+?)\}`

// Following are the constants which defines retry attempts and retry delay if there is an error in processing request
const (
	ProcessRequestRetryAttempts uint  = 2
	ProcessRequestRetryDelay    int64 = 2
)

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
