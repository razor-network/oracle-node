package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
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

func initialiseJobList(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteJobList(cmd.Flags())
}

func (*UtilsStruct) ExecuteJobList(flagSet *pflag.FlagSet) {
	cmdUtils.AssignLogFile(flagSet)

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	err = cmdUtils.GetJobList(client)
	utils.CheckError("Error in getting job list: ", err)
}
func (*UtilsStruct) GetJobList(client *ethclient.Client) error {
	jobs, err := razorUtils.GetJobs(client)

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
	var (
		LogFile string
	)

	jobListCmd.Flags().StringVarP(&LogFile, "logFile", "", "", "name of log file")

}
