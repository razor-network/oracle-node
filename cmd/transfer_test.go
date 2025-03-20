package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

func TestTransfer(t *testing.T) {
	var config types.Configurations

	type args struct {
		amount        *big.Int
		decimalAmount *big.Float
		txnOptsErr    error
		transferTxn   *Types.Transaction
		transferErr   error
		transferHash  common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "When transfer function executes successfully",
			args: args{
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "When transfer transaction fails",
			args: args{
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				transferTxn:   &Types.Transaction{},
				transferErr:   errors.New("transfer error"),
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("transfer error"),
		},
		{
			name: "When there is an error in getting txnOpts",
			args: args{
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOptsErr:    errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("CheckAmountAndBalance", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.amount)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			utilsMock.On("GetAmountInDecimal", mock.AnythingOfType("*big.Int")).Return(tt.args.decimalAmount)
			tokenManagerMock.On("Transfer", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("common.Address"), mock.AnythingOfType("*big.Int")).Return(tt.args.transferTxn, tt.args.transferErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.transferHash)

			utils := &UtilsStruct{}

			got, err := utils.Transfer(rpcParameters, config, types.TransferInput{
				ValueInWei: big.NewInt(1).Mul(big.NewInt(1), big.NewInt(1e18)),
			})
			if got != tt.want {
				t.Errorf("Txn hash for transfer function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for transfer function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for transfer function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestExecuteTransfer(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	type args struct {
		config       types.Configurations
		configErr    error
		from         string
		fromErr      error
		to           string
		toErr        error
		password     string
		balance      *big.Int
		balanceErr   error
		amount       *big.Int
		amountErr    error
		transferErr  error
		transferHash common.Hash
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test1: When ExecuteTransfer function executes successfully",
			args: args{
				config:       config,
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:   nil,
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:   nil,
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting 'from' address from flags",
			args: args{
				config:       config,
				password:     "test",
				from:         "",
				fromErr:      errors.New("address error"),
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:   nil,
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting 'to' address from flags",
			args: args{
				config:       config,
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "",
				toErr:        errors.New("address error"),
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:   nil,
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in fetching balance",
			args: args{
				config:       config,
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balanceErr:   errors.New("balance error"),
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in fetching amount",
			args: args{
				config:       config,
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				amountErr:    errors.New("amount error"),
				transferErr:  nil,
				transferHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from Transfer function",
			args: args{
				config:       config,
				password:     "test",
				from:         "0x000000000000000000000000000000000000dea1",
				fromErr:      nil,
				to:           "0x000000000000000000000000000000000000dea2",
				toErr:        nil,
				balance:      big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:   nil,
				amount:       big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				transferErr:  errors.New("transfer error"),
				transferHash: core.NilHash,
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			setupTestEndpointsEnvironment()

			utilsMock.On("IsFlagPassed", mock.Anything).Return(false)
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringFrom", flagSet).Return(tt.args.from, tt.args.fromErr)
			flagSetMock.On("GetStringTo", flagSet).Return(tt.args.to, tt.args.toErr)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.amount, tt.args.amountErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("FetchBalance", mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)
			cmdUtilsMock.On("Transfer", mock.Anything, config, mock.Anything).Return(tt.args.transferHash, tt.args.transferErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteTransfer(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The ExecuteTransfer function didn't execute as expected")
			}

		})
	}
}
