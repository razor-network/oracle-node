package utils

import (
	"math/big"
	"reflect"
	"testing"
)

func TestAllZero(t *testing.T) {
	type args struct {
		bytesValue [32]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				bytesValue: [32]byte{00000000000000000000000000000000},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				bytesValue: [32]byte{00000000000000000000000000000001},
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				bytesValue: [32]byte{},
			},
			want: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllZero(tt.args.bytesValue); got != tt.want {
				t.Errorf("AllZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToNumber(t *testing.T) {
	type args struct {
		num interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			name:    "Test int",
			args:    args{
				num: 4,
			},
			want:    big.NewFloat(4),
			wantErr: false,
		},
		{
			name:    "Test float",
			args:    args{
				num: 0.4,
			},
			want:    big.NewFloat(0.4),
			wantErr: false,
		},
		{
			name:    "Test string",
			args:    args{
				num: "4",
			},
			want:    big.NewFloat(4),
			wantErr: false,
		},
		{
			name:    "Test nil",
			args:    args{
				num: nil,
			},
			want:    big.NewFloat(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToNumber(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyFloatAndBigInt(t *testing.T) {
	type args struct {
		bigIntVal   *big.Int
		floatingVal float64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1",
			args: args{
				bigIntVal:   nil,
				floatingVal: 0,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 2",
			args: args{
				bigIntVal:   big.NewInt(1),
				floatingVal: 0,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 3",
			args: args{
				bigIntVal:   big.NewInt(0),
				floatingVal: 999.99,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 4",
			args: args{
				bigIntVal:   big.NewInt(21),
				floatingVal: 999.99,
			},
			want: big.NewInt(20999),
		},
		{
			name: "Test 5",
			args: args{
				bigIntVal:   big.NewInt(20000),
				floatingVal: 1.5,
			},
			want: big.NewInt(30000),
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyFloatAndBigInt(tt.args.bigIntVal, tt.args.floatingVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyFloatAndBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyToEightDecimals(t *testing.T) {
	type args struct {
		num *big.Float
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1",
			args: args{
				num: big.NewFloat(1.22342),
			},
			want: big.NewInt(122342000),
		},
		{
			name: "Test 2",
			args: args{
				num: big.NewFloat(0),
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 3",
			args: args{
				num: nil,
			},
			want: big.NewInt(0),
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyToEightDecimals(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyToEightDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}