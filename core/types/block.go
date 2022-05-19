package types

import (
	"math/big"
	"razor/pkg/bindings"
)

type Block struct {
	Block        bindings.StructsBlock
	BlockMedians []*big.Int
}
