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
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"
)

func TestUtilsStruct_UpdateCommission(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
		cmdUtils:          UtilsCmdMock{},
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
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When update commission executes successfully",
			args: args{
				config:              config,
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          10,
				stakerId:            1,
				UpdateCommissionErr: nil,
			},
			wantErr: nil,
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
			wantErr: errors.New("error in getting config"),
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
			wantErr: errors.New("error in fetching address"),
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
			wantErr: errors.New("error in fetching commission"),
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
			wantErr: errors.New("error in fetching the stakerId"),
		},
		{
			name: "Test 6: When there is an error in executing update commission",
			args: args{
				config:              config,
				password:            "test",
				address:             "0x000000000000000000000000000000000000dea1",
				commission:          10,
				stakerId:            0,
				UpdateCommissionErr: errors.New("error in ExecuteUpdateCommission"),
			},
			wantErr: errors.New("error in ExecuteUpdateCommission"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetConfigDataMock = func(UtilsStruct) (types.Configurations, error) {
				return tt.args.config, tt.args.configErr
			}
			AssignPasswordMock = func(set *pflag.FlagSet) string {
				return tt.args.password
			}
			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}
			GetUint8CommissionMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.commission, tt.args.commissionErr
			}
			GetStakerIdMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.stakerId, tt.args.stakerIdErr
			}
			ConnectToClientMock = func(string2 string) *ethclient.Client {
				return client
			}
			ExecuteUpdateCommissionMock = func(*ethclient.Client, types.UpdateCommissionInput, UtilsStruct) error {
				return tt.args.UpdateCommissionErr
			}
			gotErr := utilsStruct.UpdateCommission(flagSet)
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

func Test_executeUpdateCommission(t *testing.T) {

	var client *ethclient.Client
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		client                           *ethclient.Client
		input                            types.UpdateCommissionInput
		stakerInfo                       bindings.StructsStaker
		stakerInfoErr                    error
		maxCommission                    uint8
		maxCommissionErr                 error
		epochLimitForUpdateCommission    uint16
		epochLimitForUpdateCommissionErr error
		epoch                            uint32
		epochErr                         error
		UpdateCommissionTxn              *Types.Transaction
		UpdateCommissionErr              error
		WaitForBlockCompletionStatus     int
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
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				client:                        client,
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in fetching staker info",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				stakerInfoErr:                 errors.New("error in fetching stakerInfo"),
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("error in fetching stakerInfo"),
		},
		{
			name: "Test 3: When there is an error in fetching max commission",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 0,
				maxCommissionErr:              errors.New("error in fetching max commission"),
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("error in fetching max commission"),
		},
		{
			name: "Test 4: When there is an error in fetching epochLimitForUpdateCommission",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo:                       bindings.StructsStaker{},
				maxCommission:                    20,
				maxCommissionErr:                 nil,
				epochLimitForUpdateCommission:    0,
				epochLimitForUpdateCommissionErr: errors.New("error in fetching epochLimitForUpdateCommission"),
				epoch:                            11,
				UpdateCommissionTxn:              &Types.Transaction{},
				UpdateCommissionErr:              nil,
				WaitForBlockCompletionStatus:     1,
			},
			wantErr: errors.New("error in fetching epochLimitForUpdateCommission"),
		},
		{
			name: "Test 5: When there is an error in fetching epoch",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         0,
				epochErr:                      errors.New("error in fetching epoch"),
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("error in fetching epoch"),
		},
		{
			name: "Test 6: When update commission fails",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           errors.New("error in updating commission"),
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("error in updating commission"),
		},
		{
			name: "Test 7: When commission is 0",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 0,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("commission out of range"),
		},
		{
			name: "Test 8: When commission is greater than max commission",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 30,
					Config:     config,
				},
				stakerInfo:                    bindings.StructsStaker{},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 10,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("commission out of range"),
		},
		{
			name: "Test 9: When the epoch is invalid for update",
			args: args{
				input: types.UpdateCommissionInput{
					StakerId:   1,
					Address:    "0x000000000000000000000000000000000000dea1",
					Password:   "test",
					Commission: 10,
					Config:     config,
				},
				stakerInfo: bindings.StructsStaker{
					EpochCommissionLastUpdated: 1,
				},
				maxCommission:                 20,
				epochLimitForUpdateCommission: 100,
				epoch:                         11,
				UpdateCommissionTxn:           &Types.Transaction{},
				UpdateCommissionErr:           nil,
				WaitForBlockCompletionStatus:  1,
			},
			wantErr: errors.New("invalid epoch for update"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}
			GetStakerMock = func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error) {
				return tt.args.stakerInfo, tt.args.stakerInfoErr
			}
			GetMaxCommissionMock = func(*ethclient.Client) (uint8, error) {
				return tt.args.maxCommission, tt.args.maxCommissionErr
			}
			GetEpochLimitForUpdateCommissionMock = func(*ethclient.Client) (uint16, error) {
				return tt.args.epochLimitForUpdateCommission, tt.args.epochLimitForUpdateCommissionErr
			}
			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}
			UpdateCommissionMock = func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error) {
				return tt.args.UpdateCommissionTxn, tt.args.UpdateCommissionErr
			}
			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}
			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return tt.args.WaitForBlockCompletionStatus
			}
			gotErr := ExecuteUpdateCommission(tt.args.client, tt.args.input, utilsStruct)
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
