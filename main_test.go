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
		{name: "hw", args: args{input: "Hello, World!"}, want: "H1,e1,l2,o1,,1, 1,W1,o1,r1,l1,d1,!1"},
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
		{name: "aa", args: args{input: "a2"}, want: "aa"},
		{name: "a2b2", args: args{input: "a2,b2"}, want: "aabb"},
		{name: "a10b15", args: args{input: "a10,b15"}, want: "aaaaaaaaaabbbbbbbbbbbbbbb"},
		{name: "hw", args: args{input: "H1,e1,l2,o1,,1, 1,W1,o1,r1,l1,d1,!1"}, want: "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.input); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
