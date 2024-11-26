package cmd

import (
	"errors"
	"razor/core/types"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestSetConfig(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		flagInput                 string
		flagInputErr              error
		isFlagPassed              bool
		path                      string
		pathErr                   error
		isExposeMetricsFlagPassed bool
		configErr                 error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When values are passed to all flags and setConfig returns no error",
			args: args{
				flagInput:                 "http://127.0.0.1",
				path:                      "/home/config",
				isFlagPassed:              true,
				isExposeMetricsFlagPassed: false,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there are no values passed as flag and all config values are default values",
			args: args{
				path:                      "/home/config",
				isFlagPassed:              false,
				isExposeMetricsFlagPassed: false,
			},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in running metrics server",
			args: args{
				flagInput:                 "8080",
				path:                      "/home/config",
				isFlagPassed:              true,
				isExposeMetricsFlagPassed: true,
			},
		},
		{
			name: "Test 4: When there is an error in getting path",
			args: args{
				pathErr: errors.New("path error"),
			},
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 5: When there is an error in writing config",
			args: args{
				path:                      "/home/config",
				isFlagPassed:              false,
				isExposeMetricsFlagPassed: false,
				configErr:                 errors.New("config error"),
			},
			wantErr: errors.New("config error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			setupTestEndpointsEnvironment()

			cmdUtilsMock.On("GetConfigData").Return(types.Configurations{}, nil)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(false)
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			flagSetMock.On("FetchFlagInput", flagSet, mock.Anything, mock.Anything).Return(tt.args.flagInput, tt.args.flagInputErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagPassed)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isExposeMetricsFlagPassed)
			pathMock.On("GetConfigFilePath").Return(tt.args.path, tt.args.pathErr)
			viperMock.On("ViperWriteConfigAs", mock.AnythingOfType("string")).Return(tt.args.configErr)

			utils := &UtilsStruct{}
			gotErr := utils.SetConfig(flagSet)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for SetConfig function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for SetConfig function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
