package utils

import (
	"context"
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestOptionUtilsStruct_SuggestGasPriceWithRetry(t *testing.T) {
	var client *ethclient.Client

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

			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)

			clientMock.On("SuggestGasPrice", mock.AnythingOfType("*ethclient.Client"), context.Background()).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.SuggestGasPriceWithRetry(client)
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

func TestUtilsStruct_BalanceAtWithRetry(t *testing.T) {
	var client *ethclient.Client
	var account common.Address

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
			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)
			clientMock.On("BalanceAt", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("common.Address"), mock.AnythingOfType("*big.Int")).Return(tt.args.balance, tt.args.balanceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.BalanceAtWithRetry(client, account)
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

func TestUtilsStruct_EstimateGasWithRetry(t *testing.T) {
	var client *ethclient.Client
	var message ethereum.CallMsg

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
			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)
			clientMock.On("EstimateGas", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("ethereum.CallMsg")).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.EstimateGasWithRetry(client, message)
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

func TestUtilsStruct_FilterLogsWithRetry(t *testing.T) {
	var client *ethclient.Client
	var query ethereum.FilterQuery

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
			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)
			clientMock.On("FilterLogs", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("ethereum.FilterQuery")).Return(tt.args.logs, tt.args.logsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.FilterLogsWithRetry(client, query)
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

func TestUtilsStruct_GetLatestBlockWithRetry(t *testing.T) {
	var client *ethclient.Client

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
			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)
			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("*big.Int")).Return(tt.args.latestHeader, tt.args.latestHeaderErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetLatestBlockWithRetry(client)
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

func TestUtilsStruct_GetPendingNonceAtWithRetry(t *testing.T) {
	var client *ethclient.Client
	var accountAddress common.Address

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
			retryMock := new(mocks.RetryUtils)
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:  retryMock,
				ClientInterface: clientMock,
			}

			utils := StartRazor(optionsPackageStruct)
			clientMock.On("PendingNonceAt", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("common.Address")).Return(tt.args.nonce, tt.args.nonceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetPendingNonceAtWithRetry(client, accountAddress)
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
