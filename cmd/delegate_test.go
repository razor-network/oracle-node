package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestDelegate(t *testing.T) {
	var stakerId uint32 = 1

	type args struct {
		amount      *big.Int
		txnOptsErr  error
		delegateTxn *Types.Transaction
		delegateErr error
		hash        common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When delegate function executes successfully",
			args: args{
				amount:      big.NewInt(1000),
				delegateTxn: &Types.Transaction{},
				delegateErr: nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When delegate transaction fails",
			args: args{
				amount:      big.NewInt(1000),
				delegateTxn: &Types.Transaction{},
				delegateErr: errors.New("delegate error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("delegate error"),
		},
		{
			name: "Test 3: When there is an error in getting txnOpts",
			args: args{
				amount:     big.NewInt(1000),
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			stakeManagerMock.On("Delegate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.delegateTxn, tt.args.delegateErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.Delegate(rpcParameters, types.TransactionOptions{
				Amount: tt.args.amount,
			}, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for delegate function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for delegate function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for delegate function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteDelegate(t *testing.T) {
	var config types.Configurations
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	type args struct {
		config       types.Configurations
		configErr    error
		address      string
		addressErr   error
		password     string
		stakerId     uint32
		stakerIdErr  error
		balance      *big.Int
		balanceErr   error
		amount       *big.Int
		amountErr    error
		approveTxn   common.Hash
		approveErr   error
		delegateHash common.Hash
		delegateErr  error
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteDelegate() executes successfully",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				approveErr:   errors.New("address error"),
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting stakerId",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				stakerIdErr:  errors.New("stakerId error"),
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting balance",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				balanceErr:   errors.New("balance error"),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting amount",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				amountErr:    errors.New("amount error"),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from approve",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   core.NilHash,
				amountErr:    errors.New("approve error"),
				delegateHash: common.BigToHash(big.NewInt(2)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there is an error from delegate",
			args: args{
				config:       config,
				address:      "0x000000000000000000000000000000000000dead",
				password:     "test",
				stakerId:     2,
				balance:      big.NewInt(10000),
				amount:       big.NewInt(2000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				delegateHash: core.NilHash,
				delegateErr:  errors.New("delegate error"),
			},
			expectedFatal: true,
		},
	}
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
			flagSetMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			flagSetMock.On("GetUint32StakerId", flagSet).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)
			utilsMock.On("FetchBalance", mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.amount, tt.args.amountErr)
			utilsMock.On("CheckAmountAndBalance", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.amount)
			utilsMock.On("CheckEthBalanceIsZero", mock.Anything, mock.Anything).Return()
			cmdUtilsMock.On("Approve", mock.Anything, mock.Anything).Return(tt.args.approveTxn, tt.args.approveErr)
			cmdUtilsMock.On("Delegate", mock.Anything, mock.Anything, mock.AnythingOfType("uint32")).Return(tt.args.delegateHash, tt.args.delegateErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteDelegate(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteDelegate function didn't execute as expected")
			}

		})
	}
}
