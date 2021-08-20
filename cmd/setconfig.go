package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"razor/utils"
)

var setConfig = &cobra.Command{
	Use:   "setconfig",
	Short: "setconfig enables user to set the values of provider and gas multiplier",
	Long: `Setting the provider helps the CLI to know which provider to connect to.
Setting the gas multiplier value enables the CLI to multiply the gas with that value for all the transactions

Example:
  ./razor setconfig --provider https://infura/v3/matic --gasmultiplier 1.5 --buffer 20 --wait 70 --gasprice 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		gasMultiplier, _ := cmd.Flags().GetFloat32("gasmultiplier")
		bufferPercent, _ := cmd.Flags().GetInt32("buffer")
		waitTime, _ := cmd.Flags().GetInt32("wait")
		gasPrice, _ := cmd.Flags().GetInt32("gasprice")
		if provider != "" {
			viper.Set("provider", provider)
		}
		if gasMultiplier != -1 {
			viper.Set("gasmultiplier", gasMultiplier)
		}
		if bufferPercent != 0 {
			viper.Set("buffer", bufferPercent)
		}
		if waitTime != -1 {
			viper.Set("wait", waitTime)
		}
		if gasPrice != -1 {
			viper.Set("gasprice", gasPrice)
		}
		if provider == "" && gasMultiplier == -1 && bufferPercent == 0 && waitTime == -1 && gasPrice == -1 {
			viper.Set("provider", "http://127.0.0.1:8545")
			viper.Set("gasmultiplier", 1.0)
			viper.Set("buffer", 30)
			viper.Set("wait", 3)
			viper.Set("gasprice", 0)
			log.Info("Config values set to default. Use setconfig to modify the values.")
		}
		path := utils.GetDefaultPath() + "/razor.yaml"
		err := viper.WriteConfigAs(path)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setConfig)

	var (
		Provider      string
		GasMultiplier float32
		BufferPercent int32
		WaitTime      int32
		GasPrice      int32
	)
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	setConfig.Flags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	setConfig.Flags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time (in secs)")
	setConfig.Flags().Int32VarP(&GasPrice, "gasprice", "", -1, "custom gas price")
}
