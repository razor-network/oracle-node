package utils

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"math/big"
	"net/http"
	"os"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/path"
	pathMocks "razor/path/mocks"
	"razor/pkg/bindings"
	"razor/rpc"
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

var rpcManager = rpc.RPCManager{
	BestEndpoint: &rpc.RPCEndpoint{
		Client: &ethclient.Client{},
	},
}

var rpcParameters = rpc.RPCParameters{
	Ctx:        context.Background(),
	RPCManager: &rpcManager,
}

func TestAggregate(t *testing.T) {
	var (
		previousEpoch uint32
		fileInfo      fs.FileInfo
	)

	job1 := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	job2 := bindings.StructsJob{Id: 2, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	collection := bindings.StructsCollection{
		Active:            true,
		Id:                4,
		Power:             2,
		AggregationMethod: 2,
		JobIDs:            []uint16{1, 2},
		Name:              "ethCollectionMean",
	}

	type args struct {
		collection            bindings.StructsCollection
		jobCacheError         bool
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
				dataToCommit:       []*big.Int{big.NewInt(3827200), big.NewInt(3828474)},
				weight:             []uint8{1, 1},
				prevCommitmentData: big.NewInt(1),
				assetFilePath:      "",
				statErr:            nil,
			},
			want:    big.NewInt(3827837),
			wantErr: false,
		},
		{
			name: "Test 2: When the job is not present in cache",
			args: args{
				collection:         collection,
				jobCacheError:      true,
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
			commitParams := &types.CommitParams{
				JobsCache:        cache.NewJobsCache(),
				CollectionsCache: cache.NewCollectionsCache(),
			}
			if !tt.args.jobCacheError {
				commitParams.JobsCache.Jobs[job1.Id] = job1
				commitParams.JobsCache.Jobs[job2.Id] = job2
				commitParams.CollectionsCache.Collections[collection.Id] = collection
			}

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

			utilsMock.On("GetDataToCommitFromJobs", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.dataToCommit, tt.args.weight, tt.args.dataToCommitErr)
			utilsMock.On("FetchPreviousValue", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.prevCommitmentData, tt.args.prevCommitmentDataErr)
			pathUtilsMock.On("GetJobFilePath").Return(tt.args.assetFilePath, tt.args.assetFilePathErr)
			osUtilsMock.On("Stat", mock.Anything).Return(fileInfo, tt.args.statErr)
			osUtilsMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioMock.On("ReadAll", mock.Anything).Return(tt.args.fileData, tt.args.fileDataErr)
			utilsMock.On("HandleOfficialJobsFromJSONFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.overrrideJobs, tt.args.overrideJobIds)

			got, err := utils.Aggregate(rpcParameters, previousEpoch, tt.args.collection, commitParams)

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

			got, err := utils.GetActiveCollectionIds(rpcParameters)
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
	collectionEth := bindings.StructsCollection{
		Active: true, Id: 1, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2},
		Name: "ethCollectionMean",
	}

	collectionEthInactive := bindings.StructsCollection{
		Active: false, Id: 2, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2},
		Name: "ethCollectionMean",
	}

	type args struct {
		collectionId       uint16
		collectionCacheErr bool
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
				collectionId: 1,
			},
			want:    collectionEth,
			wantErr: false,
		},
		{
			name: "Test 2: When the collection is not present in cache",
			args: args{
				collectionCacheErr: true,
			},
			want:    bindings.StructsCollection{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an inactive collection",
			args: args{
				collectionId: 2,
			},
			want:    bindings.StructsCollection{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collectionCache := cache.NewCollectionsCache()
			collectionCache.Collections[collectionEth.Id] = collectionEth
			collectionCache.Collections[collectionEthInactive.Id] = collectionEthInactive

			utils := UtilsStruct{}
			got, err := utils.GetActiveCollection(collectionCache, tt.args.collectionId)
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

			got, err := utils.GetActiveJob(rpcParameters, jobId)
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

			got, err := utils.GetCollection(rpcParameters, collectionId)
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

			utilsMock.On("GetNumCollections", mock.Anything).Return(tt.args.numAssets, tt.args.numAssetsErr)
			utilsMock.On("GetCollection", mock.Anything, mock.AnythingOfType("uint16")).Return(tt.args.collection, tt.args.collectionErr)

			got, err := utils.GetAllCollections(rpcParameters)
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
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        2,
			MaxIdleConnsPerHost: 1,
		},
	}

	jobsArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 0, Weight: 10,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
		{Id: 2, SelectorType: 0, Weight: 20,
			Power: 2, Name: "ethusd_kraken", Selector: "result.XETHZUSD.c[0]",
			Url: `{"type": "GET","url": "https://api.kraken.com/0/public/Ticker?pair=ETHUSD","body": {},"header": {}}`,
		},
		{Id: 3, SelectorType: 0, Weight: 30,
			Power: 2, Name: "ethusd_kucoin", Selector: "data.ETH",
			Url: `{"type": "GET","url": "https://api.kucoin.com/api/v1/prices?base=USD&currencies=ETH","body": {},"header": {}}`,
		},
		{Id: 4, SelectorType: 0, Weight: 40,
			Power: 2, Name: "ethusd_coinbase", Selector: "data.amount",
			Url: `{"type": "GET","url": "https://api.coinbase.com/v2/prices/ETH-USD/spot","body": {},"header": {}}`,
		},
		// This job returns an error which will not add any value to data or weight array
		{Id: 5, SelectorType: 0, Weight: 10,
			Power: 2, Name: "ethusd_gemini_incorrect", Selector: "last1",
			Url: `{"type": "GET","url": "https://api.gemini.com/v1/pubticker/ethusd1","body": {},"header": {}}`,
		},
		{Id: 6, SelectorType: 0, Weight: 100,
			Power: 6, Name: "ethusd_uniswapv2", Selector: "result",
			Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","id":7269270904970082,"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","data":"0xd06ca61f0000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000050de6856358cc35f3a9a57eaaa34bd4cb707d2cd0000000000000000000000008e870d67f660d95d5be530380d0ec0bd388289e1","to":"0x7a250d5630b4cf539739df2c5dacb4c659f2488d"},"latest"]},"header": {"content-type": "application/json"}, "returnType": "hexArray[1]"}`,
		},
		{Id: 7, SelectorType: 0, Weight: 100,
			Power: 2, Name: "ethusd_uniswapv3", Selector: "result",
			Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","method":"eth_call","params":[{"to":"0xb27308f9f90d607463bb33ea1bebb41c27ce5ab6","data":"0xf7729d43000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000"},"latest"],"id":5},"header": {"content-type": "application/json"}, "returnType": "hex"}`,
		},
		// This is a duplicate job to check if cache is working correctly for POST Jobs
		{Id: 8, SelectorType: 0, Weight: 100,
			Power: 2, Name: "ethusd_uniswapv3_duplicate", Selector: "result",
			Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","method":"eth_call","params":[{"to":"0xb27308f9f90d607463bb33ea1bebb41c27ce5ab6","data":"0xf7729d43000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000"},"latest"],"id":5},"header": {"content-type": "application/json"}, "returnType": "hex"}`,
		},
		// This is a duplicate job to check if cache is working correctly for GET Jobs
		{Id: 9, SelectorType: 0, Weight: 10,
			Power: 2, Name: "ethusd_gemini_duplicate", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
	}

	type args struct {
		jobs []bindings.StructsJob
	}

	tests := []struct {
		name            string
		args            args
		wantArrayLength int
	}{
		{
			name: "Test 1: Getting values from set of jobs of length 4",
			args: args{
				jobs: jobsArray[:4],
			},
			wantArrayLength: 4,
		},
		{
			name: "Test 2: Getting values from set of jobs of length 2",
			args: args{
				jobs: jobsArray[:2],
			},
			wantArrayLength: 2,
		},
		{
			name: "Test 3: Getting values from whole set of jobs of length 9 but job at index 5 reports an error",
			args: args{
				jobs: jobsArray,
			},
			wantArrayLength: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UtilsInterface = &UtilsStruct{}
			commitParams := &types.CommitParams{
				LocalCache: cache.NewLocalCache(),
				HttpClient: httpClient,
			}

			gotDataArray, gotWeightArray := UtilsInterface.GetDataToCommitFromJobs(tt.args.jobs, commitParams)
			if len(gotDataArray) != tt.wantArrayLength || len(gotWeightArray) != tt.wantArrayLength {
				t.Errorf("GetDataToCommitFromJobs() got = %v, want %v", gotDataArray, tt.wantArrayLength)
			}
			fmt.Println("Got Data Array: ", gotDataArray)
			fmt.Println("Got WeightArray: ", gotWeightArray)
		})
	}
}

