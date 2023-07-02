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
		d     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "only A", args: args{input: "AAAAA", d: ","}, want: "A5"},
		{name: "hw", args: args{input: "Hello, World!", d: ","}, want: "H1,e1,l2,o1,,1, 1,W1,o1,r1,l1,d1,!1"},
		{name: "only comma", args: args{input: ",,,,,,,,,,,,", d: ","}, want: ",12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.input, tt.args.d); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		input string
		d     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "blank", args: args{input: "", d: ","}, want: ""},
		{name: "aa", args: args{input: "a2", d: ","}, want: "aa"},
		{name: "a2b2", args: args{input: "a2,b2", d: ","}, want: "aabb"},
		{name: "a10b15", args: args{input: "a10,b15", d: ","}, want: "aaaaaaaaaabbbbbbbbbbbbbbb"},
		{name: "hw", args: args{input: "H1,e1,l2,o1,,1, 1,W1,o1,r1,l1,d1,!1", d: ","}, want: "Hello, World!"},
		{name: "only comma", args: args{input: ",12", d: ","}, want: "11"},
		{name: "only comma with colon delimiter", args: args{input: ",12", d: ":"}, want: ",,,,,,,,,,,,"},
		{name: "only comma with colon delimiter", args: args{input: "a1:,12", d: ":"}, want: "a,,,,,,,,,,,,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.input, tt.args.d); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
