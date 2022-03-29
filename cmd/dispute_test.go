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
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	mocks2 "razor/utils/mocks"
	"testing"
)

func TestDispute(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var (
		client        *ethclient.Client
		config        types.Configurations
		account       types.Account
		epoch         uint32
		blockIndex    uint8
		proposedBlock bindings.StructsBlock
		leafId        uint16
		sortedValues  []uint32
		blockManager  *bindings.BlockManager
	)

	type args struct {
		containsStatus              bool
		positionOfCollectionInBlock *big.Int
		finalizeDisputeTxn          *Types.Transaction
		finalizeDisputeErr          error
		hash                        common.Hash
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When Dispute function executes successfully",
			args: args{
				containsStatus:     false,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: nil,
		},
		{
			name: "Test 2: When Dispute function executes successfully without executing giveSorted",
			args: args{
				containsStatus:     true,
				finalizeDisputeTxn: &Types.Transaction{},
				hash:               common.BigToHash(big.NewInt(1)),
			},
			want: nil,
		},
		{
			name: "Test 5: When FinalizeDispute transaction fails",
			args: args{
				containsStatus:     false,
				finalizeDisputeErr: errors.New("finalizeDispute error"),
			},
			want: errors.New("finalizeDispute error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			blockManagerUtilsMock := new(mocks.BlockManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			blockManagerUtils = blockManagerUtilsMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetBlockManager", mock.AnythingOfType("*ethclient.Client")).Return(blockManager)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("GiveSorted", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
			cmdUtilsMock.On("GetCollectionIdPositionInBlock", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.positionOfCollectionInBlock)
			blockManagerUtilsMock.On("FinalizeDispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.finalizeDisputeTxn, tt.args.finalizeDisputeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStruct{}

			err := utils.Dispute(client, config, account, epoch, blockIndex, proposedBlock, leafId, sortedValues)
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
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var client *ethclient.Client
	var config types.Configurations
	var account types.Account
	var epoch uint32
	var blockNumber *big.Int
	var rogueData types.Rogue

	type args struct {
		sortedProposedBlockIds    []uint32
		sortedProposedBlockIdsErr error
		biggestStake              *big.Int
		biggestStakeId            uint32
		biggestStakeErr           error
		medians                   []uint32
		revealedCollectionIds     []uint16
		revealedDataMaps          *types.RevealedDataMaps
		medainsErr                error
		proposedBlock             bindings.StructsBlock
		proposedBlockErr          error
		disputeBiggestStakeTxn    *Types.Transaction
		disputeBiggestStakeErr    error
		Hash                      common.Hash
		idDisputetxn              *Types.Transaction
		idDisputeTxnErr           error
		isEqual                   bool
		leafId                    uint16
		leafIdErr                 error
		disputeErr                error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When HandleDispute function executes successfully when there is a medians dispute case",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				medians:                []uint32{6901548, 498307},
				revealedCollectionIds:  []uint16{1},
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: nil,
					VoteWeights:          nil,
					InfluenceSum:         nil,
				},
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				isEqual:    false,
				disputeErr: nil,
			},
			want: nil,
		},
		{
			name: "Test 2: When HandleDispute function executes successfully when there is no dispute case",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				isEqual:    true,
				disputeErr: nil,
			},
			want: nil,
		},
		{
			name: "Test 3: When there is an error in getting sortedProposedBlockIds",
			args: args{
				sortedProposedBlockIdsErr: errors.New("sortedProposedBlockIds error"),
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{6701548, 478307},
				},
				isEqual:    true,
				disputeErr: nil,
			},
			want: errors.New("sortedProposedBlockIds error"),
		},
		{
			name: "Test 4: When there is an error in getting proposedBlock",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				proposedBlockErr:       errors.New("proposedBlock error"),
				isEqual:                true,
				disputeErr:             nil,
			},
			want: nil,
		},
		//{
		//	name: "Test 7: When there is an error from Dispute function",
		//	args: args{
		//		sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
		//		biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
		//		biggestStakeId:         2,
		//		proposedBlock: bindings.StructsBlock{
		//			Medians:      []uint32{6901548, 498307},
		//			BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
		//			Valid:        true,
		//		},
		//		isEqual:    false,
		//		disputeErr: errors.New("dispute error"),
		//	},
		//	want: nil,
		//},
		{
			name: "Test 8: When there is a case of Dispute but block is already disputed",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				isEqual: false,
			},
			want: nil,
		},
		{
			name: "Test 9: When HandleDispute function executes successfully when there is a biggest influence dispute case",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(4356), big.NewInt(1e18)),
				},
				disputeBiggestStakeTxn: &Types.Transaction{},
				Hash:                   common.BigToHash(big.NewInt(1)),
				isEqual:                false,
				disputeErr:             nil,
			},
			want: nil,
		},
		{
			name: "Test 10: When there is an error in getting biggestInfluenceAndId",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStakeErr:        errors.New("biggestInfluenceAndIdErr"),
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(4356), big.NewInt(1e18)),
				},
				disputeBiggestStakeTxn: &Types.Transaction{},
				Hash:                   common.BigToHash(big.NewInt(1)),
				isEqual:                false,
				disputeErr:             nil,
			},
			want: errors.New("biggestInfluenceAndIdErr"),
		},

		{
			name: "Test 11: When DisputeBiggestStakeProposed transaction fails",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(4356), big.NewInt(1e18)),
				},
				disputeBiggestStakeErr: errors.New("disputeBiggestStake error"),
				Hash:                   common.BigToHash(big.NewInt(1)),
				isEqual:                false,
				disputeErr:             nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			blockManagerUtilsMock := new(mocks.BlockManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)
			utilsPkgMock := new(mocks2.Utils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			blockManagerUtils = blockManagerUtilsMock
			transactionUtils = transactionUtilsMock
			utils.UtilsInterface = utilsPkgMock

			utilsMock.On("GetSortedProposedBlockIds", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.sortedProposedBlockIds, tt.args.sortedProposedBlockIdsErr)
			cmdUtilsMock.On("GetBiggestStakeAndId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.biggestStake, tt.args.biggestStakeId, tt.args.biggestStakeErr)
			cmdUtilsMock.On("GetLocalMediansData", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.medians, tt.args.revealedCollectionIds, tt.args.revealedDataMaps, tt.args.medainsErr)
			utilsMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.proposedBlock, tt.args.proposedBlockErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			blockManagerUtilsMock.On("DisputeBiggestStakeProposed", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeBiggestStakeTxn, tt.args.disputeBiggestStakeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.Hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			cmdUtilsMock.On("CheckDisputeForIds", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.idDisputetxn, tt.args.idDisputeTxnErr)
			utilsPkgMock.On("GetLeafIdOfACollection", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.leafId, tt.args.leafIdErr)
			cmdUtilsMock.On("Dispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeErr)

			utils := &UtilsStruct{}
			err := utils.HandleDispute(client, config, account, epoch, blockNumber, rogueData)
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

func TestGiveSorted(t *testing.T) {
	var client *ethclient.Client
	var blockManager *bindings.BlockManager
	var txnOpts *bind.TransactOpts
	var epoch uint32
	var assetId uint16
	type args struct {
		sortedStakers []uint32
		giveSorted    *Types.Transaction
		giveSortedErr error
		hash          common.Hash
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1: When Give Sorted executes successfully",
			args: args{
				sortedStakers: []uint32{2, 1, 3, 5},
				giveSorted:    &Types.Transaction{},
				hash:          common.BigToHash(big.NewInt(1)),
			},
		},
		{
			name: "Test 2: When there is an error from GiveSorted",
			args: args{
				sortedStakers: []uint32{2, 1, 3, 5},
				giveSortedErr: errors.New("giveSorted error"),
			},
		},
		{
			name: "Test 3: When sortedStakers is nil",
			args: args{
				sortedStakers: nil,
			},
		},
		{
			name: "Test 4: When error is gas limit reached",
			args: args{
				sortedStakers: []uint32{2, 1, 3, 5},
				giveSortedErr: errors.New("gas limit reached"),
				giveSorted:    &Types.Transaction{},
				hash:          common.BigToHash(big.NewInt(1)),
			},
		},
		{
			name: "Test 5: When error is gas limit reached with higher number of stakers",
			args: args{
				sortedStakers: []uint32{2, 1, 3, 5, 7, 8, 9, 10, 6, 11, 13, 12, 14, 15, 4, 20, 19, 18, 17, 16},
				giveSortedErr: errors.New("gas limit reached"),
				giveSorted:    &Types.Transaction{},
				hash:          common.BigToHash(big.NewInt(1)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			blockManagerUtilsMock := new(mocks.BlockManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			blockManagerUtils = blockManagerUtilsMock
			transactionUtils = transactionUtilsMock

			blockManagerUtilsMock.On("GiveSorted", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.giveSorted, tt.args.giveSortedErr).Once()
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			blockManagerUtilsMock.On("GiveSorted", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.giveSorted, nil)

			GiveSorted(client, blockManager, txnOpts, epoch, assetId, tt.args.sortedStakers)
		})
	}
}
