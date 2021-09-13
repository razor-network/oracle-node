package utils

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/magiconair/properties/assert"
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
			name: "Test int",
			args: args{
				num: 4,
			},
			want:    big.NewFloat(4),
			wantErr: false,
		},
		{
			name: "Test float",
			args: args{
				num: 0.4,
			},
			want:    big.NewFloat(0.4),
			wantErr: false,
		},
		{
			name: "Test string",
			args: args{
				num: "4",
			},
			want:    big.NewFloat(4),
			wantErr: false,
		},
		{
			name: "Test nil",
			args: args{
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

func TestCheckAmountAndBalance(t *testing.T) {
	type args struct {
		amount  *big.Int
		balance *big.Int
	}
	tests := []struct {
		name          string
		args          args
		want          *big.Int
		expectedFatal bool
	}{
		{
			name: "Test When amount is non-zero and less than balance",
			args: args{
				amount:  big.NewInt(1).Mul(big.NewInt(900), big.NewInt(1e18)),
				balance: big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(900), big.NewInt(1e18)),
			expectedFatal: false,
		},
		{
			name: "Test When amount is zero",
			args: args{
				amount:  big.NewInt(0),
				balance: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want:          big.NewInt(0),
			expectedFatal: false,
		},
		{
			name: "Test When amount Exceeds Balance-fatal",
			args: args{
				amount:  big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
				balance: big.NewInt(1).Mul(big.NewInt(900), big.NewInt(1e18)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(10000), big.NewInt(1e18)),
			expectedFatal: true,
		},
		{
			name: "Test When amount is equal to balance",
			args: args{
				amount:  big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
				balance: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want:          big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			expectedFatal: false,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fatal = false
			got := CheckAmountAndBalance(tt.args.amount, tt.args.balance)
			if tt.expectedFatal {
				assert.Equal(t, tt.expectedFatal, fatal)
			}

			if got.Cmp(tt.want) != 0 {
				t.Errorf("CheckAmountAndBalance() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestGetAmountInWei(t *testing.T) {
	type args struct {
		amount *big.Int
	}

	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test when amount is non-zero",
			args: args{
				big.NewInt(1000),
			},
			want: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
		},
		{
			name: "Test when amount is zero",
			args: args{
				big.NewInt(0),
			},
			want: big.NewInt(1).Mul(big.NewInt(0), big.NewInt(1e18)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAmountInWei(tt.args.amount)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAmountInWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFractionalAmountInWei(t *testing.T) {
	type args struct {
		amount *big.Int
		power  string
	}

	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test when amount is non-zero and power is non-zero",
			args: args{
				amount: big.NewInt(1000),
				power:  "17",
			},
			want: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(17), nil)),
		},
		{
			name: "Test when amount is zero and power is non-zero",
			args: args{
				amount: big.NewInt(0),
				power:  "15",
			},
			want: big.NewInt(1).Mul(big.NewInt(0), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(15), nil)),
		},
		{
			name: "Test when amount is non-zero and power is zero",
			args: args{
				amount: big.NewInt(1000),
				power:  "0",
			},
			want: big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(0), nil)),
		},
		{
			name: "Test when amount is zero and power is also zero",
			args: args{
				amount: big.NewInt(0),
				power:  "0",
			},
			want: big.NewInt(1).Mul(big.NewInt(0), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(0), nil)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFractionalAmountInWei(tt.args.amount, tt.args.power)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFractionalAmountInWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAmountInDecimal(t *testing.T) {
	type args struct {
		amount *big.Int
	}

	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{
			name: "Test1",
			args: args{
				big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e18)),
			},
			want: new(big.Float).SetInt(big.NewInt(1000)),
		},
		{
			name: "Test 2",
			args: args{
				big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1e17)),
			},
			want: big.NewFloat(100),
		},
		{
			name: "Test 3",
			args: args{
				big.NewInt(1).Mul(big.NewInt(555), big.NewInt(1e16)),
			},
			want: new(big.Float).Quo(new(big.Float).SetInt(big.NewInt(1).Mul(big.NewInt(555), big.NewInt(1e16))), new(big.Float).SetInt(big.NewInt(1e18))),
		},
		{
			name: "Test 4",
			args: args{
				big.NewInt(1).Mul(big.NewInt(123456789), big.NewInt(1e10)),
			},
			want: new(big.Float).Quo(new(big.Float).SetInt(big.NewInt(1).Mul(big.NewInt(123456789), big.NewInt(1e10))), new(big.Float).SetInt(big.NewInt(1e18))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAmountInDecimal(tt.args.amount)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("GetAmountInDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_performAggregation(t *testing.T) {
	type args struct {
		data              []*big.Int
		aggregationMethod uint32
		power             int8
	}

	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test Median for Odd Number of elements",
			args: args{
				data:              []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
				aggregationMethod: 1,
				power:             2,
			},
			want:    big.NewInt(100),
			wantErr: false,
		},
		{
			name: "Test Median for Even Number of elements",
			args: args{
				data:              []*big.Int{big.NewInt(0), big.NewInt(1)},
				aggregationMethod: 1,
				power:             3,
			},
			want:    big.NewInt(1000),
			wantErr: false,
		},
		{
			name: "Test Median for single element",
			args: args{
				data:              []*big.Int{big.NewInt(1)},
				aggregationMethod: 1,
				power:             0,
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for elements with higher value",
			args: args{
				data:              []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 1,
				power:             8,
			},
			want:    big.NewInt(150000000000),
			wantErr: false,
		},
		{
			name: "Test Median for 0 elements",
			args: args{
				data:              []*big.Int{},
				aggregationMethod: 1,
				power:             0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Mean for multiple number of elements",
			args: args{
				data:              []*big.Int{big.NewInt(0), big.NewInt(10), big.NewInt(20)},
				aggregationMethod: 2,
				power:             -1,
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Mean for single element",
			args: args{
				data:              []*big.Int{big.NewInt(100000)},
				aggregationMethod: 2,
				power:             -2,
			},
			want:    big.NewInt(1000),
			wantErr: false,
		},
		{
			name: "Test Mean for elements with higher value",
			args: args{
				data:              []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 2,
				power:             -1,
			},
			want:    big.NewInt(125),
			wantErr: false,
		},
		{
			name: "Test Mean for 0 elements",
			args: args{
				data:              []*big.Int{},
				aggregationMethod: 2,
				power:             0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test incorrect input for AggregationMethod",
			args: args{
				data:              []*big.Int{big.NewInt(1)},
				aggregationMethod: 3,
				power:             2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := performAggregation(tt.args.data, tt.args.aggregationMethod, tt.args.power)
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

func TestMultiplyWithPower(t *testing.T) {
	type args struct {
		num   *big.Float
		power int8
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test value when power is 8",
			args: args{
				num:   big.NewFloat(1.22342),
				power: 8,
			},
			want: big.NewInt(122342000),
		},
		{
			name: "Test value when number is 0",
			args: args{
				num:   big.NewFloat(0),
				power: 0,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test value when number is nil",
			args: args{
				num:   nil,
				power: 10,
			},
			want: big.NewInt(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyWithPower(tt.args.num, tt.args.power); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyWithPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
