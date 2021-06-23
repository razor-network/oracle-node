package cmd

import (
	"fmt"
	"os"
	"razor/utils"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Provider      string
	GasMultiplier float32
	BufferPercent int8
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "razor",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("Welcome to razor-cli.") },
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
	rootCmd.PersistentFlags().Float32VarP(&GasMultiplier, "gasmultiplier", "g", 1, "gas multiplier value")
	rootCmd.PersistentFlags().Int8VarP(&BufferPercent, "buffer", "b", 30, "buffer percent")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	home := utils.GetDefaultPath()
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
