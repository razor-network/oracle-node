package utils

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"razor/core/types"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestReadJSONData(t *testing.T) {
	var fileName string

	type args struct {
		fileData     []byte
		fileErr      error
		unmarshalErr error
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*types.StructsJob
		wantErr error
	}{
		{
			name: "Test 1: When ReadJSONData() executes successfully",
			args: args{
				fileData:     []byte{},
				unmarshalErr: nil,
			},
			want:    map[string]*types.StructsJob{},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is a read file error",
			args: args{
				fileErr:      errors.New("readFile error"),
				unmarshalErr: nil,
			},
			want:    nil,
			wantErr: errors.New("readFile error"),
		},
		{
			name: "Test 3: When there is unmarshal error",
			args: args{
				fileData:     []byte{},
				unmarshalErr: errors.New("unmarshal error"),
			},
			want:    nil,
			wantErr: errors.New("unmarshal error"),
		},
		{
			name: "Test 4: When unmarshal error is unexpected end of JSON input",
			args: args{
				fileData:     []byte{},
				unmarshalErr: errors.New("unexpected end of JSON input"),
			},
			want:    map[string]*types.StructsJob{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
			}
			utils := StartRazor(optionsPackageStruct)

			osMock.On("ReadFile", mock.AnythingOfType("string")).Return(tt.args.fileData, tt.args.fileErr)
			jsonMock.On("Unmarshal", mock.Anything, mock.Anything).Return(tt.args.unmarshalErr)

			got, err := utils.ReadJSONData(fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadJSONData() got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for ReadJSONData(), got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for ReadJSONData(), got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestWriteDataToJSON(t *testing.T) {
	var fileName string
	var data map[string]*types.StructsJob

	type args struct {
		jsonString   []byte
		marshalErr   error
		writeFileErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When WriteDataToJSON() executes successfully",
			args: args{
				jsonString:   []byte{1, 2, 3},
				writeFileErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is a marshal error",
			args: args{
				marshalErr: errors.New("marshal error"),
			},
			wantErr: errors.New("marshal error"),
		},
		{
			name: "Test 3: When there is a writeFile error",
			args: args{
				jsonString:   []byte{1, 2, 3},
				writeFileErr: errors.New("writeFile error"),
			},
			wantErr: errors.New("writeFile error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMock := new(mocks.JsonUtils)
			osMock := new(mocks.OSUtils)

			optionsPackageStruct := OptionsPackageStruct{
				JsonInterface: jsonMock,
				OS:            osMock,
			}
			utils := StartRazor(optionsPackageStruct)

			jsonMock.On("Marshal", mock.Anything).Return(tt.args.jsonString, tt.args.marshalErr)
			osMock.On("WriteFile", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.writeFileErr)

			gotErr := utils.WriteDataToJSON(fileName, data)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for WriteDataToJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for WriteDataToJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func TestDeleteJobFromJSON(t *testing.T) {
	var fileName string

	type args struct {
		readData     map[string]*types.StructsJob
		readDataErr  error
		jobId        string
		writeDataErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When DeleteJobFromJSON() executes successfully",
			args: args{
				readData: map[string]*types.StructsJob{"1": &types.StructsJob{
					Id: 1, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini", Selector: "last",
					Url: "https://api.gemini.com/v1/pubticker/ethusd",
				}},
				jobId:        "1",
				writeDataErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When jobId is not prent in data",
			args: args{
				readData: map[string]*types.StructsJob{"1": &types.StructsJob{
					Id: 1, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini", Selector: "last",
					Url: "https://api.gemini.com/v1/pubticker/ethusd",
				}},
				jobId: "2",
			},
			wantErr: errors.New("No job with jobId = 2 found"),
		},
		{
			name: "Test 3: When there is an error in reading data",
			args: args{
				readDataErr: errors.New("readData error"),
			},
			wantErr: errors.New("readData error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}

			utilsMock.On("ReadJSONData", mock.AnythingOfType("string")).Return(tt.args.readData, tt.args.readDataErr)
			utilsMock.On("WriteDataToJSON", mock.Anything, mock.Anything).Return(tt.args.writeDataErr)

			utils := StartRazor(optionsPackageStruct)
			gotErr := utils.DeleteJobFromJSON(fileName, tt.args.jobId)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for DeleteJobFromJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for DeleteJobFromJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}

func TestAddJobToJSON(t *testing.T) {
	var fileName string

	type args struct {
		readData     map[string]*types.StructsJob
		readDataErr  error
		writeDataErr error
		job          *types.StructsJob
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test 1: When AddJobToJSON() executes successfully",
			args: args{
				readData: map[string]*types.StructsJob{"1": &types.StructsJob{
					Id: 1, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini", Selector: "last",
					Url: "https://api.gemini.com/v1/pubticker/ethusd",
				}},
				job: &types.StructsJob{
					Id: 3, SelectorType: 1, Weight: 100,
					Power: 2, Name: "ethusd_gemini1", Selector: "last",
					Url: "https://api.gemini.com/v1/pubticker/ethusd",
				},
				writeDataErr: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in reading data",
			args: args{
				readDataErr: errors.New("readData error"),
			},
			wantErr: errors.New("readData error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				UtilsInterface: utilsMock,
			}

			utilsMock.On("ReadJSONData", mock.AnythingOfType("string")).Return(tt.args.readData, tt.args.readDataErr)
			utilsMock.On("WriteDataToJSON", mock.Anything, mock.Anything).Return(tt.args.writeDataErr)

			utils := StartRazor(optionsPackageStruct)
			gotErr := utils.AddJobToJSON(fileName, tt.args.job)
			if gotErr == nil || tt.wantErr == nil {
				if gotErr != tt.wantErr {
					t.Errorf("Error for AddJobToJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			} else {
				if gotErr.Error() != tt.wantErr.Error() {
					t.Errorf("Error for AddJobToJSON(), got = %v, want = %v", gotErr, tt.wantErr)
				}
			}
		})
	}
}
