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
	Mocks "razor/utils/mocks"
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
		bufferPercent      int32
		bufferPercentErr   error
		remainingTime      int64
		remainingTimeErr   error
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
				epoch:         4,
				numOfStakers:  3,
				remainingTime: 10,
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
				epoch:         4,
				numOfStakers:  3,
				remainingTime: 10,
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
				epoch:         4,
				numOfStakers:  3,
				remainingTime: 10,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				containsStatus:     false,
				finalizeDisputeErr: errors.New("finalizeDispute error"),
			},
			want: errors.New("finalizeDispute error"),
		},
		{
			name: "Test 6: When there is an error in getting remaining time",
			args: args{
				epoch:        4,
				numOfStakers: 3,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				remainingTimeErr: errors.New("time error"),
			},
			want: errors.New("time error"),
		},
		{
			name: "Test 7: When there is a timeout case",
			args: args{
				epoch:        4,
				numOfStakers: 10000000,
				votes: bindings.StructsVote{
					Epoch:  4,
					Values: []*big.Int{big.NewInt(100), big.NewInt(200)},
				},
				remainingTime: 1,
			},
			want: errors.New("dispute state timeout"),
		},
		{
			name: "Test 8: When there is an error in getting buffer percent",
			args: args{
				epoch:            4,
				numOfStakers:     3,
				remainingTime:    10,
				bufferPercentErr: errors.New("buffer error"),
			},
			want: errors.New("buffer error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			blockManagerUtilsMock := new(mocks.BlockManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)
			utilsPkgMock := new(Mocks.Utils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			blockManagerUtils = blockManagerUtilsMock
			transactionUtils = transactionUtilsMock
			utilsInterface = utilsPkgMock

			utilsMock.On("GetBlockManager", mock.AnythingOfType("*ethclient.Client")).Return(blockManager)
			utilsMock.On("GetNumberOfStakers", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.numOfStakers, tt.args.numOfStakersErr)
			utilsMock.On("GetVotes", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.votes, tt.args.votesErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			cmdUtilsMock.On("GiveSorted", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
			blockManagerUtilsMock.On("FinalizeDispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.finalizeDisputeTxn, tt.args.finalizeDisputeErr)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			utilsPkgMock.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(tt.args.remainingTime, tt.args.remainingTimeErr)
			cmdUtilsMock.On("GetBufferPercent").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)

			utils := &UtilsStruct{}

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
		fileName                  string
		fileNameErr               error
		epochInFile               uint32
		medianDataInFile          []*big.Int
		readFileErr               error
		mediansInUint32           []uint32
		activeAssetIds            []uint16
		activeAssetIdsErr         error
		isEqual                   bool
		iteration                 int
		disputeErr                error
		epoch                     uint32
		medians                   []uint32
		mediansErr                error
		mediansBigInt             []*big.Int
		rogue                     types.Rogue
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
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
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
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         true,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
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
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         true,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: errors.New("sortedProposedBlockIds error"),
		},
		{
			name: "Test 4: When there is an error in getting proposedBlock",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				proposedBlockErr:       errors.New("proposedBlock error"),
				mediansInUint32:        []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                true,
				disputeErr:             nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: nil,
		},
		{
			name: "Test 5: When there is an error in getting fileName",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				fileNameErr:            errors.New("fileName error"),
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: errors.New("fileName error"),
		},
		{
			name: "Test 6: When there is an error in reading file",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				readFileErr:            errors.New("readFile error"),
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: errors.New("readFile error"),
		},
		{
			name: "Test 7: When there is an error from Dispute function",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
					Valid:        true,
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      errors.New("dispute error"),
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: nil,
		},
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
				mediansInUint32: []uint32{6901548, 498307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				rogue: types.Rogue{
					IsRogue: false,
				},
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
				mediansInUint32:        []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
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
				mediansInUint32:        []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
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
				mediansInUint32:        []uint32{6701548, 478307},
				activeAssetIds:         []uint16{3, 5},
				isEqual:                false,
				iteration:              0,
				disputeErr:             nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: nil,
		},
		{
			name: "Test 12: When epochInFile is not equal to correct epoch",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				epoch:                  2,
				epochInFile:            3,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: false,
				},
			},
			want: nil,
		},
		{
			name: "Test 13: When HandleDispute function executes successfully and staker is in rogue mode",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				rogue: types.Rogue{
					IsRogue: true,
				},
			},
			want: nil,
		},
		{
			name: "Test 14: When there is an error in getting medians and staker is in rogue mode",
			args: args{
				sortedProposedBlockIds: []uint32{3, 1, 2, 5, 4},
				biggestStake:           big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				biggestStakeId:         2,
				proposedBlock: bindings.StructsBlock{
					Medians:      []uint32{6901548, 498307},
					Valid:        true,
					BiggestStake: big.NewInt(1).Mul(big.NewInt(5356), big.NewInt(1e18)),
				},
				mediansInUint32: []uint32{6701548, 478307},
				activeAssetIds:  []uint16{3, 5},
				isEqual:         false,
				iteration:       0,
				disputeErr:      nil,
				mediansErr:      errors.New("error in getting medians"),
				rogue: types.Rogue{
					IsRogue: true,
				},
			},
			want: errors.New("error in getting medians"),
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

			utilsMock.On("GetSortedProposedBlockIds", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.sortedProposedBlockIds, tt.args.sortedProposedBlockIdsErr)
			utilsMock.On("GetProposedBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32"), mock.AnythingOfType("uint32")).Return(tt.args.proposedBlock, tt.args.proposedBlockErr)
			cmdUtilsMock.On("GetBiggestStakeAndId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.AnythingOfType("uint32")).Return(tt.args.biggestStake, tt.args.biggestStakeId, tt.args.biggestStakeErr)
			blockManagerUtilsMock.On("DisputeBiggestStakeProposed", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeBiggestStakeTxn, tt.args.disputeBiggestStakeErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(txnOpts)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.Hash)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
			cmdUtilsMock.On("GetMedianDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
			utilsMock.On("ReadDataFromFile", mock.AnythingOfType("string")).Return(tt.args.epochInFile, tt.args.medianDataInFile, tt.args.readFileErr)
			cmdUtilsMock.On("MakeBlock", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string"), mock.Anything).Return(tt.args.medians, tt.args.mediansErr)
			utilsMock.On("ConvertUint32ArrayToBigIntArray", mock.Anything).Return(tt.args.mediansBigInt)
			utilsMock.On("ConvertBigIntArrayToUint32Array", mock.Anything).Return(tt.args.mediansInUint32)
			utilsMock.On("GetActiveAssetIds", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.activeAssetIds, tt.args.activeAssetIdsErr)
			cmdUtilsMock.On("Dispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.disputeErr)

			utils := &UtilsStruct{}
			err := utils.HandleDispute(client, config, account, tt.args.epoch, tt.args.rogue)
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
