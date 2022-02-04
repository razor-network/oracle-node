package cmd

import (
	"errors"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"testing"
)

func TestDeleteOverrideJob(t *testing.T) {
	var jobId uint16
	type args struct {
		jobPath              string
		jobPathErr           error
		deleteJobFromJSONErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When DeleteOverrideJob function executes successfully",
			args: args{
				jobPath: "/home/local/jobs.json",
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When DeleteOverrideJob fails due to path error",
			args: args{
				jobPath:    "/home/local/jobs.json",
				jobPathErr: errors.New("jobPath error"),
			},
			wantErr: errors.New("jobPath error"),
		},
		{
			name: "Test 3: When there is an error in deleteJobFromJSON function",
			args: args{
				jobPath:              "/home/local/jobs.json",
				deleteJobFromJSONErr: errors.New("deleteJobFromJSON error"),
			},
			wantErr: errors.New("deleteJobFromJSON error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)

			razorUtils = utilsMock

			utilsMock.On("GetJobFilePath").Return(tt.args.jobPath, tt.args.jobPathErr)
			utilsMock.On("DeleteJobFromJSON", mock.AnythingOfType("string"), mock.Anything).Return(tt.args.deleteJobFromJSONErr)

			utils := &UtilsStruct{}

			err := utils.DeleteOverrideJob(jobId)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for deleteOverride function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for deleteOverride function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteDeleteOverrideJob(t *testing.T) {
	var flagSet *pflag.FlagSet
	type args struct {
		jobId                uint16
		jobIdErr             error
		DeleteOverrideJobErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteDeleteOverrideJob function executes successfully",
			args: args{
				jobId:                1,
				DeleteOverrideJobErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting jobId from flag",
			args: args{
				jobIdErr:             errors.New("jobId error"),
				DeleteOverrideJobErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error from DeleteOverrideJob function",
			args: args{
				jobId:                1,
				DeleteOverrideJobErr: errors.New("overrideJob error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			flagSetUtils = flagSetUtilsMock
			cmdUtils = cmdUtilsMock

			flagSetUtilsMock.On("GetUint16JobId", flagSet).Return(tt.args.jobId, tt.args.jobIdErr)
			cmdUtilsMock.On("DeleteOverrideJob", mock.Anything).Return(tt.args.DeleteOverrideJobErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteDeleteOverrideJob(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteDeleteOverrideJob function didn't execute as expected")
			}

		})
	}
}
