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
	"time"
)

func Test_executeClaimBounty(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		cmdUtils:     UtilsCmdMock{},
		flagSetUtils: FlagSetMock{},
	}

	type args struct {
		config         types.Configurations
		configErr      error
		password       string
		address        string
		addressErr     error
		bountyId       uint32
		bountyIdErr    error
		claimBountyTxn common.Hash
		claimBountyErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When executeClaimBounty function executes successfully",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				configErr:      errors.New("config error"),
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				addressErr:     errors.New("address error"),
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting bountyId",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyIdErr:    errors.New("bountyId error"),
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error from claimBounty function",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyErr: errors.New("claimBounty error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetConfigDataMock = func(UtilsStruct) (types.Configurations, error) {
				return tt.args.config, tt.args.configErr
			}

			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetUint32BountyIdMock = func(*pflag.FlagSet) (uint32, error) {
				return tt.args.bountyId, tt.args.bountyIdErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			claimBountyMock = func(types.Configurations, *ethclient.Client, types.RedeemBountyInput, UtilsStruct) (common.Hash, error) {
				return tt.args.claimBountyTxn, tt.args.claimBountyErr
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return 1
			}

			fatal = false
			utilsStruct.executeClaimBounty(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}

		})
	}
}

func Test_claimBounty(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var config types.Configurations
	var client *ethclient.Client
	var bountyInput types.RedeemBountyInput
	var callOpts bind.CallOpts
	var blockTime int64

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		cmdUtils:          UtilsCmdMock{},
	}

	type args struct {
		epoch           uint32
		epochErr        error
		bountyLock      types.BountyLock
		bountyLockErr   error
		redeemBountyTxn *Types.Transaction
		redeemBountyErr error
		hash            common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When claimBounty function executes successfully",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: big.NewInt(70),
				},
				redeemBountyTxn: &Types.Transaction{},
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When claimBounty function executes successfully after waiting for few epochs",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: big.NewInt(80),
				},
				redeemBountyTxn: &Types.Transaction{},
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				epochErr: errors.New("epoch error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When there is an error in getting bounty lock",
			args: args{
				epoch:         70,
				bountyLockErr: errors.New("bountyLock error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("bountyLock error"),
		},
		{
			name: "Test 5: When the amount in bounty lock is 0",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(0),
					RedeemAfter: big.NewInt(70),
				},
			},
			want:    core.NilHash,
			wantErr: errors.New("bounty amount is 0"),
		},
		{
			name: "Test 6: When RedeemBounty transaction fails",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: big.NewInt(70),
				},
				redeemBountyErr: errors.New("redeemBounty error"),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		GetEpochMock = func(*ethclient.Client) (uint32, error) {
			return tt.args.epoch, tt.args.epochErr
		}

		GetOptionsMock = func() bind.CallOpts {
			return callOpts
		}

		GetBountyLockMock = func(*ethclient.Client, *bind.CallOpts, uint32) (types.BountyLock, error) {
			return tt.args.bountyLock, tt.args.bountyLockErr
		}

		SleepMock = func(time.Duration) {

		}

		CalculateBlockTimeMock = func(*ethclient.Client) int64 {
			return blockTime
		}

		GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
			return txnOpts
		}

		RedeemBountyMock = func(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error) {
			return tt.args.redeemBountyTxn, tt.args.redeemBountyErr
		}

		HashMock = func(*Types.Transaction) common.Hash {
			return tt.args.hash
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := claimBounty(config, client, bountyInput, utilsStruct)
			if got != tt.want {
				t.Errorf("Txn hash for claimBounty function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
