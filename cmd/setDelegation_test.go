package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestSetDelegation(t *testing.T) {
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
	}

	type args struct {
		status                     bool
		staker                     bindings.StructsStaker
		stakerErr                  error
		txnOptsErr                 error
		setDelegationAcceptanceTxn *Types.Transaction
		setDelegationAcceptanceErr error
		hash                       common.Hash
		commission                 uint8
		UpdateCommissionErr        error
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When SetDelegation function executes successfully",
			args: args{
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				stakerErr:                  nil,
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When setDelegationAcceptance transaction fails",
			args: args{
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				stakerErr:                  nil,
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: errors.New("SetDelegationAcceptance error"),
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("SetDelegationAcceptance error"),
		},
		{
			name: "Test 3: When there is an error in getting staker",
			args: args{
				stakerErr:                  errors.New("staker error"),
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 4: When stakerInfo.AcceptDelegation == delegationInput.Status",
			args: args{
				status: true,
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				stakerErr:                  nil,
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 5: When commission is non zero and UpdateCommission executes successfully",
			args: args{
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				stakerErr:                  nil,
				commission:                 10,
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: nil,
				hash:                       common.BigToHash(big.NewInt(1)),
				UpdateCommissionErr:        nil,
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 6: When commission is non zero and UpdateCommission does not executes successfully",
			args: args{
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				stakerErr:                  nil,
				commission:                 10,
				setDelegationAcceptanceTxn: &Types.Transaction{},
				setDelegationAcceptanceErr: nil,
				hash:                       common.BigToHash(big.NewInt(1)),
				UpdateCommissionErr:        errors.New("error in updating commission"),
			},
			want:    core.NilHash,
			wantErr: errors.New("error in updating commission"),
		},
		{
			name: "Test 7: When there is an error in getting txnOpts",
			args: args{
				staker: bindings.StructsStaker{
					AcceptDelegation: true,
				},
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			SetUpMockInterfaces()

			utilsMock.On("GetStaker", mock.Anything, mock.Anything).Return(tt.args.staker, tt.args.stakerErr)
			cmdUtilsMock.On("UpdateCommission", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.UpdateCommissionErr)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			stakeManagerMock.On("SetDelegationAcceptance", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.AnythingOfType("bool")).Return(tt.args.setDelegationAcceptanceTxn, tt.args.setDelegationAcceptanceErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.SetDelegation(rpcParameters, config, types.SetDelegationInput{
				Status:     tt.args.status,
				Commission: tt.args.commission,
			})
			if got != tt.want {
				t.Errorf("Txn hash for setDelegation function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for setDelegation function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for setDelegation function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteSetDelegation(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config = types.Configurations{
		Provider:      "127.0.0.1",
		GasMultiplier: 1,
		BufferPercent: 20,
		WaitTime:      1,
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
		commission                   uint8
		commissionErr                error
		SetDelegationAcceptanceTxn   *Types.Transaction
		SetDelegationAcceptanceErr   error
		setDelegationHash            common.Hash
		setDelegationErr             error
		WaitForBlockCompletionStatus int
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When SetDelegation function executes successfully",
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
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				config:                       config,
				configErr:                    errors.New("config error"),
				password:                     "test",
				address:                      "0x000000000000000000000000000000000000dea1",
				addressErr:                   nil,
				status:                       "true",
				statusErr:                    nil,
				parseStatus:                  true,
				parseStatusErr:               nil,
				stakerId:                     1,
				stakerIdErr:                  nil,
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:                       config,
				configErr:                    nil,
				password:                     "test",
				addressErr:                   errors.New("address error"),
				status:                       "true",
				statusErr:                    nil,
				parseStatus:                  true,
				parseStatusErr:               nil,
				stakerId:                     1,
				stakerIdErr:                  nil,
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting status",
			args: args{
				config:                       config,
				configErr:                    nil,
				password:                     "test",
				address:                      "0x000000000000000000000000000000000000dea1",
				addressErr:                   nil,
				statusErr:                    errors.New("status error"),
				parseStatus:                  true,
				parseStatusErr:               nil,
				stakerId:                     1,
				stakerIdErr:                  nil,
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is getting stakerId",
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
				stakerIdErr:                  errors.New("stakerId error"),
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error in parsing string status to bool",
			args: args{
				config:                       config,
				configErr:                    nil,
				password:                     "test",
				address:                      "0x000000000000000000000000000000000000dea1",
				addressErr:                   nil,
				status:                       "t",
				statusErr:                    nil,
				parseStatusErr:               errors.New("error in parsing status to bool"),
				stakerId:                     1,
				stakerIdErr:                  nil,
				setDelegationHash:            common.BigToHash(big.NewInt(1)),
				setDelegationErr:             nil,
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there is an error from SetDelegation function",
			args: args{
				config:                       config,
				configErr:                    nil,
				password:                     "test",
				address:                      "0x000000000000000000000000000000000000dea1",
				addressErr:                   nil,
				status:                       "t",
				statusErr:                    nil,
				parseStatusErr:               errors.New("error in parsing status to bool"),
				stakerId:                     1,
				stakerIdErr:                  nil,
				setDelegationHash:            core.NilHash,
				setDelegationErr:             errors.New("setDelegation error"),
				WaitForBlockCompletionStatus: 1,
			},
			expectedFatal: true,
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
			expectedFatal: true,
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
			expectedFatal: false,
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
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			flagSetMock.On("GetStringStatus", flagSet).Return(tt.args.status, tt.args.statusErr)
			flagSetMock.On("GetUint8Commission", flagSet).Return(tt.args.commission, tt.args.commissionErr)
			stringMock.On("ParseBool", mock.AnythingOfType("string")).Return(tt.args.parseStatus, tt.args.parseStatusErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("GetStakerId", mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			cmdUtilsMock.On("SetDelegation", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.setDelegationHash, tt.args.setDelegationErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteSetDelegation(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteSetDelegation function didn't execute as expected")
			}
		})
	}
}
