package cmd

import (
	"errors"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

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
