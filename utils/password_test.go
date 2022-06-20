package utils

import "testing"

func Test_strongPassword(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPassword(tt.args.input); got != tt.want {
				t.Errorf("strongPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
