package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	Mocks "razor/utils/mocks"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestPropose(t *testing.T) {

	var (
		client      *ethclient.Client
		account     types.Account
		config      types.Configurations
		staker      bindings.StructsStaker
		epoch       uint32
		blockNumber *big.Int
	)

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	salt := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	saltBytes32 := [32]byte{}
	copy(saltBytes32[:], salt)

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
		randaoHash                 [32]byte
		randaoHashErr              error
		bufferPercent              int32
		bufferPercentErr           error
		salt                       [32]byte
		saltErr                    error
		iteration                  int
		numOfProposedBlocks        uint8
		numOfProposedBlocksErr     error
		maxAltBlocks               uint8
		maxAltBlocksErr            error
		lastIteration              *big.Int
		lastProposedBlockStruct    bindings.StructsBlock
		lastProposedBlockStructErr error
		medians                    []uint32
		ids                        []uint16
		revealDataMaps             *types.RevealedDataMaps
		mediansErr                 error
		fileName                   string
		fileNameErr                error
		saveDataErr                error
		mediansBigInt              []*big.Int
		txnOpts                    *bind.TransactOpts
		proposeTxn                 *Types.Transaction
		proposeErr                 error
		hash                       common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocksErr:         errors.New("maxAltBlocks error"),
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
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
				maxAltBlocks:               2,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStructErr: errors.New("lastProposedBlockStruct error"),
				medians:                    []uint32{6701548, 478307},
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("lastProposedBlockStruct error"),
		},
		{
			name: "Test 10: When numOfProposedBlocks >= maxAltBlocks and current iteration is greater than iteration of last proposed block ",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          5,
				biggestStake:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:     2,
				salt:                saltBytes32,
				iteration:           2,
				numOfProposedBlocks: 4,
				maxAltBlocks:        2,
				lastIteration:       big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(1),
				},
				medians:    []uint32{6701548, 478307},
				txnOpts:    txnOpts,
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 11: When numOfProposedBlocks >= maxAltBlocks and current iteration is less than iteration of last proposed block and propose transaction is successful",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          5,
				biggestStake:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:     2,
				salt:                saltBytes32,
				iteration:           1,
				numOfProposedBlocks: 4,
				maxAltBlocks:        2,
				lastIteration:       big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(2),
				},
				medians:    []uint32{6701548, 478307},
				txnOpts:    txnOpts,
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 12: When there is an error in getting medians",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				mediansErr:              errors.New("makeBlock error"),
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("makeBlock error"),
		},
		{
			name: "Test 13: When Propose transaction fails",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeErr:              errors.New("propose error"),
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("propose error"),
		},
		{
			name: "Test 14: When there is an error in getting fileName",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				fileNameErr:             errors.New("fileName error"),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 15: When there is an error in saving data to file",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestStake:            big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:         2,
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				saveDataErr:             errors.New("error in saving data"),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 17: When there is an error in getting buffer percent",
			args: args{
				state:            2,
				staker:           bindings.StructsStaker{},
				numStakers:       5,
				biggestStake:     big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakerId:  2,
				bufferPercentErr: errors.New("buffer error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("buffer error"),
		},
		{
			name: "Test 18: When rogue mode is on for biggestStakerId Propose function executes successfully",
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
				salt:                    saltBytes32,
				iteration:               1,
				numOfProposedBlocks:     3,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{6701548, 478307},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
	}
	for _, tt := range tests {

		utilsMock := new(mocks.UtilsInterface)
		cmdUtilsMock := new(mocks.UtilsCmdInterface)
		blockManagerUtilsMock := new(mocks.BlockManagerInterface)
		transactionUtilsMock := new(mocks.TransactionInterface)

		razorUtils = utilsMock
		cmdUtils = cmdUtilsMock
		blockManagerUtils = blockManagerUtilsMock
		transactionUtils = transactionUtilsMock

		utilsMock.On("GetDelayedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
		utilsMock.On("GetNumberOfStakers", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.numStakers, tt.args.numStakerErr)
		cmdUtilsMock.On("GetBiggestStakeAndId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.biggestStake, tt.args.biggestStakerId, tt.args.biggestStakerIdErr)
		utilsMock.On("GetRandaoHash", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.randaoHash, tt.args.randaoHashErr)
		cmdUtilsMock.On("GetIteration", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.iteration)
		utilsMock.On("GetMaxAltBlocks", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.maxAltBlocks, tt.args.maxAltBlocksErr)
		cmdUtilsMock.On("GetSalt", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.salt, tt.args.saltErr)
		utilsMock.On("GetNumberOfProposedBlocks", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr)
		utilsMock.On("GetMaxAltBlocks", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.maxAltBlocks, tt.args.maxAltBlocksErr)
		utilsMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.lastProposedBlockStruct, tt.args.lastProposedBlockStructErr)
		cmdUtilsMock.On("MakeBlock", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.medians, tt.args.ids, tt.args.revealDataMaps, tt.args.mediansErr)
		utilsMock.On("ConvertUint32ArrayToBigIntArray", mock.Anything).Return(tt.args.mediansBigInt)
		utilsMock.On("GetProposeDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
		utilsMock.On("SaveDataToProposeJsonFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.saveDataErr)
		utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
		blockManagerUtilsMock.On("Propose", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.proposeTxn, tt.args.proposeErr)
		transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
		cmdUtilsMock.On("GetBufferPercent").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)

		utils := &UtilsStruct{}
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.Propose(client, config, account, staker, epoch, blockNumber, tt.args.rogueData)
			if got != tt.want {
				t.Errorf("Txn hash for Propose function, got = %v, want %v", got, tt.want)
			}
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
	var client *ethclient.Client
	var address string
	var epoch uint32

	type args struct {
		numOfStakers     uint32
		numOfStakersErr  error
		bufferPercent    int32
		bufferPercentErr error
		remainingTime    int64
		remainingTimeErr error
		stake            []*big.Int
		stakeErr         error
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
				numOfStakers:  2,
				remainingTime: 10,
				stake:         []*big.Int{big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18))},
			},
			wantStake: big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)),
			wantId:    1,
			wantErr:   nil,
		},
		{
			name: "Test 2: When getBiggestStakeAndId function executes successfully with more number of stakers",
			args: args{
				numOfStakers:  5,
				remainingTime: 10,
				stake:         []*big.Int{big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(32432), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(32), big.NewInt(1e18)), big.NewInt(1e18), big.NewInt(1e10)},
			},
			wantStake: big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)),
			wantId:    1,
			wantErr:   nil,
		},
		{
			name: "Test 3: When there is an error in getting numOfStakers",
			args: args{
				numOfStakersErr: errors.New("numOfStakers error"),
				remainingTime:   10,
				stake:           []*big.Int{big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18))},
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("numOfStakers error"),
		},
		{
			name: "Test 4: When there is an error in getting stake",
			args: args{
				numOfStakers:  5,
				remainingTime: 10,
				stakeErr:      errors.New("stake error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("stake error"),
		},
		{
			name: "Test 5: When there is an error in getting remaining time",
			args: args{
				numOfStakers:     2,
				remainingTime:    10,
				remainingTimeErr: errors.New("time error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("time error"),
		},
		{
			name: "Test 6: When there is a timeout case",
			args: args{
				numOfStakers:  100000,
				remainingTime: 1,
				stake:         []*big.Int{big.NewInt(1).Mul(big.NewInt(5326), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18))},
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("state timeout error"),
		},
		{
			name: "Test 7: When there is an error in getting buffer percent",
			args: args{
				numOfStakers:     2,
				bufferPercentErr: errors.New("buffer error"),
			},
			wantStake: nil,
			wantId:    0,
			wantErr:   errors.New("buffer error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(Mocks.Utils)

			razorUtils = utilsMock
			utilsInterface = utilsPkgMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("GetNumberOfStakers", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.numOfStakers, tt.args.numOfStakersErr)
			if tt.args.stake != nil {
				utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.stake[uint32(0)], tt.args.stakeErr)
			} else {
				utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(nil, tt.args.stakeErr)

			}
			utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.stake, tt.args.stakeErr)
			utilsPkgMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(tt.args.remainingTime, tt.args.remainingTimeErr)
			cmdUtilsMock.On("GetBufferPercent").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)

			utils := &UtilsStruct{}

			gotStake, gotId, err := utils.GetBiggestStakeAndId(client, address, epoch)
			if gotStake.Cmp(tt.wantStake) != 0 {
				t.Errorf("Biggest Stake from GetBiggestStakeAndId function, got = %v, want %v", gotStake, tt.wantStake)
			}
			if gotId != tt.wantId {
				t.Errorf("Staker Id of staker having biggest Influence from GetBiggestStakeAndId function, got = %v, want %v", gotId, tt.wantId)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func stakeSnapshotValue(stake string) *big.Int {
	stakeSnapshot, _ := new(big.Int).SetString(stake, 10)
	return stakeSnapshot
}

func TestGetIteration(t *testing.T) {
	var client *ethclient.Client
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
				remainingTime: 10,
			},
			want: 70183,
		},
		{
			name: "Test 2: When there is an error in getting stakeSnapshotValue",
			args: args{
				stakeSnapshot:    big.NewInt(0),
				stakeSnapshotErr: errors.New("error in getting stakeSnapshotValue"),
			},
			want: -1,
		},
		{
			name: "Test 3: When getIteration returns an invalid iteration",
			args: args{
				stakeSnapshot: big.NewInt(1),
				remainingTime: 2,
			},
			want: -1,
		},
		{
			name: "Test 4: When there is an error in getting remaining time for the state",
			args: args{
				stakeSnapshot:    stakeSnapshotValue("2592145500000000000000000"),
				remainingTimeErr: errors.New("remaining time error"),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			utilsPkgMock := new(Mocks.Utils)

			razorUtils = utilsMock
			cmdUtils = &UtilsStruct{}
			utilsInterface = utilsPkgMock

			utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(big.NewInt(1).Mul(tt.args.stakeSnapshot, big.NewInt(1e18)), tt.args.stakeSnapshotErr)
			utilsPkgMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(tt.args.remainingTime, tt.args.remainingTimeErr)

			if got := cmdUtils.GetIteration(client, proposer, bufferPercent); got != tt.want {
				t.Errorf("getIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfluencedMedian(t *testing.T) {
	type args struct {
		sortedVotes            []*big.Int
		totalInfluenceRevealed *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test if sortedVotes is empty",
			args: args{
				sortedVotes:            []*big.Int{},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
			},
			want: big.NewInt(0),
		},
		{
			name: "Test if totalInfluenceRevealed is 0",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(0),
			},
			want: big.NewInt(1).Mul(big.NewInt(2093037800), big.NewInt(1e18)),
		},
		{
			name: "Test if all the values are present",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
			},
			want: big.NewInt(498342),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils := &UtilsStruct{}
			if got := utils.InfluencedMedian(tt.args.sortedVotes, tt.args.totalInfluenceRevealed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("influencedMedian() = %v, want %v", got, tt.want)
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

func BenchmarkGetIteration(b *testing.B) {
	var client *ethclient.Client
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

	var table = []struct {
		stakeSnapshot *big.Int
		batchSize     int
	}{
		{stakeSnapshot: big.NewInt(1000), batchSize: 1000},
		{stakeSnapshot: big.NewInt(1000), batchSize: 100},
		{stakeSnapshot: big.NewInt(1000), batchSize: 10},
		{stakeSnapshot: big.NewInt(10000), batchSize: 1000},
		{stakeSnapshot: big.NewInt(100000), batchSize: 1000},
		{stakeSnapshot: big.NewInt(1000000), batchSize: 1000},
		{stakeSnapshot: big.NewInt(10000000), batchSize: 1000},
		{stakeSnapshot: big.NewInt(10000000), batchSize: 100},
		{stakeSnapshot: big.NewInt(10000000), batchSize: 10},
	}

	for _, v := range table {
		var timeRecorded []int64
		b.Run(fmt.Sprintf("Stakers_Stake_%d, BatchSize_%d", v.stakeSnapshot, v.batchSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				start := time.Now()
				utilsMock := new(mocks.UtilsInterface)
				utilsPkgMock := new(Mocks.Utils)

				razorUtils = utilsMock
				cmdUtils = &UtilsStruct{}
				utilsInterface = utilsPkgMock

				utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(big.NewInt(1).Mul(v.stakeSnapshot, big.NewInt(1e18)), nil)
				utilsPkgMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(int64(100), nil)

				core.BatchSize = v.batchSize
				cmdUtils.GetIteration(client, proposer, bufferPercent)

				timeElapsed := time.Since(start).Microseconds()
				timeRecorded = append(timeRecorded, timeElapsed)
			}
		})
		fmt.Println("Median time taken: ", CalculateMedian(timeRecorded))
		fmt.Println()
	}
}

func CalculateMedian(arr []int64) int64 {
	var median int64

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] }) // sort the numbers
	l := len(arr)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (arr[l/2-1] + arr[l/2]) / 2
	} else {
		median = arr[l/2]
	}

	return median
}

func TestMakeBlock(t *testing.T) {
	var (
		client      *ethclient.Client
		blockNumber *big.Int
		epoch       uint32
	)

	randomValue := utils.GetRogueRandomMedianValue()

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
		want    []uint32
		want1   []uint16
		want2   *types.RevealedDataMaps
		wantErr bool
	}{
		{
			name: "Test 1: When MakeBlock executes successfully and there is no rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
					VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
			},
			want:  []uint32{1},
			want1: []uint16{2},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
				VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(50)},
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
					SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
					VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"missingIds"},
				},
			},
			want:  []uint32{1},
			want1: []uint16{3},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
				VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(50)},
			},
			wantErr: false,
		},
		{
			name: "Test 5: When MakeBlock executes successfully and there is extraIds rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
					VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"extraIds"},
				},
			},
			want:  []uint32{1, randomValue},
			want1: []uint16{2, 3},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
				VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(50)},
			},
			wantErr: false,
		},
		{
			name: "Test 5: When MakeBlock executes successfully and there is medians rogue mode",
			args: args{
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
					VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
				},
				activeCollections: []uint16{1, 2},
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"medians"},
				},
			},
			want:  []uint32{randomValue},
			want1: []uint16{2},
			want2: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]uint32{1: {1, 2, 3}},
				VoteWeights:          map[uint32]*big.Int{1: big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			cmdUtilsMock.On("GetSortedRevealedValues", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.revealedDataMaps, tt.args.revealedDataMapsErr)
			utilsMock.On("GetActiveCollections", mock.Anything).Return(tt.args.activeCollections, tt.args.activeCollectionsErr)
			utilsMock.On("GetRogueRandomMedianValue").Return(randomValue)
			ut := &UtilsStruct{}
			got, got1, got2, err := ut.MakeBlock(client, blockNumber, epoch, tt.args.rogueData)
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
		client      *ethclient.Client
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
				assignedAssets: []types.RevealedStruct{{RevealedValues: []types.AssignedAsset{{LeafId: 1, Value: 100}}, Influence: big.NewInt(100)}},
			},
			want: &types.RevealedDataMaps{
				SortedRevealedValues: map[uint16][]uint32{1: {100}},
				VoteWeights:          map[uint32]*big.Int{100: big.NewInt(100)},
				InfluenceSum:         map[uint16]*big.Int{1: big.NewInt(100)},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting assignedAssets",
			args: args{
				assignedAssetsErr: errors.New("error in getting assets"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			cmdUtils = cmdUtilsMock

			cmdUtilsMock.On("IndexRevealEventsOfCurrentEpoch", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.assignedAssets, tt.args.assignedAssetsErr)
			ut := &UtilsStruct{}
			got, err := ut.GetSortedRevealedValues(client, blockNumber, epoch)
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

func BenchmarkGetBiggestStakeAndId(b *testing.B) {
	var client *ethclient.Client
	var address string
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
				utilsMock := new(mocks.UtilsInterface)
				cmdUtilsMock := new(mocks.UtilsCmdInterface)
				utilsPkgMock := new(Mocks.Utils)

				razorUtils = utilsMock
				utilsInterface = utilsPkgMock
				cmdUtils = cmdUtilsMock

				utilsMock.On("GetNumberOfStakers", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(v.numOfStakers, nil)
				utilsMock.On("GetStakeSnapshot", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(big.NewInt(10000), nil)
				utilsPkgMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(int64(150), nil)
				cmdUtilsMock.On("GetBufferPercent").Return(int32(60), nil)

				ut := &UtilsStruct{}
				_, _, err := ut.GetBiggestStakeAndId(client, address, epoch)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkInfluencedMedian(b *testing.B) {
	var table = []struct {
		numOfSortedVotes       int
		totalInfluenceRevealed *big.Int
	}{
		{numOfSortedVotes: 10, totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18))},
		{numOfSortedVotes: 100, totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(42000), big.NewInt(1e18))},
		{numOfSortedVotes: 500, totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(42000), big.NewInt(1e18))},
		{numOfSortedVotes: 1000, totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(420000), big.NewInt(1e18))},
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("Number_Of_Sorted_Votes_%d", v.numOfSortedVotes), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				utils := &UtilsStruct{}
				sortedVotes := GetDummyVotes(v.numOfSortedVotes)
				utils.InfluencedMedian(sortedVotes, v.totalInfluenceRevealed)
			}
		})
	}
}

