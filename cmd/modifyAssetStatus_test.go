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
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"testing"
)

func TestCheckCurrentStatus(t *testing.T) {

	var client *ethclient.Client
	var assetId uint16

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		assetManagerUtils: AssetManagerMock{},
	}

	type args struct {
		callOpts        bind.CallOpts
		activeStatus    bool
		activeStatusErr error
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "Test 1: When CheckCurrentStatus function executes successfully",
			args: args{
				callOpts:        bind.CallOpts{},
				activeStatus:    true,
				activeStatusErr: nil,
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "Test 2: When GetActiveStatus function gives an error",
			args: args{
				callOpts:        bind.CallOpts{},
				activeStatusErr: errors.New("activeStatus error"),
			},
			want:    false,
			wantErr: errors.New("activeStatus error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetOptionsMock = func() bind.CallOpts {
				return tt.args.callOpts
			}

			GetActiveStatusMock = func(*ethclient.Client, *bind.CallOpts, uint16) (bool, error) {
				return tt.args.activeStatus, tt.args.activeStatusErr
			}

			got, err := CheckCurrentStatus(client, assetId, utilsStruct)
			if got != tt.want {
				t.Errorf("Status from CheckCurrentStatus function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for CheckCurrentStatus function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for CheckCurrentStatus function, got = %v, want %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestModifyAssetStatus(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))

	var flagSet *pflag.FlagSet
	var config types.Configurations
	var client *ethclient.Client

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		assetManagerUtils: AssetManagerMock{},
		cmdUtils:          UtilsCmdMock{},
		transactionUtils:  TransactionMock{},
		flagSetUtils:      FlagSetMock{},
	}

	type args struct {
		address             string
		addressErr          error
		assetId             uint16
		assetIdErr          error
		status              string
		statusErr           error
		parseStatus         bool
		parseStatusErr      error
		password            string
		currentStatus       bool
		currentStatusErr    error
		epoch               uint32
		epochErr            error
		txnOpts             *bind.TransactOpts
		SetCollectionStatus *Types.Transaction
		SetAssetStatusErr   error
		hash                common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When ModifyAssetStatus executes successfully",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting address",
			args: args{
				address:             "",
				addressErr:          errors.New("address error"),
				assetId:             1,
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test 3: When there is an error in getting assetId",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetIdErr:          errors.New("assetId error"),
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("assetId error"),
		},
		{
			name: "Test 4: When there is an error in getting status string",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				statusErr:           errors.New("status error"),
				parseStatus:         true,
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("status error"),
		},
		{
			name: "Test 5: When there is an error in parsing status to bool",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				status:              "true",
				parseStatusErr:      errors.New("parsing status error"),
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("parsing status error"),
		},
		{
			name: "Test 6: When there is an error in getting current status",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatusErr:    errors.New("current status error"),
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("current status error"),
		},
		{
			name: "Test 7: When currentStatus == status",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatus:       true,
				txnOpts:             txnOpts,
				SetCollectionStatus: &Types.Transaction{},
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 8: When SetAssetStatus transaction fails",
			args: args{
				address:           "0x000000000000000000000000000000000000dea1",
				assetId:           1,
				status:            "true",
				parseStatus:       true,
				password:          "test",
				currentStatus:     false,
				txnOpts:           txnOpts,
				SetAssetStatusErr: errors.New("SetAssetStatus error"),
				hash:              common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("SetAssetStatus error"),
		},
		{
			name: "Test 9: When WaitForDisputeOrConfirmState fails",
			args: args{
				address:             "0x000000000000000000000000000000000000dea1",
				assetId:             1,
				status:              "true",
				parseStatus:         true,
				password:            "test",
				currentStatus:       false,
				txnOpts:             txnOpts,
				epochErr:            errors.New("WaitForDisputeOrConfirmState error"),
				SetCollectionStatus: &Types.Transaction{},
				SetAssetStatusErr:   nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("WaitForDisputeOrConfirmState error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetUint16AssetIdMock = func(*pflag.FlagSet) (uint16, error) {
				return tt.args.assetId, tt.args.assetIdErr
			}

			GetStringStatusMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.status, tt.args.statusErr
			}

			ParseBoolMock = func(string2 string) (bool, error) {
				return tt.args.parseStatus, tt.args.parseStatusErr
			}

			PasswordPromptMock = func() string {
				return tt.args.password
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			CheckCurrentStatusMock = func(*ethclient.Client, uint16, UtilsStruct) (bool, error) {
				return tt.args.currentStatus, tt.args.currentStatusErr
			}

			GetTxnOptsMock = func(types.TransactionOptions, utils.RazorUtilsInterface) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			WaitForAppropriateStateMock = func(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			SetCollectionStatusMock = func(*ethclient.Client, *bind.TransactOpts, bool, uint16) (*Types.Transaction, error) {
				return tt.args.SetCollectionStatus, tt.args.SetAssetStatusErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.ModifyAssetStatus(flagSet, config)
			if got != tt.want {
				t.Errorf("Txn hash for modifyAssetStatus function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for modifyAssetStatus function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for modifyAssetStatus function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
