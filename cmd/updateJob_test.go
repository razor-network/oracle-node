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

func Test_updateJob(t *testing.T) {

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	razorUtils := UtilsMock{}
	assetManagerUtils := AssetManagerMock{}
	transactionUtils := TransactionMock{}
	flagSetUtils := FlagSetMock{}

	type args struct {
		password     string
		address      string
		addressErr   error
		url          string
		urlErr       error
		selector     string
		selectorErr  error
		jobId        uint8
		jobIdErr     error
		power        int8
		powerErr     error
		txnOpts      *bind.TransactOpts
		updateJobTxn *Types.Transaction
		updateJobErr error
		hash         common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test1:  When updateJob function executes successfully",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				jobIdErr:     nil,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2: When there is an error in getting address from flags",
			args: args{
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				jobId:        1,
				jobIdErr:     nil,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test3:  When there is an error in getting jobId from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobIdErr:     errors.New("jobId error"),
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("jobId error"),
		},
		{
			name: "Test4:  When there is an error in getting url from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				jobIdErr:     nil,
				url:          "",
				urlErr:       errors.New("url error"),
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr:  errors.New("url error"),
		},
		{
			name: "Test5:  When there is an error in getting selector from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				jobIdErr:     nil,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "",
				selectorErr:  errors.New("selector error"),
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
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
				addressErr:   nil,
				jobId:        1,
				jobIdErr:     nil,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     errors.New("power error"),
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("power error"),
		},
		{
			name: "Test7:  When updateJob transaction fails",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dead",
				addressErr:   nil,
				jobId:        1,
				jobIdErr:     nil,
				url:          "https://api.gemini.com/v1/pubticker/ethusd",
				urlErr:       nil,
				selector:     "last",
				selectorErr:  nil,
				power:        1,
				powerErr:     nil,
				txnOpts:      txnOpts,
				updateJobTxn: &Types.Transaction{},
				updateJobErr: errors.New("updateJob error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("updateJob error"),
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

			GetUint8JobIdMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.jobId, tt.args.jobIdErr
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

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			UpdateJobMock = func(*ethclient.Client, *bind.TransactOpts, uint8, int8, string, string) (*Types.Transaction, error) {
				return tt.args.updateJobTxn, tt.args.updateJobErr
			}

			HashMock = func(transaction *Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err :=updateJob(flagSet, config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
			if got != tt.want {
				t.Errorf("Txn hash for updateJob function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("updateJob() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for createJob function, got = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}
