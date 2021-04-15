package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some schells",
	Long: `Stake allows user to stake schells in the razor network
	For ex:
	stake -a <amount> --address <address> --password <password>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, gasMultiplier, err := getConfigData(cmd)
		if err != nil {
			log.Error(err)
		}
		amount, _ := cmd.Flags().GetFloat32("amount")
		address, _ := cmd.Flags().GetString("address")
		password, _ := cmd.Flags().GetString("password")

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
		Amount   float32
		Address  string
		Password string
	)

	stakeCmd.Flags().Float32VarP(&Amount, "amount", "a", 0, "amount to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	stakeCmd.MarkFlagRequired("amount")
	stakeCmd.MarkFlagRequired("address")
	stakeCmd.MarkFlagRequired("password")

}
