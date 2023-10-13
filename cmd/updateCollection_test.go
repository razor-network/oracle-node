package cmd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestUpdateCollection(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var config types.Configurations
	var WaitIfCommitStateStatus uint32
	var jobIdUint16 []uint16
	var collectionInput types.CreateCollectionInput
	var collectionId uint16

	type args struct {
		txnOpts              *bind.TransactOpts
		updateCollectionTxn  *Types.Transaction
		updateCollectionErr  error
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
			name: "Test 1: When UpdateCollection function executes successfully",
			args: args{
				txnOpts:              txnOpts,
				updateCollectionTxn:  &Types.Transaction{},
				updateCollectionErr:  nil,
				waitIfCommitStateErr: nil,
				hash:                 common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When updateCollection transaction fails",
			args: args{
				txnOpts:              txnOpts,
				updateCollectionTxn:  &Types.Transaction{},
				updateCollectionErr:  errors.New("updateCollection error"),
				waitIfCommitStateErr: nil,
				hash:                 common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("updateCollection error"),
		},
		{
			name: "Test 3: When there is an error in WaitIfConfirmState",
			args: args{
				txnOpts:              txnOpts,
				updateCollectionTxn:  &Types.Transaction{},
				updateCollectionErr:  nil,
				waitIfCommitStateErr: errors.New("waitIfCommitState error"),
				hash:                 common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitIfCommitState error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("ConvertUintArrayToUint16Array", mock.Anything).Return(jobIdUint16)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("WaitIfCommitState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(WaitIfCommitStateStatus, tt.args.waitIfCommitStateErr)
			assetManagerMock.On("UpdateCollection", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.updateCollectionTxn, tt.args.updateCollectionErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.UpdateCollection(client, config, collectionInput, collectionId)

			if got != tt.want {
				t.Errorf("Txn hash for updateCollection function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for updateCollection function, got = %v, want = %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for updateCollection function, got = %v, want = %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteUpdateCollection(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	type args struct {
		config              types.Configurations
		configErr           error
		password            string
		collectionId        uint16
		collectionIdErr     error
		address             string
		addressErr          error
		aggregation         uint32
		aggregationErr      error
		power               int8
		powerErr            error
		tolerance           uint32
		toleranceErr        error
		jobId               []uint
		jobIdErr            error
		updateCollectionTxn common.Hash
		updateCollectionErr error
	}

	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteUpdateCollection function executes successfully",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				power:               0,
				tolerance:           15,
				jobId:               []uint{1, 2},
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting collection id from flags",
			args: args{
				config:              config,
				password:            "test",
				collectionIdErr:     errors.New("collectionIdErr error"),
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				power:               0,
				tolerance:           15,
				jobId:               []uint{1, 2},
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address from flags",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "",
				addressErr:          errors.New("address error"),
				aggregation:         1,
				power:               0,
				tolerance:           15,
				jobId:               []uint{1, 2},
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting aggregation method from flags",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregationErr:      errors.New("aggregation error"),
				power:               0,
				tolerance:           15,
				jobId:               []uint{1, 2},
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting power from flags",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				powerErr:            errors.New("power error"),
				tolerance:           15,
				jobId:               []uint{1, 2},
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error in getting jobIds from flags",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				tolerance:           15,
				jobIdErr:            errors.New("job Id error"),
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 7: When there ia an from UpdateCollection",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				power:               0,
				tolerance:           15,
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
				updateCollectionErr: errors.New("updateCollection error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 8: When there is an error in getting config",
			args: args{
				config:              config,
				configErr:           errors.New("config error"),
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				power:               0,
				tolerance:           15,
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 9: When there is an error in getting tolerance",
			args: args{
				config:              config,
				password:            "test",
				collectionId:        3,
				address:             "0x000000000000000000000000000000000000dead",
				aggregation:         1,
				power:               0,
				toleranceErr:        errors.New("tolerance error"),
				updateCollectionTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", flagSet).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything, mock.Anything).Return(nil)
			flagSetMock.On("GetStringAddress", flagSet).Return(tt.args.address, tt.args.addressErr)
			flagSetMock.On("GetUint16CollectionId", flagSet).Return(tt.args.collectionId, tt.args.collectionIdErr)
			flagSetMock.On("GetUintSliceJobIds", flagSet).Return(tt.args.jobId, tt.args.jobIdErr)
			flagSetMock.On("GetUint32Aggregation", flagSet).Return(tt.args.aggregation, tt.args.aggregationErr)
			flagSetMock.On("GetInt8Power", flagSet).Return(tt.args.power, tt.args.powerErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			cmdUtilsMock.On("UpdateCollection", mock.AnythingOfType("*ethclient.Client"), config, mock.Anything, mock.Anything).Return(tt.args.updateCollectionTxn, tt.args.updateCollectionErr)
			utilsMock.On("WaitForBlockCompletion", client, mock.AnythingOfType("string")).Return(nil)
			flagSetMock.On("GetUint32Tolerance", flagSet).Return(tt.args.tolerance, tt.args.toleranceErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteUpdateCollection(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteUpdateCollection function didn't execute as expected")
			}

		})
	}
}
