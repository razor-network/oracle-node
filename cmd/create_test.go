package cmd

import (
	"errors"
	accountsPkgMocks "razor/accounts/mocks"
	"razor/core/types"
	pathPkgMocks "razor/path/mocks"
	utilsPkgMocks "razor/utils/mocks"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	var password string

	nilAccount := accounts.Account{Address: common.Address{0x00},
		URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
	}

	type args struct {
		path              string
		pathErr           error
		accountManagerErr error
		account           accounts.Account
	}
	tests := []struct {
		name    string
		args    args
		want    accounts.Account
		wantErr error
	}{
		{
			name: "Test 1: When create function executes successfully",
			args: args{
				path:    "/home/local",
				pathErr: nil,
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When create fails due to path error",
			args: args{
				path:    "/home/local",
				pathErr: errors.New("path error"),
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			want:    nilAccount,
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 3: When there is an error in getting account manager",
			args: args{
				path:              "/home/local",
				pathErr:           nil,
				accountManagerErr: errors.New("account manager error"),
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			want:    nilAccount,
			wantErr: errors.New("account manager error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountsMock := new(accountsPkgMocks.AccountManagerInterface)

			var pathMock *pathPkgMocks.PathInterface
			var utilsMock *utilsPkgMocks.Utils

			pathMock = new(pathPkgMocks.PathInterface)
			pathUtils = pathMock

			utilsMock = new(utilsPkgMocks.Utils)
			razorUtils = utilsMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			utilsMock.On("AccountManagerForKeystore").Return(accountsMock, tt.args.accountManagerErr)

			accountsMock.On("CreateAccount", mock.Anything, mock.Anything).Return(tt.args.account)

			utils := &UtilsStruct{}
			got, err := utils.Create(password)

			if got.Address != tt.want.Address {
				t.Errorf("New address created, got = %v, want %v", got, tt.want.Address)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for Create function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for Create function, got = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteCreate(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		password   string
		account    accounts.Account
		accountErr error
	}

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When executeCreate executes successfully",
			args: args{
				password: "test",
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error from create function",
			args: args{
				password: "test",
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
				accountErr: errors.New("create error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			setupTestEndpointsEnvironment()

			utilsMock.On("IsFlagPassed", mock.Anything).Return(false)
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			cmdUtilsMock.On("Create", mock.AnythingOfType("string")).Return(tt.args.account, tt.args.accountErr)
			cmdUtilsMock.On("GetConfigData").Return(types.Configurations{}, nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteCreate(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeCreate function didn't execute as expected")
			}

		})
	}
}
