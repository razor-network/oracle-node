package utils

import (
	"errors"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

func (*MerkleTreeStruct) CreateMerkle(values []*big.Int) ([][][]byte, error) {
	if len(values) == 0 {
		return [][][]byte{}, errors.New("values are nil, cannot create merkle tree")
	}
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

	return tree, nil
}

func (*MerkleTreeStruct) GetProofPath(tree [][][]byte, assetId uint16) [][32]byte {
	var compactProofPath [][32]byte
	for currentLevel := len(tree) - 1; currentLevel > 0; currentLevel-- {
		currentLevelNodes := tree[currentLevel]
		currentLevelCount := len(currentLevelNodes)
		if int(assetId) == currentLevelCount-1 && currentLevelCount%2 == 1 {
			assetId = assetId / 2
			continue
		}
		var node [32]byte
		if assetId%2 == 1 {
			copy(node[:], currentLevelNodes[assetId-1])
		} else {
			copy(node[:], currentLevelNodes[assetId+1])
		}
		compactProofPath = append(compactProofPath, node)
		assetId = assetId / 2
	}
	return compactProofPath
}

func (*MerkleTreeStruct) GetMerkleRoot(tree [][][]byte) [32]byte {
	var root [32]byte
	copy(root[:], tree[0][0])
	return root
}
