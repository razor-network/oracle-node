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

func TestCreateJob(t *testing.T) {
	var jobInput types.CreateJobInput
	var config types.Configurations

	type args struct {
		txnOptsErr   error
		createJobTxn *Types.Transaction
		createJobErr error
		hash         common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1:  When createJob function executes successfully",
			args: args{
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2:  When createJob transaction fails",
			args: args{
				createJobTxn: &Types.Transaction{},
				createJobErr: errors.New("createJob error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("createJob error"),
		},
		{
			name: "Test 3:  When there is an error in getting txnOpts",
			args: args{
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			assetManagerMock.On("CreateJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.createJobTxn, tt.args.createJobErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.CreateJob(rpcParameters, config, jobInput)
			if got != tt.want {
				t.Errorf("Txn hash for createJob function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for createJob function, got = %v, want = %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for createJob function, got = %v, want = %v", got, tt.wantErr)
				}
			}

		})
	}
}

func TestExecuteCreateJob(t *testing.T) {
	var client *ethclient.Client
	var config types.Configurations
	var flagSet *pflag.FlagSet

	type args struct {
		config          types.Configurations
		configErr       error
		password        string
		address         string
		addressErr      error
		name            string
		nameErr         error
		url             string
		urlErr          error
		selector        string
		selectorErr     error
		selectorType    uint8
		selectorTypeErr error
		power           int8
		powerErr        error
		weight          uint8
		weightErr       error
		createJobHash   common.Hash
		createJobErr    error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1:  When ExecuteCreateJob function executes successfully",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				power:         1,
				weight:        10,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2:  When there is an error in getting address from flags",
			args: args{
				config:        config,
				password:      "test",
				address:       "",
				addressErr:    errors.New("address error"),
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				power:         1,
				weight:        20,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3:  When there is an error in getting name from flags",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "",
				nameErr:       errors.New("name error"),
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				power:         1,
				weight:        20,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4:  When there is an error in getting url from flags",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "",
				urlErr:        errors.New("url error"),
				selector:      "last",
				selectorType:  1,
				power:         1,
				weight:        20,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5:  When there is an error in getting selector from flags",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "",
				selectorErr:   errors.New("selector error"),
				power:         1,
				selectorType:  1,
				weight:        20,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6:  When there is an error in getting power from flag",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				powerErr:      errors.New("power error"),
				weight:        20,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7:  When there is an error in getting weight from flag",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				weightErr:     errors.New("weight error"),
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 8:  When there is an error in getting selectorType from flag",
			args: args{
				config:          config,
				password:        "test",
				address:         "0x000000000000000000000000000000000000dead",
				name:            "ETH-1",
				url:             "https://api.gemini.com/v1/pubticker/ethusd",
				selector:        "last",
				selectorTypeErr: errors.New("selectorType error"),
				createJobHash:   common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 9:  When the selector type is for XHTML link",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "BTC_COIN_GECKO",
				url:           "https://www.coingecko.com/en ",
				selector:      `table tbody tr td span[data-coin-id="1"][data-target="price.price"] span`,
				selectorType:  0,
				createJobHash: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 10:  When there is an error from CreateJob function",
			args: args{
				config:        config,
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				power:         1,
				weight:        20,
				createJobHash: core.NilHash,
				createJobErr:  errors.New("createJob error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 11:  When there is an error in getting config",
			args: args{
				config:        config,
				configErr:     errors.New("config error"),
				password:      "test",
				address:       "0x000000000000000000000000000000000000dead",
				name:          "ETH-1",
				url:           "https://api.gemini.com/v1/pubticker/ethusd",
				selector:      "last",
				selectorType:  1,
				power:         1,
				weight:        10,
				createJobHash: common.BigToHash(big.NewInt(1)),
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
			flagSetMock.On("GetStringUrl", flagSet).Return(tt.args.url, tt.args.urlErr)
			flagSetMock.On("GetStringSelector", flagSet).Return(tt.args.selector, tt.args.selectorErr)
			flagSetMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			flagSetMock.On("GetUint8Weight", flagSet).Return(tt.args.weight, tt.args.weightErr)
			flagSetMock.On("GetUint8SelectorType", flagSet).Return(tt.args.selectorType, tt.args.selectorTypeErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("CreateJob", mock.Anything, config, mock.Anything).Return(tt.args.createJobHash, tt.args.createJobErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteCreateJob(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteCreateJob function didn't execute as expected")
			}
		})
	}
}
