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

func Test_extendLock(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var flagSet *pflag.FlagSet
	var config types.Configurations
	var client *ethclient.Client

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		stakeManagerUtils: StakeManagerMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
	}

	type args struct {
		password     string
		address      string
		addressErr   error
		stakerId     uint32
		stakerIdErr  error
		txnOpts      *bind.TransactOpts
		resetLockTxn *Types.Transaction
		resetLockErr error
		hash         common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When resetLock function executes successfully",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				addressErr:   nil,
				stakerId:     1,
				stakerIdErr:  nil,
				txnOpts:      txnOpts,
				resetLockTxn: &Types.Transaction{},
				resetLockErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting address from flags",
			args: args{
				password:     "test",
				address:      "",
				addressErr:   errors.New("address error"),
				stakerId:     1,
				stakerIdErr:  nil,
				txnOpts:      txnOpts,
				resetLockTxn: &Types.Transaction{},
				resetLockErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test 3: When there is an error in getting stakerId from flags",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				addressErr:   nil,
				stakerIdErr:  errors.New("stakerId error"),
				txnOpts:      txnOpts,
				resetLockTxn: &Types.Transaction{},
				resetLockErr: nil,
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("stakerId error"),
		},
		{
			name: "Test 4: When ResetLock transaction fails",
			args: args{
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				addressErr:   nil,
				stakerId:     1,
				stakerIdErr:  nil,
				txnOpts:      txnOpts,
				resetLockTxn: &Types.Transaction{},
				resetLockErr: errors.New("resetLock error"),
				hash:         common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("resetLock error"),
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

			GetUint32StakerIdMock = func(*pflag.FlagSet) (uint32, error) {
				return tt.args.stakerId, tt.args.stakerIdErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}

			ExtendLockMock = func(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error) {
				return tt.args.resetLockTxn, tt.args.resetLockErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.extendLock(flagSet, config)
			if got != tt.want {
				t.Errorf("Txn hash for resetLock function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for resetLock function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for resetLock function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
