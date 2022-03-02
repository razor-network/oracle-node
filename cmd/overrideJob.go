package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core/types"
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
	cmdUtils.ExecuteOverrideJob(cmd.Flags())
}

func (*UtilsStruct) ExecuteOverrideJob(flagSet *pflag.FlagSet) {
	url, err := flagSetUtils.GetStringUrl(flagSet)
	utils.CheckError("Error in getting url: ", err)

	selector, err := flagSetUtils.GetStringSelector(flagSet)
	utils.CheckError("Error in getting selector: ", err)

	power, err := flagSetUtils.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	selectorType, err := flagSetUtils.GetUint8SelectorType(flagSet)
	utils.CheckError("Error in getting selector type: ", err)

	jobId, err := flagSetUtils.GetUint16JobId(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	job := &types.StructsJob{
		Id:           jobId,
		SelectorType: selectorType,
		Power:        power,
		Selector:     selector,
		Url:          url,
	}
	err = cmdUtils.OverrideJob(job)
	utils.CheckError("OverrideJob error: ", err)
	log.Info("Job added to override list successfully!")
}

func (*UtilsStruct) OverrideJob(job *types.StructsJob) error {
	jobPath, err := razorUtils.GetJobFilePath()
	if err != nil {
		return err
	}
	return razorUtils.AddJobToJSON(jobPath, job)
}

func init() {

	razorUtils = Utils{}
	cmdUtils = &UtilsStruct{}
	flagSetUtils = FLagSetUtils{}
	InitializeUtils()

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
