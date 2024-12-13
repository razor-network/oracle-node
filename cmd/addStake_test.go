package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestStakeCoins(t *testing.T) {
	txnArgs := types.TransactionOptions{
		Amount: big.NewInt(10000),
	}

	type args struct {
		txnArgs     types.TransactionOptions
		epoch       uint32
		getEpochErr error
		txnOptsErr  error
		stakeTxn    *Types.Transaction
		stakeErr    error
		hash        common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When stake transaction is successful",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				epoch:       2,
				getEpochErr: nil,
				stakeTxn:    &Types.Transaction{},
				stakeErr:    nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When waitForAppropriateState fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				epoch:       2,
				getEpochErr: errors.New("waitForAppropriateState error"),
				stakeTxn:    &Types.Transaction{},
				stakeErr:    nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForAppropriateState error"),
		},
		{
			name: "Test 3: When stake transaction fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				epoch:       2,
				getEpochErr: nil,
				stakeTxn:    &Types.Transaction{},
				stakeErr:    errors.New("stake error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("stake error"),
		},
		{
			name: "Test 4: When there is an error in getting transaction options",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				epoch:       2,
				getEpochErr: nil,
				txnOptsErr:  errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			utilsMock.On("GetEpoch", mock.Anything).Return(tt.args.epoch, tt.args.getEpochErr)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)
			stakeManagerMock.On("Stake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakeTxn, tt.args.stakeErr)

			utils := &UtilsStruct{}

			got, err := utils.StakeCoins(rpcParameters, txnArgs)
			if got != tt.want {
				t.Errorf("Txn hash for stake function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for stake function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for stake function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteStake(t *testing.T) {
	var flagSet *pflag.FlagSet
	var client *ethclient.Client
	var config types.Configurations

	type args struct {
		config          types.Configurations
		configErr       error
		password        string
		address         string
		addressErr      error
		balance         *big.Int
		balanceErr      error
		amount          *big.Int
		amountErr       error
		approveTxn      common.Hash
		approveErr      error
		minSafeRazor    *big.Int
		minSafeRazorErr error
		stakerId        uint32
		stakerIdErr     error
		staker          bindings.StructsStaker
		stakerErr       error
		stakeTxn        common.Hash
		stakeErr        error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteStake() executes successfully",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				staker:       bindings.StructsStaker{IsSlashed: false},
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
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
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
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
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting amount",
			args: args{
				config:     config,
				password:   "test",
				address:    "0x000000000000000000000000000000000000dead",
				amount:     nil,
				amountErr:  errors.New("amount error"),
				balance:    big.NewInt(10000),
				approveTxn: common.BigToHash(big.NewInt(1)),
				stakeTxn:   common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error from Approve",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				approveTxn:   core.NilHash,
				approveErr:   errors.New("approve error"),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from StakeCoins()",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     core.NilHash,
				stakeErr:     errors.New("stake error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error in getting balance",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				minSafeRazor: big.NewInt(0),
				stakerId:     1,
				balance:      nil,
				balanceErr:   errors.New("balance error"),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When stake value is less than minSafeRazor and staker has never staked",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(20),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(100),
				stakerId:     0,
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When stake value is less than minSafeRazor and staker's stake is more than the minSafeRazor already",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(20),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(100),
				stakerId:     1,
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 10: When the staker is slashed before",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(20),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(100),
				stakerId:     1,
				staker:       bindings.StructsStaker{IsSlashed: true},
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When stake value is less than minSafeRazor and staker's stake is more than the minSafeRazor already",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(20),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(100),
				stakerId:     1,
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: false,
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
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)
			utilsMock.On("FetchBalance", mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.amount, tt.args.amountErr)
			utilsMock.On("CheckAmountAndBalance", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.amount)
			utilsMock.On("CheckEthBalanceIsZero", mock.Anything, mock.Anything).Return()
			utilsMock.On("GetMinSafeRazor", mock.Anything).Return(tt.args.minSafeRazor, tt.args.minSafeRazorErr)
			utilsMock.On("GetStakerId", mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("GetStaker", mock.Anything, mock.Anything).Return(tt.args.staker, tt.args.stakerErr)
			cmdUtilsMock.On("Approve", mock.Anything, mock.Anything).Return(tt.args.approveTxn, tt.args.approveErr)
			cmdUtilsMock.On("StakeCoins", mock.Anything, mock.Anything).Return(tt.args.stakeTxn, tt.args.stakeErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteStake(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteStake function didn't execute as expected")
			}
		})
	}
}
