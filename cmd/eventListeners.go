package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"
)

func (*UtilsStruct) InitAssetCache(client *ethclient.Client) (*cache.JobsCache, *cache.CollectionsCache, error) {
	log.Info("INITIALIZING JOBS AND COLLECTIONS CACHE...")

	// Create instances of cache
	jobsCache := cache.NewJobsCache()
	collectionsCache := cache.NewCollectionsCache()

	// Initialize caches
	if err := utils.InitJobsCache(client, jobsCache); err != nil {
		log.Error("Error in initializing jobs cache: ", err)
		return nil, nil, err
	}
	if err := utils.InitCollectionsCache(client, collectionsCache); err != nil {
		log.Error("Error in initializing collections cache: ", err)
		return nil, nil, err
	}

	go scheduleResetCache(client, jobsCache, collectionsCache)

	latestHeader, err := clientUtils.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return nil, nil, err
	}
	log.Debugf("initAssetCache: Latest header value: %d", latestHeader.Number)

	fromBlock, err := razorUtils.EstimateBlockNumberAtEpochBeginning(client, latestHeader.Number)
	if err != nil {
		log.Error("Error in estimating block number at epoch beginning: ", err)
		return nil, nil, err
	}

	// Start listeners for job and collection updates, passing the caches as arguments
	go startListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval), jobsCache, collectionsCache)

	return jobsCache, collectionsCache, nil
}

// startListener starts a generic listener for blockchain events that handles multiple event types.
func startListener(client *ethclient.Client, fromBlock *big.Int, interval time.Duration, jobsCache *cache.JobsCache, collectionsCache *cache.CollectionsCache) {
	// Will start listening for asset update events from confirm state
	_, err := cmdUtils.WaitForAppropriateState(client, "start event listener for asset update events", 4)
	if err != nil {
		log.Error("Error in waiting for appropriate state for starting event listener: ", err)
		return
	}

	// Sleeping till half of the confirm state to start listening for events in interval of duration core.AssetUpdateListenerInterval
	log.Debug("Will start listening for asset update events after half of the confirm state passes, sleeping till then...")
	time.Sleep(time.Second * time.Duration(core.StateLength/2))

	collectionManagerContractABI, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Error in parsing contract ABI: %v", err)
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	eventNames := []string{"JobUpdated", "CollectionUpdated", "CollectionActivityStatus", "JobCreated", "CollectionCreated"}

	log.Debugf("Starting to listen for asset update events from now in interval of every %v ...", interval)
	for range ticker.C {
		log.Debug("Checking for asset update events...")
		toBlock, err := clientUtils.GetLatestBlockWithRetry(client)
		if err != nil {
			log.Error("Error in getting latest block to start event listener: ", err)
			continue
		}

		processEvents(client, collectionManagerContractABI, fromBlock, toBlock.Number, eventNames, jobsCache, collectionsCache)

		// Update fromBlock for the next interval
		fromBlock = new(big.Int).Add(toBlock.Number, big.NewInt(1))
	}
}

// processEvents fetches and processes logs for multiple event types.
func processEvents(client *ethclient.Client, contractABI abi.ABI, fromBlock, toBlock *big.Int, eventNames []string, jobsCache *cache.JobsCache, collectionsCache *cache.CollectionsCache) {
	logs, err := getEventLogs(client, fromBlock, toBlock)
	if err != nil {
		log.Errorf("Failed to fetch logs: %v", err)
		return
	}

	for _, eventName := range eventNames {
		eventID := contractABI.Events[eventName].ID.Hex()
		for _, vLog := range logs {
			if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == eventID {
				switch eventName {
				case "JobUpdated", "JobCreated":
					jobId := utils.ConvertHashToUint16(vLog.Topics[1])
					updatedJob, err := utils.UtilsInterface.GetActiveJob(client, jobId)
					if err != nil {
						log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
						continue
					}
					log.Debugf("RECEIVED ASSET UPDATE: Updating the job with Id %v with details %+v...", jobId, updatedJob)
					jobsCache.UpdateJob(jobId, updatedJob)
				case "CollectionUpdated", "CollectionCreated", "CollectionActivityStatus":
					collectionId := utils.ConvertHashToUint16(vLog.Topics[1])
					newCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
					if err != nil {
						log.Errorf("Error in getting collection with collection Id %v: %v", collectionId, err)
						continue
					}
					log.Debugf("RECEIVED ASSET UPDATE: Updating the collection with ID %v with details %+v", collectionId, newCollection)
					collectionsCache.UpdateCollection(collectionId, newCollection)
				}
			}
		}
	}
}

// getEventLogs is a utility function to fetch the event logs
func getEventLogs(client *ethclient.Client, fromBlock *big.Int, toBlock *big.Int) ([]Types.Log, error) {
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
		log.Errorf("Error in filter logs: %v", err)
		return []Types.Log{}, nil
	}

	return logs, nil
}

func scheduleResetCache(client *ethclient.Client, jobsCache *cache.JobsCache, collectionsCache *cache.CollectionsCache) {
	// Will not allow to start scheduling reset cache in confirm and commit state
	// As in confirm state, updating jobs/collections cache takes place
	// As in commit state, fetching of jobs/collection from cache takes place
	_, err := cmdUtils.WaitForAppropriateState(client, "schedule resetting cache", 1, 2, 3)
	if err != nil {
		log.Error("Error in waiting for appropriate state to schedule reset cache: ", err)
		return
	}

	log.Debugf("Scheduling reset asset cache now in interval of every %v ...", core.AssetCacheExpiry)
	assetCacheTicker := time.NewTicker(time.Second * time.Duration(core.AssetCacheExpiry))
	defer assetCacheTicker.Stop()

	for range assetCacheTicker.C {
		log.Info("ASSET CACHE EXPIRED! INITIALIZING JOBS AND COLLECTIONS CACHE AGAIN...")
		if err := razorUtils.ResetAssetCache(client, jobsCache, collectionsCache); err != nil {
			log.Errorf("Error resetting asset cache: %v", err)
		}
	}
}
