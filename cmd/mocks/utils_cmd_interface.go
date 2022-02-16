// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	accounts "github.com/ethereum/go-ethereum/accounts"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	bindings "razor/pkg/bindings"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	types "razor/core/types"
)

// UtilsCmdInterface is an autogenerated mock type for the UtilsCmdInterface type
type UtilsCmdInterface struct {
	mock.Mock
}

// Approve provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) Approve(_a0 types.TransactionOptions) (common.Hash, error) {
	ret := _m.Called(_a0)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.TransactionOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssignAmountInWei provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) AssignAmountInWei(_a0 *pflag.FlagSet) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AutoUnstakeAndWithdraw provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) AutoUnstakeAndWithdraw(_a0 *ethclient.Client, _a1 types.Account, _a2 *big.Int, _a3 types.Configurations) {
	_m.Called(_a0, _a1, _a2, _a3)
}

// AutoWithdraw provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) AutoWithdraw(_a0 types.TransactionOptions, _a1 uint32) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.TransactionOptions, uint32) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CalculateSecret provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) CalculateSecret(_a0 types.Account, _a1 uint32) []byte {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.Account, uint32) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// CheckCurrentStatus provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) CheckCurrentStatus(_a0 *ethclient.Client, _a1 uint16) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimBlockReward provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ClaimBlockReward(_a0 types.TransactionOptions) (common.Hash, error) {
	ret := _m.Called(_a0)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.TransactionOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimBounty provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) ClaimBounty(_a0 types.Configurations, _a1 *ethclient.Client, _a2 types.RedeemBountyInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.Configurations, *ethclient.Client, types.RedeemBountyInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Configurations, *ethclient.Client, types.RedeemBountyInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *UtilsCmdInterface) Commit(_a0 *ethclient.Client, _a1 []*big.Int, _a2 []byte, _a3 types.Account, _a4 types.Configurations) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, []*big.Int, []byte, types.Account, types.Configurations) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, []*big.Int, []byte, types.Account, types.Configurations) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) Create(_a0 string) (accounts.Account, error) {
	ret := _m.Called(_a0)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(string) accounts.Account); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCollection provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) CreateCollection(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateCollectionInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) CreateJob(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateJobInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateJobInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateJobInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delegate provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) Delegate(_a0 types.TransactionOptions, _a1 uint32) (common.Hash, error) {
	ret := _m.Called(_a0, _a1)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.TransactionOptions, uint32) common.Hash); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.TransactionOptions, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOverrideJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) DeleteOverrideJob(_a0 uint16) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispute provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *UtilsCmdInterface) Dispute(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.Account, _a3 uint32, _a4 uint8, _a5 int) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteClaimBounty provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteClaimBounty(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteCollectionList provides a mock function with given fields:
func (_m *UtilsCmdInterface) ExecuteCollectionList() {
	_m.Called()
}

// ExecuteCreate provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteCreate(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteCreateCollection provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteCreateCollection(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteCreateJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteCreateJob(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteDelegate provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteDelegate(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteDeleteOverrideJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteDeleteOverrideJob(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteExtendLock provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteExtendLock(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteImport provides a mock function with given fields:
func (_m *UtilsCmdInterface) ExecuteImport() {
	_m.Called()
}

// ExecuteJobList provides a mock function with given fields:
func (_m *UtilsCmdInterface) ExecuteJobList() {
	_m.Called()
}

// ExecuteListAccounts provides a mock function with given fields:
func (_m *UtilsCmdInterface) ExecuteListAccounts() {
	_m.Called()
}

// ExecuteModifyAssetStatus provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteModifyAssetStatus(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteOverrideJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteOverrideJob(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteSetDelegation provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteSetDelegation(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteStake provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteStake(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteStakerinfo provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteStakerinfo(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteTransfer provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteTransfer(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteUnstake provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteUnstake(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteUpdateCollection provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteUpdateCollection(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteUpdateCommission provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteUpdateCommission(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteUpdateJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteUpdateJob(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteVote provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteVote(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteWithdraw provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) ExecuteWithdraw(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExtendLock provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) ExtendLock(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.ExtendLockInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.ExtendLockInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.ExtendLockInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAmountInSRZRs provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) GetAmountInSRZRs(_a0 *ethclient.Client, _a1 string, _a2 bindings.StructsStaker, _a3 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, bindings.StructsStaker, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, bindings.StructsStaker, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBiggestStakeAndId provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) GetBiggestStakeAndId(_a0 *ethclient.Client, _a1 string, _a2 uint32) (*big.Int, uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32) uint32); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*ethclient.Client, string, uint32) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBufferPercent provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetBufferPercent() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionList provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) GetCollectionList(_a0 *ethclient.Client) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommitDataFileName provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) GetCommitDataFileName(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigData provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetConfigData() (types.Configurations, error) {
	ret := _m.Called()

	var r0 types.Configurations
	if rf, ok := ret.Get(0).(func() types.Configurations); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Configurations)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochAndState provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) GetEpochAndState(_a0 *ethclient.Client) (uint32, int64, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*ethclient.Client) int64); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*ethclient.Client) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetGasLimit provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetGasLimit() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetGasPrice() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIteration provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) GetIteration(_a0 *ethclient.Client, _a1 types.ElectedProposer) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.ElectedProposer) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetJobList provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) GetJobList(_a0 *ethclient.Client) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLastProposedEpoch provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) GetLastProposedEpoch(_a0 *ethclient.Client, _a1 *big.Int, _a2 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *big.Int, uint32) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *big.Int, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogLevel provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetLogLevel() (string, error) {
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

// GetMedianDataFileName provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) GetMedianDataFileName(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMultiplier provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetMultiplier() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvider provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetProvider() (string, error) {
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

// GetSortedVotes provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) GetSortedVotes(_a0 *ethclient.Client, _a1 string, _a2 uint16, _a3 uint32) ([]*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint16, uint32) []*big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint16, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerInfo provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) GetStakerInfo(_a0 *ethclient.Client, _a1 uint32) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWaitTime provides a mock function with given fields:
func (_m *UtilsCmdInterface) GetWaitTime() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GiveSorted provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *UtilsCmdInterface) GiveSorted(_a0 *ethclient.Client, _a1 *bindings.BlockManager, _a2 *bind.TransactOpts, _a3 uint32, _a4 uint16, _a5 []uint32) {
	_m.Called(_a0, _a1, _a2, _a3, _a4, _a5)
}

