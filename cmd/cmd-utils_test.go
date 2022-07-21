package cmd

import (
	"errors"
	"math/big"
	"razor/cmd/mocks"
	"razor/utils"
	mocks2 "razor/utils/mocks"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestGetEpochAndState(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		epoch                  uint32
		epochErr               error
		bufferPercentString    string
		bufferPercentStringErr error
		bufferPercent          int64
		bufferPercentErr       error
		state                  int64
		stateErr               error
		stateName              string
	}
	tests := []struct {
		name      string
		args      args
		wantEpoch uint32
		wantState int64
		wantErr   error
	}{
		{
			name: "Test 1: When GetEpochAndState function executes successfully",
			args: args{
				epoch:               4,
				bufferPercentString: "20",
				bufferPercent:       20,
				state:               0,
				stateName:           "commit",
			},
			wantEpoch: 4,
			wantState: 0,
			wantErr:   nil,
		},
		{
			name: "Test 2: When there is an error in getting epoch",
			args: args{
				epochErr:            errors.New("epoch error"),
				bufferPercentString: "20",
				bufferPercent:       20,
				state:               0,
				stateName:           "commit",
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("epoch error"),
		},
		{
			name: "Test 3: When there is an error in getting getConfig",
			args: args{
				epoch:                  4,
				bufferPercentStringErr: errors.New("bufferPercentString error"),
				state:                  0,
				stateName:              "commit",
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("bufferPercentString error"),
		},
		{
			name: "Test 4: When there is an error in parsing int",
			args: args{
				epoch:               4,
				bufferPercentString: "20",
				bufferPercentErr:    errors.New("bufferPercent error"),
				state:               0,
				stateName:           "commit",
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("bufferPercent error"),
		},
		{
			name: "Test 5: When there is an error in getting state",
			args: args{
				epoch:               4,
				bufferPercentString: "20",
				bufferPercent:       20,
				stateErr:            errors.New("state error"),
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("state error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(mocks2.Utils)
			stringMock := new(mocks.StringInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			utils.UtilsInterface = utilsPkgMock
			stringUtils = stringMock

			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			cmdUtilsMock.On("GetConfig", "buffer").Return(tt.args.bufferPercentString, tt.args.bufferPercentStringErr)
			stringMock.On("ParseInt", tt.args.bufferPercentString).Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			utilsMock.On("GetBufferedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
			utilsPkgMock.On("GetStateName", mock.AnythingOfType("int64")).Return(tt.args.stateName)

			utils := &UtilsStruct{}
			gotEpoch, gotState, err := utils.GetEpochAndState(client)
			if gotEpoch != tt.wantEpoch {
				t.Errorf("GetEpochAndState() got epoch = %v, want %v", gotEpoch, tt.wantEpoch)
			}
			if gotState != tt.wantState {
				t.Errorf("GetEpochAndState() got1 state = %v, want %v", gotState, tt.wantState)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetEpochAndState function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetEpochAndState function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestWaitForAppropriateState(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		epoch           uint32
		state           int64
		epochOrStateErr error
		action          string
		states          int
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr error
	}{
		{
			name: "Test 1: When WaitForAppropriateState function executes successfully for reveal state",
			args: args{
				epoch:  4,
				state:  1,
				action: "reveal",
				states: 1,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 2: When WaitForAppropriateState function executes successfully for commit state",
			args: args{
				epoch:  4,
				state:  0,
				action: "commit",
				states: 0,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 3: When WaitForAppropriateState function executes successfully for dispute state",
			args: args{
				epoch:  4,
				state:  3,
				action: "dispute",
				states: 3,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 4: When there is an error in getting epoch or state",
			args: args{
				epochOrStateErr: errors.New("error in fetching epoch and state"),
				action:          "commit",
				states:          0,
			},
			want:    0,
			wantErr: errors.New("error in fetching epoch and state"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			timeMock := new(mocks.TimeInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			timeUtils = timeMock
			cmdUtils = cmdUtilsMock

			cmdUtilsMock.On("GetEpochAndState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.epoch, tt.args.state, tt.args.epochOrStateErr)
			timeMock.On("Sleep", mock.Anything).Return()
			utils := &UtilsStruct{}
			got, err := utils.WaitForAppropriateState(client, tt.args.action, tt.args.states)
			if got != tt.want {
				t.Errorf("WaitForAppropriateState() function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for WaitForAppropriateState function, got = %v, wantErr = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for WaitForAppropriateState function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestWaitIfCommitState(t *testing.T) {
	var client *ethclient.Client
	var action string

	type args struct {
		epoch           uint32
		state           int64
		epochOrStateErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr error
	}{
		{
			name: "Test 1: When WaitIfCommitState function execute successffuly",
			args: args{
				epoch: 5,
				state: 2,
			},
			want:    5,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting epoch and state",
			args: args{
				epochOrStateErr: errors.New("error in fetching epoch and state"),
			},
			want:    0,
			wantErr: errors.New("error in fetching epoch and state"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			timeMock := new(mocks.TimeInterface)

			cmdUtils = cmdUtilsMock
			timeUtils = timeMock

			cmdUtilsMock.On("GetEpochAndState", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.state, tt.args.epochOrStateErr)
			timeMock.On("Sleep", mock.Anything).Return()

			utils := &UtilsStruct{}

			got, err := utils.WaitIfCommitState(client, action)
			if got != tt.want {
				t.Errorf("WaitIfCommitState() function, got = %v, want = %v", got, tt.want)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for WaitIfCommitState function, got = %v, wantErr = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for WaitIfCommitState function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestAssignAmountInWei1(t *testing.T) {
	var flagSet *pflag.FlagSet

	type args struct {
		amount       string
		amountErr    error
		_amount      *big.Int
		_amountErr   bool
		isFlagPassed bool
		weiRazor     bool
		weiRazorErr  error
		amountInWei  *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr error
	}{
		{
			name: "Test 1: When AssignAmountInWei executes successfully",
			args: args{
				amount:       "1000",
				_amount:      big.NewInt(1000),
				isFlagPassed: false,
				amountInWei:  big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want:    big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			wantErr: nil,
		},
		{
			name: "Test 2: When AssignAmountInWei executes successfully and weiRazor flag is passed",
			args: args{
				amount:       "1000100000000000000000",
				_amount:      big.NewInt(1).Mul(big.NewInt(10001), big.NewInt(1e17)),
				isFlagPassed: true,
				weiRazor:     true,
			},
			want:    big.NewInt(1).Mul(big.NewInt(10001), big.NewInt(1e17)),
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting amount",
			args: args{
				amountErr:    errors.New("amount error"),
				isFlagPassed: false,
			},
			want:    nil,
			wantErr: errors.New("amount error"),
		},
		{
			name: "Test 4: When there is a setString error in converting string amount",
			args: args{
				amount:       "1000A",
				_amountErr:   true,
				isFlagPassed: false,
			},
			want:    nil,
			wantErr: errors.New("SetString: error"),
		},
		{
			name: "Test 5: When there is an error in getting if weiRazor is passed or not",
			args: args{
				amount:       "10001",
				_amount:      big.NewInt(10001),
				isFlagPassed: true,
				weiRazorErr:  errors.New("weiRazor error"),
			},
			want:    nil,
			wantErr: errors.New("weiRazor error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			flagsetUtilsMock := new(mocks.FlagSetInterface)

			razorUtils = utilsMock
			flagSetUtils = flagsetUtilsMock

			flagsetUtilsMock.On("GetStringValue", flagSet).Return(tt.args.amount, tt.args.amountErr)
			flagsetUtilsMock.On("GetBoolWeiRazor", flagSet).Return(tt.args.weiRazor, tt.args.weiRazorErr)
			utilsMock.On("IsFlagPassed", mock.AnythingOfType("string")).Return(tt.args.isFlagPassed)
			utilsMock.On("GetAmountInWei", mock.AnythingOfType("*big.Int")).Return(tt.args.amountInWei)

			utils := &UtilsStruct{}
			got, err := utils.AssignAmountInWei(flagSet)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("AssignAmountInWei() function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for AssignAmountInWei function, got = %v, wantErr = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for AssignAmountInWei function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetStatesAllowed(t *testing.T) {
	type args struct {
		states    []int
		stateName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1: When states has single elements",
			args: args{
				states:    []int{1},
				stateName: "Reveal",
			},
			want: "1:Reveal",
		},
		{
			name: "Test 2: When states has multiple elements",
			args: args{
				states:    []int{1, 1},
				stateName: "Reveal",
			},
			want: "1:Reveal, 1:Reveal",
		},
		{
			name: "Test 2: When states array is nil",
			args: args{
				states: []int{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsPkgMock := new(mocks2.Utils)

			utils.UtilsInterface = utilsPkgMock

			utilsPkgMock.On("GetStateName", mock.AnythingOfType("int64")).Return(tt.args.stateName)
			if got := GetStatesAllowed(tt.args.states); got != tt.want {
				t.Errorf("GetStatesAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
