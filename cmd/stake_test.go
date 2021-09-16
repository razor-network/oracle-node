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

func Test_stakeCoins(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	razorUtils := UtilsMock{}
	transactionUtils := TransactionMock{}
	stakeManagerUtils := StakeManagerMock{}

	txnArgs := types.TransactionOptions{
		Amount: big.NewInt(10000),
	}

	type args struct {
		txnArgs               types.TransactionOptions
		txnOpts               *bind.TransactOpts
		epoch                 uint32
		waitForCommitStateErr error
		stakeTxn              *Types.Transaction
		stakeErr              error
		hash                  common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When stake transaction is successful",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				txnOpts:               txnOpts,
				epoch:                 2,
				waitForCommitStateErr: nil,
				stakeTxn:              &Types.Transaction{},
				stakeErr:              nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When waitForCommitState fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				txnOpts:               txnOpts,
				epoch:                 2,
				waitForCommitStateErr: errors.New("waitForCommitState error"),
				stakeTxn:              &Types.Transaction{},
				stakeErr:              nil,
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForCommitState error"),
		},
		{
			name: "Test 3: When stake transaction fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				txnOpts:               txnOpts,
				epoch:                 2,
				waitForCommitStateErr: nil,
				stakeTxn:              &Types.Transaction{},
				stakeErr:              errors.New("stake error"),
				hash:                  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("stake error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}
			WaitForCommitStateMock = func(*ethclient.Client, string, string) (uint32, error) {
				return tt.args.epoch, tt.args.waitForCommitStateErr
			}
			StakeMock = func(*bind.TransactOpts, uint32, *big.Int, *ethclient.Client) (*Types.Transaction, error) {
				return tt.args.stakeTxn, tt.args.stakeErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := stakeCoins(txnArgs, razorUtils, stakeManagerUtils, transactionUtils)
			if got != tt.want {
				t.Errorf("Txn hash for stake function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for stake function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for stake function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}
