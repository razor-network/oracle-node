package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
		lastProposedBlockStruct    types.Block
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
				state:                      2,
				stateErr:                   nil,
				staker:                     bindings.StructsStaker{},
				stakerErr:                  nil,
				numStakers:                 2,
				numStakerErr:               nil,
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				biggestInfluenceErr:        nil,
				randaoHash:                 [32]byte{1, 2, 3},
				randaoHashErr:              nil,
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr:                   errors.New("state error"),
				staker:                     bindings.StructsStaker{},
				stakerErr:                  nil,
				numStakers:                 2,
				numStakerErr:               nil,
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				biggestInfluenceErr:        nil,
				randaoHash:                 [32]byte{1, 2, 3},
				randaoHashErr:              nil,
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When there is an error in getting staker",
			args: args{
				state:                      2,
				stateErr:                   nil,
				stakerErr:                  errors.New("staker error"),
				numStakers:                 2,
				numStakerErr:               nil,
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				biggestInfluenceErr:        nil,
				randaoHash:                 [32]byte{1, 2, 3},
				randaoHashErr:              nil,
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("staker error"),
		},
		{
			name: "Test 4: When there is an error in getting number of stakers",
			args: args{
				state:                      2,
				stateErr:                   nil,
				staker:                     bindings.StructsStaker{},
				stakerErr:                  nil,
				numStakerErr:               errors.New("numberOfStakers error"),
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				biggestInfluenceErr:        nil,
				randaoHash:                 [32]byte{1, 2, 3},
				randaoHashErr:              nil,
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 5: When there is an error in getting biggest influence staker",
			args: args{
				state:                      2,
				stateErr:                   nil,
				staker:                     bindings.StructsStaker{},
				stakerErr:                  nil,
				numStakers:                 2,
				numStakerErr:               nil,
				biggestInfluenceErr:        errors.New("biggest influence staker error"),
				randaoHash:                 [32]byte{1, 2, 3},
				randaoHashErr:              nil,
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("biggest influence staker error"),
		},
		{
			name: "Test 6: When there is an error in getting randaoHash",
			args: args{
				state:                      2,
				stateErr:                   nil,
				staker:                     bindings.StructsStaker{},
				stakerErr:                  nil,
				numStakers:                 2,
				numStakerErr:               nil,
				biggestInfluence:           big.NewInt(1000),
				biggestInfluenceId:         2,
				biggestInfluenceErr:        nil,
				randaoHashErr:              errors.New("randao hash error"),
				iteration:                  1,
				numOfProposedBlocks:        2,
				numOfProposedBlocksErr:     nil,
				maxAltBlocks:               4,
				lastIteration:              big.NewInt(5),
				lastProposedBlockStruct:    types.Block{},
				lastProposedBlockStructErr: nil,
				medians:                    []uint32{1, 2, 3, 4},
				mediansErr:                 nil,
				txnOpts:                    txnOpts,
				proposeTxn:                 &Types.Transaction{},
				proposeErr:                 nil,
				hash:                       common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("randao hash error"),
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

		getBiggestInfluenceAndIdMock = func(*ethclient.Client, string) (*big.Int, uint32, error) {
			return tt.args.biggestInfluence, tt.args.biggestInfluenceId, tt.args.biggestInfluenceErr
		}

		GetRandaoHashMock = func(*ethclient.Client, string) ([32]byte, error) {
			return tt.args.randaoHash, tt.args.randaoHashErr
		}

		getIterationMock = func(*ethclient.Client, string, types.ElectedProposer) int {
			return tt.args.iteration
		}

		GetMaxAltBlocksMock = func(*ethclient.Client, string) (uint8, error) {
			return tt.args.maxAltBlocks, tt.args.maxAltBlocksErr
		}

		GetNumberOfProposedBlocksMock = func(*ethclient.Client, string, uint32) (uint8, error) {
			return tt.args.numOfProposedBlocks, tt.args.numOfProposedBlocksErr
		}

		GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint8) (types.Block, error) {
			return tt.args.lastProposedBlockStruct, tt.args.lastProposedBlockStructErr
		}

		MakeBlockMock = func(*ethclient.Client, string, bool) ([]uint32, error) {
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
