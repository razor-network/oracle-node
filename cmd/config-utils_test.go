package cmd

import (
	"errors"
	"razor/cmd/mocks"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestGetConfigData(t *testing.T) {
	nilConfig := types.Configurations{
		Provider:           "",
		AlternateProvider:  "",
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
		AlternateProvider:  "",
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
		alternateProvider    string
		alternateProviderErr error
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
				alternateProvider: "",
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
		{
			name: "Test 11: When there is an error in getting alternate provider",
			args: args{
				alternateProviderErr: errors.New("alternate provider error"),
			},
			want:    nilConfig,
			wantErr: errors.New("alternate provider error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			cmdUtilsMock.On("GetProvider").Return(tt.args.provider, tt.args.providerErr)
			cmdUtilsMock.On("GetAlternateProvider").Return(tt.args.alternateProvider, tt.args.alternateProviderErr)
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
		bufferPercent    int32
		bufferPercentErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr error
	}{
		{
			name: "Test 1: When getBufferPercent function executes successfully",
			args: args{
				bufferPercent: 20,
			},
			want:    20,
			wantErr: nil,
		},
		{
			name: "Test 2: When bufferPercent is 0",
			args: args{
				bufferPercent: 0,
			},
			want:    20,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting bufferPercent",
			args: args{
				bufferPercentErr: errors.New("bufferPercent error"),
			},
			want:    20,
			wantErr: errors.New("bufferPercent error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootInt32Buffer").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
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
}

func TestGetGasLimit(t *testing.T) {
	type args struct {
		gasLimit    float32
		gasLimitErr error
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr error
	}{
		{
			name: "Test 1: When getGasLimit function executes successfully",
			args: args{
				gasLimit: 4,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 2: When gasLimit is -1",
			args: args{
				gasLimit: -1,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasLimit",
			args: args{
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    2,
			wantErr: errors.New("gasLimit error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootFloat32GasLimit").Return(tt.args.gasLimit, tt.args.gasLimitErr)
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
		gasLimitOverride    uint64
		gasLimitOverrideErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error
	}{
		{
			name: "Test 1: When getGasLimitOverride function executes successfully",
			args: args{
				gasLimitOverride: 5000000,
			},
			want:    5000000,
			wantErr: nil,
		},
		{
			name: "Test 2: When gasLimitOverride is 0",
			args: args{
				gasLimitOverride: 0,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasLimitOverride",
			args: args{
				gasLimitOverrideErr: errors.New("gasLimitOverride error"),
			},
			want:    0,
			wantErr: errors.New("gasLimitOverride error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			flagSetUtils = flagSetUtilsMock

			flagSetUtilsMock.On("GetRootUint64GasLimitOverride").Return(tt.args.gasLimitOverride, tt.args.gasLimitOverrideErr)
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
		gasPrice    int32
		gasPriceErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr error
	}{
		{
			name: "Test 1: When getGasPrice function executes successfully",
			args: args{
				gasPrice: 1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 2: When gasPrice is -1",
			args: args{
				gasPrice: -1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasPrice",
			args: args{
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    1,
			wantErr: errors.New("gasPrice error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			flagSetUtils = flagSetUtilsMock

			flagSetUtilsMock.On("GetRootInt32GasPrice").Return(tt.args.gasPrice, tt.args.gasPriceErr)
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
		logLevel    string
		logLevelErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When getLogLevel function executes successfully",
			args: args{
				logLevel: "debug",
			},
			want:    "debug",
			wantErr: nil,
		},
		{
			name: "Test 2: When logLevel is nil",
			args: args{
				logLevel: "",
			},
			want:    "",
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting logLevel",
			args: args{
				logLevelErr: errors.New("logLevel error"),
			},
			want:    "",
			wantErr: errors.New("logLevel error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootStringLogLevel").Return(tt.args.logLevel, tt.args.logLevelErr)
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
		gasMultiplier    float32
		gasMultiplierErr error
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr error
	}{
		{
			name: "Test 1: When getMultiplier function executes successfully",
			args: args{
				gasMultiplier: 2,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When gasMultiplier is -1",
			args: args{
				gasMultiplier: -1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasMultiplier",
			args: args{
				gasMultiplierErr: errors.New("gasMultiplier error"),
			},
			want:    1,
			wantErr: errors.New("gasMultiplier error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootFloat32GasMultiplier").Return(tt.args.gasMultiplier, tt.args.gasMultiplierErr)
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
		provider    string
		providerErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When getProvider function execute successfully",
			args: args{
				provider: "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			},
			want:    "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			wantErr: nil,
		},
		{
			name: "Test 2: When provider has prefix https",
			args: args{
				provider: "127.0.0.1:8545",
			},
			want:    "127.0.0.1:8545",
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting provider",
			args: args{
				providerErr: errors.New("provider error"),
			},
			want:    "http://127.0.0.1:8545",
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 4: When provider is nil",
			args: args{
				provider: "",
			},
			want:    "http://127.0.0.1:8545",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootStringProvider").Return(tt.args.provider, tt.args.providerErr)
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

func TestGetAlternateProvider(t *testing.T) {
	type args struct {
		alternateProvider    string
		alternateProviderErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When getAlternateProvider function execute successfully",
			args: args{
				alternateProvider: "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			},
			want:    "https://polygon-mumbai.g.alchemy.com/v2/-Re1lE3oDIVTWchuKMfRIECn0I",
			wantErr: nil,
		},
		{
			name: "Test 2: When alternate provider has prefix https",
			args: args{
				alternateProvider: "127.0.0.1:8545",
			},
			want:    "127.0.0.1:8545",
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting alternate provider",
			args: args{
				alternateProviderErr: errors.New("alternateProvider error"),
			},
			want:    "http://127.0.0.1:8545",
			wantErr: errors.New("alternateProvider error"),
		},
		{
			name: "Test 4: When alternate provider is nil",
			args: args{
				alternateProvider: "",
			},
			want:    "http://127.0.0.1:8545",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootStringAlternateProvider").Return(tt.args.alternateProvider, tt.args.alternateProviderErr)
			utils := &UtilsStruct{}

			got, err := utils.GetAlternateProvider()
			if got != tt.want {
				t.Errorf("getAlternateProvider() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getAlternateProvider function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getAlternateProvider function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetWaitTime(t *testing.T) {
	type args struct {
		waitTime    int32
		waitTimeErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr error
	}{
		{
			name: "Test 1: When getWaitTime function executes successfully",
			args: args{
				waitTime: 4,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 2: When waitTime is -1",
			args: args{
				waitTime: -1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting waitTime",
			args: args{
				waitTimeErr: errors.New("waitTime error"),
			},
			want:    1,
			wantErr: errors.New("waitTime error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootInt32Wait").Return(tt.args.waitTime, tt.args.waitTimeErr)
			utils := &UtilsStruct{}
			got, err := utils.GetWaitTime()
			if got != tt.want {
				t.Errorf("getWaitTime() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getWaitTime function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getWaitTime function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetRPCTimeout(t *testing.T) {
	type args struct {
		rpcTimeout    int64
		rpcTimeoutErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			name: "Test 1: When getRPCTimeout function executes successfully",
			args: args{
				rpcTimeout: 12,
			},
			want:    12,
			wantErr: nil,
		},
		{
			name: "Test 2: When rpcTimeout is 0",
			args: args{
				rpcTimeout: 0,
			},
			want:    10,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting rpcTimeout",
			args: args{
				rpcTimeoutErr: errors.New("rpcTimeout error"),
			},
			want:    10,
			wantErr: errors.New("rpcTimeout error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootInt64RPCTimeout").Return(tt.args.rpcTimeout, tt.args.rpcTimeoutErr)
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
		httpTimeout    int64
		httpTimeoutErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			name: "Test 1: When getHTTPTimeout function executes successfully",
			args: args{
				httpTimeout: 12,
			},
			want:    12,
			wantErr: nil,
		},
		{
			name: "Test 2: When httpTimeout is 0",
			args: args{
				httpTimeout: 0,
			},
			want:    10,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting httpTimeout",
			args: args{
				httpTimeoutErr: errors.New("httpTimeout error"),
			},
			want:    10,
			wantErr: errors.New("httpTimeout error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			flagSetMock.On("GetRootInt64HTTPTimeout").Return(tt.args.httpTimeout, tt.args.httpTimeoutErr)
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
