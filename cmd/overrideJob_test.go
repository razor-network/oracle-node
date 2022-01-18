package cmd

import (
	"errors"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"razor/core/types"
	"testing"
)

func TestOverrideJob(t *testing.T) {
	var job types.StructsJob
	type args struct {
		jobPath         string
		jobPathErr      error
		addJobToJSONErr error
	}
	var tests = []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When OverrideJob function executes successfully",
			args: args{
				jobPath: "/home/local/jobs.json",
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When OverrideJob fails due to path error",
			args: args{
				jobPath:    "/home/local/jobs.json",
				jobPathErr: errors.New("jobPath error"),
			},
			wantErr: errors.New("jobPath error"),
		},
		{
			name: "Test 3: When there is an error in addJobToJSON function",
			args: args{
				jobPath:         "/home/local/jobs.json",
				addJobToJSONErr: errors.New("addJobToJSON error"),
			},
			wantErr: errors.New("addJobToJSON error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)

			razorUtils = utilsMock

			utilsMock.On("GetJobFilePath").Return(tt.args.jobPath, tt.args.jobPathErr)
			utilsMock.On("AddJobToJSON", mock.AnythingOfType("string"), mock.Anything).Return(tt.args.addJobToJSONErr)

			utils := &UtilsStruct{}

			err := utils.OverrideJob(&job)
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

func TestExecuteOverrideJob(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		url             string
		urlErr          error
		selector        string
		selectorErr     error
		power           int8
		powerErr        error
		selectorType    uint8
		selectorTypeErr error
		jobId           uint16
		jobIdErr        error
		OverrideJobErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteOverrideJob function executes successfully",
			args: args{
				url:            "https://api.gemini.com/v1/pubticker/ethusd",
				selector:       "last",
				power:          1,
				selectorType:   1,
				jobId:          1,
				OverrideJobErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting url from flags",
			args: args{
				url:            "",
				urlErr:         errors.New("url error"),
				selector:       "last",
				power:          1,
				selectorType:   1,
				jobId:          1,
				OverrideJobErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting selector from flags",
			args: args{
				url:            "https://api.gemini.com/v1/pubticker/ethusd",
				selector:       "",
				selectorErr:    errors.New("selector error"),
				power:          1,
				selectorType:   1,
				jobId:          1,
				OverrideJobErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting power from flag",
			args: args{
				url:            "https://api.gemini.com/v1/pubticker/ethusd",
				selector:       "last",
				powerErr:       errors.New("power error"),
				selectorType:   1,
				jobId:          1,
				OverrideJobErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting selectorType from flag",
			args: args{
				url:             "https://api.gemini.com/v1/pubticker/ethusd",
				selector:        "last",
				power:           1,
				selectorTypeErr: errors.New("selectorType error"),
				jobId:           1,
				OverrideJobErr:  nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting jobId from flag",
			args: args{
				url:            "https://api.gemini.com/v1/pubticker/ethusd",
				selector:       "last",
				power:          1,
				selectorType:   1,
				jobIdErr:       errors.New("jobId error"),
				OverrideJobErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from OverrideJob function",
			args: args{
				url:            "https://api.gemini.com/v1/pubticker/ethusd",
				selector:       "last",
				power:          1,
				selectorType:   1,
				jobId:          1,
				OverrideJobErr: errors.New("overrideJob error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagsetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			flagSetUtils = flagsetUtilsMock
			cmdUtils = cmdUtilsMock

			flagsetUtilsMock.On("GetStringUrl", flagSet).Return(tt.args.url, tt.args.urlErr)
			flagsetUtilsMock.On("GetStringSelector", flagSet).Return(tt.args.selector, tt.args.selectorErr)
			flagsetUtilsMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			flagsetUtilsMock.On("GetUint8SelectorType", flagSet).Return(tt.args.selectorType, tt.args.selectorTypeErr)
			flagsetUtilsMock.On("GetUint16JobId", flagSet).Return(tt.args.jobId, tt.args.jobIdErr)
			cmdUtilsMock.On("OverrideJob", mock.Anything).Return(tt.args.OverrideJobErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteOverrideJob(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteOverrideJob function didn't execute as expected")
			}

		})
	}
}
