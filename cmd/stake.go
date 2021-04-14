package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stakeCmd represents the stake command
var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some schells",
	Long: `Stake allows user to stake schells in the razor network
	For ex:
	stake -a <amount> --address <address> --password <password>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, _ := cmd.Flags().GetString("provider")
		gasMultiplier, _ := cmd.Flags().GetFloat32("gasmultiplier")
		amount, _ := cmd.Flags().GetFloat32("amount")
		address, _ := cmd.Flags().GetString("address")
		password, _ := cmd.Flags().GetString("password")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Warn("no config file found and no data provided either")
			} else {
				log.Warn("error in reading config file")
			}
		}
		if provider == "" {
			provider = viper.GetString("provider")
		}
		if gasMultiplier == -1 {
			gasMultiplier = float32(viper.GetFloat64("gasmultiplier"))
		}

		log.Info("provider: ", provider)
		log.Info("gasmultiplier: ", gasMultiplier)
		log.Info("amount: ", amount)
		log.Info("address: ", address)
		log.Info("password: ", password)

	},
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	var (
		Provider      string
		GasMultiplier float32
		Amount        float32
		Address       string
		Password      string
	)

	stakeCmd.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	stakeCmd.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	stakeCmd.Flags().Float32VarP(&Amount, "amount", "a", 0, "amount to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	stakeCmd.MarkFlagRequired("amount")
	stakeCmd.MarkFlagRequired("address")
	stakeCmd.MarkFlagRequired("password")

}
