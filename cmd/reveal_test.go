package cmd

import (
	"errors"
	"fmt"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

func TestCheckForLastCommitted(t *testing.T) {
	staker := bindings.StructsStaker{
		Id: 1,
	}

	type args struct {
		epoch                 uint32
		epochLastCommitted    uint32
		epochLastCommittedErr error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When CheckForLastCommitted returns no error",
			args: args{
				epoch:                 1,
				epochLastCommitted:    1,
				epochLastCommittedErr: nil,
			},
			want: nil,
		},
		{
			name: "Test 2: When there is an error in getting epochLastCommitted error",
			args: args{
				epoch:                 1,
				epochLastCommitted:    1,
				epochLastCommittedErr: errors.New("epochLastCommitted"),
			},
			want: errors.New("epochLastCommitted"),
		},
		{
			name: "Test 3: When CheckForLastCommitted returns an error when epoch != epochLastCommitted",
			args: args{
				epoch:                 3,
				epochLastCommitted:    2,
				epochLastCommittedErr: nil,
			},
			want: errors.New("commitment for this epoch not found on network.... aborting reveal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetEpochLastCommitted", mock.Anything, mock.Anything).Return(tt.args.epochLastCommitted, tt.args.epochLastCommittedErr)

			utils := &UtilsStruct{}

			err := utils.CheckForLastCommitted(rpcParameters, staker, tt.args.epoch)
			if err == nil || tt.want == nil {
				if err != tt.want {
					t.Errorf("Error for CheckForLastCommitted function, got = %v, want %v", err, tt.want)
				}
			} else {
				if err.Error() != tt.want.Error() {
					t.Errorf("Error for CheckForLastCommitted function, got = %v, want %v", err, tt.want)
				}
			}

		})
	}
}

