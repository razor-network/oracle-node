package types

import "razor/pkg/bindings"

type Block struct {
	Block        bindings.StructsBlock
	BlockMedians []uint32
}
