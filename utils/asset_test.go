package utils

import (
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestAggregate(t *testing.T) {
	var client *ethclient.Client
	var previousEpoch uint32

	job := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	collection := bindings.StructsCollection{Active: true, Id: 4, AssetIndex: 1, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2, 3}, Name: "ethCollectionMean",
	}

	type args struct {
		collection            bindings.StructsCollection
		activeJob             bindings.StructsJob
		activeJobErr          error
		dataToCommit          []*big.Int
		dataToCommitErr       error
		weight                []uint8
		prevCommitmentData    uint32
		prevCommitmentDataErr error
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
				prevCommitmentData: 1,
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
				prevCommitmentData: 1,
			},
			want:    big.NewInt(2),
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting dataToCommit",
			args: args{
				collection:         collection,
				activeJob:          job,
				dataToCommitErr:    errors.New("dataToCommit error"),
				weight:             []uint8{100},
				prevCommitmentData: 1,
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

			utilsMock.On("GetActiveJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.activeJob, tt.args.activeJobErr)
			utilsMock.On("GetDataToCommitFromJobs", mock.Anything).Return(tt.args.dataToCommit, tt.args.weight, tt.args.dataToCommitErr)
			utilsMock.On("FetchPreviousValue", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint16")).Return(tt.args.prevCommitmentData, tt.args.prevCommitmentDataErr)

			got, err := utils.Aggregate(client, previousEpoch, tt.args.collection)
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

func TestGetActiveAssetIds(t *testing.T) {
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
			name: "Test 1: When GetActiveAssetIds() executes successfully",
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("GetActiveCollections", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.activeAssetIds, tt.args.activeAssetIdsErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetActiveAssetIds(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveAssetIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveAssetIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveAssetsData(t *testing.T) {
	var client *ethclient.Client
	var epoch uint32

	collection := bindings.StructsCollection{Active: true, Id: 2, AssetIndex: 1, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2}, Name: "ethCollectionMean",
	}

	type args struct {
		numAssets           uint16
		numAssetsErr        error
		assetType           uint8
		assetTypeErr        error
		activeCollection    bindings.StructsCollection
		activeCollectionErr error
		aggregation         *big.Int
		aggregationErr      error
	}
	tests := []struct {
		name    string
		args    args
		want    []*big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetActiveAssetsData() executes successfully",
			args: args{
				numAssets:        2,
				assetType:        2,
				activeCollection: collection,
				aggregation:      big.NewInt(2),
			},
			want:    []*big.Int{big.NewInt(2), big.NewInt(2)},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numAssets",
			args: args{
				numAssetsErr:     errors.New("numAssets error"),
				assetType:        2,
				activeCollection: collection,
				aggregation:      big.NewInt(2),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting assetType",
			args: args{
				numAssets:        2,
				assetTypeErr:     errors.New("assetType error"),
				activeCollection: collection,
				aggregation:      big.NewInt(2),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting activeCollection",
			args: args{
				numAssets:           2,
				assetType:           2,
				activeCollectionErr: errors.New("activeCollection error"),
				aggregation:         big.NewInt(2),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 5: When there is an error in getting aggregation",
			args: args{
				numAssets:        2,
				assetType:        2,
				activeCollection: collection,
				aggregationErr:   errors.New("aggregation error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When there is an inactive collection",
			args: args{
				numAssets:           2,
				assetType:           2,
				activeCollectionErr: errors.New("collection inactive"),
				aggregation:         big.NewInt(2),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetNumAssets", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numAssets, tt.args.numAssetsErr)
			utilsMock.On("GetAssetType", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.assetType, tt.args.assetTypeErr)
			utilsMock.On("GetActiveCollection", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.activeCollection, tt.args.activeCollectionErr)
			utilsMock.On("Aggregate", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.aggregation, tt.args.aggregationErr)

			got, err := utils.GetActiveAssetsData(client, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveAssetsData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveAssetsData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveCollection(t *testing.T) {
	var client *ethclient.Client
	var collectionId uint16

	collectionEth := bindings.StructsCollection{Active: true, Id: 2, AssetIndex: 1, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 2}, Name: "ethCollectionMean",
	}

	collectionEthInactive := bindings.StructsCollection{Active: false, Id: 2, AssetIndex: 1, Power: 2,
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("Jobs", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.AnythingOfType("uint16")).Return(tt.args.job, tt.args.jobErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

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

func TestGetAssetType(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var assetId uint16

	job := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	collection := bindings.StructsCollection{Active: true, Id: 2, AssetIndex: 1, Power: 2,
		AggregationMethod: 2, JobIDs: []uint16{1, 3, 4}, Name: "ethCollectionMean",
	}

	assetJob := types.Asset{
		Job: job,
	}

	assetCollection := types.Asset{
		Collection: collection,
	}

	type args struct {
		activeAssets    types.Asset
		activeAssetsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "Test 1: When GetAssetType() executes successfully and assetType is job",
			args: args{
				activeAssets: assetJob,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When asssetType is collection",
			args: args{
				activeAssets: assetCollection,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting activeAssets",
			args: args{
				activeAssetsErr: errors.New("activeAssets error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("GetAsset", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.AnythingOfType("uint16")).Return(tt.args.activeAssets, tt.args.activeAssetsErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetAssetType(client, assetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAssetType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAssetType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCollection(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var collectionId uint16

	type args struct {
		asset    types.Asset
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
				asset: types.Asset{},
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("GetAsset", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.AnythingOfType("uint16")).Return(tt.args.asset, tt.args.assetErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

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

func TestGetCollections(t *testing.T) {
	var client *ethclient.Client

	collectionListArray := []bindings.StructsCollection{
		{Active: true, Id: 7, AssetIndex: 1, Power: 2,
			AggregationMethod: 2, JobIDs: []uint16{1, 2, 3}, Name: "ethCollectionMean",
		},
	}

	type args struct {
		numAssets     uint16
		numAssetsErr  error
		assetType     uint8
		assetTypeErr  error
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
			name: "Test 1: When GetCollections() executes successfully",
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
			name: "Test 3: When there is an error in getting assetType",
			args: args{
				numAssets:    1,
				assetTypeErr: errors.New("assetType error"),
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
		{
			name: "Test 4: When there is a different assetType",
			args: args{
				numAssets:  1,
				assetType:  1,
				collection: collectionListArray[0],
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetNumAssets", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numAssets, tt.args.numAssetsErr)
			utilsMock.On("GetAssetType", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.assetType, tt.args.assetTypeErr)
			utilsMock.On("GetCollection", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.collection, tt.args.collectionErr)

			got, err := utils.GetCollections(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataToCommitFromJobs(t *testing.T) {
	jobsArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		}, {Id: 2, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
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
				overrideJobData: map[string]*types.StructsJob{"1": &types.StructsJob{
					Id: 2, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini", Selector: "last",
					Url: "https://api.gemini.com/v1/pubticker/ethusd",
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
			want:    nil,
			wantErr: true,
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("GetJobFilePath").Return(tt.args.jobPath, tt.args.jobPathErr)
			utilsMock.On("ReadJSONData", mock.AnythingOfType("string")).Return(tt.args.overrideJobData, tt.args.overrideJobDataErr)
			utilsMock.On("GetDataToCommitFromJob", mock.Anything).Return(tt.args.dataToAppend, tt.args.dataToAppendErr)

			got, _, err := utils.GetDataToCommitFromJobs(jobsArray)
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
	job := bindings.StructsJob{Id: 1, SelectorType: 1, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	job2 := bindings.StructsJob{Id: 1, SelectorType: 0, Weight: 100,
		Power: 2, Name: "ethusd_gemini", Selector: "last",
		Url: "https://api.gemini.com/v1/pubticker/ethusd",
	}

	response := []byte(`{
  			"userId": 1,
  			"id": 1,
			"title": "delectus aut autem",
  			"completed": false
	}`)

	type args struct {
		job           bindings.StructsJob
		response      []byte
		responseErr   error
		parsedData    interface{}
		parsedDataErr error
		dataPoint     string
		dataPointErr  error
		datum         *big.Float
		datumErr      error
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
				job:        job,
				response:   response,
				parsedData: "abc",
				dataPoint:  "1",
				datum:      big.NewFloat(0.1),
			},
			want:    big.NewInt(10),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting response",
			args: args{
				job:         job,
				responseErr: errors.New("response error"),
				parsedData:  "abc",
				dataPoint:   "1",
				datum:       big.NewFloat(0.1),
			},
			want:    big.NewInt(10),
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting parsedData",
			args: args{
				job:           job,
				response:      response,
				parsedDataErr: errors.New("parsedData error"),
				dataPoint:     "1",
				datum:         big.NewFloat(0.1),
			},
			want:    big.NewInt(10),
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting dataPoint",
			args: args{
				job:          job,
				response:     response,
				parsedData:   "abc",
				dataPointErr: errors.New("dataPoint error"),
				datum:        big.NewFloat(0.1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test 5: When there is an error in getting datum",
			args: args{
				job:        job,
				response:   response,
				parsedData: "abc",
				dataPoint:  "1",
				datumErr:   errors.New("datum error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When there is an Unmarshal error",
			args: args{
				job:        job2,
				response:   []byte(""),
				parsedData: "abc",
				dataPoint:  "1",
				datum:      big.NewFloat(0.1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 7: When there is an error in getting response and selector type is 0",
			args: args{
				job:         job2,
				responseErr: errors.New("API error"),
				parsedData:  "abc",
				dataPoint:   "1",
				datum:       big.NewFloat(0.1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 8: When there is an error in getting parsedData and selector type is 0",
			args: args{
				job:           job2,
				response:      response,
				parsedDataErr: errors.New("parseData error"),
				dataPoint:     "1",
				datum:         big.NewFloat(0.1),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetDataFromAPI", mock.AnythingOfType("string")).Return(tt.args.response, tt.args.responseErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))
			utilsMock.On("GetDataFromJSON", mock.Anything, mock.AnythingOfType("string")).Return(tt.args.parsedData, tt.args.parsedDataErr)
			utilsMock.On("GetDataFromXHTML", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.args.dataPoint, tt.args.dataPointErr)
			optionsMock.On("ConvertToNumber", mock.Anything).Return(tt.args.datum, tt.args.datumErr)

			got, err := utils.GetDataToCommitFromJob(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataToCommitFromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataToCommitFromJob() got = %v, want %v", got, tt.want)
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
		numAssets    uint16
		numAssetsErr error
		assetType    uint8
		assetTypeErr error
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
				numAssets: 1,
				assetType: 1,
				activeJob: jobsArray[0],
			},
			want:    jobsArray,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numAssets",
			args: args{
				numAssetsErr: errors.New("numAssets error"),
				assetType:    1,
				activeJob:    jobsArray[0],
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting assetType",
			args: args{
				numAssets:    1,
				assetTypeErr: errors.New("assetType error"),
				activeJob:    jobsArray[0],
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting activeJob",
			args: args{
				numAssets:    1,
				assetType:    1,
				activeJobErr: errors.New("activeJob error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is a different assetType",
			args: args{
				numAssets: 1,
				assetType: 2,
				activeJob: jobsArray[0],
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetNumAssets", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numAssets, tt.args.numAssetsErr)
			utilsMock.On("GetAssetType", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint16")).Return(tt.args.assetType, tt.args.assetTypeErr)
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

func TestGetNumActiveAssets(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		numOfActiveAssets    *big.Int
		numOfActiveAssetsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetNumActiveAssets() executes successfully",
			args: args{
				numOfActiveAssets: big.NewInt(5),
			},
			want:    big.NewInt(5),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numOfActiveAssets",
			args: args{
				numOfActiveAssetsErr: errors.New("numOfActiveAssets error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("GetNumActiveCollections", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.numOfActiveAssets, tt.args.numOfActiveAssetsErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumActiveAssets(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumActiveAssets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumActiveAssets() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumAssets(t *testing.T) {
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
			name: "Test 1: When GetNumAssets() executes successfully",
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			optionsMock.On("GetNumAssets", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.numOfAssets, tt.args.numOfAssetsErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumAssets(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumAssets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumAssets() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAssetManagerWithOpts(t *testing.T) {

	var callOpts bind.CallOpts
	var assetManager *bindings.AssetManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetAssetManager", mock.AnythingOfType("*ethclient.Client")).Return(assetManager)

	gotAssetManager, gotCallOpts := utils.GetAssetManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetAssetManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotAssetManager, assetManager) {
		t.Errorf("GetAssetkManagerWithOpts() got assetManager = %v, want %v", gotAssetManager, assetManager)
	}
}
