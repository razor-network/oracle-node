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
	"razor/core"
	"razor/core/types"
	"testing"
)

func Test_createCollection(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var WaitForDisputeOrConfirmStateStatus uint32
	var flagSet *pflag.FlagSet
	var config types.Configurations

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		assetManagerUtils: AssetManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		password                   string
		name                       string
		nameErr                    error
		address                    string
		addressErr                 error
		jobId                      []uint
		jobIdErr                   error
		aggregation                uint32
		aggregationErr             error
		power                      int8
		powerErr                   error
		txnOpts                    *bind.TransactOpts
		jobIdUint8                 []uint8
		waitForAppropriateStateErr error
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
			name: "Test 1: When createCollection function executes successfully",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "0x000000000000000000000000000000000000dead",
				jobId:               []uint{1, 2},
				aggregation:         1,
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting name from flags",
			args: args{
				password:            "test",
				name:                "",
				nameErr:             errors.New("name error"),
				address:             "0x000000000000000000000000000000000000dead",
				jobId:               []uint{1, 2},
				aggregation:         1,
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("name error"),
		},
		{
			name: "Test 3: When there is an error in getting address from flags",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "",
				addressErr:          errors.New("address error"),
				jobId:               []uint{1, 2},
				aggregation:         1,
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test 4: When there is an error in getting jobId's from flags",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "0x000000000000000000000000000000000000dead",
				jobIdErr:            errors.New("jobId error"),
				aggregation:         1,
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("jobId error"),
		},
		{
			name: "Test 5: When there is an error in getting aggregation method from flags",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "0x000000000000000000000000000000000000dead",
				jobId:               []uint{1, 2},
				aggregationErr:      errors.New("aggregation error"),
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("aggregation error"),
		},
		{
			name: "Test 6: When there is an error in getting power from flags",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "0x000000000000000000000000000000000000dead",
				jobId:               []uint{1, 2},
				aggregation:         1,
				powerErr:            errors.New("power error"),
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("power error"),
		},
		{
			name: "Test 7: When there is an error in WaitForConfirmState",
			args: args{
				password:                   "test",
				name:                       "ETH-Collection",
				address:                    "0x000000000000000000000000000000000000dead",
				jobId:                      []uint{1, 2},
				aggregation:                1,
				power:                      0,
				txnOpts:                    txnOpts,
				jobIdUint8:                 []uint8{1, 2},
				waitForAppropriateStateErr: errors.New("waitForDisputeOrConfirmState error"),
				createCollectionTxn:        &Types.Transaction{},
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForDisputeOrConfirmState error"),
		},
		{
			name: "Test 8: When CreateCollection transaction fails",
			args: args{
				password:            "test",
				name:                "ETH-Collection",
				address:             "0x000000000000000000000000000000000000dead",
				jobId:               []uint{1, 2},
				aggregation:         1,
				power:               0,
				txnOpts:             txnOpts,
				jobIdUint8:          []uint8{1, 2},
				createCollectionTxn: &Types.Transaction{},
				createCollectionErr: errors.New("createCollection error"),
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("createCollection error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			GetStringNameMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.name, tt.args.nameErr
			}

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetUintSliceJobIdsMock = func(*pflag.FlagSet) ([]uint, error) {
				return tt.args.jobId, tt.args.jobIdErr
			}

			GetUint32AggregationMock = func(*pflag.FlagSet) (uint32, error) {
				return tt.args.aggregation, tt.args.aggregationErr
			}

			GetInt8PowerMock = func(*pflag.FlagSet) (int8, error) {
				return tt.args.power, tt.args.powerErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			ConvertUintArrayToUint8ArrayMock = func([]uint) []uint8 {
				return tt.args.jobIdUint8
			}

			WaitForAppropriateStateMock = func(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error) {
				return WaitForDisputeOrConfirmStateStatus, tt.args.waitForAppropriateStateErr
			}

			CreateCollectionMock = func(*ethclient.Client, *bind.TransactOpts, []uint8, uint32, int8, string) (*Types.Transaction, error) {
				return tt.args.createCollectionTxn, tt.args.createCollectionErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.createCollection(flagSet, config)
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
