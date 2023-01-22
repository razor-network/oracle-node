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
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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

func (*UtilsStruct) GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error) {
	activeCollection, err := UtilsInterface.GetActiveCollection(client, collectionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//Supply previous epoch to Aggregate in case if last reported value is required.
	collectionData, aggregationError := UtilsInterface.Aggregate(client, epoch-1, activeCollection)
	if aggregationError != nil {
		return nil, aggregationError
	}
	return collectionData, nil
}

func (*UtilsStruct) Aggregate(client *ethclient.Client, previousEpoch uint32, collection bindings.StructsCollection) (*big.Int, error) {
	var jobs []bindings.StructsJob
	var overriddenJobIds []uint16

	// Checks if assets.JSON file exists
	assetsFilePath, err := path.PathUtilsInterface.GetJobFilePath()
	if err != nil {
		return nil, err
	}
	if _, err := path.OSUtilsInterface.Stat(assetsFilePath); !errors.Is(err, os.ErrNotExist) {
		log.Debug("Fetching the jobs from assets.json file...")
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
			log.Debugf("Got Custom Jobs from asset.json file: %+v", customJobs)
		}
		jobs = append(jobs, customJobs...)
	}

	for _, id := range collection.JobIDs {

		// Ignoring the Jobs which are already overriden and added to jobs array
		if !Contains(overriddenJobIds, id) {
			job, err := UtilsInterface.GetActiveJob(client, id)
			if err != nil {
				log.Errorf("Error in fetching job %d: %v", id, err)
				continue
			}
			jobs = append(jobs, job)
		}
	}

	if len(jobs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	localCache := cache.NewLocalCache(time.Second * time.Duration(core.StateLength))
	dataToCommit, weight, err := UtilsInterface.GetDataToCommitFromJobs(jobs, localCache)
	if err != nil || len(dataToCommit) == 0 {
		prevCommitmentData, err := UtilsInterface.FetchPreviousValue(client, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return prevCommitmentData, nil
	}
	localCache.StopCleanup()
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
	collection, err := UtilsInterface.GetCollection(client, collectionId)
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	if !collection.Active {
		return bindings.StructsCollection{}, errors.New("collection inactive")
	}
	return collection, nil
}

func (*UtilsStruct) GetDataToCommitFromJobs(jobs []bindings.StructsJob, localCache *cache.LocalCache) ([]*big.Int, []uint8, error) {
	var (
		data   []*big.Int
		weight []uint8
	)
	for _, job := range jobs {
		dataToAppend, err := UtilsInterface.GetDataToCommitFromJob(job, localCache)
		if err != nil {
			continue
		}
		log.Debugf("Job %s gives data %s", job.Url, dataToAppend)
		data = append(data, dataToAppend)
		weight = append(weight, job.Weight)
	}
	return data, weight, nil
}

func (*UtilsStruct) GetDataToCommitFromJob(job bindings.StructsJob, localCache *cache.LocalCache) (*big.Int, error) {
	var parsedJSON map[string]interface{}
	var (
		response            []byte
		apiErr              error
		dataSourceURLStruct types.DataSourceURL
	)
	log.Debugf("Getting the data to commit for job %s having job Id %d", job.Name, job.Id)
	if strings.HasPrefix(job.Url, "{") {
		log.Debug("Job URL passed is a struct containing URL along with type of request data")
		dataSourceURLInBytes := []byte(job.Url)

		err := json.Unmarshal(dataSourceURLInBytes, &dataSourceURLStruct)
		if err != nil {
			log.Errorf("Error in unmarshalling %s: %v", job.Url, err)
			return nil, err
		}
		log.Infof("URL Struct: %+v", dataSourceURLStruct)
	} else {
		log.Debug("Job URL passed is a direct URL: ", job.Url)
		isAPIKeyRequired, err := regexp.MatchString(core.APIKeyRegex, job.Url)
		if err != nil {
			log.Error("Error in matching api key regex in job url: ", err)
			return nil, err
		}
		if isAPIKeyRequired {
			keyword, APIKey, err := GetKeyWordAndAPIKeyFromENVFile(job.Url)
			if err != nil {
				log.Error("Error in getting value from env file: ", err)
				return nil, err
			}
			log.Debug("API key: ", APIKey)
			keywordWithDollar := `$` + keyword
			log.Debug("Keyword to replace in url: ", keywordWithDollar)
			urlWithAPIKey := strings.Replace(job.Url, keywordWithDollar, APIKey, 1)
			log.Debug("URl with API key: ", urlWithAPIKey)
			job.Url = urlWithAPIKey
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
		response, apiErr = UtilsInterface.GetDataFromAPI(dataSourceURLStruct, localCache)
		if apiErr != nil {
			log.Errorf("Error in fetching data from API %s: %v", job.Url, apiErr)
			return nil, apiErr
		}
		elapsed := time.Since(start).Seconds()
		log.Debugf("Time taken to fetch the data from API : %s was %f", dataSourceURLStruct.URL, elapsed)

		err := json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
			return nil, err
		}
		parsedData, err = UtilsInterface.GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed data: ", err)
			return nil, err
		}
	} else {
		//TODO: Add retry here.
		dataPoint, err := UtilsInterface.GetDataFromXHTML(dataSourceURLStruct, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed XHTML: ", err)
			return nil, err
		}
		// remove "," and currency symbols
		parsedData = regexp.MustCompile(`[\p{Sc}, ]`).ReplaceAllString(dataPoint, "")
	}

	datum, err := UtilsInterface.ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
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

func GetKeyWordAndAPIKeyFromENVFile(url string) (string, string, error) {
	envFilePath, err := path.PathUtilsInterface.GetDotENVFilePath()
	if err != nil {
		log.Error("Error in getting env file path: ", err)
		return "", "", err
	}
	log.Debug("GetKeyWordAndAPIKeyFromENVFile: .env file path: ", envFilePath)

	log.Info("Loading env file...")
	envFileMap, err := godotenv.Read(envFilePath)
	if err != nil {
		log.Error("Error in getting env file map: ", err)
		return "", "", err
	}
	log.Debugf("GetKeyWordAndAPIKeyFromENVFile: ENV file map: %v", envFileMap)
	for keyword, APIKey := range envFileMap {
		keywordWithDollar := `$` + keyword
		isTheKeywordPresentInURL := strings.Contains(url, keywordWithDollar)
		if isTheKeywordPresentInURL {
			log.Infof("Found the keyword %s in env file: ", keyword)
			return keyword, APIKey, nil
		}
	}
	return "", "", errors.New("no value found in env file")
}
