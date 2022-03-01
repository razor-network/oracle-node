package utils

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math"
)

func (*UtilsStruct) CreateMerkle(values []string) [][][]byte {
	var tree [][][]byte
	var leaves [][]byte

	for i := 0; i < len(values); i++ {
		leaves = append(leaves, solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{values[i]}))
	}

	level := leaves
	var nextLevel [][]byte
	tree = append(tree, level)

	for len(level) != 1 {
		for i := 0; i < len(level); i += 2 {
			if i+1 < len(level) {
				nextLevel = append(nextLevel, solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{level[i], level[i+1]}))
			} else {
				nextLevel = append(nextLevel, level[i])
			}
		}
		level = nextLevel
		tree = append(tree, level)
		nextLevel = nil
	}

	// Reverse the tree
	for i, j := 0, len(tree)-1; i < j; i, j = i+1, j-1 {
		tree[i], tree[j] = tree[j], tree[i]
	}

	return tree
}

func (*UtilsStruct) GetProofPath(tree [][][]byte, assetId uint16) [][]byte {
	index := assetId - 1
	var compactProofPath [][]byte
	for currentLevel := len(tree) - 1; currentLevel > 0; currentLevel-- {
		currentLevelNodes := tree[currentLevel]
		currentLevelCount := len(currentLevelNodes)
		if int(index) == currentLevelCount-1 && currentLevelCount%2 == 1 {
			index = uint16(math.Floor(float64(index / 2)))
			continue
		}
		if index%2 == 1 {
			compactProofPath = append(compactProofPath, currentLevelNodes[index-1])
		} else {
			compactProofPath = append(compactProofPath, currentLevelNodes[index+1])
		}
		index = uint16(math.Floor(float64(index / 2)))
	}
	return compactProofPath
}

func (*UtilsStruct) GetMerkleRoot(tree [][][]byte) []byte {
	return tree[0][0]
}