// HandleBlock provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *UtilsCmdInterface) HandleBlock(_a0 *ethclient.Client, _a1 types.Account, _a2 *big.Int, _a3 types.Configurations, _a4 types.Rogue) {
	_m.Called(_a0, _a1, _a2, _a3, _a4)
}

// HandleCommitState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) HandleCommitState(_a0 *ethclient.Client, _a1 uint32, _a2 types.Rogue) ([]*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, types.Rogue) []*big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, types.Rogue) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleDispute provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) HandleDispute(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.Account, _a3 uint32) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.Account, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleExit provides a mock function with given fields:
func (_m *UtilsCmdInterface) HandleExit() {
	_m.Called()
}

// HandleRevealState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) HandleRevealState(_a0 *ethclient.Client, _a1 bindings.StructsStaker, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, bindings.StructsStaker, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImportAccount provides a mock function with given fields:
func (_m *UtilsCmdInterface) ImportAccount() (accounts.Account, error) {
	ret := _m.Called()

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func() accounts.Account); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InfluencedMedian provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) InfluencedMedian(_a0 []*big.Int, _a1 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func([]*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// IsElectedProposer provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) IsElectedProposer(_a0 types.ElectedProposer, _a1 *big.Int) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.ElectedProposer, *big.Int) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListAccounts provides a mock function with given fields:
func (_m *UtilsCmdInterface) ListAccounts() ([]accounts.Account, error) {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) MakeBlock(_a0 *ethclient.Client, _a1 string, _a2 types.Rogue) ([]uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, types.Rogue) []uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, types.Rogue) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyAssetStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) ModifyAssetStatus(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.ModifyAssetInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.ModifyAssetInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.ModifyAssetInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OverrideJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) OverrideJob(_a0 *types.StructsJob) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.StructsJob) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Propose provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *UtilsCmdInterface) Propose(_a0 *ethclient.Client, _a1 types.Account, _a2 types.Configurations, _a3 uint32, _a4 uint32, _a5 types.Rogue) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Account, types.Configurations, uint32, uint32, types.Rogue) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Account, types.Configurations, uint32, uint32, types.Rogue) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reveal provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *UtilsCmdInterface) Reveal(_a0 *ethclient.Client, _a1 []*big.Int, _a2 []byte, _a3 types.Account, _a4 string, _a5 types.Configurations) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConfig provides a mock function with given fields: flagSet
func (_m *UtilsCmdInterface) SetConfig(flagSet *pflag.FlagSet) error {
	ret := _m.Called(flagSet)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) error); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDelegation provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) SetDelegation(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.SetDelegationInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.SetDelegationInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.SetDelegationInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StakeCoins provides a mock function with given fields: _a0
func (_m *UtilsCmdInterface) StakeCoins(_a0 types.TransactionOptions) (common.Hash, error) {
	ret := _m.Called(_a0)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.TransactionOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) Transfer(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.TransferInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.TransferInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.TransferInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unstake provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) Unstake(_a0 types.Configurations, _a1 *ethclient.Client, _a2 types.UnstakeInput) (types.TransactionOptions, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.TransactionOptions
	if rf, ok := ret.Get(0).(func(types.Configurations, *ethclient.Client, types.UnstakeInput) types.TransactionOptions); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(types.TransactionOptions)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Configurations, *ethclient.Client, types.UnstakeInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCollection provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) UpdateCollection(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateCollectionInput, _a3 uint16) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput, uint16) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCommission provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) UpdateCommission(_a0 types.Configurations, _a1 *ethclient.Client, _a2 types.UpdateCommissionInput) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Configurations, *ethclient.Client, types.UpdateCommissionInput) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateJob provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) UpdateJob(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateJobInput, _a3 uint16) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateJobInput, uint16) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateJobInput, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Vote provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *UtilsCmdInterface) Vote(_a0 context.Context, _a1 types.Configurations, _a2 *ethclient.Client, _a3 types.Rogue, _a4 types.Account) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Configurations, *ethclient.Client, types.Rogue, types.Account) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForAppropriateState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) WaitForAppropriateState(_a0 *ethclient.Client, _a1 string, _a2 ...int) (uint32, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, ...int) uint32); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, ...int) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitIfCommitState provides a mock function with given fields: _a0, _a1
func (_m *UtilsCmdInterface) WaitIfCommitState(_a0 *ethclient.Client, _a1 string) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Withdraw provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterface) Withdraw(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawFunds provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsCmdInterface) WithdrawFunds(_a0 *ethclient.Client, _a1 types.Account, _a2 types.Configurations, _a3 uint32) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Account, types.Configurations, uint32) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Account, types.Configurations, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
