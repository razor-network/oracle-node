package cmd

import (
	"errors"
	"github.com/spf13/pflag"
	"testing"
)

func TestSetConfig(t *testing.T) {

	var flagSet *pflag.FlagSet

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		flagSetUtils: FlagSetMock{},
	}

	type args struct {
		provider         string
		providerErr      error
		gasmultiplier    float32
		gasmultiplierErr error
		buffer           int32
		bufferErr        error
		waitTime         int32
		waitTimeErr      error
		gasPrice         int32
		gasPriceErr      error
		logLevel         string
		logLevelErr      error
		path             string
		pathErr          error
		configErr        error
		gasLimit         int32
		gasLimitErr      error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When values are passed to all flags and setConfig returns no error",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When parameters are set to default values and setConfig returns no error",
			args: args{
				provider:      "",
				gasmultiplier: -1,
				buffer:        0,
				waitTime:      -1,
				gasPrice:      -1,
				logLevel:      "",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting provider",
			args: args{
				providerErr:   errors.New("provider error"),
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
			},
			wantErr: errors.New("provider error"),
		},
		{
			name: "Test 4: When there is an error in getting gasmultiplier",
			args: args{
				provider:         "http://127.0.0.1",
				gasmultiplierErr: errors.New("gasmultiplier error"),
				buffer:           20,
				waitTime:         2,
				gasPrice:         1,
				logLevel:         "debug",
				path:             "/home/config",
				configErr:        nil,
				gasLimit:         10,
				gasLimitErr:      nil,
			},
			wantErr: errors.New("gasmultiplier error"),
		},
		{
			name: "Test 5: When there is an error in getting buffer",
			args: args{
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				bufferErr:     errors.New("buffer error"),
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
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
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
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
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
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
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
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
				configErr:     nil,
				gasLimit:      10,
				gasLimitErr:   nil,
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
				gasLimit:      10,
				gasLimitErr:   nil,
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
				provider:      "http://127.0.0.1",
				gasmultiplier: 2,
				buffer:        20,
				waitTime:      2,
				gasPrice:      1,
				logLevel:      "debug",
				path:          "/home/config",
				configErr:     nil,
				gasLimit:      -1,
				gasLimitErr:   errors.New("gasLimit error"),
			},
			wantErr: errors.New("gasLimit error"),
		},
	}
	for _, tt := range tests {
		GetStringProviderMock = func(set *pflag.FlagSet) (string, error) {
			return tt.args.provider, tt.args.providerErr
		}

		GetFloat32GasMultiplierMock = func(set *pflag.FlagSet) (float32, error) {
			return tt.args.gasmultiplier, tt.args.gasmultiplierErr
		}

		GetInt32BufferMock = func(set *pflag.FlagSet) (int32, error) {
			return tt.args.buffer, tt.args.bufferErr
		}

		GetInt32WaitMock = func(set *pflag.FlagSet) (int32, error) {
			return tt.args.waitTime, tt.args.waitTimeErr
		}

		GetInt32GasPriceMock = func(set *pflag.FlagSet) (int32, error) {
			return tt.args.gasPrice, tt.args.gasPriceErr
		}

		GetStringLogLevelMock = func(set *pflag.FlagSet) (string, error) {
			return tt.args.logLevel, tt.args.logLevelErr
		}

		GetInt32GasLimitMock = func(set *pflag.FlagSet) (int32, error) {
			return tt.args.gasLimit, tt.args.gasLimitErr
		}

		GetConfigFilePathMock = func() (string, error) {
			return tt.args.path, tt.args.pathErr
		}

		ViperWriteConfigAsMock = func(string) error {
			return tt.args.configErr
		}

		t.Run(tt.name, func(t *testing.T) {
			gotErr := utilsStruct.SetConfig(flagSet)
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
