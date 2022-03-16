package utils

import (
	"errors"
	"math/big"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

func TestFetchPreviousValue(t *testing.T) {
	var client *ethclient.Client
	var epoch uint32

	type args struct {
		assetId  uint16
		block    bindings.StructsBlock
		blockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When FetchPreviousValue() executes successfully",
			args: args{
				assetId: 3,
				block: bindings.StructsBlock{
					Medians: []uint32{2000, 1500, 4000, 6500},
				},
			},
			want:    4000,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting proposed block",
			args: args{
				assetId:  3,
				blockErr: errors.New("block error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			blockManagerMock.On("GetBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.block, tt.args.blockErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.FetchPreviousValue(client, epoch, tt.args.assetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchPreviousValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchPreviousValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxAltBlocks(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		maxAltBlocks    uint8
		maxAltBlocksErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "Test 1: When GetMaxAltBlocks() executes successfully",
			args: args{
				maxAltBlocks: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting maxAltBlocks",
			args: args{
				maxAltBlocksErr: errors.New("maxAltBlocks error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			blockManagerMock.On("MaxAltBlocks", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.maxAltBlocks, tt.args.maxAltBlocksErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMaxAltBlocks(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxAltBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxAltBlocks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinStakeAmount(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts

	type args struct {
		minStake    *big.Int
		minStakeErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetMinStakeAmount() executes successfully",
			args: args{
				minStake: big.NewInt(1000),
			},
			want:    big.NewInt(1000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting minStake",
			args: args{
				minStakeErr: errors.New("minStake error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			blockManagerMock.On("MinStake", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minStake, tt.args.minStakeErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetMinStakeAmount(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinStakeAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinStakeAmount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberOfProposedBlocks(t *testing.T) {
	var client *ethclient.Client
	var callOpts bind.CallOpts
	var epoch uint32

	type args struct {
		numOfProposedBlocks    uint8
		numOfProposedBlocksErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "Test 1: When GetNumberOfProposedBlocks() executes successfully",
			args: args{
				numOfProposedBlocks: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numOfProposedBlocks",
			args: args{
				numOfProposedBlocksErr: errors.New("numOfProposedBlocks error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			blockManagerMock.On("GetNumProposedBlocks", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetNumberOfProposedBlocks(client, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfProposedBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumberOfProposedBlocks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProposedBlock(t *testing.T) {
	var client *ethclient.Client
	var epoch uint32
	var proposedBlockId uint32
	var callOpts bind.CallOpts

	block := bindings.StructsBlock{
		Medians: []uint32{2000, 1500, 4000, 6500},
	}

	type args struct {
		block    bindings.StructsBlock
		blockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsBlock
		wantErr bool
	}{
		{
			name: "Test 1: When GetProposedBlock() executes successfully",
			args: args{
				block: block,
			},
			want:    block,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting proposed block",
			args: args{
				blockErr: errors.New("proposedBlock error"),
			},
			want:    bindings.StructsBlock{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			blockManagerMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.block, tt.args.blockErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetProposedBlock(client, epoch, proposedBlockId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProposedBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProposedBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSortedProposedBlockId(t *testing.T) {
	var client *ethclient.Client
	var epoch uint32
	var index *big.Int
	var callOpts bind.CallOpts

	type args struct {
		sortedProposedBlockId    uint32
		sortedProposedBlockIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetSortedProposedBlockId() executes successfully",
			args: args{
				sortedProposedBlockId: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting sorted proposed block id",
			args: args{
				sortedProposedBlockIdErr: errors.New("sortedProposedBlockId error"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			utilsMock := new(mocks.Utils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				UtilsInterface:        utilsMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			utilsMock.On("GetOptions").Return(callOpts)
			blockManagerMock.On("SortedProposedBlockIds", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("*big.Int")).Return(tt.args.sortedProposedBlockId, tt.args.sortedProposedBlockIdErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetSortedProposedBlockId(client, epoch, index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSortedProposedBlockId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortedProposedBlockId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSortedProposedBlockIds(t *testing.T) {
	var client *ethclient.Client
	var epoch uint32

	type args struct {
		numOfProposedBlocks      uint8
		numOfProposedBlocksErr   error
		sortedProposedBlockId    uint32
		sortedProposedBlockIdErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetSortedProposedBlockIds() executes successfully",
			args: args{
				numOfProposedBlocks:   1,
				sortedProposedBlockId: 2,
			},
			want:    []uint32{2},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting numOfProposedBlocks",
			args: args{
				numOfProposedBlocksErr: errors.New("numOfProposedBlocks error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting sorted proposed block id",
			args: args{
				numOfProposedBlocks:      1,
				sortedProposedBlockIdErr: errors.New("sortedProposedBlockId error"),
			},
			want:    nil,
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

			utilsMock.On("GetNumberOfProposedBlocks", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr)
			utilsMock.On("GetSortedProposedBlockId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("*big.Int")).Return(tt.args.sortedProposedBlockId, tt.args.sortedProposedBlockIdErr)

			got, err := utils.GetSortedProposedBlockIds(client, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSortedProposedBlockIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortedProposedBlockIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBlockManagerWithOpts(t *testing.T) {

	var callOpts bind.CallOpts
	var blockManager *bindings.BlockManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetBlockManager", mock.AnythingOfType("*ethclient.Client")).Return(blockManager)

	gotBlockManager, gotCallOpts := utils.GetBlockManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetBlockManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotBlockManager, blockManager) {
		t.Errorf("GetBlockManagerWithOpts() got blockManager = %v, want %v", gotBlockManager, blockManager)
	}
}
