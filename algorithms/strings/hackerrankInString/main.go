package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hackerrankInString function below.
// To find 'hackerrank' word hidden among other characters
func hackerrankInString(s string) string {
	hackerRank := [10]rune{'h', 'a', 'c', 'k', 'e', 'r', 'r', 'a', 'n', 'k'}
	if len(s) < 10 {
		return "NO"
	}

	checkingIndex := 0
	for _, c := range s {
		if hackerRank[checkingIndex] == c {
			checkingIndex++
			if checkingIndex == 10 {
				return "YES" //found all characters of 'hackerrank'
			}
		}
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := hackerrankInString(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