func BenchmarkGetSortedRevealedValues(b *testing.B) {
	var (
		client      *ethclient.Client
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
				cmdUtilsMock := new(mocks.UtilsCmdInterface)

				cmdUtils = cmdUtilsMock
				asset := GetDummyRevealedValues(v.numOfRevealedValues)

				cmdUtilsMock.On("IndexRevealEventsOfCurrentEpoch", mock.Anything, mock.Anything, mock.Anything).Return(GetDummyAssignedAssets(asset, v.numOfAssignedAssets), nil)
				ut := &UtilsStruct{}
				_, err := ut.GetSortedRevealedValues(client, blockNumber, epoch)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkMakeBlock(b *testing.B) {
	var (
		client      *ethclient.Client
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
				utilsMock := new(mocks.UtilsInterface)
				cmdUtilsMock := new(mocks.UtilsCmdInterface)

				razorUtils = utilsMock
				cmdUtils = cmdUtilsMock

				votes := GetUint32DummyVotes(v.numOfVotes)

				cmdUtilsMock.On("GetSortedRevealedValues", mock.Anything, mock.Anything, mock.Anything).Return(&types.RevealedDataMaps{
					SortedRevealedValues: map[uint16][]uint32{0: votes},
					VoteWeights:          map[uint32]*big.Int{100: big.NewInt(100)},
					InfluenceSum:         map[uint16]*big.Int{0: big.NewInt(100)},
				}, nil)
				utilsMock.On("GetActiveCollections", mock.Anything).Return([]uint16{1}, nil)
				ut := &UtilsStruct{}
				_, _, _, err := ut.MakeBlock(client, blockNumber, epoch, types.Rogue{IsRogue: false})
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func GetDummyVotes(numOfVotes int) []*big.Int {
	var result []*big.Int
	for i := 0; i < numOfVotes; i++ {
		result = append(result, big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18)))
	}
	return result
}

func GetUint32DummyVotes(numOfVotes int) []uint32 {
	var result []uint32
	for i := 0; i < numOfVotes; i++ {
		result = append(result, 100)
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
			Value:  1000,
		})
	}
	return types.RevealedStruct{
		RevealedValues: revealedValues,
		Influence:      big.NewInt(1000),
	}
}
