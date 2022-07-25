//Package path provides all path related functions
package path

import (
	"os"
	pathPkg "path"
	"razor/core"
)

//This function returns the default path
func (PathUtils) GetDefaultPath() (string, error) {
	home, err := OSUtilsInterface.UserHomeDir()
	if err != nil {
		return "", err
	}
	defaultPath := pathPkg.Join(home, core.DefaultPathName)
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
	defaultPath := pathPkg.Join(razorPath, core.LogFile)
	if _, err := OSUtilsInterface.Stat(defaultPath); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(defaultPath, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}

	logFilepath := pathPkg.Join(defaultPath, fileName+".log")
	f, err := OSUtilsInterface.OpenFile(logFilepath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return logFilepath, nil
}

//This function returns the config file path
func (PathUtils) GetConfigFilePath() (string, error) {
	razorPath, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return pathPkg.Join(razorPath, core.ConfigFile), nil
}

//This function returns the job file path
func (PathUtils) GetJobFilePath() (string, error) {
	razorPath, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	filePath := pathPkg.Join(razorPath, core.AssetsDataFile)
	return filePath, nil
}

//This function returns the file name of commit data file
func (PathUtils) GetCommitDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := pathPkg.Join(razorDir, core.DataFileDirectory)
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}

	return pathPkg.Join(dataFileDir, address+core.CommitDataFile), nil
}

//This function returns the file name of propose data file
func (PathUtils) GetProposeDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := pathPkg.Join(razorDir, core.DataFileDirectory)
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return pathPkg.Join(dataFileDir, address+core.ProposeDataFile), nil
}

//This function returns the file name of dispute data file
func (PathUtils) GetDisputeDataFileName(address string) (string, error) {
	razorDir, err := PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return "", err
	}
	dataFileDir := pathPkg.Join(razorDir, core.DataFileDirectory)
	if _, err := OSUtilsInterface.Stat(dataFileDir); OSUtilsInterface.IsNotExist(err) {
		mkdirErr := OSUtilsInterface.Mkdir(dataFileDir, 0700)
		if mkdirErr != nil {
			return "", mkdirErr
		}
	}
	return pathPkg.Join(dataFileDir, address+core.DisputeDataFile), nil
}
