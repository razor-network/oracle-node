package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Staker struct {
	AcceptDelegation                bool
	IsSlashed                       bool
	Commission                      uint8
	Id                              uint32
	Age                             uint32
	Address                         common.Address
	TokenAddress                    common.Address
	EpochFirstStakedOrLastPenalized uint32
	EpochCommissionLastUpdated      uint32
	Stake                           *big.Int
}

type BountyLock struct {
	RedeemAfter  uint32
	BountyHunter common.Address
	Amount       *big.Int
}
