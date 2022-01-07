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
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"testing"
)

func TestWithdrawFunds(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	type args struct {
		lock                     types.Locks
		lockErr                  error
		withdrawReleasePeriod    uint8
		withdrawReleasePeriodErr error
		txnOpts                  *bind.TransactOpts
		epoch                    uint32
		epochErr                 error
		updatedEpoch             uint32
		updatedEpochErr          error
		withdrawHash             common.Hash
		withdrawErr              error
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When withdrawFunds function executes successfully",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				lockErr:                  errors.New("lock error"),
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("lock error"),
		},
		{
			name: "Test 3: When staker tries to withdraw without un-staking any Razors",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(0),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 4: When there is an error in getting withdrawReleasePeriod",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriodErr: errors.New("withdrawReleasePeriod error"),
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("withdrawReleasePeriod error"),
		},
		{
			name: "Test 5: When there is an error in getting epoch",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epochErr:                 errors.New("epoch error"),
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 6: When staker tries to withdraw when withdrawal period has passed",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    9,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 7: When withdraw function is not being called",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    1,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    5,
				epochErr:                 nil,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When there is a need to wait till withdrawAfter and withdraw function executes successfully",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    3,
				epochErr:                 nil,
				updatedEpoch:             5,
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 9: When there is a need to wait till withdrawAfter but there is an error in getting updated Epoch ",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                  nil,
				withdrawReleasePeriod:    4,
				withdrawReleasePeriodErr: nil,
				txnOpts:                  txnOpts,
				epoch:                    3,
				epochErr:                 nil,
				updatedEpochErr:          errors.New("updatedEpoch error"),
				withdrawHash:             common.BigToHash(big.NewInt(1)),
				withdrawErr:              nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("updatedEpoch error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterfaceMockery)
			stakeManagerUtilsMock := new(mocks.StakeManagerInterfaceMockery)
			cmdUtilsMock := new(mocks.UtilsCmdInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			stakeManagerUtilsMockery = stakeManagerUtilsMock
			cmdUtilsMockery = cmdUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.lock, tt.args.lockErr)
			utilsMock.On("GetWithdrawReleasePeriod", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetUpdatedEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.updatedEpoch, tt.args.updatedEpochErr)
			cmdUtilsMock.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawHash, tt.args.withdrawErr)
			utilsMock.On("Sleep", mock.Anything).Return()

			utils := &UtilsStructMockery{}
			got, err := utils.WithdrawFunds(client, account, configurations, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for withdrawFunds function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for withdrawFunds function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for withdrawFunds function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestWithdraw(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var stakerId uint32

	type args struct {
		withdrawTxn *Types.Transaction
		withdrawErr error
		hash        common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When withdraw function executes successfully",
			args: args{
				withdrawTxn: &Types.Transaction{},
				withdrawErr: nil,
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When Withdraw transaction fails",
			args: args{
				withdrawTxn: &Types.Transaction{},
				withdrawErr: errors.New("withdraw error"),
				hash:        common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("withdraw error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			stakeManagerUtilsMock := new(mocks.StakeManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			stakeManagerUtilsMockery = stakeManagerUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			stakeManagerUtilsMock.On("Withdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.withdrawTxn, tt.args.withdrawErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStructMockery{}
			got, err := utils.Withdraw(client, txnOpts, stakerId)
			if got != tt.want {
				t.Errorf("Txn hash for withdraw function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for withdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for withdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
