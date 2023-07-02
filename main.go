package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const defaultDelimiter = ","

func main() {
	var delimiter string
	flag.StringVar(&delimiter, "delimiter", defaultDelimiter, "Set delimiter (default: comma)")
	flag.StringVar(&delimiter, "d", defaultDelimiter, "Set delimiter by short (default: comma)")
	flag.Parse()

	// TODO: args 未入力時の挙動修正
	rle := encode(flag.Arg(0), delimiter)
	rld := decode(rle, delimiter)

	fmt.Println(rle)
	fmt.Println(rld)
}

func encode(input string, d string) string {
	count := 1
	var prev rune
	var output strings.Builder
	for _, c := range input {
		if prev == c {
			count++
			continue
		}

		if prev != 0 {
			output.WriteString(string(prev))
			output.WriteString(strconv.Itoa(count))
			output.WriteString(d)
		}

		prev = c
		count = 1
	}

	output.WriteString(string(prev))
	output.WriteString(strconv.Itoa(count))

	return output.String()
}

func decode(input string, d string) string {
	var prev, char rune
	var runes []rune
	var output strings.Builder
	var count int

	for _, c := range input {
		if string(prev) == d || prev == 0 {
			runes = []rune{}
			char = c
		} else {
			if unicode.IsDigit(c) {
				runes = append(runes, c)
			}
		}

		count, _ = strconv.Atoi(string(runes))
		if string(c) == d {
			output.WriteString(strings.Repeat(string(char), count))
			if string(prev) == d {
				output.WriteString(d)
			}
		}

		prev = c
	}

	output.WriteString(strings.Repeat(string(char), count))

	return output.String()
}
