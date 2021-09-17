package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"
)

func Test_createJob(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	razorUtils := UtilsMock{}
	assetManagerUtils := AssetManagerMock{}
	transactionUtils := TransactionMock{}

	var txnArgs types.TransactionOptions
	var power int8
	var name, selector, url string

	type args struct {
		txnOpts      *bind.TransactOpts
		createJobtxn *Types.Transaction
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
			name: "Test1: When createJob transaction is succesull",
			args: args{
				txnOpts:      txnOpts,
				createJobtxn: &Types.Transaction{},
				createJobErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2: When createJob transaction fails",
			args: args{
				txnOpts:      txnOpts,
				createJobtxn: &Types.Transaction{},
				createJobErr: errors.New("createJob error"),
				hash:         common.Hash{0x00},
			},
			want:    core.NilHash,
			wantErr: errors.New("createJob error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			CreateJobMock = func(*bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error) {
				return tt.args.createJobtxn, tt.args.createJobErr
			}

			HashMock = func(transaction *Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := createJob(txnArgs, power, name, selector, url, razorUtils, assetManagerUtils, transactionUtils)
			if got != tt.want {
				t.Errorf("Txn hash for createJob function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for createJob function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for createJob function, got = %v, want %v", got, tt.wantErr)
				}
			}

		})
	}
}
