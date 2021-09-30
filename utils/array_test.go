package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test if value is not present in the array",
			args: args{
				arr: []int{0, 1, 2},
				val: 4,
			},
			want: false,
		},
		{
			name: "Test if value is present in the array",
			args: args{
				arr: []int{0, 1, 2},
				val: 2,
			},
			want: true,
		},
		{
			name: "Test if array is empty",
			args: args{
				arr: []int{},
				val: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataInBytes(t *testing.T) {
	type args struct {
		data []*big.Int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test 1",
			args: args{
				data: []*big.Int{},
			},
			want: nil,
		},
		{
			name: "Test 2",
			args: args{
				data: []*big.Int{big.NewInt(1), big.NewInt(2)},
			},
			want: [][]byte{
				math2.U256Bytes(big.NewInt(1)),
				math2.U256Bytes(big.NewInt(2)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDataInBytes(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataInBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	type args struct {
		arr1 []uint32
		arr2 []uint32
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "Test when both arrays have same values but at different positions",
			args: args{
				arr1: []uint32{1, 1234, 2321},
				arr2: []uint32{1234, 1, 2321},
			},
			want:  false,
			want1: 1,
		},
		{
			name: "Test when both arrays have different length",
			args: args{
				arr1: []uint32{1, 1234},
				arr2: []uint32{1234, 1, 2321},
			},
			want:  false,
			want1: 3,
		},
		{
			name: "Test when both arrays are empty",
			args: args{
				arr1: []uint32{},
				arr2: []uint32{},
			},
			want:  true,
			want1: -1,
		},
		{
			name: "Test when both arrays are exactly identical",
			args: args{
				arr1: []uint32{1, 1232, 12423},
				arr2: []uint32{1, 1232, 12423},
			},
			want:  true,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsEqual(tt.args.arr1, tt.args.arr2)
			if got != tt.want {
				t.Errorf("IsEqual() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsEqual() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
