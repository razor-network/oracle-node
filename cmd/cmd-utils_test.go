package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"testing"
)

func TestGetEpochAndState(t *testing.T) {
	var client *ethclient.Client
	var address string

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
	}

	type args struct {
		epoch            uint32
		epochErr         error
		bufferPercent    int32
		bufferPercentErr error
		state            int64
		stateErr         error
		stateName        string
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
				epoch:         4,
				bufferPercent: 20,
				state:         0,
				stateName:     "commit",
			},
			wantEpoch: 4,
			wantState: 0,
			wantErr:   nil,
		},
		{
			name: "Test 2: When there is an error in getting epoch",
			args: args{
				epochErr:      errors.New("epoch error"),
				bufferPercent: 20,
				state:         0,
				stateName:     "commit",
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("epoch error"),
		},
		{
			name: "Test 3: When there is an error in getting bufferPercent",
			args: args{
				epoch:            4,
				bufferPercentErr: errors.New("bufferPercent error"),
				state:            0,
				stateName:        "commit",
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("bufferPercent error"),
		},
		{
			name: "Test 4: When there is an error in getting state",
			args: args{
				epoch:         4,
				bufferPercent: 20,
				stateErr:      errors.New("state error"),
			},
			wantEpoch: 0,
			wantState: 0,
			wantErr:   errors.New("state error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			getBufferPercentMock = func(UtilsStruct) (int32, error) {
				return tt.args.bufferPercent, tt.args.bufferPercentErr
			}

			GetDelayedStateMock = func(*ethclient.Client, int32) (int64, error) {
				return tt.args.state, tt.args.stateErr
			}

			GetStateNameMock = func(int64) string {
				return tt.args.stateName
			}

			gotEpoch, gotState, err := GetEpochAndState(client, address, utilsStruct)
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
	var address string

	utilsStruct := UtilsStruct{
		cmdUtils:   UtilsCmdMock{},
		razorUtils: UtilsMock{},
	}

	type args struct {
		epoch           uint32
		state           int64
		epochOrStateErr error
		action          string
		states          int
		contains        bool
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
				epoch:    4,
				state:    1,
				action:   "reveal",
				states:   1,
				contains: true,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 2: When WaitForAppropriateState function executes successfully for commit state",
			args: args{
				epoch:    4,
				state:    0,
				action:   "commit",
				states:   0,
				contains: true,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 3: When WaitForAppropriateState function executes successfully for dispute state",
			args: args{
				epoch:    4,
				state:    3,
				action:   "dispute",
				states:   3,
				contains: true,
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
				contains:        true,
			},
			want:    0,
			wantErr: errors.New("error in fetching epoch and state"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetEpochAndStateMock = func(*ethclient.Client, string, UtilsStruct) (uint32, int64, error) {
				return tt.args.epoch, tt.args.state, tt.args.epochOrStateErr
			}

			ContainsMock = func([]int, int) bool {
				return tt.args.contains
			}

			got, err := WaitForAppropriateState(client, address, tt.args.action, utilsStruct, tt.args.states)
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
	var address string
	var action string

	utilsStruct := UtilsStruct{
		cmdUtils:   UtilsCmdMock{},
		razorUtils: UtilsMock{},
	}
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
			GetEpochAndStateMock = func(*ethclient.Client, string, UtilsStruct) (uint32, int64, error) {
				return tt.args.epoch, tt.args.state, tt.args.epochOrStateErr
			}

			got, err := WaitIfCommitState(client, address, action, utilsStruct)
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

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		flagSetUtils: FlagSetMock{},
	}

	type args struct {
		amount                   string
		amountErr                error
		_amount                  *big.Int
		_amountErr               bool
		isFlagPassed             bool
		power                    string
		powerErr                 error
		fractionalAmountInWei    *big.Int
		fractionalAmountInWeiErr error
		amountInWei              *big.Int
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
			name: "Test 2: When AssignAmountInWei executes successfully and power flag is passed",
			args: args{
				amount:                "10001",
				_amount:               big.NewInt(10001),
				isFlagPassed:          true,
				power:                 "17",
				fractionalAmountInWei: big.NewInt(1).Mul(big.NewInt(10001), big.NewInt(1e17)),
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
			name: "Test 5: When there is an error in getting power",
			args: args{
				amount:       "10001",
				_amount:      big.NewInt(10001),
				isFlagPassed: true,
				powerErr:     errors.New("power error"),
			},
			want:    nil,
			wantErr: errors.New("power error"),
		},
		{
			name: "Test 6: When there is an error in getting fractionalAmountInWei",
			args: args{
				amount:                   "10001",
				_amount:                  big.NewInt(10001),
				isFlagPassed:             true,
				power:                    "17",
				fractionalAmountInWeiErr: errors.New("fractionalAmountInWei error"),
			},
			want:    nil,
			wantErr: errors.New("fractionalAmountInWei error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStringValueMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.amount, tt.args.amountErr
			}

			IsFlagPassedMock = func(string) bool {
				return tt.args.isFlagPassed
			}

			GetStringPowMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.power, tt.args.powerErr
			}

			GetFractionalAmountInWeiMock = func(*big.Int, string) (*big.Int, error) {
				return tt.args.fractionalAmountInWei, tt.args.fractionalAmountInWeiErr
			}

			GetAmountInWeiMock = func(*big.Int) *big.Int {
				return tt.args.amountInWei
			}

			got, err := AssignAmountInWei(flagSet, utilsStruct)
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
