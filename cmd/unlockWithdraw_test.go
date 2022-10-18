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
	"reflect"
	"testing"
)

func TestExecuteUnlockWithdraw(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	type args struct {
		config      types.Configurations
		configErr   error
		password    string
		address     string
		addressErr  error
		stakerId    uint32
		stakerIdErr error
		txn         common.Hash
		err         error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteUnlockWithdraw executes successfully",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				stakerId: 1,
				txn:      common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in fetching config",
			args: args{
				config:    types.Configurations{},
				configErr: errors.New("error in fetching config"),
				address:   "0x000000000000000000000000000000000000dead",
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				addressErr: errors.New("error in getting address"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting stakerId",
			args: args{
				config:      types.Configurations{},
				password:    "test",
				address:     "0x000000000000000000000000000000000000dead",
				stakerIdErr: errors.New("error in getting stakerId"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in HandleWithdrawLock",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				stakerId: 1,
				err:      errors.New("error in HandleWithdrawLock"),
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
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			flagSetUtils = flagSetUtilsMock
			cmdUtils = cmdUtilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsMock.On("AssignStakerId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			cmdUtilsMock.On("HandleWithdrawLock", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.txn, tt.args.err)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(nil)
			utils := &UtilsStruct{}
			utils.ExecuteUnlockWithdraw(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}
		})
	}
}

func TestHandleWithdrawLock(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var (
		client         *ethclient.Client
		account        types.Account
		configurations types.Configurations
		stakerId       uint32
	)

	type args struct {
		withdrawLock      types.Locks
		withdrawLockErr   error
		epoch             uint32
		epochErr          error
		txnOpts           *bind.TransactOpts
		unlockWithdraw    common.Hash
		unlockWithdrawErr error
		time              string
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr bool
	}{
		{
			name: "Test 1: When HandleWithdrawLock executes successfully",
			args: args{
				withdrawLock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				epoch:          5,
				txnOpts:        txnOpts,
				unlockWithdraw: common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				withdrawLockErr: errors.New("error in getting lock"),
			},
			want:    core.NilHash,
			wantErr: true,
		},
		{
			name: "Test 3: When initiateWithdraw command is not called before unlocking razors",
			args: args{
				withdrawLock: types.Locks{
					UnlockAfter: big.NewInt(0),
				},
				epoch: 5,
			},
			want:    core.NilHash,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting epoch",
			args: args{
				withdrawLock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				epochErr: errors.New("error in getting epoch"),
			},
			want:    core.NilHash,
			wantErr: true,
		},
		{
			name: "Test 5: When withdrawLock is not reached",
			args: args{
				withdrawLock: types.Locks{
					UnlockAfter: big.NewInt(4),
				},
				epoch: 3,
				time:  "20 minutes 0 seconds",
			},
			want:    core.NilHash,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.withdrawLock, tt.args.withdrawLockErr)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("UnlockWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unlockWithdraw, tt.args.unlockWithdrawErr)
			utilsMock.On("SecondsToReadableTime", mock.AnythingOfType("int")).Return(tt.args.time)
			ut := &UtilsStruct{}
			got, err := ut.HandleWithdrawLock(client, account, configurations, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleWithdrawLock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleWithdrawLock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnlockWithdraw(t *testing.T) {
	var (
		client   *ethclient.Client
		txnOpts  *bind.TransactOpts
		stakerId uint32
	)

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
			name: "Test 1: When UnlockWithdraw executes successfully",
			args: args{
				txn:  &Types.Transaction{},
				hash: common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error UnlockWithdraw",
			args: args{
				txnErr: errors.New("error in unlockWithdraw"),
			},
			want:    core.NilHash,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			stakeManagerUtilsMock.On("UnlockWithdraw", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32")).Return(tt.args.txn, tt.args.txnErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			ut := &UtilsStruct{}
			got, err := ut.UnlockWithdraw(client, txnOpts, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnlockWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnlockWithdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
