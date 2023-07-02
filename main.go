package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0))
	fmt.Println(compress(flag.Arg(0)))
}

func compress(input string) string {
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
			if count != 1 {
				output.WriteString(strconv.Itoa(count))
			}
			output.WriteString(",")
		}

		prev = c
		count = 1
	}

	output.WriteString(string(prev))
	if count != 1 {
		output.WriteString(strconv.Itoa(count))
	}

	return output.String()
}
