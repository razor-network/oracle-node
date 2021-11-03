package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core/types"
	"razor/pkg/bindings"
	"testing"
)

//func TestDispute(t *testing.T) {
//	type args struct {
//		client  *ethclient.Client
//		config  types.Configurations
//		account types.Account
//		epoch   uint32
//		blockId uint8
//		assetId int
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := Dispute(tt.args.client, tt.args.config, tt.args.account, tt.args.epoch, tt.args.blockId, tt.args.assetId); (err != nil) != tt.wantErr {
//				t.Errorf("Dispute() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestGiveSorted(t *testing.T) {
//	type args struct {
//		client        *ethclient.Client
//		blockManager  *bindings.BlockManager
//		txnOpts       *bind.TransactOpts
//		epoch         uint32
//		assetId       uint8
//		sortedStakers []uint32
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
func TestHandleDispute(t *testing.T) {
	var client *ethclient.Client
	var config types.Configurations
	var account types.Account
	var epoch uint32

	razorUtils = UtilsMock{}
	proposeUtils = ProposeUtilsMock{}
	cmdUtils = UtilsCmdMock{}

	type args struct {
		numberOfProposedBlocks    uint8
		numberOfProposedBlocksErr error
		proposedBlock             bindings.StructsBlock
		proposedBlockErr          error
		medians                   []uint32
		mediansErr                error
		activeAssetIds            []uint8
		activeAssetIdsErr         error
		isEqual                   bool
		iteration                 int
		disputeErr                error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test 1: When HandleDispute function executes successfully when there is a dispute case",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{101, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        false,
				iteration:      0,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 2: When HandleDispute function executes successfully when there is no dispute case",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{100, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 3: When there is an error in getting numberOfProposedBlocks",
			args: args{
				numberOfProposedBlocksErr: errors.New("numberOfProposedBlocks error"),
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{100, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: errors.New("numberOfProposedBlocks error"),
		},
		{
			name: "Test 4: When there is an error in getting proposedBlock",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlockErr:       errors.New("proposedBlock error"),
				medians:                []uint32{100, 200, 300},
				activeAssetIds:         []uint8{3, 4, 6},
				isEqual:                true,
				disputeErr:             nil,
			},
			want: nil,
		},
		{
			name: "Test 5: When there is an error in getting medians from MakeBlock ",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				mediansErr:     errors.New("medians error"),
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        true,
				disputeErr:     nil,
			},
			want: nil,
		},
		{
			name: "Test 6: When there is an error from Dispute function",
			args: args{
				numberOfProposedBlocks: 4,
				proposedBlock: bindings.StructsBlock{
					Medians: []uint32{100, 200, 300},
				},
				medians:        []uint32{101, 200, 300},
				activeAssetIds: []uint8{3, 4, 6},
				isEqual:        false,
				iteration:      0,
				disputeErr:     errors.New("dispute error"),
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumberOfProposedBlocksMock = func(*ethclient.Client, string, uint32) (uint8, error) {
				return tt.args.numberOfProposedBlocks, tt.args.numberOfProposedBlocksErr
			}

			GetProposedBlockMock = func(*ethclient.Client, string, uint32, uint8) (bindings.StructsBlock, error) {
				return tt.args.proposedBlock, tt.args.proposedBlockErr
			}

			MakeBlockMock = func(*ethclient.Client, string, bool) ([]uint32, error) {
				return tt.args.medians, tt.args.mediansErr
			}

			GetActiveAssetIdsMock = func(*ethclient.Client, string, uint32) ([]uint8, error) {
				return tt.args.activeAssetIds, tt.args.activeAssetIdsErr
			}

			IsEqualMock = func([]uint32, []uint32) (bool, int) {
				return tt.args.isEqual, tt.args.iteration
			}

			DisputeMock = func(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int) error {
				return tt.args.disputeErr
			}
			err := HandleDispute(client, config, account, epoch, razorUtils, proposeUtils, cmdUtils)

			if err == nil || tt.want == nil {
				if err != tt.want {
					t.Errorf("Error for HandleDispute function, got = %v, want = %v", err, tt.want)
				}
			} else {
				if err.Error() != tt.want.Error() {
					t.Errorf("Error for HandleDispute function, got = %v, want = %v", err, tt.want)
				}
			}

		})
	}

}
