package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/rpc"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (*UtilsStruct) GetCollectionManagerWithOpts(client *ethclient.Client) (*bindings.CollectionManager, bind.CallOpts) {
	return UtilsInterface.GetCollectionManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetNumCollections(rpcParameters rpc.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetNumCollections")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetNumJobs(rpcParameters rpc.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetNumJobs")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetJobs(rpcParameters rpc.RPCParameters) ([]bindings.StructsJob, error) {
	var jobs []bindings.StructsJob
	numJobs, err := UtilsInterface.GetNumJobs(rpcParameters)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= int(numJobs); i++ {
		job, err := UtilsInterface.GetActiveJob(rpcParameters, uint16(i))
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (*UtilsStruct) GetNumActiveCollections(rpcParameters rpc.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetNumActiveCollections")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetAllCollections(rpcParameters rpc.RPCParameters) ([]bindings.StructsCollection, error) {
	var collections []bindings.StructsCollection
	numCollections, err := UtilsInterface.GetNumCollections(rpcParameters)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= int(numCollections); i++ {
		collection, err := UtilsInterface.GetCollection(rpcParameters, uint16(i))
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (*UtilsStruct) GetCollection(rpcParameters rpc.RPCParameters, collectionId uint16) (bindings.StructsCollection, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetCollection", collectionId)
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsCollection), nil
}

func (*UtilsStruct) GetActiveCollectionIds(rpcParameters rpc.RPCParameters) ([]uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetActiveCollections")
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().([]uint16), nil
}

func (*UtilsStruct) GetActiveStatus(rpcParameters rpc.RPCParameters, id uint16) (bool, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetActiveStatus", id)
	if err != nil {
		return false, err
	}
	return returnedValues[0].Interface().(bool), nil
}

func (*UtilsStruct) GetAggregatedDataOfCollection(rpcParameters rpc.RPCParameters, collectionId uint16, epoch uint32, commitParams *types.CommitParams) (*big.Int, error) {
	activeCollection, err := UtilsInterface.GetActiveCollection(commitParams.CollectionsCache, collectionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//Supply previous epoch to Aggregate in case if last reported value is required.
	collectionData, aggregationError := UtilsInterface.Aggregate(rpcParameters, epoch-1, activeCollection, commitParams)
	if aggregationError != nil {
		return nil, aggregationError
	}
	return collectionData, nil
}

func (*UtilsStruct) Aggregate(rpcParameters rpc.RPCParameters, previousEpoch uint32, collection bindings.StructsCollection, commitParams *types.CommitParams) (*big.Int, error) {
	var jobs []bindings.StructsJob
	var overriddenJobIds []uint16

	// Checks if assets.JSON file exists
	assetsFilePath, err := path.PathUtilsInterface.GetJobFilePath()
	if err != nil {
		return nil, err
	}
	if _, err := path.OSUtilsInterface.Stat(assetsFilePath); !errors.Is(err, os.ErrNotExist) {
		log.Debugf("assets.json file is present, checking jobs for collection Id: %v...", collection.Id)
		jsonFile, err := path.OSUtilsInterface.Open(assetsFilePath)
		if err != nil {
			return nil, err
		}
		defer jsonFile.Close()

		data, err := IOInterface.ReadAll(jsonFile)
		if err != nil {
			return nil, err
		}
		dataString := string(data)

		powerFromJSONFile := gjson.Get(dataString, "assets.collection."+collection.Name+".power").Int()
		if powerFromJSONFile != 0 {
			collection.Power = int8(powerFromJSONFile)
		}

		// Overriding the jobs from contracts with official jobs present in asset.go
		overrideJobs, overriddenJobIdsFromJSONfile := UtilsInterface.HandleOfficialJobsFromJSONFile(collection, dataString, commitParams)
		jobs = append(jobs, overrideJobs...)
		overriddenJobIds = append(overriddenJobIds, overriddenJobIdsFromJSONfile...)

		// Also adding custom jobs to jobs array
		customJobs := GetCustomJobsFromJSONFile(collection.Name, dataString)
		if len(customJobs) != 0 {
			log.Debugf("Got Custom Jobs from asset.json file for collectionId %v: %+v", collection.Id, customJobs)
		}
		jobs = append(jobs, customJobs...)
	}

	for _, id := range collection.JobIDs {
		// Ignoring the Jobs which are already overriden and added to jobs array
		if !Contains(overriddenJobIds, id) {
			job, isPresent := commitParams.JobsCache.GetJob(id)
			if !isPresent {
				log.Errorf("Job with id %v is not present in cache", id)
				continue
			}
			jobs = append(jobs, job)
		}
	}

	if len(jobs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	dataToCommit, weight := UtilsInterface.GetDataToCommitFromJobs(jobs, commitParams)
	if len(dataToCommit) == 0 {
		prevCommitmentData, err := UtilsInterface.FetchPreviousValue(rpcParameters, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return prevCommitmentData, nil
	}
	return performAggregation(dataToCommit, weight, collection.AggregationMethod)
}

func (*UtilsStruct) GetActiveJob(rpcParameters rpc.RPCParameters, jobId uint16) (bindings.StructsJob, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "Jobs", jobId)
	if err != nil {
		return bindings.StructsJob{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsJob), nil
}

func (*UtilsStruct) GetActiveCollection(collectionsCache *cache.CollectionsCache, collectionId uint16) (bindings.StructsCollection, error) {
	collection, isPresent := collectionsCache.GetCollection(collectionId)
	if !isPresent {
		return bindings.StructsCollection{}, errors.New("collection not present in cache")
	}
	if !collection.Active {
		return bindings.StructsCollection{}, errors.New("collection inactive")
	}
	return collection, nil
}

func (*UtilsStruct) GetDataToCommitFromJobs(jobs []bindings.StructsJob, commitParams *types.CommitParams) ([]*big.Int, []uint8) {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		data   []*big.Int
		weight []uint8
	)

	for _, job := range jobs {
		wg.Add(1)
		go processJobConcurrently(&wg, &mu, &data, &weight, job, commitParams)
	}

	wg.Wait()

	return data, weight
}

func processJobConcurrently(wg *sync.WaitGroup, mu *sync.Mutex, data *[]*big.Int, weight *[]uint8, job bindings.StructsJob, commitParams *types.CommitParams) {
	defer wg.Done()

	dataToAppend, err := UtilsInterface.GetDataToCommitFromJob(job, commitParams)
	if err != nil {
		return
	}
	log.Debugf("Job ID: %d, Job %s gives data %s", job.Id, job.Url, dataToAppend)

	mu.Lock()
	defer mu.Unlock()
	*data = append(*data, dataToAppend)
	*weight = append(*weight, job.Weight)
}

func (*UtilsStruct) GetDataToCommitFromJob(job bindings.StructsJob, commitParams *types.CommitParams) (*big.Int, error) {
	var (
		response            []byte
		apiErr              error
		dataSourceURLStruct types.DataSourceURL
	)
	log.Debugf("Job ID: %d, Getting the data to commit for job %s", job.Id, job.Name)
	if isJSONCompatible(job.Url) {
		log.Debugf("Job ID: %d, Job URL passed is a struct containing URL along with type of request data", job.Id)
		dataSourceURLInBytes := []byte(job.Url)

		err := json.Unmarshal(dataSourceURLInBytes, &dataSourceURLStruct)
		if err != nil {
			log.Errorf("Job ID: %d, Error in unmarshalling %s: %v", job.Id, job.Url, err)
			return nil, err
		}
		log.Infof("Job ID: %d, URL Struct: %+v", job.Id, dataSourceURLStruct)
	} else {
		log.Debugf("Job ID: %d, Job URL passed is a direct URL: %s", job.Id, job.Url)
		re := regexp.MustCompile(core.APIKeyRegex)
		isAPIKeyRequired := re.MatchString(job.Url)
		if isAPIKeyRequired {
			job.Url = ReplaceValueWithDataFromENVFile(re, job.Url)
		}
		dataSourceURLStruct = types.DataSourceURL{
			URL:    job.Url,
			Type:   "GET",
			Body:   nil,
			Header: nil,
		}
	}
	// Fetch data from API with retry mechanism
	var parsedData interface{}
	if job.SelectorType == 0 {
		start := time.Now()
		response, apiErr = GetDataFromAPI(commitParams, dataSourceURLStruct)
		if apiErr != nil {
			log.Errorf("Job ID: %d, Error in fetching data from API %s: %v", job.Id, job.Url, apiErr)
			return nil, apiErr
		}
		elapsed := time.Since(start).Seconds()
		log.Debugf("Job ID: %d, Time taken to fetch the data from API : %s was %f", job.Id, dataSourceURLStruct.URL, elapsed)

		var parsedJSON interface{}
		err := json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Errorf("Job ID: %d, Error in parsing data from API: %v", job.Id, err)
			return nil, err
		}
		parsedData, err = parseJSONData(parsedJSON, job.Selector)
		if err != nil {
			log.Errorf("Job ID: %d, Error in parsing JSON data: %v", job.Id, err)
			return nil, err
		}
	} else {
		//TODO: Add retry here.
		dataPoint, err := GetDataFromXHTML(dataSourceURLStruct, job.Selector)
		if err != nil {
			log.Errorf("Job ID: %d, Error in fetching value from parsed XHTML: %v", job.Id, err)
			return nil, err
		}
		// remove "," and currency symbols
		parsedData = regexp.MustCompile(`[\p{Sc}, ]`).ReplaceAllString(dataPoint, "")
	}

	datum, err := ConvertToNumber(parsedData, dataSourceURLStruct.ReturnType)
	if err != nil {
		log.Errorf("Job ID: %d, Result is not a number", job.Id)
		return nil, err
	}

	return MultiplyWithPower(datum, job.Power), err
}

func (*UtilsStruct) GetAssignedCollections(rpcParameters rpc.RPCParameters, numActiveCollections uint16, seed []byte) (map[int]bool, []*big.Int, error) {
	assignedCollections := make(map[int]bool)
	var seqAllottedCollections []*big.Int
	toAssign, err := UtilsInterface.ToAssign(rpcParameters)
	if err != nil {
		return nil, nil, err
	}
	for i := 0; i < int(toAssign); i++ {
		assigned := UtilsInterface.Prng(uint32(numActiveCollections), solsha3.SoliditySHA3([]string{"bytes32", "uint256"}, []interface{}{"0x" + hex.EncodeToString(seed), big.NewInt(int64(i))}))
		assignedCollections[int(assigned.Int64())] = true
		seqAllottedCollections = append(seqAllottedCollections, assigned)
	}
	return assignedCollections, seqAllottedCollections, nil
}

func (*UtilsStruct) GetLeafIdOfACollection(rpcParameters rpc.RPCParameters, collectionId uint16) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetLeafIdOfACollection", collectionId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetCollectionIdFromIndex(rpcParameters rpc.RPCParameters, medianIndex uint16) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetCollectionIdFromIndex", medianIndex)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetCollectionIdFromLeafId(rpcParameters rpc.RPCParameters, leafId uint16) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, AssetManagerInterface, "GetCollectionIdFromLeafId", leafId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func GetCustomJobsFromJSONFile(collection string, jsonFileData string) []bindings.StructsJob {
	var collectionCustomJobs []bindings.StructsJob
	var customJob types.CustomJob

	collectionCustomJobsPath := "assets.collection." + collection + ".custom jobs"
	customJobsJSONResult := gjson.Get(jsonFileData, collectionCustomJobsPath)
	if customJobsJSONResult.Exists() {
		customJobs := customJobsJSONResult.Array()
		if len(customJobs) == 0 {
			return nil
		}
		for i := 0; i < len(customJobs); i++ {
			customJobsData := customJobs[i].String()
			url := gjson.Get(customJobsData, "URL")
			if url.Exists() {
				customJob.URL = url.String()
			}
			name := gjson.Get(customJobsData, "name")
			if name.Exists() {
				customJob.Name = name.String()
			}
			selector := gjson.Get(customJobsData, "selector")
			if selector.Exists() {
				customJob.Selector = selector.String()
			}
			power := gjson.Get(customJobsData, "power")
			if power.Exists() {
				customJob.Power = int8(power.Int())
			}
			weight := gjson.Get(customJobsData, "weight")
			if weight.Exists() {
				customJob.Weight = uint8(weight.Int())
			}
			job := ConvertCustomJobToStructJob(customJob)
			collectionCustomJobs = append(collectionCustomJobs, job)
		}
	}

	return collectionCustomJobs
}

func ConvertCustomJobToStructJob(customJob types.CustomJob) bindings.StructsJob {
	return bindings.StructsJob{
		Url:      customJob.URL,
		Name:     customJob.Name,
		Selector: customJob.Selector,
		Power:    customJob.Power,
		Weight:   customJob.Weight,
	}
}

func (*UtilsStruct) HandleOfficialJobsFromJSONFile(collection bindings.StructsCollection, dataString string, commitParams *types.CommitParams) ([]bindings.StructsJob, []uint16) {
	var overrideJobs []bindings.StructsJob
	var overriddenJobIds []uint16

	collectionName := collection.Name
	jobIds := collection.JobIDs

	for i := 0; i < len(jobIds); i++ {
		officialJobsPath := "assets.collection." + collectionName + ".official jobs." + strconv.Itoa(int(jobIds[i]))
		officialJobsJSONResult := gjson.Get(dataString, officialJobsPath)
		if officialJobsJSONResult.Exists() {
			officialJobs := officialJobsJSONResult.String()
			if officialJobs != "" {
				job, isPresent := commitParams.JobsCache.GetJob(jobIds[i])
				if !isPresent {
					log.Errorf("Job with id %v is not present in cache", jobIds[i])
					continue
				}
				log.Debugf("Overriding job %s having jobId %d from official job present in assets.json file...", job.Url, job.Id)
				url := gjson.Get(officialJobs, "URL")
				if url.Exists() {
					job.Url = url.String()
				}
				selector := gjson.Get(officialJobs, "selector")
				if selector.Exists() {
					job.Selector = selector.String()
				}
				weight := gjson.Get(officialJobs, "weight")
				if weight.Exists() {
					job.Weight = uint8(weight.Int())
				}
				power := gjson.Get(officialJobs, "power")
				if power.Exists() {
					job.Power = int8(power.Int())
				}
				overrideJobs = append(overrideJobs, job)
				overriddenJobIds = append(overriddenJobIds, jobIds[i])
			}
		} else {
			continue
		}
	}

	return overrideJobs, overriddenJobIds
}

// InitJobsCache initializes the jobs cache with data fetched from the blockchain
func InitJobsCache(rpcParameters rpc.RPCParameters, jobsCache *cache.JobsCache) error {
	jobsCache.Mu.Lock()
	defer jobsCache.Mu.Unlock()

	// Flush the jobsCache before initialization
	for k := range jobsCache.Jobs {
		delete(jobsCache.Jobs, k)
	}

	numJobs, err := UtilsInterface.GetNumJobs(rpcParameters)
	if err != nil {
		return err
	}
	for i := 1; i <= int(numJobs); i++ {
		job, err := UtilsInterface.GetActiveJob(rpcParameters, uint16(i))
		if err != nil {
			return err
		}
		jobsCache.Jobs[job.Id] = job
	}
	return nil
}

// InitCollectionsCache initializes the collections cache with data fetched from the blockchain
func InitCollectionsCache(rpcParameters rpc.RPCParameters, collectionsCache *cache.CollectionsCache) error {
	collectionsCache.Mu.Lock()
	defer collectionsCache.Mu.Unlock()

	// Flush the collectionsCache before initialization
	for k := range collectionsCache.Collections {
		delete(collectionsCache.Collections, k)
	}

	numCollections, err := UtilsInterface.GetNumCollections(rpcParameters)
	if err != nil {
		return err
	}
	for i := 1; i <= int(numCollections); i++ {
		collection, err := UtilsInterface.GetCollection(rpcParameters, uint16(i))
		if err != nil {
			return err
		}
		collectionsCache.Collections[collection.Id] = collection
	}
	return nil
}

func ReplaceValueWithDataFromENVFile(re *regexp.Regexp, value string) string {
	// substrings denotes all the occurrences of substring which satisfies APIKeyRegex
	substrings := re.FindAllString(value, -1)
	log.Debug("ReplaceValueWithDataFromENVFile: Substrings array: ", substrings)

	// Replace each found substring with its corresponding value from environment variables
	for _, keyword := range substrings {
		if keyword != "" {
			log.Debug("ReplaceValueWithDataFromENVFile: Keyword to be looked for in env file: ", keyword)
			valueForKeyword := os.ExpandEnv(keyword)
			log.Debug("Replacing keyword with its value from env file...")
			value = strings.Replace(value, keyword, valueForKeyword, -1)
		}
	}
	return value
}

func isJSONCompatible(s string) bool {
	var temp interface{}
	err := json.Unmarshal([]byte(s), &temp)
	return err == nil
}
