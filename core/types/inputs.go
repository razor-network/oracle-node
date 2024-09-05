package types

import "math/big"

type UnstakeInput struct {
	Account    Account
	ValueInWei *big.Int
	StakerId   uint32
}

type RedeemBountyInput struct {
	Account  Account
	BountyId uint32
}

type TransferInput struct {
	Account    Account
	ToAddress  string
	ValueInWei *big.Int
	Balance    *big.Int
}

type CreateJobInput struct {
	Account      Account
	Name         string
	Url          string
	Selector     string
	Power        int8
	Weight       uint8
	SelectorType uint8
}

type CreateCollectionInput struct {
	Account     Account
	Name        string
	Aggregation uint32
	Power       int8
	JobIds      []uint
	Tolerance   uint32
}

type ExtendLockInput struct {
	Account  Account
	StakerId uint32
}

type ModifyCollectionInput struct {
	Account      Account
	CollectionId uint16
	Status       bool
}

type SetDelegationInput struct {
	Account      Account
	Status       bool
	StatusString string
	StakerId     uint32
	Commission   uint8
}

type UpdateCommissionInput struct {
	Account    Account
	Commission uint8
	StakerId   uint32
}
