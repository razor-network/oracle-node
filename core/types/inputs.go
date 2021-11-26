package types

import "math/big"

type UnstakeInput struct {
	Address    string
	Password   string
	ValueInWei *big.Int
	StakerId   uint32
}
