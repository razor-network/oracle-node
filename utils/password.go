//Package utils provides the utils functions
package utils

import (
	"bufio"
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
	"os"
)

//This function prompts the password
func PasswordPrompt() string {
	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     ' ',
	}
	password, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return password
}

//This function prompts the private key
func PrivateKeyPrompt() string {
	prompt := promptui.Prompt{
		Label:    "🔑 Private Key",
		Validate: validatePrivateKey,
		Mask:     ' ',
	}
	privateKey, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

//This function validates the password
func validate(input string) error {
	if input == "" {
		return errors.New("enter a valid password")
	}
	return nil
}

//This function validates the private key
func validatePrivateKey(input string) error {
	if input == "" {
		return errors.New("enter a valid private key")
	}
	return nil
}

//This functon returns password from file
func GetPasswordFromFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Getting password from the first line of file at described location")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ""
}

//This function assigns the password
func AssignPassword(flagset *pflag.FlagSet) string {
	if UtilsInterface.IsFlagPassed("password") {
		log.Warn("Password flag is passed")
		passwordPath, _ := flagset.GetString("password")
		return GetPasswordFromFile(passwordPath)
	}
	return PasswordPrompt()
}
