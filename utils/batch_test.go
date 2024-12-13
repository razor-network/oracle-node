package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/core"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"strings"
	"testing"
)

func TestBatchCall(t *testing.T) {
	//Testing Batch call scenario for getting StakeSnapshot

	voteManagerABI, _ := abi.JSON(strings.NewReader(bindings.VoteManagerMetaData.ABI))
	stakeManagerABI, _ := abi.JSON(strings.NewReader(bindings.StakeManagerMetaData.ABI))
	numberOfArguments := 3

	type args struct {
		contractABI         *abi.ABI
		contractAddress     string
		methodName          string
		createBatchCallsErr error
		performBatchCallErr error
		results             []interface{}
		callErrors          []error
	}
	tests := []struct {
		name    string
		args    args
		want    [][]interface{}
		wantErr bool
	}{
		{
			name: "Test 1: When batch call executes successfully",
			args: args{
				contractABI:     &voteManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000a"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000c"),
				},
				callErrors: []error{nil, nil, nil},
			},
			want: [][]interface{}{
				{big.NewInt(10)},
				{big.NewInt(11)},
				{big.NewInt(12)},
			},
			wantErr: false,
		},
		{
			name: "Test 2: When one of batch calls throw an error",
			args: args{
				contractABI:     &voteManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					nil,
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000c"),
				},
				callErrors: []error{errors.New("batch call error"), nil, nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When BatchCalls receives an result of invalid type which cannot be type asserted to *string",
			args: args{
				contractABI:     &voteManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					42, // intentionally incorrect data type,
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000c"),
				},
				callErrors: []error{nil, nil, nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 4: When BatchCalls receives a nil result (empty batch call result error)",
			args: args{
				contractABI:     &voteManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					nil,
					nil,
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
				},
				callErrors: []error{nil, nil, nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 5: When BatchCalls receives an empty result (empty hex data error)",
			args: args{
				contractABI:     &voteManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					ptrString("0x"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000c"),
				},
				callErrors: []error{nil, nil, nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 6: When incorrect ABI is provided for unpacking",
			args: args{
				contractABI:     &stakeManagerABI,
				contractAddress: core.VoteManagerAddress,
				methodName:      core.GetStakeSnapshotMethod,
				results: []interface{}{
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000a"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000b"),
					ptrString("0x000000000000000000000000000000000000000000000000000000000000000c"),
				},
				callErrors: []error{nil, nil, nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 7: When there is an error in creating batch calls",
			args: args{
				contractABI:         &voteManagerABI,
				contractAddress:     core.VoteManagerAddress,
				methodName:          core.GetStakeSnapshotMethod,
				createBatchCallsErr: errors.New("create batch calls error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stakerIds []uint32
			for i := 1; i <= numberOfArguments; i++ {
				stakerIds = append(stakerIds, uint32(i))
			}

			arguments := make([][]interface{}, len(stakerIds))
			for i, stakerId := range stakerIds {
				arguments[i] = []interface{}{uint32(100), stakerId}
			}

			ClientInterface = &ClientStruct{}
			calls, err := ClientInterface.CreateBatchCalls(tt.args.contractABI, tt.args.contractAddress, tt.args.methodName, arguments)
			if err != nil {
				log.Error("Error in creating batch calls: ", err)
				return
			}
			// Mock batch call responses
			for i, result := range tt.args.results {
				if result != nil {
					calls[i].Result = result
				}
				calls[i].Error = tt.args.callErrors[i]
			}
			clientMock := new(mocks.ClientUtils)
			optionsPackageStruct := OptionsPackageStruct{
				ClientInterface: clientMock,
			}

			StartRazor(optionsPackageStruct)

			clientMock.On("CreateBatchCalls", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(calls, tt.args.createBatchCallsErr)
			clientMock.On("PerformBatchCall", mock.Anything, mock.Anything).Return(tt.args.performBatchCallErr)

			c := ClientStruct{}
			gotResults, err := c.BatchCall(rpcParameters, tt.args.contractABI, tt.args.contractAddress, tt.args.methodName, arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchCall() error = %v, but wantErr bool is %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, gotResults, tt.want) {
				t.Errorf("BatchCall() got = %v, want %v", gotResults, tt.want)
			}
		})
	}
}

func ptrString(s string) *string {
	return &s
}
