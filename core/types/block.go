//Package types include the different user defined items of possible different types in a single type
package types

import (
	"math/big"
	"razor/pkg/bindings"
)

type Block struct {
	Block        bindings.StructsBlock
	BlockMedians []*big.Int
}
