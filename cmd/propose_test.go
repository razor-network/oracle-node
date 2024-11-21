package cmd

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
)

func TestPropose(t *testing.T) {

	var (
		account     types.Account
		config      types.Configurations
		staker      bindings.StructsStaker
		epoch       uint32
		stateBuffer uint64
	)

	salt := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	saltBytes32 := [32]byte{}
	copy(saltBytes32[:], salt)

	latestHeader := &Types.Header{
		Number: big.NewInt(1001),
	}
	type args struct {
		rogueData                  types.Rogue
		state                      int64
		stateErr                   error
		staker                     bindings.StructsStaker
		numStakers                 uint32
		numStakerErr               error
		biggestStake               *big.Int
		biggestStakerId            uint32
		biggestStakerIdErr         error
		smallestStake              *big.Int
		smallestStakerId           uint32
		smallestStakerIdErr        error
		randaoHash                 [32]byte
		randaoHashErr              error
		salt                       [32]byte
		saltErr                    error
		iteration                  int
		numOfProposedBlocks        uint8
		numOfProposedBlocksErr     error
		sortedProposedBlockIds     []uint32
		sortedProposedBlocksIdsErr error
		maxAltBlocks               uint8
		maxAltBlocksErr            error
		lastIteration              *big.Int
		lastProposedBlockStruct    bindings.StructsBlock
		lastProposedBlockStructErr error
		medians                    []*big.Int
		ids                        []uint16
		revealDataMaps             *types.RevealedDataMaps
		mediansErr                 error
		fileName                   string
		fileNameErr                error
		saveDataErr                error
		mediansBigInt              []*big.Int
		txnOptsErr                 error
		proposeTxn                 *Types.Transaction
		proposeErr                 error
		hash                       common.Hash
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When Propose function executes successfully",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 1, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr:                errors.New("state error"),
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 1, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When there is an error in getting number of stakers",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakerErr:            errors.New("numberOfStakers error"),
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				sortedProposedBlockIds:  []uint32{1, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 4: When there is an error in getting biggest staker",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStakerIdErr:      errors.New("biggest staker error"),
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 1, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("biggest staker error"),
		},
		{
			name: "Test 5: When there is an error in getting randaoHash",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				saltErr:                 errors.New("salt error"),
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("salt error"),
		},
		{
			name: "Test 6: When iteration is -1",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               -1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 7: When there is an error in getting number of proposed blocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocksErr:  errors.New("numOfProposedBlocks error"),
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("numOfProposedBlocks error"),
		},
		{
			name: "Test 8: When there is an error in getting maxAltBlocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				sortedProposedBlockIds:  []uint32{0, 1},
				maxAltBlocksErr:         errors.New("maxAltBlocks error"),
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("maxAltBlocks error"),
		},
		{
			name: "Test 9: When numOfProposedBlocks >= maxAltBlocks and there is an error in getting lastProposedBlockStruct",
			args: args{
				state:                      2,
				staker:                     bindings.StructsStaker{},
				numStakers:                 5,
				biggestStake:               big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:            2,
				salt:                       saltBytes32,
				iteration:                  1,
				numOfProposedBlocks:        4,
				sortedProposedBlockIds:     []uint32{2, 0, 1, 3},
				maxAltBlocks:               2,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStructErr: errors.New("lastProposedBlockStruct error"),
				medians:                    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:                 &Types.Transaction{},
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("lastProposedBlockStruct error"),
		},
		{
			name: "Test 10: When numOfProposedBlocks >= maxAltBlocks and current iteration is greater than iteration of last proposed block ",
			args: args{
				state:                  2,
				staker:                 bindings.StructsStaker{},
				numStakers:             5,
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:        2,
				salt:                   saltBytes32,
				iteration:              2,
				numOfProposedBlocks:    4,
				sortedProposedBlockIds: []uint32{2, 0, 1, 3},
				maxAltBlocks:           2,
				lastIteration:          big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(1),
				},
				medians:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 11: When numOfProposedBlocks >= maxAltBlocks and current iteration is less than iteration of last proposed block and propose transaction is successful",
			args: args{
				state:                  2,
				staker:                 bindings.StructsStaker{},
				numStakers:             5,
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:        2,
				salt:                   saltBytes32,
				iteration:              1,
				numOfProposedBlocks:    4,
				sortedProposedBlockIds: []uint32{2, 0, 1, 3},
				maxAltBlocks:           2,
				lastIteration:          big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(2),
				},
				medians:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 12: When numOfProposedBlocks >= maxAltBlocks and there is an error in fetching sortedProposedBlockIds",
			args: args{
				state:                      2,
				staker:                     bindings.StructsStaker{},
				numStakers:                 5,
				biggestStake:               big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:            2,
				salt:                       saltBytes32,
				iteration:                  1,
				numOfProposedBlocks:        4,
				sortedProposedBlockIds:     nil,
				sortedProposedBlocksIdsErr: errors.New("error in fetching sorted proposed block ids"),
				maxAltBlocks:               2,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(2),
				},
				medians:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("error in fetching sorted proposed block ids"),
		},
		{
			name: "Test 13: When there is an error in getting medians",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				sortedProposedBlockIds:  []uint32{2, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				mediansErr:              errors.New("makeBlock error"),
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("makeBlock error"),
		},
		{
			name: "Test 14: When Propose transaction fails",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				sortedProposedBlockIds:  []uint32{2, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeErr:              errors.New("propose error"),
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("propose error"),
		},
		{
			name: "Test 15: When there is an error in getting fileName",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
				fileNameErr:             errors.New("fileName error"),
			},
			wantErr: errors.New("fileName error"),
		},
		{
			name: "Test 16: When there is an error in saving data to file",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
				saveDataErr:             errors.New("error in saving data"),
			},
			wantErr: errors.New("error in saving data"),
		},
		{
			name: "Test 17: When rogue mode is on for biggestStakerId and propose exceutes successfully",
			args: args{
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"biggestStakerId"},
				},
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				smallestStake:           big.NewInt(1000),
				smallestStakerId:        1,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: nil,
		},
		{
			name: "Test 18: When rogue mode is on for biggestStakerId and there is an error in getting smallestStakerId",
			args: args{
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"biggestStakerId"},
				},
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				smallestStakerIdErr:     errors.New("smallestStakerId error"),
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 0, 1},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			wantErr: errors.New("smallestStakerId error"),
		},
		{
			name: "Test 19: When there is an error in getting txnOpts",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				sortedProposedBlockIds:  []uint32{2, 1, 0},
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				txnOptsErr:              errors.New("txnOpts error"),
			},
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		SetUpMockInterfaces()

		utilsMock.On("GetBufferedState", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.state, tt.args.stateErr)
		utilsMock.On("GetNumberOfStakers", mock.Anything).Return(tt.args.numStakers, tt.args.numStakerErr)
		cmdUtilsMock.On("GetBiggestStakeAndId", mock.Anything, mock.Anything).Return(tt.args.biggestStake, tt.args.biggestStakerId, tt.args.biggestStakerIdErr)
		cmdUtilsMock.On("GetSmallestStakeAndId", mock.Anything, mock.Anything).Return(tt.args.smallestStake, tt.args.smallestStakerId, tt.args.smallestStakerIdErr)
		utilsMock.On("GetRandaoHash", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.randaoHash, tt.args.randaoHashErr)
		cmdUtilsMock.On("GetIteration", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.iteration)
		utilsMock.On("GetMaxAltBlocks", mock.Anything).Return(tt.args.maxAltBlocks, tt.args.maxAltBlocksErr)
		cmdUtilsMock.On("GetSalt", mock.Anything, mock.Anything).Return(tt.args.salt, tt.args.saltErr)
		cmdUtilsMock.On("GetIteration", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.iteration)
		utilsMock.On("GetNumberOfProposedBlocks", mock.Anything, mock.Anything).Return(tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr)
		utilsMock.On("GetSortedProposedBlockIds", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.sortedProposedBlockIds, tt.args.sortedProposedBlocksIdsErr)
		utilsMock.On("GetMaxAltBlocks", mock.Anything).Return(tt.args.maxAltBlocks, tt.args.maxAltBlocksErr)
		utilsMock.On("GetProposedBlock", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.lastProposedBlockStruct, tt.args.lastProposedBlockStructErr)
		cmdUtilsMock.On("MakeBlock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.medians, tt.args.ids, tt.args.revealDataMaps, tt.args.mediansErr)
		utilsMock.On("ConvertUint32ArrayToBigIntArray", mock.Anything).Return(tt.args.mediansBigInt)
		pathMock.On("GetProposeDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
		fileUtilsMock.On("SaveDataToProposeJsonFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.saveDataErr)
		utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
		blockManagerMock.On("Propose", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.proposeTxn, tt.args.proposeErr)
		transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)

		utils := &UtilsStruct{}
		t.Run(tt.name, func(t *testing.T) {
			err := utils.Propose(rpcParameters, config, account, staker, epoch, latestHeader, stateBuffer, tt.args.rogueData)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for Propose function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for Propose function, got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetBiggestStakeAndId(t *testing.T) {
	var epoch uint32

	type args struct {
		numOfStakers    uint32
		numOfStakersErr error
		stakeArray      []*big.Int
		stakeErr        error
	}
	tests := []struct {
		name      string
		args      args
		wantStake *big.Int
		wantId    uint32
		wantErr   error
	}{
		{
			name: "Test 1: When GetBiggestStakeAndId function executes successfully",
			args: args{
				numOfStakers: 7,
				stakeArray:   []*big.Int{big.NewInt(89999), big.NewInt(70000), big.NewInt(72000), big.NewInt(99999), big.NewInt(200030), big.NewInt(67777), big.NewInt(100011)},
			},
			wantStake: big.NewInt(200030),
			wantId:    5,
			wantErr:   nil,
		},
		{
			name: "Test 2: When numOfStakers is 0",
			args: args{
				numOfStakers: 0,
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("numberOfStakers is 0"),
		},
		{
			name: "Test 3: When there is an error in getting numOfStakers",
			args: args{
				numOfStakersErr: errors.New("numOfStakers error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("numOfStakers error"),
		},
		{
			name: "Test 4: When there is an error in getting stakeArray from batch calls",
			args: args{
				numOfStakers: 5,
				stakeErr:     errors.New("batch calls error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("batch calls error"),
		},
		{
			name: "Test 5: When there are large number of stakers",
			args: args{
				numOfStakers: 999,
				stakeArray:   GenerateDummyStakeSnapshotArray(999),
			},
			wantStake: big.NewInt(999000),
			wantId:    999,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetNumberOfStakers", mock.Anything).Return(tt.args.numOfStakers, tt.args.numOfStakersErr)
			cmdUtilsMock.On("BatchGetStakeSnapshotCalls", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stakeArray, tt.args.stakeErr)

			utils := &UtilsStruct{}

			gotStake, gotId, err := utils.GetBiggestStakeAndId(rpcParameters, epoch)
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			}
			if gotStake.Cmp(tt.wantStake) != 0 {
				t.Errorf("Biggest Stake from GetBiggestStakeAndId function, got = %v, want %v", gotStake, tt.wantStake)
			}
			if gotId != tt.wantId {
				t.Errorf("Staker Id of staker having biggest Influence from GetBiggestStakeAndId function, got = %v, want %v", gotId, tt.wantId)
			}

		})
	}
}

func stakeSnapshotValue(stake string) *big.Int {
	stakeSnapshot, _ := new(big.Int).SetString(stake, 10)
	return stakeSnapshot
}

func TestGetIteration(t *testing.T) {
	var bufferPercent int32

	salt := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	saltBytes32 := [32]byte{}
	copy(saltBytes32[:], salt)

	proposer := types.ElectedProposer{
		BiggestStake:    big.NewInt(1).Mul(big.NewInt(10000000), big.NewInt(1e18)),
		StakerId:        2,
		NumberOfStakers: 10,
		Salt:            saltBytes32,
	}

	type args struct {
		stakeSnapshot    *big.Int
		stakeSnapshotErr error
		stateBuffer      uint64
		stateBufferErr   error
		latestHeader     *Types.Header
		latestHeaderErr  error
		remainingTime    int64
		remainingTimeErr error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1: When getIteration returns a valid iteration",
			args: args{
				stakeSnapshot: big.NewInt(1000),
				stateBuffer:   5,
				latestHeader:  &Types.Header{},
				remainingTime: 10,
			},
			want: 70183,
		},
		{
			name: "Test 2: When there is an error in getting stakeSnapshotValue",
			args: args{
				stakeSnapshot:    big.NewInt(0),
				latestHeader:     &Types.Header{},
				stakeSnapshotErr: errors.New("error in getting stakeSnapshotValue"),
			},
			want: -1,
		},
		{
			name: "Test 3: When getIteration returns an invalid iteration",
			args: args{
				stakeSnapshot: big.NewInt(1),
				stateBuffer:   5,
				latestHeader:  &Types.Header{},
				remainingTime: 2,
			},
			want: -1,
		},
		{
			name: "Test 4: When there is an error in getting remaining time for the state",
			args: args{
				stakeSnapshot:    stakeSnapshotValue("2592145500000000000000000"),
				stateBuffer:      5,
				latestHeader:     &Types.Header{},
				remainingTimeErr: errors.New("remaining time error"),
			},
			want: -1,
		},
		{
			name: "Test 5: When there is an error in getting state buffer",
			args: args{
				stakeSnapshot:  big.NewInt(1000),
				stateBufferErr: errors.New("state buffer error"),
			},
			want: -1,
		},
		{
			name: "Test 6: When there is an error in getting latest header",
			args: args{
				stakeSnapshot:   big.NewInt(1000),
				stateBuffer:     5,
				latestHeaderErr: errors.New("latest header error"),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			SetUpMockInterfaces()
			cmdUtils = &UtilsStruct{}

			utilsMock.On("GetStakeSnapshot", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(1).Mul(tt.args.stakeSnapshot, big.NewInt(1e18)), tt.args.stakeSnapshotErr)
			utilsMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.remainingTime, tt.args.remainingTimeErr)
			utilsMock.On("GetStateBuffer", mock.Anything).Return(tt.args.stateBuffer, tt.args.stateBufferErr)
			clientUtilsMock.On("GetLatestBlockWithRetry", mock.Anything, mock.Anything).Return(tt.args.latestHeader, tt.args.latestHeaderErr)
			if got := cmdUtils.GetIteration(rpcParameters, proposer, bufferPercent); got != tt.want {
				t.Errorf("getIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsElectedProposer(t *testing.T) {

	randaoHash := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	randaoHashBytes32 := [32]byte{}
	copy(randaoHashBytes32[:], randaoHash)

	biggestStake, _ := new(big.Int).SetString("2", 10)

	type args struct {
		address      string
		proposer     types.ElectedProposer
		currentStake *big.Int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1: When staker is 3 and isElectedProposer returns true",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       0,
					Stake:           nil,
					StakerId:        3,
					BiggestStake:    biggestStake,
					NumberOfStakers: 3,
					Salt:            randaoHashBytes32,
					Epoch:           333,
				},
				currentStake: big.NewInt(10000000000),
			},
			want: true,
		},
		{
			name: "Test2: When staker is 2 and isElectedProposer returns false",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       11,
					Stake:           nil,
					StakerId:        2,
					BiggestStake:    biggestStake,
					NumberOfStakers: 3,
					Salt:            randaoHashBytes32,
					Epoch:           29,
				},
				currentStake: big.NewInt(1000000),
			},
			want: false,
		},
		{
			name: "Test3: When staker is 1 and isElectedProposer returns true",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       2,
					Stake:           nil,
					StakerId:        1,
					BiggestStake:    biggestStake,
					NumberOfStakers: 3,
					Salt:            randaoHashBytes32,
					Epoch:           333,
				},
				currentStake: big.NewInt(10000000000),
			},
			want: true,
		},
		{
			name: "Test4: When pseudoRandomNumber is not equal to proposer's stakerID",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       0,
					Stake:           nil,
					StakerId:        3,
					BiggestStake:    biggestStake,
					NumberOfStakers: 3,
					Salt:            [32]byte{},
					Epoch:           333,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		utils := &UtilsStruct{}

		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsElectedProposer(tt.args.proposer, tt.args.currentStake); got != tt.want {
				t.Errorf("IsElectedProposer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pseudoRandomNumberGenerator(t *testing.T) {
	type args struct {
		seed        []byte
		max         uint32
		blockHashes []byte
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test1",
			args: args{
				seed:        []byte{41, 13, 236, 217, 84, 139, 98, 168, 214, 3, 69, 169, 136, 56, 111, 200, 75, 166, 188, 149, 72, 64, 8, 246, 54, 47, 147, 22, 14, 243, 229, 99},
				max:         3,
				blockHashes: []byte{238, 196, 19, 129, 113, 45, 90, 98, 254, 154, 67, 248, 115, 100, 254, 121, 34, 129, 153, 210, 235, 121, 174, 197, 55, 114, 117, 71, 242, 0, 127, 107},
			},
			want: big.NewInt(00),
		},
		{
			name: "Test2",
			args: args{
				seed:        []byte{41, 13, 236, 217, 84, 139, 98, 168, 214, 3, 69, 169, 136, 56, 111, 200, 75, 166, 188, 149, 72, 64, 8, 246, 54, 47, 147, 22, 14, 243, 229, 99},
				max:         3,
				blockHashes: []byte{115, 40, 207, 108, 82, 172, 126, 50, 166, 119, 197, 130, 100, 28, 32, 116, 90, 94, 97, 221, 187, 229, 219, 58, 248, 210, 212, 124, 85, 128, 237, 31},
			},
			want: big.NewInt(0),
		},
		{
			name: "Test3",
			args: args{
				seed:        []byte{177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246},
				max:         3,
				blockHashes: []byte{28, 141, 74, 0, 129, 83, 89, 19, 163, 132, 11, 86, 189, 167, 73, 56, 94, 155, 35, 125, 134, 134, 159, 60, 66, 71, 8, 155, 92, 97, 38, 38},
			},
			want: big.NewInt(2),
		},
		{
			name: "Test4",
			args: args{
				seed:        []byte{138, 53, 172, 251, 193, 95, 248, 26, 57, 174, 125, 52, 79, 215, 9, 242, 142, 134, 0, 180, 170, 140, 101, 198, 182, 75, 254, 127, 227, 107, 209, 155},
				max:         3,
				blockHashes: []byte{28, 141, 74, 0, 129, 83, 89, 19, 163, 132, 11, 86, 189, 167, 73, 56, 94, 155, 35, 125, 134, 134, 159, 60, 66, 71, 8, 155, 92, 97, 38, 38},
			},
			want: big.NewInt(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoRandomNumberGenerator(tt.args.seed, tt.args.max, tt.args.blockHashes); got.Cmp(tt.want) != 0 {
				t.Errorf("pseudoRandomNumberGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeBlock(t *testing.T) {
	var (
		blockNumber *big.Int
		epoch       uint32
	)

	randomValue := big.NewInt(1111)

	type args struct {
		revealedDataMaps     *types.RevealedDataMaps
		revealedDataMapsErr  error
		activeCollections    []uint16
		activeCollectionsErr error
		rogueData            types.Rogue
	}
	tests := []struct {
		name    string
		args    args
		want    []*big.Int
		want1   []uint16
		want2   *types.RevealedDataMaps
		wantErr bool
	}{
		{
			name: "Test 1: When MakeBlock executes successfully and there is no rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]*big.Int{0: {big.NewInt(1), big.NewInt(1)}, 1: {big.NewInt(100), big.NewInt(100)}, 2: {big.NewInt(200), big.NewInt(200)}},
					VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(1000), big.NewInt(100).String(): big.NewInt(2000), big.NewInt(200).String(): big.NewInt(3000)},
					InfluenceSum:         map[uint16]*big.Int{0: big.NewInt(500), 1: big.NewInt(10000), 2: big.NewInt(10000), 3: big.NewInt(10000)},
				},
				activeCollections: []uint16{0, 1, 2},
			},
			want:  []*big.Int{big.NewInt(1), big.NewInt(100), big.NewInt(200)},
			want1: []uint16{0, 1, 2},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{0: {big.NewInt(1), big.NewInt(1)}, 1: {big.NewInt(100), big.NewInt(100)}, 2: {big.NewInt(200), big.NewInt(200)}},
				VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(1000), big.NewInt(1).String(): big.NewInt(1000), big.NewInt(100).String(): big.NewInt(2000), big.NewInt(100).String(): big.NewInt(2000), big.NewInt(200).String(): big.NewInt(3000), big.NewInt(200).String(): big.NewInt(3000)},
				InfluenceSum:         map[uint16]*big.Int{0: big.NewInt(250), 1: big.NewInt(2500), 2: big.NewInt(2500), 3: big.NewInt(10000)},
			},
			wantErr: false,
		},
		{
			name: "Test 2 : When there is an error in getting revealedDataMaps",
			args: args{
				revealedDataMapsErr: errors.New("error in getting revealedDataMaps"),
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "Test 3 : When there is an error in getting activeCollections",
			args: args{
				revealedDataMaps:     &types.RevealedDataMaps{},
				activeCollectionsErr: errors.New("error in getting activeCollections"),
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "Test 4: When MakeBlock executes successfully and there is missingIds rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
					VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"missingIds"},
				},
			},
			want:  []*big.Int{big.NewInt(1)},
			want1: []uint16{3},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
				VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(50)},
			},
			wantErr: false,
		},
		{
			name: "Test 5: When MakeBlock executes successfully and there is extraIds rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
					VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"extraIds"},
				},
			},
			want:  []*big.Int{big.NewInt(1), randomValue},
			want1: []uint16{2, 3},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
				VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(50)},
			},
			wantErr: false,
		},
		{
			name: "Test 5: When MakeBlock executes successfully and there is medians rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
					VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"medians"},
				},
			},
			want:  []*big.Int{randomValue},
			want1: []uint16{2},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{1: {big.NewInt(1), big.NewInt(2), big.NewInt(3)}},
				VoteWeights:          map[string]*big.Int{big.NewInt(1).String(): big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			cmdUtilsMock.On("GetSortedRevealedValues", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.revealedDataMaps, tt.args.revealedDataMapsErr)
			utilsMock.On("GetActiveCollectionIds", mock.Anything, mock.Anything).Return(tt.args.activeCollections, tt.args.activeCollectionsErr)
			utilsMock.On("GetRogueRandomValue", mock.Anything).Return(randomValue)
			ut := &UtilsStruct{}
			got, got1, got2, err := ut.MakeBlock(rpcParameters, blockNumber, epoch, tt.args.rogueData)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeBlock() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MakeBlock() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("MakeBlock() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestGetSortedRevealedValues(t *testing.T) {
	var (
		blockNumber *big.Int
		epoch       uint32
	)

	type args struct {
		assignedAssets    []types.RevealedStruct
		assignedAssetsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *types.RevealedDataMaps
		wantErr bool
	}{
		{
			name: "Test 1: When GetSortedRevealedValues executes successfully",
			args: args{
				assignedAssets: []types.RevealedStruct{
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 3, Value: big.NewInt(601)},
							{LeafId: 6, Value: big.NewInt(750)},
							{LeafId: 1, Value: big.NewInt(400)},
						},
						Influence: big.NewInt(10000000),
					},
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 10, Value: big.NewInt(1100)},
							{LeafId: 5, Value: big.NewInt(900)},
							{LeafId: 7, Value: big.NewInt(302)},
						},
						Influence: big.NewInt(20000000),
					},
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 3, Value: big.NewInt(600)},
							{LeafId: 7, Value: big.NewInt(300)},
							{LeafId: 9, Value: big.NewInt(1600)},
						},
						Influence: big.NewInt(30000000),
					},
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 10, Value: big.NewInt(1105)},
							{LeafId: 8, Value: big.NewInt(950)},
							{LeafId: 7, Value: big.NewInt(300)},
						},
						Influence: big.NewInt(40000000),
					},
				},
			},
			want: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{
					1:  {big.NewInt(400)},
					3:  {big.NewInt(600), big.NewInt(601)},
					5:  {big.NewInt(900)},
					6:  {big.NewInt(750)},
					7:  {big.NewInt(300), big.NewInt(302)},
					8:  {big.NewInt(950)},
					9:  {big.NewInt(1600)},
					10: {big.NewInt(1100), big.NewInt(1105)},
				},
				VoteWeights: map[string]*big.Int{
					"1600": big.NewInt(30000000),
					"300":  big.NewInt(70000000),
					"302":  big.NewInt(20000000),
					"400":  big.NewInt(10000000),
					"600":  big.NewInt(30000000),
					"601":  big.NewInt(10000000),
					"750":  big.NewInt(10000000),
					"1100": big.NewInt(20000000),
					"1105": big.NewInt(40000000),
					"950":  big.NewInt(40000000),
					"900":  big.NewInt(20000000),
				},
				InfluenceSum: map[uint16]*big.Int{
					1:  big.NewInt(10000000),
					3:  big.NewInt(40000000),
					5:  big.NewInt(20000000),
					6:  big.NewInt(10000000),
					7:  big.NewInt(90000000),
					8:  big.NewInt(40000000),
					9:  big.NewInt(30000000),
					10: big.NewInt(60000000),
				},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there are multiple equal and unequal vote values for single leafId",
			args: args{
				assignedAssets: []types.RevealedStruct{
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 1, Value: big.NewInt(600)},
							{LeafId: 2, Value: big.NewInt(750)},
							{LeafId: 3, Value: big.NewInt(400)},
						},
						Influence: big.NewInt(10000000),
					},
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 1, Value: big.NewInt(601)},
							{LeafId: 2, Value: big.NewInt(752)},
						},
						Influence: big.NewInt(20000000),
					},
					{
						RevealedValues: []types.AssignedAsset{
							{LeafId: 1, Value: big.NewInt(601)},
							{LeafId: 2, Value: big.NewInt(756)},
							{LeafId: 4, Value: big.NewInt(1600)},
						},
						Influence: big.NewInt(30000000),
					},
				},
			},
			want: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{
					1: {big.NewInt(600), big.NewInt(601)},
					2: {big.NewInt(750), big.NewInt(752), big.NewInt(756)},
					3: {big.NewInt(400)},
					4: {big.NewInt(1600)},
				},
				VoteWeights: map[string]*big.Int{
					"1600": big.NewInt(30000000),
					"400":  big.NewInt(10000000),
					"600":  big.NewInt(10000000),
					"601":  big.NewInt(50000000),
					"750":  big.NewInt(10000000),
					"752":  big.NewInt(20000000),
					"756":  big.NewInt(30000000),
				},
				InfluenceSum: map[uint16]*big.Int{
					1: big.NewInt(60000000),
					2: big.NewInt(60000000),
					3: big.NewInt(10000000),
					4: big.NewInt(30000000),
				},
			},
		},
		{
			name: "Test 3: When assignedAssets is empty",
			args: args{
				assignedAssets: []types.RevealedStruct{},
			},
			want: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]*big.Int{},
				VoteWeights:          map[string]*big.Int{},
				InfluenceSum:         map[uint16]*big.Int{},
			},
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting assignedAssets",
			args: args{
				assignedAssetsErr: errors.New("error in getting assets"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			cmdUtilsMock.On("IndexRevealEventsOfCurrentEpoch", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.assignedAssets, tt.args.assignedAssetsErr)
			ut := &UtilsStruct{}
			got, err := ut.GetSortedRevealedValues(rpcParameters, blockNumber, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSortedRevealedValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortedRevealedValues() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSmallestStakeAndId(t *testing.T) {
	var epoch uint32

	type args struct {
		numOfStakers    uint32
		numOfStakersErr error
		stake           *big.Int
		stakeErr        error
	}
	tests := []struct {
		name      string
		args      args
		wantStake *big.Int
		wantId    uint32
		wantErr   error
	}{
		{
			name: "Test 1: When GetSmallestStakeAndId function executes successfully",
			args: args{
				numOfStakers: 4,
				stake:        big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)),
			},
			wantStake: big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)),
			wantId:    1,
			wantErr:   nil,
		},
		{
			name: "Test 2: When the numberOfStakers is 0",
			args: args{
				numOfStakers: 0,
				stake:        big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("numberOfStakers is 0"),
		},
		{
			name: "Test 3: When there is an error in getting numOfStakers",
			args: args{
				numOfStakersErr: errors.New("numOfStakers error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("numOfStakers error"),
		},
		{
			name: "Test 4: When there is an error in getting stake",
			args: args{
				numOfStakers: 5,
				stakeErr:     errors.New("stake error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("stake error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetNumberOfStakers", mock.Anything).Return(tt.args.numOfStakers, tt.args.numOfStakersErr)
			utilsMock.On("GetStakeSnapshot", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.stake, tt.args.stakeErr)

			utils := &UtilsStruct{}

			gotStake, gotId, err := utils.GetSmallestStakeAndId(rpcParameters, epoch)
			if gotStake.Cmp(tt.wantStake) != 0 {
				t.Errorf("Smallest Stake from GetSmallestStakeAndId function, got = %v, want %v", gotStake, tt.wantStake)
			}
			if gotId != tt.wantId {
				t.Errorf("Staker Id of staker having smallest Influence from GetSmallestStakeAndId function, got = %v, want %v", gotId, tt.wantId)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetSmallestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetSmallestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestBatchGetStakeCalls(t *testing.T) {
	var epoch uint32

	voteManagerABI, _ := abi.JSON(strings.NewReader(bindings.VoteManagerMetaData.ABI))

	type args struct {
		ABI              abi.ABI
		numberOfStakers  uint32
		parseErr         error
		batchCallResults [][]interface{}
		batchCallError   error
	}
	tests := []struct {
		name       string
		args       args
		wantStakes []*big.Int
		wantErr    error
	}{
		{
			name: "Test 1: When BatchGetStakeCalls executes successfully",
			args: args{
				ABI:             voteManagerABI,
				numberOfStakers: 3,
				batchCallResults: [][]interface{}{
					{big.NewInt(10)},
					{big.NewInt(11)},
					{big.NewInt(12)},
				},
			},
			wantStakes: []*big.Int{
				big.NewInt(10),
				big.NewInt(11),
				big.NewInt(12),
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in parsing voteManager ABI",
			args: args{
				parseErr: errors.New("parse error"),
			},
			wantErr: errors.New("parse error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			abiUtilsMock.On("Parse", mock.Anything).Return(tt.args.ABI, tt.args.parseErr)
			clientUtilsMock.On("BatchCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.batchCallResults, tt.args.batchCallError)

			ut := &UtilsStruct{}
			gotStakes, err := ut.BatchGetStakeSnapshotCalls(rpcParameters, epoch, tt.args.numberOfStakers)

			if err == nil || tt.wantErr == nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}

			assert.Equal(t, tt.wantStakes, gotStakes)
		})
	}
}

func BenchmarkGetIteration(b *testing.B) {
	var bufferPercent int32

	salt := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	saltBytes32 := [32]byte{}
	copy(saltBytes32[:], salt)

	proposer := types.ElectedProposer{
		BiggestStake:    big.NewInt(1).Mul(big.NewInt(10000000), big.NewInt(1e18)),
		StakerId:        2,
		NumberOfStakers: 5,
		Salt:            saltBytes32,
	}

	var table = []struct {
		stakeSnapshot *big.Int
	}{
		{stakeSnapshot: big.NewInt(1000)},
		{stakeSnapshot: big.NewInt(10000)},
		{stakeSnapshot: big.NewInt(100000)},
		{stakeSnapshot: big.NewInt(1000000)},
		{stakeSnapshot: big.NewInt(10000000)},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("Stakers_Stake_%d", v.stakeSnapshot), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SetUpMockInterfaces()

				cmdUtils = &UtilsStruct{}

				utilsMock.On("GetStakeSnapshot", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(1).Mul(v.stakeSnapshot, big.NewInt(1e18)), nil)
				utilsMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything, mock.Anything).Return(int64(100), nil)
				utilsMock.On("GetStateBuffer", mock.Anything).Return(uint64(5), nil)
				clientUtilsMock.On("GetLatestBlockWithRetry", mock.Anything, mock.Anything).Return(&Types.Header{}, nil)

				cmdUtils.GetIteration(rpcParameters, proposer, bufferPercent)
			}
		})
	}
}

func BenchmarkGetBiggestStakeAndId(b *testing.B) {
	var epoch uint32

	var table = []struct {
		numOfStakers uint32
	}{
		{numOfStakers: 10},
		{numOfStakers: 1000},
		{numOfStakers: 100000},
		{numOfStakers: 1000000},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("Stakers_Stake_%d", v.numOfStakers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SetUpMockInterfaces()

				utilsMock.On("GetNumberOfStakers", mock.Anything).Return(v.numOfStakers, nil)
				cmdUtilsMock.On("BatchGetStakeSnapshotCalls", mock.Anything, mock.Anything, mock.Anything).Return(GenerateDummyStakeSnapshotArray(v.numOfStakers), nil)

				ut := &UtilsStruct{}
				_, _, err := ut.GetBiggestStakeAndId(rpcParameters, epoch)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkGetSortedRevealedValues(b *testing.B) {
	var (
		blockNumber *big.Int
		epoch       uint32
	)
	table := []struct {
		numOfAssignedAssets int
		numOfRevealedValues uint16
	}{
		{numOfAssignedAssets: 1, numOfRevealedValues: 10},
		{numOfAssignedAssets: 10, numOfRevealedValues: 100},
		{numOfAssignedAssets: 100, numOfRevealedValues: 1000},
		{numOfAssignedAssets: 1000, numOfRevealedValues: 10000},
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("Number_Of_Assigned_Assets_%d, Number_Of_Revealed_Votes_%d", v.numOfAssignedAssets, v.numOfRevealedValues), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SetUpMockInterfaces()

				asset := GetDummyRevealedValues(v.numOfRevealedValues)

				cmdUtilsMock.On("IndexRevealEventsOfCurrentEpoch", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(GetDummyAssignedAssets(asset, v.numOfAssignedAssets), nil)
				ut := &UtilsStruct{}
				_, err := ut.GetSortedRevealedValues(rpcParameters, blockNumber, epoch)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkMakeBlock(b *testing.B) {
	var (
		blockNumber *big.Int
		epoch       uint32
	)

	table := []struct {
		numOfVotes int
	}{
		{numOfVotes: 1},
		{numOfVotes: 100},
		{numOfVotes: 1000},
		{numOfVotes: 10000},
		{numOfVotes: 100000},
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("Number_Of_Votes_%d", v.numOfVotes), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SetUpMockInterfaces()

				votes := GetDummyVotes(v.numOfVotes)

				cmdUtilsMock.On("GetSortedRevealedValues", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]*big.Int{0: votes},
					VoteWeights:          map[string]*big.Int{(big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))).String(): big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{0: big.NewInt(100)},
				}, nil)
				utilsMock.On("GetActiveCollectionIds", mock.Anything, mock.Anything).Return([]uint16{1}, nil)
				ut := &UtilsStruct{}
				_, _, _, err := ut.MakeBlock(rpcParameters, blockNumber, epoch, types.Rogue{IsRogue: false})
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func GenerateDummyStakeSnapshotArray(numOfStakers uint32) []*big.Int {
	stakeSnapshotArray := make([]*big.Int, numOfStakers)
	for i := 0; i < int(numOfStakers); i++ {
		// For testing purposes, we will assign a stake value of (i + 1) * 1000
		stakeSnapshotArray[i] = big.NewInt(int64(i+1) * 1000)
	}
	return stakeSnapshotArray
}

func GetDummyVotes(numOfVotes int) []*big.Int {
	var result []*big.Int
	for i := 0; i < numOfVotes; i++ {
		result = append(result, big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18)))
	}
	return result
}

func GetDummyAssignedAssets(asset types.RevealedStruct, numOfAssignedAssets int) []types.RevealedStruct {
	var assignedAssets []types.RevealedStruct
	for i := 1; i <= numOfAssignedAssets; i++ {
		assignedAssets = append(assignedAssets, asset)
	}
	return assignedAssets
}

func GetDummyRevealedValues(numOfRevealedValues uint16) types.RevealedStruct {
	var revealedValues []types.AssignedAsset
	var i uint16
	for i = 1; i < numOfRevealedValues; i++ {
		revealedValues = append(revealedValues, types.AssignedAsset{
			LeafId: i,
			Value:  big.NewInt(1000),
		})
	}
	return types.RevealedStruct{
		RevealedValues: revealedValues,
		Influence:      big.NewInt(1000),
	}
}
