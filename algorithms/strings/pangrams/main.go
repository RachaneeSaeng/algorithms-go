package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the pangrams function below. 
// pangrams is one or more sentences that consist of all alphabet
func pangrams(s string) string {
	var charset [26]bool
	charCount := 0

	for _, c := range s {
		index := -1
		if c >= 65 && c <= 90 {
			index = int(c) - 65
		} else if c >= 97 && c <= 122 {
			index = int(c) - 97
		}

		if index >= 0 && !charset[index] {
			charset[index] = true
			charCount++
			if charCount == 26 {
				return "pangram"
			}
		}
	}
	return "not pangram"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

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
