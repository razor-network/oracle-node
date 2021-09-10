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
	Provider      string
	GasMultiplier float32
	BufferPercent int32
	WaitTime      int32
	GasPrice      int32
	LogLevel      string
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

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Provider, "provider", "p", "", "provider name")
	rootCmd.PersistentFlags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", -1, "gas multiplier value")
	rootCmd.PersistentFlags().Int32VarP(&BufferPercent, "buffer", "b", 0, "buffer percent")
	rootCmd.PersistentFlags().Int32VarP(&WaitTime, "wait", "w", -1, "wait time")
	rootCmd.PersistentFlags().Int32VarP(&GasPrice, "gasprice", "", -1, "gas price")
	rootCmd.PersistentFlags().StringVarP(&LogLevel, "logLevel", "", "", "log level")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := path.GetDefaultPath()
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

	config, err := GetConfigData()
	if err != nil {
		log.Fatal(err)
	}

	if config.LogLevel == "debug" {
		log.SetLevel(logrus.DebugLevel)
	}
}
