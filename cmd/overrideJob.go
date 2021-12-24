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
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			assetManagerUtils: assetManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
		}
		err := utilsStruct.executeOverrideJob(cmd.Flags())
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Job added to override list successfully!")
	},
}

func (utilsStruct UtilsStruct) executeOverrideJob(flagSet *pflag.FlagSet) error {

	url, err := utilsStruct.flagSetUtils.GetStringUrl(flagSet)
	if err != nil {
		return err
	}

	selector, err := utilsStruct.flagSetUtils.GetStringSelector(flagSet)
	if err != nil {
		return err
	}

	power, err := utilsStruct.flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return err
	}

	selectorType, err := utilsStruct.flagSetUtils.GetUint8SelectorType(flagSet)
	if err != nil {
		return err
	}

	jobId, err := utilsStruct.flagSetUtils.GetUint16JobId(flagSet)
	if err != nil {
		return err
	}

	job := &types.StructsJob{
		Id:           jobId,
		SelectorType: selectorType,
		Power:        power,
		Selector:     selector,
		Url:          url,
	}
	return utilsStruct.overrideJob(job)
}

func (utilsStruct UtilsStruct) overrideJob(job *types.StructsJob) error {
	jobPath, err := path.GetJobFilePath()
	if err != nil {
		return err
	}
	return utils.AddJobToJSON(jobPath, job)
}

func init() {
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
