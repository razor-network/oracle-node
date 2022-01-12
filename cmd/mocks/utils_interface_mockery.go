// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	time "time"

	types "razor/core/types"
)

// UtilsInterfaceMockery is an autogenerated mock type for the UtilsInterfaceMockery type
type UtilsInterfaceMockery struct {
	mock.Mock
}

// AllZero provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) AllZero(_a0 [32]byte) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func([32]byte) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AssignPassword provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) AssignPassword(_a0 *pflag.FlagSet) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AssignStakerId provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) AssignStakerId(_a0 *pflag.FlagSet, _a1 *ethclient.Client, _a2 string) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, *ethclient.Client, string) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet, *ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateBlockTime provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) CalculateBlockTime(_a0 *ethclient.Client) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// CheckAmountAndBalance provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) CheckAmountAndBalance(_a0 *big.Int, _a1 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// CheckEthBalanceIsZero provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) CheckEthBalanceIsZero(_a0 *ethclient.Client, _a1 string) {
	_m.Called(_a0, _a1)
}

// ConnectToClient provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) ConnectToClient(_a0 string) *ethclient.Client {
	ret := _m.Called(_a0)

	var r0 *ethclient.Client
	if rf, ok := ret.Get(0).(func(string) *ethclient.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethclient.Client)
		}
	}

	return r0
}

// ConvertBigIntArrayToUint32Array provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) ConvertBigIntArrayToUint32Array(_a0 []*big.Int) []uint32 {
	ret := _m.Called(_a0)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func([]*big.Int) []uint32); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	return r0
}

// ConvertRZRToSRZR provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) ConvertRZRToSRZR(_a0 *big.Int, _a1 *big.Int, _a2 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertSRZRToRZR provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) ConvertSRZRToRZR(_a0 *big.Int, _a1 *big.Int, _a2 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// ConvertUintArrayToUint16Array provides a mock function with given fields: uintArr
func (_m *UtilsInterfaceMockery) ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	ret := _m.Called(uintArr)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func([]uint) []uint16); ok {
		r0 = rf(uintArr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	return r0
}

// FetchBalance provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) FetchBalance(_a0 *ethclient.Client, _a1 string) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveAssetIds provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetActiveAssetIds(_a0 *ethclient.Client) ([]uint16, error) {
	ret := _m.Called(_a0)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []uint16); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveAssetsData provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetActiveAssetsData(_a0 *ethclient.Client, _a1 uint32) ([]*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) []*big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAmountInDecimal provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetAmountInDecimal(_a0 *big.Int) *big.Float {
	ret := _m.Called(_a0)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Float); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	return r0
}

// GetAmountInWei provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetAmountInWei(_a0 *big.Int) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetBlockManager provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetBlockManager(_a0 *ethclient.Client) *bindings.BlockManager {
	ret := _m.Called(_a0)

	var r0 *bindings.BlockManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.BlockManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.BlockManager)
		}
	}

	return r0
}

// GetCollections provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetCollections(_a0 *ethclient.Client) ([]bindings.StructsCollection, error) {
	ret := _m.Called(_a0)

	var r0 []bindings.StructsCollection
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsCollection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommitments provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetCommitments(_a0 *ethclient.Client, _a1 string) ([32]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) [32]byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFilePath provides a mock function with given fields:
func (_m *UtilsInterfaceMockery) GetConfigFilePath() (string, error) {
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

// GetDefaultPath provides a mock function with given fields:
func (_m *UtilsInterfaceMockery) GetDefaultPath() (string, error) {
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

// GetDelayedState provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetDelayedState(_a0 *ethclient.Client, _a1 int32) (int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, int32) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpoch provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetEpoch(_a0 *ethclient.Client) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastCommitted provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetEpochLastCommitted(_a0 *ethclient.Client, _a1 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastRevealed provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetEpochLastRevealed(_a0 *ethclient.Client, _a1 string, _a2 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFractionalAmountInWei provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetFractionalAmountInWei(_a0 *big.Int, _a1 string) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, string) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluenceSnapshot provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetInfluenceSnapshot(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobs provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetJobs(_a0 *ethclient.Client) ([]bindings.StructsJob, error) {
	ret := _m.Called(_a0)

	var r0 []bindings.StructsJob
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsJob); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLock provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetLock(_a0 *ethclient.Client, _a1 string, _a2 uint32) (types.Locks, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 types.Locks
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32) types.Locks); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(types.Locks)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxAltBlocks provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetMaxAltBlocks(_a0 *ethclient.Client) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumActiveAssets provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetNumActiveAssets(_a0 *ethclient.Client) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfProposedBlocks provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetNumberOfProposedBlocks(_a0 *ethclient.Client, _a1 uint32) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfStakers provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetNumberOfStakers(_a0 *ethclient.Client, _a1 string) (uint32, error) {
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

// GetOptions provides a mock function with given fields:
func (_m *UtilsInterfaceMockery) GetOptions() bind.CallOpts {
	ret := _m.Called()

	var r0 bind.CallOpts
	if rf, ok := ret.Get(0).(func() bind.CallOpts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bind.CallOpts)
	}

	return r0
}

// GetProposedBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetProposedBlock(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandaoHash provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetRandaoHash(_a0 *ethclient.Client) ([32]byte, error) {
	ret := _m.Called(_a0)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client) [32]byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRogueRandomValue provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetRogueRandomValue(_a0 int) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(int) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetSortedProposedBlockIds provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetSortedProposedBlockIds(_a0 *ethclient.Client, _a1 uint32) ([]uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) []uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakedToken provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetStakedToken(_a0 *ethclient.Client, _a1 common.Address) *bindings.StakedToken {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.StakedToken
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) *bindings.StakedToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakedToken)
		}
	}

	return r0
}

// GetStaker provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetStaker(_a0 *ethclient.Client, _a1 string, _a2 uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32) bindings.StructsStaker); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateName provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetStateName(_a0 int64) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStringAddress provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetStringAddress(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalInfluenceRevealed provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetTotalInfluenceRevealed(_a0 *ethclient.Client, _a1 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxnOpts provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetTxnOpts(_a0 types.TransactionOptions) *bind.TransactOpts {
	ret := _m.Called(_a0)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) *bind.TransactOpts); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	return r0
}

// GetUint32BountyId provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetUint32BountyId(_a0 *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdatedEpoch provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetUpdatedEpoch(_a0 *ethclient.Client) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdatedStaker provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetUpdatedStaker(_a0 *ethclient.Client, _a1 string, _a2 uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32) bindings.StructsStaker); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterfaceMockery) GetVoteValue(_a0 *ethclient.Client, _a1 uint16, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVotes provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetVotes(_a0 *ethclient.Client, _a1 uint32) (bindings.StructsVote, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsVote
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsVote); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsVote)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithdrawReleasePeriod provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetWithdrawReleasePeriod(_a0 *ethclient.Client, _a1 string) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsFlagPassed provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) IsFlagPassed(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ParseBool provides a mock function with given fields: str
func (_m *UtilsInterfaceMockery) ParseBool(str string) (bool, error) {
	ret := _m.Called(str)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sleep provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) Sleep(_a0 time.Duration) {
	_m.Called(_a0)
}

// ViperWriteConfigAs provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) ViperWriteConfigAs(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForBlockCompletion provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) WaitForBlockCompletion(_a0 *ethclient.Client, _a1 string) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
