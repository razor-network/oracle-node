package cmd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestHandleUnstakeLock(t *testing.T) {
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	type args struct {
		state                    uint32
		stateErr                 error
		lock                     types.Locks
		lockErr                  error
		withdrawReleasePeriod    uint16
		withdrawReleasePeriodErr error
		epoch                    uint32
		epochErr                 error
		time                     string
		txnOptsErr               error
		withdrawHash             common.Hash
		withdrawErr              error
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When withdrawFunds function executes successfully",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 5,
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting appropriate state",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 5,
				stateErr:              errors.New("error in getting state"),
			},
			want:    core.NilHash,
			wantErr: errors.New("error in getting state"),
		},
		{
			name: "Test 3: When there is an error in getting lock",
			args: args{
				lockErr:               errors.New("lock error"),
				withdrawReleasePeriod: 4,
				epoch:                 5,
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 4: When staker tries to withdraw without un-staking any Razors",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(0),
				},
				withdrawReleasePeriod: 4,
				epoch:                 5,
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("unstake Razors before withdrawing"),
		},
		{
			name: "Test 5: When there is an error in getting withdrawReleasePeriod",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriodErr: errors.New("withdrawReleasePeriod error"),
				epoch:                    5,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("withdrawReleasePeriod error"),
		},
		{
			name: "Test 6: When there is an error in getting epoch",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epochErr:              errors.New("epoch error"),
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 7: When staker tries to withdraw when withdrawal period has passed",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 9,
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When staker tries to withdraw when withdrawal period has not reached",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 3,
				time:                  "10 minutes 0 seconds ",
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 9: When there is an error in executing withdraw function",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 1,
				epoch:                 5,
				withdrawErr:           errors.New("withdraw error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("withdraw error"),
		},
		{
			name: "Test 10: When staker tries to withdraw when withdrawal period has not reached",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 2,
				time:                  "20 minutes 0 seconds ",
				withdrawHash:          common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 11: When there is an error in getting txnOpts",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				withdrawReleasePeriod: 4,
				epoch:                 5,
				txnOptsErr:            errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			cmdUtilsMock.On("WaitForAppropriateState", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetLock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.lock, tt.args.lockErr)
			utilsMock.On("GetWithdrawInitiationPeriod", mock.Anything, mock.Anything).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			utilsMock.On("GetEpoch", mock.Anything).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			cmdUtilsMock.On("InitiateWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)
			utilsMock.On("SecondsToReadableTime", mock.AnythingOfType("int")).Return(tt.args.time)

			utils := &UtilsStruct{}
			got, err := utils.HandleUnstakeLock(rpcParameters, account, configurations, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for withdrawFunds function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for withdrawFunds function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for withdrawFunds function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestWithdraw(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var stakerId uint32

	type args struct {
		withdrawTxn *Types.Transaction
		withdrawErr error
		hash        common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When withdraw function executes successfully",
			args: args{
				withdrawTxn: &Types.Transaction{},
				withdrawErr: nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When InitiateWithdraw transaction fails",
			args: args{
				withdrawTxn: &Types.Transaction{},
				withdrawErr: errors.New("withdraw error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("withdraw error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			stakeManagerMock.On("InitiateWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawTxn, tt.args.withdrawErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.InitiateWithdraw(rpcParameters, txnOpts, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for withdraw function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for withdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for withdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestExecuteWithdraw(t *testing.T) {
	var config types.Configurations
	var flagSet *pflag.FlagSet
	var client *ethclient.Client

	type args struct {
		config       types.Configurations
		configErr    error
		address      string
		addressErr   error
		password     string
		stakerId     uint32
		stakerIdErr  error
		withdrawHash common.Hash
		withdrawErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteInitiateWithdraw executes successfully",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				stakerId:     1,
				withdrawHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				stakerId:     1,
				withdrawHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:       config,
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				stakerId:     1,
				withdrawHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error from withdraw funds",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				stakerId:     1,
				withdrawHash: core.NilHash,
				withdrawErr:  errors.New("withdrawFunds error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting stakerId",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				stakerId:     1,
				stakerIdErr:  errors.New("stakerId error"),
				withdrawHash: common.BigToHash(big.NewInt(1)),
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
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("AssignStakerId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("HandleUnstakeLock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteInitiateWithdraw(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteInitiateWithdraw function didn't execute as expected")
			}

		})
	}
}
