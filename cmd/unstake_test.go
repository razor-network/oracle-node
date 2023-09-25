package cmd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestUnstake(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var config types.Configurations
	var client *ethclient.Client
	var address string
	var password string
	var stakerId uint32

	type args struct {
		staker         bindings.StructsStaker
		stakerErr      error
		approveHash    common.Hash
		approveHashErr error
		amount         *big.Int
		lock           types.Locks
		lockErr        error
		state          uint32
		stateErr       error
		unstakeTxn     *Types.Transaction
		unstakeErr     error
		hash           common.Hash
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
				amount:     big.NewInt(1000),
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				lockErr:    errors.New("lock error"),
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 4: When Unstake transaction fails",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
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
				unstakeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("existing unstake lock"),
		},
		{
			name: "Test 6: When there is an error in getting staker",
			args: args{
				stakerErr: errors.New("error in getting staker"),
			},
			wantErr: errors.New("error in getting staker"),
		},
		{
			name: "Test 7: When there is an error in getting approveHash",
			args: args{
				staker:         bindings.StructsStaker{},
				approveHashErr: errors.New("error in getting approveHash"),
			},
			wantErr: errors.New("error in getting approveHash"),
		},
		{
			name: "Test 8: When approveHash is not nil",
			args: args{
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				amount:      big.NewInt(1000),
				staker:      bindings.StructsStaker{},
				approveHash: common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.staker, tt.args.stakerErr)
			cmdUtilsMock.On("ApproveUnstake", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.approveHash, tt.args.approveHashErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(nil)
			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.lock, tt.args.lockErr)
			cmdUtilsMock.On("WaitForAppropriateState", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			stakeManagerMock.On("Unstake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unstakeTxn, tt.args.unstakeErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

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
		unstakeHash common.Hash
		unstakeErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteUnstake function executes successfully",
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
			expectedFatal: false,
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
		{
			name: "Test 6: When ExecuteUnstake function executes successfully and WaitForBlockCompletion executes",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				value:    big.NewInt(10000),
				stakerId: 1,
				lock: types.Locks{
					Amount: big.NewInt(0),
				},
				unstakeHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything, mock.Anything).Return(nil)
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.value, tt.args.valueErr)
			utilsMock.On("AssignStakerId", flagSet, mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.lock, tt.args.lockErr)
			cmdUtilsMock.On("Unstake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unstakeHash, tt.args.unstakeErr)
			utilsMock.On("WaitForBlockCompletion", client, mock.AnythingOfType("string")).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteUnstake(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The ExecuteUnstake function didn't execute as expected")
			}

		})
	}
}

func TestApproveUnstake(t *testing.T) {
	var (
		client             *ethclient.Client
		stakerTokenAddress common.Address
		txnArgs            types.TransactionOptions
	)

	privateKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	type args struct {
		txn    *Types.Transaction
		txnErr error
		hash   common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr bool
	}{
		{
			name: "Test 1: When ApproveUnstake executes successfully",
			args: args{
				txn: &Types.Transaction{},
			},
			want:    core.NilHash,
			wantErr: false,
		},
		{
			name: "Test 1: When there is an error in getting transaction",
			args: args{
				txnErr: errors.New("error in getting transaction"),
			},
			want:    core.NilHash,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			stakeManagerMock.On("ApproveUnstake", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.txn, tt.args.txnErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)
			ut := &UtilsStruct{}
			got, err := ut.ApproveUnstake(client, stakerTokenAddress, txnArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApproveUnstake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApproveUnstake() got = %v, want %v", got, tt.want)
			}
		})
	}
}
