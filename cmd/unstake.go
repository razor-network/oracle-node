package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//provider, gasMultiplier, err := utils.getConfigData(cmd)
		//if err != nil {
		//	log.Error(err)
		//	os.Exit(1)
		//}
		//address, _ := cmd.Flags().GetString("address")
		//password, _ := cmd.Flags().GetString("password")
		//
		//log.Info("provider: ", provider)
		//log.Info("gasmultiplier: ", gasMultiplier)
		//log.Info("address: ", address)
		//log.Info("password: ", password)
		log.Info("Unstake called")
	},
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address  string
		Password string
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	unstakeCmd.MarkFlagRequired("address")
	unstakeCmd.MarkFlagRequired("password")
}
