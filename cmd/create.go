package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"io/fs"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("dir")
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			log.Fatal("Please provide a password in order to create an account")
		}
		account := createAccount(path, password)
		log.Info("Account address: ", account.Address)
		log.Info("Keystore Path: ", account.URL)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	var (
		CustomPath string
		Password string
	)
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	defaultPath := home + "/.razor"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		os.Mkdir(defaultPath, fs.FileMode(0700))
	}
	createCmd.Flags().StringVarP(&CustomPath, "dir", "d", defaultPath, "directory path where the keystore must be stored")
	createCmd.Flags().StringVarP(&Password, "password", "", "", "password to create a new account")

	createCmd.MarkFlagRequired("password")
}
