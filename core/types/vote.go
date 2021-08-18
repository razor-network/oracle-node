package types

import "math/big"

type ElectedProposer struct {
	Iteration        int
	Stake            *big.Int
	StakerId         *big.Int
	BiggestInfluence *big.Int
	NumberOfStakers  *big.Int
	RandaoHash       [32]byte
}
