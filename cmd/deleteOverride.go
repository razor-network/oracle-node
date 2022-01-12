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
	Run:   initialiseDeleteOverrideJob,
}

func initialiseDeleteOverrideJob(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteDeleteOverrideJob(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteDeleteOverrideJob(flagSet *pflag.FlagSet) {
	jobId, err := flagSetUtilsMockery.GetUint16JobId(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	err = cmdUtilsMockery.DeleteOverrideJob(jobId)
	utils.CheckError("DeleteOverrideJob error: ", err)
	log.Info("Job removed from override list successfully!")
}

func (*UtilsStructMockery) DeleteOverrideJob(jobId uint16) error {
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
