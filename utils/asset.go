package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"razor/core"
	"razor/pkg/bindings"
	"regexp"
	"strconv"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
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
	for i := 0; i < int(numJobs); i++ {
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
		activeAssetIds []uint16
		err            error
	)
	err = retry.Do(
		func() error {
			activeAssetIds, err = AssetManagerInterface.GetActiveCollections(client)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return activeAssetIds, nil
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
	if len(collection.JobIDs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	var jobs []bindings.StructsJob
	for _, id := range collection.JobIDs {
		job, err := UtilsInterface.GetActiveJob(client, id)
		if err != nil {
			log.Errorf("Error in fetching job %d: %s", id, err)
			continue
		}
		jobs = append(jobs, job)
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
	jobPath, err := PathInterface.GetJobFilePath()
	if err != nil {
		return nil, nil, err
	}
	overrideJobData, err := UtilsInterface.ReadJSONData(jobPath)
	if err != nil {
		log.Error(err)
	}
	var (
		data   []*big.Int
		weight []uint8
	)
	for _, job := range jobs {
		if overrideJobData[strconv.Itoa(int(job.Id))] != nil {
			_job := overrideJobData[strconv.Itoa(int(job.Id))]
			job.Url = _job.Url
			job.Selector = _job.Selector
			job.SelectorType = _job.SelectorType
			job.Power = _job.Power
		}
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
			}, RetryInterface.RetryAttempts(core.MaxRetries))

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
	log.Debugf("SEED: %s", hex.EncodeToString(seed))
	for i := 0; i < int(toAssign); i++ {
		assigned := UtilsInterface.Prng(uint32(numActiveCollections), solsha3.SoliditySHA3([]string{"bytes32", "uint256"}, []interface{}{"0x" + hex.EncodeToString(seed), big.NewInt(int64(i))}))
		assignedCollections[int(assigned.Int64())] = true
		seqAllottedCollections = append(seqAllottedCollections, assigned)
	}
	return assignedCollections, seqAllottedCollections, nil
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
