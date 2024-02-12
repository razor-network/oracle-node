package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"
)

func initAssetCache(client *ethclient.Client) error {
	log.Info("INITIALIZING JOB AND COLLECTION CACHE...")
	if err := utils.InitJobsCache(client); err != nil {
		log.Error("Error in initializing jobs cache: ", err)
		return err
	}
	if err := utils.InitCollectionsCache(client); err != nil {
		log.Error("Error in initializing collections cache: ", err)
		return err
	}

	latestHeader, err := clientUtils.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return err
	}
	log.Debugf("initAssetCache: Latest header value: %d", latestHeader.Number)

	fromBlock, err := razorUtils.EstimateBlockNumberAtEpochBeginning(client, latestHeader.Number)
	if err != nil {
		log.Error("Error in estimating block number at epoch beginning: ", err)
		return err
	}

	go startJobUpdateListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval))
	go startCollectionUpdateListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval))

	return nil
}

func startJobUpdateListener(client *ethclient.Client, initialFromBlockNumber *big.Int, interval time.Duration) {
	// Initialize fromBlock to a sensible starting point, such as the contract deployment block
	var fromBlock = initialFromBlockNumber

	// Create a ticker that fires at the specified interval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Retrieve the current block number to define the toBlock
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Printf("Failed to retrieve the latest block header: %v", err)
			continue
		}

		toBlock := header.Number

		// Listen for JobUpdated events between fromBlock and toBlock
		listenForJobUpdates(client, fromBlock, toBlock)

		// Update fromBlock for the next interval to start from the next block after toBlock
		fromBlock = new(big.Int).Add(toBlock, big.NewInt(1))
	}
}

func listenForJobUpdates(client *ethclient.Client, fromBlock, toBlock *big.Int) {
	// Set up the query for filtering logs
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(core.CollectionManagerAddress),
		},
	}

	// Retrieve the logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("Failed to fetch logs for JobUpdated event: %v", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	for _, vLog := range logs {
		// Check if the log is a JobUpdated event
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["JobUpdated"].ID.Hex() {
			// Unpack the event data
			topics := vLog.Topics
			// topics[1] gives job id in data type common.Hash
			jobId := utils.ConvertHashToUint16(topics[1])

			data, err := abiUtils.Unpack(contractAbi, "JobUpdated", vLog.Data)
			if err != nil {
				log.Error("Error in JobUpdated event logs: ", err)
				return
			}

			updatedJobData := bindings.StructsJob{
				Id:           jobId,
				SelectorType: data[0].(uint8),
				Weight:       data[2].(uint8),
				Power:        data[3].(int8),
				Selector:     data[5].(string),
				Url:          data[6].(string),
			}

			log.Debugf("Updating the job with Id %v with data %+v...", jobId, updatedJobData)
			cache.UpdateJobCache(jobId, updatedJobData)
		}
	}
}

func startCollectionUpdateListener(client *ethclient.Client, initialFromBlockNumber *big.Int, interval time.Duration) {
	// Initialize fromBlock to a sensible starting point, such as the contract deployment block
	var fromBlock = initialFromBlockNumber

	// Create a ticker that fires at the specified interval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Retrieve the current block number to define the toBlock
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Printf("Failed to retrieve the latest block header: %v", err)
			continue
		}

		toBlock := header.Number

		// Listen for CollectionUpdated events between fromBlock and toBlock
		listenForCollectionUpdates(client, fromBlock, toBlock)

		// Update fromBlock for the next interval to start from the next block after toBlock
		fromBlock = new(big.Int).Add(toBlock, big.NewInt(1))
	}
}

func listenForCollectionUpdates(client *ethclient.Client, fromBlock, toBlock *big.Int) {
	// Set up the query for filtering logs
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(core.CollectionManagerAddress),
		},
	}

	// Retrieve the logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("Failed to fetch logs for CollectionUpdated event: %v", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	for _, vLog := range logs {
		// Check if the log is a CollectionUpdated event
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["CollectionUpdated"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives collection id in data type common.Hash
			collectionId := utils.ConvertHashToUint16(topics[1])

			data, err := abiUtils.Unpack(contractAbi, "CollectionUpdated", vLog.Data)
			if err != nil {
				log.Error("Error in CollectionUpdated event logs: ", err)
				return
			}

			updatedCollectionData := bindings.StructsCollection{
				Id:                collectionId,
				Power:             data[0].(int8),
				AggregationMethod: data[2].(uint32),
				Tolerance:         data[3].(uint32),
				JobIDs:            data[4].([]uint16),
			}

			log.Debugf("Updating the collection with Id %v with data %+v...", collectionId, updatedCollectionData)
			cache.UpdateCollectionCache(collectionId, updatedCollectionData)
		} else if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["CollectionActivityStatus"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives collection id in data type common.Hash
			collectionId := utils.ConvertHashToUint16(topics[1])

			data, err := abiUtils.Unpack(contractAbi, "CollectionUpdated", vLog.Data)
			if err != nil {
				log.Error("Error in CollectionUpdated event logs: ", err)
				return
			}

			activeStatus := data[0].(bool)

			log.Debugf("Updating the activity status for collection with ID %v to %v...", collectionId, activeStatus)

			// Retrieve the existing collection data from the cache
			existingCollection, found := cache.GetCollectionFromCache(collectionId)
			if !found {
				log.Errorf("Collection with ID %v not found in cache", collectionId)
				continue
			}

			// Update the 'Active' status while keeping other fields unchanged
			existingCollection.Active = activeStatus

			// Update the collection in the cache
			cache.UpdateCollectionCache(collectionId, existingCollection)
		}
	}
}
