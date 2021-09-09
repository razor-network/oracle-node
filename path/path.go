package path

import (
	"os"
)

func GetDefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultPath := home + "/.razor"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(defaultPath, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return defaultPath, nil
}

func GetLogFilePath() (string, error) {
	home, err := GetDefaultPath()
	if err != nil {
		return "", err
	}
	return home+"/razor.log", err
}

func GetConfigFilePath() (string, error) {
	home, err := GetDefaultPath()
	if err != nil {
		return "", err
	}
	return home + "/razor.yaml", nil
}
