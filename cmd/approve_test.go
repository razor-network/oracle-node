package cmd

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"
)

func TestApprove(t *testing.T) {
	type args struct {
		txnArgs         types.TransactionOptions
		callOpts        bind.CallOpts
		allowanceAmount *big.Int
		allowanceError  error
		txnOptsErr      error
		approveTxn      *Types.Transaction
		approveError    error
		hash            common.Hash
	}

	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When Allowance is smaller than amount to be approved",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(10000),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				allowanceAmount: big.NewInt(0),
				allowanceError:  nil,
				approveTxn:      &Types.Transaction{},
				approveError:    nil,
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When Allowance is greater than amount to be approved",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(1000),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				allowanceAmount: big.NewInt(10000),
				allowanceError:  nil,
				approveTxn:      &Types.Transaction{},
				approveError:    nil,
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: nil,
		},
		{
			name: "Test 3: When there is error in sending allowance ",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(10000),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				allowanceAmount: big.NewInt(0),
				allowanceError:  errors.New("allowance error"),
				approveTxn:      &Types.Transaction{},
				approveError:    nil,
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("allowance error"),
		},

		{
			name: "Test 4: When there is error in approve transaction",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(10000),
				},
				callOpts: bind.CallOpts{
					Pending:     false,
					From:        common.HexToAddress("0x000000000000000000000000000000000000dead"),
					BlockNumber: big.NewInt(1),
					Context:     context.Background(),
				},
				allowanceAmount: big.NewInt(0),
				allowanceError:  nil,
				approveTxn:      &Types.Transaction{},
				approveError:    errors.New("approve error"),
				hash:            common.BigToHash(big.NewInt(1)),
			},
			want:    common.Hash{0},
			wantErr: errors.New("approve error"),
		},
		{
			name: "Test 5: When there is error in getting txnOpts",
			args: args{
				txnArgs: types.TransactionOptions{
					Amount: big.NewInt(10000),
				},
				callOpts:        bind.CallOpts{},
				allowanceAmount: big.NewInt(0),
				txnOptsErr:      errors.New("txnOpts error"),
			},
			want:    core.NilHash,
			wantErr: errors.New("txnOpts error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUpMockInterfaces()

			utilsMock.On("GetOptions").Return(tt.args.callOpts)
			utilsMock.On("GetTxnOpts", mock.Anything, mock.Anything).Return(TxnOpts, tt.args.txnOptsErr)
			transactionMock.On("Hash", mock.Anything).Return(tt.args.hash)
			utilsMock.On("Allowance", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.allowanceAmount, tt.args.allowanceError)
			tokenManagerMock.On("Approve", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.args.approveTxn, tt.args.approveError)

			utils := &UtilsStruct{}

			got, err := utils.Approve(rpcParameters, tt.args.txnArgs)
			if got != tt.want {
				t.Errorf("Txn hash for approve function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for approve function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for approve function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
