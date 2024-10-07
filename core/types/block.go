package types

import (
	"math/big"
	"razor/pkg/bindings"
)

type Block struct {
	Block        bindings.StructsBlock
	BlockMedians []*big.Int
}

type DisputesStruct struct {
	LeafId           uint16
	LastVisitedValue *big.Int
	AccWeight        *big.Int
	Median           *big.Int
}

type ConfirmedBlock struct {
	Valid        bool
	ProposerId   uint32
	Iteration    *big.Int
	BiggestStake *big.Int
}
