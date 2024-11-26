package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestListAccounts(t *testing.T) {

	accountsList := []accounts.Account{
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea2"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
	}

	type args struct {
		path     string
		pathErr  error
		accounts []accounts.Account
	}
	tests := []struct {
		name    string
		args    args
		want    []accounts.Account
		wantErr error
	}{
		{
			name: "When listAccounts executes successfully",
			args: args{
				path:     "test/key/path",
				pathErr:  nil,
				accounts: accountsList,
			},
			want:    accountsList,
			wantErr: nil,
		},
		{
			name: "When listAccounts fails due to path error",
			args: args{
				path:     "test/key/",
				pathErr:  errors.New("path error"),
				accounts: accountsList,
			},
			want:    nil,
			wantErr: errors.New("path error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			keystoreMock.On("Accounts", mock.AnythingOfType("string")).Return(tt.args.accounts)
			utils := &UtilsStruct{}
			got, err := utils.ListAccounts()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List of accounts , got = %v, want %v", got, tt.want)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for listAccounts function, got = %v, want = %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for listAccounts function, got = %v, want = %v", got, tt.wantErr)
				}
			}

		})
	}
}

func TestExecuteListAccounts(t *testing.T) {
	var flagSet *pflag.FlagSet

	accountList := []accounts.Account{
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea2"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
	}
	type args struct {
		allAccounts    []accounts.Account
		allAccountsErr error
	}

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteListAccounts executes successfully",
			args: args{
				allAccounts: accountList,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When ExecuteListAccounts does not execute successfully",
			args: args{
				allAccountsErr: errors.New("allAccounts error"),
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

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("ListAccounts").Return(tt.args.allAccounts, tt.args.allAccountsErr)
			cmdUtilsMock.On("GetConfigData").Return(types.Configurations{}, nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteListAccounts(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteListAccounts function didn't execute as expected")
			}

		})
	}
}
