package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getConfigData(cmd *cobra.Command) (string, float32, error) {
	provider, err := rootCmd.PersistentFlags().GetString("provider")
	if err != nil {
		return "", -1, err
	}
	gasMultiplier, err := rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
	if err != nil {
		return "", -1, err
	}

	if provider == "" {
		provider = viper.GetString("provider")
	}
	if gasMultiplier == -1 {
		gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
	}

	return provider, gasMultiplier, nil
}
