package path

import (
	"os"
)

func (PathUtils) GetDefaultPath() (string, error) {
	home, err := PathUtilsInterface.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultPath := home + "/.razor"
	if _, err := PathUtilsInterface.Stat(defaultPath); PathUtilsInterface.IsNotExist(err) {
		mkdirErr := PathUtilsInterface.Mkdir(defaultPath, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return defaultPath, nil
}

func (PathUtils) GetLogFilePath() (string, error) {
	home, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return home + "/razor.log", err
}

func (PathUtils) GetConfigFilePath() (string, error) {
	home, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return home + "/razor.yaml", nil
}

func (PathUtils) GetJobFilePath() (string, error) {
	home, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	filePath := home + "/jobs.json"
	f, err := PathUtilsInterface.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return filePath, nil
}
