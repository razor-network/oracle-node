package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
	"strings"
)

func GetConfigData(utilsStruct UtilsStruct) (types.Configurations, error) {
	config := types.Configurations{
		Provider:           "",
		GasMultiplier:      0,
		BufferPercent:      0,
		WaitTime:           0,
		LogLevel:           "",
		GasLimitMultiplier: 0,
	}

	provider, err := utilsStruct.razorUtils.getProvider(utilsStruct)
	if err != nil {
		return config, err
	}
	gasMultiplier, err := utilsStruct.razorUtils.getMultiplier(utilsStruct)
	if err != nil {
		return config, err
	}
	bufferPercent, err := utilsStruct.razorUtils.getBufferPercent(utilsStruct)
	if err != nil {
		return config, err
	}
	waitTime, err := utilsStruct.razorUtils.getWaitTime(utilsStruct)
	if err != nil {
		return config, err
	}
	gasPrice, err := utilsStruct.razorUtils.getGasPrice(utilsStruct)
	if err != nil {
		return config, err
	}
	logLevel, err := utilsStruct.razorUtils.getLogLevel(utilsStruct)
	if err != nil {
		return config, err
	}
	gasLimit, err := utilsStruct.razorUtils.getGasLimit(utilsStruct)
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

func getProvider(utilsStruct UtilsStruct) (string, error) {
	provider, err := utilsStruct.flagSetUtils.GetRootStringProvider()
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

func getMultiplier(utilsStruct UtilsStruct) (float32, error) {
	gasMultiplier, err := utilsStruct.flagSetUtils.GetRootFloat32GasMultiplier()
	if err != nil {
		return 1, err
	}
	if gasMultiplier == -1 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}
	return gasMultiplier, nil
}

func getBufferPercent(utilsStruct UtilsStruct) (int32, error) {
	bufferPercent, err := utilsStruct.flagSetUtils.GetRootInt32Buffer()
	if err != nil {
		return 30, err
	}
	if bufferPercent == 0 {
		bufferPercent = viper.GetInt32("buffer")
	}
	return bufferPercent, nil
}

func getWaitTime(utilsStruct UtilsStruct) (int32, error) {
	waitTime, err := utilsStruct.flagSetUtils.GetRootInt32Wait()
	if err != nil {
		return 3, err
	}
	if waitTime == -1 {
		waitTime = viper.GetInt32("wait")
	}
	return waitTime, nil
}

func getGasPrice(utilsStruct UtilsStruct) (int32, error) {
	gasPrice, err := utilsStruct.flagSetUtils.GetRootInt32GasPrice()
	if err != nil {
		return 0, err
	}
	if gasPrice == -1 {
		gasPrice = viper.GetInt32("gasprice")
	}
	return gasPrice, nil
}

func getLogLevel(utilsStruct UtilsStruct) (string, error) {
	logLevel, err := utilsStruct.flagSetUtils.getRootStringLogLevel()
	if err != nil {
		return "", err
	}
	if logLevel == "" {
		logLevel = viper.GetString("logLevel")
	}
	return logLevel, nil
}

func getGasLimit(utilsStruct UtilsStruct) (float32, error) {
	gasLimit, err := utilsStruct.flagSetUtils.GetRootFloat32GasLimit()
	if err != nil {
		return -1, err
	}
	if gasLimit == -1 {
		gasLimit = float32(viper.GetFloat64("gasLimit"))
	}
	return gasLimit, nil
}
