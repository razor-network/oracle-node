package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"regexp"
	"strconv"
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
		jobs = append(jobs, customJobs...)
	}

	for _, id := range collection.JobIDs {

		// Ignoring the Jobs which are already overriden and added to jobs array
		if !Contains(overriddenJobIds, id) {
			job, err := UtilsInterface.GetActiveJob(client, id)
			if err != nil {
				log.Errorf("Error in fetching job %d: %s", id, err)
				continue
			}
			jobs = append(jobs, job)
		}
	}

	if len(jobs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}

	dataToCommit, weight, err := UtilsInterface.GetDataToCommitFromJobs(jobs)
	if err != nil || len(dataToCommit) == 0 {
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
	collection, err := UtilsInterface.GetCollection(client, collectionId)
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	if !collection.Active {
		return bindings.StructsCollection{}, errors.New("collection inactive")
	}
	return collection, nil
}

func (*UtilsStruct) GetDataToCommitFromJobs(jobs []bindings.StructsJob) ([]*big.Int, []uint8, error) {
	var (
		data   []*big.Int
		weight []uint8
	)
	for _, job := range jobs {
		dataToAppend, err := UtilsInterface.GetDataToCommitFromJob(job)
		if err != nil {
			continue
		}
		data = append(data, dataToAppend)
		weight = append(weight, job.Weight)
	}
	return data, weight, nil
}

func (*UtilsStruct) GetDataToCommitFromJob(job bindings.StructsJob) (*big.Int, error) {
	var parsedJSON map[string]interface{}
	var (
		response []byte
		apiErr   error
	)

	// Fetch data from API with retry mechanism
	var parsedData interface{}
	if job.SelectorType == 0 {
		start := time.Now()
		response, apiErr = UtilsInterface.GetDataFromAPI(job.Url)
		if apiErr != nil {
			log.Errorf("Error in fetching data from API %s: %s", job.Url, apiErr)
			return nil, apiErr
		}
		elapsed := time.Since(start).Seconds()
		log.Debugf("Time taken to fetch the data from API : %s was %f", job.Url, elapsed)

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
		dataPoint, err := UtilsInterface.GetDataFromXHTML(job.Url, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed XHTML: ", err)
			return nil, err
		}
		// remove "," and currency symbols
		parsedData = regexp.MustCompile(`[\p{Sc},]`).ReplaceAllString(dataPoint, "")
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

	collectionCustomJobsPath := "assets.collection." + collection + ".custom jobs"
	customJobs := gjson.Get(jsonFileData, collectionCustomJobsPath).Array()
	if len(customJobs) == 0 {
		return nil
	}

	for i := 0; i < len(customJobs); i++ {
		customJobsData := customJobs[i].String()
		url := gjson.Get(customJobsData, "URL").String()
		selector := gjson.Get(customJobsData, "selector").String()
		power := int8(gjson.Get(customJobsData, "power").Int())
		weight := uint8(gjson.Get(customJobsData, "weight").Int())
		job := ConvertCustomJobToStructJob(types.CustomJob{
			URL:      url,
			Power:    power,
			Selector: selector,
			Weight:   weight,
		})
		collectionCustomJobs = append(collectionCustomJobs, job)
	}

	return collectionCustomJobs
}

func ConvertCustomJobToStructJob(customJob types.CustomJob) bindings.StructsJob {
	return bindings.StructsJob{
		Url:      customJob.URL,
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
		officialJobs := gjson.Get(dataString, officialJobsPath).String()
		if officialJobs != "" {
			job, err := UtilsInterface.GetActiveJob(client, jobIds[i])
			if err != nil {
				continue
			}
			job.Url = gjson.Get(officialJobs, "URL").String()
			job.Selector = gjson.Get(officialJobs, "selector").String()
			job.Weight = uint8(gjson.Get(officialJobs, "weight").Int())
			job.Power = int8(gjson.Get(officialJobs, "power").Int())

			overrideJobs = append(overrideJobs, job)
			overriddenJobIds = append(overriddenJobIds, jobIds[i])
		} else {
			continue
		}
	}

	return overrideJobs, overriddenJobIds
}
