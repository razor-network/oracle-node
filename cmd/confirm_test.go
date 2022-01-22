package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/cmd/mocks"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
)

func TestClaimBlockReward(t *testing.T) {
	var options types.TransactionOptions

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	type args struct {
		epoch                     uint32
		epochErr                  error
		stakerId                  uint32
		stakerIdErr               error
		sortedProposedBlockIds    []uint32
		sortedProposedBlockIdsErr error
		selectedBlock             bindings.StructsBlock
		selectedBlockErr          error
		txnOpts                   *bind.TransactOpts
		ClaimBlockRewardTxn       *Types.Transaction
		ClaimBlockRewardErr       error
		hash                      common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When ClaimBlockReward function executes successfully",
			args: args{
				epoch:                  5,
				stakerId:               2,
				sortedProposedBlockIds: []uint32{2, 1, 3},
				selectedBlock:          bindings.StructsBlock{ProposerId: 2},
				txnOpts:                txnOpts,
				ClaimBlockRewardTxn:    &Types.Transaction{},
				ClaimBlockRewardErr:    nil,
				hash:                   common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When ClaimBlockReward transaction fails",
			args: args{
				epoch:                  5,
				stakerId:               2,
				sortedProposedBlockIds: []uint32{2, 1, 3},
				selectedBlock:          bindings.StructsBlock{ProposerId: 2},
				txnOpts:                txnOpts,
				ClaimBlockRewardTxn:    &Types.Transaction{},
				ClaimBlockRewardErr:    errors.New("claimBlockReward error"),
				hash:                   common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("claimBlockReward error"),
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
			name: "Test 4: When there is an error in getting stakerId",
			args: args{
				epoch:       5,
				stakerIdErr: errors.New("stakerId error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("stakerId error"),
		},
		{
			name: "Test 5: When there is an error in getting sortedProposedBlockIds",
			args: args{
				epoch:                     5,
				stakerId:                  2,
				sortedProposedBlockIdsErr: errors.New("sortedProposedBlockIds error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("sortedProposedBlockIds error"),
		},
		{
			name: "Test 6: When there is an error in getting proposedBlock",
			args: args{
				epoch:                  5,
				stakerId:               2,
				sortedProposedBlockIds: []uint32{2, 1, 3},
				selectedBlockErr:       errors.New("block error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("block error"),
		},
		{
			name: "Test 7: When stakerId != proposerId and ClaimBlockReward function executes successfully",
			args: args{
				epoch:                  5,
				stakerId:               3,
				sortedProposedBlockIds: []uint32{2, 1, 3},
				selectedBlock:          bindings.StructsBlock{ProposerId: 2},
				txnOpts:                txnOpts,
				ClaimBlockRewardTxn:    &Types.Transaction{},
				ClaimBlockRewardErr:    nil,
				hash:                   common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			blockManagerMock := new(mocks.BlockManagerInterface)
			transactionUtilsMock := new(mocks.TransactionInterface)

			razorUtils = utilsMock
			blockManagerUtils = blockManagerMock
			transactionUtils = transactionUtilsMock

			utilsMock.On("GetTxnOpts", options).Return(tt.args.txnOpts)
			blockManagerMock.On("ClaimBlockReward", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts")).Return(tt.args.ClaimBlockRewardTxn, tt.args.ClaimBlockRewardErr)
			transactionUtilsMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStruct{}
			got, err := utils.ClaimBlockReward(options)
			if got != tt.want {
				t.Errorf("Txn hash for ClaimBlockReward function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for ClaimBlockReward function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ClaimBlockReward function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}
