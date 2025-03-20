package cmd

import (
	"errors"
	"io/fs"
	"math/big"
	"razor/accounts"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	utilsPkgMocks "razor/utils/mocks"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
)

func TestExecuteClaimBounty(t *testing.T) {
	var client *ethclient.Client
	var flagSet *pflag.FlagSet

	type args struct {
		config               types.Configurations
		configErr            error
		password             string
		address              string
		addressErr           error
		isFlagPassed         bool
		bountyId             uint32
		bountyIdErr          error
		claimBountyTxn       common.Hash
		claimBountyErr       error
		handleClaimBountyErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When executeClaimBounty function executes successfully",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				isFlagPassed:   true,
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting config",
			args: args{
				configErr:      errors.New("config error"),
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in getting address",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				isFlagPassed:   true,
				addressErr:     errors.New("address error"),
				bountyId:       2,
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 4: When there is an error in getting bountyId",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				isFlagPassed:   true,
				bountyIdErr:    errors.New("bountyId error"),
				claimBountyTxn: common.BigToHash(big.NewInt(1)),
			},
			expectedFatal: true,
		},
		{
			name: "Test 5: When there is an error from claimBounty function",
			args: args{
				config:         types.Configurations{},
				password:       "test",
				address:        "0x000000000000000000000000000000000000dead",
				isFlagPassed:   true,
				bountyId:       2,
				claimBountyErr: errors.New("claimBounty error"),
			},
			expectedFatal: true,
		},
		{
			name: "Test 6: When there is an error from HandleClaimBounty function",
			args: args{
				config:               types.Configurations{},
				password:             "test",
				address:              "0x000000000000000000000000000000000000dead",
				isFlagPassed:         false,
				bountyId:             2,
				handleClaimBountyErr: errors.New("HandleClaimBounty error"),
			},
			expectedFatal: true,
		},
	}

	defer func() { log.LogrusInstance.ExitFunc = nil }()
	var fatal bool
	log.LogrusInstance.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()
			setupTestEndpointsEnvironment()

			fileUtilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("GetConfigData").Return(tt.args.config, tt.args.configErr)
			utilsMock.On("AssignPassword", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.password)
			utilsMock.On("CheckPassword", mock.Anything).Return(nil)
			utilsMock.On("AccountManagerForKeystore").Return(&accounts.AccountManager{}, nil)
			flagSetMock.On("GetStringAddress", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.address, tt.args.addressErr)
			flagSetMock.On("GetUint32BountyId", mock.AnythingOfType("*pflag.FlagSet")).Return(tt.args.bountyId, tt.args.bountyIdErr)
			utilsMock.On("ConnectToClient", mock.AnythingOfType("string")).Return(client)
			utilsMock.On("IsFlagPassed", mock.Anything).Return(tt.args.isFlagPassed)
			cmdUtilsMock.On("HandleClaimBounty", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.handleClaimBountyErr)
			cmdUtilsMock.On("ClaimBounty", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.claimBountyTxn, tt.args.claimBountyErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)

			fatal = false
			utils := &UtilsStruct{}
			utils.ExecuteClaimBounty(flagSet)

			if fatal != tt.expectedFatal {
				t.Error("The executeClaimBounty function didn't execute as expected")
			}

		})
	}
}

