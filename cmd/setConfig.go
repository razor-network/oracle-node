//Package cmd provides all functions related to command line
package cmd

import (
	"razor/metrics"
	"razor/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
		err := cmdUtils.SetConfig(cmd.Flags())
		utils.CheckError("SetConfig error: ", err)
	},
}

//This function returns the error if there is any and sets the config
func (*UtilsStruct) SetConfig(flagSet *pflag.FlagSet) error {
	razorUtils.AssignLogFile(flagSet)
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
	gasLimit, err := flagSetUtils.GetFloat32GasLimit(flagSet)
	if err != nil {
		return err
	}
	logFileMaxSize, err := flagSetUtils.GetIntLogFileMaxSize(flagSet)
	if err != nil {
		return err
	}
	logFileMaxBackups, err := flagSetUtils.GetIntLogFileMaxBackups(flagSet)
	if err != nil {
		return err
	}
	logFileMaxAge, err := flagSetUtils.GetIntLogFileMaxAge(flagSet)
	if err != nil {
		return err
	}

	path, pathErr := razorUtils.GetConfigFilePath()
	if pathErr != nil {
		log.Error("Error in fetching config file path")
		return pathErr
	}

	if razorUtils.IsFlagPassed("exposeMetrics") {
		port, err := flagSetUtils.GetStringExposeMetrics(flagSet)
		if err != nil {
			return err
		}

		certKey, err := flagSetUtils.GetStringCertKey(flagSet)
		if err != nil {
			return err
		}
		certFile, err := flagSetUtils.GetStringCertFile(flagSet)
		if err != nil {
			return err
		}
		viper.Set("exposeMetricsPort", port)

		configErr := viperUtils.ViperWriteConfigAs(path)
		if configErr != nil {
			log.Error("Error in writing config")
			return configErr
		}

		err = metrics.Run(port, certFile, certKey)
		if err != nil {
			logrus.Errorf("failed to start metrics http server: %s", err)
		}
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
	if logFileMaxSize != 0 {
		viper.Set("logFileMaxSize", logFileMaxSize)
	}
	if logFileMaxBackups != 0 {
		viper.Set("logFileMaxBackups", logFileMaxBackups)
	}
	if logFileMaxAge != 0 {
		viper.Set("logFileMaxAge", logFileMaxAge)
	}
	if provider == "" && gasMultiplier == -1 && bufferPercent == 0 && waitTime == -1 && gasPrice == -1 && logLevel == "" && gasLimit == -1 && logFileMaxSize == 0 && logFileMaxBackups == 0 && logFileMaxAge == 0 {
		viper.Set("provider", "http://127.0.0.1:8545")
		viper.Set("gasmultiplier", 1.0)
		viper.Set("buffer", 20)
		viper.Set("wait", 3)
		viper.Set("gasprice", 1)
		viper.Set("logLevel", "")
		viper.Set("gasLimit", 2)
		//viper.Set("exposeMetricsPort", "")
		viper.Set("logFileMaxSize", 5)
		viper.Set("logFileMaxBackups", 10)
		viper.Set("logFileMaxAge", 30)
		log.Info("Config values set to default. Use setConfig to modify the values.")
	}

	configErr := viperUtils.ViperWriteConfigAs(path)
	if configErr != nil {
		log.Error("Error in writing config")
		return configErr
	}
	return nil
}

func init() {
	rootCmd.AddCommand(setConfig)

	var (
		Provider           string
		GasMultiplier      float32
		BufferPercent      int32
		WaitTime           int32
		GasPrice           int32
		LogLevel           string
		GasLimitMultiplier float32
		ExposeMetrics      string
		CertFile           string
		CertKey            string
		LogFileMaxSize     int
		LogFileMaxBackups  int
		LogFileMaxAge      int
	)
	setConfig.Flags().StringVarP(&Provider, "provider", "p", "", "provider name")
	setConfig.Flags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	setConfig.Flags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	setConfig.Flags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time (in secs)")
	setConfig.Flags().Int32VarP(&GasPrice, "gasprice", "", -1, "custom gas price")
	setConfig.Flags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	setConfig.Flags().Float32VarP(&GasLimitMultiplier, "gasLimit", "", -1, "gas limit percentage increase")
	setConfig.Flags().StringVarP(&ExposeMetrics, "exposeMetrics", "", "", "port number")
	setConfig.Flags().StringVarP(&CertFile, "certFile", "", "", "ssl certificate path")
	setConfig.Flags().StringVarP(&CertKey, "certKey", "", "", "ssl certificate key path")
	setConfig.Flags().IntVarP(&LogFileMaxSize, "logFileMaxSize", "", 0, "max size of log file in MB")
	setConfig.Flags().IntVarP(&LogFileMaxBackups, "logFileMaxBackups", "", 0, "max number of old log files to retain")
	setConfig.Flags().IntVarP(&LogFileMaxAge, "logFileMaxAge", "", 0, "max number of days to retain old log files")

}
