package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
	"strings"
)

func (*UtilsStruct) GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
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
	config.Provider = provider
	config.GasMultiplier = gasMultiplier
	config.BufferPercent = bufferPercent
	config.WaitTime = waitTime
	config.GasPrice = gasPrice
	config.LogLevel = logLevel
	config.GasLimitMultiplier = gasLimit

	return config, nil
}

func (*UtilsStruct) GetProvider() (string, error) {
	provider, err := flagSetUtils.GetRootStringProvider()
	if err != nil {
		return "", err
	}
	if provider == "" {
		provider = viper.GetString("provider")
	}
	if !strings.HasPrefix(provider, "https") {
		log.Warn("You are not using a secure RPC URL. Switch to an https URL instead to be safe.")
	}
	return provider, nil
}

func (*UtilsStruct) GetMultiplier() (float32, error) {
	gasMultiplier, err := flagSetUtils.GetRootFloat32GasMultiplier()
	if err != nil {
		return 1, err
	}
	if gasMultiplier == -1 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}
	return gasMultiplier, nil
}

func (*UtilsStruct) GetBufferPercent() (int32, error) {
	bufferPercent, err := flagSetUtils.GetRootInt32Buffer()
	if err != nil {
		return 30, err
	}
	if bufferPercent == 0 {
		bufferPercent = viper.GetInt32("buffer")
	}
	return bufferPercent, nil
}

func (*UtilsStruct) GetWaitTime() (int32, error) {
	waitTime, err := flagSetUtils.GetRootInt32Wait()
	if err != nil {
		return 3, err
	}
	if waitTime == -1 {
		waitTime = viper.GetInt32("wait")
	}
	return waitTime, nil
}

func (*UtilsStruct) GetGasPrice() (int32, error) {
	gasPrice, err := flagSetUtils.GetRootInt32GasPrice()
	if err != nil {
		return 0, err
	}
	if gasPrice == -1 {
		gasPrice = viper.GetInt32("gasprice")
	}
	return gasPrice, nil
}

func (*UtilsStruct) GetLogLevel() (string, error) {
	logLevel, err := flagSetUtils.GetRootStringLogLevel()
	if err != nil {
		return "", err
	}
	if logLevel == "" {
		logLevel = viper.GetString("logLevel")
	}
	return logLevel, nil
}

func (*UtilsStruct) GetGasLimit() (float32, error) {
	gasLimit, err := flagSetUtils.GetRootFloat32GasLimit()
	if err != nil {
		return -1, err
	}
	if gasLimit == -1 {
		gasLimit = float32(viper.GetFloat64("gasLimit"))
	}
	return gasLimit, nil
}
