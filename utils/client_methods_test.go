package utils

import (
	"context"
	"errors"
	"math/big"
	"razor/utils/mocks"
	"reflect"
	"testing"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

func TestOptionUtilsStruct_SuggestGasPriceWithRetry(t *testing.T) {
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

			StartRazor(optionsPackageStruct)

			clientMock.On("SuggestGasPrice", mock.AnythingOfType("*ethclient.Client"), context.Background()).Return(tt.args.gasPrice, tt.args.gasPriceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.SuggestGasPriceWithRetry(rpcParameters)
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

			StartRazor(optionsPackageStruct)
			clientMock.On("BalanceAt", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("common.Address"), mock.AnythingOfType("*big.Int")).Return(tt.args.balance, tt.args.balanceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.BalanceAtWithRetry(rpcParameters, account)
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

			StartRazor(optionsPackageStruct)
			clientMock.On("EstimateGas", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("ethereum.CallMsg")).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.EstimateGasWithRetry(rpcParameters, message)
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

			StartRazor(optionsPackageStruct)
			clientMock.On("FilterLogs", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("ethereum.FilterQuery")).Return(tt.args.logs, tt.args.logsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.FilterLogsWithRetry(rpcParameters, query)
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

			StartRazor(optionsPackageStruct)
			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("*big.Int")).Return(tt.args.latestHeader, tt.args.latestHeaderErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
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

func TestUtilsStruct_GetNonceAtWithRetry(t *testing.T) {
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

			StartRazor(optionsPackageStruct)
			clientMock.On("NonceAt", mock.AnythingOfType("*ethclient.Client"), context.Background(), mock.AnythingOfType("common.Address")).Return(tt.args.nonce, tt.args.nonceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.GetNonceAtWithRetry(rpcParameters, accountAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNonceAtWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNonceAtWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientStruct_GetBlockByNumberWithRetry(t *testing.T) {
	var blockNumber *big.Int

	type args struct {
		header    *types.Header
		headerErr error
	}

	tests := []struct {
		name    string
		args    args
		want    *types.Header
		wantErr bool
	}{
		{
			name: "Test 1: When GetBlockByNumberWithRetry executes successfully",
			args: args{
				header: &types.Header{
					Number: big.NewInt(123),
				},
			},
			want: &types.Header{
				Number: big.NewInt(123),
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting block header",
			args: args{
				headerErr: errors.New("header error"),
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

			StartRazor(optionsPackageStruct)
			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), context.Background(), blockNumber).Return(tt.args.header, tt.args.headerErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			clientUtils := ClientStruct{}
			got, err := clientUtils.GetBlockByNumberWithRetry(rpcParameters, blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByNumberWithRetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) || (got != nil && got.Number.Cmp(tt.want.Number) != 0) {
				t.Errorf("GetBlockByNumberWithRetry() got = %v, want %v", got, tt.want)
			}
		})
	}
}
