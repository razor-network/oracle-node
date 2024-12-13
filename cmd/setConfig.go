// Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/metrics"
	"razor/utils"

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
  ./razor setConfig --provider https://mainnet.skalenodes.com/v1/elated-tan-skat --gasmultiplier 1 --buffer 5 --wait 1 --gasprice 0 --logLevel debug --gasLimit 2
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmdUtils.SetConfig(cmd.Flags())
		utils.CheckError("SetConfig error: ", err)
	},
}

// This function returns the error if there is any and sets the config
func (*UtilsStruct) SetConfig(flagSet *pflag.FlagSet) error {
	_, _, _, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	flagDetails := []types.FlagDetail{
		{Name: "provider", Type: "string"},
		{Name: "gasmultiplier", Type: "float32"},
		{Name: "buffer", Type: "int32"},
		{Name: "wait", Type: "int32"},
		{Name: "gasprice", Type: "int32"},
		{Name: "logLevel", Type: "string"},
		{Name: "gasLimitOverride", Type: "uint64"},
		{Name: "gasLimit", Type: "float32"},
		{Name: "rpcTimeout", Type: "int64"},
		{Name: "httpTimeout", Type: "int64"},
		{Name: "logFileMaxSize", Type: "int"},
		{Name: "logFileMaxBackups", Type: "int"},
		{Name: "logFileMaxAge", Type: "int"},
	}

	// Storing the fetched flag values in a map
	flagValues := make(map[string]interface{})
	for _, flagDetail := range flagDetails {
		flagValue, err := flagSetUtils.FetchFlagInput(flagSet, flagDetail.Name, flagDetail.Type)
		if err != nil {
			log.Errorf("Error in fetching value for flag %v: %v", flagDetail.Name, err)
			return err
		}
		flagValues[flagDetail.Name] = flagValue
	}

	configDetails := []types.ConfigDetail{
		{FlagName: "provider", Key: "provider", DefaultValue: ""},
		{FlagName: "gasmultiplier", Key: "gasmultiplier", DefaultValue: core.DefaultGasMultiplier},
		{FlagName: "buffer", Key: "buffer", DefaultValue: core.DefaultBufferPercent},
		{FlagName: "wait", Key: "wait", DefaultValue: core.DefaultWaitTime},
		{FlagName: "gasprice", Key: "gasprice", DefaultValue: core.DefaultGasPrice},
		{FlagName: "logLevel", Key: "logLevel", DefaultValue: core.DefaultLogLevel},
		{FlagName: "gasLimitOverride", Key: "gasLimitOverride", DefaultValue: core.DefaultGasLimitOverride},
		{FlagName: "gasLimit", Key: "gasLimit", DefaultValue: core.DefaultGasLimit},
		{FlagName: "rpcTimeout", Key: "rpcTimeout", DefaultValue: core.DefaultRPCTimeout},
		{FlagName: "httpTimeout", Key: "httpTimeout", DefaultValue: core.DefaultHTTPTimeout},
		{FlagName: "logFileMaxSize", Key: "logFileMaxSize", DefaultValue: core.DefaultLogFileMaxSize},
		{FlagName: "logFileMaxBackups", Key: "logFileMaxBackups", DefaultValue: core.DefaultLogFileMaxBackups},
		{FlagName: "logFileMaxAge", Key: "logFileMaxAge", DefaultValue: core.DefaultLogFileMaxAge},
	}

	var areConfigSet bool

	// Setting the respective config values in config file only if the flag was set with a value in `setConfig` command
	for _, configDetail := range configDetails {
		if flagValue, exists := flagValues[configDetail.FlagName]; exists {
			// Check if the flag was set with a value in `setConfig` command
			if flagSetUtils.Changed(flagSet, configDetail.FlagName) {
				viper.Set(configDetail.Key, flagValue)
				areConfigSet = true
			}
		}
	}

	// If no config parameter was set than all the config parameters will be set to default config values
	if !areConfigSet {
		setDefaultConfigValues(configDetails)
	}

	path, pathErr := pathUtils.GetConfigFilePath()
	if pathErr != nil {
		log.Error("Error in fetching config file path")
		return pathErr
	}

	if razorUtils.IsFlagPassed("exposeMetrics") {
		metricsErr := handleMetrics(flagSet)
		if metricsErr != nil {
			log.Error("Error in handling metrics: ", metricsErr)
			return metricsErr
		}
	}

	configErr := viperUtils.ViperWriteConfigAs(path)
	if configErr != nil {
		log.Error("Error in writing config: ", configErr)
		return configErr
	}
	return nil
}

func setDefaultConfigValues(configDetails []types.ConfigDetail) {
	log.Info("No value is set to any flag in `setConfig` command")
	log.Info("Setting the config values to default. Use `setConfig` again to modify the values.")
	for _, configDetail := range configDetails {
		viper.Set(configDetail.Key, configDetail.DefaultValue)
	}
}

func handleMetrics(flagSet *pflag.FlagSet) error {
	port, err := flagSetUtils.FetchFlagInput(flagSet, "exposeMetrics", "string")
	if err != nil {
		return err
	}
	certKey, err := flagSetUtils.FetchFlagInput(flagSet, "certKey", "string")
	if err != nil {
		return err
	}
	certFile, err := flagSetUtils.FetchFlagInput(flagSet, "certFile", "string")
	if err != nil {
		return err
	}
	viper.Set("exposeMetricsPort", port)

	err = metrics.Run(port.(string), certFile.(string), certKey.(string))
	if err != nil {
		log.Error("Failed to start metrics http server: ", err)
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
		GasLimitOverride   uint64
		RPCTimeout         int64
		HTTPTimeout        int64
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
	setConfig.Flags().Int32VarP(&WaitTime, "wait", "w", 0, "wait time (in secs)")
	setConfig.Flags().Int32VarP(&GasPrice, "gasprice", "", -1, "custom gas price")
	setConfig.Flags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	setConfig.Flags().Float32VarP(&GasLimitMultiplier, "gasLimit", "", -1, "gas limit percentage increase")
	setConfig.Flags().Uint64VarP(&GasLimitOverride, "gasLimitOverride", "", 0, "gas limit to be over ridden for a transaction")
	setConfig.Flags().Int64VarP(&RPCTimeout, "rpcTimeout", "", 0, "RPC timeout if its not responding")
	setConfig.Flags().Int64VarP(&HTTPTimeout, "httpTimeout", "", 0, "http request timeout if its not responding")
	setConfig.Flags().StringVarP(&ExposeMetrics, "exposeMetrics", "", "", "port number")
	setConfig.Flags().StringVarP(&CertFile, "certFile", "", "", "ssl certificate path")
	setConfig.Flags().StringVarP(&CertKey, "certKey", "", "", "ssl certificate key path")
	setConfig.Flags().IntVarP(&LogFileMaxSize, "logFileMaxSize", "", 0, "max size of log file in MB")
	setConfig.Flags().IntVarP(&LogFileMaxBackups, "logFileMaxBackups", "", 0, "max number of old log files to retain")
	setConfig.Flags().IntVarP(&LogFileMaxAge, "logFileMaxAge", "", 0, "max number of days to retain old log files")

}
