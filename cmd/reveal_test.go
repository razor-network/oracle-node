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
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"
)

func TestHandleRevealState(t *testing.T) {
	var client *ethclient.Client
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
			name: "Test 1: When HandleRevealState returns no error",
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
			name: "Test 3: When HandleRevealState returns an error when epoch != epochLastCommitted",
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

			utilsMock := new(mocks.UtilsInterface)
			razorUtils = utilsMock

			utilsMock.On("GetEpochLastCommitted", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("uint32")).Return(tt.args.epochLastCommitted, tt.args.epochLastCommittedErr)

			utils := &UtilsStruct{}

			err := utils.HandleRevealState(client, staker, tt.args.epoch)
			if err == nil || tt.want == nil {
				if err != tt.want {
					t.Errorf("Error for HandleRevealState function, got = %v, want %v", err, tt.want)
				}
			} else {
				if err.Error() != tt.want.Error() {
					t.Errorf("Error for HandleRevealState function, got = %v, want %v", err, tt.want)
				}
			}

		})
	}
}

func TestReveal(t *testing.T) {
	var client *ethclient.Client
	var committedData []*big.Int
	var secret []byte
	var account types.Account
	var commitAccount string
	var config types.Configurations

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		state          int64
		stateErr       error
		epoch          uint32
		epochErr       error
		commitments    [32]byte
		commitmentsErr error
		allZeroStatus  bool
		txnOpts        *bind.TransactOpts
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
				state:          1,
				stateErr:       nil,
				epoch:          1,
				epochErr:       nil,
				commitments:    [32]byte{39, 216, 48, 133, 246, 76, 27, 204, 106, 253, 89, 128, 162, 117, 198, 16, 120, 59, 207, 163, 118, 68, 154, 30, 86, 80, 42, 68, 229, 42, 231, 115},
				commitmentsErr: nil,
				allZeroStatus:  false,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      nil,
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting state",
			args: args{
				stateErr:       errors.New("state error"),
				epoch:          1,
				epochErr:       nil,
				commitments:    [32]byte{39, 216, 48, 133, 246, 76, 27, 204, 106, 253, 89, 128, 162, 117, 198, 16, 120, 59, 207, 163, 118, 68, 154, 30, 86, 80, 42, 68, 229, 42, 231, 115},
				commitmentsErr: nil,
				allZeroStatus:  false,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      nil,
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("state error"),
		},
		{
			name: "Test 3: When there is an error in getting epoch",
			args: args{
				state:          1,
				stateErr:       nil,
				epochErr:       errors.New("epoch error"),
				commitments:    [32]byte{39, 216, 48, 133, 246, 76, 27, 204, 106, 253, 89, 128, 162, 117, 198, 16, 120, 59, 207, 163, 118, 68, 154, 30, 86, 80, 42, 68, 229, 42, 231, 115},
				commitmentsErr: nil,
				allZeroStatus:  false,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      nil,
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("epoch error"),
		},
		{
			name: "Test 4: When there is an error in getting commitments",
			args: args{
				state:          1,
				stateErr:       nil,
				epoch:          1,
				epochErr:       nil,
				commitmentsErr: errors.New("commitments error"),
				allZeroStatus:  false,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      nil,
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("commitments error"),
		},
		{
			name: "Test 5: When there are zero commitments ",
			args: args{
				state:          1,
				stateErr:       nil,
				epoch:          1,
				epochErr:       nil,
				commitments:    [32]byte{39, 216, 48, 133, 246, 76, 27, 204, 106, 253, 89, 128, 162, 117, 198, 16, 120, 59, 207, 163, 118, 68, 154, 30, 86, 80, 42, 68, 229, 42, 231, 115},
				commitmentsErr: nil,
				allZeroStatus:  true,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      nil,
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 6: When Reveal transaction fails",
			args: args{
				state:          1,
				stateErr:       nil,
				epoch:          1,
				epochErr:       nil,
				commitments:    [32]byte{39, 216, 48, 133, 246, 76, 27, 204, 106, 253, 89, 128, 162, 117, 198, 16, 120, 59, 207, 163, 118, 68, 154, 30, 86, 80, 42, 68, 229, 42, 231, 115},
				commitmentsErr: nil,
				allZeroStatus:  false,
				txnOpts:        txnOpts,
				revealTxn:      &Types.Transaction{},
				revealErr:      errors.New("reveal error"),
				hash:           common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("reveal error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)
			voteManagerUtilsMock := new(mocks.VoteManagerInterface)

			razorUtils = utilsMock
			transactionUtils = transactionUtilsMock
			voteManagerUtils = voteManagerUtilsMock

			utilsMock.On("GetDelayedState", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("int32")).Return(tt.args.state, tt.args.stateErr)
			utilsMock.On("GetEpoch", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.epoch, tt.args.epochErr)
			utilsMock.On("GetTxnOpts", mock.AnythingOfType("types.TransactionOptions")).Return(tt.args.txnOpts)
			utilsMock.On("GetCommitments", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("string")).Return(tt.args.commitments, tt.args.commitmentsErr)
			utilsMock.On("AllZero", mock.Anything).Return(tt.args.allZeroStatus)
			voteManagerUtilsMock.On("Reveal", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts"), mock.AnythingOfType("uint32"), mock.Anything, mock.Anything).Return(tt.args.revealTxn, tt.args.revealErr)
			transactionUtilsMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStruct{}

			got, err := utils.Reveal(client, committedData, secret, account, commitAccount, config)
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
