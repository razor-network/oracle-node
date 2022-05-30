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
	var fileName string
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
			name: "Test 1: When GetLogFilePath() executes successfully",
			args: args{
				path: "./home/.razor",
			},
			want:    "./home/.razor/.log",
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
			name: "Test 3: When there is an error in getting file",
			args: args{
				path:    "./home/.razor",
				fileErr: errors.New("error in getting file"),
			},
			want:    "",
			wantErr: errors.New("error in getting file"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathMock := new(mocks.PathInterface)
			osMock := new(mocks.OSInterface)

			OSUtilsInterface = osMock
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			osMock.On("OpenFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.file, tt.args.fileErr)
			pa := PathUtils{}
			got, err := pa.GetLogFilePath(fileName)
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
			name: "Test 1: When GetJobFilePath executes successfully",
			args: args{
				path: "/home/.razor",
			},
			want:    "/home/.razor/assets.json",
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
			osMock := new(mocks.OSInterface)
			PathUtilsInterface = pathMock
			OSUtilsInterface = osMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
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

func TestGetCommitDataFileName(t *testing.T) {
	var fileInfo fs.FileInfo

	type args struct {
		address    string
		path       string
		pathErr    error
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
			name: "Test 1: When GetCommitDataFileName() executes successfully",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				path:    "/home",
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_CommitData.json",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 3: When dataFiles directory is not present and mkdir creates it",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_CommitData.json",
			wantErr: nil,
		},
		{
			name: "Test 4: When dataFiles directory is not present and there is an error in creating new one",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
				mkdirErr:   errors.New("mkdir error"),
			},
			want:    "",
			wantErr: errors.New("mkdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pathMock := new(mocks.PathInterface)
			osMock := new(mocks.OSInterface)

			OSUtilsInterface = osMock
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)

			pa := &PathUtils{}
			got, err := pa.GetCommitDataFileName(tt.args.address)
			if got != tt.want {
				t.Errorf("GetCommitDataFileName() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetCommitDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetCommitDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetProposeDataFileName(t *testing.T) {
	var fileInfo fs.FileInfo

	type args struct {
		address    string
		path       string
		pathErr    error
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
			name: "Test 1: When GetProposeDataFileName() executes successfully",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				path:    "/home",
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_proposedData.json",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 3: When dataFiles directory is not present and mkdir creates it",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_proposedData.json",
			wantErr: nil,
		},
		{
			name: "Test 4: When dataFiles directory is not present and there is an error in creating new one",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
				mkdirErr:   errors.New("mkdir error"),
			},
			want:    "",
			wantErr: errors.New("mkdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pathMock := new(mocks.PathInterface)
			osMock := new(mocks.OSInterface)

			OSUtilsInterface = osMock
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)

			pa := &PathUtils{}
			got, err := pa.GetProposeDataFileName(tt.args.address)
			if got != tt.want {
				t.Errorf("GetProposeDataFileName() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetProposeDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetProposeDataFileName(), got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetDisputeDataFileName(t *testing.T) {
	var fileInfo fs.FileInfo

	type args struct {
		address    string
		path       string
		pathErr    error
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
			name: "Test 1: When GetDisputeDataFileName executes successfully",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				path:    "/home",
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_disputeData.json",
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting path",
			args: args{
				address: "0x000000000000000000000000000000000000dead",
				pathErr: errors.New("path error"),
			},
			want:    "",
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 3: When dataFiles directory is not present and mkdir creates it",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
			},
			want:    "/home/dataFiles/0x000000000000000000000000000000000000dead_disputeData.json",
			wantErr: nil,
		},
		{
			name: "Test 4: When dataFiles directory is not present and there is an error in creating new one",
			args: args{
				address:    "0x000000000000000000000000000000000000dead",
				path:       "/home",
				statErr:    errors.New("not exists"),
				isNotExist: true,
				mkdirErr:   errors.New("mkdir error"),
			},
			want:    "",
			wantErr: errors.New("mkdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pathMock := new(mocks.PathInterface)
			osMock := new(mocks.OSInterface)

			OSUtilsInterface = osMock
			PathUtilsInterface = pathMock

			pathMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)

			pa := &PathUtils{}
			got, err := pa.GetDisputeDataFileName(tt.args.address)
			if got != tt.want {
				t.Errorf("GetDisputeDataFileName got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for GetDisputeDataFileName, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for GetDisputeDataFileName, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
