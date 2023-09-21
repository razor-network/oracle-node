// Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"github.com/sirupsen/logrus"
	"razor/client"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"strings"

	"github.com/spf13/viper"
)

// This function returns the config data
func (*UtilsStruct) GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
		AlternateProvider:  "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
		RPCTimeout:         0,
		HTTPTimeout:        0,
		LogFileMaxSize:     0,
		LogFileMaxBackups:  0,
		LogFileMaxAge:      0,
	}

	provider, err := cmdUtils.GetProvider()
	if err != nil {
		return config, err
	}
	alternateProvider, err := cmdUtils.GetAlternateProvider()
	if err != nil {
		return config, err
	}
	gasMultiplier, err := cmdUtils.GetMultiplier()
	if err != nil {
		return config, err
	}
	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return config, err
	}
	waitTime, err := cmdUtils.GetWaitTime()
	if err != nil {
		return config, err
	}
	gasPrice, err := cmdUtils.GetGasPrice()
	if err != nil {
		return config, err
	}
	logLevel, err := cmdUtils.GetLogLevel()
	if err != nil {
		return config, err
	}
	gasLimit, err := cmdUtils.GetGasLimit()
	if err != nil {
		return config, err
	}
	gasLimitOverride, err := cmdUtils.GetGasLimitOverride()
	if err != nil {
		return config, err
	}
	rpcTimeout, err := cmdUtils.GetRPCTimeout()
	if err != nil {
		return config, err
	}
	httpTimeout, err := cmdUtils.GetHTTPTimeout()
	if err != nil {
		return config, err
	}
	logFileMaxSize, err := cmdUtils.GetLogFileMaxSize()
	if err != nil {
		return config, err
	}
	logFileMaxBackups, err := cmdUtils.GetLogFileMaxBackups()
	if err != nil {
		return config, err
	}
	logFileMaxAge, err := cmdUtils.GetLogFileMaxAge()
	if err != nil {
		return config, err
	}
	config.Provider = provider
	config.AlternateProvider = alternateProvider
	client.SetAlternateProvider(alternateProvider)
	config.GasMultiplier = gasMultiplier
	config.BufferPercent = bufferPercent
	config.WaitTime = waitTime
	config.GasPrice = gasPrice
	config.LogLevel = logLevel
	config.GasLimitMultiplier = gasLimit
	config.GasLimitOverride = gasLimitOverride
	config.RPCTimeout = rpcTimeout
	utils.RPCTimeout = rpcTimeout
	config.HTTPTimeout = httpTimeout
	utils.HTTPTimeout = httpTimeout
	config.LogFileMaxSize = logFileMaxSize
	config.LogFileMaxBackups = logFileMaxBackups
	config.LogFileMaxAge = logFileMaxAge

	setLogLevel(config)

	return config, nil
}

