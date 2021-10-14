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
  ./razor setConfig --provider https://infura/v3/matic --gasmultiplier 1.5 --buffer 20 --wait 70 --gasprice 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := SetConfig(cmd.Flags(), razorUtils, flagSetUtils)
		utils.CheckError("SetConfig error: ", err)
	},
}

func SetConfig(flagSet *pflag.FlagSet, razorUtils utilsInterface, flagSetUtils flagSetInterface) error {
	provider, err := flagSetUtils.GetStringProvider(flagSet)
	if err != nil {
		return err
	}
	gasMultiplier, err := flagSetUtils.GetFloat32GasMultiplier(flagSet)
	if err != nil {
		return err
	}
	bufferPercent, err := flagSetUtils.GetInt32Buffer(flagSet)
	if err != nil {
		return err
	}
	waitTime, err := flagSetUtils.GetInt32Wait(flagSet)
	if err != nil {
		return err
	}
	gasPrice, err := flagSetUtils.GetInt32GasPrice(flagSet)
	if err != nil {
		return err
	}
	logLevel, err := flagSetUtils.GetStringLogLevel(flagSet)
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
	if provider == "" && gasMultiplier == -1 && bufferPercent == 0 && waitTime == -1 && gasPrice == -1 && logLevel == "" {
		viper.Set("provider", "http://127.0.0.1:8545")
		viper.Set("gasmultiplier", 1.0)
		viper.Set("buffer", 20)
		viper.Set("wait", 3)
		viper.Set("gasprice", 0)
		viper.Set("logLevel", "")
		log.Info("Config values set to default. Use setConfig to modify the values.")
	}
	path, pathErr := razorUtils.GetConfigFilePath()
	if pathErr != nil {
		log.Error("Error in fetching config file path")
		return pathErr
	}
	configErr := razorUtils.ViperWriteConfigAs(path)
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
		Provider      string
		GasMultiplier float32
		BufferPercent int32
		WaitTime      int32
		GasPrice      int32
		LogLevel      string
	)
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	setConfig.Flags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	setConfig.Flags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time (in secs)")
	setConfig.Flags().Int32VarP(&GasPrice, "gasprice", "", -1, "custom gas price")
	setConfig.Flags().StringVarP(&LogLevel, "logLevel", "", "", "log level")

}
