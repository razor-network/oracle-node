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

func TestExtendLock(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var extendLockInput types.ExtendLockInput
	var config types.Configurations
	var client *ethclient.Client

	type args struct {
		txnOpts      *bind.TransactOpts
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
				txnOpts:      txnOpts,
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
				txnOpts:      txnOpts,
				resetLockTxn: &Types.Transaction{},
				resetLockErr: errors.New("resetLock error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("resetLock error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			stakeManagerUtilsMockery = stakeManagerUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			stakeManagerUtilsMock.On("ExtendLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32")).Return(tt.args.resetLockTxn, tt.args.resetLockErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStructMockery{}

			got, err := utils.ExtendLock(client, config, extendLockInput)
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			flagSetUtilsMock := new(mocks.FlagSetInterfaceMockery)
			cmdUtilsMock := new(mocks.UtilsCmdInterfaceMockery)

			razorUtilsMockery = utilsMock
			flagSetUtilsMockery = flagSetUtilsMock
			cmdUtilsMockery = cmdUtilsMock

			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			flagSetUtilsMock.On("GetUint32StakerId", flagSet).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			cmdUtilsMock.On("ExtendLock", mock.AnythingOfType("*ethclient.Client"), config, mock.Anything).Return(tt.args.resetLockTxn, tt.args.resetLockErr)

			utils := &UtilsStructMockery{}
			fatal = false

			utils.ExecuteExtendLock(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteExtendLock function didn't execute as expected")
			}

		})
	}
}
