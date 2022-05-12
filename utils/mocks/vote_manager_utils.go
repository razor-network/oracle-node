// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	ethclient "github.com/ethereum/go-ethereum/ethclient"
	mock "github.com/stretchr/testify/mock"

	types "razor/core/types"
)

// VoteManagerUtils is an autogenerated mock type for the VoteManagerUtils type
type VoteManagerUtils struct {
	mock.Mock
}

// Commitments provides a mock function with given fields: _a0, _a1
func (_m *VoteManagerUtils) Commitments(_a0 *ethclient.Client, _a1 uint32) (types.Commitment, error) {
	ret := _m.Called(_a0, _a1)

	var r0 types.Commitment
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) types.Commitment); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(types.Commitment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastCommitted provides a mock function with given fields: client, stakerId
func (_m *VoteManagerUtils) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	ret := _m.Called(client, stakerId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastRevealed provides a mock function with given fields: client, stakerId
func (_m *VoteManagerUtils) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	ret := _m.Called(client, stakerId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluenceSnapshot provides a mock function with given fields: client, epoch, stakerId
func (_m *VoteManagerUtils) GetInfluenceSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	ret := _m.Called(client, epoch, stakerId)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(client, epoch, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, epoch, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSaltFromBlockchain provides a mock function with given fields: client
func (_m *VoteManagerUtils) GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error) {
	ret := _m.Called(client)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client) [32]byte); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakeSnapshot provides a mock function with given fields: client, epoch, stakerId
func (_m *VoteManagerUtils) GetStakeSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	ret := _m.Called(client, epoch, stakerId)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(client, epoch, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, epoch, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalInfluenceRevealed provides a mock function with given fields: client, epoch, medianIndex
func (_m *VoteManagerUtils) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	ret := _m.Called(client, epoch, medianIndex)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint16) *big.Int); ok {
		r0 = rf(client, epoch, medianIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint16) error); ok {
		r1 = rf(client, epoch, medianIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: client, epoch, stakerId, medianIndex
func (_m *VoteManagerUtils) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (uint32, error) {
	ret := _m.Called(client, epoch, stakerId, medianIndex)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32, uint16) uint32); ok {
		r0 = rf(client, epoch, stakerId, medianIndex)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32, uint16) error); ok {
		r1 = rf(client, epoch, stakerId, medianIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToAssign provides a mock function with given fields: client
func (_m *VoteManagerUtils) ToAssign(client *ethclient.Client) (uint16, error) {
	ret := _m.Called(client)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
