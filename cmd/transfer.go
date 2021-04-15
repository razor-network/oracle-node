package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, gasMultiplier, err := getConfigData(cmd)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		amount, _ := cmd.Flags().GetFloat32("amount")
		from, _ := cmd.Flags().GetString("from")
		password, _ := cmd.Flags().GetString("password")
		to, _ := cmd.Flags().GetString("to")

		log.Info("provider: ", provider)
		log.Info("gasmultiplier: ", gasMultiplier)
		log.Info("amount: ", amount)
		log.Info("from: ", from)
		log.Info("password: ", password)
		log.Info("to: ", to)

	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	var (
		Amount   float32
		From     string
		Password string
		To       string
	)

	transferCmd.Flags().Float32VarP(&Amount, "amount", "a", 0, "amount to stake")
	transferCmd.Flags().StringVarP(&From, "from", "", "", "transfer from")
	transferCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")
	transferCmd.Flags().StringVarP(&To, "to", "", "", "transfer to")

	transferCmd.MarkFlagRequired("amount")
	transferCmd.MarkFlagRequired("from")
	transferCmd.MarkFlagRequired("password")
	transferCmd.MarkFlagRequired("to")

}
