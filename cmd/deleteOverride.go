package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/utils"
	"strconv"
)

// deleteOverrideCmd represents the deleteOverride command
var deleteOverrideCmd = &cobra.Command{
	Use:   "deleteOverride",
	Short: "delete override job",
	Long:  ``,
	Run:   initialiseDeleteOverrideJob,
}

func initialiseDeleteOverrideJob(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteDeleteOverrideJob(cmd.Flags())
}

func (*UtilsStruct) ExecuteDeleteOverrideJob(flagSet *pflag.FlagSet) {
	jobId, err := flagSetUtils.GetUint16JobId(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	err = cmdUtils.DeleteOverrideJob(jobId)
	utils.CheckError("DeleteOverrideJob error: ", err)
	log.Info("Job removed from override list successfully!")
}

func (*UtilsStruct) DeleteOverrideJob(jobId uint16) error {
	jobPath, err := razorUtils.GetJobFilePath()
	if err != nil {
		return err
	}
	return razorUtils.DeleteJobFromJSON(jobPath, strconv.Itoa(int(jobId)))
}

func init() {
	rootCmd.AddCommand(deleteOverrideCmd)
	var JobId uint16

	deleteOverrideCmd.Flags().Uint16VarP(&JobId, "jobId", "j", 0, "job id to delete")
}
