package path

import (
	"io/fs"
	"os"
)

//go:generate mockery --name PathInterface --output ./mocks/ --case=underscore

var PathUtilsInterface PathInterface

type PathInterface interface {
	UserHomeDir() (string, error)
	Stat(string) (fs.FileInfo, error)
	IsNotExist(error) bool
	Mkdir(string, fs.FileMode) error
	GetDefaultPath() (string, error)
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	GetLogFilePath() (string, error)
	GetConfigFilePath() (string, error)
	GetJobFilePath() (string, error)
	Open(string) (*os.File, error)
}

type PathUtils struct{}

func (p PathUtils) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (p PathUtils) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (p PathUtils) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (p PathUtils) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

func (p PathUtils) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (p PathUtils) Open(name string) (*os.File, error) {
	return os.Open(name)
}
