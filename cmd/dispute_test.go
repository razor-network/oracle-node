package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	"reflect"
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
			name: "Test 3: When FinalizeDispute transaction fails",
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
		mediansErr                error
		proposedBlock             bindings.StructsBlock
		proposedBlockErr          error
		disputeBiggestStakeTxn    *Types.Transaction
		disputeBiggestStakeErr    error
		Hash                      common.Hash
		idDisputeTxn              *Types.Transaction
		idDisputeTxnErr           error
		status                    int
		isEqual                   bool
		misMatchIndex             int
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
			name: "Test 1: When HandleDispute function executes successfully",
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
				isEqual:    true,
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
		{
			name: "Test 5: When there is a case of Dispute but block is already disputed",
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
			name: "Test 6: When HandleDispute function executes successfully when there is a biggest influence dispute case",
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
				status:                 1,
			},
			want: nil,
		},
		{
			name: "Test 7: When there is an error in getting biggestInfluenceAndId",
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
			name: "Test 8: When DisputeBiggestStakeProposed transaction fails",
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
		{
			name: "Test 9: When there is an error in getting medians",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				mediansErr:             errors.New("error in getting medians"),
			},
			want: errors.New("error in getting medians"),
		},
		{
			name: "Test 10: When there is an error in fetching Ids from CheckDisputeForIds",
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
				idDisputeTxnErr: errors.New("error in fetching Ids from CheckDisputeForIds"),
				isEqual:         true,
				disputeErr:      nil,
			},
			want: nil,
		},
		{
			name: "Test 11: When idDisputeTxn is not nil and status is 1",
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
				idDisputeTxn: &Types.Transaction{},
				status:       1,
				isEqual:      false,
				disputeErr:   nil,
			},
			want: nil,
		},
		{
			name: "Test 12: When it is a median dispute case and error in getting leafId",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				medians:                []uint32{69015, 498307},
				revealedCollectionIds:  []uint16{1},
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: nil,
					VoteWeights:          nil,
					InfluenceSum:         nil,
				},
				proposedBlock: bindings.StructsBlock{
					Ids:          []uint16{1, 2},
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				isEqual:       false,
				misMatchIndex: 0,
				leafIdErr:     errors.New("error in getting leafId"),
				disputeErr:    nil,
			},
			want: nil,
		},
		{
			name: "Test 13: When there is an error in dispute",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				medians:                []uint32{69015, 498307},
				revealedCollectionIds:  []uint16{1},
				revealedDataMaps: &types.RevealedDataMaps{
					SortedRevealedValues: nil,
					VoteWeights:          nil,
					InfluenceSum:         nil,
				},
				proposedBlock: bindings.StructsBlock{
					Ids:          []uint16{1, 2},
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				isEqual:       false,
				misMatchIndex: 0,
				leafId:        1,
				disputeErr:    errors.New("error in dispute"),
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
			cmdUtilsMock.On("GetLocalMediansData", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.medians, tt.args.revealedCollectionIds, tt.args.revealedDataMaps, tt.args.mediansErr)
			utilsMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.proposedBlock, tt.args.proposedBlockErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			blockManagerUtilsMock.On("DisputeBiggestStakeProposed", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeBiggestStakeTxn, tt.args.disputeBiggestStakeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.Hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.status)
			cmdUtilsMock.On("CheckDisputeForIds", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.idDisputeTxn, tt.args.idDisputeTxnErr)
			utilsPkgMock.On("IsEqualUint32", mock.Anything, mock.Anything).Return(tt.args.isEqual, tt.args.misMatchIndex)
			utilsPkgMock.On("GetLeafIdOfACollection", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.leafId, tt.args.leafIdErr)
			cmdUtilsMock.On("Dispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeErr)

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

func TestGetLocalMediansData(t *testing.T) {
	var (
		client      *ethclient.Client
		account     types.Account
		blockNumber *big.Int
		rogueData   types.Rogue
	)
	type args struct {
		epoch                 uint32
		fileName              string
		fileNameErr           error
		proposedData          types.ProposeFileData
		proposeDataErr        error
		medians               []uint32
		revealedCollectionIds []uint16
		revealedDataMaps      *types.RevealedDataMaps
		mediansErr            error
		mediansBigInt         []*big.Int
		mediansInUint32       []uint32
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
			name: "Test 1: When there is an error in getting fileName",
			args: args{
				fileNameErr: errors.New("error in getting fileName"),
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting proposedData",
			args: args{
				fileName:       "",
				proposeDataErr: errors.New("error in getting proposedData"),
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: false,
		},
		{
			name: "Test 3: When file does not contain latest data",
			args: args{
				fileName:     "",
				proposedData: types.ProposeFileData{Epoch: 3},
				epoch:        5,
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting medians",
			args: args{
				mediansErr: errors.New("error in fetching medians"),
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "Test 5: When GetLocalMediansData executes successfully",
			args: args{
				medians:               []uint32{100, 200, 300},
				revealedCollectionIds: []uint16{1, 2, 3},
				revealedDataMaps:      &types.RevealedDataMaps{},
				mediansBigInt:         []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
				mediansInUint32:       []uint32{100, 200, 300},
			},
			want:    []uint32{100, 200, 300},
			want1:   []uint16{1, 2, 3},
			want2:   &types.RevealedDataMaps{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("GetProposeDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
			utilsMock.On("ReadFromProposeJsonFile", mock.Anything).Return(tt.args.proposedData, tt.args.proposeDataErr)
			cmdUtilsMock.On("MakeBlock", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything, mock.Anything).Return(tt.args.medians, tt.args.revealedCollectionIds, tt.args.revealedDataMaps, tt.args.mediansErr)
			utilsMock.On("ConvertUint32ArrayToBigIntArray", mock.Anything).Return(tt.args.mediansBigInt)
			utilsMock.On("ConvertBigIntArrayToUint32Array", mock.Anything).Return(tt.args.mediansInUint32)
			ut := &UtilsStruct{}
			got, got1, got2, err := ut.GetLocalMediansData(client, account, tt.args.epoch, blockNumber, rogueData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalMediansData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocalMediansData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetLocalMediansData() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetLocalMediansData() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestGetCollectionIdPositionInBlock(t *testing.T) {
	var client *ethclient.Client
	var leafId uint16
	type args struct {
		proposedBlock     bindings.StructsBlock
		idToBeDisputed    uint16
		idToBeDisputedErr error
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1: When GetCollectionIdPositionInBlock() executes successfully",
			args: args{
				proposedBlock: bindings.StructsBlock{
					Ids: []uint16{1, 2, 3},
				},
				idToBeDisputed: 1,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 2: When there is an error in getting idToBeDisputed",
			args: args{
				proposedBlock: bindings.StructsBlock{
					Ids: []uint16{1, 2, 3},
				},
				idToBeDisputedErr: errors.New("error in fetching idToBeDisputed"),
			},
			want: nil,
		},
		{
			name: "Test 3: When idToBeDisputes is not present in proposedBlock.Ids",
			args: args{
				proposedBlock: bindings.StructsBlock{
					Ids: []uint16{1, 2, 3},
				},
				idToBeDisputed: 4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsPkgMock := new(mocks2.Utils)

			utils.UtilsInterface = utilsPkgMock

			utilsPkgMock.On("GetCollectionIdFromLeafId", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.idToBeDisputed, tt.args.idToBeDisputedErr)
			ut := &UtilsStruct{}
			if got := ut.GetCollectionIdPositionInBlock(client, leafId, tt.args.proposedBlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollectionIdPositionInBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckDisputeForIds(t *testing.T) {
	var (
		client          *ethclient.Client
		transactionOpts types.TransactionOptions
		epoch           uint32
		blockIndex      uint8
	)

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	type args struct {
		idsInProposedBlock                    []uint16
		revealedCollectionIds                 []uint16
		isSorted                              bool
		index0                                int
		index1                                int
		DisputeOnOrderOfIds                   *Types.Transaction
		DisputeOnOrderOfIdsErr                error
		isMissing                             bool
		isMissingInt                          int
		missingCollectionId                   uint16
		DisputeCollectionIdShouldBePresent    *Types.Transaction
		DisputeCollectionIdShouldBePresentErr error
		isPresent                             bool
		positionOfPresentValue                int
		presentCollectionId                   uint16
		DisputeCollectionIdShouldBeAbsent     *Types.Transaction
		DisputeCollectionIdShouldBeAbsentErr  error
	}
	tests := []struct {
		name    string
		args    args
		want    *Types.Transaction
		wantErr bool
	}{
		{
			name: "Test 1: When CheckDisputeForIds executes successfully and check if the error is in sorted ids",
			args: args{
				idsInProposedBlock:  []uint16{1, 2, 3},
				isSorted:            false,
				DisputeOnOrderOfIds: &Types.Transaction{},
			},
			want:    &Types.Transaction{},
			wantErr: false,
		},
		{
			name: "Test 2: When CheckDisputeForIds executes successfully and check if the error is collectionIdShouldBePresent",
			args: args{
				idsInProposedBlock:                 []uint16{1, 2, 3},
				isSorted:                           true,
				isMissing:                          true,
				DisputeCollectionIdShouldBePresent: &Types.Transaction{},
			},
			want:    &Types.Transaction{},
			wantErr: false,
		},
		{
			name: "Test 3: When there is no error",
			args: args{
				idsInProposedBlock:                []uint16{1, 2, 3},
				isSorted:                          true,
				isMissing:                         false,
				isPresent:                         false,
				DisputeCollectionIdShouldBeAbsent: &Types.Transaction{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			blockManagerUtilsMock := new(mocks.BlockManagerInterface)
			utilsPkgMock := new(mocks2.Utils)

			utils.UtilsInterface = utilsPkgMock
			razorUtils = utilsMock
			blockManagerUtils = blockManagerUtilsMock
			utilsInterface = utilsPkgMock

			utilsPkgMock.On("IsSorted", mock.Anything).Return(tt.args.isSorted, tt.args.index0, tt.args.index1)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			blockManagerUtilsMock.On("DisputeOnOrderOfIds", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.DisputeOnOrderOfIds, tt.args.DisputeOnOrderOfIdsErr)
			utilsPkgMock.On("IsMissing", mock.Anything, mock.Anything).Return(tt.args.isMissing, tt.args.isMissingInt, tt.args.missingCollectionId)
			blockManagerUtilsMock.On("DisputeCollectionIdShouldBePresent", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.DisputeCollectionIdShouldBePresent, tt.args.DisputeCollectionIdShouldBePresentErr)
			utilsPkgMock.On("IsMissing", mock.Anything, mock.Anything).Return(tt.args.isPresent, tt.args.positionOfPresentValue, tt.args.presentCollectionId)
			blockManagerUtilsMock.On("DisputeCollectionIdShouldBeAbsent", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.DisputeCollectionIdShouldBeAbsent, tt.args.DisputeCollectionIdShouldBeAbsentErr)
			utilsPkgMock.On("IncreaseGasLimitValue", mock.Anything, mock.Anything, mock.Anything).Return(uint64(2000), nil)
			ut := &UtilsStruct{}
			got, err := ut.CheckDisputeForIds(client, transactionOpts, epoch, blockIndex, tt.args.idsInProposedBlock, tt.args.revealedCollectionIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDisputeForIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDisputeForIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBountyIdFromEvents(t *testing.T) {
	var (
		client       *ethclient.Client
		blockNumber  *big.Int
		bountyHunter string
	)
	type args struct {
		fromBlock      *big.Int
		fromBlockErr   error
		logs           []Types.Log
		logsErr        error
		contractABI    abi.ABI
		contractABIErr error
		data           []interface{}
		unpackErr      error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "Test 1: When GetBountyIdFromEvents() executes successfully",
			args: args{
				fromBlock: big.NewInt(0),
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractABI: abi.ABI{},
				data:        convertToSliceOfInterface([]uint32{4, 2}),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting blockNumber",
			args: args{
				fromBlockErr: errors.New("error in getting blockNumber"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting logs",
			args: args{
				fromBlock: big.NewInt(0),
				logsErr:   errors.New("error in getting logs"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting contractABI",
			args: args{
				fromBlock: big.NewInt(0),
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractABIErr: errors.New("error in contractABI"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test 5: When there is an error in unpacking",
			args: args{
				fromBlock: big.NewInt(0),
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractABI: abi.ABI{},
				unpackErr:   errors.New("error in unpacking"),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			abiMock := new(mocks.AbiInterface)
			utilsPkgMock := new(mocks2.Utils)
			abiUtilsMock := new(mocks2.ABIUtils)
			utilsPkgMock2 := new(mocks2.Utils)

			utilsInterface = utilsPkgMock2
			abiUtils = abiMock
			utils.UtilsInterface = utilsPkgMock
			utils.ABIInterface = abiUtilsMock

			utilsPkgMock.On("CalculateBlockNumberAtEpochBeginning", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.fromBlock, tt.args.fromBlockErr)
			abiUtilsMock.On("Parse", mock.Anything).Return(tt.args.contractABI, tt.args.contractABIErr)
			utilsPkgMock.On("FilterLogsWithRetry", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("ethereum.FilterQuery")).Return(tt.args.logs, tt.args.logsErr)
			abiMock.On("Unpack", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.data, tt.args.unpackErr)
			ut := &UtilsStruct{}
			got, err := ut.GetBountyIdFromEvents(client, blockNumber, bountyHunter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBountyIdFromEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBountyIdFromEvents() got = %v, want %v", got, tt.want)
			}
		})
	}
}
