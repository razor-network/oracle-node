package cmd

import (
	"errors"
	"razor/cmd/mocks"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestGetConfigData(t *testing.T) {
	config := types.Configurations{
		Provider:           "",
		ChainId:            0,
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
	}

	configData := types.Configurations{
		Provider:           "",
		ChainId:            137,
		GasMultiplier:      1,
		BufferPercent:      20,
		WaitTime:           1,
		LogLevel:           "debug",
		GasLimitMultiplier: 3,
		LogFileMaxSize:     10,
		LogFileMaxBackups:  20,
		LogFileMaxAge:      30,
		Compress:           "true",
	}

	type args struct {
		provider                   string
		providerErr                error
		chainIdString              string
		chainIdStringErr           error
		chainId                    int64
		chainIdErr                 error
		gasMultiplierString        string
		gasMultiplierStringErr     error
		gasMultiplier              float64
		gasMultiplierErr           error
		bufferPercentString        string
		bufferPercentStringErr     error
		bufferPercent              int64
		bufferPercentErr           error
		waitTimeString             string
		waitTimeStringErr          error
		waitTime                   int64
		waitTimeErr                error
		gasPriceString             string
		gasPriceStringErr          error
		gasPrice                   int64
		gasPriceErr                error
		logLevel                   string
		logLevelErr                error
		gasLimitString             string
		gasLimitStringErr          error
		gasLimit                   float64
		gasLimitErr                error
		logFileMaxSizeString       string
		logFileMaxSizeStringErr    error
		logFileMaxSize             int
		logFileMaxSizeErr          error
		logFileMaxBackupsString    string
		logFileMaxBackupsStringErr error
		logFileMaxBackups          int
		logFileMaxBackupsErr       error
		logFileMaxAgeString        string
		logFileMaxAgeStringErr     error
		logFileMaxAge              int
		logFileMaxAgeErr           error
		compress                   string
		compressErr                error
	}
	tests := []struct {
		name    string
		args    args
		want    types.Configurations
		wantErr error
	}{
		{
			name: "Test 1: When GetConfigData function executes successfully",
			args: args{
				provider:                "",
				chainIdString:           "137",
				chainId:                 137,
				gasMultiplierString:     "1",
				gasMultiplier:           1,
				bufferPercentString:     "20",
				bufferPercent:           20,
				waitTimeString:          "1",
				waitTime:                1,
				logLevel:                "debug",
				gasLimitString:          "3",
				gasLimit:                3,
				logFileMaxSizeString:    "10",
				logFileMaxBackupsString: "20",
				logFileMaxAgeString:     "30",
				logFileMaxSize:          10,
				logFileMaxBackups:       20,
				logFileMaxAge:           30,
				compress:                "true",
			},
			want:    configData,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting provider",
			args: args{
				providerErr: errors.New("provider error"),
			},
			want:    config,
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 3: When ChainId is 0",
			args: args{
				chainIdString: "0",
				chainId:       0,
			},
			want:    types.Configurations{ChainId: 0},
			wantErr: nil,
		},
		{
			name: "test 4: When there is an error in getting chainId",
			args: args{
				chainIdStringErr: errors.New("chainId error"),
			},
			want:    config,
			wantErr: errors.New("chainId error"),
		},
		{
			name: "Test 5: When there is an error in getting gasMultiplier",
			args: args{
				gasMultiplierStringErr: errors.New("gasMultiplier error"),
			},
			want:    config,
			wantErr: errors.New("gasMultiplier error"),
		},
		{
			name: "Test 6: When there is an error in getting bufferPercent",
			args: args{
				bufferPercentStringErr: errors.New("bufferPercent error"),
			},
			want:    config,
			wantErr: errors.New("bufferPercent error"),
		},
		{
			name: "Test 7: When there is an error in getting waitTime",
			args: args{
				waitTimeStringErr: errors.New("waitTime error"),
			},
			want:    config,
			wantErr: errors.New("waitTime error"),
		},
		{
			name: "Test 8: When there is an error in getting gasPrice",
			args: args{
				gasPriceStringErr: errors.New("gasPrice error"),
			},
			want:    config,
			wantErr: errors.New("gasPrice error"),
		},
		{
			name: "Test 9: When there is an error in getting logLevel",
			args: args{
				logLevelErr: errors.New("logLevel error"),
			},
			want:    config,
			wantErr: errors.New("logLevel error"),
		},
		{
			name: "Test 10: When there is an error in getting gasLimit",
			args: args{
				gasLimitStringErr: errors.New("gasLimit error"),
			},
			want:    config,
			wantErr: errors.New("gasLimit error"),
		},
		{
			name: "Test 11: When there is an error in getting logFileMaxSize",
			args: args{
				logFileMaxSizeStringErr: errors.New("logFileMaxSize error"),
			},
			want:    config,
			wantErr: errors.New("logFileMaxSize error"),
		},
		{
			name: "Test 12: When there is an error in getting logFileMaxBackups",
			args: args{
				logFileMaxBackupsStringErr: errors.New("logFileMaxBackups error"),
			},
			want:    config,
			wantErr: errors.New("logFileMaxBackups error"),
		},
		{
			name: "Test 13: When there is an error in getting logFileMaxAge",
			args: args{
				logFileMaxAgeStringErr: errors.New("logFileMaxAge error"),
			},
			want:    config,
			wantErr: errors.New("logFileMaxAge error"),
		},
		{
			name: "Test 14: When there is an error in parsing int for chainID",
			args: args{
				chainIdErr: errors.New("error in parsing"),
			},
			want:    config,
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 15: When there is an error in parsing float for gas multiplier",
			args: args{
				gasMultiplierErr: errors.New("error in parsing"),
			},
			want:    config,
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 16: When there is an error in parsing int for buffer",
			args: args{
				chainIdString:    "137",
				chainId:          137,
				bufferPercentErr: errors.New("error in parsing"),
			},
			want:    types.Configurations{ChainId: 137},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 17: When there is an error in parsing int for waitTime",
			args: args{
				chainIdString:       "137",
				chainId:             137,
				bufferPercentString: "20",
				bufferPercent:       20,
				waitTimeErr:         errors.New("error in parsing"),
			},
			want:    types.Configurations{ChainId: 137, BufferPercent: 20},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 18: When there is an error in parsing int for gasPrice",
			args: args{
				chainIdString:       "137",
				chainId:             137,
				bufferPercentString: "20",
				waitTimeString:      "5",
				bufferPercent:       20,
				waitTime:            5,
				gasPriceErr:         errors.New("error in parsing"),
			},
			want:    types.Configurations{ChainId: 137, BufferPercent: 20, WaitTime: 5},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 19: When there is an error in parsing float for gasLimit",
			args: args{
				gasMultiplierString: "1",
				gasMultiplier:       1,
				gasLimitErr:         errors.New("error in parsing"),
			},
			want:    types.Configurations{GasMultiplier: 1},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 20: When there is an error in parsing int for logFileMaxSize",
			args: args{
				logFileMaxSizeErr: errors.New("error in parsing"),
			},
			want:    config,
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 21: When there is an error in parsing int for logFileMaxBackups",
			args: args{
				logFileMaxSizeString: "10",
				logFileMaxSize:       10,
				logFileMaxBackupsErr: errors.New("error in parsing"),
			},
			want:    types.Configurations{LogFileMaxSize: 10},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 22: When there is an error in parsing int for logFileAge",
			args: args{
				logFileMaxSizeString:    "10",
				logFileMaxBackupsString: "20",
				logFileMaxAgeString:     "30",
				logFileMaxSize:          10,
				logFileMaxBackups:       20,
				logFileMaxAgeErr:        errors.New("error in parsing"),
			},
			want:    types.Configurations{LogFileMaxSize: 10, LogFileMaxBackups: 20},
			wantErr: errors.New("error in parsing"),
		},
		{
			name: "Test 23: When there is an error in getting compress",
			args: args{
				compressErr: errors.New("error in getting compress"),
			},
			want:    config,
			wantErr: errors.New("error in getting compress"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			stringMock := new(mocks.StringInterface)
			cmdUtils = cmdUtilsMock
			stringUtils = stringMock

			cmdUtilsMock.On("GetConfig", "provider").Return(tt.args.provider, tt.args.providerErr)
			cmdUtilsMock.On("GetConfig", "chainId").Return(tt.args.chainIdString, tt.args.chainIdStringErr)
			cmdUtilsMock.On("GetConfig", "gasmultiplier").Return(tt.args.gasMultiplierString, tt.args.gasMultiplierStringErr)
			cmdUtilsMock.On("GetConfig", "wait").Return(tt.args.waitTimeString, tt.args.waitTimeStringErr)
			cmdUtilsMock.On("GetConfig", "gasprice").Return(tt.args.gasPriceString, tt.args.gasPriceStringErr)
			cmdUtilsMock.On("GetConfig", "logLevel").Return(tt.args.logLevel, tt.args.logLevelErr)
			cmdUtilsMock.On("GetConfig", "gasLimit").Return(tt.args.gasLimitString, tt.args.gasLimitStringErr)
			cmdUtilsMock.On("GetConfig", "buffer").Return(tt.args.bufferPercentString, tt.args.bufferPercentStringErr)
			cmdUtilsMock.On("GetConfig", "logFileMaxSize").Return(tt.args.logFileMaxSizeString, tt.args.logFileMaxSizeStringErr)
			cmdUtilsMock.On("GetConfig", "logFileMaxBackups").Return(tt.args.logFileMaxBackupsString, tt.args.logFileMaxBackupsStringErr)
			cmdUtilsMock.On("GetConfig", "logFileMaxAge").Return(tt.args.logFileMaxAgeString, tt.args.logFileMaxAgeStringErr)
			cmdUtilsMock.On("GetConfig", "compress").Return(tt.args.compress, tt.args.compressErr)

			stringMock.On("ParseChainId", tt.args.chainIdString).Return(tt.args.chainId, tt.args.chainIdErr)
			stringMock.On("ParseFloat", tt.args.gasMultiplierString).Return(tt.args.gasMultiplier, tt.args.gasMultiplierErr)
			stringMock.On("ParseInt64", tt.args.bufferPercentString).Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			stringMock.On("ParseInt64", tt.args.waitTimeString).Return(tt.args.waitTime, tt.args.waitTimeErr)
			stringMock.On("ParseInt64", tt.args.gasPriceString).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			stringMock.On("ParseFloat", tt.args.gasLimitString).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			stringMock.On("ParseInt", tt.args.logFileMaxSizeString).Return(tt.args.logFileMaxSize, tt.args.logFileMaxSizeErr)
			stringMock.On("ParseInt", tt.args.logFileMaxBackupsString).Return(tt.args.logFileMaxBackups, tt.args.logFileMaxBackupsErr)
			stringMock.On("ParseInt", tt.args.logFileMaxAgeString).Return(tt.args.logFileMaxAge, tt.args.logFileMaxAgeErr)

			utils := &UtilsStruct{}

			got, err := utils.GetConfigData()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigData() got = %v, want %v", got, tt.want)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetConfigData function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetConfigData function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	type args struct {
		configType           string
		provider             string
		providerErr          error
		chainId              string
		chainIdErr           error
		gasMultiplier        string
		gasMultiplierErr     error
		bufferPercent        string
		bufferPercentErr     error
		waitTime             string
		waitTimeErr          error
		gasPrice             string
		gasPriceErr          error
		logLevel             string
		logLevelErr          error
		gasLimit             string
		gasLimitErr          error
		logFileMaxSize       string
		logFileMaxSizeErr    error
		logFileMaxBackups    string
		logFileMaxBackupsErr error
		logFileMaxAge        string
		logFileMaxAgeErr     error
		compress             string
		compressErr          error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: When GetConfig executes successfully for provider",
			args: args{
				configType: "provider",
				provider:   "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			},
			want:    "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			wantErr: false,
		},
		{
			name: "Test 2: When provider has prefix https",
			args: args{
				configType: "provider",
				provider:   "127.0.0.1:8545",
			},
			want:    "127.0.0.1:8545",
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting provider",
			args: args{
				configType:  "provider",
				providerErr: errors.New("provider error"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 4: When provider is nil",
			args: args{
				configType: "provider",
				provider:   "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 5: When GetConfig executes successfully for gasmultiplier",
			args: args{
				configType:    "gasmultiplier",
				gasMultiplier: "2",
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "Test 6: When gasMultiplier is -1",
			args: args{
				configType:    "gasmultiplier",
				gasMultiplier: "-1",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 7: When there is an error in getting gasMultiplier",
			args: args{
				configType:       "gasmultiplier",
				gasMultiplierErr: errors.New("gasMultiplier error"),
			},
			want:    "1",
			wantErr: true,
		},
		{
			name: "Test 8: When GetConfig executes successfully for buffer",
			args: args{
				configType:    "buffer",
				bufferPercent: "20",
			},
			want:    "20",
			wantErr: false,
		},
		{
			name: "Test 9: When bufferPercent is 0",
			args: args{
				configType:    "buffer",
				bufferPercent: "0",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 10: When there is an error in getting bufferPercent",
			args: args{
				configType:       "buffer",
				bufferPercentErr: errors.New("bufferPercent error"),
			},
			want:    "30",
			wantErr: true,
		},
		{
			name: "Test 11: When GetConfig executes successfully for wait",
			args: args{
				configType: "wait",
				waitTime:   "4",
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: "Test 12: When waitTime is -1",
			args: args{
				configType: "wait",
				waitTime:   "-1",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 13: When there is an error in getting waitTime",
			args: args{
				configType:  "wait",
				waitTimeErr: errors.New("waitTime error"),
			},
			want:    "3",
			wantErr: true,
		},
		{
			name: "Test 14: When GetConfig executes successfully for gasPrice",
			args: args{
				configType: "gasprice",
				gasPrice:   "1",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "Test 15: When gasPrice is -1",
			args: args{
				configType: "gasprice",
				gasPrice:   "-1",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 16: When there is an error in getting gasPrice",
			args: args{
				configType:  "gasprice",
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    "0",
			wantErr: true,
		},
		{
			name: "Test 17: When GetConfig executes successfully for logLevel",
			args: args{
				configType: "logLevel",
				logLevel:   "debug",
			},
			want:    "debug",
			wantErr: false,
		},
		{
			name: "Test 18: When logLevel is nil",
			args: args{
				configType: "logLevel",
				logLevel:   "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 19: When there is an error in getting logLevel",
			args: args{
				configType:  "logLevel",
				logLevelErr: errors.New("logLevel error"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 20: When GetConfig executes successfully for gasLimit",
			args: args{
				configType: "gasLimit",
				gasLimit:   "4",
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: "Test 21: When gasLimit is -1",
			args: args{
				configType: "gasLimit",
				gasLimit:   "-1",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 22: When there is an error in getting gasLimit",
			args: args{
				configType:  "gasLimit",
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    "-1",
			wantErr: true,
		},
		{
			name: "Test 23: When configType does not match with any case",
			args: args{
				configType: "abc",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 24: When GetConfig executes successfully for chainId",
			args: args{
				configType: "chainId",
				chainId:    "137",
			},
			want:    "137",
			wantErr: false,
		},
		{
			name: "Test 25: When there is an error in getting chainId",
			args: args{
				configType: "chainId",
				chainIdErr: errors.New("error in getting chainId"),
			},
			want:    "0",
			wantErr: true,
		},
		{
			name: "Test 26: When chainId is 0",
			args: args{
				configType: "chainId",
				chainId:    "0",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 27: When GetConfig executes successfully for logFileMaxSize",
			args: args{
				configType:     "logFileMaxSize",
				logFileMaxSize: "10",
			},
			want:    "10",
			wantErr: false,
		},
		{
			name: "Test 28: When there is an error in getting logFileMaxSize",
			args: args{
				configType:        "logFileMaxSize",
				logFileMaxSizeErr: errors.New("error in getting logFileMaxSize"),
			},
			want:    "5",
			wantErr: true,
		},
		{
			name: "Test 29: When logFileMaxSize is 0",
			args: args{
				configType:     "logFileMaxSize",
				logFileMaxSize: "0",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 30: When GetConfig executes successfully for logFileMaxBackups",
			args: args{
				configType:        "logFileMaxBackups",
				logFileMaxBackups: "20",
			},
			want:    "20",
			wantErr: false,
		},
		{
			name: "Test 31: When there is an error in getting logFileMaxBackups",
			args: args{
				configType:           "logFileMaxBackups",
				logFileMaxBackupsErr: errors.New("error in getting logFileMaxBackups"),
			},
			want:    "10",
			wantErr: true,
		},
		{
			name: "Test 32: When logFileMaxBackups is 0",
			args: args{
				configType:        "logFileMaxBackups",
				logFileMaxBackups: "0",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 33: When GetConfig executes successfully for logFileMaxAge",
			args: args{
				configType:    "logFileMaxAge",
				logFileMaxAge: "15",
			},
			want:    "15",
			wantErr: false,
		},
		{
			name: "Test 34: When there is an error in getting logFileMaxAge",
			args: args{
				configType:       "logFileMaxAge",
				logFileMaxAgeErr: errors.New("error in getting logFileMaxAge"),
			},
			want:    "30",
			wantErr: true,
		},
		{
			name: "Test 35: When logFileMaxAge is 0",
			args: args{
				configType:    "logFileMaxAge",
				logFileMaxAge: "0",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 36: When GetConfig executes successfully",
			args: args{
				configType: "compress",
				compress:   "true",
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "Test 37: When there is an error in getting compress",
			args: args{
				configType:  "compress",
				compressErr: errors.New("error in getting compress"),
			},
			want:    "true",
			wantErr: true,
		},
		{
			name: "Test 38: When compress value is nil",
			args: args{
				configType: "compress",
				compress:   "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			flagSetUtils = flagSetUtilsMock

			flagSetUtilsMock.On("GetRootStringConfig", "provider").Return(tt.args.provider, tt.args.providerErr)
			flagSetUtilsMock.On("GetRootStringConfig", "chainId").Return(tt.args.chainId, tt.args.chainIdErr)
			flagSetUtilsMock.On("GetRootStringConfig", "gasmultiplier").Return(tt.args.gasMultiplier, tt.args.gasMultiplierErr)
			flagSetUtilsMock.On("GetRootStringConfig", "buffer").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			flagSetUtilsMock.On("GetRootStringConfig", "wait").Return(tt.args.waitTime, tt.args.waitTimeErr)
			flagSetUtilsMock.On("GetRootStringConfig", "gasprice").Return(tt.args.gasPrice, tt.args.gasPriceErr)
			flagSetUtilsMock.On("GetRootStringConfig", "logLevel").Return(tt.args.logLevel, tt.args.logLevelErr)
			flagSetUtilsMock.On("GetRootStringConfig", "gasLimit").Return(tt.args.gasLimit, tt.args.gasLimitErr)
			flagSetUtilsMock.On("GetRootStringConfig", "logFileMaxSize").Return(tt.args.logFileMaxSize, tt.args.logFileMaxSizeErr)
			flagSetUtilsMock.On("GetRootStringConfig", "logFileMaxBackups").Return(tt.args.logFileMaxBackups, tt.args.logFileMaxBackupsErr)
			flagSetUtilsMock.On("GetRootStringConfig", "logFileMaxAge").Return(tt.args.logFileMaxAge, tt.args.logFileMaxAgeErr)
			flagSetUtilsMock.On("GetRootStringConfig", "compress").Return(tt.args.compress, tt.args.compressErr)

			ut := &UtilsStruct{}
			got, err := ut.GetConfig(tt.args.configType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
