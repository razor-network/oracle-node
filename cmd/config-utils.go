package cmd

import (
	"github.com/spf13/viper"
	"math/big"
	"razor/core/types"
)

func GetConfigData() (types.Configurations, error) {
	config := types.Configurations{
		Provider:      "",
		GasMultiplier: 0,
		ChainId:       nil,
	}
	provider, err := GetProvider()
	if err != nil {
		return config, err
	}
	gasMultiplier, err := GetMultiplier()
	if err != nil {
		return config, err
	}
	chainId, err := GetChainId()
	if err != nil {
		return config, err
	}
	config.Provider = provider
	config.GasMultiplier = gasMultiplier
	config.ChainId = chainId
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

func GetChainId() (*big.Int, error) {
	chainId, err := rootCmd.PersistentFlags().GetInt64("chainid")
	if err != nil {
		return nil, err
	}
	if chainId == 0000 {
		chainId = viper.GetInt64("chainid")
	}
	return big.NewInt(chainId), nil
}
