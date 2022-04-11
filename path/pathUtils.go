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
	GetLogFilePath(string) (string, error)
	GetConfigFilePath() (string, error)
	GetJobFilePath() (string, error)
	GetCommitDataFileName(string) (string, error)
	GetProposeDataFileName(string) (string, error)
	GetDisputeDataFileName(address string) (string, error)
}

type OSInterface interface {
	UserHomeDir() (string, error)
	Stat(string) (fs.FileInfo, error)
	IsNotExist(error) bool
	Mkdir(string, fs.FileMode) error
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	Open(string) (*os.File, error)
}

type PathUtils struct{}
type OSUtils struct{}

func (o OSUtils) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (o OSUtils) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (o OSUtils) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (o OSUtils) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

func (o OSUtils) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (o OSUtils) Open(name string) (*os.File, error) {
	return os.Open(name)
}
