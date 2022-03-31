package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/pflag"
)

func (*UtilsStruct) ConnectToClient(provider string) *ethclient.Client {
	client, err := EthClient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func (*UtilsStruct) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	coinContract := UtilsInterface.GetTokenManager(client)
	opts := UtilsInterface.GetOptions()
	return CoinInterface.BalanceOf(coinContract, &opts, address)
}

func (*UtilsStruct) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	block, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return -1, err
	}
	blockTime := uint64(block.Time)
	lowerLimit := (core.StateLength * uint64(buffer)) / 100
	upperLimit := core.StateLength - (core.StateLength*uint64(buffer))/100
	if blockTime%(core.StateLength) > upperLimit+core.StateBuffer || blockTime%(core.StateLength) < lowerLimit+core.StateBuffer {
		return -1, nil
	}
	state := blockTime / core.StateLength
	return int64(state) % core.NumberOfStates, nil
}

func (*UtilsStruct) CheckTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := ClientInterface.TransactionReceipt(client, context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func (*UtilsStruct) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	timeout := core.BlockCompletionTimeout
	for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
		log.Debug("Checking if transaction is mined....")
		transactionStatus := UtilsInterface.CheckTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			log.Error("Transaction mining unsuccessful")
			return 0
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully")
			return 1
		}
		Time.Sleep(3 * time.Second)
	}
	log.Info("Timeout Passed")
	return 0
}

func (*UtilsStruct) WaitTillNextNSecs(waitTime int32) {
	if waitTime <= 0 {
		waitTime = 1
	}
	Time.Sleep(time.Duration(waitTime) * time.Second)
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}

func (*UtilsStruct) IsFlagPassed(name string) bool {
	found := false
	for _, arg := range os.Args {
		if arg == "--"+name {
			found = true
		}
	}
	return found
}

func (*UtilsStruct) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	ethBalance, err := ClientInterface.BalanceAt(client, context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Error in fetching eth balance of the account: %s\n%s", address, err)
	}
	if ethBalance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Eth balance is 0, Aborting...")
	}
}

func GetStateName(stateNumber int64) string {
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
		stateName = "-1"
	}
	return stateName
}

func (*UtilsStruct) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	if UtilsInterface.IsFlagPassed("stakerId") {
		return UtilsInterface.GetUint32(flagSet, "stakerId")
	}
	return UtilsInterface.GetStakerId(client, address)
}

func (*UtilsStruct) GetEpoch(client *ethclient.Client) (uint32, error) {
	latestHeader, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}
	epoch := uint64(latestHeader.Time) / uint64(core.EpochLength)
	return uint32(epoch), nil
}

func (*UtilsStruct) SaveDataToFile(fileName string, epoch uint32, data []*big.Int) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	f, err := OS.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	fmt.Fprintln(f, epoch)
	for _, datum := range data {
		_, err := fmt.Fprintln(f, datum.String())
		if err != nil {
			return err
		}
	}
	defer f.Close()
	return nil
}

func (*UtilsStruct) ReadDataFromFile(fileName string) (uint32, []*big.Int, error) {
	var (
		data  []*big.Int
		epoch uint32
	)
	file, err := OS.Open(fileName)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := Bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		if lineCount > 0 {
			dataToAppend, ok := big.NewInt(0).SetString(scanner.Text(), 10)
			if ok {
				data = append(data, dataToAppend)
			}
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, nil, err
			}
			epoch = uint32(value)
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}
	return epoch, data, nil
}

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

func (*UtilsStruct) GetRemainingTimeOfCurrentState(client *ethclient.Client, bufferPercent int32) (int64, error) {
	block, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return 0, err
	}
	timeRemaining := core.StateLength - (block.Time % core.StateLength)
	upperLimit := ((core.StateLength * uint64(bufferPercent)) / 100) + core.StateBuffer

	return int64(timeRemaining - upperLimit), nil
}

