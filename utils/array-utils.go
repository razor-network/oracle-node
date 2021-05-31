package utils

import (
	"math/big"
)

func Contains(arr []*big.Int, val *big.Int) bool {
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

	for i := 0; i < len(arr1); i++ {
		if arr1[i].Cmp(arr2[i]) != 0 {
			return false
		}
	}
	return true
}
