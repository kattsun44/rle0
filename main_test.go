package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_compress(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "only A", args: args{input: "AAAAA"}, want: "A5"},
		{name: "hw", args: args{input: "Hello, World!"}, want: "H,e,l2,o,,, ,W,o,r,l,d,!"},
		{name: "only comma", args: args{input: ",,,,,,,,,,,,"}, want: ",12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.input); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
