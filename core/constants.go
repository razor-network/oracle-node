// This file contains constants, DO NOT MODIFY.
package core

import (
	"github.com/ethereum/go-ethereum/common"
)

var EpochLength uint64 = 1200
var NumberOfStates uint64 = 5
var StateLength = EpochLength / NumberOfStates
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30
var BatchSize = 1000
var NumRoutines = 10
var MaxIterations = 10000000

//DisputeGasMultiplier is used to increase the gas required from the gasLimit for the following function
var DisputeGasMultiplier float32 = 5.5
var DataFileDirectory string = "data_files"
var CommitDataFile string = "_commitData.json"
var ProposeDataFile string = "_proposeData.json"
var DisputeDataFile string = "_disputeData.json"
var AssetsDataFile string = "assets.json"
var ConfigFile string = "razor.yaml"
var LogFile string = "logs"
var DefaultPathName string = ".razor"
