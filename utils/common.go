//Package utils provides the utils functions
package utils

import (
	"context"
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"time"

	"github.com/avast/retry-go"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/pflag"
)

// ConnectToClient function helps in connecting with client
func (*UtilsStruct) ConnectToClient(provider string) *ethclient.Client {
	client, err := EthClient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

// FetchBalance function helps in fetching the balance
func (*UtilsStruct) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	var (
		balance *big.Int
		err     error
	)
	err = retry.Do(
		func() error {
			address := common.HexToAddress(accountAddress)
			erc20Contract := UtilsInterface.GetTokenManager(client)
			opts := UtilsInterface.GetOptions()
			balance, err = CoinInterface.BalanceOf(erc20Contract, &opts, address)
			if err != nil {
				log.Error("Error in fetching balance....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}

// GetBufferedState function returns the buffered state
func (*UtilsStruct) GetBufferedState(client *ethclient.Client, buffer int32) (int64, error) {
	block, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return -1, err
	}
	stateBuffer, err := UtilsInterface.GetStateBuffer(client)
	if err != nil {
		return -1, err
	}
	lowerLimit := (core.StateLength * uint64(buffer)) / 100
	upperLimit := core.StateLength - (core.StateLength*uint64(buffer))/100
	if block.Time%(core.StateLength) > upperLimit-stateBuffer || block.Time%(core.StateLength) < lowerLimit+stateBuffer {
		return -1, nil
	}
	state := block.Time / core.StateLength
	return int64(state % core.NumberOfStates), nil
}

// CheckTransactionReceipt function checks the transaction receipt
func (*UtilsStruct) CheckTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := ClientInterface.TransactionReceipt(client, context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

// WaitForBlockCompletion function waits for the block completion
func (*UtilsStruct) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error {
	timeout := core.BlockCompletionTimeout
	for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
		log.Debug("Checking if transaction is mined....")
		transactionStatus := UtilsInterface.CheckTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			err := errors.New("transaction mining unsuccessful")
			log.Error(err)
			return err
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully")
			return nil
		}
		Time.Sleep(time.Second)
	}
	log.Info("Timeout Passed")
	return errors.New("timeout passed for transaction mining")
}

// WaitTillNextNSecs function wait for next N seconds
func (*UtilsStruct) WaitTillNextNSecs(waitTime int32) {
	if waitTime <= 0 {
		waitTime = 1
	}
	Time.Sleep(time.Duration(waitTime) * time.Second)
}

// CheckError function checks the error
func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}

// IsValidERC20Address function checks if the ERC20 address is valid or not
func IsValidERC20Address(address string) bool {
	if !common.IsHexAddress(address) {
		log.Error("Invalid ERC20 Address")
		return false
	}
	return true
}

// IsFlagPassed function checks if the flag is passed or not
func (*UtilsStruct) IsFlagPassed(name string) bool {
	found := false
	for _, arg := range os.Args {
		if arg == "--"+name {
			found = true
		}
	}
	return found
}

// CheckEthBalanceIsZero function checks if the eth balance is zero or not
func (*UtilsStruct) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	ethBalance, err := ClientInterface.BalanceAt(client, context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Error in fetching sfuel balance of the account: %s\n%s", address, err)
	}
	if ethBalance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Sfuel balance is 0, Aborting...")
	}
}

// GetStateName function returns the state name
func (*UtilsStruct) GetStateName(stateNumber int64) string {
	var stateName string
	switch stateNumber {
	case 0:
		stateName = "Commit"
	case 1:
		stateName = "Reveal"
	case 2:
		stateName = "Propose"
	case 3:
		stateName = "Dispute"
	case 4:
		stateName = "Confirm"
	default:
		stateName = "Buffer"
	}
	return stateName
}

// AssignStakerId function assigns the staker Id
func (*UtilsStruct) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	if UtilsInterface.IsFlagPassed("stakerId") {
		return UtilsInterface.GetUint32(flagSet, "stakerId")
	}
	return UtilsInterface.GetStakerId(client, address)
}

// GetEpoch function returns the epoch
func (*UtilsStruct) GetEpoch(client *ethclient.Client) (uint32, error) {
	latestHeader, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}
	epoch := latestHeader.Time / uint64(core.EpochLength)
	return uint32(epoch), nil
}

// CalculateBlockTime function calculates the block time
func (*UtilsStruct) CalculateBlockTime(client *ethclient.Client) int64 {
	latestBlock, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Fatalf("Error in fetching latest Block: %s", err)
	}
	latestBlockNumber := latestBlock.Number
	lastSecondBlock, err := ClientInterface.HeaderByNumber(client, context.Background(), big.NewInt(1).Sub(latestBlockNumber, big.NewInt(1)))
	if err != nil {
		log.Fatalf("Error in fetching last second Block: %s", err)
	}
	return int64(latestBlock.Time - lastSecondBlock.Time)
}

