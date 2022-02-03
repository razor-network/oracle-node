package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"math/big"
	"os"
	"razor/utils/mocks"
	"reflect"
	"testing"
	"time"
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

//func TestCalculateBlockTime(t *testing.T) {
//	//var client *ethclient.Client
//
//	type args struct {
//		client             *ethclient.Client
//		latestBlock        *types.Header
//		latestBlockErr     error
//		lastSecondBlock    *types.Header
//		lastSecondBlockErr error
//	}
//	tests := []struct {
//		name string
//		args args
//		want int64
//	}{
//		{
//			name: "Test 1: When CalculateBlockTime() executes successfully",
//			args: args{
//				client:          &ethclient.Client{},
//				latestBlock:     &types.Header{},
//				lastSecondBlock: &types.Header{},
//			},
//			want: 0,
//		},
//		//{
//		//	name: "Test 2: When there is an error in fetching latestBlock",
//		//	args: args{
//		//		latestBlockErr:  errors.New("error in fetching latestBlock"),
//		//		lastSecondBlock: &types.Header{},
//		//	},
//		//	expectedFatal: true,
//		//},
//		//{
//		//	name: "Test 3: When there is an error in fetching lastSecondBlock",
//		//	args: args{
//		//		latestBlock:        &types.Header{},
//		//		lastSecondBlockErr: errors.New("error in fetching lastSecondBlock"),
//		//	},
//		//	expectedFatal: false,
//		//},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			optionsMock := new(mocks.OptionUtils)
//			utilsMock := new(mocks.Utils)
//
//			optionsPackageStruct := OptionsPackageStruct{
//				Options:        optionsMock,
//				UtilsInterface: utilsMock,
//			}
//			utils := StartRazor(optionsPackageStruct)
//
//			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.latestBlock, tt.args.latestBlockErr)
//			optionsMock.On("HeaderByNumber", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.lastSecondBlock, tt.args.lastSecondBlockErr)
//
//			utils = &UtilsStruct{}
//
//			got := utils.CalculateBlockTime(tt.args.client)
//			if got != tt.want {
//				t.Errorf("CalculateBlockTime() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
			name: "Test 1: When CheckEthBalanceIsZero executes successfully",
			args: args{
				ethBalance: big.NewInt(1),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When CheckEthBalanceIsZero returns zero",
			args: args{
				ethBalance: big.NewInt(0),
			},
			expectedFatal: true,
		},
		//{
		//	name: "Test 3: When there is an error in fetching ethBalance",
		//	args: args{
		//		ethBalanceErr: errors.New("error in fetching ethBalance"),
		//	},
		//	expectedFatal: true,
		//},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options: optionsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("BalanceAt", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.ethBalance, tt.args.ethBalanceErr)

			utils = &UtilsStruct{}
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
			name: "Test 2: When there is error in getting tranactionReceipt",
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
			optionsMock := new(mocks.OptionUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options: optionsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("TransactionReceipt", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.tx, tt.args.txErr)

			utils = &UtilsStruct{}
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
			name: "When ConnectToClient() executes successfully",
			args: args{
				client: &ethclient.Client{},
			},
			expectedFatal: false,
		},
		{
			name: "When there is an error in ConnectToClient() function",
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
			optionsMock := new(mocks.OptionUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options: optionsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("Dial", mock.AnythingOfType("string")).Return(tt.args.client, tt.args.clientErr)

			utils = &UtilsStruct{}
			fatal = false

			utils.ConnectToClient(provider)
			if fatal != tt.expectedFatal {
				t.Error("The ConnectToClient function didn't execute as expected")
			}
		})
	}
}

