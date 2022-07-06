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
	"razor/core/types"
	"testing"
)

func TestUtilsStruct_ClaimCommission(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	type args struct {
		config     types.Configurations
		configErr  error
		password   string
		address    string
		addressErr error
		txn        *Types.Transaction
		err        error
		hash       common.Hash
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ClaimStakeReward runs successfully",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				txn:      &Types.Transaction{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When ClaimStakeReward fails",
			args: args{
				config:   types.Configurations{},
				password: "test",
				address:  "0x000000000000000000000000000000000000dead",
				txn:      nil,
				err:      errors.New("error in claiming stake reward"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching config",
			args: args{
				config:    types.Configurations{},
				configErr: errors.New("error in fetching config"),
				address:   "0x000000000000000000000000000000000000dead",
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fatal = false

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
			utilsMock.On("AssignPassword").Return(tt.args.password)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			stakeManagerUtilsMock.On("ClaimStakeReward", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.txn, tt.args.err)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(1)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			utils.ClaimCommission(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}
		})
	}
}
