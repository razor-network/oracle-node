package cmd

import (
	"github.com/spf13/cobra"
)

// deleteOverrideCmd represents the deleteOverride command
var deleteOverrideCmd = &cobra.Command{
	Use:   "deleteOverride",
	Short: "delete override job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(deleteOverrideCmd)

	var JobId uint8

	deleteOverrideCmd.Flags().Uint8VarP(&JobId, "jobId", "j", 0, "job id to delete")
}