func TestClaimBounty(t *testing.T) {
	var config types.Configurations
	var bountyInput types.RedeemBountyInput
	var blockTime int64

	type args struct {
		epoch           uint32
		epochErr        error
		bountyLock      types.BountyLock
		bountyLockErr   error
		txnOptsErr      error
		redeemBountyTxn *Types.Transaction
		redeemBountyErr error
		hash            common.Hash
		time            string
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When claimBounty function executes successfully",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 70,
				},
				redeemBountyTxn: &Types.Transaction{},
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When claimBounty function exits successfully if lock is not reached",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 80,
				},
				redeemBountyTxn: &Types.Transaction{},
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				epochErr: errors.New("epoch error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When there is an error in getting bounty lock",
			args: args{
				epoch:         70,
				bountyLockErr: errors.New("bountyLock error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("bountyLock error"),
		},
		{
			name: "Test 5: When the amount in bounty lock is 0",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(0),
					RedeemAfter: 70,
				},
			},
			want:    core.NilHash,
			wantErr: errors.New("bounty amount is 0"),
		},
		{
			name: "Test 6: When RedeemBounty transaction fails",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 70,
				},
				redeemBountyErr: errors.New("redeemBounty error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("redeemBounty error"),
		},
		{
			name: "Test 7: When claimBounty function exits successfully if lock is not reached",
			args: args{
				epoch: 79,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 80,
				},
				redeemBountyTxn: &Types.Transaction{},
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When there is an error in getting txnOpts",
			args: args{
				epoch: 70,
				bountyLock: types.BountyLock{
					Amount:      big.NewInt(1000),
					RedeemAfter: 70,
				},
				txnOptsErr: errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stakeManagerMock := new(mocks.StakeManagerInterface)
			utilsMock := new(utilsPkgMocks.Utils)
			transactionUtilsMock := new(mocks.TransactionInterface)
			timeMock := new(mocks.TimeInterface)

			razorUtils = utilsMock
			stakeManagerUtils = stakeManagerMock
			transactionUtils = transactionUtilsMock
			timeUtils = timeMock

			utilsMock.On("GetEpoch", mock.Anything).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetBountyLock", mock.Anything, mock.Anything).Return(tt.args.bountyLock, tt.args.bountyLockErr)
			timeMock.On("Sleep", mock.AnythingOfType("time.Duration")).Return()
			utilsMock.On("CalculateBlockTime", mock.AnythingOfType("*ethclient.Client")).Return(blockTime)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			stakeManagerMock.On("RedeemBounty", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32")).Return(tt.args.redeemBountyTxn, tt.args.redeemBountyErr)
			utilsMock.On("SecondsToReadableTime", mock.AnythingOfType("int")).Return(tt.args.time)
			transactionUtilsMock.On("Hash", mock.Anything).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.ClaimBounty(rpcParameters, config, bountyInput)
			if got != tt.want {
				t.Errorf("Txn hash for claimBounty function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for claimBounty function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestHandleClaimBounty(t *testing.T) {
	var (
		config   types.Configurations
		account  types.Account
		fileInfo fs.FileInfo
	)
	type args struct {
		disputeFilePath    string
		disputeFilePathErr error
		statErr            error
		disputeData        types.DisputeFileData
		disputeDataErr     error
		claimBountyTxn     common.Hash
		claimBountyTxnErr  error
		saveDataErr        error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When HandleClaimBounty() executes successfully",
			args: args{
				disputeFilePath: "",
				statErr:         nil,
				disputeData:     types.DisputeFileData{BountyIdQueue: []uint32{1}},
				claimBountyTxn:  common.BigToHash(big.NewInt(1)),
				saveDataErr:     nil,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When HandleClaimBounty() executes successfully and there are more than one bountyId in queue",
			args: args{
				disputeFilePath: "",
				statErr:         nil,
				disputeData:     types.DisputeFileData{BountyIdQueue: []uint32{1, 2}},
				claimBountyTxn:  common.BigToHash(big.NewInt(1)),
				saveDataErr:     nil,
			},
			wantErr: false,
		},
		{
			name: "Test 3: When there is an error in getting disputeFilePath",
			args: args{
				disputeFilePathErr: errors.New("error in getting disputeFilePath"),
			},
			wantErr: true,
		},
		{
			name: "Test 6: When there is an error in getting disputeData",
			args: args{
				disputeFilePath: "",
				statErr:         nil,
				disputeDataErr:  errors.New("error in getting diapute data"),
			},
			wantErr: true,
		},
		{
			name: "When there is an error in claimBounty",
			args: args{
				disputeFilePath:   "",
				statErr:           nil,
				disputeData:       types.DisputeFileData{BountyIdQueue: []uint32{1}},
				claimBountyTxnErr: errors.New("error in claimBounty"),
			},
			wantErr: true,
		},
		{
			name: "When there is an error in saving data to file",
			args: args{
				disputeFilePath: "",
				statErr:         nil,
				disputeData:     types.DisputeFileData{BountyIdQueue: []uint32{1}},
				claimBountyTxn:  common.BigToHash(big.NewInt(1)),
				saveDataErr:     errors.New("error in saving data to file"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			pathMock.On("GetDisputeDataFileName", mock.AnythingOfType("string")).Return(tt.args.disputeFilePath, tt.args.disputeFilePathErr)
			osPathMock.On("Stat", mock.Anything).Return(fileInfo, tt.args.statErr)
			fileUtilsMock.On("ReadFromDisputeJsonFile", mock.Anything).Return(tt.args.disputeData, tt.args.disputeDataErr)
			cmdUtilsMock.On("ClaimBounty", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.claimBountyTxn, tt.args.claimBountyTxnErr)
			utilsMock.On("WaitForBlockCompletion", mock.Anything, mock.Anything).Return(nil)
			fileUtilsMock.On("SaveDataToDisputeJsonFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.saveDataErr)

			ut := &UtilsStruct{}
			if err := ut.HandleClaimBounty(rpcParameters, config, account); (err != nil) != tt.wantErr {
				t.Errorf("AutoClaimBounty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
