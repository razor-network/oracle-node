package types

import "math/big"

type Configurations struct {
	Provider      string
	GasMultiplier float32
	ChainId       *big.Int
	DefaultPath   string
}
