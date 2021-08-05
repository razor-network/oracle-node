package utils

import (
	"errors"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
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
