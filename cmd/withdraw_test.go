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

func Test_withdrawFunds(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		cmdUtils:          UtilsCmdMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
	}

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
		//{
		//	name: "Test 8: When there is a need to wait till withdrawAfter and withdraw function executes successfully",
		//	args: args{
		//		lock: types.Locks{
		//			WithdrawAfter: big.NewInt(4),
		//		},
		//		lockErr:                  nil,
		//		withdrawReleasePeriod:    4,
		//		withdrawReleasePeriodErr: nil,
		//		txnOpts:                  txnOpts,
		//		epoch:                    3,
		//		epochErr:                 nil,
		//		updatedEpoch:             5,
		//		withdrawHash:             common.BigToHash(big.NewInt(1)),
		//		withdrawErr:              nil,
		//	},
		//	want:    common.BigToHash(big.NewInt(1)),
		//	wantErr: nil,
		//},
		//{
		//	name: "Test 9: When there is a need to wait till withdrawAfter but there is an error in getting updated Epoch ",
		//	args: args{
		//		lock: types.Locks{
		//			WithdrawAfter: big.NewInt(4),
		//		},
		//		lockErr:                  nil,
		//		withdrawReleasePeriod:    4,
		//		withdrawReleasePeriodErr: nil,
		//		txnOpts:                  txnOpts,
		//		epoch:                    3,
		//		epochErr:                 nil,
		//		updatedEpochErr:          errors.New("updatedEpoch error"),
		//		withdrawHash:             common.BigToHash(big.NewInt(1)),
		//		withdrawErr:              nil,
		//	},
		//	want:    core.NilHash,
		//	wantErr: errors.New("updatedEpoch error"),
		//},
	}
	for _, tt := range tests {

		GetLockMock = func(*ethclient.Client, string, uint32) (types.Locks, error) {
			return tt.args.lock, tt.args.lockErr
		}

		GetWithdrawReleasePeriodMock = func(*ethclient.Client, string) (uint8, error) {
			return tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr
		}

		GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
			return tt.args.txnOpts
		}

		GetEpochMock = func(*ethclient.Client) (uint32, error) {
			return tt.args.epoch, tt.args.epochErr
		}

		GetUpdatedEpochMock = func(*ethclient.Client) (uint32, error) {
			return tt.args.updatedEpoch, tt.args.updatedEpochErr
		}

		WithdrawMock = func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, UtilsStruct) (common.Hash, error) {
			return tt.args.withdrawHash, tt.args.withdrawErr
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := utilsStruct.withdrawFunds(client, account, configurations, stakerId)
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

func Test_withdraw(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
	}

	var client *ethclient.Client
	var epoch uint32
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
			WithdrawContractMock = func(*ethclient.Client, *bind.TransactOpts, uint32, uint32) (*Types.Transaction, error) {
				return tt.args.withdrawTxn, tt.args.withdrawErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := withdraw(client, txnOpts, epoch, stakerId, utilsStruct)
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
