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

func Test_checkForCommitStateAndWithdraw(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var account types.Account
	var configurations types.Configurations
	var stakerId uint32

	razorUtils := UtilsMock{}
	cmdUtils := UtilsCmdMock{}
	stakeManagerUtils := StakeManagerMock{}
	transactionUtils := TransactionMock{}

	type args struct {
		lock                       types.Locks
		lockErr                    error
		withdrawReleasePeriod      uint8
		withdrawReleasePeriodErr   error
		txnOpts                    *bind.TransactOpts
		epoch                      uint32
		epochErr                   error
		commitStateEpoch           uint32
		commitStateEpochErr        error
		withdrawHash               common.Hash
		withdrawErr                error
		updatedCommitStateEpoch    uint32
		updatedCommitStateEpochErr error
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When checkForCommitStateAndWithdraw function executes successfully",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting lock",
			args: args{
				lockErr:                    errors.New("lock error"),
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
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
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
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
				lockErr:                    nil,
				withdrawReleasePeriodErr:   errors.New("withdrawReleasePeriod error"),
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
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
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epochErr:                   errors.New("epoch error"),
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
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
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      9,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 7: When there is an error in getting commitStateEpoch from WaitForCommitState ",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpochErr:        errors.New("waitForCommitState error"),
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForCommitState error"),
		},
		{
			name: "Test 8: When there is an error in getting updated commitStateEpoch from WaitForCommitStateAgain",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(5),
				},
				lockErr:                    nil,
				withdrawReleasePeriod:      4,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      4,
				epochErr:                   nil,
				commitStateEpoch:           4,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpochErr: errors.New("waitForCommitState error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("waitForCommitState error"),
		},
		{
			name: "Test 9: When withdraw function is not being called",
			args: args{
				lock: types.Locks{
					WithdrawAfter: big.NewInt(4),
				},
				lockErr:                    nil,
				withdrawReleasePeriod:      1,
				withdrawReleasePeriodErr:   nil,
				txnOpts:                    txnOpts,
				epoch:                      5,
				epochErr:                   nil,
				commitStateEpoch:           6,
				commitStateEpochErr:        nil,
				withdrawHash:               common.BigToHash(big.NewInt(1)),
				withdrawErr:                nil,
				updatedCommitStateEpoch:    7,
				updatedCommitStateEpochErr: nil,
			},
			want:    core.NilHash,
			wantErr: nil,
		},
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

		GetEpochMock = func(*ethclient.Client, string) (uint32, error) {
			return tt.args.epoch, tt.args.epochErr
		}

		WaitForCommitStateMock = func(*ethclient.Client, string, string) (uint32, error) {
			return tt.args.commitStateEpoch, tt.args.commitStateEpochErr
		}

		WithdrawMock = func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, stakeManagerInterface, transactionInterface) (common.Hash, error) {
			return tt.args.withdrawHash, tt.args.withdrawErr
		}

		WaitForCommitStateAgainMock = func(*ethclient.Client, string, string) (uint32, error) {
			return tt.args.updatedCommitStateEpoch, tt.args.updatedCommitStateEpochErr
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := checkForCommitStateAndWithdraw(client, account, configurations, stakerId, razorUtils, cmdUtils, stakeManagerUtils, transactionUtils)
			if got != tt.want {
				t.Errorf("Txn hash for checkForCommitStateAndWithdraw function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for checkForCommitStateAndWithdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for checkForCommitStateAndWithdraw function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func Test_withdraw(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	stakeManagerUtils := StakeManagerMock{}
	transactionUtils := TransactionMock{}

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

			got, err := withdraw(client, txnOpts, epoch, stakerId, stakeManagerUtils, transactionUtils)
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
