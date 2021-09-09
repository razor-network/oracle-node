package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"razor/logger"
	"razor/path"
)

var (
	Provider      string
	GasMultiplier float32
	BufferPercent int32
	WaitTime      int32
	GasPrice      int32
)

const (
	VersionMajor = 0          // Major version component of the current release
	VersionMinor = 1          // Minor version component of the current release
	VersionPatch = 6          // Patch version component of the current release
	VersionMeta  = "unstable" // Version metadata to append to the version string
)

// Version holds the textual version string.
var Version = func() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}()

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := Version
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()

var log = logger.NewLogger()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: VersionWithMeta,
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
}
