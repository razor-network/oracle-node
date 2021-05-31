package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"strconv"
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

func MultiplyToEightDecimals(num *big.Float) *big.Int {
	decimalMultiplier := big.NewFloat(float64(core.DecimalsMultiplier))
	value := big.NewFloat(1).Mul(num, decimalMultiplier)
	result := new(big.Int)
	value.Int(result)
	return result
}

func MultiplyFloatAndBigInt(gas *big.Int, val float64) *big.Int {
	value := new(big.Float)
	value.SetFloat64(val)
	conversionInt := new(big.Float)
	conversionInt.SetInt(gas)
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

func GetAmountWithChecks(amount string, balance *big.Int) *big.Int {
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		log.Fatal("SetString: error")
	}

	amountInWei := big.NewInt(1).Mul(_amount, big.NewInt(1e18))

	if amountInWei.Cmp(balance) > 0 {
		log.Fatal("Not enough balance")
	}
	return amountInWei
}
