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
	"razor/pkg/bindings"
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
	var stakerId uint32

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		amount     *big.Int
		lock       types.Locks
		lockErr    error
		epoch      uint32
		epochErr   error
		staker     bindings.StructsStaker
		stakerErr  error
		sAmount    *big.Int
		sAmountErr error
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
				staker:     bindings.StructsStaker{},
				amount:     big.NewInt(1000),
				sAmount:    big.NewInt(1000),
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
				amount:     big.NewInt(1000),
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
		{
			name: "Test 6: When there is an error in getting staker",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				epoch:     5,
				stakerErr: errors.New("staker error"),
			},
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 7: When there is an error in getting sAmount",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				epoch:      5,
				staker:     bindings.StructsStaker{},
				amount:     big.NewInt(1000),
				sAmountErr: errors.New("sAmount error"),
			},
			wantErr: errors.New("sAmount error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetLockMock = func(*ethclient.Client, string, uint32) (types.Locks, error) {
				return tt.args.lock, tt.args.lockErr
			}

			GetStakerMock = func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error) {
				return tt.args.staker, tt.args.stakerErr
			}

			WaitForAppropriateStateMock = func(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}

			GetAmountInSRZRsMock = func(*ethclient.Client, string, bindings.StructsStaker, *big.Int, UtilsStruct) (*big.Int, error) {
				return tt.args.sAmount, tt.args.sAmountErr
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
					ValueInWei: tt.args.amount,
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

			GetConfigDataMock = func(UtilsStruct) (types.Configurations, error) {
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

			AssignAmountInWeiMock = func(*pflag.FlagSet, UtilsStruct) (*big.Int, error) {
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

func TestGetAmountInSRZRs(t *testing.T) {
	var client *ethclient.Client
	var address string
	var callOpts bind.CallOpts
	var stakedToken *bindings.StakedToken

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
	}

	type args struct {
		staker         bindings.StructsStaker
		amount         *big.Int
		balance        *big.Int
		balanceErr     error
		totalSupply    *big.Int
		totalSupplyErr error
		RZR            *big.Int
		decimalAmount  *big.Float
		sRZR           *big.Int
		sRZRErr        error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr error
	}{
		{
			name: "Test 1: When GetAmountInSRZRs executes successfully",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1000),
				},
				amount:        big.NewInt(1000),
				balance:       big.NewInt(1000),
				totalSupply:   big.NewInt(1000),
				RZR:           big.NewInt(1000),
				decimalAmount: big.NewFloat(1000),
				sRZR:          big.NewInt(1000),
			},
			want:    big.NewInt(1000),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting sRZR balance",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1000),
				},
				amount:     big.NewInt(1000),
				balanceErr: errors.New("sRZR balance error"),
			},
			want:    nil,
			wantErr: errors.New("sRZR balance error"),
		},
		{
			name: "Test 3: When there is an error in getting total supply of sRZR",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1000),
				},
				amount:         big.NewInt(1000),
				balance:        big.NewInt(1000),
				totalSupplyErr: errors.New("totalSupply error"),
			},
			want:    nil,
			wantErr: errors.New("totalSupply error"),
		},
		{
			name: "Test 4: When input amount exceeds total sRZR balance",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1000),
				},
				amount:        big.NewInt(2000),
				balance:       big.NewInt(1000),
				totalSupply:   big.NewInt(1000),
				RZR:           big.NewInt(1000),
				decimalAmount: big.NewFloat(1000),
				sRZR:          big.NewInt(1000),
			},
			want:    nil,
			wantErr: errors.New("invalid amount"),
		},
		{
			name: "Test 5: When there is an error in converting RZR's to sRZR's",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1000),
				},
				amount:        big.NewInt(1000),
				balance:       big.NewInt(1000),
				totalSupply:   big.NewInt(1000),
				RZR:           big.NewInt(1000),
				decimalAmount: big.NewFloat(1000),
				sRZRErr:       errors.New("conversion RZR to sRZR error"),
			},
			want:    nil,
			wantErr: errors.New("conversion RZR to sRZR error"),
		},
		{
			name: "Test 6: When the supply is high and GetAmountInSRZRs executes successfully",
			args: args{
				staker: bindings.StructsStaker{
					Stake: big.NewInt(1).Exp(big.NewInt(10), big.NewInt(9), nil),
				},
				amount:        big.NewInt(1).Exp(big.NewInt(10), big.NewInt(6), nil),
				balance:       big.NewInt(1).Exp(big.NewInt(10), big.NewInt(7), nil),
				totalSupply:   big.NewInt(1).Exp(big.NewInt(10), big.NewInt(9), nil),
				RZR:           big.NewInt(1).Exp(big.NewInt(10), big.NewInt(7), nil),
				decimalAmount: big.NewFloat(1).Mul(big.NewFloat(10), big.NewFloat(1e5)),
				sRZR:          big.NewInt(1).Exp(big.NewInt(10), big.NewInt(6), nil),
			},
			want:    big.NewInt(1).Exp(big.NewInt(10), big.NewInt(6), nil),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStakedTokenMock = func(*ethclient.Client, common.Address) *bindings.StakedToken {
				return stakedToken
			}

			GetOptionsMock = func() bind.CallOpts {
				return callOpts
			}

			BalanceOfMock = func(*bindings.StakedToken, *bind.CallOpts, common.Address) (*big.Int, error) {
				return tt.args.balance, tt.args.balanceErr
			}

			GetTotalSupplyMock = func(*bindings.StakedToken, *bind.CallOpts) (*big.Int, error) {
				return tt.args.totalSupply, tt.args.totalSupplyErr
			}

			ConvertSRZRToRZRMock = func(*big.Int, *big.Int, *big.Int) *big.Int {
				return tt.args.RZR
			}

			GetAmountInDecimalMock = func(*big.Int) *big.Float {
				return tt.args.decimalAmount
			}

			ConvertRZRToSRZRMock = func(*big.Int, *big.Int, *big.Int) (*big.Int, error) {
				return tt.args.sRZR, tt.args.sRZRErr
			}

			got, err := GetAmountInSRZRs(client, address, tt.args.staker, tt.args.amount, utilsStruct)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("GetAmountInSRZRs() = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetAmountInSRZRs function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetAmountInSRZRs function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
