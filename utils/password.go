package utils

import (
	"bufio"
	"errors"
	"os"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
)

func (*UtilsStruct) PasswordPrompt() string {
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

func (*UtilsStruct) PrivateKeyPrompt() string {
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
	if input == "" || !strongPassword(input) {
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

func (*UtilsStruct) AssignPassword(flagSet *pflag.FlagSet) string {
	if UtilsInterface.IsFlagPassed("password") {
		log.Warn("Password flag is passed")
		log.Warn("This is a unsecure way to use razor-go")
		passwordPath, _ := flagSet.GetString("password")
		return GetPasswordFromFile(passwordPath)
	}
	return UtilsInterface.PasswordPrompt()
}

// This function checks if the password is strong enough or not
func strongPassword(input string) bool {
	l, u, p, d := 0, 0, 0, 0
	if len(input) >= 8 {
		for _, char := range input {
			switch {
			case unicode.IsUpper(char):
				u += 1
			case unicode.IsLower(char):
				l += 1
			case unicode.IsNumber(char):
				d += 1
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				p += 1
			}
		}
	}
	return (l >= 1 && u >= 1 && p >= 1 && d >= 1) && (l+u+p+d == len(input))
}
