// This file contains constants, DO NOT MODIFY.
package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength int64 = 1200
var NumberOfStates int64 = 5
var ChainId = big.NewInt(0x785B4B9847B9)
var StateLength = uint64(EpochLength / NumberOfStates)
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

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
