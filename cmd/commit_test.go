package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	mocks2 "razor/utils/mocks"
	"reflect"
	"testing"
)

func TestCommit(t *testing.T) {
	var (
		client  *ethclient.Client
		account types.Account
		config  types.Configurations
		seed    []byte
		epoch   uint32
	)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		state     int64
		stateErr  error
		root      [32]byte
		txnOpts   *bind.TransactOpts
		commitTxn *Types.Transaction
		commitErr error
		hash      common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When Commit function executes successfully",
			args: args{
				state:     0,
				stateErr:  nil,
				txnOpts:   txnOpts,
				commitTxn: &Types.Transaction{},
				commitErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr:  errors.New("state error"),
				txnOpts:   txnOpts,
				commitTxn: &Types.Transaction{},
				commitErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When Commit transaction fails",
			args: args{
				state:     0,
				stateErr:  nil,
				txnOpts:   txnOpts,
				commitTxn: &Types.Transaction{},
				commitErr: errors.New("commit error"),
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("commit error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)
			voteManagerUtilsMock := new(mocks.VoteManagerInterface)

			razorUtils = utilsMock
			transactionUtils = transactionUtilsMock
			voteManagerUtils = voteManagerUtilsMock

			utilsMock.On("GetDelayedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(tt.args.txnOpts)
			voteManagerUtilsMock.On("Commit", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.commitTxn, tt.args.commitErr)
			transactionUtilsMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.Commit(client, config, account, epoch, seed, tt.args.root)
			if got != tt.want {
				t.Errorf("Txn hash for Commit function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for Commit function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for Commit function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestHandleCommitState(t *testing.T) {
	var (
		client *ethclient.Client
		epoch  uint32
		seed   []byte
	)

	rogueValue := utils.GetRogueRandomValue(100000)

	type args struct {
		numActiveCollections    uint16
		numActiveCollectionsErr error
		assignedCollections     map[int]bool
		seqAllottedCollections  []*big.Int
		assignedCollectionsErr  error
		collectionId            uint16
		collectionIdErr         error
		collectionData          *big.Int
		collectionDataErr       error
		rogueData               types.Rogue
	}
	tests := []struct {
		name    string
		args    args
		want    types.CommitData
		wantErr error
	}{
		{
			name: "Test 1: When HandleCommitState executes successfully",
			args: args{
				numActiveCollections:   3,
				assignedCollections:    map[int]bool{1: true, 2: true},
				seqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				collectionId:           1,
				collectionData:         big.NewInt(1),
			},
			want: types.CommitData{
				AssignedCollections:    map[int]bool{1: true, 2: true},
				SeqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				Leaves:                 []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)},
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting numActiveCollections",
			args: args{
				numActiveCollectionsErr: errors.New("error in getting numActiveCollections"),
			},
			want:    types.CommitData{},
			wantErr: errors.New("error in getting numActiveCollections"),
		},
		{
			name: "Test 3: When there is an error in getting assignedCollections",
			args: args{
				numActiveCollections:   1,
				assignedCollectionsErr: errors.New("error in getting assignedCollections"),
			},
			want:    types.CommitData{},
			wantErr: errors.New("error in getting assignedCollections"),
		},
		{
			name: "Test 4: When there is an error in getting collectionId",
			args: args{
				numActiveCollections:   3,
				assignedCollections:    map[int]bool{1: true, 2: true},
				seqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				collectionIdErr:        errors.New("error in getting collectionId"),
			},
			want:    types.CommitData{},
			wantErr: errors.New("error in getting collectionId"),
		},
		{
			name: "Test 5: When there is an error in getting collectionData",
			args: args{
				numActiveCollections:   3,
				assignedCollections:    map[int]bool{1: true, 2: true},
				seqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				collectionId:           1,
				collectionDataErr:      errors.New("error in getting collectionData"),
			},
			want:    types.CommitData{},
			wantErr: errors.New("error in getting collectionData"),
		},
		{
			name: "Test 6: When rogue mode is on for commit state",
			args: args{
				numActiveCollections:   3,
				assignedCollections:    map[int]bool{1: true, 2: true},
				seqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				collectionId:           1,
				collectionData:         big.NewInt(1),
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"commit"},
				},
			},
			want: types.CommitData{
				AssignedCollections:    map[int]bool{1: true, 2: true},
				SeqAllottedCollections: []*big.Int{big.NewInt(1), big.NewInt(2)},
				Leaves:                 []*big.Int{big.NewInt(0), rogueValue, rogueValue},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsPkgMock := new(mocks2.Utils)
			utilsMock := new(mocks.UtilsInterface)

			utils.UtilsInterface = utilsPkgMock
			razorUtils = utilsMock

			utilsPkgMock.On("GetNumActiveCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numActiveCollections, tt.args.numActiveCollectionsErr)
			utilsPkgMock.On("GetAssignedCollections", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.assignedCollections, tt.args.seqAllottedCollections, tt.args.assignedCollectionsErr)
			utilsPkgMock.On("GetCollectionIdFromIndex", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.collectionId, tt.args.collectionIdErr)
			utilsPkgMock.On("GetAggregatedDataOfCollection", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.collectionData, tt.args.collectionDataErr)
			utilsMock.On("GetRogueRandomValue", mock.Anything).Return(rogueValue)

			utils := &UtilsStruct{}
			got, err := utils.HandleCommitState(client, epoch, seed, tt.args.rogueData)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data from HandleCommitState function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error from HandleCommitState function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error from HandleCommitState function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestGetSalt(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		epoch                        uint32
		numProposedBlocks            uint8
		numProposedBlocksErr         error
		blockIndexedToBeConfirmed    int8
		blockIndexedToBeConfirmedErr error
		saltFromBlockChain           [32]byte
		saltFromBlockChainErr        error
		blockId                      uint32
		blockIdErr                   error
		previousBlock                bindings.StructsBlock
		previousBlockErr             error
		salt                         [32]byte
	}
	tests := []struct {
		name    string
		args    args
		want    [32]byte
		wantErr error
	}{
		{
			name: "Test 1: When GetSalt() function executes successfully",
			args: args{
				epoch:                     2,
				numProposedBlocks:         1,
				blockIndexedToBeConfirmed: 1,
				blockId:                   1,
				previousBlock:             bindings.StructsBlock{},
				salt:                      [32]byte{},
			},
			want:    [32]byte{},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting numProposedBlocks",
			args: args{
				epoch:                2,
				numProposedBlocksErr: errors.New("error in getting numProposedBlocks"),
			},
			want:    [32]byte{},
			wantErr: errors.New("error in getting numProposedBlocks"),
		},
		{
			name: "Test 3: When there is an error in getting blockIndexedToBeConfirmed",
			args: args{
				epoch:                        2,
				numProposedBlocks:            1,
				blockIndexedToBeConfirmedErr: errors.New("error in getting blockIndexedToBeConfirmed"),
			},
			want:    [32]byte{},
			wantErr: errors.New("error in getting blockIndexedToBeConfirmed"),
		},
		{
			name: "Test 4: When numProposedBlock is zero",
			args: args{
				epoch:              2,
				numProposedBlocks:  0,
				saltFromBlockChain: [32]byte{},
			},
			want:    [32]byte{},
			wantErr: nil,
		},
		{
			name: "Test 5: When there is an error in getting blockId",
			args: args{
				epoch:                     2,
				numProposedBlocks:         1,
				blockIndexedToBeConfirmed: 1,
				blockIdErr:                errors.New("error"),
			},
			want:    [32]byte{},
			wantErr: errors.New("Error in getting blockId: error"),
		},
		{
			name: "Test 6: When there is an error in getting previousBlock",
			args: args{
				epoch:                     2,
				numProposedBlocks:         1,
				blockIndexedToBeConfirmed: 1,
				blockId:                   1,
				previousBlockErr:          errors.New("error"),
			},
			want:    [32]byte{},
			wantErr: errors.New("Error in getting previous block: error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsPkgMock := new(mocks2.Utils)
			utilsVoteManagerMock := new(mocks2.VoteManagerUtils)

			utils.UtilsInterface = utilsPkgMock
			utils.VoteManagerInterface = utilsVoteManagerMock

			utilsPkgMock.On("GetNumberOfProposedBlocks", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.numProposedBlocks, tt.args.numProposedBlocksErr)
			utilsPkgMock.On("GetBlockIndexToBeConfirmed", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.blockIndexedToBeConfirmed, tt.args.blockIndexedToBeConfirmedErr)
			utilsVoteManagerMock.On("GetSaltFromBlockchain", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.saltFromBlockChain, tt.args.saltFromBlockChainErr)
			utilsPkgMock.On("GetSortedProposedBlockId", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.blockId, tt.args.blockIdErr)
			utilsPkgMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.previousBlock, tt.args.previousBlockErr)
			utilsPkgMock.On("CalculateSalt", mock.Anything, mock.Anything).Return(tt.args.salt)

			ut := &UtilsStruct{}
			got, err := ut.GetSalt(client, tt.args.epoch)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data from GetSalt function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error from GetSalt function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error from GetSalt function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