func TestReveal(t *testing.T) {
	var (
		commitData   types.CommitData
		signature    []byte
		account      types.Account
		config       types.Configurations
		epoch        uint32
		latestHeader *Types.Header
		stateBuffer  uint64
	)

	type args struct {
		state          int64
		stateErr       error
		merkleTree     [][][]byte
		merkleTreeErr  error
		treeRevealData bindings.StructsMerkleTree
		txnOptsErr     error
		revealTxn      *Types.Transaction
		revealErr      error
		hash           common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When Reveal function executes successfully",
			args: args{
				state:     1,
				stateErr:  nil,
				revealTxn: &Types.Transaction{},
				revealErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr:  errors.New("state error"),
				revealTxn: &Types.Transaction{},
				revealErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When Reveal transaction fails",
			args: args{
				state:     1,
				stateErr:  nil,
				revealTxn: &Types.Transaction{},
				revealErr: errors.New("reveal error"),
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("reveal error"),
		},
		{
			name: "Test 4: When there is an error in getting merkle tree",
			args: args{
				state:         1,
				merkleTreeErr: errors.New("merkle tree error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("merkle tree error"),
		},
		{
			name: "Test 5: When there is an error in getting txnOpts",
			args: args{
				state:      1,
				stateErr:   nil,
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetBufferedState", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.state, tt.args.stateErr)
			merkleUtilsMock.On("CreateMerkle", mock.Anything).Return(tt.args.merkleTree, tt.args.merkleTreeErr)
			cmdUtilsMock.On("GenerateTreeRevealData", mock.Anything, mock.Anything).Return(tt.args.treeRevealData)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.AnythingOfType("types.TransactionOptions")).Return(TxnOpts, tt.args.txnOptsErr)
			voteManagerMock.On("Reveal", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32"), mock.Anything, mock.Anything).Return(tt.args.revealTxn, tt.args.revealErr)
			transactionMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.Reveal(rpcParameters, config, account, epoch, latestHeader, stateBuffer, commitData, signature)
			if got != tt.want {
				t.Errorf("Txn hash for Reveal function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for Reveal function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for Reveal function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGenerateTreeRevealData(t *testing.T) {
	type args struct {
		merkleTree [][][]byte
		commitData types.CommitData
		proof      [][32]byte
		root       [32]byte
		rootErr    error
	}
	tests := []struct {
		name string
		args args
		want bindings.StructsMerkleTree
	}{
		{
			name: "Test 1: When merkleTree and commitData is nil",
			args: args{
				merkleTree: [][][]byte{},
				commitData: types.CommitData{},
			},
			want: bindings.StructsMerkleTree{},
		},
		{
			name: "Test 2: When GenerateTreeRevealData executes successfully",
			args: args{
				merkleTree: [][][]byte{{{byte(1)}, {byte(2)}}, {{byte(3)}, {byte(4)}}, {{byte(5)}, {byte(6)}}},
				commitData: types.CommitData{
					AssignedCollections:    map[int]bool{1: true},
					SeqAllottedCollections: []*big.Int{big.NewInt(1)},
					Leaves:                 []*big.Int{big.NewInt(1), big.NewInt(2)},
				},
				proof: [][32]byte{},
				root:  [32]byte{},
			},
			want: bindings.StructsMerkleTree{
				Values: []bindings.StructsAssignedAsset{{LeafId: 1, Value: big.NewInt(2)}},
				Proofs: [][][32]byte{{}},
				Root:   [32]byte{},
			},
		},
		{
			name: "Test 3: When there is an error in getting root",
			args: args{
				merkleTree: [][][]byte{},
				commitData: types.CommitData{
					AssignedCollections:    map[int]bool{1: true},
					SeqAllottedCollections: []*big.Int{big.NewInt(1)},
					Leaves:                 []*big.Int{big.NewInt(1), big.NewInt(2)},
				},
				rootErr: errors.New("root error"),
			},
			want: bindings.StructsMerkleTree{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			merkleUtilsMock.On("GetProofPath", mock.Anything, mock.Anything).Return(tt.args.proof)
			merkleUtilsMock.On("GetMerkleRoot", mock.Anything).Return(tt.args.root, tt.args.rootErr)
			ut := &UtilsStruct{}
			if got := ut.GenerateTreeRevealData(tt.args.merkleTree, tt.args.commitData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTreeRevealData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexRevealEventsOfCurrentEpoch(t *testing.T) {
	var (
		blockNumber *big.Int
		epoch       uint32
	)

	type args struct {
		fromBlock      *big.Int
		fromBlockErr   error
		logs           []Types.Log
		logsErr        error
		contractAbi    abi.ABI
		contractAbiErr error
		data           []interface{}
		unpackErr      error
	}
	tests := []struct {
		name    string
		args    args
		want    []types.RevealedStruct
		wantErr bool
	}{
		{
			name: "Test 1: When IndexRevealEventsOfCurrentEpoch executes successfully",
			args: args{
				fromBlock:   big.NewInt(0),
				logs:        []Types.Log{},
				contractAbi: abi.ABI{},
				data:        []interface{}{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting fromBlock",
			args: args{
				fromBlockErr: errors.New("error in getting fromBlock"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting logs",
			args: args{
				fromBlock: big.NewInt(0),
				logsErr:   errors.New("error in getting logs"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When there is an error in getting contractAbi",
			args: args{
				fromBlock:      big.NewInt(0),
				logs:           []Types.Log{},
				contractAbiErr: errors.New("error in getting contractAbi"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("EstimateBlockNumberAtEpochBeginning", mock.Anything, mock.Anything).Return(tt.args.fromBlock, tt.args.fromBlockErr)
			clientUtilsMock.On("FilterLogsWithRetry", mock.Anything, mock.Anything).Return(tt.args.logs, tt.args.logsErr)
			abiUtilsMock.On("Parse", mock.Anything).Return(tt.args.contractAbi, tt.args.contractAbiErr)
			abiUtilsMock.On("Unpack", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.data, tt.args.unpackErr)
			ut := &UtilsStruct{}
			got, err := ut.IndexRevealEventsOfCurrentEpoch(rpcParameters, blockNumber, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexRevealEventsOfCurrentEpoch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexRevealEventsOfCurrentEpoch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenerateTreeRevealData(b *testing.B) {
	table := []struct {
		numOfAllottedCollections int
	}{
		{numOfAllottedCollections: 5},
		{numOfAllottedCollections: 50},
		{numOfAllottedCollections: 500},
		{numOfAllottedCollections: 5000},
		{numOfAllottedCollections: 10000},
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("Number_Of_Allotted_Collections%d", v.numOfAllottedCollections), func(b *testing.B) {
			merkleTree := [][][]byte{{{byte(1)}, {byte(2)}}, {{byte(3)}, {byte(4)}}, {{byte(5)}, {byte(6)}}}
			SetUpMockInterfaces()

			merkleUtilsMock.On("GetProofPath", mock.Anything, mock.Anything).Return([][32]byte{[32]byte{1, 2, 3}, {4, 5, 6}})
			merkleUtilsMock.On("GetMerkleRoot", mock.Anything).Return([32]byte{100}, nil)

			ut := &UtilsStruct{}
			seqAllottedCollections := getDummySeqAllottedCollection(v.numOfAllottedCollections)
			ut.GenerateTreeRevealData(merkleTree, types.CommitData{
				SeqAllottedCollections: seqAllottedCollections,
				Leaves:                 GetDummyVotes(v.numOfAllottedCollections),
			})
		})
	}
}

func getDummySeqAllottedCollection(numOfAllottedCollections int) []*big.Int {
	var result []*big.Int
	for i := 1; i <= numOfAllottedCollections; i++ {
		result = append(result, big.NewInt(1))
	}
	return result
}
