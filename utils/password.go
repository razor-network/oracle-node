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

func validate(input string) error {
	if input == "" {
		return errors.New("enter a valid password")
	}
	return nil
}
