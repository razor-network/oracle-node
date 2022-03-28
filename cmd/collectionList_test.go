package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"
)

func TestGetCollectionList(t *testing.T) {
	var client *ethclient.Client
	type fields struct {
		razorUtils Utils
	}
	testUtils := fields{
		razorUtils: Utils{},
	}

	collectionListArray := []bindings.StructsCollection{
		{Active: true,
			Id:                7,
			Power:             2,
			AggregationMethod: 2,
			JobIDs:            []uint16{1, 2, 3},
			Name:              "ethCollectionMean",
		},
		{Active: true,
			Id:                8,
			Power:             2,
			AggregationMethod: 2,
			JobIDs:            []uint16{4, 5, 6},
			Name:              "btcCollectionMean",
		},
	}

	type args struct {
		client            *ethclient.Client
		collectionList    []bindings.StructsCollection
		collectionListErr error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Test 1: When collectionList executes properly",
			fields: testUtils,
			args: args{
				client:            client,
				collectionList:    collectionListArray,
				collectionListErr: nil,
			},

			wantErr: nil,
		},
		{
			name:   "Test 2: When there is a error fetching collection list ",
			fields: testUtils,
			args: args{
				client:            client,
				collectionListErr: errors.New("error in fetching collection list"),
			},
			wantErr: errors.New("error in fetching collection list"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			razorUtils = utilsMock

			utilsMock.On("GetCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.collectionList, tt.args.collectionListErr)
			utils := &UtilsStruct{}

			err := utils.GetCollectionList(tt.args.client)

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for collectionList function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for collectionList function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteCollectionList(t *testing.T) {
	var config types.Configurations
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	type args struct {
		config            types.Configurations
		configErr         error
		collectionListErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1:  When ExecuteCollectionList function executes successfully",
			args: args{
				config:            config,
				configErr:         nil,
				collectionListErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2:  When there is an error in getting config",
			args: args{
				config:            config,
				configErr:         errors.New("config error"),
				collectionListErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3:  When there is an error in getting GetCollectionList function",
			args: args{
				config:            config,
				configErr:         nil,
				collectionListErr: errors.New("collectionList error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("GetCollectionList", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.collectionListErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteCollectionList(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteCollectionList function didn't execute as expected")
			}

		})
	}
}
