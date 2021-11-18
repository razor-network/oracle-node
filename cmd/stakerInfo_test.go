package cmd

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
	"testing"
)

func TestUtilsStruct_GetStakerInfo(t *testing.T) {
	var client *ethclient.Client
	stake, _ := new(big.Int).SetString("10000000000000000000000", 10)

	type fields struct {
		razorUtils        UtilsMock
		stakeManagerUtils StakeManagerMock
	}

	testUtils := fields{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
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
			utilsStruct := &UtilsStruct{
				razorUtils:        tt.fields.razorUtils,
				stakeManagerUtils: tt.fields.stakeManagerUtils,
			}

			GetOptionsMock = func(bool, string, string) bind.CallOpts {
				return tt.args.callOpts
			}

			StakerInfoMock = func(*ethclient.Client, *bind.CallOpts, uint32) (types.Staker, error) {
				return tt.args.stakerInfo, tt.args.stakerInfoErr
			}

			GetMaturityMock = func(*ethclient.Client, *bind.CallOpts, uint32) (uint16, error) {
				return tt.args.maturity, tt.args.maturityErr
			}

			GetInfluenceSnapshotMock = func(*ethclient.Client, string, uint32, uint32) (*big.Int, error) {
				return tt.args.influence, tt.args.influenceErr
			}

			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			err := utilsStruct.GetStakerInfo(tt.args.client, tt.args.stakerId)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for approve function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for approve function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}
