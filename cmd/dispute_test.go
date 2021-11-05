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
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"
)

func TestDispute(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var blockManager *bindings.BlockManager
	var client *ethclient.Client
	var config types.Configurations
	var account types.Account
	var blockId uint8
	var assetId int

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		cmdUtils:          UtilsCmdMock{},
		blockManagerUtils: BlockManagerMock{},
		transactionUtils:  TransactionMock{},
	}

	type args struct {
		epoch              uint32
		numOfStakers       uint32
		numOfStakersErr    error
		votes              bindings.StructsVote
		votesErr           error
		containsStatus     bool
		finalizeDisputeTxn *Types.Transaction
		finalizeDisputeErr error
		hash               common.Hash
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When Dispute function executes successfully",
			args: args{
				epoch:        4,
				numOfStakers: 3,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				containsStatus:     false,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: nil,
		},
		{
			name: "Test 2: When Dispute function executes successfully without executing giveSorted",
			args: args{
				epoch:        4,
				numOfStakers: 3,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				containsStatus:     true,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: nil,
		},
		{
			name: "Test 3: When there is an error in getting number of stakers",
			args: args{
				epoch:           4,
				numOfStakersErr: errors.New("numberOfStakers error"),
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				containsStatus:     false,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: errors.New("numberOfStakers error"),
		},
		{
			name: "Test 4: When there is an error in getting votes",
			args: args{
				epoch:              4,
				numOfStakers:       3,
				votesErr:           errors.New("votes error"),
				containsStatus:     false,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: errors.New("votes error"),
		},
		{
			name: "Test 5: When FinalizeDispute transaction fails",
			args: args{
				epoch:        4,
				numOfStakers: 3,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				containsStatus:     false,
				finalizeDisputeErr: errors.New("finalizeDispute error"),
			},
			want: errors.New("finalizeDispute error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetBlockManagerMock = func(*ethclient.Client) *bindings.BlockManager {
				return blockManager
			}

			GetNumberOfStakersMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.numOfStakers, tt.args.numOfStakersErr
			}

			GetVotesMock = func(*ethclient.Client, string, uint32) (bindings.StructsVote, error) {
				return tt.args.votes, tt.args.votesErr
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return txnOpts
			}

			ContainsMock = func([]int, int) bool {
				return tt.args.containsStatus
			}

			GiveSortedMock = func(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {

			}

			FinalizeDisputeMock = func(*ethclient.Client, *bind.TransactOpts, uint32, uint8) (*Types.Transaction, error) {
				return tt.args.finalizeDisputeTxn, tt.args.finalizeDisputeErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			WaitForBlockCompletionMock = func(*ethclient.Client, string) int {
				return 1
			}

			err := Dispute(client, config, account, tt.args.epoch, blockId, assetId, utilsStruct)
			if err == nil || tt.want == nil {
				if err != tt.want {
					t.Errorf("Error for Dispute function, got = %v, want = %v", err, tt.want)
				}
			} else {
				if err.Error() != tt.want.Error() {
					t.Errorf("Error for Dispute function, got = %v, want = %v", err, tt.want)
				}
			}
		})
	}
}

func TestHandleDispute(t *testing.T) {
	var client *ethclient.Client
	var config types.Configurations
	var account types.Account
	var epoch uint32

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		proposeUtils:      ProposeUtilsMock{},
		cmdUtils:          UtilsCmdMock{},
		blockManagerUtils: BlockManagerMock{},
		transactionUtils:  TransactionMock{},
	}

	type args struct {
		numberOfProposedBlocks    uint8
		numberOfProposedBlocksErr error
		proposedBlock             bindings.StructsBlock
		proposedBlockErr          error
		medians                   []uint32
		mediansErr                error
		activeAssetIds            []uint8
		activeAssetIdsErr         error
		isEqual                   bool
		iteration                 int
		disputeErr                error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When HandleDispute function executes successfully when there is a dispute case",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
					Valid:   true,
				},
				medians:        []uint32{101, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        false,
				iteration:      0,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 2: When HandleDispute function executes successfully when there is no dispute case",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{100, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 3: When there is an error in getting numberOfProposedBlocks",
			args: args{
				numberOfProposedBlocksErr: errors.New("numberOfProposedBlocks error"),
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{100, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: errors.New("numberOfProposedBlocks error"),
		},
		{
			name: "Test 4: When there is an error in getting proposedBlock",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlockErr:       errors.New("proposedBlock error"),
				medians:                []uint32{100, 200, 300},
				activeAssetIds:         []uint8{3, 4, 6},
				isEqual:                true,
				disputeErr:             nil,
			},
			want: nil,
		},
		{
			name: "Test 5: When there is an error in getting medians from MakeBlock ",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				mediansErr:     errors.New("medians error"),
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 6: When there is an error from Dispute function",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
					Valid:   true,
				},
				medians:        []uint32{101, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        false,
				iteration:      0,
				disputeErr:     errors.New("dispute error"),
			},
			want: nil,
		},
		{
			name: "Test 6: When there is a case of Dispute but block is already disputed",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{101, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        false,
				iteration:      0,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumberOfProposedBlocksMock = func(*ethclient.Client, string, uint32) (uint8, error) {
				return tt.args.numberOfProposedBlocks, tt.args.numberOfProposedBlocksErr
			}

			GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint8) (bindings.StructsBlock, error) {
				return tt.args.proposedBlock, tt.args.proposedBlockErr
			}

			MakeBlockMock = func(*ethclient.Client, string, bool, utilsInterface, proposeUtilsInterface) ([]uint32, error) {
				return tt.args.medians, tt.args.mediansErr
			}

			GetActiveAssetIdsMock = func(*ethclient.Client, string, uint32) ([]uint8, error) {
				return tt.args.activeAssetIds, tt.args.activeAssetIdsErr
			}

			IsEqualMock = func([]uint32, []uint32) (bool, int) {
				return tt.args.isEqual, tt.args.iteration
			}

			DisputeMock = func(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int, UtilsStruct) error {
				return tt.args.disputeErr
			}

			err := HandleDispute(client, config, account, epoch, utilsStruct)
			if err == nil || tt.want == nil {
				if err != tt.want {
					t.Errorf("Error for HandleDispute function, got = %v, want = %v", err, tt.want)
				}
			} else {
				if err.Error() != tt.want.Error() {
					t.Errorf("Error for HandleDispute function, got = %v, want = %v", err, tt.want)
				}
			}
		})
	}
}