// GetRemainingTimeOfCurrentState function returns the remaining time of current state
func (*UtilsStruct) GetRemainingTimeOfCurrentState(client *ethclient.Client, bufferPercent int32) (int64, error) {
	block, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return 0, err
	}
	stateBuffer, err := UtilsInterface.GetStateBuffer(client)
	if err != nil {
		return 0, err
	}
	timeRemaining := core.StateLength - (block.Time % core.StateLength)
	upperLimit := ((core.StateLength * uint64(bufferPercent)) / 100) + stateBuffer

	return int64(timeRemaining - upperLimit), nil
}

// CalculateSalt function calculates the salt
func (*UtilsStruct) CalculateSalt(epoch uint32, medians []*big.Int) [32]byte {
	salt := solsha3.SoliditySHA3([]string{"uint32", "uint256"}, []interface{}{epoch, medians})
	var saltInBytes32 [32]byte
	copy(saltInBytes32[:], salt)
	return saltInBytes32
}

// Prng function returns the prng
func (*UtilsStruct) Prng(max uint32, prngHashes []byte) *big.Int {
	sum := big.NewInt(0).SetBytes(prngHashes)
	maxBigInt := big.NewInt(int64(max))
	return sum.Mod(sum, maxBigInt)
}

// EstimateBlockNumberAtEpochBeginning function calculates the block number at epoch beginning
func (*UtilsStruct) EstimateBlockNumberAtEpochBeginning(client *ethclient.Client, currentBlockNumber *big.Int) (*big.Int, error) {
	block, err := ClientInterface.HeaderByNumber(client, context.Background(), currentBlockNumber)
	if err != nil {
		log.Errorf("Error in fetching block : %s", err)
		return nil, err
	}
	currentEpoch := block.Time / uint64(core.EpochLength)
	previousBlockNumber := block.Number.Uint64() - core.StateLength

	previousBlock, err := ClientInterface.HeaderByNumber(client, context.Background(), big.NewInt(int64(previousBlockNumber)))
	if err != nil {
		log.Errorf("Err in fetching Previous block : %s", err)
		return nil, err
	}
	previousBlockActualTimestamp := previousBlock.Time
	previousBlockAssumedTimestamp := block.Time - uint64(core.EpochLength)
	previousEpoch := previousBlockActualTimestamp / uint64(core.EpochLength)
	if previousBlockActualTimestamp > previousBlockAssumedTimestamp && previousEpoch != currentEpoch-1 {
		return UtilsInterface.EstimateBlockNumberAtEpochBeginning(client, big.NewInt(int64(previousBlockNumber)))
	}
	return big.NewInt(int64(previousBlockNumber)), nil

}

// SaveDataToCommitJsonFile function saves data to commit JSON file
func (*UtilsStruct) SaveDataToCommitJsonFile(filePath string, epoch uint32, commitData types.CommitData) error {

	var data types.CommitFileData
	data.Epoch = epoch
	data.AssignedCollections = commitData.AssignedCollections
	data.SeqAllottedCollections = commitData.SeqAllottedCollections
	data.Leaves = commitData.Leaves

	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

// ReadFromCommitJsonFile function reads from commit JSON file
func (*UtilsStruct) ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.CommitFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.CommitFileData{}, err
	}
	var commitedData types.CommitFileData

	err = JsonInterface.Unmarshal(byteValue, &commitedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.CommitFileData{}, err
	}
	return commitedData, nil
}

// AssignLogFile function assigns the log file
func (*UtilsStruct) AssignLogFile(flagSet *pflag.FlagSet) {
	if UtilsInterface.IsFlagPassed("logFile") {
		fileName, err := FlagSetInterface.GetLogFileName(flagSet)
		if err != nil {
			log.Fatalf("Error in getting file name : ", err)
		}
		logger.InitializeLogger(fileName)
	}
}

// SaveDataToProposeJsonFile function saves data to propose JSON file
func (*UtilsStruct) SaveDataToProposeJsonFile(filePath string, epoch uint32, proposeData types.ProposeData) error {

	var data types.ProposeFileData
	data.Epoch = epoch
	data.MediansData = proposeData.MediansData
	data.RevealedCollectionIds = proposeData.RevealedCollectionIds
	data.RevealedDataMaps = proposeData.RevealedDataMaps

	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

// ReadFromProposeJsonFile function reads from propose JSON file
func (*UtilsStruct) ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.ProposeFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.ProposeFileData{}, err
	}
	var proposedData types.ProposeFileData

	err = JsonInterface.Unmarshal(byteValue, &proposedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.ProposeFileData{}, err
	}
	return proposedData, nil
}

// SaveDataToDisputeJsonFile function saves data to Dispute JSON file
func (*UtilsStruct) SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error {
	var data types.DisputeFileData

	data.BountyIdQueue = bountyIdQueue
	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

// ReadFromDisputeJsonFile function reads from dispute JSON file
func (*UtilsStruct) ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.DisputeFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.DisputeFileData{}, err
	}
	var disputeData types.DisputeFileData

	err = JsonInterface.Unmarshal(byteValue, &disputeData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.DisputeFileData{}, err
	}
	return disputeData, nil
}
