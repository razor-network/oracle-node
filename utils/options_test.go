package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/utils/mocks"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

func Test_getGasPrice(t *testing.T) {
	type args struct {
		suggestedGasPrice    *big.Int
		suggestedGasPriceErr error
		config               types.Configurations
		multipliedGasPrice   *big.Int
	}
	tests := []struct {
		name          string
		args          args
		want          *big.Int
		expectedFatal bool
	}{
		{
			name: "Test 1: When config gas price is greater than suggested gas price",
			args: args{
				config: types.Configurations{
					GasPrice:      2,
					GasMultiplier: 2,
				},
				suggestedGasPrice:  big.NewInt(1).Mul(big.NewInt(1), big.NewInt(1e9)),
				multipliedGasPrice: big.NewInt(1).Mul(big.NewInt(4), big.NewInt(1e9)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(4), big.NewInt(1e9)),
			expectedFatal: false,
		},
		{
			name: "Test 2: When config gas price is less than suggested gas price",
			args: args{
				config: types.Configurations{
					GasPrice:      2,
					GasMultiplier: 2,
				},
				suggestedGasPrice:  big.NewInt(1).Mul(big.NewInt(4), big.NewInt(1e9)),
				multipliedGasPrice: big.NewInt(1).Mul(big.NewInt(8), big.NewInt(1e9)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(8), big.NewInt(1e9)),
			expectedFatal: false,
		},
		{
			name: "Test 3: When config gas price is 0",
			args: args{
				config: types.Configurations{
					GasPrice:      0,
					GasMultiplier: 2,
				},
				suggestedGasPrice:  big.NewInt(1).Mul(big.NewInt(4), big.NewInt(1e9)),
				multipliedGasPrice: big.NewInt(1).Mul(big.NewInt(8), big.NewInt(1e9)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(8), big.NewInt(1e9)),
			expectedFatal: false,
		},
		{
			name: "Test 3: When suggest gas price throws an error and config gas price has a non-zero value",
			args: args{
				config: types.Configurations{
					GasPrice:      1,
					GasMultiplier: 2,
				},
				suggestedGasPriceErr: errors.New("error in fetching gas price"),
				multipliedGasPrice:   big.NewInt(1).Mul(big.NewInt(2), big.NewInt(1e9)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(2), big.NewInt(1e9)),
			expectedFatal: false,
		},
		{
			name: "Test 4: When suggest gas price throws an error and config gas price is 0",
			args: args{
				config: types.Configurations{
					GasPrice:      0,
					GasMultiplier: 2,
				},
				suggestedGasPriceErr: errors.New("error in fetching gas price"),
				multipliedGasPrice:   big.NewInt(1).Mul(big.NewInt(2), big.NewInt(1e9)),
			},
			want:          big.NewInt(0),
			expectedFatal: false,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UtilsMock := new(mocks.Utils)
			clientUtilsMock := new(mocks.ClientUtils)

			clientUtilsMock.On("SuggestGasPriceWithRetry", mock.Anything).Return(tt.args.suggestedGasPrice, tt.args.suggestedGasPriceErr)
			UtilsMock.On("MultiplyFloatAndBigInt", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("float64")).Return(tt.args.multipliedGasPrice)

			fatal = false

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:  UtilsMock,
				ClientInterface: clientUtilsMock,
			}
			StartRazor(optionsPackageStruct)
			gasUtils := GasStruct{}
			got := gasUtils.GetGasPrice(rpcParameters, tt.args.config)
			if fatal != tt.expectedFatal {
				if got.Cmp(tt.want) != 0 {
					t.Errorf("getGasPrice() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_utils_GetTxnOpts(t *testing.T) {
	var gasPrice *big.Int

	privateKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		address         string
		nonce           uint64
		nonceErr        error
		txnOpts         *bind.TransactOpts
		txnOptsErr      error
		gasLimit        uint64
		gasLimitErr     error
		latestHeader    *Types.Header
		latestHeaderErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *bind.TransactOpts
		wantErr bool
	}{
		{
			name: "Test 1: When GetTxnOptions execute successfully",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:    2,
				txnOpts:  txnOpts,
				gasLimit: 1,
			},
			want:    txnOpts,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting private key as address is not present in keystore",
			args: args{
				address: "0x77Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:   2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When the accountManager is nil",
			args: args{
				address:  "",
				nonce:    2,
				txnOpts:  txnOpts,
				gasLimit: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting nonce",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:    2,
				nonceErr: errors.New("nonce error"),
				txnOpts:  txnOpts,
				gasLimit: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 5: When there is an error in getting transactor",
			args: args{
				address:    "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:      2,
				txnOpts:    txnOpts,
				txnOptsErr: errors.New("transactor error"),
				gasLimit:   1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When there is an error in getting gasLimit",
			args: args{
				address:     "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:       2,
				txnOpts:     txnOpts,
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    txnOpts,
			wantErr: false,
		},
		{
			name: "Test 7: When there is an rpc error in getting gasLimit",
			args: args{
				address:     "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:       2,
				txnOpts:     txnOpts,
				gasLimitErr: errors.New("504 gateway error"),
				latestHeader: &Types.Header{
					GasLimit: 500,
				},
			},
			want:    txnOpts,
			wantErr: false,
		},
		{
			name: "Test 8: When there is an rpc error in getting gasLimit and than error in getting latest header",
			args: args{
				address:     "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				nonce:       2,
				txnOpts:     txnOpts,
				gasLimitErr: errors.New("504 gateway error"),
				latestHeader: &Types.Header{
					GasLimit: 0,
				},
				latestHeaderErr: errors.New("latest header error"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	originalExitFunc := log.LogrusInstance.ExitFunc                   // Preserve the original ExitFunc
	defer func() { log.LogrusInstance.ExitFunc = originalExitFunc }() // Ensure it's reset after tests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var account types.Account
			accountManager := accounts.NewAccountManager("test_accounts")
			if tt.args.address != "" {
				account = accounts.InitAccountStruct(tt.args.address, "Test@123", accountManager)
			} else {
				account = types.Account{}
			}
			transactionData := types.TransactionOptions{
				Account: account,
			}
			utilsMock := new(mocks.Utils)
			pathMock := new(mocks.PathUtils)
			bindMock := new(mocks.BindUtils)
			clientMock := new(mocks.ClientUtils)
			gasMock := new(mocks.GasUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:  utilsMock,
				PathInterface:   pathMock,
				BindInterface:   bindMock,
				ClientInterface: clientMock,
				GasInterface:    gasMock,
			}

			utils := StartRazor(optionsPackageStruct)

			clientMock.On("GetNonceAtWithRetry", mock.Anything, mock.Anything).Return(tt.args.nonce, tt.args.nonceErr)
			gasMock.On("GetGasPrice", mock.Anything, mock.Anything).Return(gasPrice)
			bindMock.On("NewKeyedTransactorWithChainID", mock.AnythingOfType("*ecdsa.PrivateKey"), mock.AnythingOfType("*big.Int")).Return(tt.args.txnOpts, tt.args.txnOptsErr)
			gasMock.On("GetGasLimit", mock.Anything, transactionData, txnOpts).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			utilsMock.On("MultiplyFloatAndBigInt", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("float64")).Return(big.NewInt(1))
			clientMock.On("GetLatestBlockWithRetry", mock.Anything).Return(tt.args.latestHeader, tt.args.latestHeaderErr)

			got, err := utils.GetTxnOpts(rpcParameters, transactionData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTxnOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTxnOpts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetGasLimit(t *testing.T) {
	txnOpts := &bind.TransactOpts{
		GasPrice: big.NewInt(1),
		Value:    big.NewInt(1000),
	}
	var parsedData abi.ABI
	var inputData []byte
	var reader = strings.NewReader("")

	type args struct {
		transactionData     types.TransactionOptions
		parsedData          abi.ABI
		parseErr            error
		inputData           []byte
		packErr             error
		toAssign            uint16
		toAssignErr         error
		gasLimit            uint64
		gasLimitErr         error
		increaseGasLimit    uint64
		increaseGasLimitErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error
	}{
		{
			name: "Test 1: When getGasLimit executes successfully",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "stake",
					Config:     types.Configurations{GasLimitMultiplier: 2},
				},
				parsedData:       parsedData,
				inputData:        inputData,
				gasLimit:         1,
				increaseGasLimit: 2,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When method name is nil",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "",
				},
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in parsing data",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "stake",
					Config:     types.Configurations{GasLimitMultiplier: 2},
				},
				parseErr: errors.New("parse error"),
			},
			want:    0,
			wantErr: errors.New("parse error"),
		},
		{
			name: "Test 4: When there is a pack error",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "stake",
					Config:     types.Configurations{GasLimitMultiplier: 2},
				},
				parsedData: parsedData,
				packErr:    errors.New("pack error"),
			},
			want:    0,
			wantErr: errors.New("pack error"),
		},
		{
			name: "Test 5: When there is an error in estimating gasLimit",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "stake",
					Config: types.Configurations{
						GasLimitMultiplier: 2,
						GasLimitOverride:   5000000,
					},
				},
				parsedData:  parsedData,
				inputData:   inputData,
				gasLimitErr: errors.New("gasLimit error"),
			},
			want:    5000000,
			wantErr: nil,
		},
		{
			name: "Test 6: When the transaction is reveal and we get the calculated gasLimit for reveal txn",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "reveal",
					Config: types.Configurations{
						GasLimitMultiplier: 2,
						GasLimitOverride:   5000000,
					},
				},
				parsedData:       parsedData,
				inputData:        inputData,
				toAssign:         3,
				increaseGasLimit: 963728,
			},
			want:    963728,
			wantErr: nil,
		},
		{
			name: "Test 6: When the transaction is reveal and we get the error in getting toAssign during calculating gasLimit for reveal txn",
			args: args{
				transactionData: types.TransactionOptions{
					MethodName: "reveal",
					Config: types.Configurations{
						GasLimitMultiplier: 2,
						GasLimitOverride:   5000000,
					},
				},
				parsedData:  parsedData,
				inputData:   inputData,
				toAssignErr: errors.New("toAssign error"),
			},
			want:    5000000,
			wantErr: errors.New("toAssign error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.Utils)
			abiMock := new(mocks.ABIUtils)
			clientUtilsMock := new(mocks.ClientUtils)
			gasUtilsMock := new(mocks.GasUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:  utilsMock,
				ABIInterface:    abiMock,
				ClientInterface: clientUtilsMock,
				GasInterface:    gasUtilsMock,
			}

			StartRazor(optionsPackageStruct)

			abiMock.On("Parse", reader).Return(tt.args.parsedData, tt.args.parseErr)
			abiMock.On("Pack", parsedData, mock.AnythingOfType("string"), mock.Anything).Return(tt.args.inputData, tt.args.packErr)
			clientUtilsMock.On("EstimateGasWithRetry", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.gasLimit, tt.args.gasLimitErr)
			gasUtilsMock.On("IncreaseGasLimitValue", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.increaseGasLimit, tt.args.increaseGasLimitErr)
			utilsMock.On("ToAssign", mock.Anything).Return(tt.args.toAssign, tt.args.toAssignErr)

			gasUtils := GasStruct{}
			got, err := gasUtils.GetGasLimit(rpcParameters, tt.args.transactionData, txnOpts)
			if got != tt.want {
				t.Errorf("getGasLimit() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getGasLimit function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getGasLimit function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestUtilsStruct_IncreaseGasLimitValue(t *testing.T) {
	type args struct {
		gasLimit           uint64
		gasLimitMultiplier float32
		latestBlock        *Types.Header
		blockErr           error
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error
	}{
		{
			name: "Test 1: When increaseGasLimitValue() executes successfully",
			args: args{
				gasLimit:           1,
				gasLimitMultiplier: 2,
				latestBlock: &Types.Header{
					GasLimit: 3,
				},
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "Test 2: When gasLimit > latestBlock.GasLimit",
			args: args{
				gasLimit:           1,
				gasLimitMultiplier: 3,
				latestBlock: &Types.Header{
					GasLimit: 1,
				},
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 3: When gasLimit is 0",
			args: args{
				gasLimit:           0,
				gasLimitMultiplier: 2,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 4: When gasMultiplier is 0",
			args: args{
				gasLimit:           1,
				gasLimitMultiplier: 0,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Test 5: When there is an error in getting latest header",
			args: args{
				gasLimit:           1,
				gasLimitMultiplier: 2,
				blockErr:           errors.New("block error"),
			},
			want:    0,
			wantErr: errors.New("block error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientUtilsMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientUtilsMock,
			}

			StartRazor(optionsPackageStruct)

			clientUtilsMock.On("GetLatestBlockWithRetry", mock.Anything, mock.Anything).Return(tt.args.latestBlock, tt.args.blockErr)

			gasUtils := GasStruct{}
			got, err := gasUtils.IncreaseGasLimitValue(rpcParameters, tt.args.gasLimit, tt.args.gasLimitMultiplier)
			if got != tt.want {
				t.Errorf("increaseGasLimitValue() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for increaseGasLimitValue function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for increaseGasLimitValue function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetOptions(t *testing.T) {
	callOpts := bind.CallOpts{
		Pending:     false,
		BlockNumber: nil,
		Context:     context.Background(),
	}
	tests := []struct {
		name string
		want bind.CallOpts
	}{
		{
			name: "Test 1: When GetOptionsExecutes successfully",
			want: callOpts,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}

			utils := StartRazor(optionsPackageStruct)
			if got := utils.GetOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