func TestGetDataToCommitFromJob(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        2,
			MaxIdleConnsPerHost: 1,
		},
	}

	job := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_kraken", Selector: "result.XETHZUSD.c[0]",
		Url: `{"type": "GET","url": "https://api.kraken.com/0/public/Ticker?pair=ETHUSD","body": {},"header": {}}`,
	}

	job1 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_sample", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd/apiKey=${SAMPLE_API_KEY_NEW}",
	}

	postJobUniswapV3 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_sample", Selector: "result",
		Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","method":"eth_call","params":[{"to":"0xb27308f9f90d607463bb33ea1bebb41c27ce5ab6","data":"0xf7729d43000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000"},"latest"],"id":5},"header": {"content-type": "application/json"}, "returnType": "hex"}`,
	}

	postJobUniswapV2 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 6, Name: "ethusd_sample", Selector: "result",
		Url: `{"type": "POST","url": "https://rpc.ankr.com/eth","body": {"jsonrpc":"2.0","id":7269270904970082,"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","data":"0xd06ca61f0000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000050de6856358cc35f3a9a57eaaa34bd4cb707d2cd0000000000000000000000008e870d67f660d95d5be530380d0ec0bd388289e1","to":"0x7a250d5630b4cf539739df2c5dacb4c659f2488d"},"latest"]},"header": {"content-type": "application/json"}, "returnType": "hexArray[1]"}`,
	}

	arrayOfObjectsJob := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_bitfinex", Selector: "last_price",
		Url: "https://api.bitfinex.com/v1/pubticker/ethusd",
	}

	arrayOfArraysJob := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_bitfinex_v2", Selector: "last_price",
		Url: "https://api-pub.bitfinex.com/v2/tickers?symbols=tXDCUSD",
	}

	invalidDataSourceStructJob := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_sample", Selector: "result",
		Url: `{"type": true,"url1": {}}`,
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
				job: postJobUniswapV3,
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
		{
			name: "Test 5: When there is an error in unmarshalling invalid dataSourceStruct",
			args: args{
				job: invalidDataSourceStructJob,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When GetDataToCommitFromJob() executes successfully for a POST uniswap v2 Job",
			args: args{
				job: postJobUniswapV2,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test 7: When GetDataToCommitFromJob() executes successfully for job returning response of type array of objects",
			args: args{
				job: arrayOfObjectsJob,
			},
			wantErr: false,
		},
		{
			name: "Test 8: When GetDataToCommitFromJob() fails for job returning response of type arrays of arrays as element in array is not a json object",
			args: args{
				job: arrayOfArraysJob,
			},
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

			commitParams := &types.CommitParams{
				LocalCache: cache.NewLocalCache(),
				HttpClient: httpClient,
			}

			data, err := utils.GetDataToCommitFromJob(tt.args.job, commitParams)
			fmt.Println("JOB returns data: ", data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataToCommitFromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetJobs(t *testing.T) {
	jobsArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
	}

	type args struct {
		numJobs      uint16
		numJobsErr   error
		assetType    uint8
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

			utilsMock.On("GetNumJobs", mock.Anything).Return(tt.args.numJobs, tt.args.numJobsErr)
			utilsMock.On("GetActiveJob", mock.Anything, mock.Anything).Return(tt.args.activeJob, tt.args.activeJobErr)

			got, err := utils.GetJobs(rpcParameters)
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

			got, err := utils.GetNumActiveCollections(rpcParameters)
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
			assetManagerMock.On("GetNumCollections", mock.Anything, mock.Anything).Return(tt.args.numOfAssets, tt.args.numOfAssetsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumCollections(rpcParameters)
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
	ethCollection := bindings.StructsCollection{
		Active: true, Id: 7, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1}, Name: "ethCollection",
	}

	ethCollection1 := bindings.StructsCollection{
		Active: true, Id: 7, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2}, Name: "ethCollection",
	}

	job1 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 0,
		Power: 2, Name: "ethusd_kucoin", Selector: "last",
		Url: "http://kucoin.com/eth",
	}

	job2 := bindings.StructsJob{Id: 2, SelectorType: 0, Weight: 2,
		Power: 3, Name: "ethusd_coinbase", Selector: "eth2",
		Url: "http://api.coinbase.com/eth",
	}

	type args struct {
		collection    bindings.StructsCollection
		dataString    string
		addJobToCache bool
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
				collection:    ethCollection,
				dataString:    jsonDataString,
				addJobToCache: true,
			},
			want: []bindings.StructsJob{
				{
					Id:           1,
					SelectorType: 0,
					Name:         "ethusd_kucoin",
					Url:          "http://kucoin.com/eth1",
					Selector:     "eth1",
					Power:        2,
					Weight:       2,
				},
			},
			wantOverrideJobIds: []uint16{1},
		},
		{
			name: "Test 2: When officialJobs for collection is not present in assets.json",
			args: args{
				collection: ethCollection,
				dataString: "",
			},
			want:               nil,
			wantOverrideJobIds: nil,
		},
		{
			name: "Test 3: When there is an error from GetActiveJob()",
			args: args{
				collection:    ethCollection,
				dataString:    jsonDataString,
				addJobToCache: false,
			},
			want:               nil,
			wantOverrideJobIds: nil,
		},
		{
			name: "Test 4: When multiple jobIds are needed to be overridden from official jobs",
			args: args{
				collection:    ethCollection1,
				dataString:    jsonDataString,
				addJobToCache: true,
			},
			want: []bindings.StructsJob{
				{
					Id:       1,
					Name:     "ethusd_kucoin",
					Url:      "http://kucoin.com/eth1",
					Selector: "eth1",
					Power:    2,
					Weight:   2,
				},
				{
					Id:       2,
					Name:     "ethusd_coinbase",
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
			commitParams := &types.CommitParams{
				JobsCache: cache.NewJobsCache(),
			}
			if tt.args.addJobToCache {
				commitParams.JobsCache.Jobs[job1.Id] = job1
				commitParams.JobsCache.Jobs[job2.Id] = job2
			}

			utils := &UtilsStruct{}

			gotJobs, gotOverrideJobIds := utils.HandleOfficialJobsFromJSONFile(tt.args.collection, tt.args.dataString, commitParams)
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
			utilsMock.On("Aggregate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.collectionData, tt.args.aggregationErr)

			got, err := utils.GetAggregatedDataOfCollection(rpcParameters, collectionId, epoch, &types.CommitParams{HttpClient: &http.Client{}})
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

			got, got1, err := utils.GetAssignedCollections(rpcParameters, numActiveCollections, seed)
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
	var collectionId uint16
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

			got, err := utils.GetLeafIdOfACollection(rpcParameters, collectionId)
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
	var medianIndex uint16
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

			got, err := utils.GetCollectionIdFromIndex(rpcParameters, medianIndex)
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
	var leafId uint16
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
			assetManagerMock.On("GetCollectionIdFromLeafId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.collectionId, tt.args.collectionIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetCollectionIdFromLeafId(rpcParameters, leafId)
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

func TestIsJSONCompatible(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid JSON object",
			input:    `{"key": "value"}`,
			expected: true,
		},
		{
			name:     "Valid JSON array",
			input:    `["item1", "item2"]`,
			expected: true,
		},
		{
			name:     "Valid JSON number",
			input:    "1234",
			expected: true,
		},
		{
			name:     "Invalid JSON missing closing brace",
			input:    `{"key": "value"`,
			expected: false,
		},
		{
			name:     "Invalid JSON missing key",
			input:    `{: "value"}`,
			expected: false,
		},
		{
			name:     "direct URL passed",
			input:    "https://api.gemini.com/v1/pubticker/ethusd",
			expected: false,
		},
		{
			name:     "URL passed as a dataSourceStruct",
			input:    `{"type": "GET","url": "https://api.kraken.com/0/public/Ticker?pair=ETHUSD","body": {},"header": {}}`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isJSONCompatible(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for input: %s", tc.expected, result, tc.input)
			}
		})
	}
}
