// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"
)

// CoinUtils is an autogenerated mock type for the CoinUtils type
type CoinUtils struct {
	mock.Mock
}

// Allowance provides a mock function with given fields: client, owner, spender
func (_m *CoinUtils) Allowance(client *ethclient.Client, owner common.Address, spender common.Address) (*big.Int, error) {
	ret := _m.Called(client, owner, spender)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address, common.Address) (*big.Int, error)); ok {
		return rf(client, owner, spender)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address, common.Address) *big.Int); ok {
		r0 = rf(client, owner, spender)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address, common.Address) error); ok {
		r1 = rf(client, owner, spender)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BalanceOf provides a mock function with given fields: client, account
func (_m *CoinUtils) BalanceOf(client *ethclient.Client, account common.Address) (*big.Int, error) {
	ret := _m.Called(client, account)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) (*big.Int, error)); ok {
		return rf(client, account)
	}
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) *big.Int); ok {
		r0 = rf(client, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address) error); ok {
		r1 = rf(client, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCoinUtils creates a new instance of CoinUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCoinUtils(t interface {
	mock.TestingT
	Cleanup(func())
}) *CoinUtils {
	mock := &CoinUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
