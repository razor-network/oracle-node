package utils

import (
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestGetEpochLimitForUpdateCommission(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		epochLimitForUpdateCommission    uint16
		epochLimitForUpdateCommissionErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "When GetEpochLimitForUpdateCommission() exectues successfully",
			args: args{
				epochLimitForUpdateCommission: 100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "When there is an error in getting epochLimitForUpdateCommission",
			args: args{
				epochLimitForUpdateCommissionErr: errors.New("epochLimitForUpdateCommission error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("EpochLimitForUpdateCommission", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.epochLimitForUpdateCommission, tt.args.epochLimitForUpdateCommissionErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetEpochLimitForUpdateCommission(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLimitForUpdateCommission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLimitForUpdateCommission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLock(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var address string
	var stakerId uint32

	type args struct {
		staker    bindings.StructsStaker
		stakerErr error
		locks     types.Locks
		locksErr  error
	}
	tests := []struct {
		name    string
		args    args
		want    types.Locks
		wantErr bool
	}{
		{
			name: "Test 1: When GetLock() executes successfully",
			args: args{
				staker: bindings.StructsStaker{},
				locks:  types.Locks{},
			},
			want:    types.Locks{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting staker",
			args: args{
				stakerErr: errors.New("staker error"),
			},
			want:    types.Locks{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an getting locks",
			args: args{
				staker:   bindings.StructsStaker{},
				locksErr: errors.New("locks error"),
			},
			want:    types.Locks{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))
			stakeManagerMock.On("Locks", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.Anything, mock.Anything).Return(tt.args.locks, tt.args.locksErr)

			got, err := utils.GetLock(client, address, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxCommission(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		maxCommission    uint8
		maxCommissionErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "When GetMaxCommission() executes successfully",
			args: args{
				maxCommission: 20,
			},
			want:    20,
			wantErr: false,
		},
		{
			name: "When there is an error in getting maxCommission",
			args: args{
				maxCommissionErr: errors.New("maxCommission error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("MaxCommission", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.maxCommission, tt.args.maxCommissionErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMaxCommission(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMaxCommission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMaxCommission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberOfStakers(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		numStakers    uint32
		numStakersErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "When GetNumberOfStakers() executes successfully",
			args: args{
				numStakers: 1000,
			},
			want:    1000,
			wantErr: false,
		},
		{
			name: "When there is an error in getting numStakers",
			args: args{
				numStakersErr: errors.New("numStakers error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("GetNumStakers", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.numStakers, tt.args.numStakersErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumberOfStakers(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfStakers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfStakers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStake(t *testing.T) {
	var client *ethclient.Client
	var stakerId uint32

	type args struct {
		staker    bindings.StructsStaker
		stakerErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "When GetStake() exectues successfully",
			args: args{
				staker: bindings.StructsStaker{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When there is an error in getting staker",
			args: args{
				stakerErr: errors.New("staker error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStake(client, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStaker(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var stakerId uint32

	type args struct {
		staker    bindings.StructsStaker
		stakerErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsStaker
		wantErr bool
	}{
		{
			name: "When GetStaker() exectues successfully",
			args: args{
				staker: bindings.StructsStaker{},
			},
			want:    bindings.StructsStaker{},
			wantErr: false,
		},
		{
			name: "When there is an error in getting staker",
			args: args{
				stakerErr: errors.New("staker error"),
			},
			want:    bindings.StructsStaker{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStaker(client, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStakerId(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var account string

	type args struct {
		stakerId    uint32
		stakerIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "When GetStakerId() exectues successfully",
			args: args{
				stakerId: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "When there is an error in getting stakerId",
			args: args{
				stakerIdErr: errors.New("stakerId error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), &callOpts, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStakerId(client, account)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStakerId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStakerId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWithdrawReleasePeriod(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		withdrawReleasePeriod    uint8
		withdrawReleasePeriodErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "Test 1: When GetWithdrawReleasePeriod() executes successfully",
			args: args{
				withdrawReleasePeriod: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting withdrawReleasePeriod",
			args: args{
				withdrawReleasePeriodErr: errors.New("withdrawReleasePeriood error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:               optionsMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("WithdrawReleasePeriod", mock.AnythingOfType("*ethclient.Client"), &callOpts).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetWithdrawReleasePeriod(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithdrawReleasePeriod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWithdrawReleasePeriod() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStakeManagerWithOpts(t *testing.T) {
	var callOpts bind.CallOpts
	var stakeManager *bindings.StakeManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetStakeManager", mock.AnythingOfType("*ethclient.Client")).Return(stakeManager)

	gotStakeManager, gotCallOpts := utils.GetStakeManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetStakeManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotStakeManager, stakeManager) {
		t.Errorf("GetStakeManagerWithOpts() got stakeManager = %v, want %v", gotStakeManager, stakeManager)
	}
}
