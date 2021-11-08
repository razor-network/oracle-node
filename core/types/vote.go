package types

import "math/big"

type ElectedProposer struct {
	Iteration        int
	Stake            *big.Int
	StakerId         uint32
	BiggestInfluence *big.Int
	NumberOfStakers  uint32
	RandaoHash       [32]byte
	Epoch            uint32
}

type Commitment struct {
	Epoch          uint32
	CommitmentHash [32]byte
}
