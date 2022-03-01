package types

import "math/big"

type ElectedProposer struct {
	Iteration       int
	Stake           *big.Int
	StakerId        uint32
	BiggestStake    *big.Int
	NumberOfStakers uint32
	Salt            [32]byte
	Epoch           uint32
}

type Commitment struct {
	Epoch          uint32
	CommitmentHash [32]byte
}

type Rogue struct {
	IsRogue   bool
	RogueMode []string
}

type CommitData struct {
	AssignedCollections   map[int]bool
	SeqAllottedCollections []*big.Int
	Leaves                []string
}
