package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		arr interface{}
		val interface{}
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
		{
			name: "Test for string values",
			args: args{
				arr: []string{"vote", "commit", "reveal"},
				val: "commit",
			},
			want: true,
		},
		{
			name: "Test for string values not present in the array",
			args: args{
				arr: []string{"vote", "commit", "reveal"},
				val: "propose",
			},
			want: false,
		},
		{
			name: "Test for string array and int value",
			args: args{
				arr: []string{"vote", "commit", "reveal"},
				val: 42,
			},
			want: false,
		},
		{
			name: "Test for int array and string value",
			args: args{
				arr: []int{0, 1, 2},
				val: "commit",
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
			want1: 0,
		},
		{
			name: "Test when both arrays have different length and len(arr1) < len(arr2)",
			args: args{
				arr1: []uint32{1, 1234},
				arr2: []uint32{1234, 1, 2321},
			},
			want:  false,
			want1: 2,
		},
		{
			name: "Test when both arrays have different length and len(arr1) > len(arr2)",
			args: args{
				arr1: []uint32{1234, 1, 2321},
				arr2: []uint32{1, 1234},
			},
			want:  false,
			want1: 2,
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

func TestCalculateSumOfUint8Array(t *testing.T) {
	type args struct {
		data []uint8
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test1: If the data is provided",
			args: args{
				data: []uint8{100, 100, 100},
			},
			want: uint(300),
		},
		{
			name: "Test2: If the data is not provided",
			args: args{
				data: nil,
			},
			want: uint(0),
		},
		{
			name: "Test3: If the data is empty",
			args: args{
				data: []uint8{},
			},
			want: uint(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSumOfUint8Array(tt.args.data); got != tt.want {
				t.Errorf("CalculateSumOfUint8Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertBigIntArrayToUint32Array(t *testing.T) {
	type args struct {
		data []*big.Int
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "Test when array has length more than 1",
			args: args{
				data: []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
			},
			want: []uint32{100, 200, 300},
		},
		{
			name: "Test when array has length 1",
			args: args{
				data: []*big.Int{big.NewInt(100)},
			},
			want: []uint32{100},
		},
		{
			name: "Test when array is nil",
			args: args{
				data: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBigIntArrayToUint32Array(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertBigIntArrayToUint32Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUint32ArrayToBigIntArray(t *testing.T) {
	type args struct {
		data []uint32
	}
	tests := []struct {
		name string
		args args
		want []*big.Int
	}{
		{
			name: "Test when array has length more than 1",
			args: args{
				data: []uint32{100, 200, 300},
			},
			want: []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
		},
		{
			name: "Test when array has length 1",
			args: args{
				data: []uint32{100},
			},
			want: []*big.Int{big.NewInt(100)},
		},
		{
			name: "Test when array is nil",
			args: args{
				data: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUint32ArrayToBigIntArray(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertUint32ArrayToBigIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUintArrayToUint16Array(t *testing.T) {
	type args struct {
		data []uint
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "Test when array has length more than 1",
			args: args{
				data: []uint{100, 150, 200},
			},
			want: []uint16{100, 150, 200},
		},
		{
			name: "Test when array has length 1",
			args: args{
				data: []uint{100},
			},
			want: []uint16{100},
		},
		{
			name: "Test when array is nil",
			args: args{
				data: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUintArrayToUint16Array(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertUintArrayToUint16Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStringFromArray(t *testing.T) {
	type args struct {
		source         string
		subStringArray []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1: When one string in array is substring of source string",
			args: args{
				source:         "This is a substring check",
				subStringArray: []string{"apple", "go", "check"},
			},
			want: true,
		},
		{
			name: "Test 2: When substring array is nil",
			args: args{
				source:         "This is a substring check",
				subStringArray: nil,
			},
			want: false,
		},
		{
			name: "Test 3: When source string is nil string",
			args: args{
				source:         "",
				subStringArray: []string{"apple", "go", "check"},
			},
			want: false,
		},
		{
			name: "Test 4: When source string and substring array is nil",
			args: args{
				source:         "",
				subStringArray: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsStringFromArray(tt.args.source, tt.args.subStringArray); got != tt.want {
				t.Errorf("ContainsStringFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
