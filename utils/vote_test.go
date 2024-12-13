package utils

import (
	"errors"
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

func TestGetCommitments(t *testing.T) {
	var callOpts bind.CallOpts
	var address string

	type args struct {
		stakerId      uint32
		stakerIdErr   error
		commitments   types.Commitment
		commitmentErr error
	}
	tests := []struct {
		name    string
		args    args
		want    types.Commitment
		wantErr bool
	}{
		{
			name: "Test 1: When GetCommitment() executes successfully",
			args: args{
				stakerId:    1,
				commitments: types.Commitment{},
			},
			want:    types.Commitment{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting stakerId",
			args: args{
				stakerIdErr: errors.New("stakerId error"),
				commitments: types.Commitment{},
			},
			want:    types.Commitment{},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting commitments",
			args: args{
				stakerId:      1,
				commitmentErr: errors.New("commitments error"),
			},
			want:    types.Commitment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			utilsMock.On("GetStakerId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakerId, tt.args.stakerIdErr)
			voteManagerMock.On("GetCommitment", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.commitments, tt.args.commitmentErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetCommitment(rpcParameters, address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommitment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpochLastCommitted(t *testing.T) {
	var callOpts bind.CallOpts
	var stakerId uint32

	type args struct {
		epochLastCommitted    uint32
		epochLastCommittedErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetEpochLastCommitted() executes successfully",
			args: args{
				epochLastCommitted: 100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting epochLastCommitted",
			args: args{
				epochLastCommittedErr: errors.New("epochLastCommitted error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			voteManagerMock.On("GetEpochLastCommitted", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastCommitted, tt.args.epochLastCommittedErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetEpochLastCommitted(rpcParameters, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLastCommitted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLastCommitted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpochLastRevealed(t *testing.T) {
	var stakerId uint32

	type args struct {
		epochLastRevealed    uint32
		epochLastRevealedErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetEpochLastRevealed() executes successfully",
			args: args{
				epochLastRevealed: 100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting epochLastRevealed",
			args: args{
				epochLastRevealedErr: errors.New("epochLastRevealed error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			voteManagerMock.On("GetEpochLastRevealed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastRevealed, tt.args.epochLastRevealedErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetEpochLastRevealed(rpcParameters, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLastRevealed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLastRevealed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInfluenceSnapshot(t *testing.T) {
	var callOpts bind.CallOpts
	var stakerId uint32
	var epoch uint32

	type args struct {
		influenceSnapshot *big.Int
		influenceErr      error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetInfluenceSnapshot() executes successfully",
			args: args{
				influenceSnapshot: big.NewInt(1),
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting influence",
			args: args{
				influenceErr: errors.New("influence error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			voteManagerMock.On("GetInfluenceSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.influenceSnapshot, tt.args.influenceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetInfluenceSnapshot(rpcParameters, stakerId, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfluenceSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfluenceSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStakeSnapshot(t *testing.T) {
	var callOpts bind.CallOpts
	var stakerId uint32
	var epoch uint32

	type args struct {
		stakeSnapshot *big.Int
		snapshotErr   error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetStakeSnapshot() executes successfully",
			args: args{
				stakeSnapshot: big.NewInt(10000),
			},
			want:    big.NewInt(10000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting snapshot",
			args: args{
				snapshotErr: errors.New("snapshot error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			voteManagerMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.stakeSnapshot, tt.args.snapshotErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStakeSnapshot(rpcParameters, stakerId, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStakeSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStakeSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTotalInfluenceRevealed(t *testing.T) {
	var (
		callOpts    bind.CallOpts
		epoch       uint32
		medianIndex uint16
	)

	type args struct {
		totalInfluenceRevealed *big.Int
		influenceErr           error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetTotalInfluenceRevealed() executes successfully",
			args: args{
				totalInfluenceRevealed: big.NewInt(100000),
			},
			want:    big.NewInt(100000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting influence",
			args: args{
				influenceErr: errors.New("influence error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			voteManagerMock.On("GetTotalInfluenceRevealed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint16")).Return(tt.args.totalInfluenceRevealed, tt.args.influenceErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetTotalInfluenceRevealed(rpcParameters, epoch, medianIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTotalInfluenceRevealed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalInfluenceRevealed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVoteValue(t *testing.T) {
	var (
		callOpts    bind.CallOpts
		medianIndex uint16
		stakerId    uint32
		epoch       uint32
	)

	type args struct {
		voteValue    *big.Int
		voteValueErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetVoteValue() executes successfully",
			args: args{
				voteValue: big.NewInt(50000),
			},
			want:    big.NewInt(50000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting voteValue",
			args: args{
				voteValueErr: errors.New("voteValue error"),
			},
			want:    big.NewInt(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				UtilsInterface:       utilsMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			voteManagerMock.On("GetVoteValue", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint16")).Return(tt.args.voteValue, tt.args.voteValueErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetVoteValue(rpcParameters, epoch, stakerId, medianIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoteValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("GetVoteValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVoteManagerWithOpts(t *testing.T) {
	var callOpts bind.CallOpts
	var voteManager *bindings.VoteManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetVoteManager", mock.AnythingOfType("*ethclient.Client")).Return(voteManager)

	gotVoteManager, gotCallOpts := utils.GetVoteManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetVoteManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotVoteManager, voteManager) {
		t.Errorf("GetVoteManagerWithOpts() got voteManager = %v, want %v", gotVoteManager, voteManager)
	}
}

func TestToAssign(t *testing.T) {
	type args struct {
		toAssign    uint16
		toAssignErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "Test 1: When ToAssign() executes successfully",
			args: args{
				toAssign: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting toAssign",
			args: args{
				toAssignErr: errors.New("error in toAssign"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			voteManagerMock := new(mocks.VoteManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				VoteManagerInterface: voteManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			voteManagerMock.On("ToAssign", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.toAssign, tt.args.toAssignErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.ToAssign(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAssign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToAssign() got = %v, want %v", got, tt.want)
			}
		})
	}
}
