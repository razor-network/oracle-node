package path

import (
	"io"
	"os"
	"time"
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

func GetLogFilePath() (io.Writer, error) {
	home, err := GetDefaultPath()
	if err != nil {
		return nil, err
	}
	dt := time.Now().Format("2006-01-02_15.04.05")
	logFilePath, err := os.OpenFile(home+"/"+dt+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return logFilePath, nil
}

func GetConfigFilePath() (string, error) {
	home, err := GetDefaultPath()
	if err != nil {
		return "", err
	}
	return home + "/razor.yaml", nil
}
