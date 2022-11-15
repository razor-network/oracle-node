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
	"razor/core"
	"razor/core/types"
	utilsPkgMocks "razor/utils/mocks"
	"testing"
)

func TestUpdateJob(t *testing.T) {

	var client *ethclient.Client
	var config types.Configurations
	var WaitIfCommitStateStatus uint32
	var jobInput types.CreateJobInput
	var jobId uint16

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		txnOpts              *bind.TransactOpts
		updateJobTxn         *Types.Transaction
		updateJobErr         error
		waitIfCommitStateErr error
		hash                 common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1:  When UpdateJob function executes successfully",
			args: args{
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2:  When updateJob transaction fails",
			args: args{
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: errors.New("updateJob error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("updateJob error"),
		},
		{
			name: "Test 3:  When there is an error in WaitIfConfirmState",
			args: args{
				txnOpts:              txnOpts,
				updateJobTxn:         &Types.Transaction{},
				waitIfCommitStateErr: errors.New("waitIfCommitState error"),
				hash:                 common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitIfCommitState error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(utilsPkgMocks.Utils)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			assetManagerUtilsMock := new(mocks.AssetManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			assetManagerUtils = assetManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("WaitIfCommitState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(WaitIfCommitStateStatus, tt.args.waitIfCommitStateErr)
			assetManagerUtilsMock.On("UpdateJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.updateJobTxn, tt.args.updateJobErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.UpdateJob(client, config, jobInput, jobId)
			if got != tt.want {
				t.Errorf("Txn hash for updateJob function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("updateJob() error = %v, wantErr = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for createJob function, got = %v, want = %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteUpdateJob(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	type args struct {
		config          types.Configurations
		configErr       error
		password        string
		address         string
		addressErr      error
		url             string
		urlErr          error
		selector        string
		selectorErr     error
		selectorType    uint8
		selectorTypeErr error
		jobId           uint16
		jobIdErr        error
		power           int8
		powerErr        error
		weight          uint8
		weightErr       error
		updateJobTxn    common.Hash
		updateJobErr    error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1:  When ExecuteUpdateJob function executes successfully",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting address from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting jobId from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobIdErr:     errors.New("jobId error"),
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting url from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				url:          "",
				urlErr:       errors.New("url error"),
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting selector from flags",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "",
				selectorErr:  errors.New("selector error"),
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting power from flag",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				powerErr:     errors.New("power error"),
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there is an error in getting weight from flag",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weightErr:    errors.New("weight error"),
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there an error from UpdateJob",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: core.NilHash,
				updateJobErr: errors.New("updateJob error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When there is an error in getting selectorType",
			args: args{
				config:          config,
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				jobId:           1,
				url:             "https://api.gemini.com/v1/pubticker/ethusd",
				selector:        "last",
				selectorTypeErr: errors.New("selectorType error"),
				power:           1,
				weight:          10,
				updateJobTxn:    common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 10: When selectorType is of XHTML",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 0,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 11: When there is an error in getting config",
			args: args{
				config:       config,
				configErr:    errors.New("config error"),
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				jobId:        1,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				selectorType: 1,
				power:        1,
				weight:       10,
				updateJobTxn: common.BigToHash(big.NewInt(1)),
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
			flagsetUtilsMock := new(mocks.FlagSetInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			fileUtilsMock := new(utilsPkgMocks.FileUtils)

			razorUtils = utilsMock
			flagSetUtils = flagsetUtilsMock
			cmdUtils = cmdUtilsMock
			fileUtils = fileUtilsMock

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			flagsetUtilsMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			flagsetUtilsMock.On("GetStringUrl", flagSet).Return(tt.args.url, tt.args.urlErr)
			flagsetUtilsMock.On("GetStringSelector", flagSet).Return(tt.args.selector, tt.args.selectorErr)
			flagsetUtilsMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			flagsetUtilsMock.On("GetUint16JobId", flagSet).Return(tt.args.jobId, tt.args.jobIdErr)
			flagsetUtilsMock.On("GetUint8Weight", flagSet).Return(tt.args.weight, tt.args.weightErr)
			flagsetUtilsMock.On("GetUint8SelectorType", flagSet).Return(tt.args.selectorType, tt.args.selectorTypeErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("UpdateJob", mock.AnythingOfType("*ethclient.Client"), config, mock.Anything, mock.Anything).Return(tt.args.updateJobTxn, tt.args.updateJobErr)
			utilsMock.On("WaitForBlockCompletion", client, mock.AnythingOfType("string")).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteUpdateJob(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteUpdateJob function didn't execute as expected")
			}
		})
	}
}
