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

func TestSetDelegation(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		config                       types.Configurations
		configErr                    error
		password                     string
		address                      string
		addressErr                   error
		status                       string
		statusErr                    error
		parseStatus                  bool
		parseStatusErr               error
		stakerId                     uint32
		stakerIdErr                  error
		staker                       bindings.StructsStaker
		stakerErr                    error
		commission                   uint8
		commissionErr                error
		SetDelegationAcceptanceTxn   *Types.Transaction
		SetDelegationAcceptanceErr   error
		hash                         common.Hash
		WaitForBlockCompletionStatus int
		executeUpdateCommissionErr   error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When SetDelegation function executes successfully",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				config:         config,
				configErr:      errors.New("config error"),
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("config error"),
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				addressErr:     errors.New("address error"),
				status:         "true",
				statusErr:      nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("address error"),
		},
		{
			name: "Test 4: When there is an error in getting status",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				statusErr:      errors.New("status error"),
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("status error"),
		},
		{
			name: "Test 5: When there is getting stakerId",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerIdErr:    errors.New("stakerId error"),
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("stakerId error"),
		},
		{
			name: "Test 7: When there is an error in parsing string status to bool",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "t",
				statusErr:      nil,
				parseStatusErr: errors.New("error in parsing status to bool"),
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("error in parsing status to bool"),
		},
		{
			name: "Test 8: When there is an error in getting staker",
			args: args{
				config:                       config,
				configErr:                    nil,
				password:                     "test",
				address:                      "0x000000000000000000000000000000000000dea1",
				addressErr:                   nil,
				status:                       "true",
				statusErr:                    nil,
				parseStatus:                  true,
				parseStatusErr:               nil,
				stakerId:                     1,
				stakerIdErr:                  nil,
				stakerErr:                    errors.New("staker error"),
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 10: When SetDelegationAcceptance transaction fails",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   errors.New("SetDelegationAcceptance error"),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("SetDelegationAcceptance error"),
		},
		{
			name: "Test 11: When there is an error in fetching commission",
			args: args{
				config:      config,
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				status:      "true",
				parseStatus: true,
				stakerId:    1,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				commissionErr:                errors.New("error in fetching commission"),
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("error in fetching commission"),
		},
		{
			name: "Test 12: When commission is non zero",
			args: args{
				config:      config,
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				status:      "true",
				parseStatus: true,
				stakerId:    1,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				commission:                   12,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: nil,
		},
		{
			name: "Test 13: When executeUpdateCommission throws an error",
			args: args{
				config:      config,
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				status:      "true",
				parseStatus: true,
				stakerId:    1,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
				},
				commission:                   12,
				executeUpdateCommissionErr:   errors.New("error in updating commission"),
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("error in updating commission"),
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

			GetStringStatusMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.status, tt.args.statusErr
			}

			GetUint8CommissionMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.commission, tt.args.commissionErr
			}

			ParseBoolMock = func(string) (bool, error) {
				return tt.args.parseStatus, tt.args.parseStatusErr
			}

			ConnectToClientMock = func(string2 string) *ethclient.Client {
				return client
			}

			GetStakerIdMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.stakerId, tt.args.stakerIdErr
			}

			GetStakerMock = func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error) {
				return tt.args.staker, tt.args.stakerErr
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}

			SetDelegationAcceptanceMock = func(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error) {
				return tt.args.SetDelegationAcceptanceTxn, tt.args.SetDelegationAcceptanceErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			ExecuteUpdateCommissionMock = func(*ethclient.Client, types.UpdateCommissionInput, UtilsStruct) error {
				return tt.args.executeUpdateCommissionErr
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return tt.args.WaitForBlockCompletionStatus
			}

			gotErr := utilsStruct.SetDelegation(flagSet)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for SetDelegation function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for SetDelegation function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
