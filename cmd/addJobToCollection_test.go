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

func Test_addJobToCollection(t *testing.T) {

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
		password              string
		address               string
		addressErr            error
		jobId                 uint8
		jobIdErr              error
		collectionId          uint8
		collectionIdErr       error
		txnOpts               *bind.TransactOpts
		addJobToCollectionTxn *Types.Transaction
		addJobToCollectionErr error
		hash                  common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test1: When addJobToCollection function executes successfully",
			args: args{
				password:              "test",
				address:               "0x000000000000000000000000000000000000dead",
				addressErr:            nil,
				jobId:                 1,
				jobIdErr:              nil,
				collectionId:          1,
				collectionIdErr:       nil,
				txnOpts:               txnOpts,
				addJobToCollectionTxn: &Types.Transaction{},
				addJobToCollectionErr: nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2: When there is an error in getting address from flags",
			args: args{
				password:              "test",
				address:               "",
				addressErr:            errors.New("address error"),
				jobId:                 1,
				jobIdErr:              nil,
				collectionId:          1,
				collectionIdErr:       nil,
				txnOpts:               txnOpts,
				addJobToCollectionTxn: &Types.Transaction{},
				addJobToCollectionErr: nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test3: When there is an error in getting jobId from flags",
			args: args{
				password:              "test",
				address:               "0x000000000000000000000000000000000000dead",
				addressErr:            nil,
				jobIdErr:              errors.New("jobId error"),
				collectionId:          1,
				collectionIdErr:       nil,
				txnOpts:               txnOpts,
				addJobToCollectionTxn: &Types.Transaction{},
				addJobToCollectionErr: nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("jobId error"),
		},
		{
			name: "Test4: When there is an error in getting collectionId from flags",
			args: args{
				password:              "test",
				address:               "0x000000000000000000000000000000000000dead",
				addressErr:            nil,
				jobId:                 1,
				jobIdErr:              nil,
				collectionIdErr:       errors.New("collectionId error"),
				txnOpts:               txnOpts,
				addJobToCollectionTxn: &Types.Transaction{},
				addJobToCollectionErr: nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("collectionId error"),
		},
		{
			name: "Test5: When AddJobToCollection transaction fails",
			args: args{
				password:              "test",
				address:               "0x000000000000000000000000000000000000dead",
				addressErr:            nil,
				jobId:                 1,
				jobIdErr:              nil,
				collectionId:          1,
				collectionIdErr:       nil,
				txnOpts:               txnOpts,
				addJobToCollectionTxn: &Types.Transaction{},
				addJobToCollectionErr: errors.New("addJobToCollection error"),
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("addJobToCollection error"),
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

			GetUint8CollectionIdMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.collectionId, tt.args.collectionIdErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			AddJobToCollectionMock = func(*ethclient.Client, *bind.TransactOpts, uint8, uint8) (*Types.Transaction, error) {
				return tt.args.addJobToCollectionTxn, tt.args.addJobToCollectionErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := addJobToCollection(flagSet, config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
			if got != tt.want {
				t.Errorf("Txn hash for addJobToCollection function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for addJobToCollection function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for addJobToCollection function, got = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}
