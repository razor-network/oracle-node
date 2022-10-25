package cmd

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core/types"
	utilsPkgMocks "razor/utils/mocks"
	"testing"
)

func TestUtilsStruct_GetStakerInfo(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	stake, _ := new(big.Int).SetString("10000000000000000000000", 10)

	type fields struct {
		razorUtils        Utils
		stakeManagerUtils StakeManagerUtils
	}

	testUtils := fields{
		razorUtils:        Utils{},
		stakeManagerUtils: StakeManagerUtils{},
	}

	type args struct {
		client        *ethclient.Client
		stakerId      uint32
		callOpts      bind.CallOpts
		stakerInfo    types.Staker
		stakerInfoErr error
		maturity      uint16
		maturityErr   error
		influence     *big.Int
		influenceErr  error
		epoch         uint32
		epochErr      error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Test 1: When StakerInfo executes properly",
			fields: testUtils,
			args: args{
				client:   client,
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				stakerInfo: types.Staker{
					AcceptDelegation:                false,
					Commission:                      0,
					Address:                         common.HexToAddress("0x000000000000000000000000000000000000dead"),
					TokenAddress:                    common.HexToAddress("0x00000000000000000000000000000000deadcoin"),
					Id:                              1,
					Age:                             10000,
					EpochFirstStakedOrLastPenalized: 0,
					Stake:                           stake,
				},
				stakerInfoErr: nil,
				maturity:      uint16(70),
				maturityErr:   nil,
				influence:     big.NewInt(0),
				influenceErr:  nil,
				epoch:         1,
				epochErr:      nil,
			},
			wantErr: nil,
		},
		{
			name:   "Test 2: When there is error fetching staker info",
			fields: testUtils,
			args: args{
				client:   client,
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				stakerInfo: types.Staker{
					AcceptDelegation:                false,
					Commission:                      0,
					Address:                         common.Address{},
					TokenAddress:                    common.Address{},
					Id:                              0,
					Age:                             0,
					EpochFirstStakedOrLastPenalized: 0,
					Stake:                           nil,
				},
				stakerInfoErr: errors.New("error in fetching staker info"),
				maturity:      uint16(70),
				maturityErr:   nil,
				influence:     big.NewInt(0),
				influenceErr:  nil,
				epoch:         1,
				epochErr:      nil,
			},
			wantErr: errors.New("error in fetching staker info"),
		},
		{
			name:   "Test 3: When there is error fetching maturity",
			fields: testUtils,
			args: args{
				client:   client,
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				stakerInfo: types.Staker{
					AcceptDelegation:                false,
					Commission:                      0,
					Address:                         common.HexToAddress("0x000000000000000000000000000000000000dead"),
					TokenAddress:                    common.HexToAddress("0x00000000000000000000000000000000deadcoin"),
					Id:                              1,
					Age:                             10000,
					EpochFirstStakedOrLastPenalized: 0,
					Stake:                           stake,
				},
				stakerInfoErr: nil,
				maturity:      uint16(0),
				maturityErr:   errors.New("error in fetching maturity"),
				influence:     big.NewInt(0),
				influenceErr:  nil,
				epoch:         1,
				epochErr:      nil,
			},
			wantErr: errors.New("error in fetching maturity"),
		},
		{
			name:   "Test 4: When there is error fetching influence",
			fields: testUtils,
			args: args{
				client:   client,
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				stakerInfo: types.Staker{
					AcceptDelegation:                false,
					Commission:                      0,
					Address:                         common.HexToAddress("0x000000000000000000000000000000000000dead"),
					TokenAddress:                    common.HexToAddress("0x00000000000000000000000000000000deadcoin"),
					Id:                              1,
					Age:                             10000,
					EpochFirstStakedOrLastPenalized: 0,
					Stake:                           stake,
				},
				stakerInfoErr: nil,
				maturity:      uint16(70),
				maturityErr:   nil,
				influence:     big.NewInt(0),
				influenceErr:  errors.New("error in fetching influence"),
				epoch:         1,
				epochErr:      nil,
			},
			wantErr: errors.New("error in fetching influence"),
		},
		{
			name:   "Test 5: When there is error fetching epoch",
			fields: testUtils,
			args: args{
				client:   client,
				stakerId: 1,
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				stakerInfo: types.Staker{
					AcceptDelegation:                false,
					Commission:                      0,
					Address:                         common.HexToAddress("0x000000000000000000000000000000000000dead"),
					TokenAddress:                    common.HexToAddress("0x00000000000000000000000000000000deadcoin"),
					Id:                              1,
					Age:                             10000,
					EpochFirstStakedOrLastPenalized: 0,
					Stake:                           stake,
				},
				stakerInfoErr: nil,
				maturity:      uint16(70),
				maturityErr:   nil,
				influence:     big.NewInt(0),
				influenceErr:  nil,
				epoch:         0,
				epochErr:      errors.New("error in fetching epoch"),
			},
			wantErr: errors.New("error in fetching epoch"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(utilsPkgMocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerMock

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("StakerInfo", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.CallOpts"), mock.AnythingOfType("uint32")).Return(tt.args.stakerInfo, tt.args.stakerInfoErr)
			stakeManagerMock.On("GetMaturity", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.CallOpts"), mock.AnythingOfType("uint32")).Return(tt.args.maturity, tt.args.maturityErr)
			utilsMock.On("GetInfluenceSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.influence, tt.args.influenceErr)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utils := &UtilsStruct{}
			err := utils.GetStakerInfo(tt.args.client, tt.args.stakerId)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for StakerInfo function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for StakerInfo function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestUtilsStruct_ExecuteStakerinfo(t *testing.T) {
	var config types.Configurations
	var flagSet *pflag.FlagSet
	var client *ethclient.Client

	type args struct {
		config        types.Configurations
		configErr     error
		stakerId      uint32
		stakerIdErr   error
		stakerInfoErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1:  When ExecuteStakerinfo function executes successfully",
			args: args{
				config:        config,
				configErr:     nil,
				stakerId:      1,
				stakerIdErr:   nil,
				stakerInfoErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2:  When there is an error in getting config",
			args: args{
				config:        config,
				configErr:     errors.New("config error"),
				stakerId:      1,
				stakerIdErr:   nil,
				stakerInfoErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3:  When there is an error in getting stakerId",
			args: args{
				config:        config,
				configErr:     nil,
				stakerId:      1,
				stakerIdErr:   errors.New("stakerId error"),
				stakerInfoErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4:  When there is an error in getting GetStakerInfo function",
			args: args{
				config:        config,
				configErr:     nil,
				stakerId:      1,
				stakerIdErr:   nil,
				stakerInfoErr: errors.New("stakerInfo error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(utilsPkgMocks.Utils)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			flagSetUtilsMock := new(mocks.FlagSetInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			flagSetUtils = flagSetUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			flagSetUtilsMock.On("GetUint32StakerId", flagSet).Return(tt.args.stakerId, tt.args.stakerIdErr)
			cmdUtilsMock.On("GetStakerInfo", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.stakerInfoErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteStakerinfo(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteStakerinfo function didn't execute as expected")
			}

		})
	}
}
