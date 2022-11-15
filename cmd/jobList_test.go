package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"razor/core/types"
	"razor/pkg/bindings"
	utilsPkgMocks "razor/utils/mocks"
	"testing"
)

func TestGetJobList(t *testing.T) {
	var client *ethclient.Client
	type fields struct {
		razorUtils Utils
	}
	testUtils := fields{
		razorUtils: Utils{},
	}

	jobListArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
		{Id: 2, SelectorType: 1, Weight: 100,
			Power: 2, Name: "btcusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/btcusd",
		},
	}

	type args struct {
		client     *ethclient.Client
		jobList    []bindings.StructsJob
		jobListErr error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Test 1: When jobList executes properly",
			fields: testUtils,
			args: args{
				client:     client,
				jobList:    jobListArray,
				jobListErr: nil,
			},

			wantErr: nil,
		},
		{
			name:   "Test 2: When there is a error fetching job list ",
			fields: testUtils,
			args: args{
				client:     client,
				jobListErr: errors.New("error in fetching job list"),
			},
			wantErr: errors.New("error in fetching job list"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(utilsPkgMocks.Utils)
			razorUtils = utilsMock

			utilsMock.On("GetJobs", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.jobList, tt.args.jobListErr)
			utils := &UtilsStruct{}

			err := utils.GetJobList(tt.args.client)

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for jobList function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for jobList function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteJobList(t *testing.T) {
	var config types.Configurations
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	type args struct {
		config     types.Configurations
		configErr  error
		jobListErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1:  When ExecuteJobList function executes successfully",
			args: args{
				config:     config,
				configErr:  nil,
				jobListErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2:  When there is an error in getting config",
			args: args{
				config:     config,
				configErr:  errors.New("config error"),
				jobListErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3:  When there is an error in getting GetJobList function",
			args: args{
				config:     config,
				configErr:  nil,
				jobListErr: errors.New("jobList error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(utilsPkgMocks.Utils)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			fileUtilsMock := new(utilsPkgMocks.FileUtils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			fileUtils = fileUtilsMock

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("GetJobList", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.jobListErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteJobList(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteJobList function didn't execute as expected")
			}

		})
	}
}
