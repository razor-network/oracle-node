package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/path"
	"razor/utils"
	"strconv"
)

// deleteOverrideCmd represents the deleteOverride command
var deleteOverrideCmd = &cobra.Command{
	Use:   "deleteOverride",
	Short: "delete override job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			assetManagerUtils: assetManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			packageUtils:      packageUtils,
		}
		err := utilsStruct.executeDeleteOverrideJob(cmd.Flags())
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Job removed from override list successfully!")
	},
}

func (utilsStruct UtilsStruct) executeDeleteOverrideJob(flagSet *pflag.FlagSet) error {
	jobId, err := utilsStruct.flagSetUtils.GetUint16JobId(flagSet)
	if err != nil {
		return err
	}
	return utilsStruct.deleteOverrideJob(jobId)
}

func (utilsStruct UtilsStruct) deleteOverrideJob(jobId uint16) error {
	jobPath, err := path.GetJobFilePath()
	if err != nil {
		return err
	}
	return utils.DeleteJobFromJSON(jobPath, strconv.Itoa(int(jobId)))
}

func init() {
	rootCmd.AddCommand(deleteOverrideCmd)
	var JobId uint16

	deleteOverrideCmd.Flags().Uint16VarP(&JobId, "jobId", "j", 0, "job id to delete")
}
