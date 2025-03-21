// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// BlockManagerInterface is an autogenerated mock type for the BlockManagerInterface type
type BlockManagerInterface struct {
	mock.Mock
}

// ClaimBlockReward provides a mock function with given fields: client, opts
func (_m *BlockManagerInterface) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(client, opts)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(client, opts)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(client, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts) error); ok {
		r1 = rf(client, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeBiggestStakeProposed provides a mock function with given fields: client, opts, epoch, blockIndex, correctBiggestStakerId
func (_m *BlockManagerInterface) DisputeBiggestStakeProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestStakerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, blockIndex, correctBiggestStakerId)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, blockIndex, correctBiggestStakerId)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, blockIndex, correctBiggestStakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) error); ok {
		r1 = rf(client, opts, epoch, blockIndex, correctBiggestStakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeCollectionIdShouldBeAbsent provides a mock function with given fields: client, opts, epoch, blockIndex, id, positionOfCollectionInBlock
func (_m *BlockManagerInterface) DisputeCollectionIdShouldBeAbsent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16, positionOfCollectionInBlock *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, blockIndex, id, positionOfCollectionInBlock)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16, *big.Int) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, blockIndex, id, positionOfCollectionInBlock)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, blockIndex, id, positionOfCollectionInBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16, *big.Int) error); ok {
		r1 = rf(client, opts, epoch, blockIndex, id, positionOfCollectionInBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeCollectionIdShouldBePresent provides a mock function with given fields: client, opts, epoch, blockIndex, id
func (_m *BlockManagerInterface) DisputeCollectionIdShouldBePresent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, blockIndex, id)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, blockIndex, id)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, blockIndex, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint16) error); ok {
		r1 = rf(client, opts, epoch, blockIndex, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeOnOrderOfIds provides a mock function with given fields: client, opts, epoch, blockIndex, index0, index1
func (_m *BlockManagerInterface) DisputeOnOrderOfIds(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, index0 *big.Int, index1 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, blockIndex, index0, index1)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int, *big.Int) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, blockIndex, index0, index1)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, blockIndex, index0, index1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int, *big.Int) error); ok {
		r1 = rf(client, opts, epoch, blockIndex, index0, index1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeDispute provides a mock function with given fields: client, opts, epoch, blockIndex, positionOfCollectionInBlock
func (_m *BlockManagerInterface) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, positionOfCollectionInBlock *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, blockIndex, positionOfCollectionInBlock)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, blockIndex, positionOfCollectionInBlock)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, blockIndex, positionOfCollectionInBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, *big.Int) error); ok {
		r1 = rf(client, opts, epoch, blockIndex, positionOfCollectionInBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GiveSorted provides a mock function with given fields: client, opts, epoch, leafId, sortedValues
func (_m *BlockManagerInterface) GiveSorted(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []*big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, leafId, sortedValues)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint16, []*big.Int) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, leafId, sortedValues)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint16, []*big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, leafId, sortedValues)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint16, []*big.Int) error); ok {
		r1 = rf(client, opts, epoch, leafId, sortedValues)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Propose provides a mock function with given fields: client, opts, epoch, ids, medians, iteration, biggestInfluencerId
func (_m *BlockManagerInterface) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, ids []uint16, medians []*big.Int, iteration *big.Int, biggestInfluencerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch, ids, medians, iteration, biggestInfluencerId)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, []uint16, []*big.Int, *big.Int, uint32) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch, ids, medians, iteration, biggestInfluencerId)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, []uint16, []*big.Int, *big.Int, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, epoch, ids, medians, iteration, biggestInfluencerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, []uint16, []*big.Int, *big.Int, uint32) error); ok {
		r1 = rf(client, opts, epoch, ids, medians, iteration, biggestInfluencerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetDispute provides a mock function with given fields: client, opts, epoch
func (_m *BlockManagerInterface) ResetDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, epoch)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) (*types.Transaction, error)); ok {
		return rf(client, opts, epoch)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlockManagerInterface creates a new instance of BlockManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockManagerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockManagerInterface {
	mock := &BlockManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
