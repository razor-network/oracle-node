package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"testing"
)

func TestCreate(t *testing.T) {
	var password string

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		accountUtils: AccountMock{},
	}

	type args struct {
		path    string
		pathErr error
		account accounts.Account
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
			want: accounts.Account{Address: common.Address{0x00},
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			wantErr: errors.New("path error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetDefaultPathMock = func() (string, error) {
				return tt.args.path, tt.args.pathErr
			}

			CreateAccountMock = func(string, string) accounts.Account {
				return accounts.Account{
					Address: tt.args.account.Address,
					URL:     accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				}
			}

			got, err := Create(password, utilsStruct)

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

func Test_executeCreate(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		password   string
		account    accounts.Account
		accountErr error
	}

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
		cmdUtils:   UtilsCmdMock{},
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			CreateMock = func(string, UtilsStruct) (accounts.Account, error) {
				return tt.args.account, tt.args.accountErr
			}

			fatal = false

			utilsStruct.executeCreate(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeCreate function didn't execute as expected")
			}

		})
	}
}
