package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"razor/utils"
)

var setConfig = &cobra.Command{
	Use:   "setConfig",
	Short: "setConfig enables user to set the values of provider and gas multiplier",
	Long: `Setting the provider helps the CLI to know which provider to connect to.
Setting the gas multiplier value enables the CLI to multiply the gas with that value for all the transactions

Example:
  ./razor setConfig --provider https://infura/v3/matic --gasmultiplier 1.5 --buffer 20 --wait 70 --gasprice 1 --logLevel debug --gasLimit 5
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:   razorUtils,
			flagSetUtils: flagSetUtils,
		}
		err := utilsStruct.SetConfig(cmd.Flags())
		utils.CheckError("SetConfig error: ", err)
	},
}

func (utilsStruct UtilsStruct) SetConfig(flagSet *pflag.FlagSet) error {
	provider, err := utilsStruct.flagSetUtils.GetStringProvider(flagSet)
	if err != nil {
		return err
	}
	gasMultiplier, err := utilsStruct.flagSetUtils.GetFloat32GasMultiplier(flagSet)
	if err != nil {
		return err
	}
	bufferPercent, err := utilsStruct.flagSetUtils.GetInt32Buffer(flagSet)
	if err != nil {
		return err
	}
	waitTime, err := utilsStruct.flagSetUtils.GetInt32Wait(flagSet)
	if err != nil {
		return err
	}
	gasPrice, err := utilsStruct.flagSetUtils.GetInt32GasPrice(flagSet)
	if err != nil {
		return err
	}
	logLevel, err := utilsStruct.flagSetUtils.GetStringLogLevel(flagSet)
	if err != nil {
		return err
	}
	gasLimit, err := flagSetUtils.GetFloat32GasLimit(flagSet)
	if err != nil {
		return err
	}
	if provider != "" {
		viper.Set("provider", provider)
	}
	if gasMultiplier != -1 {
		viper.Set("gasmultiplier", gasMultiplier)
	}
	if bufferPercent != 0 {
		viper.Set("buffer", bufferPercent)
	}
	if waitTime != -1 {
		viper.Set("wait", waitTime)
	}
	if gasPrice != -1 {
		viper.Set("gasprice", gasPrice)
	}
	if logLevel != "" {
		viper.Set("logLevel", logLevel)
	}
	if gasLimit != -1 {
		viper.Set("gasLimit", gasLimit)
	}
	if provider == "" && gasMultiplier == -1 && bufferPercent == 0 && waitTime == -1 && gasPrice == -1 && logLevel == "" && gasLimit == -1 {
		viper.Set("provider", "http://127.0.0.1:8545")
		viper.Set("gasmultiplier", 1.0)
		viper.Set("buffer", 20)
		viper.Set("wait", 3)
		viper.Set("gasprice", 0)
		viper.Set("logLevel", "")
		viper.Set("gasLimit", 2)
		log.Info("Config values set to default. Use setConfig to modify the values.")
	}
	path, pathErr := utilsStruct.razorUtils.GetConfigFilePath()
	if pathErr != nil {
		log.Error("Error in fetching config file path")
		return pathErr
	}
	configErr := utilsStruct.razorUtils.ViperWriteConfigAs(path)
	if configErr != nil {
		log.Error("Error in writing config")
		return configErr
	}
	return nil
}

func init() {

	razorUtils = Utils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(setConfig)

	var (
		Provider           string
		GasMultiplier      float32
		BufferPercent      int32
		WaitTime           int32
		GasPrice           int32
		LogLevel           string
		GasLimitMultiplier float32
	)
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	setConfig.Flags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	setConfig.Flags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time (in secs)")
	setConfig.Flags().Int32VarP(&GasPrice, "gasprice", "", -1, "custom gas price")
	setConfig.Flags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	setConfig.Flags().Float32VarP(&GasLimitMultiplier, "gasLimit", "", -1, "gas limit percentage increase")

}
