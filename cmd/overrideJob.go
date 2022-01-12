package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core/types"
	"razor/path"
	"razor/utils"
)

// overrideJobCmd represents the overrideJob command
var overrideJobCmd = &cobra.Command{
	Use:   "overrideJob",
	Short: "overrideJob can be used to override existing job",
	Long:  ``,
	Run:   initialiseOverrideJob,
}

func initialiseOverrideJob(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteOverrideJob(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteOverrideJob(flagSet *pflag.FlagSet) {
	url, err := flagSetUtilsMockery.GetStringUrl(flagSet)
	utils.CheckError("Error in getting url: ", err)

	selector, err := flagSetUtilsMockery.GetStringSelector(flagSet)
	utils.CheckError("Error in getting selector: ", err)

	power, err := flagSetUtilsMockery.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	selectorType, err := flagSetUtilsMockery.GetUint8SelectorType(flagSet)
	utils.CheckError("Error in getting selector type: ", err)

	jobId, err := flagSetUtilsMockery.GetUint16JobId(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	job := &types.StructsJob{
		Id:           jobId,
		SelectorType: selectorType,
		Power:        power,
		Selector:     selector,
		Url:          url,
	}
	err = cmdUtilsMockery.OverrideJob(job)
	utils.CheckError("OverrideJob error: ", err)
	log.Info("Job added to override list successfully!")
}

func (*UtilsStructMockery) OverrideJob(job *types.StructsJob) error {
	jobPath, err := path.GetJobFilePath()
	if err != nil {
		return err
	}
	return utils.AddJobToJSON(jobPath, job)
}

func init() {

	cmdUtilsMockery = &UtilsStructMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}

	rootCmd.AddCommand(overrideJobCmd)
	var (
		JobId        uint16
		URL          string
		Selector     string
		SelectorType uint8
		Power        int8
	)

	overrideJobCmd.Flags().Uint16VarP(&JobId, "jobId", "j", 0, "job id to override")
	overrideJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	overrideJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath/XHTML selector)")
	overrideJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	overrideJobCmd.Flags().Uint8VarP(&SelectorType, "selectorType", "", 1, "selector type (1 for json, 2 for XHTML)")
}
