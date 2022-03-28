package cmd

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/accounts"
	accountMocks "razor/accounts/mocks"
	"razor/cmd/mocks"
	"razor/core/types"
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

			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
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

//func TestHandleBlock(t *testing.T) {
//
//	var client *ethclient.Client
//	var account types.Account
//	var blockNumber *big.Int
//
//	//randomNum := utils.GetRogueRandomValue(10000000)
//
//	type args struct {
//		config               types.Configurations
//		rogueData            types.Rogue
//		state                int64
//		stateName            string
//		stateErr             error
//		epoch                uint32
//		epochErr             error
//		stakerId             uint32
//		stakerIdErr          error
//		ethBalance           *big.Int
//		actualStake          *big.Float
//		actualStakeErr       error
//		ethBalanceErr        error
//		actualBalance        *big.Float
//		actualBalanceErr     error
//		minStake             *big.Int
//		minStakeErr          error
//		staker               bindings.StructsStaker
//		stakerErr            error
//		sRZR                 *big.Int
//		sRZRErr              error
//		commitTxn            common.Hash
//		revealTxn            common.Hash
//		proposeTxn           common.Hash
//		handleDisputeErr     error
//		blockConfirmed       uint32
//		claimBlockRewardHash common.Hash
//		claimBlockRewardErr  error
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr error
//	}{
//		{
//			name: "Test 1: When the state is commit and there is an error to save committed date to file",
//			args: args{
//				rogueData: types.Rogue{
//					IsRogue: false,
//				},
//				state:       0,
//				stateName:   "commit",
//				epoch:       5,
//				stakerId:    2,
//				actualStake: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//				//stake:         big.NewInt(2000),
//				ethBalance:    big.NewInt(1),
//				actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//				minStake:      big.NewInt(1000),
//				staker: bindings.StructsStaker{
//					IsSlashed: false,
//				},
//				//epochLastCommitted: 4,
//				//secret:             []byte{1, 2, 3},
//				//commitData:         []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//				commitTxn: common.BigToHash(big.NewInt(1)),
//				//commitFile:         "commit.txt",
//				//saveDataErr:        errors.New("error while saving committed data in file"),
//			},
//			wantErr: nil,
//		},
//{
//	name: "Test 2: When there is an error in getting state",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		stateErr: errors.New("state error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 3: When there is an error in getting epoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:     0,
//		stateName: "commit",
//		epochErr:  errors.New("epoch error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 4: When there is an error in getting stakerId",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:       0,
//		stateName:   "commit",
//		epoch:       5,
//		stakerIdErr: errors.New("stakerId error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 5: When there is an error in getting ethBalance",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalanceErr: errors.New("ethBalance error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 6: When there is an error in getting min stake",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStakeErr:   errors.New("min stake error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 7: When there is an error in getting staker",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		stakerErr:     errors.New("staker error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 8: When there is an error in getting actual balance",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:            0,
//		stateName:        "commit",
//		epoch:            5,
//		stakerId:         2,
//		stake:            big.NewInt(2000),
//		ethBalance:       big.NewInt(1),
//		actualBalanceErr: errors.New("converting to eth error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 9: When the state is commit and there is an error in getting epoch last committed",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommittedErr: errors.New("epoch last committed error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 10: When the state is commit and secret is nil",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 4,
//		secret:             nil,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 11: When the state is commit and lastEpochCommited >= epoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 6,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 12: When the state is commit and there is an error from HandleCommitState()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 4,
//		secret:             []byte{1, 2, 3},
//		commitDataErr:      errors.New("handle commit state error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 13: When the state is commit and there is an error from Commit()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 4,
//		secret:             []byte{1, 2, 3},
//		commitData:         []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		commitErr:          errors.New("commit error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 14: When the state is commit and there is an error in getting file name where commit data needs to be saved",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 4,
//		secret:             []byte{1, 2, 3},
//		commitData:         []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		commitTxn:          common.BigToHash(big.NewInt(1)),
//		commitFileErr:      errors.New("error in getting file name"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 15: When the state is Reveal and there is an error in getting epoch last revealed",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealedErr: errors.New("epoch last revealed error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 16: When the state is Reveal and lastEpochRevealed  >= epoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed: 6,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 17: When the state is Reveal and secret is nil",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed: 4,
//		commitData:        []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:            nil,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 18: When the state is Reveal and there is an error from HandleRevealState()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed:    4,
//		commitData:           []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:               []byte{1, 2, 3},
//		handleRevealStateErr: errors.New("handle reveal state error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 19: When the state is Reveal and there is an error from Reveal()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed:    4,
//		commitData:           []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:               []byte{1, 2, 3},
//		handleRevealStateErr: nil,
//		revealErr:            errors.New("reveal error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 20: When the state is Reveal and committed data is nil",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed:  4,
//		commitData:         nil,
//		commitFile:         "commit.txt",
//		epochInFile:        5,
//		commitedDataInFile: []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:             nil,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 21: When the state is Reveal and committed data is nil and there is an error in getting committed data file name",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed: 4,
//		commitData:        nil,
//		commitFileErr:     errors.New("error in getting committed data file name"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 22: When the state is Reveal and committed data is nil and there is an error in getting committed data from file",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed: 4,
//		commitData:        nil,
//		commitFile:        "commit.txt",
//		readFileErr:       errors.New("error in getting committed data from file"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 23: When the state is Reveal and committed data is nil amd epochInfile != currentEpoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed:  4,
//		commitData:         nil,
//		commitFile:         "commit.txt",
//		epochInFile:        4,
//		commitedDataInFile: []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:             nil,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 24: When the state is Reveal and rogue mode is ON for reveal state",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue:   true,
//			RogueMode: []string{"reveal"},
//		},
//		state:         1,
//		stateName:     "reveal",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastRevealed:    4,
//		commitFile:           "commit.txt",
//		commitData:           []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		secret:               []byte{1, 2, 3},
//		handleRevealStateErr: nil,
//		revealTxn:            common.BigToHash(big.NewInt(1)),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 25: When the state is propose and there is an error in getting epoch last proposed",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposedErr: errors.New("epoch last proposed error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 26: When the state is propose and epochLastProposed >= currentEpoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposed: 5,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 27: When the state is propose and there is an error in getting epochLastRevealed",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposed:    4,
//		epochLastRevealedErr: errors.New("epoch last revealed error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 28: When the state is propose and lastReveal < epoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposed: 4,
//		epochLastRevealed: 4,
//	},
//},
//{
//	name: "Test 29: When the state is propose and there is an error from Propose()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposed: 4,
//		epochLastRevealed: 5,
//		proposeErr:        errors.New("propose error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 30: When the state is dispute and lastVerification >= epoch",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         3,
//		stateName:     "dispute",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		lastVerification: 5,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 31: When the state is dispute and there is an error from HandleDispute()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         3,
//		stateName:     "dispute",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		lastVerification: 4,
//		handleDisputeErr: errors.New("handle dispute error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 32: When the state is confirm and there is an error from ClaimBlockReward()",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         4,
//		stateName:     "confirm",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		lastVerification:    5,
//		blockConfirmed:      4,
//		claimBlockRewardErr: errors.New("claimBlockReward error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 33: When the state is -1 and config.waitTime > 5",
//	args: args{
//		config: types.Configurations{
//			WaitTime: 6,
//		},
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         -1,
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 34: When there is no error in confirm state",
//	args: args{
//		config: types.Configurations{
//			WaitTime: 6,
//		},
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		stateName:     "confirm",
//		state:         4,
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		lastVerification:     5,
//		blockConfirmed:       4,
//		claimBlockRewardHash: common.BigToHash(big.NewInt(1)),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 35: When there is no error in commit",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastCommitted: 4,
//		secret:             []byte{1, 2, 3},
//		commitData:         []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
//		commitTxn:          common.BigToHash(big.NewInt(1)),
//		commitFile:         "commit.txt",
//		saveDataErr:        nil,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 36: When there is no error in propose state",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         2,
//		stateName:     "propose",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(1000),
//		staker: bindings.StructsStaker{
//			IsSlashed: false,
//		},
//		epochLastProposed: 4,
//		epochLastRevealed: 5,
//		proposeTxn:        common.BigToHash(big.NewInt(1)),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 37: When stakerId is 0 ",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:     2,
//		stateName: "propose",
//		epoch:     5,
//		stakerId:  0,
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 38: When there is an error in getting stake ",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:     2,
//		stateName: "propose",
//		epoch:     5,
//		stakerId:  2,
//		stakeErr:  errors.New("stake error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 39: When stakedAmount < minStake",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(2000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(10000),
//		stakerErr:     errors.New("staker error"),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 39: When stakedAmount is 0 ",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(0),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(10000),
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 40: When staker is slashed",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(0),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(10000),
//		staker: bindings.StructsStaker{
//			IsSlashed: true,
//		},
//	},
//	wantErr: nil,
//},
//{
//	name: "Test 41: When there is an error in getting sRZR balance",
//	args: args{
//		rogueData: types.Rogue{
//			IsRogue: false,
//		},
//		state:         0,
//		stateName:     "commit",
//		epoch:         5,
//		stakerId:      2,
//		stake:         big.NewInt(10000),
//		ethBalance:    big.NewInt(1),
//		actualBalance: new(big.Float).SetInt(big.NewInt(1)).Quo(big.NewFloat(1), big.NewFloat(1e18)).SetPrec(32),
//		minStake:      big.NewInt(10000),
//		sRZRErr:       errors.New("sRZR error"),
//	},
//	wantErr: nil,
//},
//}
//for _, tt := range tests {
//t.Run(tt.name, func (t *testing.T) {
//
//utilsMock := new(mocks.UtilsInterface)
//cmdUtilsMock := new(mocks.UtilsCmdInterface)
//voteManagerUtilsMock := new(mocks.VoteManagerInterface)
//utilsPkgMock := new(Mocks.Utils)
//timeMock := new(mocks.TimeInterface)
//osMock := new(mocks.OSInterface)
//
//razorUtils = utilsMock
//cmdUtils = cmdUtilsMock
//voteManagerUtils = voteManagerUtilsMock
//timeUtils = timeMock
//osUtils = osMock
//utils.UtilsInterface = utilsPkgMock
//
//utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
//utilsMock.On("GetDelayedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
//utilsMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
//utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
//utilsPkgMock.On("BalanceAtWithRetry", mock.AnythingOfType("*ethclient.Client"), mock.Anything).Return(tt.args.ethBalance, tt.args.ethBalanceErr)
//utilsPkgMock.On("GetMinStakeAmount", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.minStake, tt.args.minStakeErr)
//utilsMock.On("ConvertWeiToEth", mock.AnythingOfType("*big.Int")).Return(tt.args.actualBalance, tt.args.actualBalanceErr)
//utilsMock.On("GetStateName", mock.AnythingOfType("int64")).Return(tt.args.stateName)
//cmdUtilsMock.On("AutoUnstakeAndWithdraw", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
////utilsMock.On("GetStaker", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.staker, tt.args.stakerErr)
//utilsMock.On("GetStakerSRZRBalance", mock.Anything, mock.Anything).Return(tt.args.sRZR, tt.args.sRZRErr)
////utilsMock.On("GetEpochLastCommitted", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastCommitted, tt.args.epochLastCommittedErr)
////cmdUtilsMock.On("CalculateSecret", mock.Anything, mock.Anything).Return(tt.args.secret)
////cmdUtilsMock.On("GetCommitDataFileName", mock.AnythingOfType("string")).Return(tt.args.commitFile, tt.args.commitFileErr)
////cmdUtilsMock.On("HandleCommitState", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.commitData, tt.args.commitDataErr)
//cmdUtilsMock.On("InitiateCommit", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.commitTxn)
//utilsMock.On("WaitForBlockCompletion", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(1)
////utilsMock.On("SaveDataToFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.saveDataErr)
////utilsMock.On("GetEpochLastRevealed", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastRevealed, tt.args.epochLastRevealedErr)
////utilsMock.On("ReadDataFromFile", mock.AnythingOfType("string")).Return(tt.args.epochInFile, tt.args.commitedDataInFile, tt.args.readFileErr)
////cmdUtilsMock.On("HandleRevealState", mock.AnythingOfType("*ethclient.Client"), mock.Anything, mock.AnythingOfType("uint32")).Return(tt.args.handleRevealStateErr)
//cmdUtilsMock.On("InitiateReveal", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.revealTxn)
////utilsMock.On("GetRogueRandomValue", mock.AnythingOfType("int")).Return(randomNum)
//utilsMock.On("WaitTillNextNSecs", mock.AnythingOfType("int32")).Return()
////cmdUtilsMock.On("GetLastProposedEpoch", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastProposed, tt.args.epochLastProposedErr)
//cmdUtilsMock.On("InitiatePropose", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.proposeTxn)
//cmdUtilsMock.On("HandleDispute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.handleDisputeErr)
//cmdUtilsMock.On("ClaimBlockReward", mock.Anything).Return(tt.args.claimBlockRewardHash, tt.args.claimBlockRewardErr)
//timeMock.On("Sleep", mock.AnythingOfType("time.Duration")).Return()
//osMock.On("Exit", mock.AnythingOfType("int")).Return()
//
////_committedData = tt.args.commitData
////lastVerification = tt.args.lastVerification
//blockConfirmed = tt.args.blockConfirmed
//
//utils := &UtilsStruct{}
//utils.HandleBlock(client, account, blockNumber, tt.args.config, tt.args.rogueData)
//
//})
//}
//}

