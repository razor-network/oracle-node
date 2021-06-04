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
	if num == nil {
		return big.NewInt(0)
	}
	decimalMultiplier := big.NewFloat(float64(core.DecimalsMultiplier))
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
