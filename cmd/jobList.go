package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {

		utilsStruct := UtilsStruct{
			razorUtils:   razorUtils,
			flagSetUtils: flagSetUtils,
		}

		config, err := GetConfigData(utilsStruct)
		utils.CheckError("Error in getting config: ", err)

		client := utils.ConnectToClient(config.Provider)

		err = utilsStruct.GetJobList(client)

		if err != nil {
			log.Error("Error in getting job list: ", err)
		}

	},
}

func (utilsStruct *UtilsStruct) GetJobList(client *ethclient.Client) error {
	jobs, err := utilsStruct.razorUtils.GetJobs(client)

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
	razorUtils = Utils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(jobListCmd)

}
