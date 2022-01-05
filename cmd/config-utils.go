package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
	"strings"
)

func (*UtilsStructMockery) GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
	}

	provider, err := cmdUtilsMockery.GetProvider()
	if err != nil {
		return config, err
	}
	gasMultiplier, err := cmdUtilsMockery.GetMultiplier()
	if err != nil {
		return config, err
	}
	bufferPercent, err := cmdUtilsMockery.GetBufferPercent()
	if err != nil {
		return config, err
	}
	waitTime, err := cmdUtilsMockery.GetWaitTime()
	if err != nil {
		return config, err
	}
	gasPrice, err := cmdUtilsMockery.GetGasPrice()
	if err != nil {
		return config, err
	}
	logLevel, err := cmdUtilsMockery.GetLogLevel()
	if err != nil {
		return config, err
	}
	gasLimit, err := cmdUtilsMockery.GetGasLimit()
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

func (*UtilsStructMockery) GetProvider() (string, error) {
	provider, err := flagSetUtilsMockery.GetRootStringProvider()
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

func (*UtilsStructMockery) GetMultiplier() (float32, error) {
	gasMultiplier, err := flagSetUtilsMockery.GetRootFloat32GasMultiplier()
	if err != nil {
		return 1, err
	}
	if gasMultiplier == -1 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}
	return gasMultiplier, nil
}

func (*UtilsStructMockery) GetBufferPercent() (int32, error) {
	bufferPercent, err := flagSetUtilsMockery.GetRootInt32Buffer()
	if err != nil {
		return 30, err
	}
	if bufferPercent == 0 {
		bufferPercent = viper.GetInt32("buffer")
	}
	return bufferPercent, nil
}

func (*UtilsStructMockery) GetWaitTime() (int32, error) {
	waitTime, err := flagSetUtilsMockery.GetRootInt32Wait()
	if err != nil {
		return 3, err
	}
	if waitTime == -1 {
		waitTime = viper.GetInt32("wait")
	}
	return waitTime, nil
}

func (*UtilsStructMockery) GetGasPrice() (int32, error) {
	gasPrice, err := flagSetUtilsMockery.GetRootInt32GasPrice()
	if err != nil {
		return 0, err
	}
	if gasPrice == -1 {
		gasPrice = viper.GetInt32("gasprice")
	}
	return gasPrice, nil
}

func (*UtilsStructMockery) GetLogLevel() (string, error) {
	logLevel, err := flagSetUtilsMockery.GetRootStringLogLevel()
	if err != nil {
		return "", err
	}
	if logLevel == "" {
		logLevel = viper.GetString("logLevel")
	}
	return logLevel, nil
}

func (*UtilsStructMockery) GetGasLimit() (float32, error) {
	gasLimit, err := flagSetUtilsMockery.GetRootFloat32GasLimit()
	if err != nil {
		return -1, err
	}
	if gasLimit == -1 {
		gasLimit = float32(viper.GetFloat64("gasLimit"))
	}
	return gasLimit, nil
}
