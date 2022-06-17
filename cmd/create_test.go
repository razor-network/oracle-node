package cmd

import (
	"errors"
	"github.com/awnumar/memguard"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	razorAccounts "razor/accounts"
	mocks2 "razor/utils/mocks"

	//"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	Mocks "razor/accounts/mocks"
	//razorAccounts "razor/accounts"
	"razor/cmd/mocks"
	"testing"
)

func TestCreate(t *testing.T) {
	var password string

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

			utilsMock := new(mocks.UtilsInterface)
			accountUtilsMock := new(Mocks.AccountInterface)

			razorUtils = utilsMock
			razorAccounts.AccountUtilsInterface = accountUtilsMock

			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			accountUtilsMock.On("CreateAccount", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(accounts.Account{
				Address: tt.args.account.Address,
				URL:     accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			})

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
		password       string
		keyBuffer      *memguard.LockedBuffer
		keyBufferBytes []byte
		decryptData    []byte
		account        accounts.Account
		accountErr     error
	}

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When executeCreate executes successfully",
			args: args{
				password:       "test",
				keyBufferBytes: []byte("test"),
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

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(mocks2.Utils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			utilsInterface = utilsPkgMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			utilsPkgMock.On("InterruptAndPurge")
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsPkgMock.On("KeyBuffer", mock.Anything).Return(tt.args.keyBuffer)
			utilsPkgMock.On("KeyBufferBytes", mock.Anything).Return(tt.args.keyBufferBytes)
			utilsPkgMock.On("Decrypt", mock.Anything).Return(tt.args.decryptData)
			cmdUtilsMock.On("Create", mock.AnythingOfType("string")).Return(tt.args.account, tt.args.accountErr)
			utilsPkgMock.On("DestroyKeyBuffer", mock.Anything)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteCreate(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeCreate function didn't execute as expected")
			}

		})
	}
}
