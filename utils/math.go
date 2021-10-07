package utils

import (
	"errors"
	"math"
	"math/big"
	"razor/core/types"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/ethclient"
	"modernc.org/sortutil"
)

func ConvertToNumber(num interface{}) (*big.Float, error) {
	if num == nil {
		return big.NewFloat(0), errors.New("no data provided")
	}
	switch v := num.(type) {
	case int:
		return big.NewFloat(float64(v)), nil
	case float64:
		return big.NewFloat(v), nil
	case string:
		convertedNumber, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Error("Error in converting from string to float: ", err)
			return big.NewFloat(0), nil
		}
		return big.NewFloat(convertedNumber), nil
	}
	return big.NewFloat(0), nil
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

func MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
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

func CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	if amountInWei.Cmp(balance) > 0 {
		log.Fatal("Not enough balance")
	}
	return amountInWei
}

func GetAmountInWei(amount *big.Int) *big.Int {
	amountInWei := big.NewInt(1).Mul(amount, big.NewInt(1e18))
	return amountInWei
}

func GetFractionalAmountInWei(amount *big.Int, power string) *big.Int {
	_power, err := new(big.Int).SetString(power, 10)
	if !err {
		log.Fatal("SetString: error")
	}
	amountInWei := big.NewInt(1).Mul(amount, big.NewInt(1).Exp(big.NewInt(10), _power, nil))
	return amountInWei
}

func AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	amount, err := flagSet.GetString("value")
	if err != nil {
		log.Fatal("Error in reading value", err)
	}
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		log.Fatal("SetString: error")
	}
	var amountInWei *big.Int
	if IsFlagPassed("pow") {
		power, _ := flagSet.GetString("pow")
		amountInWei = GetFractionalAmountInWei(_amount, power)
	} else {
		amountInWei = GetAmountInWei(_amount)
	}
	return amountInWei
}

func GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(amountInWei), new(big.Float).SetInt(big.NewInt(1e18)))
}

func Aggregate(client *ethclient.Client, address string, previousEpoch uint32, collection types.Collection) (*big.Int, error) {
	if len(collection.JobIDs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	var jobs []types.Job
	for _, id := range collection.JobIDs {
		job, err := GetActiveJob(client, address, id)
		if err != nil {
			log.Errorf("Error in fetching job %d: %s", id, err)
			continue
		}
		jobs = append(jobs, job)
	}
	dataToCommit, err := GetDataToCommitFromJobs(jobs)
	if err != nil || len(dataToCommit) == 0 {
		prevCommitmentData, err := FetchPreviousValue(client, address, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return big.NewInt(int64(prevCommitmentData)), nil
	}
	return performAggregation(dataToCommit, collection.AggregationMethod)
}

func performAggregation(data []*big.Int, aggregationMethod uint32) (*big.Int, error) {
	if len(data) == 0 {
		return nil, errors.New("aggregation cannot be performed for nil data")
	}
	// convention is 1 for median and 2 for mean
	switch aggregationMethod {
	case 1:
		sortutil.BigIntSlice.Sort(data)
		median := data[len(data)/2]
		return big.NewInt(median.Int64()), nil
	case 2:
		sum := CalculateSumOfArray(data)
		mean := sum.Div(sum, big.NewInt(int64(len(data))))
		return big.NewInt(mean.Int64()), nil
	}
	return nil, errors.New("invalid aggregation method")
}

func ConvertWeiToEth(data *big.Int) (*big.Float, error) {
	if data.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0), errors.New("cannot divide by 0")
	}
	dataInFloat := new(big.Float).SetInt(data)
	return dataInFloat.Quo(dataInFloat, big.NewFloat(1e18)).SetPrec(32), nil
}
