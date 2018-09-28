package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumNumber function below.
// Return the minimum number of characters to make the password strong
func minimumNumber(n int32, password string) int32 {
	var digit, lowercase, uppercase, special bool
	for _, c := range password {
		if c >= 48 && c <= 57 {
			digit = true
		} else if c >= 65 && c <= 90 {
			uppercase = true
		} else if c >= 97 && c <= 122 {
			lowercase = true
		} else {
			switch c {
			case 33, 64, 35, 36, 37, 94, 38, 40, 41, 42, 43, 45:
				special = true
			}
		}

		if digit && lowercase && uppercase && special {
			break
		}
	}

	var missing int32
	if !digit {
		missing++
	}
	if !lowercase {
		missing++
	}
	if !uppercase {
		missing++
	}
	if !special {
		missing++
	}

	if missing > (6 - n) {
		return missing
	}
	return 6 - n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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
