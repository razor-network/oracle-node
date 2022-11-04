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
	utilsPkgMocks "razor/utils/mocks"
	"testing"
)

func TestUtilsStruct_ClaimCommission(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var callOpts bind.CallOpts

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	type args struct {
		config        types.Configurations
		configErr     error
		password      string
		address       string
		addressErr    error
		stakerInfo    types.Staker
		stakerInfoErr error
		stakerId      uint32
		stakerIdErr   error
		callOpts      bind.CallOpts
		txn           *Types.Transaction
		err           error
		hash          common.Hash
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ClaimStakerReward runs successfully",
			args: args{
				config:   types.Configurations{},
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				stakerInfo: types.Staker{
					StakerReward: big.NewInt(100),
				},
				stakerInfoErr: nil,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				txn:           &Types.Transaction{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in fetching staker id",
			args: args{
				config:      types.Configurations{},
				stakerId:    0,
				stakerIdErr: errors.New("error in getting staker id"),
				stakerInfo: types.Staker{
					StakerReward: big.NewInt(0),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				password: "test",
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching config",
			args: args{
				config:    types.Configurations{},
				configErr: errors.New("error in fetching config"),
				address:   "0x000000000000000000000000000000000000dead",
				stakerInfo: types.Staker{
					StakerReward: big.NewInt(0),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
			},

			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in fetching stakerInfo",
			args: args{
				config: types.Configurations{},
				stakerInfo: types.Staker{
					Address:      common.Address{},
					TokenAddress: common.Address{},
					Stake:        nil,
					StakerReward: big.NewInt(0),
				},
				stakerInfoErr: errors.New("error in fetching staker info"),
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				stakerId:    1,
				stakerIdErr: nil,
				password:    "test",
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in claiming stake reward",
			args: args{
				config: types.Configurations{},
				stakerInfo: types.Staker{
					Address:      common.Address{},
					TokenAddress: common.Address{},
					Stake:        nil,
					StakerReward: big.NewInt(100),
				},
				stakerInfoErr: nil,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				stakerId:    1,
				stakerIdErr: nil,
				password:    "test",
				err:         errors.New("error in claiming stake reward"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in mining block",
			args: args{
				config: types.Configurations{},
				stakerInfo: types.Staker{
					Address:      common.Address{},
					TokenAddress: common.Address{},
					Stake:        nil,
					StakerReward: big.NewInt(100),
				},
				stakerInfoErr: nil,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				stakerId:    1,
				stakerIdErr: nil,
				password:    "test",
				err:         errors.New("error in wait for blockCompletion for claim commission"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is no commission to claim",
			args: args{
				config: types.Configurations{},
				stakerInfo: types.Staker{
					Address:      common.Address{},
					TokenAddress: common.Address{},
					Stake:        nil,
					StakerReward: big.NewInt(0),
				},
				stakerInfoErr: nil,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
				},
				stakerId:    1,
				stakerIdErr: nil,
				password:    "test",
				err:         errors.New("no commission to claim"),
			},
			expectedFatal: false,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fatal = false

			utilsMock := new(utilsPkgMocks.Utils)
			flagSetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			flagSetUtils = flagSetUtilsMock
			cmdUtils = cmdUtilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			utilsMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("GetOptions").Return(callOpts)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(nil)

			stakeManagerUtilsMock.On("StakerInfo", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.CallOpts"), mock.AnythingOfType("uint32")).Return(tt.args.stakerInfo, tt.args.stakerInfoErr)
			stakeManagerUtilsMock.On("ClaimStakerReward", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.txn, tt.args.err)

			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			utils.ClaimCommission(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}
		})
	}
}
