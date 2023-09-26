// Package path provides all path related functions
package path

import (
	"io/fs"
	"os"
)

//go:generate mockery --name PathInterface --output ./mocks/ --case=underscore
//go:generate mockery --name OSInterface --output ./mocks/ --case=underscore

var PathUtilsInterface PathInterface
var OSUtilsInterface OSInterface

type PathInterface interface {
	GetDefaultPath() (string, error)
	GetLogFilePath(fileName string) (string, error)
	GetConfigFilePath() (string, error)
	GetJobFilePath() (string, error)
	GetCommitDataFileName(address string) (string, error)
	GetProposeDataFileName(address string) (string, error)
	GetDisputeDataFileName(address string) (string, error)
}

type OSInterface interface {
	UserHomeDir() (string, error)
	Stat(name string) (fs.FileInfo, error)
	IsNotExist(err error) bool
	Mkdir(name string, perm fs.FileMode) error
	OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error)
	Open(name string) (*os.File, error)
}

type PathUtils struct{}
type OSUtils struct{}

// This function returns the home directory of user
func (o OSUtils) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// This function is used to get status of the specified file path
func (o OSUtils) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

// This function returns the bool which shows the particular file exists or not
func (o OSUtils) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// This function is used to make a new directory
func (o OSUtils) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

// This function is used to open the file and this is generalized open call
func (o OSUtils) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// This function is used to open the file
func (o OSUtils) Open(name string) (*os.File, error) {
	return os.Open(name)
}