// This function returns the provider
func (*UtilsStruct) GetProvider() (string, error) {
	provider, err := flagSetUtils.GetRootStringProvider()
	if err != nil {
		return "", err
	}
	if provider == "" {
		if viper.IsSet("provider") {
			provider = viper.GetString("provider")
		} else {
			log.Error("Provider is not set in config file")
			return "", errors.New("provider is not set")
		}
	}
	if !strings.HasPrefix(provider, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return provider, nil
}

// This function returns the alternate provider
func (*UtilsStruct) GetAlternateProvider() (string, error) {
	alternateProvider, err := flagSetUtils.GetRootStringAlternateProvider()
	if err != nil {
		return "", err
	}
	if alternateProvider == "" {
		if viper.IsSet("alternateProvider") {
			alternateProvider = viper.GetString("alternateProvider")
		} else {
			alternateProvider = ""
			log.Debug("alternate provider is not set, taking its nil value ", alternateProvider)
		}
	}
	if !strings.HasPrefix(alternateProvider, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return alternateProvider, nil
}

// This function returns the multiplier
func (*UtilsStruct) GetMultiplier() (float32, error) {
	gasMultiplier, err := flagSetUtils.GetRootFloat32GasMultiplier()
	if err != nil {
		return float32(core.DefaultGasMultiplier), err
	}
	if gasMultiplier == -1 {
		if viper.IsSet("gasmultiplier") {
			gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
		} else {
			gasMultiplier = float32(core.DefaultGasMultiplier)
			log.Debug("GasMultiplier is not set, taking its default value ", gasMultiplier)
		}
	}
	return gasMultiplier, nil
}

// This function returns the buffer percent
func (*UtilsStruct) GetBufferPercent() (int32, error) {
	bufferPercent, err := flagSetUtils.GetRootInt32Buffer()
	if err != nil {
		return int32(core.DefaultBufferPercent), err
	}
	if bufferPercent == 0 {
		if viper.IsSet("buffer") {
			bufferPercent = viper.GetInt32("buffer")
		} else {
			bufferPercent = int32(core.DefaultBufferPercent)
			log.Debug("BufferPercent is not set, taking its default value ", bufferPercent)
		}
	}
	return bufferPercent, nil
}

// This function returns the wait time
func (*UtilsStruct) GetWaitTime() (int32, error) {
	waitTime, err := flagSetUtils.GetRootInt32Wait()
	if err != nil {
		return int32(core.DefaultWaitTime), err
	}
	if waitTime == -1 {
		if viper.IsSet("wait") {
			waitTime = viper.GetInt32("wait")
		} else {
			waitTime = int32(core.DefaultWaitTime)
			log.Debug("WaitTime is not set, taking its default value ", waitTime)
		}
	}
	return waitTime, nil
}

// This function returns the gas price
func (*UtilsStruct) GetGasPrice() (int32, error) {
	gasPrice, err := flagSetUtils.GetRootInt32GasPrice()
	if err != nil {
		return int32(core.DefaultGasPrice), err
	}
	if gasPrice == -1 {
		if viper.IsSet("gasprice") {
			gasPrice = viper.GetInt32("gasprice")
		} else {
			gasPrice = int32(core.DefaultGasPrice)
			log.Debug("GasPrice is not set, taking its default value ", gasPrice)

		}
	}
	return gasPrice, nil
}

// This function returns the log level
func (*UtilsStruct) GetLogLevel() (string, error) {
	logLevel, err := flagSetUtils.GetRootStringLogLevel()
	if err != nil {
		return core.DefaultLogLevel, err
	}
	if logLevel == "" {
		if viper.IsSet("logLevel") {
			logLevel = viper.GetString("logLevel")
		} else {
			logLevel = core.DefaultLogLevel
			log.Debug("LogLevel is not set, taking its default value ", logLevel)
		}
	}
	return logLevel, nil
}

// This function returns the gas limit
func (*UtilsStruct) GetGasLimit() (float32, error) {
	gasLimit, err := flagSetUtils.GetRootFloat32GasLimit()
	if err != nil {
		return float32(core.DefaultGasLimit), err
	}
	if gasLimit == -1 {
		if viper.IsSet("gasLimit") {
			gasLimit = float32(viper.GetFloat64("gasLimit"))
		} else {
			gasLimit = float32(core.DefaultGasLimit)
			log.Debug("GasLimit is not set, taking its default value ", gasLimit)
		}
	}
	return gasLimit, nil
}

// This function returns the gas limit to override
func (*UtilsStruct) GetGasLimitOverride() (uint64, error) {
	gasLimitOverride, err := flagSetUtils.GetRootUint64GasLimitOverride()
	if err != nil {
		return uint64(core.DefaultGasLimitOverride), err
	}
	if gasLimitOverride == 0 {
		if viper.IsSet("gasLimitOverride") {
			gasLimitOverride = viper.GetUint64("gasLimitOverride")
		} else {
			gasLimitOverride = uint64(core.DefaultGasLimitOverride)
			log.Debug("GasLimitOverride is not set, taking its default value ", gasLimitOverride)
		}
	}
	return gasLimitOverride, nil
}

// This function returns the RPC timeout
func (*UtilsStruct) GetRPCTimeout() (int64, error) {
	rpcTimeout, err := flagSetUtils.GetRootInt64RPCTimeout()
	if err != nil {
		return int64(core.DefaultRPCTimeout), err
	}
	if rpcTimeout == 0 {
		if viper.IsSet("rpcTimeout") {
			rpcTimeout = viper.GetInt64("rpcTimeout")
		} else {
			rpcTimeout = int64(core.DefaultRPCTimeout)
			log.Debug("RPCTimeout is not set, taking its default value ", rpcTimeout)
		}
	}
	return rpcTimeout, nil
}

func (*UtilsStruct) GetHTTPTimeout() (int64, error) {
	httpTimeout, err := flagSetUtils.GetRootInt64HTTPTimeout()
	if err != nil {
		return int64(core.DefaultHTTPTimeout), err
	}
	if httpTimeout == 0 {
		if viper.IsSet("httpTimeout") {
			httpTimeout = viper.GetInt64("httpTimeout")
		} else {
			httpTimeout = int64(core.DefaultRPCTimeout)
			log.Debug("HTTPTimeout is not set, taking its default value ", httpTimeout)
		}
	}
	return httpTimeout, nil
}

func (*UtilsStruct) GetLogFileMaxSize() (int, error) {
	logFileMaxSize, err := flagSetUtils.GetRootIntLogFileMaxSize()
	if err != nil {
		return core.DefaultLogFileMaxSize, err
	}
	if logFileMaxSize == 0 {
		if viper.IsSet("logFileMaxSize") {
			logFileMaxSize = viper.GetInt("logFileMaxSize")
		} else {
			logFileMaxSize = core.DefaultLogFileMaxSize
			log.Debug("logFileMaxSize is not set, taking its default value ", logFileMaxSize)
		}
	}
	return logFileMaxSize, nil
}

func (*UtilsStruct) GetLogFileMaxBackups() (int, error) {
	logFileMaxBackups, err := flagSetUtils.GetRootIntLogFileMaxBackups()
	if err != nil {
		return core.DefaultLogFileMaxBackups, err
	}
	if logFileMaxBackups == 0 {
		if viper.IsSet("logFileMaxBackups") {
			logFileMaxBackups = viper.GetInt("logFileMaxBackups")
		} else {
			logFileMaxBackups = core.DefaultLogFileMaxBackups
			log.Debug("logFileMaxBackups is not set, taking its default value ", logFileMaxBackups)
		}
	}
	return logFileMaxBackups, nil
}

func (*UtilsStruct) GetLogFileMaxAge() (int, error) {
	logFileMaxAge, err := flagSetUtils.GetRootIntLogFileMaxAge()
	if err != nil {
		return core.DefaultLogFileMaxAge, err
	}
	if logFileMaxAge == 0 {
		if viper.IsSet("logFileMaxAge") {
			logFileMaxAge = viper.GetInt("logFileMaxAge")
		} else {
			logFileMaxAge = core.DefaultLogFileMaxAge
			log.Debug("logFileMaxAge is not set, taking its default value ", logFileMaxAge)
		}
	}
	return logFileMaxAge, nil
}

// This function sets the log level
func setLogLevel(config types.Configurations) {
	if config.LogLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Debug("Config details: ")
	log.Debugf("Provider: %s", config.Provider)
	log.Debugf("Alternate Provider: %s", config.AlternateProvider)
	log.Debugf("Gas Multiplier: %.2f", config.GasMultiplier)
	log.Debugf("Buffer Percent: %d", config.BufferPercent)
	log.Debugf("Wait Time: %d", config.WaitTime)
	log.Debugf("Gas Price: %d", config.GasPrice)
	log.Debugf("Log Level: %s", config.LogLevel)
	log.Debugf("Gas Limit: %.2f", config.GasLimitMultiplier)
	log.Debugf("Gas Limit Override: %d", config.GasLimitOverride)
	log.Debugf("RPC Timeout: %d", config.RPCTimeout)
	log.Debugf("HTTP Timeout: %d", config.HTTPTimeout)

	if razorUtils.IsFlagPassed("logFile") {
		log.Debugf("Log File Max Size: %d MB", config.LogFileMaxSize)
		log.Debugf("Log File Max Backups (max number of old log files to retain): %d", config.LogFileMaxBackups)
		log.Debugf("Log File Max Age (max number of days to retain old log files): %d", config.LogFileMaxAge)
	}
}