//func TestFetchBalance(t *testing.T) {
//	var client *ethclient.Client
//	var accountAddress string
//	var callOpts bind.CallOpts
//
//	type args struct {
//		coinContract *bindings.RAZOR
//	}
//	tests := []struct {
//		name          string
//		args          args
//		expectedFatal bool
//	}{
//		{
//			name: "When FetchBalance() executes successfully",
//			args: args{
//				coinContract: &bindings.RAZOR{},
//			},
//			expectedFatal: true,
//		},
//	}
//
//	defer func() { log.ExitFunc = nil }()
//	var fatal bool
//	log.ExitFunc = func(int) { fatal = true }
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			utilsMock := new(mocks.Utils)
//
//			optionsPackageStruct := OptionsPackageStruct{
//				UtilsInterface: utilsMock,
//			}
//			utils := StartRazor(optionsPackageStruct)
//
//			utilsMock.On("GetTokenManager", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.coinContract)
//			utilsMock.On("GetOptions").Return(callOpts)
//
//			utils = &UtilsStruct{}
//			fatal = false
//
//			utils.FetchBalance(client, accountAddress)
//			if fatal != tt.expectedFatal {
//				t.Error("The ConnectToClient function didn't execute as expected")
//			}
//		})
//	}
//}

func TestGetDelayedState(t *testing.T) {

	type args struct {
		client *ethclient.Client
		buffer int32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, err := ut.GetDelayedState(tt.args.client, tt.args.buffer)
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

//func TestGetEpoch(t *testing.T) {
//	var client *ethclient.Client
//
//	type args struct {
//		latestHeader    *types.Header
//		latestHeaderErr error
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    uint32
//		wantErr bool
//	}{
//		{
//			name: "Test 1",
//			args: args{
//				latestHeader: &types.Header{},
//			},
//			want:    0,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			utilsMock := new(mocks.Utils)
//
//			optionsPackageStruct := OptionsPackageStruct{
//				UtilsInterface: utilsMock,
//			}
//			utils := StartRazor(optionsPackageStruct)
//
//			utilsMock.On("GetLatestBlockWithRetry", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.latestHeader, tt.args.latestHeaderErr)
//
//			got, err := utils.GetEpoch(client)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetEpoch() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetEpoch() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
			name: "Test 1",
			args: args{
				stateNumber: 0,
			},
			want: "Commit",
		},
		{
			name: "Test 2",
			args: args{
				stateNumber: 1,
			},
			want: "Reveal",
		},
		{
			name: "Test 3",
			args: args{
				stateNumber: 2,
			},
			want: "Propose",
		},
		{
			name: "Test 4",
			args: args{
				stateNumber: 3,
			},
			want: "Dispute",
		},
		{
			name: "Test 5",
			args: args{
				stateNumber: 4,
			},
			want: "Confirm",
		},
		{
			name: "Test 6",
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
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		want1   []*big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, got1, err := ut.ReadCommittedDataFromFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCommittedDataFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadCommittedDataFromFile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadCommittedDataFromFile() got1 = %v, want %v", got1, tt.want1)
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
			name: "Test 1",
			args: args{
				committedData: []*big.Int{big.NewInt(2)},
				file:          file,
			},
			wantErr: true,
		},
		{
			name: "Test 2",
			args: args{
				committedData: []*big.Int{big.NewInt(2)},
				fileErr:       errors.New("error in fetching file"),
			},
			wantErr: true,
		},
		{
			name: "Test 3",
			args: args{
				committedData: []*big.Int{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options: optionsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("OpenFile", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(tt.args.file, tt.args.fileErr)

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
			name: "test 1",
			args: args{
				transactionStatus: 0,
			},
			want: 0,
		},
		{
			name: "test 2",
			args: args{
				transactionStatus: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("CheckTransactionReceipt", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.transactionStatus)

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
			name: "test 1",
			args: args{
				name: "password",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFlagPassed(tt.args.name); got != tt.want {
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
			name: "Test 1",
			args: args{
				waitTime: 1,
			},
		},
		{
			name: "Test 2",
			args: args{
				waitTime: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitTillNextNSecs(tt.args.waitTime)
		})
	}
}

func TestSleep(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				duration: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sleep(tt.args.duration)
		})
	}
}
