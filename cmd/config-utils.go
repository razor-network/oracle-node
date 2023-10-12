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

func getConfigValueForKey(key string, dataType string) interface{} {
	switch dataType {
	case "string":
		return viper.GetString(key)
	case "float32": // Note: viper doesn't have GetFloat32
		return float32(viper.GetFloat64(key))
	case "float64":
		return viper.GetFloat64(key)
	case "int":
		return viper.GetInt(key)
	case "int32":
		return viper.GetInt32(key)
	case "int64":
		return viper.GetInt64(key)
	case "uint64":
		return viper.GetUint64(key)
	default:
		log.Fatalf("Unsupported data type: %s", dataType)
		return nil
	}
}

func getConfigValue(flagName string, dataType string, defaultReturnValue interface{}, viperKey string) (interface{}, error) {
	// Check if the config parameter was passed as a root flag in the command.
	if rootCmd.Flags().Changed(flagName) {
		// Getting the root flag input
		rootFlagValue, err := flagSetUtils.FetchRootFlagInput(flagName, dataType)
		if err != nil {
			log.Errorf("Error in getting value from root flag")
			return defaultReturnValue, err
		}
		log.Debugf("%v flag passed as root flag, Taking value of config %v = %v ", flagName, flagName, rootFlagValue)
		return rootFlagValue, nil
	}

	// Checking if value of config parameter is present in config file
	if viper.IsSet(viperKey) {
		valueForKey := getConfigValueForKey(viperKey, dataType)
		log.Debugf("Taking value of config %v = %v from config file", viperKey, valueForKey)
		return valueForKey, nil
	}
	log.Debugf("%v config is not set, taking its default value %v", viperKey, defaultReturnValue)
	return defaultReturnValue, nil
}

// This function returns the provider
func (*UtilsStruct) GetProvider() (string, error) {
	provider, err := getConfigValue("provider", "string", "", "provider")
	if err != nil {
		return "", err
	}
	providerString := provider.(string)
	if providerString == "" {
		return "", errors.New("provider not set")
	}
	if !strings.HasPrefix(providerString, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return providerString, nil
}

// This function returns the alternate provider
func (*UtilsStruct) GetAlternateProvider() (string, error) {
	alternateProvider, err := getConfigValue("alternateProvider", "string", "", "alternateProvider")
	if err != nil {
		return "", err
	}
	alternateProviderString := alternateProvider.(string)
	if !strings.HasPrefix(alternateProviderString, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return alternateProviderString, nil
}

// This function returns the multiplier
func (*UtilsStruct) GetMultiplier() (float32, error) {
	gasMultiplier, err := getConfigValue("gasmultiplier", "float32", core.DefaultGasMultiplier, "gasmultiplier")
	if err != nil {
		return float32(core.DefaultGasMultiplier), err
	}
	return gasMultiplier.(float32), nil
}

// This function returns the buffer percent
func (*UtilsStruct) GetBufferPercent() (int32, error) {
	bufferPercent, err := getConfigValue("buffer", "int32", core.DefaultBufferPercent, "buffer")
	if err != nil {
		return int32(core.DefaultBufferPercent), err
	}
	return bufferPercent.(int32), nil
}

// This function returns the wait time
func (*UtilsStruct) GetWaitTime() (int32, error) {
	waitTime, err := getConfigValue("wait", "int32", core.DefaultWaitTime, "wait")
	if err != nil {
		return int32(core.DefaultWaitTime), err
	}
	return waitTime.(int32), nil
}

// This function returns the gas price
func (*UtilsStruct) GetGasPrice() (int32, error) {
	gasPrice, err := getConfigValue("gasprice", "int32", core.DefaultGasPrice, "gasprice")
	if err != nil {
		return int32(core.DefaultGasPrice), err
	}
	return gasPrice.(int32), nil
}

// This function returns the log level
func (*UtilsStruct) GetLogLevel() (string, error) {
	logLevel, err := getConfigValue("logLevel", "string", core.DefaultLogLevel, "logLevel")
	if err != nil {
		return core.DefaultLogLevel, err
	}
	return logLevel.(string), nil
}

// This function returns the gas limit
func (*UtilsStruct) GetGasLimit() (float32, error) {
	gasLimit, err := getConfigValue("gasLimit", "float32", core.DefaultGasLimit, "gasLimit")
	if err != nil {
		return float32(core.DefaultGasLimit), err
	}
	return gasLimit.(float32), nil
}

// This function returns the gas limit to override
func (*UtilsStruct) GetGasLimitOverride() (uint64, error) {
	gasLimitOverride, err := getConfigValue("gasLimitOverride", "uint64", core.DefaultGasLimitOverride, "gasLimitOverride")
	if err != nil {
		return uint64(core.DefaultGasLimitOverride), err
	}
	return gasLimitOverride.(uint64), nil
}

// This function returns the RPC timeout
func (*UtilsStruct) GetRPCTimeout() (int64, error) {
	rpcTimeout, err := getConfigValue("rpcTimeout", "int64", core.DefaultRPCTimeout, "rpcTimeout")
	if err != nil {
		return int64(core.DefaultRPCTimeout), err
	}
	return rpcTimeout.(int64), nil
}

func (*UtilsStruct) GetHTTPTimeout() (int64, error) {
	httpTimeout, err := getConfigValue("httpTimeout", "int64", core.DefaultHTTPTimeout, "httpTimeout")
	if err != nil {
		return int64(core.DefaultHTTPTimeout), err
	}
	return httpTimeout.(int64), nil
}

func (*UtilsStruct) GetLogFileMaxSize() (int, error) {
	logFileMaxSize, err := getConfigValue("logFileMaxSize", "int", core.DefaultLogFileMaxSize, "logFileMaxSize")
	if err != nil {
		return core.DefaultLogFileMaxSize, err
	}
	return logFileMaxSize.(int), nil
}

func (*UtilsStruct) GetLogFileMaxBackups() (int, error) {
	logFileMaxBackups, err := getConfigValue("logFileMaxBackups", "int", core.DefaultLogFileMaxBackups, "logFileMaxBackups")
	if err != nil {
		return core.DefaultLogFileMaxBackups, err
	}
	return logFileMaxBackups.(int), nil
}

func (*UtilsStruct) GetLogFileMaxAge() (int, error) {
	logFileMaxAge, err := getConfigValue("logFileMaxAge", "int", core.DefaultLogFileMaxAge, "logFileMaxAge")
	if err != nil {
		return core.DefaultLogFileMaxAge, err
	}
	return logFileMaxAge.(int), nil
}

// This function sets the log level
func setLogLevel(config types.Configurations) {
	if config.LogLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Debugf("Config details: %v", config)

	if razorUtils.IsFlagPassed("logFile") {
		log.Debugf("Log File Max Size: %d MB", config.LogFileMaxSize)
		log.Debugf("Log File Max Backups (max number of old log files to retain): %d", config.LogFileMaxBackups)
		log.Debugf("Log File Max Age (max number of days to retain old log files): %d", config.LogFileMaxAge)
	}
}
