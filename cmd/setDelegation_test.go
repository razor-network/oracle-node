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
	"razor/utils"
	"testing"
)

func TestDecreaseCommission(t *testing.T) {

	var client *ethclient.Client
	var stakerId uint32
	var commission uint8

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		decreaseCommissionPrompt     bool
		decreaseCommissionTxn        *Types.Transaction
		decreaseCommissionErr        error
		hash                         common.Hash
		WaitForBlockCompletionStatus int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When DecreaseCommission executes successfully",
			args: args{
				decreaseCommissionPrompt:     true,
				decreaseCommissionTxn:        &Types.Transaction{},
				decreaseCommissionErr:        nil,
				hash:                         common.BigToHash(big.NewInt(1)),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When DecreaseCommission transaction fails",
			args: args{
				decreaseCommissionPrompt:     true,
				decreaseCommissionTxn:        &Types.Transaction{},
				decreaseCommissionErr:        errors.New("decreaseCommission error"),
				hash:                         common.BigToHash(big.NewInt(1)),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("decreaseCommission error"),
		},
		{
			name: "Test 3: When decreaseCommissionPrompt is false",
			args: args{
				decreaseCommissionPrompt:     false,
				decreaseCommissionTxn:        &Types.Transaction{},
				decreaseCommissionErr:        nil,
				hash:                         common.BigToHash(big.NewInt(1)),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			DecreaseCommissionPromptMock = func() bool {
				return tt.args.decreaseCommissionPrompt
			}

			DecreaseCommissionContractMock = func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error) {
				return tt.args.decreaseCommissionTxn, tt.args.decreaseCommissionErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return tt.args.WaitForBlockCompletionStatus
			}

			gotErr := DecreaseCommission(client, stakerId, txnOpts, commission, utilsStruct)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for DecreaseCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for DecreaseCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func TestSetCommission(t *testing.T) {
	var client *ethclient.Client
	var stakerID uint32
	var commission uint8

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
	}

	type args struct {
		setCommissionTxn             *Types.Transaction
		setCommissionErr             error
		hash                         common.Hash
		WaitForBlockCompletionStatus int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When SetCommission function executes successfully",
			args: args{
				setCommissionTxn:             &Types.Transaction{},
				setCommissionErr:             nil,
				hash:                         common.BigToHash(big.NewInt(1)),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When SetCommission transaction fails",
			args: args{
				setCommissionTxn:             &Types.Transaction{},
				setCommissionErr:             errors.New("setCommission error"),
				hash:                         common.BigToHash(big.NewInt(1)),
				WaitForBlockCompletionStatus: 1,
			},
			wantErr: errors.New("setCommission error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCommissionContractMock = func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error) {
				return tt.args.setCommissionTxn, tt.args.setCommissionErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return tt.args.WaitForBlockCompletionStatus
			}

			gotErr := SetCommission(client, stakerID, txnOpts, commission, utilsStruct)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for SetCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for SetCommission function, got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

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
		commission                   uint8
		commissionErr                error
		parseStatus                  bool
		parseStatusErr               error
		stakerId                     uint32
		stakerIdErr                  error
		staker                       bindings.StructsStaker
		stakerErr                    error
		SetDelegationAcceptanceTxn   *Types.Transaction
		SetDelegationAcceptanceErr   error
		hash                         common.Hash
		WaitForBlockCompletionStatus int
		updatedStaker                bindings.StructsStaker
		updatedStakerErr             error
		SetCommissionErr             error
		DecreaseCommissionErr        error
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
			},
			wantErr: errors.New("stakerId error"),
		},
		{
			name: "Test 6: When there is an error in getting commission",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				commissionErr:  errors.New("commission error"),
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
			},
			wantErr: errors.New("commission error"),
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
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
				commission:                   5,
				commissionErr:                nil,
				parseStatus:                  true,
				parseStatusErr:               nil,
				stakerId:                     1,
				stakerIdErr:                  nil,
				stakerErr:                    errors.New("staker error"),
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: nil,
			},
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 9: When there is an error in getting updated staker",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				commission:     5,
				commissionErr:  nil,
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
				updatedStakerErr:             errors.New("updated staker"),
				SetCommissionErr:             nil,
				DecreaseCommissionErr:        nil,
			},
			wantErr: errors.New("updated staker"),
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
				commission:     5,
				commissionErr:  nil,
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
				updatedStakerErr:             nil,
				SetCommissionErr:             nil,
				DecreaseCommissionErr:        nil,
			},
			wantErr: errors.New("SetDelegationAcceptance error"),
		},
		{
			name: "Test 11: When SetCommission function fails",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				commission:     5,
				commissionErr:  nil,
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
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      errors.New("setCommission error"),
				DecreaseCommissionErr: nil,
			},
			wantErr: errors.New("setCommission error"),
		},
		{
			name: "Test 12: When DecreaseCommission function fails",
			args: args{
				config:         config,
				configErr:      nil,
				password:       "test",
				address:        "0x000000000000000000000000000000000000dea1",
				addressErr:     nil,
				status:         "true",
				statusErr:      nil,
				commission:     2,
				commissionErr:  nil,
				parseStatus:    true,
				parseStatusErr: nil,
				stakerId:       1,
				stakerIdErr:    nil,
				staker: bindings.StructsStaker{
					AcceptDelegation: false,
					Commission:       5,
				},
				stakerErr:                    nil,
				SetDelegationAcceptanceTxn:   &Types.Transaction{},
				SetDelegationAcceptanceErr:   nil,
				WaitForBlockCompletionStatus: 1,
				updatedStaker: bindings.StructsStaker{
					AcceptDelegation: true,
					Commission:       5,
				},
				updatedStakerErr:      nil,
				SetCommissionErr:      nil,
				DecreaseCommissionErr: errors.New("decreaseCommission error"),
			},
			wantErr: errors.New("decreaseCommission error"),
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

			GetTxnOptsMock = func(types.TransactionOptions, utils.Utils) *bind.TransactOpts {
				return txnOpts
			}

			SetDelegationAcceptanceMock = func(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error) {
				return tt.args.SetDelegationAcceptanceTxn, tt.args.SetDelegationAcceptanceErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return tt.args.WaitForBlockCompletionStatus
			}

			GetUpdatedStakerMock = func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error) {
				return tt.args.updatedStaker, tt.args.updatedStakerErr
			}

			SetCommissionMock = func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error {
				return tt.args.SetCommissionErr
			}

			DecreaseCommissionMock = func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error {
				return tt.args.DecreaseCommissionErr
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
