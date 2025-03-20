package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"reflect"
	"testing"
)

func TestCommit(t *testing.T) {
	var (
		account      types.Account
		config       types.Configurations
		latestHeader *Types.Header
		stateBuffer  uint64
		epoch        uint32
		commitment   [32]byte
	)

	type args struct {
		state      int64
		stateErr   error
		txnOptsErr error
		commitTxn  *Types.Transaction
		commitErr  error
		hash       common.Hash
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
				commitTxn: &Types.Transaction{},
				commitErr: errors.New("commit error"),
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("commit error"),
		},
		{
			name: "Test 4: When there is an error in getting txnOpts",
			args: args{
				state:      0,
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
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			voteManagerMock.On("Commit", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32"), mock.Anything).Return(tt.args.commitTxn, tt.args.commitErr)
			transactionMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.Commit(rpcParameters, config, account, epoch, latestHeader, stateBuffer, commitment)
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
		epoch uint32
		seed  []byte
	)

	rogueValue := big.NewInt(1111)

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
			localCache := cache.NewLocalCache()
			commitParams := &types.CommitParams{
				LocalCache: localCache,
			}

			SetUpMockInterfaces()

			utilsMock.On("GetNumActiveCollections", mock.Anything).Return(tt.args.numActiveCollections, tt.args.numActiveCollectionsErr)
			utilsMock.On("GetAssignedCollections", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.assignedCollections, tt.args.seqAllottedCollections, tt.args.assignedCollectionsErr)
			utilsMock.On("GetCollectionIdFromIndex", mock.Anything, mock.Anything).Return(tt.args.collectionId, tt.args.collectionIdErr)
			utilsMock.On("GetAggregatedDataOfCollection", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.collectionData, tt.args.collectionDataErr)
			utilsMock.On("GetRogueRandomValue", mock.Anything).Return(rogueValue)

			utils := &UtilsStruct{}
			got, err := utils.HandleCommitState(rpcParameters, epoch, seed, commitParams, tt.args.rogueData)
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
			SetUpMockInterfaces()

			utilsMock.On("GetNumberOfProposedBlocks", mock.Anything, mock.Anything).Return(tt.args.numProposedBlocks, tt.args.numProposedBlocksErr)
			utilsMock.On("GetBlockIndexToBeConfirmed", mock.Anything).Return(tt.args.blockIndexedToBeConfirmed, tt.args.blockIndexedToBeConfirmedErr)
			utilsMock.On("GetSaltFromBlockchain", mock.Anything).Return(tt.args.saltFromBlockChain, tt.args.saltFromBlockChainErr)
			utilsMock.On("GetSortedProposedBlockId", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.blockId, tt.args.blockIdErr)
			utilsMock.On("GetProposedBlock", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.previousBlock, tt.args.previousBlockErr)
			utilsMock.On("CalculateSalt", mock.Anything, mock.Anything).Return(tt.args.salt)

			ut := &UtilsStruct{}
			got, err := ut.GetSalt(rpcParameters, tt.args.epoch)
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

func BenchmarkHandleCommitState(b *testing.B) {
	var (
		epoch uint32
		seed  []byte
	)

	rogueValue := big.NewInt(1111)

	var table = []struct {
		numActiveCollections uint16
		assignedCollections  map[int]bool
	}{
		{numActiveCollections: 5, assignedCollections: map[int]bool{1: true, 2: true, 3: true}},
		{numActiveCollections: 10, assignedCollections: map[int]bool{1: true, 2: true, 3: true, 4: true}},
		{numActiveCollections: 20, assignedCollections: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true}},
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("Number_Of_Active_Collections%d", v.numActiveCollections), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				localCache := cache.NewLocalCache()
				commitParams := &types.CommitParams{
					LocalCache: localCache,
				}

				SetUpMockInterfaces()

				utilsMock.On("GetNumActiveCollections", mock.Anything).Return(v.numActiveCollections, nil)
				utilsMock.On("GetAssignedCollections", mock.Anything, mock.Anything, mock.Anything).Return(v.assignedCollections, nil, nil)
				utilsMock.On("GetCollectionIdFromIndex", mock.Anything, mock.Anything).Return(uint16(1), nil)
				utilsMock.On("GetAggregatedDataOfCollection", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(1000), nil)
				utilsMock.On("GetRogueRandomValue", mock.Anything).Return(rogueValue)

				ut := &UtilsStruct{}
				_, err := ut.HandleCommitState(rpcParameters, epoch, seed, commitParams, types.Rogue{IsRogue: false})
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

func TestCalculateCommitment(t *testing.T) {
	type args struct {
		seed   []byte
		values []*big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    string // Changed type from [32]byte to string
		wantErr bool
	}{
		{
			name: "Test 1: When there the values for seed and values are valid",
			args: args{
				seed:   []byte("5ab3bd027e66773306cc8c889dc48b17753d7ac6e400e066e91c3f8119540c6c"),
				values: []*big.Int{big.NewInt(200), big.NewInt(100)},
			},
			want:    "61fc5d313bb53f669154b2778a5c93859c4eb389f799166104691135869d7947",
			wantErr: false,
		},
		{
			name: "Test 2: When length of values array is 0",
			args: args{
				seed:   []byte("5ab3bd027e66773306cc8c889dc48b17753d7ac6e400e066e91c3f8119540c6c"),
				values: []*big.Int{},
			},
			want:    "0000000000000000000000000000000000000000000000000000000000000000",
			wantErr: true,
		},
		{
			name: "Test 3: When seed is empty",
			args: args{
				seed:   []byte{},
				values: []*big.Int{big.NewInt(200), big.NewInt(100)},
			},
			want:    "643e39018427c8db4cc8bbfdb9f04cd485032b6bc924db1bbf6b019391d032e9",
			wantErr: false,
		},
		{
			name: "Test 4: when When length of values array is 0 and seed is empty",
			args: args{
				seed:   []byte{},
				values: []*big.Int{},
			},
			want:    "0000000000000000000000000000000000000000000000000000000000000000",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utils.MerkleInterface = &utils.MerkleTreeStruct{}
			merkleUtils = utils.MerkleInterface

			got, err := CalculateCommitment(tt.args.seed, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateCommitment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			gotString := hex.EncodeToString(got[:]) // Convert [32]byte to hex string for comparison
			fmt.Println(gotString)
			if !reflect.DeepEqual(gotString, tt.want) {
				t.Errorf("CalculateCommitment() got = %v, want %v", gotString, tt.want)
			}
		})
	}
}

func TestVerifyCommitment(t *testing.T) {
	var (
		account types.Account
	)
	type args struct {
		commitmentFetchedString string
		commitmentHashString    string
		commitmentErr           error
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test 1: When commitmentFetched matches the commitment in the epoch",
			args: args{
				commitmentHashString:    "22c9ba074e44d0009116b244a5cece9e9ade85af486e1f4f8db8e304e6605bea",
				commitmentFetchedString: "22c9ba074e44d0009116b244a5cece9e9ade85af486e1f4f8db8e304e6605bea",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test 2: When commitmentFetched does not match the commitment in the epoch",
			args: args{
				commitmentHashString:    "23cabb074e44d0009116b244a5cece9e9ade85af486e1f4f8db8e304e6605bea",
				commitmentFetchedString: "22c9ba074e44d0009116b244a5cece9e9ade85af486e1f4f8db8e304e6605bea",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in fetching commitment from the blockchain",
			args: args{
				commitmentErr: errors.New("getCommitment error"),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var commitmentHash [32]byte
			var commitmentFetched [32]byte

			// Convert the commitmentFetchedString to a [32]byte format
			if tt.args.commitmentFetchedString != "" {
				var err error
				commitmentFetched, err = convertStringToByte32(tt.args.commitmentFetchedString)
				if err != nil {
					t.Errorf("Error in decoding commitmentFetchedString: %v", err)
					return
				}
			}

			// Convert the commitmentHashString to a [32]byte format
			if tt.args.commitmentHashString != "" {
				var err error
				commitmentHash, err = convertStringToByte32(tt.args.commitmentHashString)
				if err != nil {
					t.Errorf("Error in decoding commitmentHashString: %v", err)
					return
				}
			}

			SetUpMockInterfaces()

			utils.MerkleInterface = &utils.MerkleTreeStruct{}
			merkleUtils = utils.MerkleInterface

			utilsMock.On("GetCommitment", mock.Anything, mock.Anything).Return(types.Commitment{CommitmentHash: commitmentHash}, tt.args.commitmentErr)

			got, err := VerifyCommitment(rpcParameters, account, commitmentFetched)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyCommitment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyCommitment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSeed(t *testing.T) {
	var (
		account      types.Account
		keystorePath string
		epoch        uint32
	)
	type args struct {
		secret    string
		secretErr error
		salt      string
		saltErr   error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "When both secret and seed are valid",
			args: args{
				secret: "0f7f6290794dae00bf7c673d36fa2a5b447d2c8c60e9a4220b7ab65be80547a9",
				salt:   "03bceb412a8c973dbb960f1353ba91cf6ca10dfde21c911054cf1e61f0d28e0b",
			},
			want:    "8f81216409d9ecf1fbbab41cc3941c504e5c3b170cb3e4de6477974de4a9fd37",
			wantErr: false,
		},
		{
			name: "When there is an error in getting secret",
			args: args{
				secretErr: errors.New("secret error"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "When there is an error in getting salt",
			args: args{
				secret:  "0f7f6290794dae00bf7c673d36fa2a5b447d2c8c60e9a4220b7ab65be80547a9",
				saltErr: errors.New("salt error"),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var salt [32]byte
			var secret []byte

			if tt.args.secret != "" {
				var err error
				secret, err = hex.DecodeString(tt.args.secret)
				if err != nil {
					t.Errorf("Error in decoding secret: %v", err)
					return
				}
			}
			if tt.args.salt != "" {
				var err error
				salt, err = convertStringToByte32(tt.args.salt)
				if err != nil {
					t.Errorf("Error in decoding salt: %v", err)
					return
				}
			}

			SetUpMockInterfaces()

			cmdUtilsMock.On("CalculateSecret", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, secret, tt.args.secretErr)
			cmdUtilsMock.On("GetSalt", mock.Anything, mock.Anything).Return(salt, tt.args.saltErr)
			got, err := CalculateSeed(rpcParameters, account, keystorePath, epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("CalculateSeed() got = %v, want %v", hex.EncodeToString(got), tt.want)
			}
		})
	}
}

func convertStringToByte32(value string) ([32]byte, error) {
	decodedValue, err := hex.DecodeString(value)
	if err != nil {
		log.Error("Error in decoding string:", err)
		return [32]byte{}, err
	}
	if len(decodedValue) != 32 {
		return [32]byte{}, errors.New("decoded string is not 32 bytes long")
	}
	var decodedValueByte32 [32]byte
	copy(decodedValueByte32[:], decodedValue)
	return decodedValueByte32, nil
}
