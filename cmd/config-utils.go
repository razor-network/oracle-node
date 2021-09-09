package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
	"strings"
)

func GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:      "",
		GasMultiplier: 0,
		BufferPercent: 0,
		WaitTime:      0,
	}
	provider, err := getProvider()
	if err != nil {
		return config, err
	}
	gasMultiplier, err := getMultiplier()
	if err != nil {
		return config, err
	}
	bufferPercent, err := getBufferPercent()
	if err != nil {
		return config, err
	}
	waitTime, err := getWaitTime()
	if err != nil {
		return config, err
	}
	gasPrice, err := getGasPrice()
	if err != nil {
		return config, err
	}
	config.Provider = provider
	config.GasMultiplier = gasMultiplier
	config.BufferPercent = bufferPercent
	config.WaitTime = waitTime
	config.GasPrice = gasPrice
	return config, nil
}

func getProvider() (string, error) {
	provider, err := rootCmd.PersistentFlags().GetString("provider")
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

func getMultiplier() (float32, error) {
	gasMultiplier, err := rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
	if err != nil {
		return 1, err
	}
	if gasMultiplier == -1 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}
	return gasMultiplier, nil
}

func getBufferPercent() (int32, error) {
	bufferPercent, err := rootCmd.PersistentFlags().GetInt32("buffer")
	if err != nil {
		return 30, err
	}
	if bufferPercent == 0 {
		bufferPercent = viper.GetInt32("buffer")
	}
	return bufferPercent, nil
}

func getWaitTime() (int32, error) {
	waitTime, err := rootCmd.PersistentFlags().GetInt32("wait")
	if err != nil {
		return 3, err
	}
	if waitTime == -1 {
		waitTime = viper.GetInt32("wait")
	}
	return waitTime, nil
}

func getGasPrice() (int32, error) {
	gasPrice, err := rootCmd.PersistentFlags().GetInt32("gasprice")
	if err != nil {
		return 0, err
	}
	if gasPrice == -1 {
		gasPrice = viper.GetInt32("gasprice")
	}
	return gasPrice, nil
}
