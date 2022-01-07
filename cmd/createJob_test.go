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
	"testing"
)

func TestCreateJob(t *testing.T) {
	var client *ethclient.Client
	var jobInput types.CreateJobInput
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		txnOpts      *bind.TransactOpts
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
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2:  When createJob transaction fails",
			args: args{
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				createJobErr: errors.New("createJob error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("createJob error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			assetManagerUtilsMock := new(mocks.AssetManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			assetManagerUtilsMockery = assetManagerUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			assetManagerUtilsMock.On("CreateJob", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.createJobTxn, tt.args.createJobErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStructMockery{}
			got, err := utils.CreateJob(client, config, jobInput)
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			flagsetUtilsMock := new(mocks.FlagSetInterfaceMockery)
			cmdUtilsMock := new(mocks.UtilsCmdInterfaceMockery)

			razorUtilsMockery = utilsMock
			flagSetUtilsMockery = flagsetUtilsMock
			cmdUtilsMockery = cmdUtilsMock

			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			flagsetUtilsMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			flagsetUtilsMock.On("GetStringName", flagSet).Return(tt.args.name, tt.args.nameErr)
			flagsetUtilsMock.On("GetStringUrl", flagSet).Return(tt.args.url, tt.args.urlErr)
			flagsetUtilsMock.On("GetStringSelector", flagSet).Return(tt.args.selector, tt.args.selectorErr)
			flagsetUtilsMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			flagsetUtilsMock.On("GetUint8Weight", flagSet).Return(tt.args.weight, tt.args.weightErr)
			flagsetUtilsMock.On("GetUint8SelectorType", flagSet).Return(tt.args.selectorType, tt.args.selectorTypeErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("CreateJob", mock.AnythingOfType("*ethclient.Client"), config, mock.Anything).Return(tt.args.createJobHash, tt.args.createJobErr)
			utilsMock.On("WaitForBlockCompletion", client, mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStructMockery{}
			fatal = false

			utils.ExecuteCreateJob(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteCreateJob function didn't execute as expected")
			}
		})
	}
}
