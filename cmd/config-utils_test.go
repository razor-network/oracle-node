package cmd

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"os"
	"path/filepath"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"reflect"
	"strings"
	"testing"
)

var tempConfigPath = "test_config.yaml"

func createTestConfig(t *testing.T, viperKey string, value interface{}) {

	// Set some values
	viper.Set(viperKey, value)

	// Write the temporary config
	if err := viper.WriteConfigAs(tempConfigPath); err != nil {
		t.Fatalf("Failed to write temp config: %s", err)
	}

	viper.SetConfigName(strings.TrimSuffix(tempConfigPath, filepath.Ext(tempConfigPath)))
	viper.AddConfigPath(".")
}

func removeTestConfig(path string) {
	os.RemoveAll(path)
}

func TestGetConfigData(t *testing.T) {
	nilConfig := types.Configurations{
		Provider:           "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
		RPCTimeout:         0,
		HTTPTimeout:        0,
		LogFileMaxSize:     0,
		LogFileMaxBackups:  0,
		LogFileMaxAge:      0,
	}

	configData := types.Configurations{
		Provider:           "",
		GasMultiplier:      1,
		BufferPercent:      20,
		WaitTime:           1,
		LogLevel:           "debug",
		GasLimitMultiplier: 3,
		GasLimitOverride:   1000000,
		RPCTimeout:         10,
		HTTPTimeout:        10,
		LogFileMaxSize:     5,
		LogFileMaxBackups:  10,
		LogFileMaxAge:      30,
	}

	type args struct {
		provider             string
		providerErr          error
		gasMultiplier        float32
		gasMultiplierErr     error
		bufferPercent        int32
		bufferPercentErr     error
		waitTime             int32
		waitTimeErr          error
		gasPrice             int32
		gasPriceErr          error
		logLevel             string
		logLevelErr          error
		gasLimit             float32
		gasLimitOverride     uint64
		gasLimitOverrideErr  error
		rpcTimeout           int64
		rpcTimeoutErr        error
		httpTimeout          int64
		httpTimeoutErr       error
		gasLimitErr          error
		logFileMaxSize       int
		logFileMaxSizeErr    error
		logFileMaxBackups    int
		logFileMaxBackupsErr error
		logFileMaxAge        int
		logFileMaxAgeErr     error
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
				provider:          "",
				gasMultiplier:     1,
				bufferPercent:     20,
				waitTime:          1,
				logLevel:          "debug",
				gasLimit:          3,
				gasLimitOverride:  1000000,
				rpcTimeout:        10,
				httpTimeout:       10,
				logFileMaxSize:    5,
				logFileMaxBackups: 10,
				logFileMaxAge:     30,
			},
			want:    configData,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting provider",
			args: args{
				providerErr: errors.New("provider error"),
			},
			want:    nilConfig,
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 3: When there is an error in getting gasMultiplier",
			args: args{
				gasMultiplierErr: errors.New("gasMultiplier error"),
			},
			want:    nilConfig,
			wantErr: errors.New("gasMultiplier error"),
		},
		{
			name: "Test 4: When there is an error in getting bufferPercent",
			args: args{
				bufferPercentErr: errors.New("bufferPercent error"),
			},
			want:    nilConfig,
			wantErr: errors.New("bufferPercent error"),
		},
		{
			name: "Test 5: When there is an error in getting waitTime",
			args: args{
				waitTimeErr: errors.New("waitTime error"),
			},
			want:    nilConfig,
			wantErr: errors.New("waitTime error"),
		},
		{
			name: "Test 6: When there is an error in getting gasPrice",
			args: args{
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    nilConfig,
			wantErr: errors.New("gasPrice error"),
		},
		{
			name: "Test 7: When there is an error in getting logLevel",
			args: args{
				logLevelErr: errors.New("logLevel error"),
			},
			want:    nilConfig,
			wantErr: errors.New("logLevel error"),
		},
		{
			name: "Test 8: When there is an error in getting gasLimit",
			args: args{
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    nilConfig,
			wantErr: errors.New("gasLimit error"),
		},
		{
			name: "Test 9: When there is an error in getting rpcTimeout",
			args: args{
				rpcTimeoutErr: errors.New("rpcTimeout error"),
			},
			want:    nilConfig,
			wantErr: errors.New("rpcTimeout error"),
		},
		{
			name: "Test 10: When there is an error in getting httpTimeout",
			args: args{
				httpTimeoutErr: errors.New("httpTimeout error"),
			},
			want:    nilConfig,
			wantErr: errors.New("httpTimeout error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			cmdUtilsMock.On("GetProvider").Return(tt.args.provider, tt.args.providerErr)
			cmdUtilsMock.On("GetMultiplier").Return(tt.args.gasMultiplier, tt.args.gasMultiplierErr)
			cmdUtilsMock.On("GetWaitTime").Return(tt.args.waitTime, tt.args.waitTimeErr)
			cmdUtilsMock.On("GetGasPrice").Return(tt.args.gasPrice, tt.args.gasPriceErr)
			cmdUtilsMock.On("GetLogLevel").Return(tt.args.logLevel, tt.args.logLevelErr)
			cmdUtilsMock.On("GetGasLimit").Return(tt.args.gasLimit, tt.args.gasLimitErr)
			cmdUtilsMock.On("GetGasLimitOverride").Return(tt.args.gasLimitOverride, tt.args.gasLimitOverrideErr)
			cmdUtilsMock.On("GetBufferPercent").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			cmdUtilsMock.On("GetRPCTimeout").Return(tt.args.rpcTimeout, tt.args.rpcTimeoutErr)
			cmdUtilsMock.On("GetHTTPTimeout").Return(tt.args.httpTimeout, tt.args.httpTimeoutErr)
			cmdUtilsMock.On("GetLogFileMaxSize").Return(tt.args.logFileMaxSize, tt.args.logFileMaxSizeErr)
			cmdUtilsMock.On("GetLogFileMaxBackups").Return(tt.args.logFileMaxBackups, tt.args.logFileMaxBackupsErr)
			cmdUtilsMock.On("GetLogFileMaxAge").Return(tt.args.logFileMaxAge, tt.args.logFileMaxAgeErr)
			utilsMock.On("IsFlagPassed", mock.AnythingOfType("string")).Return(true)
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

func TestGetBufferPercent(t *testing.T) {
	type args struct {
		isFlagSet          bool
		bufferPercent      int32
		bufferPercentErr   error
		bufferInTestConfig int32
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               int32
		wantErr            error
	}{
		{
			name: "Test 1: When buffer percent is fetched from root flag",
			args: args{
				isFlagSet:     true,
				bufferPercent: 10,
			},
			want:    10,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching buffer from root flag",
			args: args{
				isFlagSet:        true,
				bufferPercentErr: errors.New("buffer percent error"),
			},
			want:    core.DefaultBufferPercent,
			wantErr: errors.New("buffer percent error"),
		},
		{
			name:               "Test 3: When buffer value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				bufferInTestConfig: 6,
			},
			want:    6,
			wantErr: nil,
		},
		{
			name:    "Test 4: When buffer is not passed in root nor set in config",
			want:    core.DefaultBufferPercent,
			wantErr: nil,
		},
		{
			name:               "Test 5: When buffer value is out of a valid range",
			useDummyConfigFile: true,
			args: args{
				bufferInTestConfig: 0,
			},
			want:    core.DefaultBufferPercent,
			wantErr: errors.New("invalid buffer percent"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "buffer", tt.args.bufferInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetBufferPercent()
			if got != tt.want {
				t.Errorf("getBufferPercent() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getBufferPercent function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getBufferPercent function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestGetGasLimit(t *testing.T) {
	type args struct {
		isFlagSet            bool
		gasLimit             float32
		gasLimitErr          error
		gasLimitInTestConfig float32
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               float32
		wantErr            error
	}{
		{
			name: "Test 1: When gasLimit is fetched from root flag",
			args: args{
				isFlagSet: true,
				gasLimit:  2,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching gasLimit from root flag",
			args: args{
				isFlagSet:   true,
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    core.DefaultGasLimit,
			wantErr: errors.New("gasLimit error"),
		},
		{
			name:               "Test 3: When gas value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				gasLimitInTestConfig: 2.5,
			},
			want:    2.5,
			wantErr: nil,
		},
		{
			name:    "Test 4: When gasLimit is not passed in root nor set in config",
			want:    core.DefaultGasLimit,
			wantErr: nil,
		},
		{
			name:               "Test 5: When gas limit value is out of valid range",
			useDummyConfigFile: true,
			args: args{
				gasLimitInTestConfig: 3.5,
			},
			want:    3.5,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "gasLimit", tt.args.gasLimitInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetGasLimit()
			if got != tt.want {
				t.Errorf("getGasLimit() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getGasLimit function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getGasLimit function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetGasLimitOverride(t *testing.T) {
	type args struct {
		isFlagSet                bool
		gasLimitOverride         uint64
		gasLimitOverrideErr      error
		gasLimitOverrideInConfig uint64
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               uint64
		wantErr            error
	}{
		{
			name: "Test 1: When gasLimitOverride is fetched from root flag",
			args: args{
				isFlagSet:        true,
				gasLimitOverride: 40000000,
			},
			want:    40000000,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching gasLimitOverride from root flag",
			args: args{
				isFlagSet:           true,
				gasLimitOverrideErr: errors.New("gasLimitOverride error"),
			},
			want:    core.DefaultGasLimitOverride,
			wantErr: errors.New("gasLimitOverride error"),
		},
		{
			name:               "Test 3: When gasLimitOverride is fetched from config",
			useDummyConfigFile: true,
			args: args{
				gasLimitOverrideInConfig: 30000000,
			},
			want:    30000000,
			wantErr: nil,
		},
		{
			name:    "Test 4: When gasLimitOverride is not passed in root nor set in config",
			want:    core.DefaultGasLimitOverride,
			wantErr: nil,
		},
		{
			name:               "Test 3: When gasLimitOverride is fetched from config",
			useDummyConfigFile: true,
			args: args{
				gasLimitOverrideInConfig: 60000000,
			},
			want:    core.DefaultGasLimitOverride,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "gasLimitOverride", tt.args.gasLimitOverrideInConfig)
				defer removeTestConfig(tempConfigPath)
			}

			flagSetUtilsMock := new(mocks.FlagSetInterface)
			flagSetUtils = flagSetUtilsMock

			flagSetUtilsMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.gasLimitOverride, tt.args.gasLimitOverrideErr)
			flagSetUtilsMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetGasLimitOverride()
			if got != tt.want {
				t.Errorf("getGasLimitOverride() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getGasLimitOverride function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getGasLimitOverride function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetGasPrice(t *testing.T) {
	type args struct {
		isFlagSet            bool
		gasPrice             int32
		gasPriceErr          error
		gasPriceInTestConfig int32
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               int32
		wantErr            error
	}{
		{
			name: "Test 1: When gasPrice is fetched from root flag",
			args: args{
				isFlagSet: true,
				gasPrice:  1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching gasPrice from root flag",
			args: args{
				isFlagSet:   true,
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    core.DefaultGasPrice,
			wantErr: errors.New("gasPrice error"),
		},
		{
			name:               "Test 3: When gasPrice value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				gasPriceInTestConfig: 0,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "Test 4: When gasPrice is not passed in root nor set in config",
			want:    core.DefaultGasPrice,
			wantErr: nil,
		},
		{
			name:               "Test 5: When gasPrice is out of valid range",
			useDummyConfigFile: true,
			args: args{
				gasPriceInTestConfig: 3,
			},
			want:    core.DefaultGasPrice,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "gasprice", tt.args.gasPriceInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			flagSetUtilsMock := new(mocks.FlagSetInterface)
			flagSetUtils = flagSetUtilsMock

			flagSetUtilsMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			flagSetUtilsMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetGasPrice()
			if got != tt.want {
				t.Errorf("getGasPrice() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getGasPrice function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getGasPrice function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetLogLevel(t *testing.T) {
	type args struct {
		isFlagSet            bool
		logLevel             string
		logLevelErr          error
		logLevelInTestConfig string
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               string
		wantErr            error
	}{
		{
			name: "Test 1: When logLevel is fetched from root flag",
			args: args{
				isFlagSet: true,
				logLevel:  "debug",
			},
			want:    "debug",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching logLevel from root flag",
			args: args{
				isFlagSet:   true,
				logLevelErr: errors.New("logLevel error"),
			},
			want:    core.DefaultLogLevel,
			wantErr: errors.New("logLevel error"),
		},
		{
			name:               "Test 3: When logLevel value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				logLevelInTestConfig: "info",
			},
			want:    "info",
			wantErr: nil,
		},
		{
			name:    "Test 4: When logLevel is not passed in root nor set in config",
			want:    core.DefaultLogLevel,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "logLevel", tt.args.logLevelInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.logLevel, tt.args.logLevelErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetLogLevel()
			if got != tt.want {
				t.Errorf("getLogLevel() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getLogLevel function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getLogLevel function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetMultiplier(t *testing.T) {
	type args struct {
		isFlagSet                 bool
		gasMultiplier             float32
		gasMultiplierErr          error
		gasMultiplierInTestConfig float32
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               float32
		wantErr            error
	}{
		{
			name: "Test 1: When gasMultiplier is fetched from root flag",
			args: args{
				isFlagSet:     true,
				gasMultiplier: 2,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching gasMultiplier from root flag",
			args: args{
				isFlagSet:        true,
				gasMultiplierErr: errors.New("gasMultiplier error"),
			},
			want:    core.DefaultGasMultiplier,
			wantErr: errors.New("gasMultiplier error"),
		},
		{
			name:               "Test 3: When gasMultiplier value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				gasMultiplierInTestConfig: 3,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name:    "Test 4: When gasMultiplier is not passed in root nor set in config",
			want:    core.DefaultGasMultiplier,
			wantErr: nil,
		},
		{
			name:               "Test 5: When gasMultiplier is out of a valid range",
			useDummyConfigFile: true,
			args: args{
				gasMultiplierInTestConfig: 4,
			},
			want:    core.DefaultGasMultiplier,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "gasmultiplier", tt.args.gasMultiplierInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.gasMultiplier, tt.args.gasMultiplierErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetMultiplier()
			if got != tt.want {
				t.Errorf("getMultiplier() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getMultiplier function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getMultiplier function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetProvider(t *testing.T) {
	type args struct {
		provider             string
		providerErr          error
		isFlagSet            bool
		providerInTestConfig string
	}
	tests := []struct {
		name               string
		args               args
		useDummyConfigFile bool
		want               string
		wantErr            error
	}{
		{
			name: "Test 1: When provider is fetched from root flag",
			args: args{
				provider:  "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
				isFlagSet: true,
			},
			want:    "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching provider from root flag",
			args: args{
				providerErr: errors.New("provider error"),
				isFlagSet:   true,
			},
			want:    "",
			wantErr: errors.New("provider error"),
		},
		{
			name:               "Test 3: When provider value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				providerInTestConfig: "https://config-provider-url.com",
			},
			want:    "https://config-provider-url.com",
			wantErr: nil,
		},
		{
			name:    "Test 4: When provider is neither passed in root nor set in config",
			args:    args{},
			want:    "",
			wantErr: errors.New("provider is not set"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "provider", tt.args.providerInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.provider, tt.args.providerErr)
			flagSetMock.On("Changed", mock.Anything, "provider").Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetProvider()
			if got != tt.want {
				t.Errorf("getProvider() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getProvider function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getProvider function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetRPCTimeout(t *testing.T) {
	type args struct {
		isFlagSet              bool
		rpcTimeout             int64
		rpcTimeoutErr          error
		rpcTimeoutInTestConfig int64
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               int64
		wantErr            error
	}{
		{
			name: "Test 1: When rpcTimeout is fetched from root flag",
			args: args{
				isFlagSet:  true,
				rpcTimeout: 6,
			},
			want:    6,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching rpcTimeout from root flag",
			args: args{
				isFlagSet:     true,
				rpcTimeoutErr: errors.New("rpcTimeout error"),
			},
			want:    core.DefaultRPCTimeout,
			wantErr: errors.New("rpcTimeout error"),
		},
		{
			name:               "Test 3: When rpcTimeout value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				rpcTimeoutInTestConfig: 7,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name:    "Test 4: When rpcTimeout is not passed in root nor set in config",
			want:    core.DefaultRPCTimeout,
			wantErr: nil,
		},
		{
			name:               "Test 5: When rpcTimeout value is out of a valid range",
			useDummyConfigFile: true,
			args: args{
				rpcTimeoutInTestConfig: 70,
			},
			want:    core.DefaultRPCTimeout,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "rpcTimeout", tt.args.rpcTimeoutInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.rpcTimeout, tt.args.rpcTimeoutErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetRPCTimeout()
			if got != tt.want {
				t.Errorf("getRPCTimeout() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getRPCTimeout function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getRPCTimeout function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetHTTPTimeout(t *testing.T) {
	type args struct {
		isFlagSet               bool
		httpTimeout             int64
		httpTimeoutErr          error
		httpTimeoutInTestConfig int64
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               int64
		wantErr            error
	}{
		{
			name: "Test 1: When httpTimeout is fetched from root flag",
			args: args{
				isFlagSet:   true,
				httpTimeout: 6,
			},
			want:    6,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching httpTimeout from root flag",
			args: args{
				isFlagSet:      true,
				httpTimeoutErr: errors.New("httpTimeout error"),
			},
			want:    core.DefaultHTTPTimeout,
			wantErr: errors.New("httpTimeout error"),
		},
		{
			name:               "Test 3: When httpTimeout value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				httpTimeoutInTestConfig: 7,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name:    "Test 4: When httpTimeout is not passed in root nor set in config",
			want:    core.DefaultHTTPTimeout,
			wantErr: nil,
		},
		{
			name:               "Test 5: When httpTimeout is out of valid range",
			useDummyConfigFile: true,
			args: args{
				httpTimeoutInTestConfig: 70,
			},
			want:    core.DefaultHTTPTimeout,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "httpTimeout", tt.args.httpTimeoutInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.httpTimeout, tt.args.httpTimeoutErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetHTTPTimeout()
			if got != tt.want {
				t.Errorf("getHTTPTimeout() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getHTTPTimeout function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getHTTPTimeout function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetWaitTime(t *testing.T) {
	type args struct {
		isFlagSet        bool
		waitTime         int32
		waitTimeErr      error
		waitInTestConfig int32
	}
	tests := []struct {
		name               string
		useDummyConfigFile bool
		args               args
		want               int32
		wantErr            error
	}{
		{
			name: "Test 1: When wait time is fetched from root flag",
			args: args{
				isFlagSet: true,
				waitTime:  2,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching wait time from root flag",
			args: args{
				isFlagSet:   true,
				waitTimeErr: errors.New("wait time error"),
			},
			want:    core.DefaultWaitTime,
			wantErr: errors.New("wait time error"),
		},
		{
			name:               "Test 3: When wait time value is fetched from config",
			useDummyConfigFile: true,
			args: args{
				waitInTestConfig: 3,
			},
			want:    3,
			wantErr: nil,
		},
		{
			name:    "Test 4: When wait time is not passed in root nor set in config",
			want:    core.DefaultWaitTime,
			wantErr: nil,
		},
		{
			name:               "Test 5: When wait time value is out of valid range",
			useDummyConfigFile: true,
			args: args{
				waitInTestConfig: 40,
			},
			want:    core.DefaultWaitTime,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset() // Reset viper state

			if tt.useDummyConfigFile {
				createTestConfig(t, "wait", tt.args.waitInTestConfig)
				defer removeTestConfig(tempConfigPath)
			}

			SetUpMockInterfaces()

			flagSetMock.On("FetchRootFlagInput", mock.Anything, mock.Anything).Return(tt.args.waitTime, tt.args.waitTimeErr)
			flagSetMock.On("Changed", mock.Anything, mock.Anything).Return(tt.args.isFlagSet)

			utils := &UtilsStruct{}
			got, err := utils.GetWaitTime()
			if got != tt.want {
				t.Errorf("GetWaitTime() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetWaitTime function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetWaitTime function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestValidateBufferPercentLimit(t *testing.T) {
	type args struct {
		bufferPercent  int32
		stateBuffer    uint64
		stateBufferErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Buffer percent less than max buffer percent",
			args: args{
				stateBuffer:   10,
				bufferPercent: 20,
			},
			wantErr: nil,
		},
		{
			name: "Buffer percent greater than max buffer percent",
			args: args{
				stateBuffer:   10,
				bufferPercent: 60,
			},
			wantErr: errors.New("buffer percent exceeds limit"),
		},
		{
			name: "GetStateBuffer returns an error",
			args: args{
				stateBufferErr: errors.New("state buffer error"),
				bufferPercent:  10,
			},
			wantErr: errors.New("state buffer error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetStateBuffer", mock.Anything).Return(tt.args.stateBuffer, tt.args.stateBufferErr)

			err := ValidateBufferPercentLimit(rpcParameters, tt.args.bufferPercent)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetEpochAndState function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetEpochAndState function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
