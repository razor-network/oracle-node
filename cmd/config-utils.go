// Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"razor/core"
	"razor/core/types"
	"razor/rpc"
	"razor/utils"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//This function returns the config data
func (*UtilsStruct) GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
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
	if flagSetUtils.Changed(rootCmd.Flags(), flagName) {
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

//This function returns the provider
func (*UtilsStruct) GetProvider() (string, error) {
	provider, err := getConfigValue("provider", "string", "", "provider")
	if err != nil {
		return "", err
	}
	providerString := provider.(string)
	if providerString == "" {
		return "", errors.New("provider is not set")
	}
	if !strings.HasPrefix(providerString, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return providerString, nil
}

//This function returns the multiplier
func (*UtilsStruct) GetMultiplier() (float32, error) {
	const (
		MinMultiplier = 1.0 // Minimum multiplier value
		MaxMultiplier = 3.0 // Maximum multiplier value
	)

	gasMultiplier, err := getConfigValue("gasmultiplier", "float32", core.DefaultGasMultiplier, "gasmultiplier")
	if err != nil {
		return core.DefaultGasMultiplier, err
	}

	multiplierFloat32 := gasMultiplier.(float32)

	// Validate multiplier range
	if multiplierFloat32 < MinMultiplier || multiplierFloat32 > MaxMultiplier {
		log.Infof("GasMultiplier %.2f is out of the valid range (%.1f-%.1f), using default value %.2f", multiplierFloat32, MinMultiplier, MaxMultiplier, core.DefaultGasMultiplier)
		return core.DefaultGasMultiplier, nil
	}

	return multiplierFloat32, nil
}

//This function returns the buffer percent
func (*UtilsStruct) GetBufferPercent() (int32, error) {
	bufferPercent, err := getConfigValue("buffer", "int32", core.DefaultBufferPercent, "buffer")
	if err != nil {
		return core.DefaultBufferPercent, err
	}

	bufferPercentInt32 := bufferPercent.(int32)

	// bufferPercent cannot be less than core.DefaultBufferPercent else through an error
	if bufferPercentInt32 < core.DefaultBufferPercent {
		log.Infof("BufferPercent should be greater than or equal to %v", core.DefaultBufferPercent)
		return core.DefaultBufferPercent, errors.New("invalid buffer percent")
	}

	return bufferPercentInt32, nil
}

//This function returns the wait time
func (*UtilsStruct) GetWaitTime() (int32, error) {
	const (
		MinWaitTime = 1  // Minimum wait time in seconds
		MaxWaitTime = 15 // Maximum wait time in seconds
	)

	waitTime, err := getConfigValue("wait", "int32", core.DefaultWaitTime, "wait")
	if err != nil {
		return core.DefaultWaitTime, err
	}

	waitTimeInt32 := waitTime.(int32)

	// Validate waitTime range
	if waitTimeInt32 < MinWaitTime || waitTimeInt32 > MaxWaitTime {
		log.Infof("WaitTime %d is out of the valid range (%d-%d), using default value %d", waitTimeInt32, MinWaitTime, MaxWaitTime, core.DefaultWaitTime)
		return core.DefaultWaitTime, nil
	}

	return waitTimeInt32, nil
}

//This function returns the gas price
func (*UtilsStruct) GetGasPrice() (int32, error) {
	gasPrice, err := getConfigValue("gasprice", "int32", core.DefaultGasPrice, "gasprice")
	if err != nil {
		return core.DefaultGasPrice, err
	}

	gasPriceInt32 := gasPrice.(int32)

	// Validate gasPrice value
	if gasPriceInt32 != 0 && gasPriceInt32 != 1 {
		log.Infof("GasPrice %d is invalid, using default value %d", gasPriceInt32, core.DefaultGasPrice)
		return core.DefaultGasPrice, nil
	}

	return gasPriceInt32, nil
}

//This function returns the log level
func (*UtilsStruct) GetLogLevel() (string, error) {
	logLevel, err := getConfigValue("logLevel", "string", core.DefaultLogLevel, "logLevel")
	if err != nil {
		return core.DefaultLogLevel, err
	}
	return logLevel.(string), nil
}

//This function returns the gas limit
func (*UtilsStruct) GetGasLimit() (float32, error) {
	//gasLimit in the config acts as a gasLimit multiplier
	const (
		MinGasLimit = 1.0 // Minimum gas limit
		MaxGasLimit = 3.0 // Maximum gas limit
	)

	gasLimit, err := getConfigValue("gasLimit", "float32", core.DefaultGasLimit, "gasLimit")
	if err != nil {
		return core.DefaultGasLimit, err
	}

	gasLimitFloat32 := gasLimit.(float32)

	// Validate gasLimit range
	if gasLimitFloat32 < MinGasLimit || gasLimitFloat32 > MaxGasLimit {
		log.Warnf("GasLimit %.2f is out of the suggested range (%.1f-%.1f), using default value %.2f", gasLimitFloat32, MinGasLimit, MaxGasLimit, core.DefaultGasLimit)
	}

	return gasLimitFloat32, nil
}

//This function returns the gas limit to override
func (*UtilsStruct) GetGasLimitOverride() (uint64, error) {
	const (
		MinGasLimitOverride = 10000000 // Minimum gas limit override
		MaxGasLimitOverride = 50000000 // Maximum gas limit override
	)

	gasLimitOverride, err := getConfigValue("gasLimitOverride", "uint64", core.DefaultGasLimitOverride, "gasLimitOverride")
	if err != nil {
		return core.DefaultGasLimitOverride, err
	}

	gasLimitOverrideUint64 := gasLimitOverride.(uint64)

	// Validate gasLimitOverride range
	if gasLimitOverrideUint64 < MinGasLimitOverride || gasLimitOverrideUint64 > MaxGasLimitOverride {
		log.Infof("GasLimitOverride %d is out of the valid range (%d-%d), using default value %d", gasLimitOverrideUint64, MinGasLimitOverride, MaxGasLimitOverride, core.DefaultGasLimitOverride)
		return core.DefaultGasLimitOverride, nil
	}

	return gasLimitOverrideUint64, nil
}

//This function returns the RPC timeout
func (*UtilsStruct) GetRPCTimeout() (int64, error) {
	const (
		MinRPCTimeout = 1  // Minimum RPC timeout in seconds
		MaxRPCTimeout = 10 // Maximum RPC timeout in seconds
	)

	rpcTimeout, err := getConfigValue("rpcTimeout", "int64", core.DefaultRPCTimeout, "rpcTimeout")
	if err != nil {
		return core.DefaultRPCTimeout, err
	}

	rpcTimeoutInt64 := rpcTimeout.(int64)

	// Validate rpcTimeout range
	if rpcTimeoutInt64 < MinRPCTimeout || rpcTimeoutInt64 > MaxRPCTimeout {
		log.Infof("RPCTimeout %d is out of the valid range (%d-%d), using default value %d", rpcTimeoutInt64, MinRPCTimeout, MaxRPCTimeout, core.DefaultRPCTimeout)
		return core.DefaultRPCTimeout, nil
	}

	return rpcTimeoutInt64, nil
}

func (*UtilsStruct) GetHTTPTimeout() (int64, error) {
	const (
		MinHTTPTimeout = 3 // Minimum HTTP timeout in seconds
		MaxHTTPTimeout = 8 // Maximum HTTP timeout in seconds
	)

	httpTimeout, err := getConfigValue("httpTimeout", "int64", core.DefaultHTTPTimeout, "httpTimeout")
	if err != nil {
		return core.DefaultHTTPTimeout, err
	}

	httpTimeoutInt64 := httpTimeout.(int64)

	// Validate httpTimeout range
	if httpTimeoutInt64 < MinHTTPTimeout || httpTimeoutInt64 > MaxHTTPTimeout {
		log.Infof("HTTPTimeout %d is out of the valid range (%d-%d), using default value %d", httpTimeoutInt64, MinHTTPTimeout, MaxHTTPTimeout, core.DefaultHTTPTimeout)
		return core.DefaultHTTPTimeout, nil
	}

	return httpTimeoutInt64, nil
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

//This function sets the log level
func setLogLevel(config types.Configurations) {
	if config.LogLevel == "debug" {
		log.SetLogLevel(logrus.DebugLevel)
	}

	if razorUtils.IsFlagPassed("logFile") {
		log.Debugf("Log File Max Size: %d MB", config.LogFileMaxSize)
		log.Debugf("Log File Max Backups (max number of old log files to retain): %d", config.LogFileMaxBackups)
		log.Debugf("Log File Max Age (max number of days to retain old log files): %d", config.LogFileMaxAge)
	}
}

func ValidateBufferPercentLimit(rpcParameters rpc.RPCParameters, bufferPercent int32) error {
	stateBuffer, err := razorUtils.GetStateBuffer(rpcParameters)
	if err != nil {
		return err
	}
	maxBufferPercent := calculateMaxBufferPercent(stateBuffer, core.StateLength)
	if bufferPercent >= maxBufferPercent {
		log.Errorf("Buffer percent %v is greater than or equal to maximum possible buffer percent", maxBufferPercent)
		return errors.New("buffer percent exceeds limit")
	}
	return nil
}

// calculateMaxBuffer calculates the maximum buffer percent value.
func calculateMaxBufferPercent(stateBuffer, stateLength uint64) int32 {
	if stateLength == 0 {
		return 0
	}

	// The formula is derived from the condition:
	// 2(maxBuffer % stateLength) < (stateLength - 2*StateBuffer)

	// Perform the calculation with float64 for precision
	maxBufferPercent := 50 * (1 - (float64(2*stateBuffer) / float64(stateLength)))
	return int32(maxBufferPercent)
}
