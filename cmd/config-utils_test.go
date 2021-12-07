package cmd

import (
	"errors"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestGetConfigData(t *testing.T) {
	config := types.Configurations{
		Provider:           "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
	}

	configData := types.Configurations{
		Provider:           "",
		GasMultiplier:      1,
		BufferPercent:      20,
		WaitTime:           1,
		LogLevel:           "debug",
		GasLimitMultiplier: 3,
	}

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		flagSetUtils: FlagSetMock{},
	}
	type args struct {
		provider         string
		providerErr      error
		gasMultiplier    float32
		gasMultiplierErr error
		bufferPercent    int32
		bufferPercentErr error
		waitTime         int32
		waitTimeErr      error
		gasPrice         int32
		gasPriceErr      error
		logLevel         string
		logLevelErr      error
		gasLimit         float32
		gasLimitErr      error
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
				provider:      "",
				gasMultiplier: 1,
				bufferPercent: 20,
				waitTime:      1,
				logLevel:      "debug",
				gasLimit:      3,
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
			name: "Test 3: When there is an error in getting gasMultiplier",
			args: args{
				gasMultiplierErr: errors.New("gasMultiplier error"),
			},
			want:    config,
			wantErr: errors.New("gasMultiplier error"),
		},
		{
			name: "Test 4: When there is an error in getting bufferPercent",
			args: args{
				bufferPercentErr: errors.New("bufferPercent error"),
			},
			want:    config,
			wantErr: errors.New("bufferPercent error"),
		},
		{
			name: "Test 5: When there is an error in getting waitTime",
			args: args{
				waitTimeErr: errors.New("waitTime error"),
			},
			want:    config,
			wantErr: errors.New("waitTime error"),
		},
		{
			name: "Test 6: When there is an error in getting gasPrice",
			args: args{
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    config,
			wantErr: errors.New("gasPrice error"),
		},
		{
			name: "Test 7: When there is an error in getting logLevel",
			args: args{
				logLevelErr: errors.New("logLevel error"),
			},
			want:    config,
			wantErr: errors.New("logLevel error"),
		},
		{
			name: "Test 8: When there is an error in getting gasLimit",
			args: args{
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    config,
			wantErr: errors.New("gasLimit error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getProviderMock = func(UtilsStruct) (string, error) {
				return tt.args.provider, tt.args.providerErr
			}

			getMultiplierMock = func(UtilsStruct) (float32, error) {
				return tt.args.gasMultiplier, tt.args.gasMultiplierErr
			}

			getWaitTimeMock = func(UtilsStruct) (int32, error) {
				return tt.args.waitTime, tt.args.waitTimeErr
			}

			getGasPriceMock = func(UtilsStruct) (int32, error) {
				return tt.args.gasPrice, tt.args.gasPriceErr
			}

			getLogLevelMock = func(UtilsStruct) (string, error) {
				return tt.args.logLevel, tt.args.logLevelErr
			}

			getGasLimitMock = func(UtilsStruct) (float32, error) {
				return tt.args.gasLimit, tt.args.gasLimitErr
			}

			getBufferPercentMock = func(UtilsStruct) (int32, error) {
				return tt.args.bufferPercent, tt.args.bufferPercentErr
			}

			got, err := GetConfigData(utilsStruct)
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

func Test_getBufferPercent(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}

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
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting bufferPercent",
			args: args{
				bufferPercentErr: errors.New("bufferPercent error"),
			},
			want:    30,
			wantErr: errors.New("bufferPercent error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRootInt32BufferMock = func() (int32, error) {
				return tt.args.bufferPercent, tt.args.bufferPercentErr
			}

			got, err := getBufferPercent(utilsStruct)
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

func Test_getGasLimit(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}
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
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasLimit",
			args: args{
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    -1,
			wantErr: errors.New("gasLimit error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRootFloat32GasLimitMock = func() (float32, error) {
				return tt.args.gasLimit, tt.args.gasLimitErr
			}

			got, err := getGasLimit(utilsStruct)
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

func Test_getGasPrice(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}
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
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting gasPrice",
			args: args{
				gasPriceErr: errors.New("gasPrice error"),
			},
			want:    0,
			wantErr: errors.New("gasPrice error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRootInt32GasPriceMock = func() (int32, error) {
				return tt.args.gasPrice, tt.args.gasPriceErr
			}

			got, err := getGasPrice(utilsStruct)
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

func Test_getLogLevel(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}

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
			getRootStringLogLevelMock = func() (string, error) {
				return tt.args.logLevel, tt.args.logLevelErr
			}

			got, err := getLogLevel(utilsStruct)
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

func Test_getMultiplier(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}
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
			want:    0,
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
			GetRootFloat32GasMultiplierMock = func() (float32, error) {
				return tt.args.gasMultiplier, tt.args.gasMultiplierErr
			}

			got, err := getMultiplier(utilsStruct)
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

func Test_getProvider(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}
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
			want:    "",
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 2: When provider is nil",
			args: args{
				provider: "",
			},
			want:    "",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRootStringProviderMock = func() (string, error) {
				return tt.args.provider, tt.args.providerErr
			}

			got, err := getProvider(utilsStruct)
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

func Test_getWaitTime(t *testing.T) {
	utilsStruct := UtilsStruct{
		flagSetUtils: FlagSetMock{},
	}

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
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting waitTime",
			args: args{
				waitTimeErr: errors.New("waitTime error"),
			},
			want:    3,
			wantErr: errors.New("waitTime error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRootInt32WaitMock = func() (int32, error) {
				return tt.args.waitTime, tt.args.waitTimeErr
			}

			got, err := getWaitTime(utilsStruct)
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
