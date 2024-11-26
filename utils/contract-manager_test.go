package utils

import (
	"errors"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

func TestGetStakeManager(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		stakeManager    *bindings.StakeManager
		stakeManagerErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetStakeManager() executes successfully",
			args: args{
				stakeManager: &bindings.StakeManager{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting stakeManager",
			args: args{
				stakeManagerErr: errors.New("stakeManager error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewStakeManager", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.stakeManager, tt.args.stakeManagerErr)

			fatal = false
			utils.GetStakeManager(client)

			if fatal != tt.expectedFatal {
				t.Error("The GetStakeManager() function didn't execute as expected")
			}
		})
	}
}

func TestGetStakedToken(t *testing.T) {
	var client *ethclient.Client
	var address common.Address

	type args struct {
		stakedToken    *bindings.StakedToken
		stakedTokenErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetStakedToken() executes successfully",
			args: args{
				stakedToken: &bindings.StakedToken{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting stakedToken",
			args: args{
				stakedTokenErr: errors.New("stakedToken error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewStakedToken", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.stakedToken, tt.args.stakedTokenErr)

			fatal = false
			utils.GetStakedToken(client, address)

			if fatal != tt.expectedFatal {
				t.Error("The GetStakeToken() function didn't execute as expected")
			}
		})
	}
}

func TestGetTokenManager(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		tokenManager    *bindings.RAZOR
		tokenManagerErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetTokenManager() executes successfully",
			args: args{
				tokenManager: &bindings.RAZOR{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting tokenManager",
			args: args{
				tokenManagerErr: errors.New("tokenManager error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewRAZOR", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.tokenManager, tt.args.tokenManagerErr)

			fatal = false
			utils.GetTokenManager(client)

			if fatal != tt.expectedFatal {
				t.Error("The GetTokenManager() function didn't execute as expected")
			}
		})
	}
}

func TestGetAssetManager(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		assetManager    *bindings.CollectionManager
		assetManagerErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetCollectionManager() executes successfully",
			args: args{
				assetManager: &bindings.CollectionManager{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting assetManager",
			args: args{
				assetManagerErr: errors.New("assetManager error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewCollectionManager", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.assetManager, tt.args.assetManagerErr)

			fatal = false
			utils.GetCollectionManager(client)

			if fatal != tt.expectedFatal {
				t.Error("The GetCollectionManager() function didn't execute as expected")
			}
		})
	}
}

func TestGetBlockManager(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		blockManager    *bindings.BlockManager
		blockManagerErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetBlockManager() executes successfully",
			args: args{
				blockManager: &bindings.BlockManager{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting blockManager",
			args: args{
				blockManagerErr: errors.New("blockManager error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewBlockManager", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.blockManager, tt.args.blockManagerErr)

			fatal = false
			utils.GetBlockManager(client)

			if fatal != tt.expectedFatal {
				t.Error("The GetBlockManager() function didn't execute as expected")
			}
		})
	}
}

func TestGetVoteManager(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		voteManager    *bindings.VoteManager
		voteManagerErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetVoteManager() executes successfully",
			args: args{
				voteManager: &bindings.VoteManager{},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting voteManager",
			args: args{
				voteManagerErr: errors.New("voteManager error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bindingsMock := new(mocks.BindingsUtils)

			optionsPackageStruct := OptionsPackageStruct{
				BindingsInterface: bindingsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bindingsMock.On("NewVoteManager", mock.Anything, mock.AnythingOfType("*ethclient.Client")).Return(tt.args.voteManager, tt.args.voteManagerErr)

			fatal = false
			utils.GetVoteManager(client)

			if fatal != tt.expectedFatal {
				t.Error("The GetVoteManager() function didn't execute as expected")
			}
		})
	}
}
