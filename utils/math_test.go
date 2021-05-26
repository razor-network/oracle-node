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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		gas *big.Int
		val float64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyFloatAndBigInt(tt.args.gas, tt.args.val); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyToEightDecimals(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyToEightDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
