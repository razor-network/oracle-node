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
		{
			name: "When strongPassword() returns true and every condition is met",
			args: args{
				input: "Qwerty12@",
			},
			want: true,
		},
		{
			name: "When strongPassword() returns false",
			args: args{
				input: "test",
			},
			want: false,
		},
		{
			name: "When password does not contain uppercase letters",
			args: args{
				input: "qwerty12@",
			},
			want: false,
		},
		{
			name: "When password does not contain lowercase letters",
			args: args{
				input: "QWERTY12@",
			},
			want: false,
		},
		{
			name: "When password does not contain digits",
			args: args{
				input: "qwerty!#%@",
			},
			want: false,
		},
		{
			name: "When password does not contain special character",
			args: args{
				input: "qwerty1234",
			},
			want: false,
		},
		{
			name: "When password is not long enough",
			args: args{
				input: "Qw1@",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPassword(tt.args.input); got != tt.want {
				t.Errorf("strongPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
