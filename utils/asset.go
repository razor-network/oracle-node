package utils

import (
	"encoding/json"
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"regexp"
	"strconv"
)

func (*UtilsStruct) GetAssetManagerWithOpts(client *ethclient.Client) (*bindings.AssetManager, bind.CallOpts) {
	return UtilsInterface.GetAssetManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetNumAssets(client *ethclient.Client) (uint16, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		numAssets uint16
		err       error
	)
	err = retry.Do(
		func() error {
			numAssets, err = Options.GetNumAssets(client, &callOpts)
			if err != nil {
				log.Error("Error in fetching numAssets.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return numAssets, nil
}

func (*UtilsStruct) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	var jobs []bindings.StructsJob
	var JobIDs []uint16

	numAssets, err := UtilsInterface.GetNumAssets(client)
	if err != nil {
		return nil, err
	}
	for i := uint16(1); i <= numAssets; i++ {
		assetType, err := UtilsInterface.GetAssetType(client, i)
		if err != nil {
			return nil, err
		}
		if assetType == 1 {
			JobIDs = append(JobIDs, i)
		} else {
			continue
		}

	}

	for i := 0; i < len(JobIDs); i++ {
		jobId := JobIDs[i]
		job, err := UtilsInterface.GetActiveJob(client, jobId)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil

}

func (*UtilsStruct) GetNumActiveAssets(client *ethclient.Client) (*big.Int, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		numActiveAssets *big.Int
		err             error
	)
	err = retry.Do(
		func() error {
			numActiveAssets, err = Options.GetNumActiveCollections(client, &callOpts)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return numActiveAssets, nil
}

func (*UtilsStruct) GetAssetType(client *ethclient.Client, assetId uint16) (uint8, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		activeAsset types.Asset
		err         error
	)
	err = retry.Do(
		func() error {
			activeAsset, err = Options.GetAsset(client, &callOpts, assetId)
			if err != nil {
				log.Error("Error in fetching asset.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	if activeAsset.Job.Id == 0 {
		return 2, nil
	}
	return 1, nil
}

func (*UtilsStruct) GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	var collections []bindings.StructsCollection
	var CollectionIDs []uint16

	numAssets, err := UtilsInterface.GetNumAssets(client)
	if err != nil {
		return nil, err
	}
	for i := uint16(1); i <= numAssets; i++ {
		assetType, err := UtilsInterface.GetAssetType(client, i)
		if err != nil {
			return nil, err
		}
		if assetType == 2 {
			CollectionIDs = append(CollectionIDs, i)
		} else {
			continue
		}

	}

	for i := 0; i < len(CollectionIDs); i++ {
		collectionId := CollectionIDs[i]
		collection, err := UtilsInterface.GetCollection(client, collectionId)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}

	return collections, nil

}

func (*UtilsStruct) GetCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		asset types.Asset
		err   error
	)
	err = retry.Do(
		func() error {
			asset, err = Options.GetAsset(client, &callOpts, collectionId)
			if err != nil {
				log.Error("Error in fetching collection.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	return asset.Collection, nil
}

func (*UtilsStruct) GetActiveAssetIds(client *ethclient.Client) ([]uint16, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		activeAssetIds []uint16
		err            error
	)
	err = retry.Do(
		func() error {
			activeAssetIds, err = Options.GetActiveCollections(client, &callOpts)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return activeAssetIds, nil
}

func (*UtilsStruct) GetActiveAssetsData(client *ethclient.Client, epoch uint32) ([]*big.Int, error) {
	var data []*big.Int

	numOfAssets, err := UtilsInterface.GetNumAssets(client)
	if err != nil {
		return data, err
	}

	for assetIndex := 1; assetIndex <= int(numOfAssets); assetIndex++ {
		assetType, err := UtilsInterface.GetAssetType(client, uint16(assetIndex))
		if err != nil {
			log.Error("Error in fetching asset type: ", assetType)
			return nil, err
		}
		if assetType == 2 {
			activeCollection, err := UtilsInterface.GetActiveCollection(client, uint16(assetIndex))
			if err != nil {
				log.Error(err)
				if err.Error() == errors.New("collection inactive").Error() {
					continue
				}
				return nil, err
			}
			//Supply previous epoch to Aggregate in case if last reported value is required.
			collectionData, aggregationError := UtilsInterface.Aggregate(client, epoch-1, activeCollection)
			if aggregationError != nil {
				return nil, aggregationError
			}
			data = append(data, collectionData)
		}
	}
	return data, nil
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

		data, err := Options.ReadAll(jsonFile)
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
		if !Contains(overriddenJobIds, jobs) {
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
		return big.NewInt(int64(prevCommitmentData)), nil
	}
	return performAggregation(dataToCommit, weight, collection.AggregationMethod)
}

func (*UtilsStruct) GetActiveJob(client *ethclient.Client, jobId uint16) (bindings.StructsJob, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		job bindings.StructsJob
		err error
	)
	err = retry.Do(
		func() error {
			job, err = Options.Jobs(client, &callOpts, jobId)
			if err != nil {
				log.Errorf("Error in fetching job %d.... Retrying", jobId)
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
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
		apiErr = retry.Do(
			func() error {
				response, apiErr = UtilsInterface.GetDataFromAPI(job.Url)
				if apiErr != nil {
					log.Error("Error in fetching data from API: ", apiErr)
					return apiErr
				}
				return nil
			}, Options.RetryAttempts(core.MaxRetries))

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

	datum, err := Options.ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return nil, err
	}

	return MultiplyWithPower(datum, job.Power), err
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
