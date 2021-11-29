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

func Test_createJob(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		assetManagerUtils: AssetManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
	}

	type args struct {
		password     string
		address      string
		addressErr   error
		name         string
		nameErr      error
		url          string
		urlErr       error
		selector     string
		selectorErr  error
		power        int8
		powerErr     error
		weight       uint8
		weightErr    error
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
			name: "Test1:  When createJob function executes successfully",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				power:        1,
				weight:       10,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2:  When there is an error in getting address from flags",
			args: args{
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				power:        1,
				weight:       20,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test3:  When there is an error in getting name from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "",
				nameErr:      errors.New("name error"),
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				power:        1,
				weight:       20,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("name error"),
		},
		{
			name: "Test4:  When there is an error in getting url from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "",
				urlErr:       errors.New("url error"),
				selector:     "last",
				power:        1,
				weight:       20,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("url error"),
		},
		{
			name: "Test5:  When there is an error in getting selector from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "",
				selectorErr:  errors.New("selector error"),
				power:        1,
				weight:       20,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("selector error"),
		},
		{
			name: "Test6:  When there is an error in getting power from flag",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				powerErr:     errors.New("power error"),
				weight:       20,
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("power error"),
		},
		{
			name: "Test7:  When there is an error in getting weight from flag",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				weightErr:    errors.New("weight error"),
				txnOpts:      txnOpts,
				createJobTxn: &Types.Transaction{},
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("weight error"),
		},
		{
			name: "Test8:  When createJob transaction fails",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				name:         "ETH-1",
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				selector:     "last",
				power:        1,
				weight:       20,
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
			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetStringNameMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.name, tt.args.nameErr
			}

			GetStringUrlMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.url, tt.args.urlErr
			}

			GetStringSelectorMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.selector, tt.args.selectorErr
			}

			GetInt8PowerMock = func(*pflag.FlagSet) (int8, error) {
				return tt.args.power, tt.args.powerErr
			}

			GetUint8WeightMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.weight, tt.args.weightErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions, UtilsStruct) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			CreateJobMock = func(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) (*Types.Transaction, error) {
				return tt.args.createJobTxn, tt.args.createJobErr
			}

			HashMock = func(transaction *Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.createJob(flagSet, config)
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
