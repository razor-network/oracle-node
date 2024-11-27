package utils

import (
	"errors"
	"math/big"
	"os"
	"razor/accounts"
	Types "razor/core/types"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestCheckError(t *testing.T) {
	type args struct {
		msg string
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1: When CheckError() executes successfully",
			args: args{
				msg: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckError(tt.args.msg, tt.args.err)
		})
	}
}

func TestCalculateBlockTime(t *testing.T) {
	type args struct {
		latestBlock        *types.Header
		latestBlockErr     error
		lastSecondBlock    *types.Header
		lastSecondBlockErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When CalculateBlockTime() executes successfully",
			args: args{
				latestBlock: &types.Header{
					Time:   123,
					Number: big.NewInt(100),
				},
				lastSecondBlock: &types.Header{
					Time: 120,
				},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in fetching latestBlock",
			args: args{
				latestBlock: &types.Header{
					Time:   123,
					Number: big.NewInt(100),
				},
				latestBlockErr: errors.New("error in fetching latestBlock"),
				lastSecondBlock: &types.Header{
					Time: 120,
				},
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching lastSecondBlock",
			args: args{
				latestBlock: &types.Header{
					Time:   123,
					Number: big.NewInt(100),
				},
				lastSecondBlock: &types.Header{
					Time: 120,
				},
				lastSecondBlockErr: errors.New("error in fetching lastSecondBlock"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientUtilsMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientUtilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientUtilsMock.On("GetLatestBlockWithRetry", mock.Anything).Return(tt.args.latestBlock, tt.args.latestBlockErr)
			clientUtilsMock.On("GetBlockByNumberWithRetry", mock.Anything, mock.Anything).Return(tt.args.lastSecondBlock, tt.args.lastSecondBlockErr)

			fatal = false

			utils.CalculateBlockTime(rpcParameters)
			if fatal != tt.expectedFatal {
				t.Error("The CalculateBlockTime function didn't execute as expected")
			}
		})
	}
}

func TestCheckEthBalanceIsZero(t *testing.T) {
	var address string

	type args struct {
		ethBalance    *big.Int
		ethBalanceErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When CheckEthBalanceIsZero() executes successfully",
			args: args{
				ethBalance: big.NewInt(1),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When CheckEthBalanceIsZero() returns zero",
			args: args{
				ethBalance: big.NewInt(0),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching ethBalance",
			args: args{
				ethBalance:    big.NewInt(1),
				ethBalanceErr: errors.New("error in fetching ethBalance"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientMock.On("BalanceAtWithRetry", mock.Anything, mock.Anything).Return(tt.args.ethBalance, tt.args.ethBalanceErr)

			fatal = false

			utils.CheckEthBalanceIsZero(rpcParameters, address)
			if fatal != tt.expectedFatal {
				t.Error("The CheckEthBalanceIsZero function didn't execute as expected")
			}
		})
	}
}

func TestCheckTransactionReceipt(t *testing.T) {
	var txHash string

	type args struct {
		tx    *types.Receipt
		txErr error
	}
	tests := []struct {
		name          string
		args          args
		want          int
		expectedFatal bool
	}{
		{
			name: "Test 1: When CheckTransactionReceipt() executes successfully",
			args: args{
				tx: &types.Receipt{},
			},
			want:          0,
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is error in getting transactionReceipt",
			args: args{
				txErr: errors.New("error in fetching transactionReceipt"),
			},
			want:          0,
			expectedFatal: false,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientMock := new(mocks.ClientUtils)
			retryMock := new(mocks.RetryUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientMock,
				RetryInterface:  retryMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientMock.On("TransactionReceipt", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.tx, tt.args.txErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			fatal = false

			utils.CheckTransactionReceipt(rpcParameters, txHash)
			if fatal != tt.expectedFatal {
				t.Error("The CheckTransactionReceipt function didn't execute as expected")
			}
		})
	}
}

func TestConnectToClient(t *testing.T) {
	var provider string
	type args struct {
		client    *ethclient.Client
		clientErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ConnectToClient() executes successfully",
			args: args{
				client: &ethclient.Client{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in ConnectToClient() function",
			args: args{
				clientErr: errors.New("error in connecting to client"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ethClientMock := new(mocks.EthClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				EthClient: ethClientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			ethClientMock.On("Dial", mock.AnythingOfType("string")).Return(tt.args.client, tt.args.clientErr)

			fatal = false

			utils.ConnectToClient(provider)
			if fatal != tt.expectedFatal {
				t.Error("The ConnectToClient function didn't execute as expected")
			}
		})
	}
}

func TestFetchBalance(t *testing.T) {
	var accountAddress string

	type args struct {
		erc20Contract *bindings.RAZOR
		balance       *big.Int
		balanceErr    error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "When FetchBalance() executes successfully",
			args: args{
				erc20Contract: &bindings.RAZOR{},
				balance:       big.NewInt(1),
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "When there is an error in fetching balance",
			args: args{
				erc20Contract: &bindings.RAZOR{},
				balanceErr:    errors.New("error in fetching balance"),
			},
			want:    big.NewInt(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			erc20Mock := new(mocks.CoinUtils)
			retryMock := new(mocks.RetryUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
				CoinInterface:  erc20Mock,
				RetryInterface: retryMock,
			}
			utils := StartRazor(optionsPackageStruct)

			erc20Mock.On("BalanceOf", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.FetchBalance(rpcParameters, accountAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBufferedState(t *testing.T) {
	type args struct {
		block       *types.Header
		buffer      int32
		stateBuffer uint64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test 1: When GetBufferedState() executes successfully",
			args: args{
				block: &types.Header{
					Time: 100,
				},
				buffer:      2,
				stateBuffer: 5,
			},

			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When blockNumber%(core.StateLength) is greater than lowerLimit",
			args: args{
				block: &types.Header{
					Time: 1080,
				},
				buffer:      2,
				stateBuffer: 5,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Test 3: When GetBufferedState() executes successfully and state we get is other than 0",
			args: args{
				block: &types.Header{
					Time: 900,
				},
				buffer:      2,
				stateBuffer: 5,
			},

			want:    -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils := StartRazor(OptionsPackageStruct{})

			got, err := utils.GetBufferedState(tt.args.block, tt.args.stateBuffer, tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBufferedState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBufferedState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpoch(t *testing.T) {
	type args struct {
		latestHeader    *types.Header
		latestHeaderErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetEpoch() executes successfully",
			args: args{
				latestHeader: &types.Header{
					Number: big.NewInt(100),
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting latestHeader",
			args: args{
				latestHeader: &types.Header{
					Number: big.NewInt(100),
				},
				latestHeaderErr: errors.New("latestHeader error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientUtilsMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientUtilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientUtilsMock.On("GetLatestBlockWithRetry", mock.Anything).Return(tt.args.latestHeader, tt.args.latestHeaderErr)

			got, err := utils.GetEpoch(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpoch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpoch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStateName(t *testing.T) {
	type args struct {
		stateNumber int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1: When state is commit",
			args: args{
				stateNumber: 0,
			},
			want: "Commit",
		},
		{
			name: "Test 2: When state is reveal",
			args: args{
				stateNumber: 1,
			},
			want: "Reveal",
		},
		{
			name: "Test 3: When state is propose",
			args: args{
				stateNumber: 2,
			},
			want: "Propose",
		},
		{
			name: "Test 4: When state is dispute",
			args: args{
				stateNumber: 3,
			},
			want: "Dispute",
		},
		{
			name: "Test 5: When state is confirm",
			args: args{
				stateNumber: 4,
			},
			want: "Confirm",
		},
		{
			name: "Test 6: When state is none of the above",
			args: args{
				stateNumber: 5,
			},
			want: "Buffer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStateName(tt.args.stateNumber); got != tt.want {
				t.Errorf("GetStateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitForBlockCompletion(t *testing.T) {
	var hashToRead string

	type args struct {
		transactionStatus int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When WaitForBlockCompletion() executes successfully",
			args: args{
				transactionStatus: 0,
			},
			want: errors.New("transaction mining unsuccessful"),
		},
		{
			name: "Test 2: When transactionStatus is 1",
			args: args{
				transactionStatus: 1,
			},
			want: nil,
		},
		{
			name: "Test 3: When transactionStatus is neither 1 nor 0",
			args: args{
				transactionStatus: 2,
			},
			want: errors.New("maximum attempts failed for transaction mining"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			timeMock := new(mocks.TimeUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
				Time:           timeMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("CheckTransactionReceipt", mock.Anything, mock.Anything).Return(tt.args.transactionStatus)
			timeMock.On("Sleep", mock.Anything).Return()

			gotErr := utils.WaitForBlockCompletion(rpcParameters, hashToRead)
			if gotErr == nil || tt.want == nil {
				if gotErr != tt.want {
					t.Errorf("Error for WaitForBlockCompletion function, got = %v, want %v", gotErr, tt.want)
				}
			} else {
				if gotErr.Error() != tt.want.Error() {
					t.Errorf("Error for WaitForBlockCompletion function, got = %v, want %v", gotErr, tt.want)
				}
			}
		})
	}
}

func TestIsFlagPassed(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1: When IsFlagPassed() executes successfully",
			args: args{
				name: "password",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsPackageStruct := OptionsPackageStruct{}
			utils := StartRazor(optionsPackageStruct)
			if got := utils.IsFlagPassed(tt.args.name); got != tt.want {
				t.Errorf("IsFlagPassed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitTillNextNSecs(t *testing.T) {
	type args struct {
		waitTime int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1: When WaitTillNextNSecs() executes successfully",
			args: args{
				waitTime: 1,
			},
		},
		{
			name: "Test 2: When waitTime is negative",
			args: args{
				waitTime: -1,
			},
		},
	}
	for _, tt := range tests {
		timeMock := new(mocks.TimeUtils)

		optionsPackageStruct := OptionsPackageStruct{
			Time: timeMock,
		}
		utils := StartRazor(optionsPackageStruct)
		timeMock.On("Sleep", mock.Anything).Return()

		t.Run(tt.name, func(t *testing.T) {
			utils.WaitTillNextNSecs(tt.args.waitTime)
		})
	}
}

func TestIsValidAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1: When correct address is passed",
			args: args{
				address: "0x8797EA6306881D74c4311C08C0Ca2C0a76dDC90e",
			},
			want: true,
		},
		{
			name: "Test 2: When incorrect address is passed",
			args: args{
				address: "0x8797EA6306881D74c4311C08C0Ca2C0a76dDC90z",
			},
			want: false,
		},
		{
			name: "Test 2: When nil is passed",
			args: args{
				address: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAddress(tt.args.address); got != tt.want {
				t.Errorf("IsValidAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidateAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
		wantErr     error
	}{
		{
			name: "Test 1: When correct address is passed",
			args: args{
				address: "0x8797EA6306881D74c4311C08C0Ca2C0a76dDC90e",
			},
			wantAddress: "0x8797EA6306881D74c4311C08C0Ca2C0a76dDC90e",
			wantErr:     nil,
		},
		{
			name: "Test 2: When incorrect address is passed",
			args: args{
				address: "0x8797EA6306881D74c4311C08C0Ca2C0a76dDC90z",
			},
			wantAddress: "",
			wantErr:     errors.New("invalid address"),
		},
		{
			name: "Test 2: When nil is passed",
			args: args{
				address: "",
			},
			wantAddress: "",
			wantErr:     errors.New("invalid address"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ValidateAddress(tt.args.address)
			if got != tt.wantAddress {
				t.Errorf("ValidateAddress() returns address = %v, want address = %v", got, tt.wantAddress)
			}
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for ValidateAddress(), got error = %v, want error = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ValidateAddress(), got error = %v, want error = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func TestAssignStakerId(t *testing.T) {
	var flagSet *pflag.FlagSet
	var address string

	type args struct {
		flagPassed         bool
		flagSetStakerId    uint32
		flagSetStakerIdErr error
		stakerId           uint32
		stakerIdErr        error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
		wantErr       error
	}{
		{
			name: "Test 1: When AssignStakerId() executes successfully and flag is not passed",
			args: args{
				flagPassed: false,
				stakerId:   1,
			},
			expectedFatal: false,
			wantErr:       nil,
		},
		{
			name: "Test 2: When AssignStakerId() executes successfully and flag is passed",
			args: args{
				flagPassed:      true,
				flagSetStakerId: 1,
			},
			expectedFatal: false,
			wantErr:       nil,
		},
		{
			name: "Test 3: When there is an error in getting stakerId",
			args: args{
				flagPassed:  false,
				stakerIdErr: errors.New("stakerId error"),
			},
			expectedFatal: false,
			wantErr:       errors.New("stakerId error"),
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("IsFlagPassed", mock.AnythingOfType("string")).Return(tt.args.flagPassed)
			utilsMock.On("GetUint32", mock.Anything, mock.AnythingOfType("string")).Return(tt.args.flagSetStakerId, tt.args.flagSetStakerIdErr)
			utilsMock.On("GetStakerId", mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)

			fatal = false

			_, err := utils.AssignStakerId(rpcParameters, flagSet, address)
			if fatal != tt.expectedFatal {
				t.Error("The AssignStakerId function didn't execute as expected")
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for AssignStakerId function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for AssignStakerId function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestAssignLogFile(t *testing.T) {
	var flagSet *pflag.FlagSet
	type args struct {
		isFlagPassed bool
		fileName     string
		fileNameErr  error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When AssignLogFile() executes successfully",
			args: args{
				isFlagPassed: true,
				fileName:     "",
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting logFile name",
			args: args{
				isFlagPassed: true,
				fileNameErr:  errors.New("fileName error"),
				fileName:     "",
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			flagSetMock := new(mocks.FlagSetUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:   utilsMock,
				FlagSetInterface: flagSetMock,
			}
			StartRazor(optionsPackageStruct)
			fatal = false

			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
			flagSetMock.On("GetLogFileName", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.fileName, tt.args.fileNameErr)

			fileUtils := FileStruct{}
			fileUtils.AssignLogFile(flagSet, Types.Configurations{})
			if fatal != tt.expectedFatal {
				t.Error("The AssignLogFile function didn't execute as expected")
			}
		})
	}
}

func TestGetRemainingTimeOfCurrentState(t *testing.T) {
	type args struct {
		block         *types.Header
		stateBuffer   uint64
		bufferPercent int32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test 1: When GetRemainingTimeOfCurrentState() executes successfully",
			args: args{
				block:       &types.Header{},
				stateBuffer: 5,
			},
			want:    85,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils := StartRazor(OptionsPackageStruct{})
			got, err := utils.GetRemainingTimeOfCurrentState(tt.args.block, tt.args.stateBuffer, tt.args.bufferPercent)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemainingTimeOfCurrentState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRemainingTimeOfCurrentState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSalt(t *testing.T) {
	type args struct {
		epoch   uint32
		medians []*big.Int
	}
	tests := []struct {
		name string
		args args
		want [32]byte
	}{
		{
			name: "When CalculateSalt() is executed successfully",
			args: args{
				epoch:   1,
				medians: []*big.Int{},
			},
			want: [32]byte{3, 188, 235, 65, 42, 140, 151, 61, 187, 150, 15, 19, 83, 186, 145, 207, 108, 161, 13, 253, 226, 28, 145, 16, 84, 207, 30, 97, 240, 210, 142, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			if got := ut.CalculateSalt(tt.args.epoch, tt.args.medians); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrng(t *testing.T) {
	type args struct {
		max        uint32
		prngHashes []byte
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "When Prng() is calculated correctly",
			args: args{
				max:        10,
				prngHashes: []byte{},
			},
			want: big.NewInt(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			if got := ut.Prng(tt.args.max, tt.args.prngHashes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prng() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEstimateBlockNumberAtEpochBeginning(t *testing.T) {
	var currentBlockNumber *big.Int
	type args struct {
		block            *types.Header
		blockErr         error
		previousBlock    *types.Header
		previousBlockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When EstimateBlockNumberAtEpochBeginning() is executed successfully",
			args: args{
				block:         &types.Header{Time: 1, Number: big.NewInt(100)},
				previousBlock: &types.Header{Time: 20, Number: big.NewInt(12)},
			},
			want:    big.NewInt(10),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting block",
			args: args{
				blockErr: errors.New("error in getting block"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:  utilsMock,
				ClientInterface: clientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientMock.On("GetBlockByNumberWithRetry", mock.Anything, mock.Anything).Return(tt.args.block, tt.args.blockErr)
			clientMock.On("GetBlockByNumberWithRetry", mock.Anything, mock.Anything).Return(tt.args.previousBlock, tt.args.previousBlockErr)
			got, err := utils.EstimateBlockNumberAtEpochBeginning(rpcParameters, currentBlockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstimateBlockNumberAtEpochBeginning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EstimateBlockNumberAtEpochBeginning() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveDataToCommitJsonFile(t *testing.T) {
	var (
		filePath   string
		epoch      uint32
		commitData Types.CommitData
		commitment [32]byte
	)
	type args struct {
		jsonData     []byte
		jsonDataErr  error
		writeFileErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When SaveDataToCommitJsonFile() executes successfully",
			args: args{
				jsonData: []byte{},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonData",
			args: args{
				jsonDataErr: errors.New("error in getting jsonData"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in writing file",
			args: args{
				jsonData:     []byte{},
				writeFileErr: errors.New("error in writing file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
			}
			StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)

			fileUtils := FileStruct{}
			if err := fileUtils.SaveDataToCommitJsonFile(filePath, epoch, commitData, commitment); (err != nil) != tt.wantErr {
				t.Errorf("SaveDataToCommitJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveDataToProposeJsonFile(t *testing.T) {
	var (
		filePath    string
		proposeData Types.ProposeFileData
	)

	type args struct {
		jsonData     []byte
		jsonDataErr  error
		writeFileErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When SaveDataToProposeJsonFile() executes successfully",
			args: args{
				jsonData: []byte{},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonData",
			args: args{
				jsonDataErr: errors.New("error in getting jsonData"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in writing file",
			args: args{
				jsonData:     []byte{},
				writeFileErr: errors.New("error in writing file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
			}
			StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)

			fileUtils := FileStruct{}
			if err := fileUtils.SaveDataToProposeJsonFile(filePath, proposeData); (err != nil) != tt.wantErr {
				t.Errorf("SaveDataToProposeJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveDataToDisputeJsonFile(t *testing.T) {
	var (
		filePath      string
		bountyIdQueue []uint32
	)
	type args struct {
		jsonData     []byte
		jsonDataErr  error
		writeFileErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When SaveDataToDisputeJsonFile() executes successfully",
			args: args{
				jsonData: []byte{},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonData",
			args: args{
				jsonDataErr: errors.New("error in getting jsonData"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in writing file",
			args: args{
				jsonData:     []byte{},
				writeFileErr: errors.New("error in writing file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
			}
			StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)

			fileUtils := FileStruct{}
			if err := fileUtils.SaveDataToDisputeJsonFile(filePath, bountyIdQueue); (err != nil) != tt.wantErr {
				t.Errorf("SaveDataToDisputeJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadFromCommitJsonFile(t *testing.T) {
	var filePath string
	type args struct {
		jsonFile     *os.File
		jsonFileErr  error
		byteValue    []byte
		byteValueErr error
		unmarshalErr error
	}
	tests := []struct {
		name    string
		args    args
		want    Types.CommitFileData
		wantErr bool
	}{
		{
			name: "Test 1: When ReadFromCommitJsonFile() executes successfully",
			args: args{
				jsonFile:  &os.File{},
				byteValue: []byte{},
			},
			want:    Types.CommitFileData{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonFile",
			args: args{
				jsonFileErr: errors.New("error in getting jsonFile"),
			},
			want:    Types.CommitFileData{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting byteValue",
			args: args{
				jsonFile:     &os.File{},
				byteValueErr: errors.New("error in getting byteValue"),
			},
			want:    Types.CommitFileData{},
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in unmarshal",
			args: args{
				jsonFile:     &os.File{},
				byteValue:    []byte{},
				unmarshalErr: errors.New("error in unmarshal"),
			},
			want:    Types.CommitFileData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)
			ioMock := new(mocks.IOUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
				IOInterface:   ioMock,
			}
			StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			fileUtils := FileStruct{}
			got, err := fileUtils.ReadFromCommitJsonFile(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromCommitJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFromCommitJsonFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFromProposeJsonFile(t *testing.T) {
	var filePath string
	type args struct {
		jsonFile     *os.File
		jsonFileErr  error
		byteValue    []byte
		byteValueErr error
		unmarshalErr error
	}
	tests := []struct {
		name    string
		args    args
		want    Types.ProposeFileData
		wantErr bool
	}{
		{
			name: "Test 1: When ReadFromProposeJsonFile() executes successfully",
			args: args{
				jsonFile:  &os.File{},
				byteValue: []byte{},
			},
			want:    Types.ProposeFileData{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonFile",
			args: args{
				jsonFileErr: errors.New("error in getting jsonFile"),
			},
			want:    Types.ProposeFileData{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting byteValue",
			args: args{
				jsonFile:     &os.File{},
				byteValueErr: errors.New("error in getting byteValue"),
			},
			want:    Types.ProposeFileData{},
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in unmarshal",
			args: args{
				jsonFile:     &os.File{},
				byteValue:    []byte{},
				unmarshalErr: errors.New("error in unmarshal"),
			},
			want:    Types.ProposeFileData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)
			ioMock := new(mocks.IOUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
				IOInterface:   ioMock,
			}
			StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			fileUtils := FileStruct{}
			got, err := fileUtils.ReadFromProposeJsonFile(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromProposeJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFromProposeJsonFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFromDisputeJsonFile(t *testing.T) {
	var filePath string
	type args struct {
		jsonFile     *os.File
		jsonFileErr  error
		byteValue    []byte
		byteValueErr error
		unmarshalErr error
	}
	tests := []struct {
		name    string
		args    args
		want    Types.DisputeFileData
		wantErr bool
	}{
		{
			name: "Test 1: When ReadFromDisputeJsonFile() executes successfully",
			args: args{
				jsonFile:  &os.File{},
				byteValue: []byte{},
			},
			want:    Types.DisputeFileData{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting jsonFile",
			args: args{
				jsonFileErr: errors.New("error in getting jsonFile"),
			},
			want:    Types.DisputeFileData{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting byteValue",
			args: args{
				jsonFile:     &os.File{},
				byteValueErr: errors.New("error in getting byteValue"),
			},
			want:    Types.DisputeFileData{},
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in unmarshal",
			args: args{
				jsonFile:     &os.File{},
				byteValue:    []byte{},
				unmarshalErr: errors.New("error in unmarshal"),
			},
			want:    Types.DisputeFileData{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)
			ioMock := new(mocks.IOUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
				IOInterface:   ioMock,
			}
			StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			fileUtils := FileStruct{}
			got, err := fileUtils.ReadFromDisputeJsonFile(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromDisputeJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFromDisputeJsonFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_CheckPassword(t *testing.T) {
	type args struct {
		account Types.Account
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When password is correct",
			args: args{
				account: Types.Account{
					Address:        "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
					Password:       "Test@123",
					AccountManager: accounts.NewAccountManager("test_accounts"),
				},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When password is incorrect",
			args: args{
				account: Types.Account{
					Address:        "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
					Password:       "Test@456",
					AccountManager: accounts.NewAccountManager("test_accounts"),
				},
			},
			wantErr: true,
		},
		{
			name: "Test 3: When address or keystore path provided is not present",
			args: args{
				account: Types.Account{
					Address:        "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
					Password:       "Test@123",
					AccountManager: accounts.NewAccountManager("test_accounts_1"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			if err := ut.CheckPassword(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUtilsStruct_AccountManagerForKeystore(t *testing.T) {
	type args struct {
		path    string
		pathErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When account manager for keystore is returned successfully",
			args: args{
				path: "test_accounts",
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				pathErr: errors.New("path error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathMock := new(mocks.PathUtils)
			optionsPackageStruct := OptionsPackageStruct{
				PathInterface: pathMock,
			}

			utils := StartRazor(optionsPackageStruct)

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)

			_, err := utils.AccountManagerForKeystore()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountManagerForKeystore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
