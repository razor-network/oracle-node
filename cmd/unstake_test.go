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
	"razor/pkg/bindings"
	"testing"
)

func TestUnstake(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var config types.Configurations
	var client *ethclient.Client
	var address string
	var password string
	var stakerId uint32

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

			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			cmdUtils = cmdUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.lock, tt.args.lockErr)
			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			cmdUtilsMock.On("WaitForAppropriateState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("GetAmountInSRZRs", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.sAmount, tt.args.sAmountErr)
			stakeManagerUtilsMock.On("Unstake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unstakeTxn, tt.args.unstakeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			_, gotErr := utils.Unstake(config, client,
				types.UnstakeInput{
					Address:    address,
					Password:   password,
					StakerId:   stakerId,
					ValueInWei: tt.args.amount,
				})
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

func TestExecuteUnstake(t *testing.T) {

	var client *ethclient.Client
	var hash common.Hash
	var flagSet *pflag.FlagSet

	type args struct {
		config      types.Configurations
		configErr   error
		password    string
		address     string
		addressErr  error
		value       *big.Int
		valueErr    error
		stakerId    uint32
		stakerIdErr error
		lock        types.Locks
		lockErr     error
		unstakeErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When inputUnstake function executes successfully",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				value:    big.NewInt(10000),
				stakerId: 1,
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
				configErr: errors.New("config error"),
				password:  "test",
				address:   "0x000000000000000000000000000000000000dead",
				value:     big.NewInt(10000),
				stakerId:  1,
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
				config:     types.Configurations{},
				password:   "test",
				addressErr: errors.New("address error"),
				value:      big.NewInt(10000),
				stakerId:   1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting stakerId",
			args: args{
				config:      types.Configurations{},
				password:    "test",
				address:     "0x000000000000000000000000000000000000dead",
				value:       big.NewInt(10000),
				stakerIdErr: errors.New("stakerId error"),
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an existing lock",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				value:    big.NewInt(10000),
				stakerId: 1,
				lock: types.Locks{
					Amount: big.NewInt(1000),
				},
				unstakeErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from Unstake function",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				value:    big.NewInt(10000),
				stakerId: 1,
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

			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)
			flagSetUtilsMock := new(mocks.FlagSetInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			cmdUtils = cmdUtilsMock
			transactionUtils = transactionUtilsMock
			flagSetUtils = flagSetUtilsMock

			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.value, tt.args.valueErr)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsMock.On("AssignStakerId", flagSet, mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.lock, tt.args.lockErr)
			cmdUtilsMock.On("Unstake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(hash, tt.args.unstakeErr)
			utilsMock.On("WaitForBlockCompletion", client, mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteUnstake(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The ExecuteUnstake function didn't execute as expected")
			}

		})
	}
}

func TestAutoWithdraw(t *testing.T) {

	var txnArgs types.TransactionOptions
	var stakerId uint32

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
			name: "Test 1: When AutoWithdraw function executes successfully",
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

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			timeMock := new(mocks.TimeInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			timeUtils = timeMock

			cmdUtilsMock.On("HandleUnstakeLock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawFundsHash, tt.args.withdrawFundsErr)
			timeMock.On("Sleep", mock.Anything).Return()
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

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

func TestGetAmountInSRZRs(t *testing.T) {
	var client *ethclient.Client
	var address string
	var callOpts bind.CallOpts
	var stakedToken *bindings.StakedToken

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

			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock

			utilsMock.On("GetStakedToken", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("common.Address")).Return(stakedToken)
			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerUtilsMock.On("BalanceOf", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)
			stakeManagerUtilsMock.On("GetTotalSupply", mock.Anything, mock.Anything).Return(tt.args.totalSupply, tt.args.totalSupplyErr)
			utilsMock.On("ConvertSRZRToRZR", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.RZR)
			utilsMock.On("GetAmountInDecimal", mock.AnythingOfType("*big.Int")).Return(tt.args.decimalAmount)
			utilsMock.On("ConvertRZRToSRZR", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.sRZR, tt.args.sRZRErr)

			utils := &UtilsStruct{}

			got, err := utils.GetAmountInSRZRs(client, address, tt.args.staker, tt.args.amount)
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
