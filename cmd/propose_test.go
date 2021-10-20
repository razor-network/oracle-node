package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	randMath "math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"testing"
)

func TestPropose(t *testing.T) {

	var client *ethclient.Client
	var account types.Account
	var config types.Configurations
	var stakerId uint32
	var epoch uint32
	var rogueMode bool

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	razorUtils := UtilsMock{}
	proposeUtils := ProposeUtilsMock{}
	blockManagerUtils := BlockManagerMock{}
	transactionUtils := TransactionMock{}

	type args struct {
		state                      int64
		stateErr                   error
		staker                     bindings.StructsStaker
		stakerErr                  error
		numStakers                 uint32
		numStakerErr               error
		biggestInfluence           *big.Int
		biggestInfluenceId         uint32
		biggestInfluenceErr        error
		randaoHash                 [32]byte
		randaoHashErr              error
		iteration                  int
		numOfProposedBlocks        uint8
		numOfProposedBlocksErr     error
		maxAltBlocks               uint8
		maxAltBlocksErr            error
		lastIteration              *big.Int
		lastProposedBlockStruct    bindings.StructsBlock
		lastProposedBlockStructErr error
		medians                    []uint32
		mediansErr                 error
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
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
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
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When there is an error in getting staker",
			args: args{
				state:                   2,
				stakerErr:               errors.New("staker error"),
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 4: When there is an error in getting number of stakers",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakerErr:            errors.New("numberOfStakers error"),
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 5: When there is an error in getting biggest influence staker",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluenceErr:     errors.New("biggest influence staker error"),
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("biggest influence staker error"),
		},
		{
			name: "Test 6: When there is an error in getting randaoHash",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHashErr:           errors.New("randao hash error"),
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("randao hash error"),
		},
		{
			name: "Test 7: When iteration is -1",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               -1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When there is an error in getting number of proposed blocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocksErr:  errors.New("numOfProposedBlocks error"),
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("numOfProposedBlocks error"),
		},
		{
			name: "Test 9: When there is an error in getting maxAltBlocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocksErr:         errors.New("maxAltBlocks error"),
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeTxn:              &Types.Transaction{},
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("maxAltBlocks error"),
		},
		{
			name: "Test 10: When numOfProposedBlocks >= maxAltBlocks and there is an error in getting lastProposedBlockStruct",
			args: args{
				state:                      2,
				staker:                     bindings.StructsStaker{},
				numStakers:                 2,
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				randaoHash:                 [32]byte{1, 2, 3},
				iteration:                  1,
				numOfProposedBlocks:        4,
				maxAltBlocks:               2,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStructErr: errors.New("lastProposedBlockStruct error"),
				medians:                    []uint32{1, 2, 3, 4},
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("lastProposedBlockStruct error"),
		},
		{
			name: "Test 11: When numOfProposedBlocks >= maxAltBlocks and current iteration is greater than iteration of last proposed block ",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          2,
				biggestInfluence:    big.NewInt(1000),
				biggestInfluenceId:  2,
				randaoHash:          [32]byte{1, 2, 3},
				iteration:           2,
				numOfProposedBlocks: 4,
				maxAltBlocks:        2,
				lastIteration:       big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(1),
				},
				medians:    []uint32{1, 2, 3, 4},
				txnOpts:    txnOpts,
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 12: When numOfProposedBlocks >= maxAltBlocks and current iteration is less than iteration of last proposed block and propose transaction is successful",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          2,
				biggestInfluence:    big.NewInt(1000),
				biggestInfluenceId:  2,
				randaoHash:          [32]byte{1, 2, 3},
				iteration:           1,
				numOfProposedBlocks: 4,
				maxAltBlocks:        2,
				lastIteration:       big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{
					Iteration: big.NewInt(2),
				},
				medians:    []uint32{1, 2, 3, 4},
				txnOpts:    txnOpts,
				proposeTxn: &Types.Transaction{},
				hash:       common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 13: When there is an error in getting medians",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
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
			name: "Test 14: When Propose transaction fails",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              2,
				biggestInfluence:        big.NewInt(1000),
				biggestInfluenceId:      2,
				randaoHash:              [32]byte{1, 2, 3},
				iteration:               1,
				numOfProposedBlocks:     2,
				maxAltBlocks:            4,
				lastIteration:           big.NewInt(5),
				lastProposedBlockStruct: bindings.StructsBlock{},
				medians:                 []uint32{1, 2, 3, 4},
				txnOpts:                 txnOpts,
				proposeErr:              errors.New("propose error"),
				hash:                    common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("propose error"),
		},
	}
	for _, tt := range tests {
		GetDelayedStateMock = func(*ethclient.Client, int32) (int64, error) {
			return tt.args.state, tt.args.stateErr
		}

		GetStakerMock = func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error) {
			return tt.args.staker, tt.args.stakerErr
		}

		GetNumberOfStakersMock = func(*ethclient.Client, string) (uint32, error) {
			return tt.args.numStakers, tt.args.numStakerErr
		}

		getBiggestInfluenceAndIdMock = func(*ethclient.Client, string, utilsInterface) (*big.Int, uint32, error) {
			return tt.args.biggestInfluence, tt.args.biggestInfluenceId, tt.args.biggestInfluenceErr
		}

		GetRandaoHashMock = func(*ethclient.Client, string) ([32]byte, error) {
			return tt.args.randaoHash, tt.args.randaoHashErr
		}

		getIterationMock = func(*ethclient.Client, string, types.ElectedProposer, proposeUtilsInterface) int {
			return tt.args.iteration
		}

		GetMaxAltBlocksMock = func(*ethclient.Client, string) (uint8, error) {
			return tt.args.maxAltBlocks, tt.args.maxAltBlocksErr
		}

		GetNumberOfProposedBlocksMock = func(*ethclient.Client, string, uint32) (uint8, error) {
			return tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr
		}

		GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint8) (bindings.StructsBlock, error) {
			return tt.args.lastProposedBlockStruct, tt.args.lastProposedBlockStructErr
		}

		MakeBlockMock = func(*ethclient.Client, string, bool, utilsInterface, proposeUtilsInterface) ([]uint32, error) {
			return tt.args.medians, tt.args.mediansErr
		}

		GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
			return tt.args.txnOpts
		}

		ProposeMock = func(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error) {
			return tt.args.proposeTxn, tt.args.proposeErr
		}

		HashMock = func(*Types.Transaction) common.Hash {
			return tt.args.hash
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := Propose(client, account, config, stakerId, epoch, rogueMode, razorUtils, proposeUtils, blockManagerUtils, transactionUtils)
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

func Test_getBiggestInfluenceAndId(t *testing.T) {
	var client *ethclient.Client
	var address string

	razorUtils := UtilsMock{}

	type args struct {
		numOfStakers    uint32
		numOfStakersErr error
		influence       *big.Int
		influenceErr    error
	}
	tests := []struct {
		name          string
		args          args
		wantInfluence *big.Int
		wantId        uint32
		wantErr       error
	}{
		{
			name: "Test 1: When getBiggestInfluenceAndId function executes successfully",
			args: args{
				numOfStakers:    3,
				numOfStakersErr: nil,
				influence:       big.NewInt(1000),
				influenceErr:    nil,
			},
			wantInfluence: big.NewInt(1000),
			wantId:        1,
			wantErr:       nil,
		},
		{
			name: "Test 2: When there is an error in getting numOfStakers",
			args: args{
				numOfStakersErr: errors.New("numOfStakers error"),
				influence:       big.NewInt(1000),
				influenceErr:    nil,
			},
			wantInfluence: nil,
			wantId:        0,
			wantErr:       errors.New("numOfStakers error"),
		},
		{
			name: "Test 3: When there is an error in getting influence",
			args: args{
				numOfStakers:    3,
				numOfStakersErr: nil,
				influenceErr:    errors.New("influence error"),
			},
			wantInfluence: nil,
			wantId:        0,
			wantErr:       errors.New("influence error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumberOfStakersMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.numOfStakers, tt.args.numOfStakersErr
			}

			GetInfluenceMock = func(*ethclient.Client, string, uint32) (*big.Int, error) {
				return tt.args.influence, tt.args.influenceErr
			}

			gotInfluence, gotId, err := getBiggestInfluenceAndId(client, address, razorUtils)
			if gotInfluence.Cmp(tt.wantInfluence) != 0 {
				t.Errorf("Biggest Influence from getBiggestInfluenceAndId function, got = %v, want %v", gotInfluence, tt.wantInfluence)
			}
			if gotId != tt.wantId {
				t.Errorf("Staker Id of staker having biggest Influence from getBiggestInfluenceAndId function, got = %v, want %v", gotId, tt.wantId)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getBiggestInfluenceAndId function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getBiggestInfluenceAndId function, got = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func Test_getIteration(t *testing.T) {
	var client *ethclient.Client
	var address string
	var proposer types.ElectedProposer

	proposeUtils := ProposeUtilsMock{}

	type args struct {
		isElectedProposer bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1: When getIteration returns a valid iteration",
			args: args{
				isElectedProposer: true,
			},
			want: 0,
		},
		//{
		//	name: "Test 2: When getIteration returns an invalid iteration",
		//	args: args{
		//		isElectedProposer: false,
		//	},
		//	want: -1,
		//},
	}
	for _, tt := range tests {
		isElectedProposerMock = func(*ethclient.Client, string, types.ElectedProposer) bool {
			return tt.args.isElectedProposer
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := getIteration(client, address, proposer, proposeUtils); got != tt.want {
				t.Errorf("getIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeBlock(t *testing.T) {

	var client *ethclient.Client
	var address string

	rogueModeMedian := big.NewInt(int64(randMath.Intn(10000000)))

	razorUtils := UtilsMock{}
	proposeUtils := ProposeUtilsMock{}

	type args struct {
		numAssets                 *big.Int
		numAssetsErr              error
		epoch                     uint32
		epochErr                  error
		sortedVotes               []*big.Int
		sortedVotesErr            error
		totalInfluenceRevealed    *big.Int
		totalInfluenceRevealedErr error
		rogueMode                 bool
		influencedMedian          *big.Int
		mediansInUint32           []uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []uint32
		wantErr error
	}{
		{
			name: "Test 1: When rogueMode is true and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1000)},
				totalInfluenceRevealed: big.NewInt(2000),
				rogueMode:              true,
				mediansInUint32:        []uint32{uint32(rogueModeMedian.Int64())},
			},
			want:    []uint32{uint32(rogueModeMedian.Int64())},
			wantErr: nil,
		},
		{
			name: "Test 2: When rogueMode is false and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(2000), big.NewInt(4000)},
				totalInfluenceRevealed: big.NewInt(1000),
				rogueMode:              false,
				influencedMedian:       big.NewInt(6),
				mediansInUint32:        []uint32{uint32(big.NewInt(6).Int64())},
			},
			want:    []uint32{6},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting number of assets",
			args: args{
				numAssetsErr:           errors.New("numAssets error"),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(2000)},
				totalInfluenceRevealed: big.NewInt(1000),
				rogueMode:              false,
				influencedMedian:       big.NewInt(2),
				mediansInUint32:        []uint32{uint32(big.NewInt(2).Int64())},
			},
			want:    nil,
			wantErr: errors.New("numAssets error"),
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				numAssets:              big.NewInt(1),
				epochErr:               errors.New("epoch error"),
				sortedVotes:            []*big.Int{big.NewInt(2000)},
				totalInfluenceRevealed: big.NewInt(1000),
				rogueMode:              false,
				influencedMedian:       big.NewInt(2),
				mediansInUint32:        []uint32{uint32(big.NewInt(2).Int64())},
			},
			want:    nil,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When there is an error in getting sorted votes",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotesErr:         errors.New("sorted votes error"),
				totalInfluenceRevealed: big.NewInt(1000),
				mediansInUint32:        nil,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Test 5: When there is an error in getting totalInfluenceRevealed",
			args: args{
				numAssets:                 big.NewInt(1),
				epoch:                     4,
				sortedVotes:               []*big.Int{big.NewInt(2000)},
				totalInfluenceRevealedErr: errors.New("totalInfluenceRevealed error"),
				mediansInUint32:           nil,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Test 6: When number of assets is greater than 1 and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(2),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(2000), big.NewInt(4000)},
				totalInfluenceRevealed: big.NewInt(1000),
				rogueMode:              false,
				influencedMedian:       big.NewInt(6),
				mediansInUint32:        []uint32{uint32(big.NewInt(6).Int64()), uint32(big.NewInt(6).Int64())},
			},
			want:    []uint32{6, 6},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumActiveAssetsMock = func(*ethclient.Client, string) (*big.Int, error) {
				return tt.args.numAssets, tt.args.numAssetsErr
			}

			GetEpochMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			getSortedVotesMock = func(*ethclient.Client, string, uint8, uint32, utilsInterface) ([]*big.Int, error) {
				return tt.args.sortedVotes, tt.args.sortedVotesErr
			}

			GetTotalInfluenceRevealedMock = func(*ethclient.Client, string, uint32) (*big.Int, error) {
				return tt.args.totalInfluenceRevealed, tt.args.totalInfluenceRevealedErr
			}

			influencedMedianMock = func([]*big.Int, *big.Int) *big.Int {
				return tt.args.influencedMedian
			}

			ConvertBigIntArrayToUint32ArrayMock = func([]*big.Int) []uint32 {
				return tt.args.mediansInUint32
			}

			got, err := MakeBlock(client, address, tt.args.rogueMode, razorUtils, proposeUtils)
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data from MakeBlock function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error from MakeBlock function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error from MakeBlock function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func Test_getSortedVotes(t *testing.T) {

	var client *ethclient.Client
	var address string
	var assetId uint8

	razorUtils := UtilsMock{}

	type args struct {
		numberOfStakers      uint32
		numberOfStakersErr   error
		epoch                uint32
		epochLastRevealed    uint32
		epochLastRevealedErr error
		vote                 *big.Int
		voteErr              error
		influence            *big.Int
		influenceErr         error
	}
	tests := []struct {
		name    string
		args    args
		want    []*big.Int
		wantErr error
	}{
		{
			name: "Test 1: When getSortedVotes function executes successfully",
			args: args{
				numberOfStakers:   2,
				epoch:             4,
				epochLastRevealed: 4,
				vote:              big.NewInt(1000),
				influence:         big.NewInt(2000),
			},
			want:    []*big.Int{big.NewInt(2000000), big.NewInt(2000000)},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting numberOfStakers",
			args: args{
				numberOfStakersErr: errors.New("numberOfStakers error"),
				epoch:              4,
				epochLastRevealed:  4,
				vote:               big.NewInt(1000),
				influence:          big.NewInt(2000),
			},
			want:    nil,
			wantErr: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 3: When there is an error in getting epochLastRevealed",
			args: args{
				numberOfStakers:      2,
				epoch:                4,
				epochLastRevealedErr: errors.New("epochLastRevealed error"),
				vote:                 big.NewInt(1000),
				influence:            big.NewInt(2000),
			},
			want:    nil,
			wantErr: errors.New("epochLastRevealed error"),
		},
		{
			name: "Test 4: When there is an error in getting vote value",
			args: args{
				numberOfStakers:   2,
				epoch:             4,
				epochLastRevealed: 4,
				voteErr:           errors.New("vote error"),
				influence:         big.NewInt(2000),
			},
			want:    nil,
			wantErr: errors.New("vote error"),
		},
		{
			name: "Test 5: When there is an error in getting influence",
			args: args{
				numberOfStakers:   2,
				epoch:             4,
				epochLastRevealed: 4,
				vote:              big.NewInt(1000),
				influenceErr:      errors.New("influence error"),
			},
			want:    nil,
			wantErr: errors.New("influence error"),
		},
		{
			name: "Test 6: When epoch != epochLastRevealed",
			args: args{
				numberOfStakers:   2,
				epoch:             5,
				epochLastRevealed: 4,
				vote:              big.NewInt(1000),
				influence:         big.NewInt(2000),
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Test 7: When numberOfStakers is 0",
			args: args{
				numberOfStakers:   0,
				epoch:             4,
				epochLastRevealed: 4,
				vote:              big.NewInt(1000),
				influence:         big.NewInt(2000),
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumberOfStakersMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.numberOfStakers, tt.args.numberOfStakersErr
			}

			GetEpochLastRevealedMock = func(*ethclient.Client, string, uint32) (uint32, error) {
				return tt.args.epochLastRevealed, tt.args.epochLastRevealedErr
			}

			GetVoteValueMock = func(*ethclient.Client, string, uint8, uint32) (*big.Int, error) {
				return tt.args.vote, tt.args.voteErr
			}

			GetInfluenceSnapshotMock = func(*ethclient.Client, string, uint32, uint32) (*big.Int, error) {
				return tt.args.influence, tt.args.influenceErr
			}

			got, err := getSortedVotes(client, address, assetId, tt.args.epoch, razorUtils)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data from getSortedVotes function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error from getSortedVotes function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error from getSortedVotes function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_influencedMedian(t *testing.T) {
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
				totalInfluenceRevealed: big.NewInt(100000000),
			},
			want: big.NewInt(0),
		},
		{
			name: "Test if totalInfluenceRevealed is 0",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(100)},
				totalInfluenceRevealed: big.NewInt(0),
			},
			want: big.NewInt(100),
		},
		{
			name: "Test if all the values are present",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(100), big.NewInt(110), big.NewInt(115), big.NewInt(118)},
				totalInfluenceRevealed: big.NewInt(400),
			},
			want: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := influencedMedian(tt.args.sortedVotes, tt.args.totalInfluenceRevealed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("influencedMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
