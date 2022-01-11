package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
	randMath "math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestPropose(t *testing.T) {

	var (
		client   *ethclient.Client
		account  types.Account
		config   types.Configurations
		stakerId uint32
		epoch    uint32
		rogue    types.Rogue
	)

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	randaoHash := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	randaoHashBytes32 := [32]byte{}
	copy(randaoHashBytes32[:], randaoHash)

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		proposeUtils:      ProposeUtilsMock{},
		blockManagerUtils: BlockManagerMock{},
		transactionUtils:  TransactionMock{},
	}

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
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			name: "Test 3: When there is an error in getting staker",
			args: args{
				state:                   2,
				stakerErr:               errors.New("staker error"),
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 4: When there is an error in getting number of stakers",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakerErr:            errors.New("numberOfStakers error"),
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			name: "Test 5: When there is an error in getting biggest influence staker",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluenceErr:     errors.New("biggest influence staker error"),
				randaoHash:              randaoHashBytes32,
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
			wantErr: errors.New("biggest influence staker error"),
		},
		{
			name: "Test 6: When there is an error in getting randaoHash",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHashErr:           errors.New("randao hash error"),
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
			wantErr: errors.New("randao hash error"),
		},
		{
			name: "Test 7: When iteration is -1",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			name: "Test 8: When there is an error in getting number of proposed blocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			name: "Test 9: When there is an error in getting maxAltBlocks",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
			name: "Test 10: When numOfProposedBlocks >= maxAltBlocks and there is an error in getting lastProposedBlockStruct",
			args: args{
				state:                      2,
				staker:                     bindings.StructsStaker{},
				numStakers:                 5,
				biggestInfluence:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:         2,
				randaoHash:                 randaoHashBytes32,
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
			name: "Test 11: When numOfProposedBlocks >= maxAltBlocks and current iteration is greater than iteration of last proposed block ",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          5,
				biggestInfluence:    big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:  2,
				randaoHash:          randaoHashBytes32,
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
			name: "Test 12: When numOfProposedBlocks >= maxAltBlocks and current iteration is less than iteration of last proposed block and propose transaction is successful",
			args: args{
				state:               2,
				staker:              bindings.StructsStaker{},
				numStakers:          5,
				biggestInfluence:    big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:  2,
				randaoHash:          randaoHashBytes32,
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
			name: "Test 13: When there is an error in getting medians",
			args: args{
				state:                   2,
				staker:                  bindings.StructsStaker{},
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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
				numStakers:              5,
				biggestInfluence:        big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestInfluenceId:      2,
				randaoHash:              randaoHashBytes32,
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

		getBiggestStakeAndIdMock = func(*ethclient.Client, string, uint32, UtilsStruct) (*big.Int, uint32, error) {
			return tt.args.biggestInfluence, tt.args.biggestInfluenceId, tt.args.biggestInfluenceErr
		}

		GetRandaoHashMock = func(*ethclient.Client) ([32]byte, error) {
			return tt.args.randaoHash, tt.args.randaoHashErr
		}

		getIterationMock = func(*ethclient.Client, types.ElectedProposer, UtilsStruct) int {
			return tt.args.iteration
		}

		GetMaxAltBlocksMock = func(*ethclient.Client, string) (uint8, error) {
			return tt.args.maxAltBlocks, tt.args.maxAltBlocksErr
		}

		GetNumberOfProposedBlocksMock = func(*ethclient.Client, string, uint32) (uint8, error) {
			return tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr
		}

		GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint32) (bindings.StructsBlock, error) {
			return tt.args.lastProposedBlockStruct, tt.args.lastProposedBlockStructErr
		}

		MakeBlockMock = func(*ethclient.Client, string, types.Rogue, UtilsStruct) ([]uint32, error) {
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
			got, err := utilsStruct.Propose(client, account, config, stakerId, epoch, rogue)
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
	var epoch uint32

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
	}

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
			name: "Test 1: When getBiggestStakeAndId function executes successfully",
			args: args{
				numOfStakers:    5,
				numOfStakersErr: nil,
				influence:       big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				influenceErr:    nil,
			},
			wantInfluence: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
			wantId:        1,
			wantErr:       nil,
		},
		{
			name: "Test 2: When there is an error in getting numOfStakers",
			args: args{
				numOfStakersErr: errors.New("numOfStakers error"),
				influence:       big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				influenceErr:    nil,
			},
			wantInfluence: nil,
			wantId:        0,
			wantErr:       errors.New("numOfStakers error"),
		},
		{
			name: "Test 3: When there is an error in getting influence",
			args: args{
				numOfStakers:    5,
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

			GetInfluenceSnapshotMock = func(*ethclient.Client, uint32, uint32) (*big.Int, error) {
				return tt.args.influence, tt.args.influenceErr
			}

			gotStake, gotId, err := getBiggestStakeAndId(client, address, epoch, utilsStruct)
			if gotStake.Cmp(tt.wantInfluence) != 0 {
				t.Errorf("Biggest Influence from getBiggestStakeAndId function, got = %v, want %v", gotStake, tt.wantInfluence)
			}
			if gotId != tt.wantId {
				t.Errorf("Staker Id of staker having biggest Influence from getBiggestStakeAndId function, got = %v, want %v", gotId, tt.wantId)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for getBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for getBiggestStakeAndId function, got = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func Test_getIteration(t *testing.T) {
	var client *ethclient.Client
	var proposer types.ElectedProposer

	utilsStruct := UtilsStruct{
		proposeUtils: ProposeUtilsMock{},
		razorUtils:   UtilsMock{},
	}

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
		{
			name: "Test 2: When getIteration returns an invalid iteration",
			args: args{
				isElectedProposer: false,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		isElectedProposerMock = func(*ethclient.Client, types.ElectedProposer, UtilsStruct) bool {
			return tt.args.isElectedProposer
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := getIteration(client, proposer, utilsStruct); got != tt.want {
				t.Errorf("getIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeBlock(t *testing.T) {

	var client *ethclient.Client
	var address string

	rogueMedian := big.NewInt(int64(randMath.Intn(10000000)))

	utilsStruct := UtilsStruct{
		razorUtils:   UtilsMock{},
		proposeUtils: ProposeUtilsMock{},
	}

	type args struct {
		numAssets                 *big.Int
		numAssetsErr              error
		epoch                     uint32
		epochErr                  error
		sortedVotes               []*big.Int
		sortedVotesErr            error
		totalInfluenceRevealed    *big.Int
		totalInfluenceRevealedErr error
		rogue                     types.Rogue
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
			name: "Test 1: When rogue is true and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
				rogue: types.Rogue{
					IsRogue:   true,
					RogueMode: nil,
				},
				mediansInUint32: []uint32{uint32(rogueMedian.Int64())},
			},
			want:    []uint32{uint32(rogueMedian.Int64())},
			wantErr: nil,
		},
		{
			name: "Test 2: When rogue is false and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
				rogue:                  types.Rogue{IsRogue: false},
				influencedMedian:       big.NewInt(498342),
				mediansInUint32:        []uint32{uint32(big.NewInt(498342).Int64())},
			},
			want:    []uint32{498342},
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting number of assets",
			args: args{
				numAssetsErr:           errors.New("numAssets error"),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
				rogue:                  types.Rogue{IsRogue: false},
				influencedMedian:       big.NewInt(498342),
				mediansInUint32:        []uint32{uint32(big.NewInt(498342).Int64()), uint32(big.NewInt(498342).Int64())},
			},
			want:    nil,
			wantErr: errors.New("numAssets error"),
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				numAssets:              big.NewInt(1),
				epochErr:               errors.New("epoch error"),
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
				rogue:                  types.Rogue{IsRogue: false},
				influencedMedian:       big.NewInt(498342),
				mediansInUint32:        []uint32{uint32(big.NewInt(498342).Int64()), uint32(big.NewInt(498342).Int64())},
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
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
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
				sortedVotes:               []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealedErr: errors.New("totalInfluenceRevealed error"),
				mediansInUint32:           nil,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Test 6: When number of assets are more than 1 and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(3),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697718000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(4200), big.NewInt(1e18)),
				rogue:                  types.Rogue{IsRogue: false},
				influencedMedian:       big.NewInt(498342),
				mediansInUint32:        []uint32{uint32(big.NewInt(498342).Int64()), uint32(big.NewInt(498342).Int64()), uint32(big.NewInt(498342).Int64())},
			},
			want:    []uint32{498342, 498342, 498342},
			wantErr: nil,
		},
		{
			name: "Test 7: When rogue is true in propose mode and MakeBlock function executes successfully",
			args: args{
				numAssets:              big.NewInt(1),
				epoch:                  4,
				sortedVotes:            []*big.Int{big.NewInt(1).Mul(big.NewInt(697690000), big.NewInt(1e18))},
				totalInfluenceRevealed: big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
				rogue: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"propose"},
				},
				mediansInUint32: []uint32{uint32(rogueMedian.Int64())},
			},
			want:    []uint32{uint32(rogueMedian.Int64())},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumActiveAssetsMock = func(*ethclient.Client) (*big.Int, error) {
				return tt.args.numAssets, tt.args.numAssetsErr
			}

			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			getSortedVotesMock = func(*ethclient.Client, string, uint16, uint32, UtilsStruct) ([]*big.Int, error) {
				return tt.args.sortedVotes, tt.args.sortedVotesErr
			}

			GetTotalInfluenceRevealedMock = func(*ethclient.Client, uint32) (*big.Int, error) {
				return tt.args.totalInfluenceRevealed, tt.args.totalInfluenceRevealedErr
			}

			influencedMedianMock = func([]*big.Int, *big.Int) *big.Int {
				return tt.args.influencedMedian
			}

			ConvertBigIntArrayToUint32ArrayMock = func([]*big.Int) []uint32 {
				return tt.args.mediansInUint32
			}

			got, err := MakeBlock(client, address, tt.args.rogue, utilsStruct)
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
	var assetId uint16

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
	}

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
				numberOfStakers:   3,
				epoch:             4,
				epochLastRevealed: 4,
				vote:              big.NewInt(498307),
				influence:         big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
			},
			want:    []*big.Int{big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18)), big.NewInt(1).Mul(big.NewInt(697629800), big.NewInt(1e18))},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting numberOfStakers",
			args: args{
				numberOfStakersErr: errors.New("numberOfStakers error"),
				epoch:              4,
				epochLastRevealed:  4,
				vote:               big.NewInt(498307),
				influence:          big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
			},
			want:    nil,
			wantErr: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 3: When there is an error in getting epochLastRevealed",
			args: args{
				numberOfStakers:      3,
				epoch:                4,
				epochLastRevealedErr: errors.New("epochLastRevealed error"),
				vote:                 big.NewInt(498307),
				influence:            big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
			},
			want:    nil,
			wantErr: errors.New("epochLastRevealed error"),
		},
		{
			name: "Test 4: When there is an error in getting vote value",
			args: args{
				numberOfStakers:   3,
				epoch:             4,
				epochLastRevealed: 4,
				voteErr:           errors.New("vote error"),
				influence:         big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
			},
			want:    nil,
			wantErr: errors.New("vote error"),
		},
		{
			name: "Test 5: When there is an error in getting influence",
			args: args{
				numberOfStakers:   3,
				epoch:             4,
				epochLastRevealed: 4,
				vote:              big.NewInt(498307),
				influenceErr:      errors.New("influence error"),
			},
			want:    nil,
			wantErr: errors.New("influence error"),
		},
		{
			name: "Test 6: When epoch != epochLastRevealed",
			args: args{
				numberOfStakers:   3,
				epoch:             5,
				epochLastRevealed: 4,
				vote:              big.NewInt(498307),
				influence:         big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
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
				vote:              big.NewInt(498307),
				influence:         big.NewInt(1).Mul(big.NewInt(1400), big.NewInt(1e18)),
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

			GetVoteValueMock = func(*ethclient.Client, uint16, uint32) (*big.Int, error) {
				return tt.args.vote, tt.args.voteErr
			}

			GetInfluenceSnapshotMock = func(*ethclient.Client, uint32, uint32) (*big.Int, error) {
				return tt.args.influence, tt.args.influenceErr
			}

			got, err := getSortedVotes(client, address, assetId, tt.args.epoch, utilsStruct)
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
			if got := influencedMedian(tt.args.sortedVotes, tt.args.totalInfluenceRevealed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("influencedMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func influenceSnapshotValue(infl string) *big.Int {
	influence, _ := new(big.Int).SetString(infl, 10)
	return influence
}

func Test_isElectedProposer(t *testing.T) {
	var client *ethclient.Client

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
	}

	randaoHash := []byte{142, 170, 157, 83, 109, 43, 34, 152, 21, 154, 159, 12, 195, 119, 50, 186, 218, 57, 39, 173, 228, 135, 20, 100, 149, 27, 169, 158, 34, 113, 66, 64}
	randaoHashBytes32 := [32]byte{}
	copy(randaoHashBytes32[:], randaoHash)

	biggestInfluence, _ := new(big.Int).SetString("2592145500000000000000000", 10)

	type args struct {
		client               *ethclient.Client
		address              string
		proposer             types.ElectedProposer
		influenceSnapshot    *big.Int
		influenceSnapshotErr error
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1: When staker is 3 and isElectedProposer returns true",
			args: args{
				client:  client,
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       0,
					Stake:           nil,
					StakerId:        3,
					BiggestStake:    biggestInfluence,
					NumberOfStakers: 3,
					RandaoHash:      randaoHashBytes32,
					Epoch:           333,
				},
				influenceSnapshot:    influenceSnapshotValue("2592145500000000000000000"),
				influenceSnapshotErr: nil,
			},
			want: true,
		},
		{
			name: "Test2: When staker is 2 and isElectedProposer returns false",
			args: args{
				client:  client,
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       11,
					Stake:           nil,
					StakerId:        2,
					BiggestStake:    biggestInfluence,
					NumberOfStakers: 3,
					RandaoHash:      randaoHashBytes32,
					Epoch:           29,
				},
				influenceSnapshot:    influenceSnapshotValue("529422500000000000000000"),
				influenceSnapshotErr: nil,
			},
			want: false,
		},
		{
			name: "When staker is 1 and isElectedProposer returns true",
			args: args{
				client:  client,
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       2,
					Stake:           nil,
					StakerId:        1,
					BiggestStake:    biggestInfluence,
					NumberOfStakers: 3,
					RandaoHash:      randaoHashBytes32,
					Epoch:           333,
				},
				influenceSnapshot:    influenceSnapshotValue("2592145500000000000000000"),
				influenceSnapshotErr: nil,
			},
			want: true,
		},
		{
			name: "Test4: When there is an error getting influence snapshot",
			args: args{
				client:  client,
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       0,
					Stake:           nil,
					StakerId:        3,
					BiggestStake:    biggestInfluence,
					NumberOfStakers: 3,
					RandaoHash:      randaoHashBytes32,
					Epoch:           333,
				},
				influenceSnapshot:    nil,
				influenceSnapshotErr: errors.New("error in getting influence snapshot"),
			},
			want: false,
		},
		{
			name: "Test5: When pseudoRandomNumber is not equal to proposer's stakerID",
			args: args{
				client:  client,
				address: "0x000000000000000000000000000000000000dead",
				proposer: types.ElectedProposer{
					Iteration:       0,
					Stake:           nil,
					StakerId:        3,
					BiggestStake:    biggestInfluence,
					NumberOfStakers: 3,
					RandaoHash:      [32]byte{},
					Epoch:           333,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {

		GetInfluenceSnapshotMock = func(*ethclient.Client, uint32, uint32) (*big.Int, error) {
			return tt.args.influenceSnapshot, tt.args.influenceSnapshotErr
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := isElectedProposer(tt.args.client, tt.args.proposer, utilsStruct); got != tt.want {
				t.Errorf("isElectedProposer() = %v, want %v", got, tt.want)
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
