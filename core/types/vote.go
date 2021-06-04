package types

import "math/big"

type ElectedProposer struct {
	Iteration       int
	Stake           *big.Int
	StakerId        *big.Int
	BiggestStake    *big.Int
	NumberOfStakers *big.Int
	BlockHashes     []byte
}
