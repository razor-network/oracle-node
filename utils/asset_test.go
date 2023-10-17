package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"math/big"
	"os"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/path"
	pathMocks "razor/path/mocks"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

func TestAggregate(t *testing.T) {
	var client *ethclient.Client
	var previousEpoch uint32
	var fileInfo fs.FileInfo

	job := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	collection := bindings.StructsCollection{
		Active:            true,
		Id:                4,
		Power:             2,
		AggregationMethod: 2,
		JobIDs:            []uint16{1, 2, 3},
		Name:              "ethCollectionMean",
	}

	type args struct {
		collection            bindings.StructsCollection
		activeJob             bindings.StructsJob
		activeJobErr          error
		dataToCommit          []*big.Int
		dataToCommitErr       error
		weight                []uint8
		prevCommitmentData    *big.Int
		prevCommitmentDataErr error
		assetFilePath         string
		assetFilePathErr      error
		statErr               error
		jsonFile              *os.File
		jsonFileErr           error
		fileData              []byte
		fileDataErr           error
		overrrideJobs         []bindings.StructsJob
		overrideJobIds        []uint16
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When Aggregate() executes successfully",
			args: args{
				collection:         collection,
				activeJob:          job,
				dataToCommit:       []*big.Int{big.NewInt(2)},
				weight:             []uint8{100},
				prevCommitmentData: big.NewInt(1),
				assetFilePath:      "",
				statErr:            nil,
			},
			want:    big.NewInt(2),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting activeJob",
			args: args{
				collection:         collection,
				activeJobErr:       errors.New("activeJob error"),
				dataToCommit:       []*big.Int{big.NewInt(2)},
				weight:             []uint8{100},
				prevCommitmentData: big.NewInt(1),
				assetFilePath:      "",
				statErr:            errors.New(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting dataToCommit",
			args: args{
				collection:         collection,
				activeJob:          job,
				dataToCommitErr:    errors.New("dataToCommit error"),
				weight:             []uint8{100},
				prevCommitmentData: big.NewInt(1),
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting prevCommitmentData",
			args: args{
				collection:            collection,
				activeJob:             job,
				dataToCommit:          []*big.Int{big.NewInt(2)},
				weight:                []uint8{100},
				prevCommitmentDataErr: errors.New("prevCommitmentData error"),
			},
			want:    big.NewInt(2),
			wantErr: false,
		},
		{
			name: "Test 5: When there is an error in getting prevCommitmentData",
			args: args{
				collection:            collection,
				activeJob:             job,
				dataToCommitErr:       errors.New("dataToCommit error"),
				weight:                []uint8{100},
				prevCommitmentDataErr: errors.New("prevCommitmentData error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When there is a nil collection",
			args: args{
				collection: bindings.StructsCollection{},
				statErr:    errors.New("file does not exist"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 7: When there is an error in getting asset file path",
			args: args{
				assetFilePathErr: errors.New("path error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 8: When there is an error in opening json file",
			args: args{
				assetFilePath: "./razor/assets.json",
				jsonFileErr:   errors.New("open file error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 9: When assets are fetched from json file and there is an error in reading json file",
			args: args{
				assetFilePath: "./razor/assets.json",
				jsonFile:      &os.File{},
				statErr:       nil,
				fileDataErr:   errors.New("reading file error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			pathUtilsMock := new(pathMocks.PathInterface)
			osUtilsMock := new(pathMocks.OSInterface)
			ioMock := new(mocks.IOUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
				IOInterface:    ioMock,
			}
			path.PathUtilsInterface = pathUtilsMock
			path.OSUtilsInterface = osUtilsMock
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetActiveJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.activeJob, tt.args.activeJobErr)
			utilsMock.On("GetDataToCommitFromJobs", mock.Anything, mock.Anything).Return(tt.args.dataToCommit, tt.args.weight, tt.args.dataToCommitErr)
			utilsMock.On("FetchPreviousValue", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint16")).Return(tt.args.prevCommitmentData, tt.args.prevCommitmentDataErr)
			pathUtilsMock.On("GetJobFilePath").Return(tt.args.assetFilePath, tt.args.assetFilePathErr)
			osUtilsMock.On("Stat", mock.Anything).Return(fileInfo, tt.args.statErr)
			osUtilsMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioMock.On("ReadAll", mock.Anything).Return(tt.args.fileData, tt.args.fileDataErr)
			utilsMock.On("HandleOfficialJobsFromJSONFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.overrrideJobs, tt.args.overrideJobIds)

			got, err := utils.Aggregate(client, previousEpoch, tt.args.collection, &cache.LocalCache{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Aggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveCollectionIds(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		activeAssetIds    []uint16
		activeAssetIdsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetActiveCollections() executes successfully",
			args: args{
				activeAssetIds: []uint16{1, 2},
			},
			want:    []uint16{1, 2},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting activeAssetIds",
			args: args{
				activeAssetIdsErr: errors.New("activeAssetIds error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			assetManagerMock := new(mocks.AssetManagerUtils)
			retryMock := new(mocks.RetryUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetManagerMock,
				RetryInterface:        retryMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			assetManagerMock.On("GetActiveCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.activeAssetIds, tt.args.activeAssetIdsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetActiveCollectionIds(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveCollection(t *testing.T) {
	var client *ethclient.Client
	var collectionId uint16

	collectionEth := bindings.StructsCollection{Active: true,
		Id:                2,
		Power:             2,
		AggregationMethod: 2,
		JobIDs:            []uint16{1, 2},
		Name:              "ethCollectionMean",
	}

	collectionEthInactive := bindings.StructsCollection{Active: false, Id: 2, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2}, Name: "ethCollectionMean",
	}

	type args struct {
		collection    bindings.StructsCollection
		collectionErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsCollection
		wantErr bool
	}{
		{
			name: "Test 1: When GetActiveCollection() executes successfully",
			args: args{
				collection: collectionEth,
			},
			want:    collectionEth,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting collection",
			args: args{
				collectionErr: errors.New("collection error"),
			},
			want:    bindings.StructsCollection{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an inactive collection",
			args: args{
				collection: collectionEthInactive,
			},
			want:    bindings.StructsCollection{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetCollection", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.collection, tt.args.collectionErr)

			got, err := utils.GetActiveCollection(client, collectionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveJob(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var jobId uint16

	jobEth := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	type args struct {
		job    bindings.StructsJob
		jobErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsJob
		wantErr bool
	}{
		{
			name: "Test 1: When GetActiveJob() executes successfully",
			args: args{
				job: jobEth,
			},
			want:    jobEth,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting job",
			args: args{
				jobErr: errors.New("job error"),
			},
			want:    bindings.StructsJob{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			assetManagerMock.On("Jobs", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.job, tt.args.jobErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetActiveJob(client, jobId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCollection(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var collectionId uint16

	type args struct {
		asset    bindings.StructsCollection
		assetErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsCollection
		wantErr bool
	}{
		{
			name: "Test 1: When GetCollection() executes successfully",
			args: args{
				asset: bindings.StructsCollection{},
			},
			want:    bindings.StructsCollection{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting asset",
			args: args{
				assetErr: errors.New("asset error"),
			},
			want:    bindings.StructsCollection{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			assetManagerMock.On("GetCollection", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.asset, tt.args.assetErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetCollection(client, collectionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCollections(t *testing.T) {
	var client *ethclient.Client

	collectionListArray := []bindings.StructsCollection{
		{
			Active:            true,
			Id:                7,
			Power:             2,
			AggregationMethod: 2,
			JobIDs:            []uint16{1, 2, 3},
			Name:              "ethCollectionMean",
		},
	}

	type args struct {
		numAssets    uint16
		numAssetsErr error
		assetType    uint8
		//assetTypeErr  error
		collection    bindings.StructsCollection
		collectionErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []bindings.StructsCollection
		wantErr bool
	}{
		{
			name: "Test 1: When GetAllCollections() executes successfully",
			args: args{
				numAssets:  1,
				assetType:  2,
				collection: collectionListArray[0],
			},
			want:    collectionListArray,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numAssets",
			args: args{
				numAssetsErr: errors.New("numAssets error"),
				assetType:    2,
				collection:   collectionListArray[0],
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting collection",
			args: args{
				numAssets:     1,
				assetType:     2,
				collectionErr: errors.New("collection error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			assetMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetNumCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numAssets, tt.args.numAssetsErr)
			assetMock.On("GetCollection", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.collection, tt.args.collectionErr)

			got, err := utils.GetAllCollections(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataToCommitFromJobs(t *testing.T) {
	jobsArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: `{"type": "GET","url": "https://api.gemini.com/v1/pubticker/ethusd","body": {},"header": {}}`,
		}, {Id: 2, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: `{"type": "GET","url": "https://api.gemini.com/v1/pubticker/ethusd","body": {},"header": {}}`,
		},
	}

	type args struct {
		jobPath            string
		jobPathErr         error
		overrideJobData    map[string]*types.StructsJob
		overrideJobDataErr error
		dataToAppend       *big.Int
		dataToAppendErr    error
	}
	tests := []struct {
		name    string
		args    args
		want    []*big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetDataToCommitFromJobs() executes successfully",
			args: args{
				jobPath: "",
				overrideJobData: map[string]*types.StructsJob{"1": {
					Id: 2, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini", Selector: "last",
					Url: `{"type": "GET","url": "https://api.gemini.com/v1/pubticker/ethusd","body": {},"header": {}}`,
				}},
				dataToAppend: big.NewInt(1),
			},
			want:    []*big.Int{big.NewInt(1), big.NewInt(1)},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting overrideJobData",
			args: args{
				jobPath:            "",
				overrideJobDataErr: errors.New("overrideJobData error"),
				dataToAppend:       big.NewInt(1),
			},
			want:    []*big.Int{big.NewInt(1), big.NewInt(1)},
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting jobPath",
			args: args{
				jobPathErr:      errors.New("jobPath error"),
				overrideJobData: map[string]*types.StructsJob{},
				dataToAppend:    big.NewInt(1),
			},
			want:    []*big.Int{big.NewInt(1), big.NewInt(1)},
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting dataToAppend",
			args: args{
				jobPath:         "",
				overrideJobData: map[string]*types.StructsJob{},
				dataToAppendErr: errors.New("dataToAppend error"),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			pathMock := new(mocks.PathUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
				PathInterface:  pathMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetDataToCommitFromJob", mock.Anything, mock.Anything).Return(tt.args.dataToAppend, tt.args.dataToAppendErr)

			got, _, err := utils.GetDataToCommitFromJobs(jobsArray, &cache.LocalCache{})
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataToCommitFromJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataToCommitFromJobs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataToCommitFromJob(t *testing.T) {
	job := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_kraken", Selector: "result.XETHZUSD.c[0]",
		Url: `{"type": "GET","url": "https://api.kraken.com/0/public/Ticker?pair=ETHUSD","body": {},"header": {}}`,
	}

	job1 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_sample", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd/apiKey=${SAMPLE_API_KEY_NEW}",
	}

	postJob := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_sample", Selector: "result",
		Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","method":"eth_call","params":[{"to":"0xb27308f9f90d607463bb33ea1bebb41c27ce5ab6","data":"0xf7729d43000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000"},"latest"],"id":5},"header": {"content-type": "application/json"}, "returnType": "hex"}`,
	}

	invalidXHTMLJob := bindings.StructsJob{Id: 3, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: `{"type1": "GET","url1": "https://api.gemini.com/v1/pubticker/ethusd","body1": {},"header1": {}}`,
	}

	type args struct {
		job bindings.StructsJob
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetDataToCommitFromJob() executes successfully",
			args: args{
				job: job,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is a case to pick up API key from environment variable file and keyword is not present",
			args: args{
				job: job1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When GetDataToCommitFromJob() executes successfully for a POST Job",
			args: args{
				job: postJob,
			},
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in unmarshalling invalid dataSourceStruct for XHTML job",
			args: args{
				job: invalidXHTMLJob,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			pathUtilsMock := new(pathMocks.PathInterface)
			path.PathUtilsInterface = pathUtilsMock

			pathUtilsMock.On("GetDotENVFilePath", mock.Anything).Return("$HOME/.razor/.env", nil)
			lc := cache.NewLocalCache(time.Second * 10)
			data, err := utils.GetDataToCommitFromJob(tt.args.job, lc)
			fmt.Println("JOB returns data: ", data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataToCommitFromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetJobs(t *testing.T) {
	var client *ethclient.Client

	jobsArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
	}

	type args struct {
		numJobs    uint16
		numJobsErr error
		assetType  uint8
		//assetTypeErr error
		activeJob    bindings.StructsJob
		activeJobErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []bindings.StructsJob
		wantErr bool
	}{
		{
			name: "Test 1: When GetJobs() executes successfully",
			args: args{
				numJobs:   1,
				assetType: 1,
				activeJob: jobsArray[0],
			},
			want:    jobsArray,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numJobs",
			args: args{
				numJobsErr: errors.New("numJobs error"),
				assetType:  1,
				activeJob:  jobsArray[0],
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting activeJob",
			args: args{
				numJobs:      1,
				assetType:    1,
				activeJobErr: errors.New("activeJob error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assetMock := new(mocks.AssetManagerUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetMock,
			}
			utils := StartRazor(optionsPackageStruct)

			assetMock.On("GetNumJobs", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numJobs, tt.args.numJobsErr)
			utilsMock.On("GetActiveJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.activeJob, tt.args.activeJobErr)

			got, err := utils.GetJobs(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJobs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumActiveCollections(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		numOfActiveAssets    uint16
		numOfActiveAssetsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetNumActiveCollections() executes successfully",
			args: args{
				numOfActiveAssets: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numOfActiveAssets",
			args: args{
				numOfActiveAssetsErr: errors.New("numOfActiveAssets error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			assetManagerMock.On("GetNumActiveCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numOfActiveAssets, tt.args.numOfActiveAssetsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumActiveCollections(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumActiveCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumActiveCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumCollections(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		numOfAssets    uint16
		numOfAssetsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetNumCollections() executes successfully",
			args: args{
				numOfAssets: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numOfAssets",
			args: args{
				numOfAssetsErr: errors.New("numOfAssets error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			assetManagerMock.On("GetNumCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numOfAssets, tt.args.numOfAssetsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumCollections(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAssetManagerWithOpts(t *testing.T) {

	var callOpts bind.CallOpts
	var assetManager *bindings.CollectionManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetCollectionManager", mock.AnythingOfType("*ethclient.Client")).Return(assetManager)

	gotAssetManager, gotCallOpts := utils.GetCollectionManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetCollectionManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotAssetManager, assetManager) {
		t.Errorf("GetAssetkManagerWithOpts() got assetManager = %v, want %v", gotAssetManager, assetManager)
	}
}

func TestGetCustomJobsFromJSONFile(t *testing.T) {
	type args struct {
		collection   string
		jsonFileData string
	}
	tests := []struct {
		name string
		args args
		want []bindings.StructsJob
	}{
		{
			name: "Test 1: When collection is present in json file string",
			args: args{
				collection:   "ethCollection",
				jsonFileData: jsonDataString,
			},
			want: []bindings.StructsJob{
				{
					Url:      "http://127.0.0.1/eth1",
					Selector: "eth1",
					Power:    2,
					Weight:   3,
				},
				{
					Url:      "http://127.0.0.1/eth2",
					Selector: "eth2",
					Power:    2,
					Weight:   2,
				},
			},
		},
		{
			name: "Test 2: When collection is not present in json file string",
			args: args{
				collection:   "btcCollection",
				jsonFileData: jsonDataString,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCustomJobsFromJSONFile(tt.args.collection, tt.args.jsonFileData)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomJobsFromJSONFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertCustomJobToStructJob(t *testing.T) {
	type args struct {
		customJob types.CustomJob
	}
	tests := []struct {
		name string
		args args
		want bindings.StructsJob
	}{
		{
			name: "Test 1: Converting customJob struct to bindings.Job struct",
			args: args{
				customJob: types.CustomJob{
					URL:    "http://api.coinbase.com/eth2",
					Name:   "eth_coinBase",
					Power:  3,
					Weight: 2,
				},
			},
			want: bindings.StructsJob{
				Url:    "http://api.coinbase.com/eth2",
				Name:   "eth_coinBase",
				Power:  3,
				Weight: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertCustomJobToStructJob(tt.args.customJob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertCustomJobToStructJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleOfficialJobsFromJSONFile(t *testing.T) {
	var client *ethclient.Client
	ethCollection := bindings.StructsCollection{
		Active: true, Id: 7, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1}, Name: "ethCollection",
	}

	ethCollection1 := bindings.StructsCollection{
		Active: true, Id: 7, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2, 3}, Name: "ethCollection",
	}

	type args struct {
		collection bindings.StructsCollection
		dataString string
		job        bindings.StructsJob
		jobErr     error
	}
	tests := []struct {
		name               string
		args               args
		want               []bindings.StructsJob
		wantOverrideJobIds []uint16
	}{
		{
			name: "Test 1: When officialJobs for collection is present in assets.json",
			args: args{
				collection: ethCollection,
				dataString: jsonDataString,
				job: bindings.StructsJob{
					Id: 1,
				},
			},
			want: []bindings.StructsJob{
				{
					Id:       1,
					Url:      "http://kucoin.com/eth1",
					Selector: "eth1",
					Power:    2,
					Weight:   2,
				},
			},
			wantOverrideJobIds: []uint16{1},
		},
		{
			name: "Test 2: When officialJobs for collection is not present in assets.json",
			args: args{
				collection: ethCollection,
				dataString: "",
				job: bindings.StructsJob{
					Id: 1,
				},
			},
			want:               nil,
			wantOverrideJobIds: nil,
		},
		{
			name: "Test 3: When there is an error from GetActiveJob()",
			args: args{
				collection: ethCollection,
				dataString: jsonDataString,
				jobErr:     errors.New("job error"),
			},
			want:               nil,
			wantOverrideJobIds: nil,
		},
		{
			name: "Test 4: When multiple jobIds are needed to be overridden from official jobs",
			args: args{
				collection: ethCollection1,
				dataString: jsonDataString,
				job: bindings.StructsJob{
					Id:       1,
					Url:      "http://kraken.com/eth1",
					Selector: "data.ETH",
					Power:    3,
					Weight:   1,
				},
			},
			want: []bindings.StructsJob{
				{
					Id:       1,
					Url:      "http://kucoin.com/eth1",
					Selector: "eth1",
					Power:    2,
					Weight:   2,
				},
				{
					Id:       1,
					Url:      "http://api.coinbase.com/eth2",
					Selector: "eth2",
					Power:    3,
					Weight:   2,
				},
			},
			wantOverrideJobIds: []uint16{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			utilsMock.On("GetActiveJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.job, tt.args.jobErr)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			gotJobs, gotOverrideJobIds := utils.HandleOfficialJobsFromJSONFile(client, tt.args.collection, tt.args.dataString)
			if !reflect.DeepEqual(gotJobs, tt.want) {
				t.Errorf("HandleOfficialJobsFromJSONFile() gotJobs = %v, want %v", gotJobs, tt.want)
			}
			if !reflect.DeepEqual(gotOverrideJobIds, tt.wantOverrideJobIds) {
				t.Errorf("HandleOfficialJobsFromJSONFile() gotOverrideJobIds = %v, want %v", gotOverrideJobIds, tt.wantOverrideJobIds)
			}
		})
	}
}

var jsonDataString = `{
  "assets": {
    "collection": {
      "ethCollection": {
        "power": 2,
        "official jobs": {
          "1": {
            "URL": "http://kucoin.com/eth1",
            "selector": "eth1",
            "power": 2,
            "weight": 2
          },
          "2": {
            "URL": "http://api.coinbase.com/eth2",
            "selector": "eth2",
            "power": 3,
            "weight": 2
          }
        },
        "custom jobs": [
          {
            "URL": "http://127.0.0.1/eth1",
            "selector": "eth1",
            "power": 2,
            "weight": 3
          },
          {
            "URL": "http://127.0.0.1/eth2",
            "selector": "eth2",
            "power": 2,
            "weight": 2
          },
        ]
      }
    }
  }
}`

func TestGetAggregatedDataOfCollection(t *testing.T) {
	var (
		client       *ethclient.Client
		collectionId uint16
		epoch        uint32
	)
	type args struct {
		activeCollection    bindings.StructsCollection
		activeCollectionErr error
		collectionData      *big.Int
		aggregationErr      error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetAggregatedDataOfCollection() executes successfully",
			args: args{
				activeCollection: bindings.StructsCollection{},
				collectionData:   &big.Int{},
			},
			want:    &big.Int{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting activeCollection",
			args: args{
				activeCollectionErr: errors.New("error in getting activeCollection"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an aggregation error",
			args: args{
				activeCollection: bindings.StructsCollection{},
				aggregationErr:   errors.New("error in aggregation"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetActiveCollection", mock.Anything, mock.Anything).Return(tt.args.activeCollection, tt.args.activeCollectionErr)
			utilsMock.On("Aggregate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.collectionData, tt.args.aggregationErr)

			got, err := utils.GetAggregatedDataOfCollection(client, collectionId, epoch, &cache.LocalCache{})
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAggregatedDataOfCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAggregatedDataOfCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAssignedCollections(t *testing.T) {
	var (
		client               *ethclient.Client
		numActiveCollections uint16
		seed                 []byte
	)
	type args struct {
		toAssign    uint16
		toAssignErr error
		assigned    *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]bool
		want1   []*big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetAssignedCollections() executes successfully",
			args: args{
				toAssign: 1,
				assigned: &big.Int{},
			},
			want:    map[int]bool{0: true},
			want1:   []*big.Int{big.NewInt(0)},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting toAssign",
			args: args{
				toAssignErr: errors.New("error in getting toAssign"),
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("ToAssign", mock.Anything).Return(tt.args.toAssign, tt.args.toAssignErr)
			utilsMock.On("Prng", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.assigned)

			got, got1, err := utils.GetAssignedCollections(client, numActiveCollections, seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAssignedCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAssignedCollections() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAssignedCollections() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetLeafIdOfACollection(t *testing.T) {
	var (
		client       *ethclient.Client
		collectionId uint16
	)
	type args struct {
		leafId    uint16
		leafIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetLeafIdOfACollection() executes successfully",
			args: args{
				leafId: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting leafId",
			args: args{
				leafIdErr: errors.New("error in getting leafId"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)
			assetManagerMock.On("GetLeafIdOfACollection", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.leafId, tt.args.leafIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetLeafIdOfACollection(client, collectionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLeafIdOfACollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLeafIdOfACollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCollectionIdFromIndex(t *testing.T) {
	var (
		client      *ethclient.Client
		medianIndex uint16
	)
	type args struct {
		collectionId    uint16
		collectionIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetCollectionIdFromIndex() executes successfully",
			args: args{
				collectionId: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting collectionId",
			args: args{
				collectionIdErr: errors.New("error in getting collectionId"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)
			assetManagerMock.On("GetCollectionIdFromIndex", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.collectionId, tt.args.collectionIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetCollectionIdFromIndex(client, medianIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollectionIdFromIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCollectionIdFromIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCollectionIdFromLeafId(t *testing.T) {
	var (
		client *ethclient.Client
		leafId uint16
	)
	type args struct {
		collectionId    uint16
		collectionIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetCollectionIdFromLeafId() executes successfully",
			args: args{
				collectionId: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting collectionId",
			args: args{
				collectionIdErr: errors.New("error in getting collectionId"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			assetManagerMock := new(mocks.AssetManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				AssetManagerInterface: assetManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)
			assetManagerMock.On("GetCollectionIdFromLeafId", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.collectionId, tt.args.collectionIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetCollectionIdFromLeafId(client, leafId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollectionIdFromLeafId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCollectionIdFromLeafId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceValueWithDataFromENVFile(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		envVariables   map[string]string
		expectedOutput string
	}{
		{
			name:           "single env variable",
			value:          "API key is ${API_KEY}",
			envVariables:   map[string]string{"API_KEY": "12345"},
			expectedOutput: "API key is 12345",
		},
		{
			name:           "multiple env variables",
			value:          "API key is ${API_KEY} and secret is ${SECRET}",
			envVariables:   map[string]string{"API_KEY": "12345", "SECRET": "abcdef"},
			expectedOutput: "API key is 12345 and secret is abcdef",
		},
		{
			name:           "no env variable in value",
			value:          "API key is present",
			envVariables:   map[string]string{"API_KEY": "12345"},
			expectedOutput: "API key is present",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, val := range tt.envVariables {
				os.Setenv(key, val)
			}

			// Execute function
			var re = regexp.MustCompile(core.APIKeyRegex)
			result := ReplaceValueWithDataFromENVFile(re, tt.value)

			// Check result
			if result != tt.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", tt.expectedOutput, result)
			}

			// Cleanup environment variables
			for key := range tt.envVariables {
				os.Unsetenv(key)
			}
		})
	}
}
