package utils

import (
	"math/big"
	"reflect"
	"testing"
)

func TestMerkleTreeStructCreateMerkle(t *testing.T) {
	type args struct {
		values []*big.Int
	}
	tests := []struct {
		name string
		args args
		want [][][]byte
	}{
		{
			name: "Test 1: When CreateMerkle() executes successfully",
			args: args{
				values: []*big.Int{big.NewInt(1)},
			},
			want: [][][]byte{{{177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246}}},
		},
		{
			name: "Test 2: When it contains multiple values",
			args: args{
				values: []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
			},
			want: [][][]byte{{{52, 34, 71, 138, 66, 123, 171, 13, 20, 101, 234, 111, 34, 80, 119, 150, 191, 7, 189, 164, 50, 31, 68, 19, 51, 45, 101, 139, 128, 81, 172, 181}}, {{182, 253, 152, 238, 183, 224, 143, 197, 33, 241, 21, 17, 40, 154, 254, 77, 142, 135, 63, 215, 163, 251, 118, 171, 117, 127, 164, 124, 35, 245, 150, 233}, {177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246}}, {{177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246}, {177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246}, {177, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246}}},
		},
		{
			name: "Test 3: When CreateMerkle() executes successfully and it contain even value",
			args: args{
				values: []*big.Int{big.NewInt(2)},
			},
			want: [][][]byte{{{64, 87, 135, 250, 18, 168, 35, 224, 242, 183, 99, 28, 196, 27, 59, 168, 130, 139, 51, 33, 202, 129, 17, 17, 250, 117, 205, 58, 163, 187, 90, 206}}},
		},
		{
			name: "Test 4: When CreateMerkle() executes successfully and it contain odd value",
			args: args{
				values: []*big.Int{big.NewInt(3)},
			},
			want: [][][]byte{{{194, 87, 90, 14, 158, 89, 60, 0, 249, 89, 248, 201, 47, 18, 219, 40, 105, 195, 57, 90, 59, 5, 2, 208, 94, 37, 22, 68, 111, 113, 248, 91}}},
		},
		{
			name: "Test 5: When CreateMerkle() executes successfully and if the value is zero",
			args: args{
				values: []*big.Int{big.NewInt(0)},
			},
			want: [][][]byte{{{41, 13, 236, 217, 84, 139, 98, 168, 214, 3, 69, 169, 136, 56, 111, 200, 75, 166, 188, 149, 72, 64, 8, 246, 54, 47, 147, 22, 14, 243, 229, 99}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &MerkleTreeStruct{}
			if got := me.CreateMerkle(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMerkle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerkleTreeStructGetMerkleRoot(t *testing.T) {
	type args struct {
		tree [][][]byte
	}
	tests := []struct {
		name string
		args args
		want [32]byte
	}{
		{
			name: "When GetMerkleRoot() executes successfully",
			args: args{
				tree: [][][]byte{{{2}, {1}, {3}}},
			},
			want: [32]byte{2},
		},
		{
			name: "When tree is nil or it does not contain any element",
			args: args{
				tree: [][][]byte{{{}, {}, {}}},
			},
			want: [32]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &MerkleTreeStruct{}
			if got := me.GetMerkleRoot(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMerkleRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerkleTreeStructGetProofPath(t *testing.T) {
	type args struct {
		tree    [][][]byte
		assetId uint16
	}
	tests := []struct {
		name string
		args args
		want [][32]byte
	}{
		{
			name: "Test 1: When GetProofPath() executes successfully and assetId is even",
			args: args{
				tree:    [][][]byte{{{1, 10}, {2, 12}, {3, 13}}, {{4, 14}, {5, 15}, {6, 16}}, {{7, 17}, {8, 18}, {9, 19}}},
				assetId: 2,
			},
			want: [][32]byte{{4, 14}},
		},
		{
			name: "Test 2: When GetProofPath() executes successfully and assetId is odd",
			args: args{
				tree:    [][][]byte{{{1, 10}, {2, 12}, {3, 13}}, {{4, 14}, {5, 15}, {6, 16}}, {{7, 17}, {8, 18}, {9, 19}}},
				assetId: 1,
			},
			want: [][32]byte{{7, 17}, {5, 15}},
		},
		{
			name: "Test 3: When tree is nil",
			args: args{
				tree:    [][][]byte{},
				assetId: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &MerkleTreeStruct{}
			if got := me.GetProofPath(tt.args.tree, tt.args.assetId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProofPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
