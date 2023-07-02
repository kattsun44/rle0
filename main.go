package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const defaultStr = "AAAAAAABBBBBBBCCCCDEEE"
	var str = flag.String("str", defaultStr, "use")
	flag.Parse()
	fmt.Println(*str)
	fmt.Println(compress(*str))
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
			output.WriteString(strconv.Itoa(count))
		}

		prev = c
		count = 1
	}

	output.WriteString(string(prev))
	output.WriteString(strconv.Itoa(count))

	return output.String()
}
