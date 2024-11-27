//Package cmd provides all functions related to command line
package cmd

import (
	"fmt"
	"os"
	"razor/core"
	"razor/logger"
	"razor/path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Provider           string
	GasMultiplier      float32
	BufferPercent      int32
	WaitTime           int32
	GasPrice           int32
	LogLevel           string
	GasLimitMultiplier float32
	GasLimitOverride   uint64
	LogFile            string
	RPCTimeout         int64
	HTTPTimeout        int64
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
)

var log = logger.GetLogger()

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
	rootCmd.PersistentFlags().Uint64VarP(&GasLimitOverride, "gasLimitOverride", "", 0, "gas limit to be over ridden for a transaction")
	rootCmd.PersistentFlags().StringVarP(&LogFile, "logFile", "", "", "name of log file")
	rootCmd.PersistentFlags().Int64VarP(&RPCTimeout, "rpcTimeout", "", 0, "RPC timeout if its not responding")
	rootCmd.PersistentFlags().Int64VarP(&HTTPTimeout, "httpTimeout", "", 0, "HTTP request timeout if its not responding")
	rootCmd.PersistentFlags().IntVarP(&LogFileMaxSize, "logFileMaxSize", "", 0, "max size of log file MB")
	rootCmd.PersistentFlags().IntVarP(&LogFileMaxBackups, "logFileMaxBackups", "", 0, "max number of old log files to retain")
	rootCmd.PersistentFlags().IntVarP(&LogFileMaxAge, "logFileMaxAge", "", 0, "max number of days to retain old log files")
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
}
