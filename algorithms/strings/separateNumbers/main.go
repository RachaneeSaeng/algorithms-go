package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the separateNumbers function below.
func separateNumbers(s string) {
	strLenght := len(s)
	if strLenght <= 1 || s[0] == '0' {
		fmt.Println("NO")
		return
	}

	for digit := 1; digit*2 < strLenght; digit++ {
		firstValue := s[0:digit]
		patternDigit := digit
		if firstValue == strings.Repeat("9", digit) {
			patternDigit = digit + 1
		}

		current := digit
		value, _ := strconv.ParseInt(firstValue, 10, 32)
		for current+patternDigit <= strLenght {
			nextValue, _ := strconv.ParseInt(s[current:current+patternDigit], 10, 32)
			if value+1 == nextValue {
				value = nextValue
				current = current + patternDigit
			} else {
				break
			}
		}

		if current == strLenght { // sequence
			fmt.Println("YES", firstValue)
			return
		}
	}
	fmt.Println("NO")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		separateNumbers(s)
	}
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
