//Package types include the different user defined items of possible different types in a single type
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

type CreateJobInput struct {
	Address      string
	Password     string
	Name         string
	Url          string
	Selector     string
	Power        int8
	Weight       uint8
	SelectorType uint8
}

type CreateCollectionInput struct {
	Address     string
	Name        string
	Password    string
	Aggregation uint32
	Power       int8
	JobIds      []uint
	Tolerance   uint32
}

type ExtendLockInput struct {
	Address  string
	Password string
	StakerId uint32
}

type ModifyCollectionInput struct {
	Address      string
	Password     string
	CollectionId uint16
	Status       bool
}

type SetDelegationInput struct {
	Address      string
	Password     string
	Status       bool
	StatusString string
	StakerId     uint32
	Commission   uint8
}

type UpdateCommissionInput struct {
	Address    string
	Password   string
	Commission uint8
	StakerId   uint32
}
