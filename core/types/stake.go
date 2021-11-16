package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Staker struct {
	AcceptDelegation                bool
	Commission                      uint8
	Address                         common.Address
	TokenAddress                    common.Address
	Id                              uint32
	Age                             uint32
	EpochFirstStakedOrLastPenalized uint32
	Stake                           *big.Int
}
