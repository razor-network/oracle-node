package cmd

import (
	"errors"
	"razor/cmd/mocks"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestSetConfig(t *testing.T) {

	var flagSet *pflag.FlagSet

	type args struct {
		provider              string
		providerErr           error
		gasmultiplier         float32
		gasmultiplierErr      error
		buffer                int32
		bufferErr             error
		waitTime              int32
		waitTimeErr           error
		gasPrice              int32
		gasPriceErr           error
		logLevel              string
		logLevelErr           error
		path                  string
		pathErr               error
		configErr             error
		gasLimitMultiplier    float32
		gasLimitMultiplierErr error
		isFlagPassed          bool
		port                  string
		portErr               error
		certFile              string
		certFileErr           error
		certKey               string
		certKeyErr            error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When values are passed to all flags and setConfig returns no error",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When parameters are set to default values and setConfig returns no error",
			args: args{
				provider:              "",
				gasmultiplier:         -1,
				buffer:                0,
				waitTime:              -1,
				gasPrice:              -1,
				logLevel:              "",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting provider",
			args: args{
				providerErr:           errors.New("provider error"),
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 4: When there is an error in getting gasmultiplier",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplierErr:      errors.New("gasmultiplier error"),
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("gasmultiplier error"),
		},
		{
			name: "Test 5: When there is an error in getting buffer",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				bufferErr:             errors.New("buffer error"),
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("buffer error"),
		},
		{
			name: "Test 6: When there is an error in getting waitTime",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTimeErr:           errors.New("waitTime error"),
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("waitTime error"),
		},
		{
			name: "Test 7: When there is an error in getting gasprice",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPriceErr:           errors.New("gasprice error"),
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("gasprice error"),
		},
		{
			name: "Test 8: When there is an error in getting logLevel",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevelErr:           errors.New("logLevel error"),
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("logLevel error"),
		},
		{
			name: "Test 9: When there is an error in getting path",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				pathErr:               errors.New("path error"),
				configErr:             nil,
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 10: When there is an error in writing config",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             errors.New("writing config error"),
				gasLimitMultiplier:    10,
				gasLimitMultiplierErr: nil,
			},
			wantErr: errors.New("writing config error"),
		},
		{
			name: "Test 11: When only one of the flags is passed",
			args: args{
				gasmultiplier: 2,
				path:          "/home/config",
				configErr:     nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 12: When there is an error in getting gas limit",
			args: args{
				provider:              "http://127.0.0.1",
				gasmultiplier:         2,
				buffer:                20,
				waitTime:              2,
				gasPrice:              1,
				logLevel:              "debug",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    -1,
				gasLimitMultiplierErr: errors.New("gasLimitMultiplier error"),
			},
			wantErr: errors.New("gasLimitMultiplier error"),
		},
		{
			name: "Test 13: When default nil values are passed",
			args: args{
				provider:              "",
				gasmultiplier:         -1,
				buffer:                0,
				waitTime:              -1,
				gasPrice:              -1,
				logLevel:              "",
				path:                  "/home/config",
				configErr:             nil,
				gasLimitMultiplier:    -1,
				gasLimitMultiplierErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 14: When exposeMetrics flag is passed",
			args: args{
				isFlagPassed: true,
				port:         "",
				configErr:    errors.New("config error"),
			},
			wantErr: errors.New("config error"),
		},
		{
			name: "Test 15: When there is an error in getting port",
			args: args{
				isFlagPassed: true,
				portErr:      errors.New("error in getting port"),
			},
			wantErr: errors.New("error in getting port"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			viperMock := new(mocks.ViperInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			flagSetUtils = flagSetUtilsMock
			viperUtils = viperMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			flagSetUtilsMock.On("GetStringProvider", flagSet).Return(tt.args.provider, tt.args.providerErr)
			flagSetUtilsMock.On("GetFloat32GasMultiplier", flagSet).Return(tt.args.gasmultiplier, tt.args.gasmultiplierErr)
			flagSetUtilsMock.On("GetInt32Buffer", flagSet).Return(tt.args.buffer, tt.args.bufferErr)
			flagSetUtilsMock.On("GetInt32Wait", flagSet).Return(tt.args.waitTime, tt.args.waitTimeErr)
			flagSetUtilsMock.On("GetInt32GasPrice", flagSet).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			flagSetUtilsMock.On("GetStringLogLevel", flagSet).Return(tt.args.logLevel, tt.args.logLevelErr)
			flagSetUtilsMock.On("GetFloat32GasLimit", flagSet).Return(tt.args.gasLimitMultiplier, tt.args.gasLimitMultiplierErr)
			flagSetUtilsMock.On("GetStringExposeMetrics", flagSet).Return(tt.args.port, tt.args.portErr)
			flagSetUtilsMock.On("GetStringCertFile", flagSet).Return(tt.args.certFile, tt.args.certFileErr)
			flagSetUtilsMock.On("GetStringCertKey", flagSet).Return(tt.args.certKey, tt.args.certKeyErr)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
			utilsMock.On("GetConfigFilePath").Return(tt.args.path, tt.args.pathErr)
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
