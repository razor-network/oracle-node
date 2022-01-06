// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TokenManagerInterfaceMockery is an autogenerated mock type for the TokenManagerInterfaceMockery type
type TokenManagerInterfaceMockery struct {
	mock.Mock
}

// Allowance provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *TokenManagerInterfaceMockery) Allowance(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 common.Address, _a3 common.Address) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Approve provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *TokenManagerInterfaceMockery) Approve(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 common.Address, _a3 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *TokenManagerInterfaceMockery) Transfer(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 common.Address, _a3 *big.Int) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
