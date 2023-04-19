package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestSetConfig(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		provider              string
		providerErr           error
		alternateProvider     string
		alternateProviderErr  error
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
		gasLimitOverride      uint64
		gasLimitOverrideErr   error
		rpcTimeout            int64
		rpcTimeoutErr         error
		httpTimeout           int64
		httpTimeoutErr        error
		isFlagPassed          bool
		port                  string
		portErr               error
		certFile              string
		certFileErr           error
		certKey               string
		certKeyErr            error
		logFileMaxSize        int
		logFileMaxSizeErr     error
		logFileMaxBackups     int
		logFileMaxBackupsErr  error
		logFileMaxAge         int
		logFileMaxAgeErr      error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When values are passed to all flags and setConfig returns no error",
			args: args{
				provider:           "http://127.0.0.1",
				alternateProvider:  "http://127.0.0.1:8545",
				gasmultiplier:      2,
				buffer:             20,
				waitTime:           2,
				gasPrice:           1,
				logLevel:           "debug",
				path:               "/home/config",
				gasLimitMultiplier: 10,
				rpcTimeout:         10,
				httpTimeout:        20,
				logFileMaxSize:     6,
				logFileMaxBackups:  11,
				logFileMaxAge:      31,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When parameters are set to default values and setConfig returns no error",
			args: args{
				provider:           "",
				gasmultiplier:      -1,
				buffer:             0,
				waitTime:           -1,
				gasPrice:           -1,
				logLevel:           "",
				path:               "/home/config",
				gasLimitMultiplier: 10,
				rpcTimeout:         0,
				httpTimeout:        0,
			},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting provider",
			args: args{
				providerErr: errors.New("provider error"),
			},
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 4: When there is an error in getting gasmultiplier",
			args: args{
				provider:         "http://127.0.0.1",
				gasmultiplierErr: errors.New("gasmultiplier error"),
			},
			wantErr: errors.New("gasmultiplier error"),
		},
		{
			name: "Test 5: When there is an error in getting buffer",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				bufferErr:     errors.New("buffer error"),
			},
			wantErr: errors.New("buffer error"),
		},
		{
			name: "Test 6: When there is an error in getting waitTime",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTimeErr:   errors.New("waitTime error"),
			},
			wantErr: errors.New("waitTime error"),
		},
		{
			name: "Test 7: When there is an error in getting gasprice",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPriceErr:   errors.New("gasprice error"),
			},
			wantErr: errors.New("gasprice error"),
		},
		{
			name: "Test 8: When there is an error in getting logLevel",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevelErr:   errors.New("logLevel error"),
			},
			wantErr: errors.New("logLevel error"),
		},
		{
			name: "Test 9: When there is an error in getting path",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				pathErr:       errors.New("path error"),
			},
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 10: When there is an error in writing config",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     errors.New("writing config error"),
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
				gasLimitMultiplier:    -1,
				gasLimitMultiplierErr: errors.New("gasLimitMultiplier error"),
			},
			wantErr: errors.New("gasLimitMultiplier error"),
		},
		{
			name: "Test 13: When default nil values are passed",
			args: args{
				provider:           "",
				gasmultiplier:      -1,
				buffer:             0,
				waitTime:           -1,
				gasPrice:           -1,
				logLevel:           "",
				rpcTimeout:         0,
				httpTimeout:        0,
				path:               "/home/config",
				gasLimitMultiplier: -1,
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
		{
			name: "Test 16: When there is an error in getting RPC timeout",
			args: args{
				provider:           "http://127.0.0.1",
				gasmultiplier:      2,
				buffer:             20,
				waitTime:           2,
				gasPrice:           1,
				logLevel:           "debug",
				path:               "/home/config",
				gasLimitMultiplier: -1,
				rpcTimeoutErr:      errors.New("rpcTimeout error"),
			},
			wantErr: errors.New("rpcTimeout error"),
		},
		{
			name: "Test 17: When there is an error in getting gas limit to overrride",
			args: args{
				provider:            "http://127.0.0.1",
				gasmultiplier:       2,
				buffer:              20,
				waitTime:            2,
				gasPrice:            1,
				logLevel:            "debug",
				path:                "/home/config",
				gasLimitMultiplier:  -1,
				gasLimitOverrideErr: errors.New("gasLimitOverride error"),
			},
			wantErr: errors.New("gasLimitOverride error"),
		},
		{
			name: "Test 18: When there is an error in getting HTTP timeout",
			args: args{
				provider:           "http://127.0.0.1",
				gasmultiplier:      2,
				buffer:             20,
				waitTime:           2,
				gasPrice:           1,
				logLevel:           "debug",
				path:               "/home/config",
				gasLimitMultiplier: -1,
				rpcTimeout:         10,
				httpTimeoutErr:     errors.New("httpTimeout error"),
			},
			wantErr: errors.New("httpTimeout error"),
		},
		{
			name: "Test 18: When there is an error in getting alternate provider",
			args: args{
				provider:             "http://127.0.0.1",
				alternateProviderErr: errors.New("alternate provider error"),
			},
			wantErr: errors.New("alternate provider error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			flagSetMock.On("GetStringProvider", flagSet).Return(tt.args.provider, tt.args.providerErr)
			flagSetMock.On("GetStringAlternateProvider", flagSet).Return(tt.args.alternateProvider, tt.args.alternateProviderErr)
			flagSetMock.On("GetFloat32GasMultiplier", flagSet).Return(tt.args.gasmultiplier, tt.args.gasmultiplierErr)
			flagSetMock.On("GetInt32Buffer", flagSet).Return(tt.args.buffer, tt.args.bufferErr)
			flagSetMock.On("GetInt32Wait", flagSet).Return(tt.args.waitTime, tt.args.waitTimeErr)
			flagSetMock.On("GetInt32GasPrice", flagSet).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			flagSetMock.On("GetStringLogLevel", flagSet).Return(tt.args.logLevel, tt.args.logLevelErr)
			flagSetMock.On("GetFloat32GasLimit", flagSet).Return(tt.args.gasLimitMultiplier, tt.args.gasLimitMultiplierErr)
			flagSetMock.On("GetUint64GasLimitOverride", flagSet).Return(tt.args.gasLimitOverride, tt.args.gasLimitOverrideErr)
			flagSetMock.On("GetInt64RPCTimeout", flagSet).Return(tt.args.rpcTimeout, tt.args.rpcTimeoutErr)
			flagSetMock.On("GetInt64HTTPTimeout", flagSet).Return(tt.args.httpTimeout, tt.args.httpTimeoutErr)
			flagSetMock.On("GetStringExposeMetrics", flagSet).Return(tt.args.port, tt.args.portErr)
			flagSetMock.On("GetStringCertFile", flagSet).Return(tt.args.certFile, tt.args.certFileErr)
			flagSetMock.On("GetStringCertKey", flagSet).Return(tt.args.certKey, tt.args.certKeyErr)
			flagSetMock.On("GetIntLogFileMaxSize", mock.Anything).Return(tt.args.logFileMaxSize, tt.args.logFileMaxSizeErr)
			flagSetMock.On("GetIntLogFileMaxBackups", mock.Anything).Return(tt.args.logFileMaxBackups, tt.args.logFileMaxBackupsErr)
			flagSetMock.On("GetIntLogFileMaxAge", mock.Anything).Return(tt.args.logFileMaxAge, tt.args.logFileMaxAgeErr)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
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
