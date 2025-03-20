package cmd

import (
	"errors"
	"razor/accounts"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestUpdateCommission(t *testing.T) {
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	type args struct {
		commission                       uint8
		stakerInfo                       bindings.StructsStaker
		stakerInfoErr                    error
		maxCommission                    uint8
		maxCommissionErr                 error
		epochLimitForUpdateCommission    uint16
		epochLimitForUpdateCommissionErr error
		epoch                            uint32
		epochErr                         error
		time                             string
		txnOptsErr                       error
		UpdateCommissionTxn              *Types.Transaction
		UpdateCommissionErr              error
		hash                             common.Hash
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When update commission executes successfully",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching staker info",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				stakerInfoErr:                 errors.New("error in fetching stakerInfo"),
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("error in fetching stakerInfo"),
		},
		{
			name: "Test 3: When there is an error in fetching max commission",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 0,
				maxCommissionErr:              errors.New("error in fetching max commission"),
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("error in fetching max commission"),
		},
		{
			name: "Test 4: When there is an error in fetching epochLimitForUpdateCommission",
			args: args{
				commission:                       10,
				stakerInfo:                       bindings.StructsStaker{},
				maxCommission:                    20,
				maxCommissionErr:                 nil,
				epochLimitForUpdateCommission:    0,
				epochLimitForUpdateCommissionErr: errors.New("error in fetching epochLimitForUpdateCommission"),
				epoch:                            11,
				UpdateCommissionTxn:              &Types.Transaction{},
				UpdateCommissionErr:              nil,
			},
			wantErr: errors.New("error in fetching epochLimitForUpdateCommission"),
		},
		{
			name: "Test 5: When there is an error in fetching epoch",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         0,
				epochErr:                      errors.New("error in fetching epoch"),
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("error in fetching epoch"),
		},
		{
			name: "Test 6: When update commission fails",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           errors.New("error in updating commission"),
			},
			wantErr: errors.New("error in updating commission"),
		},
		{
			name: "Test 7: When commission is 0",
			args: args{
				commission:                    0,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("commission out of range"),
		},
		{
			name: "Test 8: When commission is greater than max commission",
			args: args{
				commission:                    30,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("commission out of range"),
		},
		{
			name: "Test 9: When the epoch is invalid for update",
			args: args{
				commission: 10,
				stakerInfo: bindings.StructsStaker{
					EpochCommissionLastUpdated: 1,
				},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 100,
				epoch:                         1,
				time:                          "8 hours 25 minutes 0 second ",
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("invalid epoch for update"),
		},
		{
			name: "Test 10: When only one epoch is remaining for update commission",
			args: args{
				commission: 10,
				stakerInfo: bindings.StructsStaker{
					EpochCommissionLastUpdated: 1,
				},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 100,
				epoch:                         101,
				time:                          "5 minutes 0 second ",
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
			},
			wantErr: errors.New("invalid epoch for update"),
		},
		{
			name: "Test 11: When there is an error in getting txnOpts",
			args: args{
				commission:                    10,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				txnOptsErr:                    errors.New("txnOpts error"),
			},
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetStaker", mock.Anything, mock.Anything).Return(tt.args.stakerInfo, tt.args.stakerInfoErr)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			utilsMock.On("GetEpoch", mock.Anything).Return(tt.args.epoch, tt.args.epochErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)
			utilsMock.On("GetMaxCommission", mock.Anything, mock.Anything).Return(tt.args.maxCommission, tt.args.maxCommissionErr)
			utilsMock.On("GetEpochLimitForUpdateCommission", mock.Anything).Return(tt.args.epochLimitForUpdateCommission, tt.args.epochLimitForUpdateCommissionErr)
			utilsMock.On("SecondsToReadableTime", mock.AnythingOfType("int")).Return(tt.args.time)
			stakeManagerMock.On("UpdateCommission", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.UpdateCommissionTxn, tt.args.UpdateCommissionErr)

			utils := &UtilsStruct{}
			gotErr := utils.UpdateCommission(rpcParameters, config, types.UpdateCommissionInput{
				Commission: tt.args.commission,
			})
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for UpdateCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for UpdateCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteUpdateCommission(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	type args struct {
		config              types.Configurations
		configErr           error
		password            string
		address             string
		addressErr          error
		commission          uint8
		commissionErr       error
		stakerId            uint32
		stakerIdErr         error
		UpdateCommissionErr error
	}

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteUpdateCommission() executes successfully",
			args: args{
				config:     config,
				password:   "test",
				address:    "0x000000000000000000000000000000000000dea1",
				commission: 10,
				stakerId:   1,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in fetching config",
			args: args{
				config:              types.Configurations{},
				configErr:           errors.New("error in getting config"),
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          10,
				stakerId:            1,
				UpdateCommissionErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching address",
			args: args{
				config:              config,
				password:            "test",
				address:             "",
				addressErr:          errors.New("error in fetching address"),
				commission:          10,
				stakerId:            1,
				UpdateCommissionErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in fetching commission",
			args: args{
				config:              config,
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          0,
				commissionErr:       errors.New("error in fetching commission"),
				stakerId:            1,
				UpdateCommissionErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in fetching stakerId",
			args: args{
				config:              config,
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          10,
				stakerId:            0,
				stakerIdErr:         errors.New("error in fetching the stakerId"),
				UpdateCommissionErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from updateCommission",
			args: args{
				config:              config,
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          10,
				stakerId:            1,
				UpdateCommissionErr: errors.New("error in updating commission"),
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

			utilsMock.On("IsFlagPassed", mock.Anything).Return(true)
			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetUint8Commission", flagSet).Return(tt.args.commission, tt.args.commissionErr)
			utilsMock.On("GetStakerId", mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("UpdateCommission", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.UpdateCommissionErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteUpdateCommission(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteUpdateCommission function didn't execute as expected")
			}
		})
	}
}
