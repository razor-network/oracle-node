package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getConfigData(cmd *cobra.Command) (string, float32, error) {
	provider, err := rootCmd.PersistentFlags().GetString("provider")
	if err != nil {
		return "", 0, err
	}

	gasMultiplier, err := rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
	if err != nil {
		return "", 0, err
	}

	if provider == "" {
		provider = viper.GetString("provider")
	}

	if gasMultiplier == 0 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}

	if provider == "" || gasMultiplier == 0 {
		return "", 0, errors.New("provider and gas multiplier value not set")
	}

	return provider, gasMultiplier, nil
}
