package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestHandleUnstakeLock(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        big.NewInt(configurations.ChainId),
		Config:         configurations,
	}

	type args struct {
		state                    uint32
		stateErr                 error
		lock                     types.Locks
		lockErr                  error
		withdrawReleasePeriod    uint16
		withdrawReleasePeriodErr error
		txnOpts                  *bind.TransactOpts
		txnArgs                  types.TransactionOptions
		epoch                    uint32
		epochErr                 error
		time                     string
		withdrawHash             common.Hash
		withdrawErr              error
		blockCompletionErr       error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When withdrawFunds function executes successfully",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				txnArgs:                  txnArgs,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2:When there is an error in getting state",
			args: args{
				stateErr: errors.New("error in getting state"),
			},
			wantErr: errors.New("error in getting state"),
		},
		{
			name: "Test 3: When there is an error in getting locks",
			args: args{
				lockErr:                  errors.New("lock error"),
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawErr:              nil,
			},
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 4: When staker tries to withdraw without un-staking any Razors",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(0),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawErr:              nil,
			},
			wantErr: errors.New("unstake Razors before withdrawing"),
		},
		{
			name: "Test 5: When there is an error in getting withdrawReleasePeriod",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriodErr: errors.New("withdrawReleasePeriod error"),
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawErr:              nil,
			},
			wantErr: errors.New("withdrawReleasePeriod error"),
		},
		{
			name: "Test 6: When there is an error in getting epoch",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epochErr:                 errors.New("epoch error"),
				withdrawErr:              nil,
			},
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 7: When staker tries to withdraw when withdraw Initiation period has passed",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    9,
				epochErr:                 nil,
				withdrawErr:              nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 8: When staker tries to withdraw when withdraw Initiation period has not reached",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    3,
				epochErr:                 nil,
				time:                     "10 minutes 0 seconds ",
				withdrawErr:              nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 9: When there is an error in executing withdraw function",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    1,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawErr:              errors.New("withdraw error"),
			},
			wantErr: errors.New("withdraw error"),
		},
		{
			name: "Test 10: When staker tries to withdraw when withdrawal period has not reached",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    2,
				epochErr:                 nil,
				time:                     "20 minutes 0 seconds ",
				withdrawErr:              nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 10: When there is error in mining the transaction",
			args: args{
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				txnArgs:                  txnArgs,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
				blockCompletionErr:       errors.New("error in BlockCompletion for initiateWithdraw"),
			},
			wantErr: errors.New("error in BlockCompletion for initiateWithdraw"),
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

			cmdUtilsMock.On("WaitForAppropriateState", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.lock, tt.args.lockErr)
			utilsMock.On("GetWithdrawInitiationPeriod", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("InitiateWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.blockCompletionErr)
			utilsMock.On("SecondsToReadableTime", mock.AnythingOfType("int")).Return(tt.args.time)

			utils := &UtilsStruct{}
			_, err := utils.HandleUnstakeLock(client, account, configurations, stakerId)
			//if got != tt.want {
			//	t.Errorf("Txn hash for withdrawFunds function, got = %v, want = %v", got, tt.want)
			//}
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

			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			stakeManagerUtilsMock.On("InitiateWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawTxn, tt.args.withdrawErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.InitiateWithdraw(client, txnOpts, stakerId)
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
	var account types.Account

	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        big.NewInt(config.ChainId),
		Config:         config,
	}

	type args struct {
		config          types.Configurations
		configErr       error
		address         string
		addressErr      error
		password        string
		stakerId        uint32
		stakerIdErr     error
		autoWithdraw    bool
		autoWithdrawErr error
		txnArgs         types.TransactionOptions
		withdrawErr     error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteInitiateWithdraw executes successfully",
			args: args{
				config:          config,
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				stakerId:        1,
				txnArgs:         txnArgs,
				autoWithdraw:    true,
				autoWithdrawErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				config:          config,
				configErr:       errors.New("config error"),
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				stakerId:        1,
				autoWithdraw:    true,
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:          config,
				password:        "test",
				address:         "",
				addressErr:      errors.New("address error"),
				stakerId:        1,
				autoWithdraw:    true,
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error from withdraw funds",
			args: args{
				config:          config,
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				stakerId:        1,
				withdrawErr:     errors.New("withdrawFunds error"),
				autoWithdraw:    true,
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting stakerId",
			args: args{
				config:          config,
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				stakerId:        1,
				stakerIdErr:     errors.New("stakerId error"),
				autoWithdraw:    true,
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting autoWithdraw status",
			args: args{
				config:          types.Configurations{},
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				autoWithdraw:    true,
				autoWithdrawErr: errors.New("autoWithdraw error"),
				stakerId:        1,
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

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword").Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsMock.On("AssignStakerId", flagSet, mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			flagSetUtilsMock.On("GetBoolAutoWithdraw", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.autoWithdraw, tt.args.autoWithdrawErr)
			cmdUtilsMock.On("HandleUnstakeLock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.txnArgs, tt.args.withdrawErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(nil)
			cmdUtilsMock.On("AutoWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.autoWithdrawErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteInitiateWithdraw(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteInitiateWithdraw function didn't execute as expected")
			}

		})
	}
}

func TestAutoWithdraw(t *testing.T) {

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        big.NewInt(configurations.ChainId),
		Config:         configurations,
	}

	type args struct {
		withdrawLockErr    error
		lock               types.Locks
		lockErr            error
		stakerId           uint32
		txnArgs            types.TransactionOptions
		txn                common.Hash
		blockCompletionErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When AutoWithdraw function executes successfully",
			args: args{
				withdrawLockErr: nil,
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:  nil,
				stakerId: stakerId,
				txnArgs:  txnArgs,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error from withdrawFunds",
			args: args{
				withdrawLockErr: errors.New("withdrawFunds error"),
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:  nil,
				stakerId: stakerId,
				txnArgs:  txnArgs,
			},
			wantErr: errors.New("withdrawFunds error"),
		},
		{
			name: "Test 3: When there is an error in fetching Withdraw Lock",
			args: args{
				withdrawLockErr: nil,
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:  errors.New("error in fetching withdrawLock"),
				stakerId: stakerId,
				txnArgs:  txnArgs,
			},
			wantErr: errors.New("error in fetching withdrawLock"),
		},
		{
			name: "Test 4: When there is an error in mining block for AutoWithdraw",
			args: args{
				withdrawLockErr: nil,
				lock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				lockErr:            nil,
				stakerId:           stakerId,
				txnArgs:            txnArgs,
				txn:                common.BigToHash(big.NewInt(1)),
				blockCompletionErr: errors.New("error in BlockCompletion for AutoWithdraw"),
			},
			wantErr: errors.New("error in BlockCompletion for AutoWithdraw"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			timeMock := new(mocks.TimeInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			timeUtils = timeMock

			cmdUtilsMock.On("HandleWithdrawLock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.txn, tt.args.withdrawLockErr)
			timeMock.On("Sleep", mock.Anything).Return()
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.blockCompletionErr)
			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.lock, tt.args.lockErr)

			utils := &UtilsStruct{}
			gotErr := utils.AutoWithdraw(txnArgs, stakerId)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for AutoWithdraw function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for AutoWithdraw function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
