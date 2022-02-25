package path

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"os"
	"razor/path/mocks"
	"testing"
)

func TestGetDefaultPath(t *testing.T) {
	var fileInfo fs.FileInfo
	type args struct {
		homeDir    string
		homeDirErr error
		statErr    error
		isNotExist bool
		mkdirErr   error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When GetDefaultPath executes successfully",
			args: args{
				homeDir:    "/home",
				statErr:    nil,
				isNotExist: false,
			},
			want:    "/home/.razor",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting user home directory",
			args: args{
				homeDirErr: errors.New("homeDir error"),
			},
			want:    "",
			wantErr: errors.New("homeDir error"),
		},
		{
			name: "Test 3: When there is an error from Stat()",
			args: args{
				homeDir:    "/home",
				statErr:    errors.New("stat error"),
				isNotExist: true,
				mkdirErr:   nil,
			},
			want:    "/home/.razor",
			wantErr: nil,
		},
		{
			name: "Test 4: When there is an error from Stat() and than there is an error from Mkdir()",
			args: args{
				homeDir:    "/home",
				statErr:    errors.New("stat error"),
				isNotExist: true,
				mkdirErr:   errors.New("mkdir error"),
			},
			want:    "",
			wantErr: errors.New("mkdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osMock := new(mocks.OSInterface)
			OSUtilsInterface = osMock

			osMock.On("UserHomeDir").Return(tt.args.homeDir, tt.args.homeDirErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)

			pa := PathUtils{}
			got, err := pa.GetDefaultPath()
			if got != tt.want {
				t.Errorf("Path from GetDefaultPath function, got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetDefaultPath function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetDefaultPath function, got = %v, want = %v", err, tt.wantErr)
				}
			}

		})
	}
}

func TestGetLogFilePath(t *testing.T) {
	type args struct {
		path    string
		pathErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When GetLogFilePath() executes successfully",
			args: args{
				path: "./home/.razor",
			},
			want:    "./home/.razor/razor.log",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting home path",
			args: args{
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathMock := new(mocks.PathInterface)
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			pa := PathUtils{}
			got, err := pa.GetLogFilePath()
			if got != tt.want {
				t.Errorf("GetLogFilePath(), got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetLogFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetLogFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetConfigFilePath(t *testing.T) {
	type args struct {
		path    string
		pathErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When GetConfigFilePath() executes successfully",
			args: args{
				path: "./home/.razor",
			},
			want:    "./home/.razor/razor.yaml",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting home path",
			args: args{
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathMock := new(mocks.PathInterface)
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			pa := PathUtils{}
			got, err := pa.GetConfigFilePath()
			if got != tt.want {
				t.Errorf("GetConfigFilePath(), got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetConfigFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetConfigFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetJobFilePath(t *testing.T) {
	var file *os.File
	type args struct {
		path    string
		pathErr error
		file    *os.File
		fileErr error
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Test 1: When GetJobFilePath executes successfully",
			args: args{
				path: "/home/.razor",
				file: file,
			},
			want:    "/home/.razor/jobs.json",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting home path",
			args: args{
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 1: When there is an error in opening the file",
			args: args{
				path:    "/home/.razor",
				fileErr: errors.New("open file error"),
			},
			want:    "",
			wantErr: errors.New("open file error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathMock := new(mocks.PathInterface)
			osMock := new(mocks.OSInterface)
			PathUtilsInterface = pathMock
			OSUtilsInterface = osMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			osMock.On("OpenFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.file, tt.args.fileErr)
			pa := PathUtils{}
			got, err := pa.GetJobFilePath()
			if got != tt.want {
				t.Errorf("GetJobFilePath(), got = %v, want = %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetJobFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetJobFilePath function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
