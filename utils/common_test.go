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
	"razor/pkg/bindings"
	"razor/utils/mocks"
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			clientMock := new(mocks.ClientUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
				Client:         clientMock,
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
				Client: clientMock,
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
				Client: clientMock,
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

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetTokenManager", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.coinContract)
			utilsMock.On("GetOptions").Return(callOpts)
			utilsMock.On("BalanceOf", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.balance, tt.args.balanceErr)

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
					Number: big.NewInt(100),
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
					Number: big.NewInt(900),
				},
				buffer: 2,
			},
			want:    -1,
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
			ut := &UtilsStruct{}
			if got := ut.GetStateName(tt.args.stateNumber); got != tt.want {
				t.Errorf("GetStateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadCommittedDataFromFile(t *testing.T) {
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
			name: "Test 1: When ReadCommittedDataFromFile() executes successfully",
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

			_, _, err := utils.ReadCommittedDataFromFile(fileName)
			if fatal != tt.expectedFatal {
				t.Error("The ReadCommittedDataFromFile function didn't execute as expected")
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for ReadCommittedDataFromFile function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ReadCommittedDataFromFile function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestSaveCommittedDataToFile(t *testing.T) {
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
			name: "Test 1: When SaveCommittedDataToFile() executes successfully",
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

			if err := utils.SaveCommittedDataToFile(fileName, epoch, tt.args.committedData); (err != nil) != tt.wantErr {
				t.Errorf("SaveCommittedDataToFile() error = %v, wantErr %v", err, tt.wantErr)
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
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
