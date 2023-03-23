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
)

var (
	Provider           string
	GasMultiplier      float32
	BufferPercent      int32
	WaitTime           int32
	GasPrice           int32
	LogLevel           string
	GasLimitMultiplier float32
	LogFile            string
	RPCTimeout         int64
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
	rootCmd.PersistentFlags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	rootCmd.PersistentFlags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	rootCmd.PersistentFlags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time")
	rootCmd.PersistentFlags().Int32VarP(&GasPrice, "gasprice", "", -1, "gas price")
	rootCmd.PersistentFlags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	rootCmd.PersistentFlags().Float32VarP(&GasLimitMultiplier, "gasLimit", "", -1, "gas limit percentage increase")
	rootCmd.PersistentFlags().StringVarP(&LogFile, "logFile", "", "", "name of log file")
	rootCmd.PersistentFlags().Int64VarP(&RPCTimeout, "rpcTimeout", "", 0, "RPC timeout if its not responding")
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
	log.Debugf("Gas Multiplier: %.2f", config.GasMultiplier)
	log.Debugf("Buffer Percent: %d", config.BufferPercent)
	log.Debugf("Wait Time: %d", config.WaitTime)
	log.Debugf("Gas Price: %d", config.GasPrice)
	log.Debugf("Log Level: %s", config.LogLevel)
	log.Debugf("Gas Limit: %.2f", config.GasLimitMultiplier)
	log.Debugf("RPC Timeout: %d", config.RPCTimeout)
}
