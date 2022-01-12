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
		txnOpts             *bind.TransactOpts
		ClaimBlockRewardTxn *Types.Transaction
		ClaimBlockRewardErr error
		hash                common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test1: When ClaimBlockReward function executes successfully",
			args: args{
				txnOpts:             txnOpts,
				ClaimBlockRewardTxn: &Types.Transaction{},
				ClaimBlockRewardErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test2: When ClaimBlockReward transaction fails",
			args: args{
				txnOpts:             txnOpts,
				ClaimBlockRewardTxn: &Types.Transaction{},
				ClaimBlockRewardErr: errors.New("claimBlockReward error"),
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("claimBlockReward error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			blockManagerMock := new(mocks.BlockManagerInterfaceMockery)
			transactionUtilsMock := new(mocks.TransactionInterfaceMockery)

			razorUtilsMockery = utilsMock
			blockManagerUtilsMockery = blockManagerMock
			transactionUtilsMockery = transactionUtilsMock

			utilsMock.On("GetTxnOpts", options).Return(tt.args.txnOpts)
			blockManagerMock.On("ClaimBlockReward", mock.AnythingOfType("*ethclient.Client"), mock.AnythingOfType("*bind.TransactOpts")).Return(tt.args.ClaimBlockRewardTxn, tt.args.ClaimBlockRewardErr)
			transactionUtilsMock.On("Hash", mock.AnythingOfType("*types.Transaction")).Return(tt.args.hash)

			utils := &UtilsStructMockery{}
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
