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
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"
	"time"
)

func TestUnstake(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var config types.Configurations
	var client *ethclient.Client
	var address string
	var password string
	var valueInWei *big.Int
	var stakerId uint32
	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		lock       types.Locks
		lockErr    error
		epoch      uint32
		epochErr   error
		unstakeTxn *Types.Transaction
		unstakeErr error
		hash       common.Hash
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When unstake function executes successfully",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				epoch:      5,
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				lockErr:    errors.New("lock error"),
				epoch:      5,
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				epochErr:   errors.New("epoch error"),
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When Unstake transaction fails",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				epoch:      5,
				unstakeErr: errors.New("unstake error"),
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("unstake error"),
		},
		{
			name: "Test 5: When there is an existing lock",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(1000),
				},
				epoch:      5,
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("existing lock"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetLockMock = func(*ethclient.Client, string, uint32) (types.Locks, error) {
				return tt.args.lock, tt.args.lockErr
			}

			WaitForAppropriateStateMock = func(*ethclient.Client, string, string, ...int) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}

			UnstakeContractMock = func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error) {
				return tt.args.unstakeTxn, tt.args.unstakeErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return 1
			}

			_, gotErr := Unstake(config, client,
				types.UnstakeInput{
					Address:    address,
					Password:   password,
					StakerId:   stakerId,
					ValueInWei: valueInWei,
				}, utilsStruct)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for Unstake function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for Unstake function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func Test_executeUnstake(t *testing.T) {

	var client *ethclient.Client
	var txnArgs types.TransactionOptions
	var flagSet *pflag.FlagSet

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		cmdUtils:          UtilsCmdMock{},
		flagSetUtils:      FlagSetMock{},
	}

	type args struct {
		config              types.Configurations
		configErr           error
		password            string
		address             string
		addressErr          error
		autoWithdraw        bool
		autoWithdrawErr     error
		value               *big.Int
		valueErr            error
		stakerId            uint32
		stakerIdErr         error
		lock                types.Locks
		lockErr             error
		unstakeErr          error
		autoWithdrawFuncErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When inputUnstake function executes successfully",
			args: args{
				config:       types.Configurations{},
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerId:     1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				configErr:    errors.New("config error"),
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerId:     1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:       types.Configurations{},
				password:     "test",
				addressErr:   errors.New("address error"),
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerId:     1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting autoWithdraw status",
			args: args{
				config:          types.Configurations{},
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				autoWithdrawErr: errors.New("autoWithdraw error"),
				value:           big.NewInt(10000),
				stakerId:        1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting stakerId",
			args: args{
				config:       types.Configurations{},
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerIdErr:  errors.New("stakerId error"),
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an existing lock",
			args: args{
				config:       types.Configurations{},
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerId:     1,
				lock: types.Locks{
					Amount: big.NewInt(1000),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from Unstake function",
			args: args{
				config:       types.Configurations{},
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				autoWithdraw: true,
				value:        big.NewInt(10000),
				stakerId:     1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: errors.New("unstake error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetConfigDataMock = func() (types.Configurations, error) {
				return tt.args.config, tt.args.configErr
			}

			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetBoolAutoWithdrawMock = func(*pflag.FlagSet) (bool, error) {
				return tt.args.autoWithdraw, tt.args.autoWithdrawErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			AssignAmountInWeiMock = func(*pflag.FlagSet) (*big.Int, error) {
				return tt.args.value, tt.args.valueErr
			}

			CheckEthBalanceIsZeroMock = func(*ethclient.Client, string) {

			}

			AssignStakerIdMock = func(*pflag.FlagSet, *ethclient.Client, string) (uint32, error) {
				return tt.args.stakerId, tt.args.stakerIdErr
			}

			GetLockMock = func(*ethclient.Client, string, uint32) (types.Locks, error) {
				return tt.args.lock, tt.args.lockErr
			}

			UnstakeMock = func(types.Configurations, *ethclient.Client, types.UnstakeInput, UtilsStruct) (types.TransactionOptions, error) {
				return txnArgs, tt.args.unstakeErr
			}

			AutoWithdrawMock = func(types.TransactionOptions, uint32, UtilsStruct) error {
				return tt.args.autoWithdrawFuncErr
			}

			fatal = false

			utilsStruct.executeUnstake(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The inputUnstake function didn't execute as expected")
			}

		})
	}
}

func TestAutoWithdraw(t *testing.T) {
	var txnArgs types.TransactionOptions
	var stakerId uint32

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
		cmdUtils:   UtilsCmdMock{},
	}

	type args struct {
		withdrawFundsHash common.Hash
		withdrawFundsErr  error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When AutoWithdraw function exceutes successfully",
			args: args{
				withdrawFundsHash: common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error from withdrawFunds",
			args: args{
				withdrawFundsErr: errors.New("withdrawFunds error"),
			},
			wantErr: errors.New("withdrawFunds error"),
		},
		{
			name: "Test 3: When withdrawFundsTxn is 0x00",
			args: args{
				withdrawFundsHash: core.NilHash,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withdrawFundsMock = func(*ethclient.Client, types.Account, types.Configurations, uint32, UtilsStruct) (common.Hash, error) {
				return tt.args.withdrawFundsHash, tt.args.withdrawFundsErr
			}

			SleepMock = func(time.Duration) {

			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return 1
			}
			gotErr := AutoWithdraw(txnArgs, stakerId, utilsStruct)
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
