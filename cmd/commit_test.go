package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"math/big"
	randMath "math/rand"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"reflect"
	"testing"
)

func TestCommit(t *testing.T) {
	var client *ethclient.Client
	var data []*big.Int
	var secret []byte
	var account types.Account
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:       UtilsMock{},
		voteManagerUtils: VoteManagerMock{},
		transactionUtils: TransactionMock{},
	}

	type args struct {
		state     int64
		stateErr  error
		epoch     uint32
		epochErr  error
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
				epoch:     1,
				epochErr:  nil,
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
				epoch:     1,
				epochErr:  nil,
				txnOpts:   txnOpts,
				commitTxn: &Types.Transaction{},
				commitErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				state:     0,
				stateErr:  nil,
				epochErr:  errors.New("epoch error"),
				txnOpts:   txnOpts,
				commitTxn: &Types.Transaction{},
				commitErr: nil,
				hash:      common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When Commit transaction fails",
			args: args{
				state:     0,
				stateErr:  nil,
				epoch:     1,
				epochErr:  nil,
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

			GetDelayedStateMock = func(*ethclient.Client, int32) (int64, error) {
				return tt.args.state, tt.args.stateErr
			}

			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			GetTxnOptsMock = func(types.TransactionOptions, utils.RazorUtilsInterface) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			CommitMock = func(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error) {
				return tt.args.commitTxn, tt.args.commitErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.Commit(client, data, secret, account, config)
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
	)

	rogueValue := big.NewInt(int64(randMath.Intn(10000000)))

	utilsStruct := UtilsStruct{
		razorUtils: UtilsMock{},
	}

	type args struct {
		data               []*big.Int
		dataErr            error
		numActiveAssets    *big.Int
		numActiveAssetsErr error
		rogueValue         *big.Int
		rogue              types.Rogue
	}
	tests := []struct {
		name    string
		args    args
		want    []*big.Int
		wantErr error
	}{
		{
			name: "Test 1: When HandleCommitState executes successfully",
			args: args{
				data:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
				dataErr: nil,
				rogue:   types.Rogue{IsRogue: false},
			},
			want:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting data from getActiveAssetData",
			args: args{
				rogue:   types.Rogue{IsRogue: false},
				dataErr: errors.New("data error"),
			},
			want:    nil,
			wantErr: errors.New("data error"),
		},
		{
			name: "Test 3: When rogue mode is activated but not for commit ",
			args: args{
				rogue: types.Rogue{IsRogue: true, RogueMode: []string{"propose"}},
				data:  []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
			},
			want:    []*big.Int{big.NewInt(6701548), big.NewInt(478307)},
			wantErr: nil,
		},
		{
			name: "Test 4: When rogue mode is activated for commit ",
			args: args{
				numActiveAssets: big.NewInt(1),
				rogue:           types.Rogue{IsRogue: true, RogueMode: []string{"propose", "commit"}},
				rogueValue:      rogueValue,
				data:            []*big.Int{rogueValue},
			},
			want:    []*big.Int{rogueValue},
			wantErr: nil,
		},
		{
			name: "Test 5: When there is an error in fetching numActiveAssets",
			args: args{
				numActiveAssets:    nil,
				numActiveAssetsErr: errors.New("Error in fetching active assets"),
				rogue:              types.Rogue{IsRogue: true, RogueMode: []string{"propose", "commit"}},
			},
			want:    nil,
			wantErr: errors.New("Error in fetching active assets"),
		},
	}
	for _, tt := range tests {

		GetActiveAssetsDataMock = func(*ethclient.Client, uint32) ([]*big.Int, error) {
			return tt.args.data, tt.args.dataErr
		}

		GetNumActiveAssetsMock = func(*ethclient.Client) (*big.Int, error) {
			return tt.args.numActiveAssets, tt.args.numActiveAssetsErr
		}

		GetRogueRandomValueMock = func(int) *big.Int {
			return tt.args.rogueValue
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := utilsStruct.HandleCommitState(client, epoch, tt.args.rogue)
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
