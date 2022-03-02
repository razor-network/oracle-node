// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	accounts "razor/accounts"

	abi "github.com/ethereum/go-ethereum/accounts/abi"

	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	bindings "razor/pkg/bindings"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	ecdsa "crypto/ecdsa"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	ethereum "github.com/ethereum/go-ethereum"

	fs "io/fs"

	io "io"

	mock "github.com/stretchr/testify/mock"

	retry "github.com/avast/retry-go"

	types "razor/core/types"
)

// OptionUtils is an autogenerated mock type for the OptionUtils type
type OptionUtils struct {
	mock.Mock
}

// Commitments provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) Commitments(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (types.Commitment, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.Commitment
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) types.Commitment); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(types.Commitment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertToNumber provides a mock function with given fields: _a0
func (_m *OptionUtils) ConvertToNumber(_a0 interface{}) (*big.Float, error) {
	ret := _m.Called(_a0)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(interface{}) *big.Float); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpochLimitForUpdateCommission provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) EpochLimitForUpdateCommission(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint16, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint16); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) EstimateGas(_a0 *ethclient.Client, _a1 context.Context, _a2 ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) FilterLogs(_a0 *ethclient.Client, _a1 context.Context, _a2 ethereum.FilterQuery) ([]coretypes.Log, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []coretypes.Log
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, ethereum.FilterQuery) []coretypes.Log); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coretypes.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveCollections provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetActiveCollections(_a0 *ethclient.Client, _a1 *bind.CallOpts) ([]uint16, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) []uint16); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAsset provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetAsset(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint16) (types.Asset, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.Asset
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint16) types.Asset); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(types.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetBlock(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultPath provides a mock function with given fields:
func (_m *OptionUtils) GetDefaultPath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastCommitted provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetEpochLastCommitted(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastRevealed provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetEpochLastRevealed(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluenceSnapshot provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) GetInfluenceSnapshot(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobFilePath provides a mock function with given fields:
func (_m *OptionUtils) GetJobFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumActiveCollections provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetNumActiveCollections(_a0 *ethclient.Client, _a1 *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumAssets provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetNumAssets(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint16, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint16); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumProposedBlocks provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetNumProposedBlocks(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (uint8, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint8); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumStakers provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetNumStakers(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrivateKey provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) GetPrivateKey(_a0 string, _a1 string, _a2 string, _a3 accounts.AccountInterface) *ecdsa.PrivateKey {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *ecdsa.PrivateKey
	if rf, ok := ret.Get(0).(func(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	return r0
}

// GetProposedBlock provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) GetProposedBlock(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandaoHash provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetRandaoHash(_a0 *ethclient.Client, _a1 *bind.CallOpts) ([32]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) [32]byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakeSnapshot provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) GetStakeSnapshot(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaker provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetStaker(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) bindings.StructsStaker); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerId provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetStakerId(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 common.Address) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, common.Address) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, common.Address) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalInfluenceRevealed provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetTotalInfluenceRevealed(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVote provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) GetVote(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (bindings.StructsVote, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsVote
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) bindings.StructsVote); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsVote)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) GetVoteValue(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint16, _a3 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint16, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint16, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Jobs provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) Jobs(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint16) (bindings.StructsJob, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsJob
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint16) bindings.StructsJob); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsJob)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Locks provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) Locks(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 common.Address, _a3 common.Address) (types.Locks, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 types.Locks
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) types.Locks); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(types.Locks)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Marshal provides a mock function with given fields: _a0
func (_m *OptionUtils) Marshal(_a0 interface{}) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxAltBlocks provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) MaxAltBlocks(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxCommission provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) MaxCommission(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinStake provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) MinStake(_a0 *ethclient.Client, _a1 *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAssetManager provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewAssetManager(_a0 common.Address, _a1 *ethclient.Client) (*bindings.AssetManager, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.AssetManager
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.AssetManager); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.AssetManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlockManager provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewBlockManager(_a0 common.Address, _a1 *ethclient.Client) (*bindings.BlockManager, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.BlockManager
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.BlockManager); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.BlockManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyedTransactorWithChainID provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewKeyedTransactorWithChainID(_a0 *ecdsa.PrivateKey, _a1 *big.Int) (*bind.TransactOpts, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, *big.Int) *bind.TransactOpts); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecdsa.PrivateKey, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRAZOR provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewRAZOR(_a0 common.Address, _a1 *ethclient.Client) (*bindings.RAZOR, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.RAZOR
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.RAZOR); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.RAZOR)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStakeManager provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewStakeManager(_a0 common.Address, _a1 *ethclient.Client) (*bindings.StakeManager, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.StakeManager
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.StakeManager); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakeManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStakedToken provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewStakedToken(_a0 common.Address, _a1 *ethclient.Client) (*bindings.StakedToken, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.StakedToken
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.StakedToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakedToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVoteManager provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) NewVoteManager(_a0 common.Address, _a1 *ethclient.Client) (*bindings.VoteManager, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.VoteManager
	if rf, ok := ret.Get(0).(func(common.Address, *ethclient.Client) *bindings.VoteManager); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.VoteManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *ethclient.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Pack provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) Pack(_a0 abi.ABI, _a1 string, _a2 ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(abi.ABI, string, ...interface{}) []byte); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abi.ABI, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: _a0
func (_m *OptionUtils) Parse(_a0 io.Reader) (abi.ABI, error) {
	ret := _m.Called(_a0)

	var r0 abi.ABI
	if rf, ok := ret.Get(0).(func(io.Reader) abi.ABI); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(abi.ABI)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingNonceAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) PendingNonceAt(_a0 *ethclient.Client, _a1 context.Context, _a2 common.Address) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, common.Address) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, common.Address) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: _a0
func (_m *OptionUtils) ReadAll(_a0 io.ReadCloser) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(io.ReadCloser) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.ReadCloser) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFile provides a mock function with given fields: _a0
func (_m *OptionUtils) ReadFile(_a0 string) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetryAttempts provides a mock function with given fields: _a0
func (_m *OptionUtils) RetryAttempts(_a0 uint) retry.Option {
	ret := _m.Called(_a0)

	var r0 retry.Option
	if rf, ok := ret.Get(0).(func(uint) retry.Option); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(retry.Option)
		}
	}

	return r0
}

// SortedProposedBlockIds provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *OptionUtils) SortedProposedBlockIds(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 *big.Int) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) uint32); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) SuggestGasPrice(_a0 *ethclient.Client, _a1 context.Context) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unmarshal provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) Unmarshal(_a0 []byte, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithdrawReleasePeriod provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) WithdrawReleasePeriod(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteFile provides a mock function with given fields: _a0, _a1, _a2
func (_m *OptionUtils) WriteFile(_a0 string, _a1 []byte, _a2 fs.FileMode) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, fs.FileMode) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
