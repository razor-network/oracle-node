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

func TestCreateCollection(t *testing.T) {
	var WaitForDisputeOrConfirmStateStatus uint32
	var config types.Configurations
	var collectionInput types.CreateCollectionInput

	type args struct {
		jobIdUint8                 []uint16
		waitForAppropriateStateErr error
		txnOptsErr                 error
		createCollectionTxn        *Types.Transaction
		createCollectionErr        error
		hash                       common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When CreateCollection function executes successfully",
			args: args{
				jobIdUint8:          []uint16{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in WaitForConfirmState",
			args: args{
				jobIdUint8:                 []uint16{1, 2},
				waitForAppropriateStateErr: errors.New("waitForDisputeOrConfirmState error"),
				createCollectionTxn:        &Types.Transaction{},
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForDisputeOrConfirmState error"),
		},
		{
			name: "Test 3: When CreateCollection transaction fails",
			args: args{
				jobIdUint8:          []uint16{1, 2},
				createCollectionTxn: &Types.Transaction{},
				createCollectionErr: errors.New("createCollection error"),
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("createCollection error"),
		},
		{
			name: "Test 4: When there is an error in getting txnOpts",
			args: args{
				jobIdUint8: []uint16{1, 2},
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("ConvertUintArrayToUint16Array", mock.Anything).Return(tt.args.jobIdUint8)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			cmdUtilsMock.On("WaitForAppropriateState", mock.Anything, mock.Anything, mock.Anything).Return(WaitForDisputeOrConfirmStateStatus, tt.args.waitForAppropriateStateErr)
			assetManagerMock.On("CreateCollection", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.createCollectionTxn, tt.args.createCollectionErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.CreateCollection(rpcParameters, config, collectionInput)
			if got != tt.want {
				t.Errorf("Txn hash for createCollection function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for createCollection function, got = %v, want = %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for createCollection function, got = %v, want = %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteCreateCollection(t *testing.T) {
	var client *ethclient.Client
	var config types.Configurations
	var flagSet *pflag.FlagSet

	type args struct {
		config               types.Configurations
		configErr            error
		password             string
		name                 string
		nameErr              error
		address              string
		addressErr           error
		jobId                []uint
		jobIdErr             error
		aggregation          uint32
		aggregationErr       error
		power                int8
		powerErr             error
		tolerance            uint32
		toleranceErr         error
		createCollectionErr  error
		createCollectionHash common.Hash
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteCreateCollection function executes successfully",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting name from flags",
			args: args{
				config:               config,
				password:             "test",
				name:                 "",
				nameErr:              errors.New("name error"),
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address from flags",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "",
				addressErr:           errors.New("address error"),
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting jobId's from flags",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobIdErr:             errors.New("jobId error"),
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting aggregation method from flags",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregationErr:       errors.New("aggregation error"),
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting power from flags",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				powerErr:             errors.New("power error"),
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error from CreateCollection",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionErr:  errors.New("createCollection error"),
				createCollectionHash: core.NilHash,
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there is an error in getting config",
			args: args{
				config:               config,
				configErr:            errors.New("config error"),
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				tolerance:            20,
				createCollectionHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When there is an error in getting tolerance",
			args: args{
				config:               config,
				password:             "test",
				name:                 "ETH-Collection",
				address:              "0x000000000000000000000000000000000000dead",
				jobId:                []uint{1, 2},
				aggregation:          1,
				power:                0,
				toleranceErr:         errors.New("tolerance error"),
				createCollectionHash: common.BigToHash(big.NewInt(1)),
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
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			flagSetMock.On("GetStringName", flagSet).Return(tt.args.name, tt.args.nameErr)
			flagSetMock.On("GetUintSliceJobIds", flagSet).Return(tt.args.jobId, tt.args.jobIdErr)
			flagSetMock.On("GetUint32Aggregation", flagSet).Return(tt.args.aggregation, tt.args.aggregationErr)
			flagSetMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			flagSetMock.On("GetUint32Tolerance", flagSet).Return(tt.args.tolerance, tt.args.toleranceErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("CreateCollection", mock.Anything, config, mock.Anything).Return(tt.args.createCollectionHash, tt.args.createCollectionErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteCreateCollection(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteCreateCollection function didn't execute as expected")
			}
		})
	}
}
