package utils

import (
	"errors"
	"math/big"
	"razor/utils/mocks"
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
		{
			name: "Test incorrect string",
			args: args{
				num: "4w",
			},
			want:    big.NewFloat(0),
			wantErr: true,
		},
		{
			name: "Test when variable type is out of switch case",
			args: args{
				num: big.NewInt(4),
			},
			want:    big.NewFloat(0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UtilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: UtilsMock,
			}
			utils := StartRazor(optionsPackageStruct)
			got, err := utils.ConvertToNumber(tt.args.num)
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
			UtilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: UtilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			if got := utils.MultiplyFloatAndBigInt(tt.args.bigIntVal, tt.args.floatingVal); !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    *big.Int
		wantErr error
	}{
		{
			name: "Test when amount is non-zero and power is non-zero",
			args: args{
				amount: big.NewInt(1000),
				power:  "17",
			},
			want:    big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(17), nil)),
			wantErr: nil,
		},
		{
			name: "Test when amount is zero and power is non-zero",
			args: args{
				amount: big.NewInt(0),
				power:  "15",
			},
			want:    big.NewInt(1).Mul(big.NewInt(0), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(15), nil)),
			wantErr: nil,
		},
		{
			name: "Test when amount is non-zero and power is zero",
			args: args{
				amount: big.NewInt(1000),
				power:  "0",
			},
			want:    big.NewInt(1).Mul(big.NewInt(1000), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(0), nil)),
			wantErr: nil,
		},
		{
			name: "Test when amount is zero and power is also zero",
			args: args{
				amount: big.NewInt(0),
				power:  "0",
			},
			want:    big.NewInt(1).Mul(big.NewInt(0), big.NewInt(1).Exp(big.NewInt(10), big.NewInt(0), nil)),
			wantErr: nil,
		},
		{
			name: "Test when there is an incorrect input type for power",
			args: args{
				amount: big.NewInt(0),
				power:  "A",
			},
			want:    nil,
			wantErr: errors.New("SetString: error"),
		},
		{
			name: "Test when there is a decimal input type for power",
			args: args{
				amount: big.NewInt(0),
				power:  "11.5",
			},
			want:    nil,
			wantErr: errors.New("SetString: error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFractionalAmountInWei(tt.args.amount, tt.args.power)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFractionalAmountInWei() = %v, want %v", got, tt.want)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetFractionalAmountInWei function, got = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetFractionalAmountInWei function, got = %v, want %v", err, tt.wantErr)
				}
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
		{
			name: "Test 5",
			args: args{
				big.NewInt(0),
			},
			want: big.NewFloat(0),
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

func TestConvertWeiToEth(t *testing.T) {
	type args struct {
		data *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			name:    "Test if the value is 0",
			args:    args{big.NewInt(0)},
			want:    big.NewFloat(0),
			wantErr: true,
		},
		{
			name:    "Test if data is bigger than 1e18",
			args:    args{big.NewInt(2 * 1e18)},
			want:    big.NewFloat(2).SetPrec(32),
			wantErr: false,
		},
		{
			name:    "Test if data is smaller than 1e18",
			args:    args{big.NewInt(234 * 1e12)},
			want:    big.NewFloat(234 * 1e-6).SetPrec(32),
			wantErr: false,
		},
		{
			name:    "Test if data is in the order of 1e18",
			args:    args{big.NewInt(392 * 1e16)},
			want:    big.NewFloat(3.92).SetPrec(32),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertWeiToEth(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertWeiToEth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertWeiToEth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_performAggregation(t *testing.T) {
	type args struct {
		data              []*big.Int
		weight            []uint8
		aggregationMethod uint32
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
				weight:            []uint8{1, 1, 1},
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for Even Number of elements",
			args: args{
				data:              []*big.Int{big.NewInt(0), big.NewInt(1)},
				aggregationMethod: 1,
				weight:            []uint8{1, 1},
			},
			want:    big.NewInt(0),
			wantErr: false,
		},
		{
			name: "Test Median for single element",
			args: args{
				data:              []*big.Int{big.NewInt(1)},
				aggregationMethod: 1,
				weight:            []uint8{1},
			},
			want:    big.NewInt(1),
			wantErr: false,
		},
		{
			name: "Test Median for elements with higher value",
			args: args{
				data:              []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 1,
				weight:            []uint8{100, 100, 100, 100},
			},
			want:    big.NewInt(1000),
			wantErr: false,
		},
		{
			name: "Test Median for 0 elements",
			args: args{
				data:              []*big.Int{},
				aggregationMethod: 1,
				weight:            []uint8{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Mean for multiple number of elements",
			args: args{
				data:              []*big.Int{big.NewInt(0), big.NewInt(10), big.NewInt(20)},
				aggregationMethod: 2,
				weight:            []uint8{10, 20, 30},
			},
			want:    big.NewInt(13),
			wantErr: false,
		},
		{
			name: "Test Mean for single element",
			args: args{
				data:              []*big.Int{big.NewInt(100000)},
				aggregationMethod: 2,
				weight:            []uint8{99},
			},
			want:    big.NewInt(100000),
			wantErr: false,
		},
		{
			name: "Test Mean for elements with higher value",
			args: args{
				data:              []*big.Int{big.NewInt(500), big.NewInt(1000), big.NewInt(1500), big.NewInt(2000)},
				aggregationMethod: 2,
				weight:            []uint8{100, 100, 100, 100},
			},
			want:    big.NewInt(1250),
			wantErr: false,
		},
		{
			name: "Test Mean for 0 elements",
			args: args{
				data:              []*big.Int{},
				aggregationMethod: 2,
				weight:            []uint8{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test incorrect input for AggregationMethod",
			args: args{
				data:              []*big.Int{big.NewInt(1)},
				aggregationMethod: 3,
				weight:            []uint8{1},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := performAggregation(tt.args.data, tt.args.weight, tt.args.aggregationMethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("performAggregation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("performAggregation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateWeightedMedian(t *testing.T) {
	type args struct {
		data        []*big.Int
		weight      []uint8
		totalWeight uint
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1: Weighted median for even number of elements",
			args: args{
				data:        []*big.Int{big.NewInt(4), big.NewInt(1), big.NewInt(3), big.NewInt(2)},
				weight:      []uint8{25, 49, 25, 1},
				totalWeight: 100,
			},
			want: big.NewInt(2),
		},
		{
			name: "Test 2: Weighted median for odd number of elements",
			args: args{
				data:        []*big.Int{big.NewInt(5), big.NewInt(1), big.NewInt(3), big.NewInt(2), big.NewInt(4)},
				weight:      []uint8{25, 15, 20, 10, 30},
				totalWeight: 100,
			},
			want: big.NewInt(4),
		},
		{
			name: "Test 3: Weighted median for eth values",
			args: args{
				data:        []*big.Int{big.NewInt(423469), big.NewInt(423322), big.NewInt(423402)},
				weight:      []uint8{100, 100, 100},
				totalWeight: 300,
			},
			want: big.NewInt(423402),
		},
		{
			name: "Test 4: When the data array is empty",
			args: args{
				data:        []*big.Int{},
				weight:      []uint8{100, 100, 100},
				totalWeight: 300,
			},
			want: nil,
		},
		{
			name: "Test 5:  When the weight array is empty",
			args: args{
				data:        []*big.Int{big.NewInt(423469), big.NewInt(423322), big.NewInt(423402)},
				weight:      []uint8{},
				totalWeight: 0,
			},
			want: nil,
		},
		{
			name: "Test 6:  When the total weight is 0",
			args: args{
				data:        []*big.Int{big.NewInt(423469), big.NewInt(423322), big.NewInt(423402)},
				weight:      []uint8{100, 100, 100},
				totalWeight: 0,
			},
			want: nil,
		},
		{
			name: "Test 7:  When very high total weight is being passed so sum of weights is never greater than or equal to 0.5",
			args: args{
				data:        []*big.Int{big.NewInt(5), big.NewInt(1), big.NewInt(3), big.NewInt(2), big.NewInt(4)},
				weight:      []uint8{25, 15, 20, 10, 30},
				totalWeight: 10000,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeightedMedian(tt.args.data, tt.args.weight, tt.args.totalWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateWeightedMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFractionalWeight(t *testing.T) {
	type args struct {
		weights     []uint8
		totalWeight uint
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "Test 1: Odd number of array elements if total weight = 100",
			args: args{
				weights:     []uint8{25, 15, 20, 10, 30},
				totalWeight: 100,
			},
			want: []float32{0.25, 0.15, 0.2, 0.1, 0.3},
		},
		{
			name: "Test 2: Even number of array elements if total weight = 100",
			args: args{
				weights:     []uint8{25, 49, 25, 1},
				totalWeight: 100,
			},
			want: []float32{0.25, 0.49, 0.25, 0.01},
		},
		{
			name: "Test 3: Weight is more than 100",
			args: args{
				weights:     []uint8{100, 100, 100},
				totalWeight: 300,
			},
			want: []float32{0.33333334, 0.33333334, 0.33333334},
		},
		{
			name: "Test 4: Weight array is empty",
			args: args{
				weights:     []uint8{},
				totalWeight: 100,
			},
			want: nil,
		},
		{
			name: "Test 5: Total weight it 0",
			args: args{
				weights:     []uint8{25, 49, 25, 1},
				totalWeight: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFractionalWeight(tt.args.weights, tt.args.totalWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFractionalWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertRZRToSRZR(t *testing.T) {
	type args struct {
		amount       *big.Int
		currentStake *big.Int
		totalSupply  *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When currentStake and totalSupply are equal",
			args: args{
				amount:       big.NewInt(500),
				currentStake: big.NewInt(2000),
				totalSupply:  big.NewInt(2000),
			},
			want:    big.NewInt(500),
			wantErr: false,
		},
		{
			name: "Test 2: When totalSupply < currentStake ",
			args: args{
				amount:       big.NewInt(500),
				currentStake: big.NewInt(4000),
				totalSupply:  big.NewInt(2000),
			},
			want:    big.NewInt(250),
			wantErr: false,
		},
		{
			name: "Test 3: When currentStake is 0",
			args: args{
				amount:       big.NewInt(500),
				currentStake: big.NewInt(0),
				totalSupply:  big.NewInt(2000),
			},
			want:    big.NewInt(0),
			wantErr: true,
		},
		{
			name: "Test 4: When values are high",
			args: args{
				amount:       big.NewInt(1).Mul(big.NewInt(500), big.NewInt(1e7)),
				currentStake: big.NewInt(1).Mul(big.NewInt(400), big.NewInt(1e8)),
				totalSupply:  big.NewInt(1).Mul(big.NewInt(4000), big.NewInt(1e8)),
			},
			want:    big.NewInt(1).Mul(big.NewInt(500), big.NewInt(1e8)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertRZRToSRZR(tt.args.amount, tt.args.currentStake, tt.args.totalSupply)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertRZRToSRZR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("ConvertRZRToSRZR() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertSRZRToRZR(t *testing.T) {
	type args struct {
		sAmount      *big.Int
		currentStake *big.Int
		totalSupply  *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1: When current stake totalSupply are equal",
			args: args{
				sAmount:      big.NewInt(500),
				currentStake: big.NewInt(1000),
				totalSupply:  big.NewInt(1000),
			},
			want: big.NewInt(500),
		},
		{
			name: "Test 2: When totalSupply < currentStake",
			args: args{
				sAmount:      big.NewInt(500),
				currentStake: big.NewInt(2000),
				totalSupply:  big.NewInt(1000),
			},
			want: big.NewInt(1000),
		},
		{
			name: "Test 3: When values are high",
			args: args{
				sAmount:      big.NewInt(1).Mul(big.NewInt(500), big.NewInt(1e8)),
				currentStake: big.NewInt(1).Mul(big.NewInt(400), big.NewInt(1e8)),
				totalSupply:  big.NewInt(1).Mul(big.NewInt(4000), big.NewInt(1e8)),
			},
			want: big.NewInt(1).Mul(big.NewInt(500), big.NewInt(1e7)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSRZRToRZR(tt.args.sAmount, tt.args.currentStake, tt.args.totalSupply); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertSRZRToRZR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRogueRandomValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test 1: Given a value, the function generates a random value less than or equal to that value",
			args: args{
				value: 10,
			},
			want: big.NewInt(10),
		},
		{
			name: "Test 2: Test for negative value",
			args: args{
				value: -10,
			},
			want: big.NewInt(0),
		},
		{
			name: "Test 3: Test for 0 value",
			args: args{
				value: 0,
			},
			want: big.NewInt(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRogueRandomValue(tt.args.value)
			if got.Cmp(tt.want) > 0 {
				t.Errorf("GetRogueRandomValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
