//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"razor/RPC"
	"razor/core/types"
	"razor/logger"
	"razor/utils"
	"strconv"
)

// jobListCmd represents the jobList command
var jobListCmd = &cobra.Command{
	Use:   "jobList",
	Short: "list of all jobs",
	Long: `Provides the list of all jobs with their name, weight, power etc.

Example:
	./razor jobList`,
	Run: initialiseJobList,
}

//This function initialises the ExecuteJobList function
func initialiseJobList(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteJobList(cmd.Flags())
}

//This function sets the flags appropriately and executes the GetJobList function
func (*UtilsStruct) ExecuteJobList(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteJobList: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)
	logger.SetLoggerParameters(client, "")

	rpcManager, err := RPC.InitializeRPCManager(config.Provider)
	utils.CheckError("Error in initializing RPC Manager: ", err)

	rpcParameters := types.RPCParameters{
		RPCManager: rpcManager,
		Ctx:        context.Background(),
	}

	log.Debug("ExecuteJobList: Calling JobList()...")
	err = cmdUtils.GetJobList(rpcParameters)
	utils.CheckError("Error in getting job list: ", err)
}

//This function provides the list of all jobs
func (*UtilsStruct) GetJobList(rpcParameters types.RPCParameters) error {
	jobs, err := razorUtils.GetJobs(rpcParameters)
	log.Debugf("JobList: Jobs: %+v", jobs)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Job Id", "Selector Type", "Weight", "Power", "Name", "Selector", "Url"})
	for i := 0; i < len(jobs); i++ {
		table.Append([]string{
			strconv.Itoa(int(jobs[i].Id)),
			strconv.Itoa(int(jobs[i].SelectorType)),
			strconv.Itoa(int(jobs[i].Weight)),
			strconv.Itoa(int(jobs[i].Power)),
			jobs[i].Name,
			jobs[i].Selector,
			jobs[i].Url,
		})

	}

	table.Render()
	return nil

}

func init() {
	rootCmd.AddCommand(jobListCmd)

}
