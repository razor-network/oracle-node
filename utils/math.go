package utils

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	mathRand "math/rand"
	"razor/core"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ConvertToNumber(num interface{}, returnType string) (*big.Float, error) {
	if num == nil {
		return big.NewFloat(0), errors.New("no data provided")
	}
	switch v := num.(type) {
	case int:
		return big.NewFloat(float64(v)), nil
	case float64:
		return big.NewFloat(v), nil
	case string:
		if strings.ToLower(returnType) == core.HexReturnType {
			return ConvertHexToBigFloat(v)
		}
		if isHexArrayPattern(returnType) {
			return HandleHexArray(v, returnType)
		}
		convertedNumber, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Error("Error in converting from string to float: ", err)
			return big.NewFloat(0), err
		}
		return big.NewFloat(convertedNumber), nil
	}
	return big.NewFloat(0), errors.New("unsupported type provided")
}

func ConvertHexToBigFloat(hexString string) (*big.Float, error) {
	hexValue := strings.TrimPrefix(hexString, "0x")
	hexValueUint64, err := strconv.ParseUint(hexValue, 16, 64)
	if err != nil {
		log.Errorf("Error in converting hex value %v to uint64: %v", hexValue, err)
		return big.NewFloat(0), err
	}
	return big.NewFloat(float64(hexValueUint64)), nil
}

func MultiplyWithPower(num *big.Float, power int8) *big.Int {
	if num == nil {
		return big.NewInt(0)
	}
	decimalMultiplier := big.NewFloat(math.Pow(10, float64(power)))
	value := big.NewFloat(1).Mul(num, decimalMultiplier)
	result := new(big.Int)
	value.Int(result)
	return result
}

func (*UtilsStruct) MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
	if bigIntVal == nil || floatingVal == 0 {
		return big.NewInt(0)
	}
	value := new(big.Float)
	value.SetFloat64(floatingVal)
	conversionInt := new(big.Float)
	conversionInt.SetInt(bigIntVal)
	value.Mul(value, conversionInt)
	result := new(big.Int)
	value.Int(result)
	return result
}

func AllZero(bytesValue [32]byte) bool {
	for _, value := range bytesValue {
		if value != 0 {
			return false
		}
	}
	return true
}

func (*UtilsStruct) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	if amountInWei.Cmp(balance) > 0 {
		log.Fatal("Not enough razor balance")
	}
	return amountInWei
}

func GetAmountInWei(amount *big.Int) *big.Int {
	amountInWei := big.NewInt(1).Mul(amount, big.NewInt(1e18))
	return amountInWei
}

func GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(amountInWei), new(big.Float).SetInt(big.NewInt(1e18)))
}

func performAggregation(data []*big.Int, weight []uint8, aggregationMethod uint32) (*big.Int, error) {
	if len(data) == 0 {
		return nil, errors.New("aggregation cannot be performed for nil data")
	}
	totalWeight := CalculateSumOfUint8Array(weight)
	// convention is 1 for median and 2 for mean
	switch aggregationMethod {
	case 1:
		return calculateWeightedMedian(data, weight, totalWeight), nil
	case 2:
		weightedSum := CalculateWeightedSum(data, weight)
		weightedMean := weightedSum.Div(weightedSum, big.NewInt(int64(totalWeight)))
		return weightedMean, nil
	}
	return nil, errors.New("invalid aggregation method")
}
func calculateWeightedMedian(data []*big.Int, weight []uint8, totalWeight uint) *big.Int {
	if len(data) == 0 || len(weight) == 0 || totalWeight == 0 {
		return nil
	}
	fractionalWeights := getFractionalWeight(weight, totalWeight)
	//Create a pair of [data, weight]
	var pairs [][]interface{}
	for i := 0; i < len(data); i++ {
		pairs = append(pairs, []interface{}{data[i], fractionalWeights[i]})
	}
	//Sort the weight according to the data in increasing order
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i][0].(*big.Int).Cmp(pairs[j][0].(*big.Int)) < 0
	})

	sum := float32(0)
	for _, pair := range pairs {
		//Calculate the sum of weights from the sorted pair
		sum += pair[1].(float32)
		//If the sum exceeds 0.5 then that pair contains the median data
		if sum >= 0.5 {
			return pair[0].(*big.Int)
		}
	}
	return nil
}

func getFractionalWeight(weights []uint8, totalWeight uint) []float32 {
	if len(weights) == 0 || totalWeight == 0 {
		return nil
	}
	var fractionalWeight []float32
	for _, weight := range weights {
		fractionalWeight = append(fractionalWeight, float32(weight)/float32(totalWeight))
	}
	return fractionalWeight
}

