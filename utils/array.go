package utils

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"modernc.org/sortutil"
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

func IsEqual(arr1 []*big.Int, arr2 []*big.Int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sortutil.BigIntSlice.Sort(arr1)
	sortutil.BigIntSlice.Sort(arr2)
	for i := 0; i < len(arr1); i++ {
		if arr1[i].Cmp(arr2[i]) != 0 {
			return false
		}
	}
	return true
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
