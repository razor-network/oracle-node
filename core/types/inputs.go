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

type TransferInput struct {
	FromAddress string
	ToAddress   string
	Password    string
	ValueInWei  *big.Int
	Balance     *big.Int
}
