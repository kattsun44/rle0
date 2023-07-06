package main

import (
	"flag"
	"fmt"
	"os"
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

	arg := flag.Arg(0)
	_, err := os.Stat(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rl0 : %v\n", err)
		os.Exit(1)
	}

	f, err := os.Open(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rl0 : %v\n", err)
		os.Exit(1)
	}

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rl0 : %v\n", err)
		os.Exit(1)
	}

	rle := encode(string(data[:count]), delimiter)
	rld := decode(rle, delimiter)

	fmt.Println(rld)
	fmt.Println(rle)

	// .rl 拡張子がつく新規ファイルを作り、エンコード結果を書き込む (拡張子の形式は暫定)
	nf, err := os.Create(arg + ".rl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "rl0 : %v\n", err)
		os.Exit(1)
	}

	nCount, err := nf.Write([]byte(rle))
	if err != nil {
		fmt.Fprintf(os.Stderr, "rl0 : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d bytes -> %d bytes (%.4g%%)\n", count, nCount, float64(nCount)/float64(count)*100)

	defer f.Close()
	defer nf.Close()
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
