package cmd

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"os"
	"path"
	"razor/cmd/mocks"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	mocks2 "razor/utils/mocks"
	"reflect"
	"testing"
)

func TestExecuteVote(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	type args struct {
		config       types.Configurations
		configErr    error
		password     string
		rogueStatus  bool
		rogueErr     error
		rogueMode    []string
		rogueModeErr error
		address      string
		addressErr   error
		voteErr      error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteVote() executes successfully",
			args: args{
				config:      config,
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				rogueStatus: true,
				rogueMode:   []string{"propose", "commit"},
				voteErr:     nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				configErr:   errors.New("config error"),
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				rogueStatus: true,
				rogueMode:   []string{"propose", "commit"},
				voteErr:     nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:      config,
				password:    "test",
				address:     "",
				addressErr:  errors.New("address error"),
				rogueStatus: true,
				rogueMode:   []string{"propose", "commit"},
				voteErr:     nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting rogue status",
			args: args{
				config:    config,
				password:  "test",
				address:   "0x000000000000000000000000000000000000dea1",
				rogueErr:  errors.New("rogue status error"),
				rogueMode: []string{"propose", "commit"},
				voteErr:   nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error in getting rogue modes",
			args: args{
				config:       config,
				password:     "test",
				address:      "0x000000000000000000000000000000000000dea1",
				rogueStatus:  true,
				rogueMode:    nil,
				rogueModeErr: errors.New("rogueModes error"),
				voteErr:      nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from Vote()",
			args: args{
				config:      config,
				password:    "test",
				address:     "0x000000000000000000000000000000000000dea1",
				rogueStatus: true,
				rogueMode:   []string{"propose", "commit"},
				voteErr:     errors.New("vote error"),
			},
			expectedFatal: false,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			flagSetUtilsMock := new(mocks.FlagSetInterface)
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			osMock := new(mocks.OSInterface)

			flagSetUtils = flagSetUtilsMock
			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			osUtils = osMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword").Return(tt.args.password)
			flagSetUtilsMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			flagSetUtilsMock.On("GetBoolRogue", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.rogueStatus, tt.args.rogueErr)
			flagSetUtilsMock.On("GetStringSliceRogueMode", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.rogueMode, tt.args.rogueModeErr)
			cmdUtilsMock.On("HandleExit").Return()
			cmdUtilsMock.On("Vote", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.voteErr)
			osMock.On("Exit", mock.AnythingOfType("int")).Return()

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteVote(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteVote function didn't execute as expected")
			}
		})
	}
}

func TestGetLastProposedEpoch(t *testing.T) {
	var client *ethclient.Client
	blockNumber := big.NewInt(20)

	type args struct {
		fromBlock        *big.Int
		fromBlockErr     error
		stakerId         uint32
		logs             []Types.Log
		logsErr          error
		contractAbi      abi.ABI
		parseErr         error
		unpackedData     []interface{}
		unpackErr        error
		bufferPercent    int32
		bufferPercentErr error
		time             int64
		timeErr          error
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr error
	}{
		{
			name: "Test 1: When GetLastProposedBlock() executes successfully",
			args: args{
				fromBlock: big.NewInt(0),
				stakerId:  2,
				logs: []Types.Log{
					{
						Data:   []byte{4, 2},
						Topics: []common.Hash{common.BigToHash(big.NewInt(1000)), common.BigToHash(big.NewInt(2))},
					},
				},
				contractAbi:   abi.ABI{},
				unpackedData:  convertToSliceOfInterface([]uint32{4, 2}),
				bufferPercent: 1,
				time:          0,
			},
			want:    4,
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting logs",
			args: args{
				logsErr: errors.New("logs error"),
			},
			want:    0,
			wantErr: errors.New("logs error"),
		},
		{
			name: "Test 3: When there is an error in getting contractAbi while parsing",
			args: args{
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				parseErr:     errors.New("parse error"),
				unpackedData: convertToSliceOfInterface([]uint32{4, 2}),
			},
			want:    0,
			wantErr: errors.New("parse error"),
		},
		{
			name: "Test 4: When there is an error in unpacking",
			args: args{
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractAbi: abi.ABI{},
				unpackErr:   errors.New("unpack error"),
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Test 5: When there is an error in fetching blocks",
			args: args{
				fromBlockErr: errors.New("error in fetching blocks"),
			},
			want:    0,
			wantErr: errors.New("Not able to Fetch Block: error in fetching blocks"),
		},
		{
			name: "Test 6: When there is an error in getting bufferPercent",
			args: args{
				fromBlock: big.NewInt(0),
				stakerId:  2,
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractAbi:      abi.ABI{},
				unpackedData:     convertToSliceOfInterface([]uint32{4, 2}),
				bufferPercentErr: errors.New("error in getting buffer percent"),
			},
			want:    0,
			wantErr: errors.New("error in getting buffer percent"),
		},
		{
			name: "Test 7: When there is an error in getting remaining time",
			args: args{
				fromBlock: big.NewInt(0),
				stakerId:  2,
				logs: []Types.Log{
					{
						Data: []byte{4, 2},
					},
				},
				contractAbi:   abi.ABI{},
				unpackedData:  convertToSliceOfInterface([]uint32{4, 2}),
				bufferPercent: 1,
				timeErr:       errors.New("error in getting time"),
			},
			want:    0,
			wantErr: errors.New("error in getting time"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			abiMock := new(mocks.AbiInterface)
			utilsPkgMock := new(mocks2.Utils)
			abiUtilsMock := new(mocks2.ABIUtils)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsMock := new(mocks.UtilsInterface)
			utilsPkgMock2 := new(mocks2.Utils)

			utilsInterface = utilsPkgMock2
			razorUtils = utilsMock
			abiUtils = abiMock
			utils.UtilsInterface = utilsPkgMock
			utils.ABIInterface = abiUtilsMock
			cmdUtils = cmdUtilsMock

			utilsPkgMock.On("CalculateBlockNumberAtEpochBeginning", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.Anything).Return(tt.args.fromBlock, tt.args.fromBlockErr)
			abiUtilsMock.On("Parse", mock.Anything).Return(tt.args.contractAbi, tt.args.parseErr)
			utilsPkgMock.On("FilterLogsWithRetry", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("ethereum.FilterQuery")).Return(tt.args.logs, tt.args.logsErr)
			abiMock.On("Unpack", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unpackedData, tt.args.unpackErr)
			cmdUtilsMock.On("GetBufferPercent").Return(tt.args.bufferPercent, tt.args.bufferPercentErr)
			utilsPkgMock2.On("GetRemainingTimeOfCurrentState", mock.Anything, mock.Anything).Return(tt.args.time, tt.args.timeErr)

			utils := &UtilsStruct{}
			got, err := utils.GetLastProposedEpoch(client, blockNumber, tt.args.stakerId)
			if got != tt.want {
				t.Errorf("GetLastProposedEpoch() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetLastProposedEpoch(), got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetLastProposedEpoch(), got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func convertToSliceOfInterface(arr []uint32) []interface{} {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	return s
}

func TestCalculateSecret(t *testing.T) {
	dir, _ := os.Getwd()
	razorPath := path.Dir(dir)
	testKeystorePath := path.Join(razorPath, "utils/test_accounts")

	type args struct {
		address  string
		password string
		epoch    uint32
		chainId  *big.Int
	}
	tests := []struct {
		name          string
		args          args
		wantSignature string
		wantSecret    string
		wantErr       bool
	}{
		{
			name: "Test 1 - Address 1 with SKALE chainId",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				password: "Test@123",
				epoch:    9021,
				chainId:  big.NewInt(0x785B4B9847B9),
			},
			wantSignature: "be151a0d3890dec990ecc47923df44f1f63e7159db9694712836b75e3f48c95802e096c7554f17e865491a7e01aeffee0e6e20f5e31fa573c6e9640efd1f86ee1b",
			wantSecret:    "0f7f6290794dae00bf7c673d36fa2a5b447d2c8c60e9a4220b7ab65be80547a9",
			wantErr:       false,
		},
		{
			name: "Test 2 - Address 2 with SKALE chainId",
			args: args{
				address:  "0xBd3e0a1d11163934DF10501c9E1a18fbAA9ecAf4",
				password: "Test@123",
				epoch:    9021,
				chainId:  big.NewInt(0x785B4B9847B9),
			},
			wantSignature: "e89df172968b577ab60192503949bb19751bc5d50f4bd11fc98a5b3089e31a945494ac10874254caec540a3d360179c62ac2a0fe8a380bd77ab113925e37024c1b",
			wantSecret:    "b3e7edd43fae5b925a33494f75e1d38484c3c0d8be29b7a8ff71ce17f65fc542",
			wantErr:       false,
		},
		{
			name: "Test 3 - Address 1 with Hardhat chainId",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				epoch:    9021,
				password: "Test@123",
				chainId:  big.NewInt(31337),
			},
			wantSignature: "a98ef5a1cec4e319580acc579b6e56d49158d2f10b66bd6a573861f17b3640ee3a7f720869c48c1c42b4bcb67c2119f0250f8fad7a70ef2de839564166117af31b",
			wantSecret:    "34653d009bf1af9ff85cfd432a1bc6e2128ab307090ff38332ba0909e599c9fa",
			wantErr:       false,
		},
		{
			name: "Test 4 - Address 1 with epoch = 0",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				password: "Test@123",
				epoch:    0,
				chainId:  big.NewInt(31337),
			},
			wantSignature: "761c52de33ed3ae5185e79d872ff42f38dca720ec7ccbe66df4ec188d03448e234a95571b602247f2245da60baa0605a76d680cbc4921117170c9e2e1e673c3e1c",
			wantSecret:    "a64a7ac998067f775a819dff2adc94c5d6427fbb4759cdc4460e69592c5463d8",
			wantErr:       false,
		},
		{
			name: "Test 5 - When address is nil",
			args: args{
				address:  "",
				password: "Test@123",
				epoch:    0,
				chainId:  big.NewInt(31337),
			},
			wantSignature: "",
			wantSecret:    "",
			wantErr:       true,
		},
		{
			name: "Test 6 - When password is wrong",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				password: "Test",
				epoch:    0,
				chainId:  big.NewInt(31337),
			},
			wantSignature: "",
			wantSecret:    "",
			wantErr:       true,
		},
		{
			name: "Test 7 - When ChainId is nil",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				password: "Test@123",
				epoch:    0,
				chainId:  nil,
			},
			wantSignature: "",
			wantSecret:    "",
			wantErr:       true,
		},
		{
			name: "Test 6 - When password is nil",
			args: args{
				address:  "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
				password: "",
				epoch:    9021,
				chainId:  big.NewInt(31337),
			},
			wantSignature: "",
			wantSecret:    "",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeInterfaces()
			gotSignature, gotSecret, err := cmdUtils.CalculateSecret(types.Account{Address: tt.args.address,
				Password: tt.args.password}, tt.args.epoch, testKeystorePath, tt.args.chainId)

			gotSignatureInHash := hex.EncodeToString(gotSignature)
			gotSecretInHash := hex.EncodeToString(gotSecret)
			if !reflect.DeepEqual(gotSignatureInHash, tt.wantSignature) {
				t.Errorf("CalculateSecret() Signature = %v, want %v", gotSignatureInHash, tt.wantSignature)
			}
			if !reflect.DeepEqual(gotSecretInHash, tt.wantSecret) {
				t.Errorf("CalculateSecret() Secret = %v, want %v", gotSecretInHash, tt.wantSecret)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInitiateCommit(t *testing.T) {
	var (
		client    *ethclient.Client
		config    types.Configurations
		account   types.Account
		stakerId  uint32
		rogueData types.Rogue
	)
	type args struct {
		staker                    bindings.StructsStaker
		stakerErr                 error
		minStakeAmount            *big.Int
		minStakeAmountErr         error
		epoch                     uint32
		lastCommit                uint32
		lastCommitErr             error
		secret                    []byte
		secretErr                 error
		signature                 []byte
		salt                      [32]byte
		saltErr                   error
		path                      string
		pathErr                   error
		commitData                types.CommitData
		commitDataErr             error
		merkleTree                [][][]byte
		merkleRoot                [32]byte
		commitTxn                 common.Hash
		commitTxnErr              error
		waitForBlockCompletionErr error
		fileName                  string
		fileNameErr               error
		saveErr                   error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When InitiateCommit executes successfully",
			args: args{
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmount: big.NewInt(100),
				epoch:          5,
				lastCommit:     2,
				signature:      []byte{2},
				secret:         []byte{1},
				salt:           [32]byte{},
				commitData: types.CommitData{
					AssignedCollections:    nil,
					SeqAllottedCollections: nil,
					Leaves:                 nil,
				},
				merkleTree: [][][]byte{},
				commitTxn:  common.BigToHash(big.NewInt(1)),
				fileName:   "",
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting staker",
			args: args{
				stakerErr: errors.New("error in getting staker"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When there is an error in getting minStakeAmount",
			args: args{
				staker:            bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmountErr: errors.New("error in getting minStakeAmount"),
			},
			wantErr: true,
		},
		{
			name: "Test 4: When staker's stake is less than minStakeAmount",
			args: args{
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10)},
				minStakeAmount: big.NewInt(100),
				epoch:          5,
				lastCommit:     2,
				signature:      []byte{2},
				secret:         []byte{1},
				salt:           [32]byte{},
				commitData: types.CommitData{
					AssignedCollections:    nil,
					SeqAllottedCollections: nil,
					Leaves:                 nil,
				},
				merkleTree: [][][]byte{},
				commitTxn:  common.BigToHash(big.NewInt(1)),
				fileName:   "",
			},
			wantErr: false,
		},
		{
			name: "Test 5: When there is an error in getting last commit epoch",
			args: args{
				epoch:         5,
				lastCommitErr: errors.New("error in getting last commit epoch"),
			},
			wantErr: true,
		},
		{
			name: "Test 6: When lastCommittedEpoch is greater or equal to current epoch",
			args: args{
				epoch:      5,
				lastCommit: 6,
			},
			wantErr: false,
		},
		{
			name: "Test 7: When there is an error in getting secret",
			args: args{
				epoch:      5,
				lastCommit: 2,
				secretErr:  errors.New("error in getting secret"),
			},
			wantErr: true,
		},
		{
			name: "Test 8: When there is an error in getting salt",
			args: args{
				epoch:      5,
				lastCommit: 2,
				secret:     []byte{1},
				saltErr:    errors.New("error in getting salt"),
			},
			wantErr: true,
		},
		{
			name: "Test 9: When there is an error in getting commitData",
			args: args{
				epoch:         5,
				lastCommit:    2,
				secret:        []byte{1},
				salt:          [32]byte{},
				commitDataErr: errors.New("error in getting commitData"),
			},
			wantErr: true,
		},
		{
			name: "Test 10: When there is an erron in commit",
			args: args{
				epoch:      5,
				lastCommit: 2,
				secret:     []byte{1},
				salt:       [32]byte{},
				commitData: types.CommitData{
					AssignedCollections:    nil,
					SeqAllottedCollections: nil,
					Leaves:                 nil,
				},
				merkleTree:   [][][]byte{},
				commitTxnErr: errors.New("error in commit"),
			},
			wantErr: true,
		},
		{
			name: "Test 11: When there is an error in getting fileName",
			args: args{
				epoch:      5,
				lastCommit: 2,
				secret:     []byte{1},
				salt:       [32]byte{},
				commitData: types.CommitData{
					AssignedCollections:    nil,
					SeqAllottedCollections: nil,
					Leaves:                 nil,
				},
				fileNameErr: errors.New("error in getting fileName"),
			},
			wantErr: true,
		},
		{
			name: "Test 12: When there is an error in sending commit transaction",
			args: args{
				epoch:      5,
				lastCommit: 2,
				secret:     []byte{1},
				salt:       [32]byte{},
				commitData: types.CommitData{
					AssignedCollections:    nil,
					SeqAllottedCollections: nil,
					Leaves:                 nil,
				},
				merkleTree:                [][][]byte{},
				commitTxn:                 common.BigToHash(big.NewInt(1)),
				waitForBlockCompletionErr: errors.New("transaction mining unsuccessful"),
			},
			wantErr: true,
		},
		{
			name: "Test 13: When there is an error in getting path",
			args: args{
				epoch:      5,
				lastCommit: 2,
				pathErr:    errors.New("path error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			merkleInterface := new(mocks2.MerkleTreeInterface)
			utilsPkgMock := new(mocks2.Utils)

			utils.MerkleInterface = merkleInterface
			utils.UtilsInterface = utilsPkgMock
			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			utilsPkgMock.On("GetMinStakeAmount", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minStakeAmount, tt.args.minStakeAmountErr)
			utilsMock.On("GetEpochLastCommitted", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.lastCommit, tt.args.lastCommitErr)
			cmdUtilsMock.On("CalculateSecret", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.signature, tt.args.secret, tt.args.secretErr)
			cmdUtilsMock.On("GetSalt", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.salt, tt.args.saltErr)
			cmdUtilsMock.On("HandleCommitState", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.commitData, tt.args.commitDataErr)
			merkleInterface.On("CreateMerkle", mock.Anything).Return(tt.args.merkleTree)
			merkleInterface.On("GetMerkleRoot", mock.Anything).Return(tt.args.merkleRoot)
			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			cmdUtilsMock.On("Commit", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.commitTxn, tt.args.commitTxnErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.waitForBlockCompletionErr)
			utilsMock.On("GetCommitDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
			utilsMock.On("SaveDataToCommitJsonFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.saveErr)
			ut := &UtilsStruct{}
			if err := ut.InitiateCommit(client, config, account, tt.args.epoch, stakerId, rogueData); (err != nil) != tt.wantErr {
				t.Errorf("InitiateCommit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitiateReveal(t *testing.T) {
	var (
		client  *ethclient.Client
		config  types.Configurations
		account types.Account
	)

	randomNum := utils.GetRogueRandomValue(10000000)

	type args struct {
		staker                   bindings.StructsStaker
		minStakeAmount           *big.Int
		minStakeAmountErr        error
		epoch                    uint32
		lastReveal               uint32
		lastRevealErr            error
		revealStateErr           error
		fileName                 string
		fileNameErr              error
		committedDataFromFile    types.CommitFileData
		committedDataFromFileErr error
		path                     string
		pathErr                  error
		signature                []byte
		secret                   []byte
		secretErr                error
		revealTxn                common.Hash
		revealTxnErr             error
		rogueData                types.Rogue
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When InitiateReveal executes successfully",
			args: args{
				staker:                bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmount:        big.NewInt(100),
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 5},
				signature:             []byte{1},
				secret:                []byte{},
				revealTxn:             common.BigToHash(big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting minStakeAmount",
			args: args{
				staker:            bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmountErr: errors.New("error in getting minStakeAmount"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When staker's stake is less than minStakeAmount",
			args: args{
				staker:                bindings.StructsStaker{Id: 1, Stake: big.NewInt(10)},
				minStakeAmount:        big.NewInt(100),
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 5},
				signature:             []byte{1},
				secret:                []byte{},
				revealTxn:             common.BigToHash(big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting lastReveal",
			args: args{
				epoch:         5,
				lastRevealErr: errors.New("error in getting lastReveal"),
			},
			wantErr: true,
		},
		{
			name: "Test 4: When lastRevealedEpoch is greater or equal to current epoch",
			args: args{
				epoch:      5,
				lastReveal: 6,
			},
			wantErr: false,
		},
		{
			name: "Test 5: When there is an error in handleRevealState",
			args: args{
				epoch:          5,
				lastReveal:     2,
				revealStateErr: errors.New("error in handleRevealState"),
			},
			wantErr: true,
		},
		{
			name: "Test 6: When there is an error in getting fileName",
			args: args{
				epoch:       5,
				lastReveal:  2,
				fileNameErr: errors.New("error in getting fileName"),
			},
			wantErr: true,
		},
		{
			name: "Test 7: When there is an error in getting data from file",
			args: args{
				epoch:                    5,
				lastReveal:               2,
				fileName:                 "",
				committedDataFromFileErr: errors.New("error in reading data from file"),
			},
			wantErr: true,
		},
		{
			name: "Test 8: When file does not contain the latest data",
			args: args{
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 3},
			},
			wantErr: true,
		},
		{
			name: "Test 9: When there is an error in getting secret",
			args: args{
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 5},
				secretErr:             errors.New("error in getting secret"),
			},
			wantErr: true,
		},
		{
			name: "Test 10: When there is an error in reveal",
			args: args{
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 5},
				secret:                []byte{},
				revealTxnErr:          errors.New("error in reveal"),
			},
			wantErr: true,
		},
		{
			name: "Test 11: When InitiateReveal executes successfully and rogueMode is in reveal",
			args: args{
				epoch: 5,
				rogueData: types.Rogue{
					IsRogue:   true,
					RogueMode: []string{"reveal"},
				},
				lastReveal: 2,
				fileName:   "",
				committedDataFromFile: types.CommitFileData{
					Epoch:  5,
					Leaves: []*big.Int{big.NewInt(1), big.NewInt(2)},
				},
				secret:    []byte{},
				revealTxn: common.BigToHash(big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test 12: When there is an error in getting path",
			args: args{
				epoch:                 5,
				lastReveal:            2,
				fileName:              "",
				committedDataFromFile: types.CommitFileData{Epoch: 5},
				pathErr:               errors.New("path error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(mocks2.Utils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			utils.UtilsInterface = utilsPkgMock

			utilsPkgMock.On("GetMinStakeAmount", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minStakeAmount, tt.args.minStakeAmountErr)
			utilsMock.On("GetEpochLastRevealed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.lastReveal, tt.args.lastRevealErr)
			cmdUtilsMock.On("HandleRevealState", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.AnythingOfType("uint32")).Return(tt.args.revealStateErr)
			utilsMock.On("GetCommitDataFileName", mock.AnythingOfType("string")).Return(tt.args.fileName, tt.args.fileNameErr)
			utilsMock.On("ReadFromCommitJsonFile", mock.Anything).Return(tt.args.committedDataFromFile, tt.args.committedDataFromFileErr)
			utilsMock.On("GetRogueRandomValue", mock.AnythingOfType("int")).Return(randomNum)
			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			cmdUtilsMock.On("CalculateSecret", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.signature, tt.args.secret, tt.args.secretErr)
			cmdUtilsMock.On("Reveal", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.revealTxn, tt.args.revealTxnErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(nil)
			ut := &UtilsStruct{}
			if err := ut.InitiateReveal(client, config, account, tt.args.epoch, tt.args.staker, tt.args.rogueData); (err != nil) != tt.wantErr {
				t.Errorf("InitiateReveal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitiatePropose(t *testing.T) {
	var (
		client      *ethclient.Client
		config      types.Configurations
		account     types.Account
		blockNumber *big.Int
		rogueData   types.Rogue
	)
	type args struct {
		staker            bindings.StructsStaker
		minStakeAmount    *big.Int
		minStakeAmountErr error
		epoch             uint32
		lastProposal      uint32
		lastProposalErr   error
		lastReveal        uint32
		lastRevealErr     error
		proposeTxn        common.Hash
		proposeTxnErr     error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When InitiatePropose executes successfully",
			args: args{
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmount: big.NewInt(100),
				epoch:          5,
				lastProposal:   4,
				lastReveal:     6,
				proposeTxn:     common.BigToHash(big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting minStakeAmount",
			args: args{
				staker:            bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				minStakeAmountErr: errors.New("error in getting minStakeAmount"),
			},
			wantErr: true,
		},
		{
			name: "Test 3: When staker's stake is less than minStakeAmount",
			args: args{
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10)},
				minStakeAmount: big.NewInt(100),
				epoch:          5,
				lastProposal:   4,
				lastReveal:     6,
				proposeTxn:     common.BigToHash(big.NewInt(1)),
			},
			wantErr: false,
		},
		{
			name: "Test 4: When there is an error in getting last proposed epoch",
			args: args{
				epoch:           5,
				lastProposalErr: errors.New("error in getting last proposed epoch"),
			},
			wantErr: true,
		},
		{
			name: "Test 5: When last proposed epoch is greater then current epoch",
			args: args{
				epoch:        5,
				lastProposal: 6,
			},
			wantErr: false,
		},
		{
			name: "Test 6: When last revealed epoch is less than current epoch",
			args: args{
				epoch:        5,
				lastProposal: 4,
				lastReveal:   3,
			},
			wantErr: false,
		},
		{
			name: "Test 7: When there is an error in getting last revealed epoch",
			args: args{
				epoch:         5,
				lastProposal:  4,
				lastRevealErr: errors.New("error in getting last revealed epoch"),
			},
			wantErr: true,
		},
		{
			name: "Test 8: When there is an error in propose",
			args: args{
				epoch:         5,
				lastProposal:  4,
				lastReveal:    6,
				proposeTxnErr: errors.New("error in propose"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(mocks2.Utils)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			utils.UtilsInterface = utilsPkgMock

			utilsPkgMock.On("GetMinStakeAmount", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minStakeAmount, tt.args.minStakeAmountErr)
			cmdUtilsMock.On("GetLastProposedEpoch", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("uint32")).Return(tt.args.lastProposal, tt.args.lastProposalErr)
			utilsMock.On("GetEpochLastRevealed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.lastReveal, tt.args.lastRevealErr)
			cmdUtilsMock.On("Propose", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.proposeTxn, tt.args.proposeTxnErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(nil)
			ut := &UtilsStruct{}
			if err := ut.InitiatePropose(client, config, account, tt.args.epoch, tt.args.staker, blockNumber, rogueData); (err != nil) != tt.wantErr {
				t.Errorf("InitiatePropose() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandleBlock(t *testing.T) {
	var (
		client      *ethclient.Client
		account     types.Account
		blockNumber *big.Int
		rogueData   types.Rogue
	)

	type args struct {
		config               types.Configurations
		state                int64
		stateErr             error
		epoch                uint32
		epochErr             error
		stateName            string
		stakerId             uint32
		stakerIdErr          error
		staker               bindings.StructsStaker
		stakerErr            error
		ethBalance           *big.Int
		ethBalanceErr        error
		actualStake          *big.Float
		actualStakeErr       error
		actualBalance        *big.Float
		sRZRBalance          *big.Int
		sRZRBalanceErr       error
		sRZRInEth            *big.Float
		initiateCommitErr    error
		initiateRevealErr    error
		initiateProposeErr   error
		handleDisputeErr     error
		claimBlockRewardTxn  common.Hash
		claimBlockRewardErr  error
		lastVerification     uint32
		isFlagPassed         bool
		handleClaimBountyErr error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1: When HandleBlock executes successfully and state is commit",
			args: args{
				state:         0,
				epoch:         1,
				stateName:     "commit",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:    big.NewInt(1000),
				actualStake:   big.NewFloat(10000),
				actualBalance: big.NewFloat(1000),
				sRZRBalance:   big.NewInt(10000),
				sRZRInEth:     big.NewFloat(100),
			},
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr: errors.New("error in getting state"),
			},
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				state:    0,
				epochErr: errors.New("error in getting epoch"),
			},
		},
		{
			name: "Test 4: When there is an error in getting stakerId",
			args: args{
				state:       0,
				epoch:       1,
				stakerIdErr: errors.New("error in getting stakerId"),
			},
		},
		{
			name: "Test 5: When stakerId is 0",
			args: args{
				state:     0,
				epoch:     1,
				stateName: "commit",
				stakerId:  0,
			},
		},
		{
			name: "Test 6: When there is an error in getting staker",
			args: args{
				state:     0,
				epoch:     1,
				stateName: "commit",
				stakerId:  1,
				stakerErr: errors.New("error in getting staker"),
			},
		},
		{
			name: "Test 7: When there is an error in getting ethBalance",
			args: args{
				state:         0,
				epoch:         1,
				stateName:     "commit",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalanceErr: errors.New("error in getting ethBalance"),
			},
		},
		{
			name: "Test 8: When there is error in converting stakedAmount",
			args: args{
				state:          0,
				epoch:          1,
				stateName:      "commit",
				stakerId:       1,
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:     big.NewInt(1000),
				actualStakeErr: errors.New("error in converting stakedAmount"),
			},
		},

		{
			name: "Test 9: When there is an error in getting sRZR Balance",
			args: args{
				state:          0,
				epoch:          1,
				stateName:      "commit",
				stakerId:       1,
				staker:         bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:     big.NewInt(1000),
				actualStake:    big.NewFloat(10000),
				actualBalance:  big.NewFloat(1000),
				sRZRBalanceErr: errors.New("error in getting sRZR Balance"),
			},
		},
		{
			name: "Test 10: When stake is less than the misStake",
			args: args{
				state:         0,
				epoch:         1,
				stateName:     "commit",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(100)},
				ethBalance:    big.NewInt(1000),
				actualStake:   big.NewFloat(100),
				actualBalance: big.NewFloat(1000),
				sRZRBalance:   big.NewInt(100),
				sRZRInEth:     big.NewFloat(100),
			},
		},
		{
			name: "Test 11: When stake has already been withdrwan by staker",
			args: args{
				state:         0,
				epoch:         1,
				stateName:     "commit",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(0)},
				ethBalance:    big.NewInt(1000),
				actualStake:   big.NewFloat(0),
				actualBalance: big.NewFloat(1000),
				sRZRBalance:   big.NewInt(0),
				sRZRInEth:     big.NewFloat(100),
			},
		},
		{
			name: "Test 12: When staker is already slashed",
			args: args{
				state:         0,
				epoch:         1,
				stateName:     "commit",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000), IsSlashed: true},
				ethBalance:    big.NewInt(1000),
				actualStake:   big.NewFloat(10000),
				actualBalance: big.NewFloat(1000),
				sRZRBalance:   big.NewInt(10000),
				sRZRInEth:     big.NewFloat(100),
			},
		},
		{
			name: "Test 13: When there is an error in initiateCommit",
			args: args{
				state:             0,
				epoch:             1,
				stateName:         "commit",
				stakerId:          1,
				staker:            bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:        big.NewInt(1000),
				actualStake:       big.NewFloat(10000),
				actualBalance:     big.NewFloat(1000),
				sRZRBalance:       big.NewInt(10000),
				sRZRInEth:         big.NewFloat(100),
				initiateCommitErr: errors.New("error in initiateCommit"),
			},
		},
		{
			name: "Test 14: When there is an error in initiateReveal",
			args: args{
				state:             1,
				epoch:             1,
				stateName:         "reveal",
				stakerId:          1,
				staker:            bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:        big.NewInt(1000),
				actualStake:       big.NewFloat(10000),
				actualBalance:     big.NewFloat(1000),
				sRZRBalance:       big.NewInt(10000),
				sRZRInEth:         big.NewFloat(100),
				initiateRevealErr: errors.New("error in initiateReveal"),
			},
		},
		{
			name: "Test 15: When there is an error in initiatePropose",
			args: args{
				state:              2,
				epoch:              1,
				stateName:          "propose",
				stakerId:           1,
				staker:             bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:         big.NewInt(1000),
				actualStake:        big.NewFloat(10000),
				actualBalance:      big.NewFloat(1000),
				sRZRBalance:        big.NewInt(10000),
				sRZRInEth:          big.NewFloat(100),
				initiateProposeErr: errors.New("error in initiatePropose"),
			},
		},
		{
			name: "Test 16: When there is an error in handleDispute",
			args: args{
				state:            3,
				epoch:            1,
				stateName:        "dispute",
				stakerId:         1,
				staker:           bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:       big.NewInt(1000),
				actualStake:      big.NewFloat(10000),
				actualBalance:    big.NewFloat(1000),
				sRZRBalance:      big.NewInt(10000),
				sRZRInEth:        big.NewFloat(100),
				handleDisputeErr: errors.New("error in handleDispute"),
			},
		},
		{
			name: "Test 17: When there is no error in dispute and HandleClaimBounty flag is passed",
			args: args{
				state:         3,
				epoch:         1,
				stateName:     "dispute",
				stakerId:      1,
				staker:        bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:    big.NewInt(1000),
				actualStake:   big.NewFloat(10000),
				actualBalance: big.NewFloat(1000),
				sRZRBalance:   big.NewInt(10000),
				sRZRInEth:     big.NewFloat(100),
				isFlagPassed:  true,
			},
		},
		{
			name: "Test 18: When there is no error in dispute but HandleClaimBounty throws error",
			args: args{
				state:                3,
				epoch:                1,
				stateName:            "dispute",
				stakerId:             1,
				staker:               bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:           big.NewInt(1000),
				actualStake:          big.NewFloat(10000),
				actualBalance:        big.NewFloat(1000),
				sRZRBalance:          big.NewInt(10000),
				sRZRInEth:            big.NewFloat(100),
				isFlagPassed:         true,
				handleClaimBountyErr: errors.New("error in handleClaimBounty"),
			},
		},
		{
			name: "Test 19: When claimBlockReward executes successfully in confirm state",
			args: args{
				state:               4,
				epoch:               1,
				stateName:           "confirm",
				lastVerification:    1,
				stakerId:            1,
				staker:              bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:          big.NewInt(1000),
				actualStake:         big.NewFloat(10000),
				actualBalance:       big.NewFloat(1000),
				sRZRBalance:         big.NewInt(10000),
				sRZRInEth:           big.NewFloat(100),
				claimBlockRewardTxn: common.BigToHash(big.NewInt(1)),
			},
		},
		{
			name: "Test 20: When there is an error in claimBlockReward",
			args: args{
				state:               4,
				epoch:               2,
				stateName:           "confirm",
				lastVerification:    1,
				stakerId:            2,
				staker:              bindings.StructsStaker{Id: 2, Stake: big.NewInt(10000)},
				ethBalance:          big.NewInt(1000),
				actualStake:         big.NewFloat(10000),
				actualBalance:       big.NewFloat(1000),
				sRZRBalance:         big.NewInt(10000),
				sRZRInEth:           big.NewFloat(100),
				claimBlockRewardErr: errors.New("error in claimBlockReward"),
			},
		},
		{
			name: "Test 21: When lastVerification is greater than the current epoch in dispute state",
			args: args{
				state:            3,
				epoch:            1,
				lastVerification: 4,
				stateName:        "dispute",
				stakerId:         1,
				staker:           bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:       big.NewInt(1000),
				actualStake:      big.NewFloat(10000),
				actualBalance:    big.NewFloat(1000),
				sRZRBalance:      big.NewInt(10000),
				sRZRInEth:        big.NewFloat(100),
			},
		},
		{
			name: "Test 22: When waitTime is more than 5 in -1 state",
			args: args{
				state:            -1,
				epoch:            1,
				lastVerification: 4,
				stateName:        "",
				stakerId:         1,
				staker:           bindings.StructsStaker{Id: 1, Stake: big.NewInt(10000)},
				ethBalance:       big.NewInt(1000),
				actualStake:      big.NewFloat(10000),
				actualBalance:    big.NewFloat(1000),
				sRZRBalance:      big.NewInt(10000),
				sRZRInEth:        big.NewFloat(100),
				config:           types.Configurations{WaitTime: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsPkgMock := new(mocks2.Utils)
			osMock := new(mocks.OSInterface)
			timeMock := new(mocks.TimeInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock
			utils.UtilsInterface = utilsPkgMock
			utilsInterface = utilsPkgMock
			osUtils = osMock
			timeUtils = timeMock

			utilsMock.On("GetDelayedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
			utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
			utilsPkgMock.On("BalanceAtWithRetry", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.ethBalance, tt.args.ethBalanceErr)
			utilsMock.On("ConvertWeiToEth", mock.AnythingOfType("*big.Int")).Return(tt.args.actualStake, tt.args.actualStakeErr)
			utilsMock.On("GetStakerSRZRBalance", mock.Anything, mock.Anything).Return(tt.args.sRZRBalance, tt.args.sRZRBalanceErr)
			utilsPkgMock.On("GetStateName", mock.AnythingOfType("int64")).Return(tt.args.stateName)
			osMock.On("Exit", mock.AnythingOfType("int")).Return()
			cmdUtilsMock.On("InitiateCommit", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.initiateCommitErr)
			cmdUtilsMock.On("InitiateReveal", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.initiateRevealErr)
			cmdUtilsMock.On("InitiatePropose", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.initiateProposeErr)
			cmdUtilsMock.On("HandleDispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.handleDisputeErr)
			utilsPkgMock.On("IsFlagPassed", mock.AnythingOfType("string")).Return(tt.args.isFlagPassed)
			cmdUtilsMock.On("HandleClaimBounty", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.handleClaimBountyErr)
			cmdUtilsMock.On("ClaimBlockReward", mock.Anything).Return(tt.args.claimBlockRewardTxn, tt.args.claimBlockRewardErr)
			utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(nil)
			timeMock.On("Sleep", mock.Anything).Return()
			utilsMock.On("WaitTillNextNSecs", mock.AnythingOfType("int32")).Return()
			lastVerification = tt.args.lastVerification
			ut := &UtilsStruct{}
			ut.HandleBlock(client, account, blockNumber, tt.args.config, rogueData)
		})
	}
}
