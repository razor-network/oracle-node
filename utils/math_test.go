package utils

import (
	"github.com/magiconair/properties/assert"
	log "github.com/sirupsen/logrus"
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

func TestGetAmountWithChecks(t *testing.T) {
	type args struct {
		amount string
		balance *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
		expectedFatal bool
	}{
		{
			name: "Test When amount is non-zero and less than balance",
			args: args{
				amount: "900",
				balance: big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
			},
			want: big.NewInt(1).Mul(big.NewInt(900), big.NewInt(1e18)),
			expectedFatal: false,
		},
		{
			name: "Test When amount is zero",
			args: args{
				amount: "0",
				balance: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want: big.NewInt(0),
			expectedFatal: false,
		},
		{
			name: "Test When amount Exceeds Balance-fatal",
			args: args{
				amount: "10000",
				balance: big.NewInt(1).Mul(big.NewInt(900), big.NewInt(1e18)),
			},
			want: big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
			expectedFatal: true,
		},
		{
			name: "Test When amount is equal to balance",
			args: args{
				amount: "1000",
				balance: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			expectedFatal: false,
		},
	}

	defer func() { log.StandardLogger().ExitFunc = nil }()
	var fatal bool
	log.StandardLogger().ExitFunc = func(int){ fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fatal = false
			got := GetAmountWithChecks(tt.args.amount, tt.args.balance)
			if tt.expectedFatal {
				assert.Equal(t, tt.expectedFatal, fatal)
			}

			if got.Cmp(tt.want)!=0 {
				t.Errorf("GetAmountWithChecks() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_performAggregation(t *testing.T) {
	type args struct {
		data []*big.Int
		aggregationMethod uint32
	}

	tests := []struct {
		name string
		args args
		want *big.Int
		wantErr bool
	}{
		{
			name: "Test Median for Odd Number of elements",
			args: args{
				data: []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				aggregationMethod: 1,
			},
			want: big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for Even Number of elements",
			args: args{
				data: []*big.Int{big.NewInt(0), big.NewInt(1)},
				aggregationMethod: 1,
			},
			want: big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for single element",
			args: args{
				data: []*big.Int{big.NewInt(1)},
				aggregationMethod: 1,
			},
			want: big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for elements with higher value",
			args: args{
				data: []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 1,
			},
			want: big.NewInt(1500),
			wantErr: false,
		},
		{
			name: "Test Median for 0 elements",
			args: args{
				data: []*big.Int{},
				aggregationMethod: 1,
			},
			want: nil ,
			wantErr: true,
		},
		{
			name: "Test Mean for multiple number of elements",
			args: args{
				data: []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				aggregationMethod: 2,
			},
			want: big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Mean for single element",
			args: args{
				data: []*big.Int{big.NewInt(1)},
				aggregationMethod: 2,
			},
			want: big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Mean for elements with higher value",
			args: args{
				data: []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 2,
			},
			want: big.NewInt(1250),
			wantErr: false,
		},
		{
			name: "Test Mean for 0 elements",
			args: args{
				data: []*big.Int{},
				aggregationMethod: 2,
			},
			want: nil ,
			wantErr: true,
		},
		{
			name: "Test incorrect input for AggregationMethod",
			args: args{
				data: []*big.Int{big.NewInt(1)},
				aggregationMethod: 3,
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,err := performAggregation(tt.args.data, tt.args.aggregationMethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("performAggregation() = %v, want %v", got, tt.want)
			}
		})
	}
}
