package cmd

import (
	"github.com/spf13/pflag"
	"razor/core/types"
)

//go:generate mockery --name UtilsInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name FlagSetInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name UtilsCmdInterfaceMockery --output ./mocks/ --case=underscore

var razorUtilsMockery UtilsInterfaceMockery
var flagSetUtilsMockery FlagSetInterfaceMockery
var cmdUtilsMockery UtilsCmdInterfaceMockery

type UtilsInterfaceMockery interface {
	GetConfigFilePath() (string, error)
	ViperWriteConfigAs(string) error
}

type FlagSetInterfaceMockery interface {
	GetStringProvider(*pflag.FlagSet) (string, error)
	GetFloat32GasMultiplier(*pflag.FlagSet) (float32, error)
	GetInt32Buffer(*pflag.FlagSet) (int32, error)
	GetInt32Wait(*pflag.FlagSet) (int32, error)
	GetInt32GasPrice(*pflag.FlagSet) (int32, error)
	GetFloat32GasLimit(set *pflag.FlagSet) (float32, error)
	GetStringLogLevel(*pflag.FlagSet) (string, error)
	GetBoolAutoWithdraw(*pflag.FlagSet) (bool, error)
	GetUint32BountyId(*pflag.FlagSet) (uint32, error)
	GetRootStringProvider() (string, error)
	GetRootFloat32GasMultiplier() (float32, error)
	GetRootInt32Buffer() (int32, error)
	GetRootInt32Wait() (int32, error)
	GetRootInt32GasPrice() (int32, error)
	GetRootStringLogLevel() (string, error)
	GetRootFloat32GasLimit() (float32, error)
}

type UtilsCmdInterfaceMockery interface {
	SetConfig(flagSet *pflag.FlagSet) error
	GetProvider() (string, error)
	GetMultiplier() (float32, error)
	GetWaitTime() (int32, error)
	GetGasPrice() (int32, error)
	GetLogLevel() (string, error)
	GetGasLimit() (float32, error)
	GetBufferPercent() (int32, error)
	GetConfigData() (types.Configurations, error)
}

type UtilsMockery struct{}
type FLagSetUtilsMockery struct{}
type UtilsStructMockery struct{}
