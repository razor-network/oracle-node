package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func Contains(arr []*big.Int, val *big.Int) bool {
	if val == nil || len(arr) == 0 {
		return false
	}
	for _, value := range arr {
		if value.Cmp(val) == 0 {
			return true
		}
	}
	return false
}

func IsEqual(arr1 []uint32, arr2 []uint32) (bool, int) {
	if len(arr1) > len(arr2) {
		return false, len(arr2) + 1
	} else if len(arr1) < len(arr2) {
		return false, len(arr1) + 1
	}
	for i := 0; i < len(arr1); i++ {
		if arr2[i] != arr1[i] {
			return false, i + 1
		}
	}
	return true, -1
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

func ConvertToBigIntArray(data []string) []*big.Int {
	var bigIntArray []*big.Int

	for _, datum := range data {
		bigData, ok := new(big.Int).SetString(datum, 10)
		if !ok {
			log.Fatal("SetString: error")
		}
		bigIntArray = append(bigIntArray, bigData)
	}

	return bigIntArray
}

func ConvertBigIntArrayToUint32Array(bigIntArray []*big.Int) []uint32 {
	var arr []uint32
	for _, datum := range bigIntArray {
		arr = append(arr, uint32(datum.Int64()))
	}
	return arr
}

func CalculateSumOfArray(data []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, datum := range data {
		sum.Add(sum, datum)
	}
	return sum
}

func ConvertUintArrayToUint8Array(uintArr []uint) []uint8 {
	var arr []uint8
	for _, datum := range uintArr {
		arr = append(arr, uint8(datum))
	}
	return arr
}
