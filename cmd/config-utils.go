package cmd

import (
	"errors"
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
		return "", errors.New("provider value is not set")
	}
	return provider, nil
}

func GetMultiplier() (float32, error) {
	gasMultiplier, err := rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
	if err != nil {
		return 0, err
	}
	if gasMultiplier == 0 {
		return 0, errors.New("gas multiplier value not set")
	}
	return gasMultiplier, nil
}

func GetChainId() (*big.Int, error) {
	chainId, err := rootCmd.PersistentFlags().GetInt64("chainid")
	if err != nil {
		return nil, err
	}
	if chainId == 0000 {
		return nil, errors.New("chain id not set")
	}
	return big.NewInt(chainId), nil
}