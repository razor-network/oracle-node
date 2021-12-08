package utils

import (
	"context"
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reflect"
	"testing"
)

func TestBalanceAtWithRetry(t *testing.T) {
	var client *ethclient.Client
	var account common.Address

	razorUtils := RazorUtilsMock{}
	type args struct {
		balance    *big.Int
		balanceErr error
	}

	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When BalanceAtWithRetry() executes successfully",
			args: args{
				balance: big.NewInt(1000),
			},
			want:    big.NewInt(1000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting balance",
			args: args{
				balanceErr: errors.New("balance error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BalanceAtMock = func(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error) {
				return tt.args.balance, tt.args.balanceErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}

			got, err := BalanceAtWithRetry(client, account, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceAtWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BalanceAtWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEstimateGasWithRetry(t *testing.T) {
	var client *ethclient.Client
	var message ethereum.CallMsg
	razorUtils := RazorUtilsMock{}

	type args struct {
		gasLimit    uint64
		gasLimitErr error
	}

	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Test 1: When EstimateGasWithRetry executes successfully",
			args: args{
				gasLimit: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting gasLimit",
			args: args{
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EstimateGasMock = func(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error) {
				return tt.args.gasLimit, tt.args.gasLimitErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}

			got, err := EstimateGasWithRetry(client, message, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstimateGasWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EstimateGasWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLogsWithRetry(t *testing.T) {
	var client *ethclient.Client
	var query ethereum.FilterQuery

	razorUtils := RazorUtilsMock{}
	type args struct {
		logs    []types.Log
		logsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Log
		wantErr bool
	}{
		{
			name: "Test 1: When FilterLogsWithRetry executes successfully",
			args: args{
				logs: []types.Log{},
			},
			want:    []types.Log{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting filterLogs",
			args: args{
				logsErr: errors.New("logs error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FilterLogsMock = func(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return tt.args.logs, tt.args.logsErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}

			got, err := FilterLogsWithRetry(client, query, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterLogsWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterLogsWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLatestBlockWithRetry(t *testing.T) {
	var client *ethclient.Client

	razorUtils := RazorUtilsMock{}

	type args struct {
		latestHeader    *types.Header
		latestHeaderErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Header
		wantErr bool
	}{
		{
			name: "Test 1: When GetLatestBlockWithRetry executes successfully",
			args: args{
				latestHeader: &types.Header{},
			},
			want:    &types.Header{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting latestHeader",
			args: args{
				latestHeaderErr: errors.New("header error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeaderByNumberMock = func(*ethclient.Client, context.Context, *big.Int) (*types.Header, error) {
				return tt.args.latestHeader, tt.args.latestHeaderErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}
			got, err := GetLatestBlockWithRetry(client, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestBlockWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatestBlockWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPendingNonceAtWithRetry(t *testing.T) {
	var client *ethclient.Client
	var accountAddress common.Address

	razorUtils := RazorUtilsMock{}
	type args struct {
		nonce    uint64
		nonceErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Test 1: When BalanceAtWithRetry() executes successfully",
			args: args{
				nonce: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting nonce",
			args: args{
				nonceErr: errors.New("nonce error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PendingNonceAtMock = func(*ethclient.Client, context.Context, common.Address) (uint64, error) {
				return tt.args.nonce, tt.args.nonceErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}

			got, err := GetPendingNonceAtWithRetry(client, accountAddress, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingNonceAtWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPendingNonceAtWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuggestGasPriceWithRetry(t *testing.T) {
	var client *ethclient.Client
	razorUtils := RazorUtilsMock{}

	type args struct {
		gasPrice    *big.Int
		gasPriceErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When SuggestGasPriceWithRetry() executes successfully",
			args: args{
				gasPrice: big.NewInt(1),
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting gasPrice",
			args: args{
				gasPriceErr: errors.New("gasPrice Error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuggestGasPriceMock = func(*ethclient.Client, context.Context) (*big.Int, error) {
				return tt.args.gasPrice, tt.args.gasPriceErr
			}

			RetryAttemptsMock = func(uint) retry.Option {
				return retry.Attempts(1)
			}

			got, err := SuggestGasPriceWithRetry(client, razorUtils)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuggestGasPriceWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuggestGasPriceWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}
