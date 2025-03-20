package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

func TestGetEpochLimitForUpdateCommission(t *testing.T) {
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
			name: "When GetEpochLimitForUpdateCommission() executes successfully",
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			stakeManagerMock.On("EpochLimitForUpdateCommission", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epochLimitForUpdateCommission, tt.args.epochLimitForUpdateCommissionErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetEpochLimitForUpdateCommission(rpcParameters)
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
	var (
		address  string
		stakerId uint32
		lockType uint8
	)

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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetStaker", mock.Anything, mock.Anything).Return(tt.args.staker, tt.args.stakerErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))
			stakeManagerMock.On("Locks", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.AnythingOfType("uint8")).Return(tt.args.locks, tt.args.locksErr)

			got, err := utils.GetLock(rpcParameters, address, stakerId, lockType)
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("MaxCommission", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.maxCommission, tt.args.maxCommissionErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMaxCommission(rpcParameters)
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetNumStakers", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numStakers, tt.args.numStakersErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumberOfStakers(rpcParameters)
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStake(rpcParameters, stakerId)
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStaker(rpcParameters, stakerId)
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStakerId(rpcParameters, account)
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
	type args struct {
		withdrawReleasePeriod    uint16
		withdrawReleasePeriodErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetWithdrawInitiationPeriod() executes successfully",
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
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("WithdrawInitiationPeriod", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.withdrawReleasePeriod, tt.args.withdrawReleasePeriodErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetWithdrawInitiationPeriod(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithdrawInitiationPeriod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWithdrawInitiationPeriod() got = %v, want %v", got, tt.want)
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

func TestGetMinSafeRazor(t *testing.T) {
	type args struct {
		minSafeRazor    *big.Int
		minSafeRazorErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetMinSafeRazor() executes successfully",
			args: args{
				minSafeRazor: big.NewInt(100),
			},
			want:    big.NewInt(100),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting minSafeRazor",
			args: args{
				minSafeRazorErr: errors.New("minSafeRazor error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("MinSafeRazor", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minSafeRazor, tt.args.minSafeRazorErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMinSafeRazor(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinSafeRazor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinSafeRazor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStakerInfo(t *testing.T) {
	type args struct {
		staker    types.Staker
		stakerErr error
	}
	tests := []struct {
		name    string
		args    args
		want    types.Staker
		wantErr bool
	}{
		{
			name: "Test 1: When StakerInfo() executes successfully",
			args: args{
				staker: types.Staker{
					Id: 1,
				},
			},
			want: types.Staker{
				Id: 1,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting staker info",
			args: args{
				stakerErr: errors.New("staker info error"),
			},
			want:    types.Staker{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("StakerInfo", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.StakerInfo(rpcParameters, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("StakerInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StakerInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaturity(t *testing.T) {
	type args struct {
		maturity    uint16
		maturityErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetMaturity() executes successfully",
			args: args{
				maturity: 100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting maturity",
			args: args{
				maturityErr: errors.New("maturity error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetMaturity", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.maturity, tt.args.maturityErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMaturity(rpcParameters, 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMaturity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMaturity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBountyLock(t *testing.T) {
	type args struct {
		bountyLock    types.BountyLock
		bountyLockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    types.BountyLock
		wantErr bool
	}{
		{
			name: "Test 1: When GetBountyLock() executes successfully",
			args: args{
				bountyLock: types.BountyLock{
					BountyHunter: common.HexToAddress("0xaA"),
				},
			},
			want: types.BountyLock{
				BountyHunter: common.HexToAddress("0xaA"),
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting bounty lock",
			args: args{
				bountyLockErr: errors.New("bounty lock error"),
			},
			want:    types.BountyLock{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			stakeManagerMock := new(mocks.StakeManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				StakeManagerInterface: stakeManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			stakeManagerMock.On("GetBountyLock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.bountyLock, tt.args.bountyLockErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetBountyLock(rpcParameters, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBountyLock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBountyLock() got = %v, want %v", got, tt.want)
			}
		})
	}
}
