//Package cmd provides all functions related to command line
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"razor/core"
	"razor/logger"
	"razor/path"
	"razor/utils"
)

var (
	Provider           string
	ChainId            string
	GasMultiplier      string
	BufferPercent      string
	WaitTime           string
	GasPrice           string
	LogLevel           string
	GasLimitMultiplier string
	LogFile            string
	LogFileMaxSize     string
	LogFileMaxBackups  string
	LogFileMaxAge      string
	Compress           string
)

var log = logger.NewLogger()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: core.VersionWithMeta,
	Use:     "razor [command] [flags]",
	Short:   "Official node for running stakers in Golang",
	Long:    `Razor can be used by the stakers to stake, delegate and vote on the razorscan. Stakers can vote correctly and earn rewards.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to razor-go.")
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//This function add the following command to the root command
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Provider, "provider", "p", "", "provider name")
	rootCmd.PersistentFlags().StringVarP(&ChainId, "chainId", "c", "0", "chainId")
	rootCmd.PersistentFlags().StringVarP(&GasMultiplier, "gasmultiplier", "g", "-1", "gas multiplier value")
	rootCmd.PersistentFlags().StringVarP(&BufferPercent, "buffer", "b", "0", "buffer percent")
	rootCmd.PersistentFlags().StringVarP(&WaitTime, "wait", "w", "-1", "wait time")
	rootCmd.PersistentFlags().StringVarP(&GasPrice, "gasprice", "", "-1", "gas price")
	rootCmd.PersistentFlags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	rootCmd.PersistentFlags().StringVarP(&GasLimitMultiplier, "gasLimit", "", "-1", "gas limit percentage increase")
	rootCmd.PersistentFlags().StringVarP(&LogFile, "logFile", "", "", "name of log file")
	rootCmd.PersistentFlags().StringVarP(&LogFileMaxSize, "logFileMaxSize", "", "0", "max size of log file MB")
	rootCmd.PersistentFlags().StringVarP(&LogFileMaxBackups, "logFileMaxBackups", "", "0", "max number of old log files to retain")
	rootCmd.PersistentFlags().StringVarP(&LogFileMaxAge, "logFileMaxAge", "", "0", "max number of days to retain old log files")
	rootCmd.PersistentFlags().StringVarP(&Compress, "compress", "", "true", "compression of logs file is true or false")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := path.PathUtilsInterface.GetDefaultPath()
	if err != nil {
		log.Fatal("Error in fetching .razor directory: ", err)
	}
	// Search config in home directory with name "razor.yaml".
	viper.AddConfigPath(home)
	viper.SetConfigName("razor")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("No config file found")
		} else {
			log.Warn("error in reading config")
		}
	}

	setLogLevel()
}

//This function sets the log level
func setLogLevel() {
	config, err := cmdUtils.GetConfigData()
	if err != nil {
		log.Fatal(err)
	}

	if config.LogLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Debug("Config details: ")
	log.Debugf("Provider: %s", config.Provider)
	log.Debugf("ChainId: %d", config.ChainId)
	log.Debugf("Gas Multiplier: %.2f", config.GasMultiplier)
	log.Debugf("Buffer Percent: %d", config.BufferPercent)
	log.Debugf("Wait Time: %d", config.WaitTime)
	log.Debugf("Gas Price: %d", config.GasPrice)
	log.Debugf("Log Level: %s", config.LogLevel)
	log.Debugf("Gas Limit: %.2f", config.GasLimitMultiplier)

	if utils.UtilsInterface.IsFlagPassed("logFile") {
		log.Debugf("Log File Max Size: %d MB", config.LogFileMaxSize)
		log.Debugf("Log File Max Backups (max number of old log files to retain): %d", config.LogFileMaxBackups)
		log.Debugf("Log File Max Age (max number of days to retain old log files): %d", config.LogFileMaxAge)
		log.Debugf("Compression for logFile: %s", config.Compress)
	}
}
