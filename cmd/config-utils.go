//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
	"strings"
)

//This function returns the config data
func (*UtilsStruct) GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
		ChainId:            0,
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
		LogFileMaxSize:     0,
		LogFileMaxBackups:  0,
		LogFileMaxAge:      0,
	}
	provider, err := cmdUtils.GetConfig("provider")
	if err != nil {
		return config, err
	}
	chainIdString, err := cmdUtils.GetConfig("chainId")
	if err != nil {
		return config, err
	}
	gasMultiplierString, err := cmdUtils.GetConfig("gasmultiplier")
	if err != nil {
		return config, err
	}
	bufferPercentString, err := cmdUtils.GetConfig("buffer")
	if err != nil {
		return config, err
	}
	waitTimeString, err := cmdUtils.GetConfig("wait")
	if err != nil {
		return config, err
	}
	gasPriceString, err := cmdUtils.GetConfig("gasprice")
	if err != nil {
		return config, err
	}
	logLevel, err := cmdUtils.GetConfig("logLevel")
	if err != nil {
		return config, err
	}
	gasLimitString, err := cmdUtils.GetConfig("gasLimit")
	if err != nil {
		return config, err
	}
	logFileMaxSizeString, err := cmdUtils.GetConfig("logFileMaxSize")
	if err != nil {
		return config, err
	}
	logFileMaxBackupsString, err := cmdUtils.GetConfig("logFileMaxBackups")
	if err != nil {
		return config, err
	}
	logFileMaxAgeString, err := cmdUtils.GetConfig("logFileMaxAge")
	if err != nil {
		return config, err
	}
	config.Provider = provider
	chainId, err := stringUtils.ParseChainId(chainIdString)
	if err != nil {
		return config, err
	}
	config.ChainId = chainId
	gasMultiplier, err := stringUtils.ParseFloat(gasMultiplierString)
	if err != nil {
		return config, err
	}
	config.GasMultiplier = float32(gasMultiplier)
	bufferPercent, err := stringUtils.ParseInt64(bufferPercentString)
	if err != nil {
		return config, err
	}
	config.BufferPercent = int32(bufferPercent)
	waitTime, err := stringUtils.ParseInt64(waitTimeString)
	if err != nil {
		return config, err
	}
	config.WaitTime = int32(waitTime)
	gasPrice, err := stringUtils.ParseInt64(gasPriceString)
	if err != nil {
		return config, err
	}
	config.GasPrice = int32(gasPrice)
	config.LogLevel = logLevel
	gasLimit, err := stringUtils.ParseFloat(gasLimitString)
	if err != nil {
		return config, err
	}
	config.GasLimitMultiplier = float32(gasLimit)
	logFileMaxSize, err := stringUtils.ParseInt(logFileMaxSizeString)
	if err != nil {
		return config, err
	}
	config.LogFileMaxSize = logFileMaxSize
	logFileMaxBackups, err := stringUtils.ParseInt(logFileMaxBackupsString)
	if err != nil {
		return config, err
	}
	config.LogFileMaxBackups = logFileMaxBackups
	logFileMaxAge, err := stringUtils.ParseInt(logFileMaxAgeString)
	if err != nil {
		return config, err
	}
	config.LogFileMaxAge = logFileMaxAge

	return config, nil
}

//This function returns the config value in form of string taking configType as input
func (*UtilsStruct) GetConfig(configType string) (string, error) {
	switch configType {
	case "provider":
		provider, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "", err
		}
		if provider == "" {
			provider = viper.GetString(configType)
		}
		if !strings.HasPrefix(provider, "https") {
			log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
		}
		return provider, nil

	case "chainId":
		chainId, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "0", err
		}
		if chainId == "0" {
			chainId = viper.GetString(configType)
		}
		return chainId, nil

	case "gasmultiplier":
		gasMultiplier, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "1", err
		}
		if gasMultiplier == "-1" {
			gasMultiplier = viper.GetString(configType)
		}
		return gasMultiplier, nil

	case "buffer":
		bufferPercent, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "30", err
		}
		if bufferPercent == "0" {
			bufferPercent = viper.GetString(configType)
		}
		return bufferPercent, nil

	case "wait":
		waitTime, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "3", err
		}
		if waitTime == "-1" {
			waitTime = viper.GetString(configType)
		}
		return waitTime, nil

	case "gasprice":
		gasPrice, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "0", err
		}
		if gasPrice == "-1" {
			gasPrice = viper.GetString(configType)
		}
		return gasPrice, nil

	case "logLevel":
		logLevel, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "", err
		}
		if logLevel == "" {
			logLevel = viper.GetString(configType)
		}
		return logLevel, nil

	case "gasLimit":
		gasLimit, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "-1", err
		}
		if gasLimit == "-1" {
			gasLimit = viper.GetString(configType)
		}
		return gasLimit, nil

	case "logFileMaxSize":
		logFileMaxSize, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "5", err
		}
		if logFileMaxSize == "0" {
			logFileMaxSize = viper.GetString(configType)
		}
		return logFileMaxSize, nil

	case "logFileMaxBackups":
		logFileMaxBackups, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "10", err
		}
		if logFileMaxBackups == "0" {
			logFileMaxBackups = viper.GetString(configType)
		}
		return logFileMaxBackups, nil

	case "logFileMaxAge":
		logFileMaxAge, err := flagSetUtils.GetRootStringConfig(configType)
		if err != nil {
			return "30", err
		}
		if logFileMaxAge == "0" {
			logFileMaxAge = viper.GetString(configType)
		}
		return logFileMaxAge, nil
	}
	return "", nil
}
