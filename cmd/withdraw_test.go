package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"testing"
)

func TestWithdrawFunds(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32
	var blockTime int64

	type args struct {
		lock                     types.Locks
		lockErr                  error
		withdrawReleasePeriod    uint8
		withdrawReleasePeriodErr error
		txnOpts                  *bind.TransactOpts
		epoch                    uint32
		epochErr                 error
		time                     string
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
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				lockErr:                  errors.New("lock error"),
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 3: When staker tries to withdraw without un-staking any Razors",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(0),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 4: When there is an error in getting withdrawReleasePeriod",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriodErr: errors.New("withdrawReleasePeriod error"),
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("withdrawReleasePeriod error"),
		},
		{
			name: "Test 5: When there is an error in getting epoch",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epochErr:                 errors.New("epoch error"),
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 6: When staker tries to withdraw when withdrawal period has passed",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    9,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 7: When staker tries to withdraw when withdrawal period has not reached",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    3,
				epochErr:                 nil,
				time:                     "10 minutes 0 seconds ",
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When there is an error in executing withdraw function",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    1,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawErr:              errors.New("withdraw error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("withdraw error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			cmdUtils = cmdUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.lock, tt.args.lockErr)
			utilsMock.On("GetWithdrawReleasePeriod", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("CalculateBlockTime", mock.AnythingOfType("*ethclient.Client")).Return(blockTime)
			utilsMock.On("SecondsToHuman", mock.AnythingOfType("int")).Return(tt.args.time)
			cmdUtilsMock.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)

			utils := &UtilsStruct{}
			got, err := utils.WithdrawFunds(client, account, configurations, stakerId)
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
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
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
			name: "Test 2: When Withdraw transaction fails",
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

			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			stakeManagerUtilsMock.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawTxn, tt.args.withdrawErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.Withdraw(client, txnOpts, stakerId)
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
			name: "Test 1: When ExecuteWithdraw executes successfully",
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			flagSetUtilsMock := new(mocks.FlagSetInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			flagSetUtils = flagSetUtilsMock

			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsMock.On("AssignStakerId", flagSet, mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("WithdrawFunds", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteWithdraw(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteWithdraw function didn't execute as expected")
			}

		})
	}
}
