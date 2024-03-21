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
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (*UtilsStruct) GetCollectionManagerWithOpts(client *ethclient.Client) (*bindings.CollectionManager, bind.CallOpts) {
	return UtilsInterface.GetCollectionManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetNumCollections(client *ethclient.Client) (uint16, error) {
	var (
		numCollections uint16
		err            error
	)
	err = retry.Do(
		func() error {
			numCollections, err = AssetManagerInterface.GetNumCollections(client)
			if err != nil {
				log.Error("Error in fetching numCollections.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return numCollections, nil
}

func (*UtilsStruct) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	var jobs []bindings.StructsJob
	numJobs, err := AssetManagerInterface.GetNumJobs(client)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= int(numJobs); i++ {
		job, err := UtilsInterface.GetActiveJob(client, uint16(i))
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (*UtilsStruct) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	var (
		numActiveAssets uint16
		err             error
	)
	err = retry.Do(
		func() error {
			numActiveAssets, err = AssetManagerInterface.GetNumActiveCollections(client)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return numActiveAssets, nil
}

func (*UtilsStruct) GetAllCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	var collections []bindings.StructsCollection
	numCollections, err := UtilsInterface.GetNumCollections(client)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= int(numCollections); i++ {
		collection, err := AssetManagerInterface.GetCollection(client, uint16(i))
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (*UtilsStruct) GetCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error) {
	var (
		collection bindings.StructsCollection
		err        error
	)
	err = retry.Do(
		func() error {
			collection, err = AssetManagerInterface.GetCollection(client, collectionId)
			if err != nil {
				log.Error("Error in fetching collection.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	return collection, nil
}

func (*UtilsStruct) GetActiveCollectionIds(client *ethclient.Client) ([]uint16, error) {
	var (
		activeCollectionIds []uint16
		err                 error
	)
	err = retry.Do(
		func() error {
			activeCollectionIds, err = AssetManagerInterface.GetActiveCollections(client)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return activeCollectionIds, nil
}

func (*UtilsStruct) GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32, localCache *cache.LocalCache) (*big.Int, error) {
	activeCollection, err := UtilsInterface.GetActiveCollection(client, collectionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//Supply previous epoch to Aggregate in case if last reported value is required.
	collectionData, aggregationError := UtilsInterface.Aggregate(client, epoch-1, activeCollection, localCache)
	if aggregationError != nil {
		return nil, aggregationError
	}
	return collectionData, nil
}

func (*UtilsStruct) Aggregate(client *ethclient.Client, previousEpoch uint32, collection bindings.StructsCollection, localCache *cache.LocalCache) (*big.Int, error) {
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
		overrideJobs, overriddenJobIdsFromJSONfile := UtilsInterface.HandleOfficialJobsFromJSONFile(client, collection, dataString)
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
			job, isPresent := cache.GetJobFromCache(id)
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
	dataToCommit, weight := UtilsInterface.GetDataToCommitFromJobs(jobs, localCache)
	if len(dataToCommit) == 0 {
		prevCommitmentData, err := UtilsInterface.FetchPreviousValue(client, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return prevCommitmentData, nil
	}
	return performAggregation(dataToCommit, weight, collection.AggregationMethod)
}

func (*UtilsStruct) GetActiveJob(client *ethclient.Client, jobId uint16) (bindings.StructsJob, error) {
	var (
		job bindings.StructsJob
		err error
	)
	err = retry.Do(
		func() error {
			job, err = AssetManagerInterface.Jobs(client, jobId)
			if err != nil {
				log.Errorf("Error in fetching job %d.... Retrying", jobId)
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsJob{}, err
	}
	return job, nil
}

func (*UtilsStruct) GetActiveCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error) {
	collection, isPresent := cache.GetCollectionFromCache(collectionId)
	if !isPresent {
		return bindings.StructsCollection{}, errors.New("collection not present in cache")
	}
	if !collection.Active {
		return bindings.StructsCollection{}, errors.New("collection inactive")
	}
	return collection, nil
}

func (*UtilsStruct) GetDataToCommitFromJobs(jobs []bindings.StructsJob, localCache *cache.LocalCache) ([]*big.Int, []uint8) {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		data   []*big.Int
		weight []uint8
	)

	for _, job := range jobs {
		wg.Add(1)
		go processJobConcurrently(&wg, &mu, &data, &weight, job, localCache)
	}

	wg.Wait()

	return data, weight
}

func processJobConcurrently(wg *sync.WaitGroup, mu *sync.Mutex, data *[]*big.Int, weight *[]uint8, job bindings.StructsJob, localCache *cache.LocalCache) {
	defer wg.Done()

	dataToAppend, err := UtilsInterface.GetDataToCommitFromJob(job, localCache)
	if err != nil {
		return
	}
	log.Debugf("Job ID: %d, Job %s gives data %s", job.Id, job.Url, dataToAppend)

	mu.Lock()
	defer mu.Unlock()
	*data = append(*data, dataToAppend)
	*weight = append(*weight, job.Weight)
}

func (*UtilsStruct) GetDataToCommitFromJob(job bindings.StructsJob, localCache *cache.LocalCache) (*big.Int, error) {
	var parsedJSON map[string]interface{}
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
		response, apiErr = GetDataFromAPI(dataSourceURLStruct, localCache)
		if apiErr != nil {
			log.Errorf("Job ID: %d, Error in fetching data from API %s: %v", job.Id, job.Url, apiErr)
			return nil, apiErr
		}
		elapsed := time.Since(start).Seconds()
		log.Debugf("Job ID: %d, Time taken to fetch the data from API : %s was %f", job.Id, dataSourceURLStruct.URL, elapsed)

		err := json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Errorf("Job ID: %d, Error in parsing data from API: %v", job.Id, err)
			return nil, err
		}
		parsedData, err = GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Errorf("Job ID: %d, Error in fetching value from parsed data: %v", job.Id, err)
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

func (*UtilsStruct) GetAssignedCollections(client *ethclient.Client, numActiveCollections uint16, seed []byte) (map[int]bool, []*big.Int, error) {
	assignedCollections := make(map[int]bool)
	var seqAllottedCollections []*big.Int
	toAssign, err := UtilsInterface.ToAssign(client)
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

func (*UtilsStruct) GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error) {
	var (
		leafId uint16
		err    error
	)
	err = retry.Do(
		func() error {
			leafId, err = AssetManagerInterface.GetLeafIdOfACollection(client, collectionId)
			if err != nil {
				log.Error("Error in fetching collection id.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return leafId, nil
}

func (*UtilsStruct) GetCollectionIdFromIndex(client *ethclient.Client, medianIndex uint16) (uint16, error) {
	var (
		collectionId uint16
		err          error
	)
	err = retry.Do(
		func() error {
			collectionId, err = AssetManagerInterface.GetCollectionIdFromIndex(client, medianIndex)
			if err != nil {
				log.Error("Error in fetching collection id.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return collectionId, nil
}

func (*UtilsStruct) GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error) {
	var (
		collectionId uint16
		err          error
	)
	err = retry.Do(
		func() error {
			collectionId, err = AssetManagerInterface.GetCollectionIdFromLeafId(client, leafId)
			if err != nil {
				log.Error("Error in fetching collection id.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return collectionId, nil
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

func (*UtilsStruct) HandleOfficialJobsFromJSONFile(client *ethclient.Client, collection bindings.StructsCollection, dataString string) ([]bindings.StructsJob, []uint16) {
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
				job, err := UtilsInterface.GetActiveJob(client, jobIds[i])
				if err != nil {
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

func InitJobsCache(client *ethclient.Client) error {
	cache.JobsCache.Mu.Lock()
	defer cache.JobsCache.Mu.Unlock()

	// Flush the JobsCache before initialization
	for k := range cache.JobsCache.Jobs {
		delete(cache.JobsCache.Jobs, k)
	}

	numJobs, err := AssetManagerInterface.GetNumJobs(client)
	if err != nil {
		return err
	}
	for i := 1; i <= int(numJobs); i++ {
		job, err := UtilsInterface.GetActiveJob(client, uint16(i))
		if err != nil {
			return err
		}
		cache.JobsCache.Jobs[job.Id] = job
	}
	return nil
}

func InitCollectionsCache(client *ethclient.Client) error {
	cache.CollectionsCache.Mu.Lock()
	defer cache.CollectionsCache.Mu.Unlock()

	// Flush the CollectionsCacheStruct before initialization
	for k := range cache.CollectionsCache.Collections {
		delete(cache.CollectionsCache.Collections, k)
	}

	numCollections, err := UtilsInterface.GetNumCollections(client)
	if err != nil {
		return err
	}
	for i := 1; i <= int(numCollections); i++ {
		collection, err := AssetManagerInterface.GetCollection(client, uint16(i))
		if err != nil {
			return err
		}
		cache.CollectionsCache.Collections[collection.Id] = collection
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
