package types

import "math/big"

type UnstakeInput struct {
	Address    string
	Password   string
	ValueInWei *big.Int
	StakerId   uint32
}

type RedeemBountyInput struct {
	Address  string
	Password string
	BountyId uint32
}

type UpdateCommissionInput struct {
	StakerId   uint32
	Address    string
	Password   string
	Commission uint8
	Config     Configurations
}
