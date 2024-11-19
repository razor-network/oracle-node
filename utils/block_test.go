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

func TestFetchPreviousValue(t *testing.T) {
	var epoch uint32

	type args struct {
		assetId  uint16
		block    bindings.StructsBlock
		blockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When FetchPreviousValue() executes successfully",
			args: args{
				assetId: 3,
				block: bindings.StructsBlock{
					Medians: []*big.Int{big.NewInt(2000), big.NewInt(1500), big.NewInt(4000), big.NewInt(6500)},
				},
			},
			want:    big.NewInt(4000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting proposed block",
			args: args{
				assetId:  3,
				blockErr: errors.New("block error"),
			},
			want:    big.NewInt(0),
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

			utilsMock.On("GetBlock", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.block, tt.args.blockErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.FetchPreviousValue(rpcParameters, epoch, tt.args.assetId)
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

			got, err := utils.GetMaxAltBlocks(rpcParameters)
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

			got, err := utils.GetMinStakeAmount(rpcParameters)
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

			got, err := utils.GetNumberOfProposedBlocks(rpcParameters, epoch)
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
	var epoch uint32
	var proposedBlockId uint32
	var callOpts bind.CallOpts

	block := bindings.StructsBlock{
		Medians: []*big.Int{big.NewInt(2000), big.NewInt(1500), big.NewInt(4000), big.NewInt(6500)},
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

			got, err := utils.GetProposedBlock(rpcParameters, epoch, proposedBlockId)
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

			got, err := utils.GetSortedProposedBlockId(rpcParameters, epoch, index)
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

			utilsMock.On("GetNumberOfProposedBlocks", mock.Anything, mock.Anything).Return(tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr)
			utilsMock.On("GetSortedProposedBlockId", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.sortedProposedBlockId, tt.args.sortedProposedBlockIdErr)

			got, err := utils.GetSortedProposedBlockIds(rpcParameters, epoch)
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

func TestGetBlock(t *testing.T) {
	var epoch uint32
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
			name: "Test 1: When GetBlock() executes successfully",
			args: args{
				block: bindings.StructsBlock{},
			},
			want:    bindings.StructsBlock{},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting block",
			args: args{
				blockErr: errors.New("error in getting block"),
			},
			want:    bindings.StructsBlock{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			blockManagerMock.On("GetBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.block, tt.args.blockErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetBlock(rpcParameters, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBlockIndexToBeConfirmed(t *testing.T) {
	type args struct {
		blockIndex    int8
		blockIndexErr error
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		{
			name: "Test 1: When GetBlockIndexToBeConfirmed() executes successfully",
			args: args{
				blockIndex: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting blockIndex",
			args: args{
				blockIndexErr: errors.New("error in getting blockIndex"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			blockManagerMock.On("GetBlockIndexToBeConfirmed", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.blockIndex, tt.args.blockIndexErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetBlockIndexToBeConfirmed(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockIndexToBeConfirmed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBlockIndexToBeConfirmed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStateBuffer(t *testing.T) {
	type args struct {
		stateBuffer    uint8
		stateBufferErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "When GetStateBuffer() executes successfully",
			args: args{
				stateBuffer: 5,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "When there is an error in getting stateBuffer",
			args: args{
				stateBufferErr: errors.New("error in getting stateBuffer"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			blockManagerMock := new(mocks.BlockManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:        retryMock,
				BlockManagerInterface: blockManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			blockManagerMock.On("StateBuffer", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.stateBuffer, tt.args.stateBufferErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStateBuffer(rpcParameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStateBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStateBuffer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpochLastProposed(t *testing.T) {
	var callOpts bind.CallOpts
	var stakerId uint32

	type args struct {
		epochLastProposed    uint32
		epochLastProposedErr error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetEpochLastProposed() executes successfully",
			args: args{
				epochLastProposed: 100,
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting epochLastProposed",
			args: args{
				epochLastProposedErr: errors.New("epochLastProposed error"),
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
			blockManagerMock.On("GetEpochLastProposed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastProposed, tt.args.epochLastProposedErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetEpochLastProposed(rpcParameters, stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLastProposed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLastProposed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfirmedBlocks(t *testing.T) {
	var callOpts bind.CallOpts
	var epoch uint32

	type args struct {
		confirmedBlock    types.ConfirmedBlock
		confirmedBlockErr error
	}
	tests := []struct {
		name    string
		args    args
		want    types.ConfirmedBlock
		wantErr bool
	}{
		{
			name: "Test 1: When GetConfirmedBlocks() executes successfully",
			args: args{
				confirmedBlock: types.ConfirmedBlock{
					ProposerId: 1,
				},
			},
			want: types.ConfirmedBlock{
				ProposerId: 1,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting confirmedBlock",
			args: args{
				confirmedBlockErr: errors.New("confirmedBlock error"),
			},
			want:    types.ConfirmedBlock{},
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
			blockManagerMock.On("GetConfirmedBlocks", mock.Anything, mock.Anything).Return(tt.args.confirmedBlock, tt.args.confirmedBlockErr)
			retryMock.On("RetryAttempts", mock.Anything).Return(retry.Attempts(1))

			got, err := utils.GetConfirmedBlocks(rpcParameters, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfirmedBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetConfirmedBlocks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisputes(t *testing.T) {
	var callOpts bind.CallOpts
	var epoch uint32
	var address common.Address

	type args struct {
		disputes    types.DisputesStruct
		disputesErr error
	}
	tests := []struct {
		name    string
		args    args
		want    types.DisputesStruct
		wantErr bool
	}{
		{
			name: "Test 1: When Disputes() executes successfully",
			args: args{
				disputes: types.DisputesStruct{
					LeafId: 1,
				},
			},
			want: types.DisputesStruct{
				LeafId: 1,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting disputes",
			args: args{
				disputesErr: errors.New("disputes error"),
			},
			want:    types.DisputesStruct{},
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
			blockManagerMock.On("Disputes", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputes, tt.args.disputesErr)
			retryMock.On("RetryAttempts", mock.Anything).Return(retry.Attempts(1))

			got, err := utils.Disputes(rpcParameters, epoch, address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Disputes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Disputes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
