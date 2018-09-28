package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the camelcase function below.
// check if string is in format camelCase
func camelcase(s string) int32 {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	var wordCount int32
	for _, c := range s {
		if !(c >= 65 && c <= 90) && !(c >= 97 && c <= 122) {
			return 0 // invalid character
		}

		if wordCount == 0 {
			if c >= 97 && c <= 122 {
				wordCount++
			} else {
				return 0 // invalid format (not start with lowercase)
			}
		} else if c >= 65 && c <= 90 {
			wordCount++
		}
	}
	return wordCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
