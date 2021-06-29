package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		arr []*big.Int
		val *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				arr: []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				val: nil,
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				arr: []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				val: big.NewInt(4),
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				arr: []*big.Int{},
				val: big.NewInt(4),
			},
			want: false,
		},
		{
			name: "Test 4",
			args: args{
				arr: []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				val: big.NewInt(1),
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				arr: []*big.Int{},
				val: nil,
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
		arr1 []*big.Int
		arr2 []*big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				arr1: []*big.Int{big.NewInt(1), big.NewInt(1234), big.NewInt(2321)},
				arr2: []*big.Int{big.NewInt(1234), big.NewInt(1), big.NewInt(2321)},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				arr1: []*big.Int{big.NewInt(1), big.NewInt(1234)},
				arr2: []*big.Int{big.NewInt(1234), big.NewInt(1), big.NewInt(2321)},
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				arr1: []*big.Int{},
				arr2: []*big.Int{},
			},
			want: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqual(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