func (*UtilsStruct) CalculateSalt(epoch uint32, medians []uint32) [32]byte {
	salt := solsha3.SoliditySHA3([]string{"uint32", "uint32[]"}, []interface{}{epoch, medians})
	var saltInBytes32 [32]byte
	copy(saltInBytes32[:], salt)
	return saltInBytes32
}

func (*UtilsStruct) Prng(max uint32, prngHashes []byte) *big.Int {
	sum := big.NewInt(0).SetBytes(prngHashes)
	maxBigInt := big.NewInt(int64(max))
	return sum.Mod(sum, maxBigInt)
}

func CalculateBlockNumberAtEpochBeginning(client *ethclient.Client, epochLength int64, currentBlockNumber *big.Int) (*big.Int, error) {
	block, err := ClientInterface.HeaderByNumber(client, context.Background(), currentBlockNumber)
	if err != nil {
		log.Errorf("Error in fetching block : %s", err)
		return nil, err
	}
	current_epoch := block.Time / uint64(core.EpochLength)
	previousBlockNumber := block.Number.Uint64() - core.StateLength

	previousBlock, err := ClientInterface.HeaderByNumber(client, context.Background(), big.NewInt(int64(previousBlockNumber)))
	if err != nil {
		log.Errorf("Err in fetching Previous block : %s", err)
		return nil, err
	}
	previousBlockActualTimestamp := previousBlock.Time
	previousBlockAssumedTimestamp := block.Time - uint64(core.EpochLength)
	previous_epoch := previousBlockActualTimestamp / uint64(core.EpochLength)
	if previousBlockActualTimestamp > previousBlockAssumedTimestamp && previous_epoch != current_epoch-1 {
		return CalculateBlockNumberAtEpochBeginning(client, core.EpochLength, big.NewInt(int64(previousBlockNumber)))

	}
	return big.NewInt(int64(previousBlockNumber)), nil

}

func (*UtilsStruct) SaveDataToCommitJsonFile(filePath string, epoch uint32, commitData types.CommitData) error {

	var data types.CommitFileData
	data.Epoch = epoch
	data.AssignedCollections = commitData.AssignedCollections
	data.SeqAllottedCollections = commitData.SeqAllottedCollections
	data.Leaves = commitData.Leaves

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func (*UtilsStruct) ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.CommitFileData{}, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.CommitFileData{}, err
	}
	var commitedData types.CommitFileData

	err = json.Unmarshal(byteValue, &commitedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.CommitFileData{}, err
	}
	return commitedData, nil
}

func (*UtilsStruct) AssignLogFile(flagSet *pflag.FlagSet) {
	if UtilsInterface.IsFlagPassed("logFile") {
		fileName, err := FlagSetInterface.GetLogFileName(flagSet)
		if err != nil {
			log.Fatalf("Error in getting file name : ", err)
		}
		logger.InitializeLogger(fileName)
	}
}

func (*UtilsStruct) SaveDataToProposeJsonFile(filePath string, epoch uint32, proposeData types.ProposeData) error {

	var data types.ProposeFileData
	data.Epoch = epoch
	data.MediansData = proposeData.MediansData
	data.RevealedCollectionIds = proposeData.RevealedCollectionIds
	data.RevealedDataMaps = proposeData.RevealedDataMaps

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func (*UtilsStruct) ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.ProposeFileData{}, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.ProposeFileData{}, err
	}
	var proposedData types.ProposeFileData

	err = json.Unmarshal(byteValue, &proposedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.ProposeFileData{}, err
	}
	return proposedData, nil
}

func SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error {
	var data types.DisputeFileData

	data.BountyIdQueue = bountyIdQueue
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func ReadFromJsonFile(filePath string) (types.DisputeFileData, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.DisputeFileData{}, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.DisputeFileData{}, err
	}
	var disputeData types.DisputeFileData

	err = json.Unmarshal(byteValue, &disputeData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.DisputeFileData{}, err
	}
	return disputeData, nil
}
