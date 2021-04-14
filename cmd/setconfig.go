package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		if provider != "" {
			viper.Set("provider", provider)
		}
		if gasMultiplier != -1 {
			viper.Set("gasmultiplier", gasMultiplier)
		}
		pwd, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path := pwd + "/.razor.yaml"
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setConfig)

	var Provider string
	var GasMultiplier float32
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
}
