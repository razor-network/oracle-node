package cmd

import (
	"github.com/spf13/viper"
	"razor/core/types"
)

func GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:      "",
		GasMultiplier: 0,
	}
	provider, err := GetProvider()
	if err != nil {
		return config, err
	}
	gasMultiplier, err := GetMultiplier()
	if err != nil {
		return config, err
	}
	config.Provider = provider
	config.GasMultiplier = gasMultiplier
	return config, nil
}

func GetProvider() (string, error) {
	provider, err := rootCmd.PersistentFlags().GetString("provider")
	if err != nil {
		return "", err
	}
	if provider == "" {
		provider = viper.GetString("provider")
	}
	return provider, nil
}

func GetMultiplier() (float32, error) {
	gasMultiplier, err := rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
	if err != nil {
		return 0, err
	}
	if gasMultiplier == 0 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}
	return gasMultiplier, nil
}