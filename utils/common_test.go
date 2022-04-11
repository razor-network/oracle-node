package utils

import (
	"bufio"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"os"
	Types "razor/core/types"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"
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
	var client *ethclient.Client

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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:  utilsMock,
				ClientInterface: clientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.latestBlock, tt.args.latestBlockErr)
			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.lastSecondBlock, tt.args.lastSecondBlockErr)

			fatal = false

			utils.CalculateBlockTime(client)
			if fatal != tt.expectedFatal {
				t.Error("The CalculateBlockTime function didn't execute as expected")
			}
		})
	}
}

func TestCheckEthBalanceIsZero(t *testing.T) {
	var client *ethclient.Client
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientMock.On("BalanceAt", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.ethBalance, tt.args.ethBalanceErr)

			fatal = false

			utils.CheckEthBalanceIsZero(client, address)
			if fatal != tt.expectedFatal {
				t.Error("The CheckEthBalanceIsZero function didn't execute as expected")
			}
		})
	}
}

func TestCheckTransactionReceipt(t *testing.T) {
	var client *ethclient.Client
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientMock,
			}
			utils := StartRazor(optionsPackageStruct)

			clientMock.On("TransactionReceipt", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.tx, tt.args.txErr)

			fatal = false

			utils.CheckTransactionReceipt(client, txHash)
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

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
	var client *ethclient.Client
	var accountAddress string
	var callOpts bind.CallOpts

	type args struct {
		coinContract *bindings.RAZOR
		balance      *big.Int
		balanceErr   error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
		wantErr       error
	}{
		{
			name: "When FetchBalance() executes successfully",
			args: args{
				coinContract: &bindings.RAZOR{},
				balance:      big.NewInt(1),
			},
			expectedFatal: false,
			wantErr:       nil,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.Utils)
			coinMock := new(mocks.CoinUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
				CoinInterface:  coinMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetTokenManager", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.coinContract)
			utilsMock.On("GetOptions").Return(callOpts)
			coinMock.On("BalanceOf", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)

			fatal = false

			_, err := utils.FetchBalance(client, accountAddress)
			if fatal != tt.expectedFatal {
				t.Error("The FetchBalance function didn't execute as expected")
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for FetchBalance function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for fetchBalance function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetDelayedState(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		block    *types.Header
		blockErr error
		buffer   int32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test 1: When GetDelayedState() executes successfully",
			args: args{
				block: &types.Header{
					Time: 100,
				},
				buffer: 2,
			},

			want:    0,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting block",
			args: args{
				block: &types.Header{
					Number: big.NewInt(100),
				},
				blockErr: errors.New("block error"),
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "Test 3: When blockNumber%(core.StateLength) is greater than lowerLimit",
			args: args{
				block: &types.Header{
					Time: 1080,
				},
				buffer: 2,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Test 4: When GetDelayedState() executes successfully and state we get is other than 0",
			args: args{
				block: &types.Header{
					Time: 900,
				},
				buffer: 2,
			},

			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}

			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.block, tt.args.blockErr)

			got, err := utils.GetDelayedState(client, tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDelayedState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDelayedState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpoch(t *testing.T) {
	var client *ethclient.Client

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
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.latestHeader, tt.args.latestHeaderErr)

			got, err := utils.GetEpoch(client)
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
			want: "-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)
			if got := utils.GetStateName(tt.args.stateNumber); got != tt.want {
				t.Errorf("GetStateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDataFromFile(t *testing.T) {
	var fileName string

	type args struct {
		file    *os.File
		fileErr error
		scanner *bufio.Scanner
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
		wantErr       error
	}{
		{
			name: "Test 1: When ReadDataFromFile() executes successfully",
			args: args{
				file:    &os.File{},
				scanner: &bufio.Scanner{},
			},
			expectedFatal: false,
			wantErr:       errors.New("bufio.Scanner: token too long"),
		},
		{
			name: "Test 2: When there is an error in getting file",
			args: args{
				fileErr: errors.New("error in getting file"),
			},
			expectedFatal: false,
			wantErr:       errors.New("error in getting file"),
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osMock := new(mocks.OSUtils)
			bufioMock := new(mocks.BufioUtils)

			optionsPackageStruct := OptionsPackageStruct{
				OS:    osMock,
				Bufio: bufioMock,
			}
			utils := StartRazor(optionsPackageStruct)

			osMock.On("Open", mock.AnythingOfType("string")).Return(tt.args.file, tt.args.fileErr)
			bufioMock.On("NewScanner", mock.Anything).Return(tt.args.scanner)

			fatal = false

			_, _, err := utils.ReadDataFromFile(fileName)
			if fatal != tt.expectedFatal {
				t.Error("The ReadDataFromFile function didn't execute as expected")
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for ReadDataFromFile function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ReadDataFromFile function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestSaveDataToFile(t *testing.T) {
	var file *os.File
	var fileName string
	var epoch uint32

	type args struct {
		committedData []*big.Int
		file          *os.File
		fileErr       error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When SaveDataToFile() executes successfully",
			args: args{
				committedData: []*big.Int{big.NewInt(2)},
				file:          file,
			},
			wantErr: true,
		},
		{
			name: "Test 2: When there is an error inn getting file",
			args: args{
				committedData: []*big.Int{big.NewInt(2)},
				fileErr:       errors.New("error in fetching file"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an empty committedData",
			args: args{
				committedData: []*big.Int{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				OS: osMock,
			}
			utils := StartRazor(optionsPackageStruct)

			osMock.On("OpenFile", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(tt.args.file, tt.args.fileErr)

			if err := utils.SaveDataToFile(fileName, epoch, tt.args.committedData); (err != nil) != tt.wantErr {
				t.Errorf("SaveDataToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWaitForBlockCompletion(t *testing.T) {
	var client *ethclient.Client
	var hashToRead string

	type args struct {
		transactionStatus int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1: When WaitForBlockCompletion() executes successfully",
			args: args{
				transactionStatus: 0,
			},
			want: 0,
		},
		{
			name: "Test 2: When transactionStatus is 1",
			args: args{
				transactionStatus: 1,
			},
			want: 1,
		},
		{
			name: "Test 3: When transactionStatus is neither 1 nor 0",
			args: args{
				transactionStatus: 2,
			},
			want: 0,
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

			utilsMock.On("CheckTransactionReceipt", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.transactionStatus)
			timeMock.On("Sleep", mock.Anything).Return()

			if got := utils.WaitForBlockCompletion(client, hashToRead); got != tt.want {
				t.Errorf("WaitForBlockCompletion() = %v, want %v", got, tt.want)
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

func TestAssignStakerId(t *testing.T) {
	var flagSet *pflag.FlagSet
	var client *ethclient.Client
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

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("IsFlagPassed", mock.AnythingOfType("string")).Return(tt.args.flagPassed)
			utilsMock.On("GetUint32", mock.Anything, mock.AnythingOfType("string")).Return(tt.args.flagSetStakerId, tt.args.flagSetStakerIdErr)
			utilsMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)

			fatal = false

			_, err := utils.AssignStakerId(flagSet, client, address)
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
		name string
		args args
	}{
		{
			name: "Test 1: When AssignLogFile() executes successfully",
			args: args{
				isFlagPassed: true,
				fileName:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			flagSetMock := new(mocks.FlagSetUtils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface:   utilsMock,
				FlagSetInterface: flagSetMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
			flagSetMock.On("GetLogFileName", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.fileName, tt.args.fileNameErr)
			utils.AssignLogFile(flagSet)
		})
	}
}

func TestGetRemainingTimeOfCurrentState(t *testing.T) {
	var (
		client        *ethclient.Client
		bufferPercent int32
	)
	type args struct {
		block    *types.Header
		blockErr error
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
				block: &types.Header{},
			},
			want:    355,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting block",
			args: args{
				blockErr: errors.New("error in getting block"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.block, tt.args.blockErr)
			got, err := utils.GetRemainingTimeOfCurrentState(client, bufferPercent)
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
		medians []uint32
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
				medians: []uint32{},
			},
			want: [32]byte{81, 248, 27, 205, 252, 50, 74, 13, 255, 43, 91, 236, 157, 146, 226, 28, 190, 188, 77, 94, 41, 211, 163, 211, 13, 227, 224, 63, 190, 171, 141, 127},
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

func TestCalculateBlockNumberAtEpochBeginning(t *testing.T) {
	var (
		client             *ethclient.Client
		epochLength        int64
		currentBlockNumber *big.Int
	)
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
			name: "Test 1: When CalculateBlockNumberAtEpochBeginning() is executed successfully",
			args: args{
				block:         &types.Header{Time: 1, Number: big.NewInt(1)},
				previousBlock: &types.Header{Time: 20, Number: big.NewInt(1)},
			},
			want:    big.NewInt(-359),
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

			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.block, tt.args.blockErr)
			clientMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.previousBlock, tt.args.previousBlockErr)
			got, err := utils.CalculateBlockNumberAtEpochBeginning(client, epochLength, currentBlockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateBlockNumberAtEpochBeginning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateBlockNumberAtEpochBeginning() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveDataToCommitJsonFile(t *testing.T) {
	var (
		filePath   string
		epoch      uint32
		commitData Types.CommitData
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
			utils := StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)

			if err := utils.SaveDataToCommitJsonFile(filePath, epoch, commitData); (err != nil) != tt.wantErr {
				t.Errorf("SaveDataToCommitJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveDataToProposeJsonFile(t *testing.T) {
	var (
		filePath    string
		epoch       uint32
		proposeData Types.ProposeData
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
			utils := StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)
			if err := utils.SaveDataToProposeJsonFile(filePath, epoch, proposeData); (err != nil) != tt.wantErr {
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
			utils := StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonData, tt.args.jsonDataErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)
			if err := utils.SaveDataToDisputeJsonFile(filePath, bountyIdQueue); (err != nil) != tt.wantErr {
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
			ioutilMock := new(mocks.IoutilUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface:   jsonMock,
				OS:              osMock,
				IoutilInterface: ioutilMock,
			}
			utils := StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioutilMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)
			got, err := utils.ReadFromCommitJsonFile(filePath)
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
			ioutilMock := new(mocks.IoutilUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface:   jsonMock,
				OS:              osMock,
				IoutilInterface: ioutilMock,
			}
			utils := StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioutilMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			got, err := utils.ReadFromProposeJsonFile(filePath)
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
			ioutilMock := new(mocks.IoutilUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface:   jsonMock,
				OS:              osMock,
				IoutilInterface: ioutilMock,
			}
			utils := StartRazor(optionsPackageStruct)
			osMock.On("Open", mock.Anything).Return(tt.args.jsonFile, tt.args.jsonFileErr)
			ioutilMock.On("ReadAll", mock.Anything).Return(tt.args.byteValue, tt.args.byteValueErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			got, err := utils.ReadFromDisputeJsonFile(filePath)
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
