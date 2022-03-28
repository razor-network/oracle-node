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
	"razor/utils"
	mocks2 "razor/utils/mocks"
	"testing"
)

func TestStakeCoins(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	txnArgs := types.TransactionOptions{
		Amount: big.NewInt(10000),
	}

	type args struct {
		txnArgs     types.TransactionOptions
		txnOpts     *bind.TransactOpts
		epoch       uint32
		getEpochErr error
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
				txnOpts:     txnOpts,
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
				txnOpts:     txnOpts,
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
				txnOpts:     txnOpts,
				epoch:       2,
				getEpochErr: nil,
				stakeTxn:    &Types.Transaction{},
				stakeErr:    errors.New("stake error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("stake error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.getEpochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			stakeManagerUtilsMock.On("Stake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakeTxn, tt.args.stakeErr)

			utils := &UtilsStruct{}

			got, err := utils.StakeCoins(txnArgs)
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
		config           types.Configurations
		configErr        error
		password         string
		address          string
		addressErr       error
		balance          *big.Int
		balanceErr       error
		amount           *big.Int
		amountErr        error
		approveTxn       common.Hash
		approveErr       error
		minSafeRazor     *big.Int
		minSafeRazorErr  error
		stakeTxn         common.Hash
		stakeErr         error
		isFlagPassed     bool
		autoVote         bool
		autoVoteErr      error
		isRogue          bool
		isRogueErr       error
		rogueMode        []string
		rogueModeErr     error
		voteErr          error
		revealedDataMaps types.RevealedDataMaps
		//revealedDataErr  error
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
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
				revealedDataMaps: types.RevealedDataMaps{
					SortedRevealedValues: nil,
					VoteWeights:          nil,
					InfluenceSum:         nil,
				},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When autoVote flag is passed and ExecuteStake() executes successfully",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: true,
				autoVote:     true,
				isRogue:      true,
				rogueMode:    []string{"propose"},
				voteErr:      nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 3: When there is an error in getting config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting address",
			args: args{
				config:       config,
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting amount",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       nil,
				amountErr:    errors.New("amount error"),
				balance:      big.NewInt(10000),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from Approve",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   core.NilHash,
				approveErr:   errors.New("approve error"),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from StakeCoins()",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     core.NilHash,
				stakeErr:     errors.New("stake error"),
				isFlagPassed: false,
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there is an error in getting autoVote status",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: true,
				autoVoteErr:  errors.New("autoVote error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When there is an error in getting rogue status",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: true,
				autoVote:     true,
				isRogueErr:   errors.New("rogue error"),
				rogueMode:    []string{"propose"},
				voteErr:      nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 10: When there is an error in getting rogue modes",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				balance:      big.NewInt(10000),
				minSafeRazor: big.NewInt(0),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: true,
				autoVote:     true,
				isRogue:      true,
				rogueMode:    nil,
				rogueModeErr: errors.New("rogueModes error"),
				voteErr:      nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 11: When there is an error in getting balance",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				amount:       big.NewInt(2000),
				minSafeRazor: big.NewInt(0),
				balance:      nil,
				balanceErr:   errors.New("balance error"),
				approveTxn:   common.BigToHash(big.NewInt(1)),
				stakeTxn:     common.BigToHash(big.NewInt(2)),
				isFlagPassed: false,
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
			utilsPkgMock := new(mocks2.Utils)

			razorUtils = utilsMock
			flagSetUtils = flagSetUtilsMock
			cmdUtils = cmdUtilsMock
			utils.UtilsInterface = utilsPkgMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			utilsMock.On("FetchBalance", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.balance, tt.args.balanceErr)
			cmdUtilsMock.On("AssignAmountInWei", flagSet).Return(tt.args.amount, tt.args.amountErr)
			utilsMock.On("CheckAmountAndBalance", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(tt.args.amount)
			utilsMock.On("CheckEthBalanceIsZero", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return()
			utilsPkgMock.On("GetMinSafeRazor", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minSafeRazor, tt.args.minSafeRazorErr)
			cmdUtilsMock.On("Approve", mock.Anything).Return(tt.args.approveTxn, tt.args.approveErr)
			cmdUtilsMock.On("StakeCoins", mock.Anything).Return(tt.args.stakeTxn, tt.args.stakeErr)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
			flagSetUtilsMock.On("GetBoolAutoVote", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.autoVote, tt.args.autoVoteErr)
			flagSetUtilsMock.On("GetBoolRogue", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.isRogue, tt.args.isRogueErr)
			flagSetUtilsMock.On("GetStringSliceRogueMode", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.rogueMode, tt.args.rogueModeErr)
			cmdUtilsMock.On("Vote", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.voteErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteStake(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteStake function didn't execute as expected")
			}
		})
	}
}
