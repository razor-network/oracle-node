package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/mock"
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
)

func TestTransfer(t *testing.T) {
	var client *ethclient.Client
	var config types.Configurations
	var transferInput types.TransferInput

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31000))

	type args struct {
		amount        *big.Int
		decimalAmount *big.Float
		txnOpts       *bind.TransactOpts
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
				txnOpts:       txnOpts,
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
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   errors.New("transfer error"),
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("transfer error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			tokenManangerUtilsMock := new(mocks.TokenManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			tokenManagerUtils = tokenManangerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("CheckAmountAndBalance", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.amount)
			utilsMock.On("GetTxnOpts", mock.Anything).Return(tt.args.txnOpts)
			utilsMock.On("GetAmountInDecimal", mock.AnythingOfType("*big.Int")).Return(tt.args.decimalAmount)
			tokenManangerUtilsMock.On("Transfer", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("common.Address"), mock.AnythingOfType("*big.Int")).Return(tt.args.transferTxn, tt.args.transferErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.transferHash)

			utils := &UtilsStruct{}

			got, err := utils.Transfer(client, config, transferInput)
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			flagsetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			flagSetUtils = flagsetUtilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			flagsetUtilsMock.On("GetStringFrom", flagSet).Return(tt.args.from, tt.args.fromErr)
			flagsetUtilsMock.On("GetStringTo", flagSet).Return(tt.args.to, tt.args.toErr)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.amount, tt.args.amountErr)
			utilsMock.On("AssignPassword", flagSet).Return()
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("FetchBalance", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.balance, tt.args.balanceErr)
			cmdUtilsMock.On("Transfer", mock.AnythingOfType("*ethclient.Client"), config, mock.AnythingOfType("types.TransferInput")).Return(tt.args.transferHash, tt.args.transferErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteTransfer(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The ExecuteTransfer function didn't execute as expected")
			}

		})
	}
}
