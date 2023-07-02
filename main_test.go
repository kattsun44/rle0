package main

import (
	"testing"
)

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

func Test_encode(t *testing.T) {
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
			if got := encode(tt.args.input); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "blank", args: args{input: ""}, want: ""},
		// {name: "aa", args: args{input: "a2"}, want: "aa"},
		// {name: "foobar", args: args{input: "f,o2,b,a,r"}, want: "foobar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.input); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
