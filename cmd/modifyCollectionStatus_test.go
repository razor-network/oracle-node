package cmd

import (
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestModifyAssetStatus(t *testing.T) {
	var config types.Configurations

	type args struct {
		status              bool
		currentStatus       bool
		currentStatusErr    error
		epoch               uint32
		epochErr            error
		txnOptsErr          error
		SetCollectionStatus *Types.Transaction
		SetAssetStatusErr   error
		hash                common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When ModifyAssetStatus executes successfully",
			args: args{
				status:              true,
				currentStatus:       false,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting current status",
			args: args{
				status:              true,
				currentStatusErr:    errors.New("current status error"),
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("current status error"),
		},
		{
			name: "Test 3: When currentStatus == status",
			args: args{
				status:              true,
				currentStatus:       true,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 4: When SetAssetStatus transaction fails",
			args: args{
				status:            true,
				currentStatus:     false,
				SetAssetStatusErr: errors.New("SetAssetStatus error"),
				hash:              common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("SetAssetStatus error"),
		},
		{
			name: "Test 5: When WaitForAppropriateState fails",
			args: args{
				status:              true,
				currentStatus:       false,
				epochErr:            errors.New("WaitForAppropriateState error"),
				SetCollectionStatus: &Types.Transaction{},
				SetAssetStatusErr:   nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("WaitForAppropriateState error"),
		},
		{
			name: "Test 6: When there is an error in getting txnOpts",
			args: args{
				status:              true,
				currentStatus:       false,
				SetCollectionStatus: &Types.Transaction{},
				txnOptsErr:          errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetActiveStatus", mock.Anything, mock.Anything).Return(tt.args.currentStatus, tt.args.currentStatusErr)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			cmdUtilsMock.On("WaitForAppropriateState", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.epoch, tt.args.epochErr)
			assetManagerMock.On("SetCollectionStatus", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.SetCollectionStatus, tt.args.SetAssetStatusErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.ModifyCollectionStatus(rpcParameters, config, types.ModifyCollectionInput{
				Status: tt.args.status,
			})
			if got != tt.want {
				t.Errorf("Txn hash for modifyAssetStatus function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for modifyAssetStatus function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for modifyAssetStatus function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteModifyAssetStatus(t *testing.T) {

	var flagSet *pflag.FlagSet
	var config types.Configurations
	var client *ethclient.Client

	type args struct {
		config                     types.Configurations
		configErr                  error
		address                    string
		addressErr                 error
		collectionId               uint16
		collectionIdErr            error
		status                     string
		statusErr                  error
		parseStatus                bool
		parseStatusErr             error
		password                   string
		ModifyCollectionStatusHash common.Hash
		ModifyCollectionStatusErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ModifyAssetStatus executes successfully",
			args: args{
				config:                     config,
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionId:               1,
				status:                     "true",
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting address",
			args: args{
				config:                     config,
				address:                    "",
				addressErr:                 errors.New("address error"),
				collectionId:               1,
				status:                     "true",
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting collectionId",
			args: args{
				config:                     config,
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionIdErr:            errors.New("assetId error"),
				status:                     "true",
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting status string",
			args: args{
				config:                     config,
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionId:               1,
				statusErr:                  errors.New("status error"),
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in parsing status to bool",
			args: args{
				config:                     config,
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionId:               1,
				status:                     "true",
				parseStatusErr:             errors.New("parsing status error"),
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from ModifyAssetStatus",
			args: args{
				config:                     config,
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionId:               1,
				status:                     "true",
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: core.NilHash,
				ModifyCollectionStatusErr:  errors.New("ModifyAssetStatus error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there ia n error in getting config",
			args: args{
				config:                     config,
				configErr:                  errors.New("config error"),
				address:                    "0x000000000000000000000000000000000000dea1",
				collectionId:               1,
				status:                     "true",
				parseStatus:                true,
				password:                   "test",
				ModifyCollectionStatusHash: common.BigToHash(big.NewInt(1)),
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
			flagSetMock.On("GetUint16CollectionId", flagSet).Return(tt.args.collectionId, tt.args.collectionIdErr)
			flagSetMock.On("GetStringStatus", flagSet).Return(tt.args.status, tt.args.statusErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			stringMock.On("ParseBool", mock.AnythingOfType("string")).Return(tt.args.parseStatus, tt.args.parseStatusErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("ModifyCollectionStatus", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.ModifyCollectionStatusHash, tt.args.ModifyCollectionStatusErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteModifyCollectionStatus(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The ExecuteModifyAssetStatus function didn't execute as expected")
			}

		})
	}
}
