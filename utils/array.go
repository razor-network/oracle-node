package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"strings"
)

func Contains(slice interface{}, val interface{}) bool {
	switch slice := slice.(type) {
	case []int:
		for _, value := range slice {
			if value == val {
				return true
			}
		}
	case []uint32:
		for _, value := range slice {
			if value == val {
				return true
			}
		}
	case []string:
		for _, value := range slice {
			if value == val {
				return true
			}
		}
	case []uint16:
		for _, value := range slice {
			if value == val {
				return true
			}
		}
	}
	return false
}

func ContainsBigInteger(arr []*big.Int, num *big.Int) bool {
	if num == nil {
		return false
	}
	for _, value := range arr {
		if value.Cmp(num) == 0 {
			return true
		}
	}
	return false
}

func IsEqual(arr1 []*big.Int, arr2 []*big.Int) (bool, int) {
	if len(arr1) > len(arr2) {
		return false, len(arr2)
	} else if len(arr1) < len(arr2) {
		return false, len(arr1)
	}
	for i := 0; i < len(arr1); i++ {
		if arr2[i].Cmp(arr1[i]) != 0 {
			return false, i
		}
	}
	return true, -1
}

func IsEqualByte(arr1 []byte, arr2 []byte) (bool, int) {
	if len(arr1) > len(arr2) {
		return false, len(arr2)
	} else if len(arr1) < len(arr2) {
		return false, len(arr1)
	}
	for i := 0; i < len(arr1); i++ {
		if arr2[i] != arr1[i] {
			return false, i
		}
	}
	return true, -1
}

// IsMissing checks for elements present in 1st array but not in second
func IsMissing(arr1 []uint16, arr2 []uint16) (bool, int, uint16) {
	arrayMap := make(map[uint16]bool)
	for i := 0; i < len(arr2); i++ {
		arrayMap[arr2[i]] = true
	}
	for i := 0; i < len(arr1); i++ {
		if !arrayMap[arr1[i]] {
			return true, i, arr1[i]
		}
	}
	return false, -1, 0
}

func IsSorted(values []uint16) (bool, int, int) {
	if values == nil {
		return true, -1, -1
	}
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			return false, i, i + 1
		}
	}
	return true, -1, -1
}

func IndexOf(array []uint32, value uint32) int {
	for arrayIndex, arrayVal := range array {
		if arrayVal == value {
			return arrayIndex
		}
	}
	return -1
}

func GetDataInBytes(data []*big.Int) [][]byte {
	if len(data) == 0 {
		return nil
	}
	var dataInBytes [][]byte
	for _, datum := range data {
		dataInBytes = append(dataInBytes, math2.U256Bytes(datum))
	}
	return dataInBytes
}

func ConvertUint32ArrayToBigIntArray(uint32Array []uint32) []*big.Int {
	var arr []*big.Int
	for _, datum := range uint32Array {
		arr = append(arr, big.NewInt(int64(datum)))
	}
	return arr
}

func CalculateWeightedSum(data []*big.Int, weight []uint8) *big.Int {
	sum := big.NewInt(0)
	for index, datum := range data {
		weighedData := big.NewInt(1).Mul(datum, big.NewInt(int64(weight[index])))
		sum.Add(sum, weighedData)
	}
	return sum
}

func CalculateSumOfUint8Array(data []uint8) uint {
	sum := uint(0)
	if len(data) == 0 {
		return sum
	}
	for _, datum := range data {
		sum += uint(datum)
	}
	return sum
}

func ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	var arr []uint16
	for _, datum := range uintArr {
		arr = append(arr, uint16(datum))
	}
	return arr
}

func ContainsStringFromArray(source string, subStringArray []string) bool {
	for i := 0; i < len(subStringArray); i++ {
		if strings.Contains(source, subStringArray[i]) {
			return true
		}
	}
	return false
}
