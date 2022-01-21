package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClaimBlockReward(t *testing.T) {

	var options types.TransactionOptions

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	utilsStruct := UtilsStruct{
		razorUtils:        UtilsMock{},
		blockManagerUtils: BlockManagerMock{},
		transactionUtils:  TransactionMock{},
	}

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
			GetEpochMock = func(*ethclient.Client) (uint32, error) {
				return tt.args.epoch, tt.args.epochErr
			}

			GetStakerIdMock = func(*ethclient.Client, string) (uint32, error) {
				return tt.args.stakerId, tt.args.stakerIdErr
			}

			GetSortedProposedBlockIdsMock = func(*ethclient.Client, string, uint32) ([]uint32, error) {
				return tt.args.sortedProposedBlockIds, tt.args.sortedProposedBlockIdsErr
			}

			GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint32) (bindings.StructsBlock, error) {
				return tt.args.selectedBlock, tt.args.selectedBlockErr
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			ClaimBlockRewardMock = func(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error) {
				return tt.args.ClaimBlockRewardTxn, tt.args.ClaimBlockRewardErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := utilsStruct.ClaimBlockReward(options)
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
