package utils

import (
	"bufio"
	"errors"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

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

func validate(input string) error {
	if input == "" {
		return errors.New("enter a valid password")
	}
	return nil
}

func validatePrivateKey(input string) error {
	if input == "" {
		return errors.New("enter a valid private key")
	}
	return nil
}

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

func AssignPassword(flagset *pflag.FlagSet) string {
	if IsFlagPassed("password") {
		passwordPath, _ := flagset.GetString("password")
		prompt := promptui.Prompt{
			Label:     "Setting password from file at described path, are you sure?",
			IsConfirm: true,
		}
		result, err := prompt.Run()
		CheckError(result, err)
		if strings.ToLower(result) == "y" {
			return GetPasswordFromFile(passwordPath)
		}
	}
	return PasswordPrompt()
}
