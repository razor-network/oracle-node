package path

import "os"

func (PathUtils) GetDefaultPath() (string, error) {
	home, err := OSUtilsInterface.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultPath := home + "/.razor"
	if _, err := OSUtilsInterface.Stat(defaultPath); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(defaultPath, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return defaultPath, nil
}

func (PathUtils) GetLogFilePath(fileName string) (string, error) {
	home, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	filePath := home + "/" + fileName + ".log"
	f, err := OSUtilsInterface.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return filePath, nil
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
	filePath := home + "/assets.json"
	return filePath, nil
}
