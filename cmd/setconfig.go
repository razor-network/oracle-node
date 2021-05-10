package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"razor/utils"
)

var setConfig = &cobra.Command{
	Use:   "setconfig",
	Short: "setconfig enables user to set the values of provider and gas multiplier",
	Long: `Setting the provider helps the CLI to know which provider to connect to.
Setting the gas multiplier value enables the CLI to multiply the gas with that value for all the transactions
`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		gasMultiplier, _ := cmd.Flags().GetFloat32("gasmultiplier")
		chainId, _ := cmd.Flags().GetInt64("chainid")

		if provider != "" {
			viper.Set("provider", provider)
		}
		if gasMultiplier != -1 {
			viper.Set("gasmultiplier", gasMultiplier)
		}
		if chainId != 0000 {
			viper.Set("chainid", chainId)
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
		ChainId       int64
	)
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	setConfig.Flags().Int64VarP(&ChainId, "chainid", "c", 0000, "chain id of the blockchain network")
}
