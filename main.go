package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	flag.Parse()
	fmt.Println(encode(flag.Arg(0)))
	fmt.Println(decode(encode(flag.Arg(0))))
}

func encode(input string) string {
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
			output.WriteString(",")
		}

		prev = c
		count = 1
	}

	output.WriteString(string(prev))
	output.WriteString(strconv.Itoa(count))

	return output.String()
}

func decode(input string) string {
	var prev, char rune
	var runes []rune
	var output strings.Builder
	var count int

	for _, c := range input {
		if prev == ',' || prev == 0 {
			runes = []rune{}
			char = c
		} else {
			if unicode.IsDigit(c) {
				runes = append(runes, c)
			}
		}

		count, _ = strconv.Atoi(string(runes))
		if c == ',' {
			output.WriteString(strings.Repeat(string(char), count))
			if prev == ',' {
				output.WriteString(",")
			}
		}

		prev = c
	}

	output.WriteString(strings.Repeat(string(char), count))

	return output.String()
}
