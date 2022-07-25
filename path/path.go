//Package path provides all path related functions
package path

import (
	"os"
	"path/filepath"
)

//This function returns the default path
func (PathUtils) GetDefaultPath() (string, error) {
	home, err := OSUtilsInterface.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultPath := filepath.Join(home, ".razor")
	if _, err := OSUtilsInterface.Stat(defaultPath); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(defaultPath, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return defaultPath, nil
}

//This function returns the log file path
func (PathUtils) GetLogFilePath(fileName string) (string, error) {
	razorPath, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(razorPath, fileName+".log")
	f, err := OSUtilsInterface.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return filePath, nil
}

//This function returns the config file path
func (PathUtils) GetConfigFilePath() (string, error) {
	razorPath, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(razorPath, "razor.yaml"), nil
}

//This function returns the job file path
func (PathUtils) GetJobFilePath() (string, error) {
	razorPath, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(razorPath, "assets.json")
	return filePath, nil
}

//This function returns the file name of commit data file
func (PathUtils) GetCommitDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := filepath.Join(razorDir, "data_files")
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}

	return filepath.Join(dataFileDir, address+"_CommitData.json"), nil
}

//This function returns the file name of propose data file
func (PathUtils) GetProposeDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := filepath.Join(razorDir, "data_files")
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return filepath.Join(dataFileDir, address+"_proposedData.json"), nil
}

//This function returns the file name of dispute data file
func (PathUtils) GetDisputeDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := filepath.Join(razorDir, "data_files")
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return filepath.Join(dataFileDir, address+"_disputeData.json"), nil
}
