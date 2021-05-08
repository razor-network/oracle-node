package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
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
		log.Info("Withdraw called")
	},
}

func init() {
	rootCmd.AddCommand(withdrawCmd)

	var (
		Address  string
		Password string
	)

	withdrawCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	withdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	withdrawCmd.MarkFlagRequired("address")
	withdrawCmd.MarkFlagRequired("password")
}
