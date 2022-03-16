package utils

import "testing"

func TestSecondsToReadableTime(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "Test 1",
			args: args{
				input: 100,
			},
			wantResult: "1 minute 40 seconds ",
		},
		{
			name: "Test 2",
			args: args{
				input: 10000000000,
			},
			wantResult: "45 years 11 months 4 weeks 2 days 17 hours 46 minutes 40 seconds ",
		},
		{
			name: "Test 3",
			args: args{
				input: 100000000,
			},
			wantResult: "5 months 15 weeks 2 days 9 hours 46 minutes 40 seconds ",
		},
		{
			name: "Test 4",
			args: args{
				input: 1000000,
			},
			wantResult: "1 week 4 days 13 hours 46 minutes 40 seconds ",
		},
		{
			name: "Test 5",
			args: args{
				input: 100000,
			},
			wantResult: "1 day 3 hours 46 minutes 40 seconds ",
		},
		{
			name: "Test 6",
			args: args{
				input: 10000,
			},
			wantResult: "2 hours 46 minutes 40 seconds ",
		},
		{
			name: "Test 7",
			args: args{
				input: 10,
			},
			wantResult: "10 seconds ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			if gotResult := ut.SecondsToReadableTime(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("SecondsToHuman() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
