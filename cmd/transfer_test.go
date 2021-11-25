package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

func Test_transfer(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		transactionUtils:  TransactionMock{},
		tokenManagerUtils: TokenManagerMock{},
		flagSetUtils:      FlagSetMock{},
	}

	type args struct {
		from          string
		fromErr       error
		to            string
		toErr         error
		password      string
		balance       *big.Int
		balanceErr    error
		amount        *big.Int
		amountErr     error
		decimalAmount *big.Float
		txnOpts       *bind.TransactOpts
		transferTxn   *Types.Transaction
		transferErr   error
		transferHash  common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "When transfer function executes successfully",
			args: args{
				password:      "test",
				from:          "0x000000000000000000000000000000000000dea1",
				fromErr:       nil,
				to:            "0x000000000000000000000000000000000000dea2",
				toErr:         nil,
				balance:       big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:    nil,
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "When there is an error in getting 'from' address from flags",
			args: args{
				password:      "test",
				from:          "",
				fromErr:       errors.New("address error"),
				to:            "0x000000000000000000000000000000000000dea2",
				toErr:         nil,
				balance:       big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:    nil,
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "When there is an error in getting 'to' address from flags",
			args: args{
				password:      "test",
				from:          "0x000000000000000000000000000000000000dea1",
				fromErr:       nil,
				to:            "",
				toErr:         errors.New("address error"),
				balance:       big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:    nil,
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "When there is an error in fetching balance",
			args: args{
				password:      "test",
				from:          "0x000000000000000000000000000000000000dea1",
				fromErr:       nil,
				to:            "0x000000000000000000000000000000000000dea2",
				toErr:         nil,
				balanceErr:    errors.New("balance error"),
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("balance error"),
		},
		{
			name: "When there is an error in fetching amount",
			args: args{
				password:      "test",
				from:          "0x000000000000000000000000000000000000dea1",
				fromErr:       nil,
				to:            "0x000000000000000000000000000000000000dea2",
				toErr:         nil,
				balance:       big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				amountErr:     errors.New("amount error"),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   nil,
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("amount error"),
		},
		{
			name: "When transfer transaction fails",
			args: args{
				password:      "test",
				from:          "0x000000000000000000000000000000000000dea1",
				fromErr:       nil,
				to:            "0x000000000000000000000000000000000000dea2",
				toErr:         nil,
				balance:       big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balanceErr:    nil,
				amount:        big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				decimalAmount: big.NewFloat(1000),
				txnOpts:       txnOpts,
				transferTxn:   &Types.Transaction{},
				transferErr:   errors.New("transfer error"),
				transferHash:  common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("transfer error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssignPasswordMock = func(set *pflag.FlagSet) string {
				return tt.args.password
			}

			GetStringFromMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.from, tt.args.fromErr
			}

			GetStringToMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.to, tt.args.toErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}
			FetchBalanceMock = func(*ethclient.Client, string) (*big.Int, error) {
				return tt.args.balance, tt.args.balanceErr
			}
			AssignAmountInWeiMock = func(set *pflag.FlagSet) (*big.Int, error) {
				return tt.args.amount, tt.args.amountErr
			}
			CheckAmountAndBalanceMock = func(*big.Int, *big.Int) *big.Int {
				return tt.args.amount
			}
			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}
			GetAmountInDecimalMock = func(*big.Int) *big.Float {
				return tt.args.decimalAmount
			}
			TransferMock = func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error) {
				return tt.args.transferTxn, tt.args.transferErr
			}
			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.transferHash
			}

			got, err := utilsStruct.transfer(flagSet, config)
			if got != tt.want {
				t.Errorf("Txn hash for transfer function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for transfer function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for transfer function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
