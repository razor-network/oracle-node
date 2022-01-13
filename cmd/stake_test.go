package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"testing"
)

func Test_stakeCoins(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	txnArgs := types.TransactionOptions{
		Amount: big.NewInt(10000),
	}

	type args struct {
		txnArgs     types.TransactionOptions
		txnOpts     *bind.TransactOpts
		epoch       uint32
		getEpochErr error
		stakeTxn    *Types.Transaction
		stakeErr    error
		hash        common.Hash
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
				txnOpts:     txnOpts,
				epoch:       2,
				getEpochErr: nil,
				stakeTxn:    &Types.Transaction{},
				stakeErr:    nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When waitForAppropriateState fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				txnOpts:     txnOpts,
				epoch:       2,
				getEpochErr: errors.New("waitForAppropriateState error"),
				stakeTxn:    &Types.Transaction{},
				stakeErr:    nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForAppropriateState error"),
		},
		{
			name: "Test 3: When stake transaction fails",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				txnOpts:     txnOpts,
				epoch:       2,
				getEpochErr: nil,
				stakeTxn:    &Types.Transaction{},
				stakeErr:    errors.New("stake error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("stake error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.getEpochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			stakeManagerUtilsMock.On("Stake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakeTxn, tt.args.stakeErr)

			utils := &UtilsStruct{}

			got, err := utils.StakeCoins(txnArgs)
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
