package main_test

import (
	main "github.com/evgensr/practicum1/cmd/shortener"
	"testing"
)

func Test_getHash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pass md5",
			args: args{"1"},
			want: "c4ca4238a0b923820dcc509a6f75849b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.GetHash(tt.args.text); got != tt.want {
				t.Errorf("getHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
