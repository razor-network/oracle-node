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

func TestClaimBlockReward(t *testing.T) {

	var options types.TransactionOptions

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	razorUtils := UtilsMock{}
	blockManagerUtils := BlockManagerMock{}
	transactionUtils := TransactionMock{}

	type args struct {
		txnOpts             *bind.TransactOpts
		ClaimBlockRewardTxn *Types.Transaction
		ClaimBlockRewardErr error
		hash                common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test1: When ClaimBlockReward function executes successfully",
			args: args{
				txnOpts:             txnOpts,
				ClaimBlockRewardTxn: &Types.Transaction{},
				ClaimBlockRewardErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2: When ClaimBlockReward transaction fails",
			args: args{
				txnOpts:             txnOpts,
				ClaimBlockRewardTxn: &Types.Transaction{},
				ClaimBlockRewardErr: errors.New("claimBlockReward error"),
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("claimBlockReward error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			ClaimBlockRewardMock = func(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error) {
				return tt.args.ClaimBlockRewardTxn, tt.args.ClaimBlockRewardErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := ClaimBlockReward(options, razorUtils, blockManagerUtils, transactionUtils)
			if got != tt.want {
				t.Errorf("Txn hash for ClaimBlockReward function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for ClaimBlockReward function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ClaimBlockReward function, got = %v, want %v", got, tt.wantErr)
				}
			}

		})
	}
}
