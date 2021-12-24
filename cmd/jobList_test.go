package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/pkg/bindings"
	"testing"
)

func TestUtilsStruct_GetJobList(t *testing.T) {
	var client *ethclient.Client
	type fields struct {
		razorUtils UtilsMock
	}
	testUtils := fields{
		razorUtils: UtilsMock{},
	}

	jobListArray := []bindings.StructsJob{
		{Id: 1, SelectorType: 1, Weight: 100,
			Power: 2, Name: "ethusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/ethusd",
		},
		{Id: 2, SelectorType: 1, Weight: 100,
			Power: 2, Name: "btcusd_gemini", Selector: "last",
			Url: "https://api.gemini.com/v1/pubticker/btcusd",
		},
	}

	type args struct {
		client     *ethclient.Client
		jobList    []bindings.StructsJob
		jobListErr error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Test 1: When jobList executes properly",
			fields: testUtils,
			args: args{
				client:     client,
				jobList:    jobListArray,
				jobListErr: nil,
			},

			wantErr: nil,
		},
		{
			name:   "Test 2: When there is a error fetching job list ",
			fields: testUtils,
			args: args{
				client:     client,
				jobListErr: errors.New("error in fetching job list"),
			},
			wantErr: errors.New("error in fetching job list"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsStruct := &UtilsStruct{
				razorUtils: tt.fields.razorUtils,
			}

			GetJobsMock = func(*ethclient.Client) ([]bindings.StructsJob, error) {
				return tt.args.jobList, tt.args.jobListErr
			}
			err := utilsStruct.GetJobList(tt.args.client)

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for jobList function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for jobList function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
