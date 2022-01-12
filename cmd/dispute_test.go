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

			utilsMock := new(mocks.UtilsInterfaceMockery)
			cmdUtilsMock := new(mocks.UtilsCmdInterfaceMockery)
			blockManagerUtilsMock := new(mocks.BlockManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			cmdUtilsMockery = cmdUtilsMock
			blockManagerUtilsMockery = blockManagerUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetBlockManager", mock.AnythingOfType("*ethclient.Client")).Return(blockManager)
			utilsMock.On("GetNumberOfStakers", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.numOfStakers, tt.args.numOfStakersErr)
			utilsMock.On("GetVotes", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.votes, tt.args.votesErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("GiveSorted", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
			blockManagerUtilsMock.On("FinalizeDispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.finalizeDisputeTxn, tt.args.finalizeDisputeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)

			utils := &UtilsStructMockery{}

			err := utils.Dispute(client, config, account, tt.args.epoch, blockId, assetId)
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

	type args struct {
		sortedProposedBlockIds    []uint32
		sortedProposedBlockIdsErr error
		proposedBlock             bindings.StructsBlock
		proposedBlockErr          error
		biggestStake              *big.Int
		biggestStakeId            uint32
		biggestStakeErr           error
		disputeBiggestStakeTxn    *Types.Transaction
		disputeBiggestStakeErr    error
		Hash                      common.Hash
		medians                   []uint32
		mediansErr                error
		activeAssetIds            []uint16
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
			name: "Test 1: When HandleDispute function executes successfully when there is a medians dispute case",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				medians:        []uint32{6701548, 478307},
				activeAssetIds: []uint16{3, 5},
				isEqual:        false,
				iteration:      0,
				disputeErr:     nil,
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
				medians:        []uint32{6701548, 478307},
				activeAssetIds: []uint16{3, 5},
				isEqual:        true,
				disputeErr:     nil,
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
				medians:        []uint32{6701548, 478307},
				activeAssetIds: []uint16{3, 5},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: errors.New("sortedProposedBlockIds error"),
		},
		{
			name: "Test 4: When there is an error in getting proposedBlock",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				proposedBlockErr:       errors.New("proposedBlock error"),
				medians:                []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                true,
				disputeErr:             nil,
			},
			want: nil,
		},
		{
			name: "Test 5: When there is an error in getting medians from MakeBlock",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansErr:     errors.New("medians error"),
				activeAssetIds: []uint16{3, 5},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: errors.New("medians error"),
		},
		{
			name: "Test 6: When there is an error from Dispute function",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
					Valid:        true,
				},
				medians:        []uint32{6701548, 478307},
				activeAssetIds: []uint16{3, 5},
				isEqual:        false,
				iteration:      0,
				disputeErr:     errors.New("dispute error"),
			},
			want: nil,
		},
		{
			name: "Test 7: When there is a case of Dispute but block is already disputed",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6701548, 478307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				medians:        []uint32{6901548, 498307},
				activeAssetIds: []uint16{3, 5},
				isEqual:        false,
				iteration:      0,
			},
			want: nil,
		},
		{
			name: "Test 8: When HandleDispute function executes successfully when there is a biggest influence dispute case",
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
				medians:                []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
			},
			want: nil,
		},
		{
			name: "Test 9: When there is an error in getting biggestInfluenceAndId",
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
				medians:                []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
			},
			want: errors.New("biggestInfluenceAndIdErr"),
		},

		{
			name: "Test 10: When DisputeBiggestStakeProposed transaction fails",
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
				medians:                []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			cmdUtilsMock := new(mocks.UtilsCmdInterfaceMockery)
			blockManagerUtilsMock := new(mocks.BlockManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			cmdUtilsMockery = cmdUtilsMock
			blockManagerUtilsMockery = blockManagerUtilsMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetSortedProposedBlockIds", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.sortedProposedBlockIds, tt.args.sortedProposedBlockIdsErr)
			utilsMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.proposedBlock, tt.args.proposedBlockErr)
			cmdUtilsMock.On("GetBiggestStakeAndId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.biggestStake, tt.args.biggestStakeId, tt.args.biggestStakeErr)
			blockManagerUtilsMock.On("DisputeBiggestStakeProposed", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeBiggestStakeTxn, tt.args.disputeBiggestStakeErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.Hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			cmdUtilsMock.On("MakeBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.Anything).Return(tt.args.medians, tt.args.mediansErr)
			utilsMock.On("GetActiveAssetIds", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.activeAssetIds, tt.args.activeAssetIdsErr)
			cmdUtilsMock.On("Dispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeErr)

			utils := &UtilsStructMockery{}
			err := utils.HandleDispute(client, config, account, epoch)
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
