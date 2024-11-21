package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestExtendLock(t *testing.T) {
	var extendLockInput types.ExtendLockInput
	var config types.Configurations

	type args struct {
		txnOptsErr   error
		resetLockTxn *Types.Transaction
		resetLockErr error
		hash         common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When resetLock function executes successfully",
			args: args{
				resetLockTxn: &Types.Transaction{},
				resetLockErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When ResetLock transaction fails",
			args: args{
				resetLockTxn: &Types.Transaction{},
				resetLockErr: errors.New("resetLock error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("resetLock error"),
		},
		{
			name: "Test 3: When there is an error in getting txnOpts",
			args: args{
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			stakeManagerMock.On("ResetUnstakeLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32")).Return(tt.args.resetLockTxn, tt.args.resetLockErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.ResetUnstakeLock(rpcParameters, config, extendLockInput)
			if got != tt.want {
				t.Errorf("Txn hash for resetLock function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for resetLock function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for resetLock function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestExecuteExtendLock(t *testing.T) {

	var flagSet *pflag.FlagSet
	var config types.Configurations
	var client *ethclient.Client

	type args struct {
		config       types.Configurations
		configErr    error
		password     string
		address      string
		addressErr   error
		stakerId     uint32
		stakerIdErr  error
		resetLockTxn common.Hash
		resetLockErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When resetLock function executes successfully",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				stakerId:     1,
				resetLockTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting address from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				stakerId:     1,
				resetLockTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting stakerId from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				stakerIdErr:  errors.New("stakerId error"),
				resetLockTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When ResetLock transaction fails",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				stakerId:     1,
				resetLockTxn: core.NilHash,
				resetLockErr: errors.New("resetLock error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				stakerId:     1,
				resetLockTxn: common.BigToHash(big.NewInt(1)),
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

			utilsMock.On("IsFlagPassed", mock.Anything).Return(true)
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("AssignStakerId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)
			cmdUtilsMock.On("ResetUnstakeLock", mock.Anything, config, mock.Anything).Return(tt.args.resetLockTxn, tt.args.resetLockErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteExtendLock(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteExtendLock function didn't execute as expected")
			}

		})
	}
}