func TestAutoUnstakeAndWithdraw(t *testing.T) {
	var client *ethclient.Client
	var account types.Account
	var amount *big.Int
	var config types.Configurations
	var hash common.Hash

	type args struct {
		stakerId        uint32
		stakerIdErr     error
		unstakeErr      error
		autoWithdrawErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When AutoUnstakeAndWithdraw() executes successfully",
			args: args{
				stakerId:        2,
				unstakeErr:      nil,
				autoWithdrawErr: nil,
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in gettin stakerId",
			args: args{
				stakerIdErr:     errors.New("stakerId error"),
				unstakeErr:      nil,
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error from Unstake()",
			args: args{
				stakerId:        2,
				unstakeErr:      errors.New("unstake error"),
				autoWithdrawErr: nil,
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error from AutoWithdraw()",
			args: args{
				stakerId:        2,
				unstakeErr:      nil,
				autoWithdrawErr: errors.New("autoWithdraw error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cmdUtilsMock := new(mocks.UtilsCmdInterface)
			utilsMock := new(mocks.UtilsInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("GetStakerId", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.stakerId, tt.args.stakerIdErr)
			cmdUtilsMock.On("Unstake", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(hash, tt.args.unstakeErr)
			cmdUtilsMock.On("AutoWithdraw", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.autoWithdrawErr)

			utils := &UtilsStruct{}
			fatal = false

			utils.AutoUnstakeAndWithdraw(client, account, amount, config)
			if fatal != tt.expectedFatal {
				t.Error("The AutoUnstakeAndWithdraw function didn't execute as expected")
			}
		})
	}
}

//func TestGetLastProposedEpoch(t *testing.T) {
//	var client *ethclient.Client
//	blockNumber := big.NewInt(20)
//
//	type args struct {
//		stakerId     uint32
//		logs         []Types.Log
//		logsErr      error
//		contractAbi  abi.ABI
//		parseErr     error
//		unpackedData []interface{}
//		unpackErr    error
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    uint32
//		wantErr error
//	}{
//		{
//			name: "Test 1: When GetLastProposedBlock() executes successfully",
//			args: args{
//				stakerId: 2,
//				logs: []Types.Log{
//					{
//						Data: []byte{4, 2},
//					},
//				},
//				contractAbi:  abi.ABI{},
//				unpackedData: convertToSliceOfInterface([]uint32{4, 2}),
//			},
//			want:    4,
//			wantErr: nil,
//		},
//		{
//			name: "Test 2: When there is an error in getting logs",
//			args: args{
//				stakerId: 2,
//				logsErr:  errors.New("logs error"),
//			},
//			want:    0,
//			wantErr: errors.New("logs error"),
//		},
//		{
//			name: "Test 3: When there is an error in getting contractAbi while parsing",
//			args: args{
//				stakerId: 2,
//				logs: []Types.Log{
//					{
//						Data: []byte{4, 2},
//					},
//				},
//				parseErr:     errors.New("parse error"),
//				unpackedData: convertToSliceOfInterface([]uint32{4, 2}),
//			},
//			want:    0,
//			wantErr: errors.New("parse error"),
//		},
//		{
//			name: "Test 4: When there is an error in unpacking",
//			args: args{
//				stakerId: 2,
//				logs: []Types.Log{
//					{
//						Data: []byte{4, 2},
//					},
//				},
//				contractAbi: abi.ABI{},
//				unpackErr:   errors.New("unpack error"),
//			},
//			want:    0,
//			wantErr: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			abiMock := new(mocks.AbiInterface)
//			utilsPkgMock := new(Mocks.Utils)
//			abiUtilsMock := new(Mocks.ABIUtils)
//
//			abiUtils = abiMock
//			utils.UtilsInterface = utilsPkgMock
//			utils.ABIInterface = abiUtilsMock
//
//			abiUtilsMock.On("Parse", mock.Anything).Return(tt.args.contractAbi, tt.args.parseErr)
//			utilsPkgMock.On("FilterLogsWithRetry", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("ethereum.FilterQuery")).Return(tt.args.logs, tt.args.logsErr)
//			abiMock.On("Unpack", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.unpackedData, tt.args.unpackErr)
//
//			utils := &UtilsStruct{}
//			got, err := utils.GetLastProposedEpoch(client, blockNumber, tt.args.stakerId)
//			if got != tt.want {
//				t.Errorf("GetLastProposedEpoch() got = %v, want %v", got, tt.want)
//			}
//			if err == nil || tt.wantErr == nil {
//				if err != tt.wantErr {
//					t.Errorf("Error for GetLastProposedEpoch(), got = %v, want = %v", err, tt.wantErr)
//				}
//			} else {
//				if err.Error() != tt.wantErr.Error() {
//					t.Errorf("Error for GetLastProposedEpoch(), got = %v, want = %v", err, tt.wantErr)
//				}
//			}
//		})
//	}
//}

//func convertToSliceOfInterface(arr []uint32) []interface{} {
//	s := make([]interface{}, len(arr))
//	for i, v := range arr {
//		s[i] = v
//	}
//	return s
//}

func TestGetCommitDataFileName(t *testing.T) {
	type args struct {
		address string
		path    string
		pathErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When GetCommitDataFileName() executes successfully",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				path:    "/home",
			},
			want:    "/home/0x000000000000000000000000000000000000dead_data",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			razorUtils = utilsMock

			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)

			utils := &UtilsStruct{}
			got, err := utils.GetCommitDataFileName(tt.args.address)
			if got != tt.want {
				t.Errorf("GetCommitDataFileName() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetCommitDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetCommitDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestCalculateSecret(t *testing.T) {
	var account types.Account
	var epoch uint32

	type args struct {
		path        string
		pathErr     error
		signedData  []byte
		signDataErr error
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test 1: When CalculateSecret executes successfully",
			args: args{
				path:       "/home/razor",
				signedData: []byte{234, 211},
			},
			want: solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString([]byte{234, 211})}),
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				pathErr:    errors.New("path error"),
				signedData: []byte{234, 211},
			},
			want: nil,
		},
		{
			name: "Test 3: When there is an error in getting signed data",
			args: args{
				path:        "/home/razor",
				signDataErr: errors.New("sign data error"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.UtilsInterface)
			accountUtilsMock := new(accountMocks.AccountInterface)

			razorUtils = utilsMock
			accounts.AccountUtilsInterface = accountUtilsMock

			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			accountUtilsMock.On("SignData", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.signedData, tt.args.signDataErr)

			utils := &UtilsStruct{}
			if got, _ := utils.CalculateSecret(account, epoch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
