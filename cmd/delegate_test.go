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
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"
)

func Test_delegate(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		transactionUtils:  TransactionMock{},
		stakeManagerUtils: StakeManagerMock{},
	}

	var txnArgs types.TransactionOptions
	var stakerId uint32 = 1

	type args struct {
		amount      *big.Float
		txnOpts     *bind.TransactOpts
		epoch       uint32
		epochErr    error
		delegateTxn *Types.Transaction
		delegateErr error
		hash        common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When delegate function executes successfully",
			args: args{
				amount:      big.NewFloat(1000),
				txnOpts:     txnOpts,
				epoch:       1,
				epochErr:    nil,
				delegateTxn: &Types.Transaction{},
				delegateErr: nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When delegate transaction fails",
			args: args{
				amount:      big.NewFloat(1000),
				txnOpts:     txnOpts,
				epoch:       1,
				epochErr:    nil,
				delegateTxn: &Types.Transaction{},
				delegateErr: errors.New("delegate error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("delegate error"),
		},
		{
			name: "Test 3: When GetEpoch fails",
			args: args{
				amount:      big.NewFloat(1000),
				txnOpts:     txnOpts,
				epoch:       1,
				epochErr:    errors.New("GetEpoch error"),
				delegateTxn: &Types.Transaction{},
				delegateErr: nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("GetEpoch error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAmountInDecimalMock = func(*big.Int) *big.Float {
				return tt.args.amount
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			GetEpochMock = func(client *ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			DelegateMock = func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error) {
				return tt.args.delegateTxn, tt.args.delegateErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.delegate(txnArgs, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for delegate function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for delegate function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for delegate function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}
