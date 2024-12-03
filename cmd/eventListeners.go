package cmd

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"strings"
)

func (*UtilsStruct) InitJobAndCollectionCache(rpcParameters rpc.RPCParameters) (*cache.JobsCache, *cache.CollectionsCache, *big.Int, error) {
	initAssetCacheBlock, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return nil, nil, nil, err
	}
	log.Debugf("InitJobAndCollectionCache: Latest header value when initializing jobs and collections cache: %d", initAssetCacheBlock.Number)

	log.Info("INITIALIZING JOBS AND COLLECTIONS CACHE...")

	// Create instances of cache
	jobsCache := cache.NewJobsCache()
	collectionsCache := cache.NewCollectionsCache()

	// Initialize caches
	if err := utils.InitJobsCache(rpcParameters, jobsCache); err != nil {
		log.Error("Error in initializing jobs cache: ", err)
		return nil, nil, nil, err
	}
	if err := utils.InitCollectionsCache(rpcParameters, collectionsCache); err != nil {
		log.Error("Error in initializing collections cache: ", err)
		return nil, nil, nil, err
	}

	return jobsCache, collectionsCache, initAssetCacheBlock.Number, nil
}

// CheckForJobAndCollectionEvents checks for specific job and collections event that were emitted.
func CheckForJobAndCollectionEvents(rpcParameters rpc.RPCParameters, commitParams *types.CommitParams) error {
	collectionManagerContractABI, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Error in parsing collection manager contract ABI: %v", err)
		return err
	}

	eventNames := []string{core.JobUpdatedEvent, core.CollectionUpdatedEvent, core.CollectionActivityStatusEvent, core.JobCreatedEvent, core.CollectionCreatedEvent}

	log.Debug("Checking for Job/Collection update events...")
	toBlock, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
	if err != nil {
		log.Error("Error in getting latest block to start event listener: ", err)
		return err
	}

	// Process events and update the fromBlock for the next iteration
	newFromBlock, err := processEvents(rpcParameters, collectionManagerContractABI, commitParams.FromBlockToCheckForEvents, toBlock.Number, eventNames, commitParams.JobsCache, commitParams.CollectionsCache)
	if err != nil {
		return err
	}

	// Update the commitParams with the new fromBlock
	commitParams.FromBlockToCheckForEvents = new(big.Int).Add(newFromBlock, big.NewInt(1))

	return nil
}

// processEvents fetches and processes logs for multiple event types.
func processEvents(rpcParameters rpc.RPCParameters, contractABI abi.ABI, fromBlock, toBlock *big.Int, eventNames []string, jobsCache *cache.JobsCache, collectionsCache *cache.CollectionsCache) (*big.Int, error) {
	logs, err := getEventLogs(rpcParameters, fromBlock, toBlock)
	if err != nil {
		log.Errorf("Failed to fetch logs: %v", err)
		return nil, err
	}

	for _, eventName := range eventNames {
		eventID := contractABI.Events[eventName].ID.Hex()
		for _, vLog := range logs {
			if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == eventID {
				switch eventName {
				case core.JobUpdatedEvent, core.JobCreatedEvent:
					jobId := utils.ConvertHashToUint16(vLog.Topics[1])
					updatedJob, err := razorUtils.GetActiveJob(rpcParameters, jobId)
					if err != nil {
						log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
						continue
					}
					log.Debugf("RECEIVED JOB EVENT: Updating the job with Id %v with details %+v...", jobId, updatedJob)
					jobsCache.UpdateJob(jobId, updatedJob)
				case core.CollectionUpdatedEvent, core.CollectionCreatedEvent, core.CollectionActivityStatusEvent:
					collectionId := utils.ConvertHashToUint16(vLog.Topics[1])
					newCollection, err := razorUtils.GetCollection(rpcParameters, collectionId)
					if err != nil {
						log.Errorf("Error in getting collection with collection Id %v: %v", collectionId, err)
						continue
					}
					log.Debugf("RECEIVED COLLECTION EVENT: Updating the collection with ID %v with details %+v", collectionId, newCollection)
					collectionsCache.UpdateCollection(collectionId, newCollection)
				}
			}
		}
	}

	// Return the new toBlock for the next iteration
	return toBlock, nil
}

// getEventLogs is a utility function to fetch the event logs
func getEventLogs(rpcParameters rpc.RPCParameters, fromBlock *big.Int, toBlock *big.Int) ([]Types.Log, error) {
	log.Debugf("Checking for events from block %v to block %v...", fromBlock, toBlock)

	// Set up the query for filtering logs
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(core.CollectionManagerAddress),
		},
	}

	// Retrieve the logs
	logs, err := clientUtils.FilterLogsWithRetry(rpcParameters, query)
	if err != nil {
		log.Errorf("Error in filter logs: %v", err)
		return []Types.Log{}, err
	}

	return logs, nil
}