func ConvertWeiToEth(data *big.Int) (*big.Float, error) {
	if data.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0), errors.New("cannot divide by 0")
	}
	dataInFloat := new(big.Float).SetInt(data)
	return dataInFloat.Quo(dataInFloat, big.NewFloat(1e18)).SetPrec(32), nil
}

func ConvertRZRToSRZR(amount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error) {
	if currentStake.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), errors.New("current stake is 0")
	}
	return big.NewInt(1).Div(big.NewInt(1).Mul(amount, totalSupply), currentStake), nil
}

func ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int {
	return big.NewInt(1).Div(big.NewInt(1).Mul(sAmount, currentStake), totalSupply)
}

func (*UtilsStruct) GetRogueRandomValue(value int) *big.Int {
	if value <= 0 {
		return big.NewInt(0)
	}
	rogueRandomValue, _ := rand.Int(rand.Reader, big.NewInt(int64(value)))
	return rogueRandomValue
}

func GetRogueRandomMedianValue() uint32 {
	rogueRandomMedianValue, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return uint32(rogueRandomMedianValue.Int64())
}

func Shuffle(slice []uint32) []uint32 {
	copiedSlice := make([]uint32, len(slice))
	copy(copiedSlice, slice)
	r := mathRand.New(mathRand.NewSource(time.Now().Unix()))
	for n := len(copiedSlice); n > 0; n-- {
		randIndex := r.Intn(n)
		copiedSlice[n-1], copiedSlice[randIndex] = copiedSlice[randIndex], copiedSlice[n-1]
	}
	return copiedSlice
}

func HandleHexArray(hexStr string, returnType string) (*big.Float, error) {
	decodedHexArray, err := decodeHexString(hexStr)
	if err != nil {
		log.Error("Error in decoding hex array: ", err)
		return big.NewFloat(0), err
	}
	log.Info("HandleHexArray: decoded hex array: ", decodedHexArray)

	index, err := extractIndex(returnType)
	if err != nil {
		log.Error("Error in extracting value from decoded hex array: ", err)
		return big.NewFloat(0), err
	}
	log.Debug("HandleHexArray: extracted index: ", index)

	// Check if index is within the bounds of decodedHexArray
	if index < 0 || index >= len(decodedHexArray) {
		log.Error("extracted index is out of bounds for decoded hex array")
		return big.NewFloat(0), errors.New("extracted index is out of bounds")
	}

	// decodedHexArray[index] returns value in wei, so it needs to be converted to eth
	valueInEth, err := ConvertWeiToEth(decodedHexArray[index])
	if err != nil {
		log.Error("Error in converting wei to eth: ", err)
		return big.NewFloat(0), err
	}

	return valueInEth, nil
}

func decodeHexString(hexStr string) ([]*big.Int, error) {
	// Remove the "0x" prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")
	// The length of uint256 in hex (32 bytes)
	const uint256HexLength = 64

	// Make sure the string length is at least enough for the offset and length
	if len(hexStr) < 2*uint256HexLength {
		return nil, errors.New("hex string too short to contain valid data")
	}

	// Getting the starting position of the array data (skipping this step as per Ethereum ABI encoding)
	// Skip the length of the array (next 32 bytes)
	lengthStr := hexStr[uint256HexLength : 2*uint256HexLength]
	length, success := new(big.Int).SetString(lengthStr, 16)
	if !success {
		log.Error("Invalid length of the array from the hex string")
		return nil, errors.New("invalid length")
	}

	// The remaining part of the string are the uint256 values
	valuesStr := hexStr[2*uint256HexLength:]

	// Each value is 32 bytes long, so check if the length matches
	if len(valuesStr) != int(length.Int64())*uint256HexLength {
		return nil, errors.New("data length does not match length specifier")
	}

	var values []*big.Int
	for i := 0; i < int(length.Int64()); i++ {
		start := i * uint256HexLength
		end := start + uint256HexLength
		n := new(big.Int)
		n, success := n.SetString(valuesStr[start:end], 16)
		if !success {
			log.Errorf("Invalid uint256 value at index %d", i)
			return nil, errors.New("invalid uint256 value at index")
		}
		values = append(values, n)
	}

	return values, nil
}

func extractIndex(s string) (int, error) {
	re := regexp.MustCompile(core.HexArrayExtractIndexRegex)

	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, errors.New("no index found in string")
	}

	// Converting the captured substring to an integer
	index, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, errors.New("invalid index format in string")
	}

	return index, nil
}

func isHexArrayPattern(s string) bool {
	pattern := core.HexArrayReturnType
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
