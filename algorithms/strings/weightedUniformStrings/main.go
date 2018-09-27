package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type weightedUniform struct {
	weight, count int32
}

// Complete the weightedUniformStrings function below.
func weightedUniformStrings(s string, queries []int32) []string {
	uniforms := []weightedUniform{}

	var checkingCharacter rune
	var characterCount int32
	for _, c := range s {
		if checkingCharacter != c {
			if checkingCharacter > 0 {
				// save record of previous uniform
				weight := checkingCharacter - 96
				uniforms = append(uniforms, weightedUniform{int32(weight), characterCount})
			}

			checkingCharacter = c
			characterCount = 1
		} else {
			characterCount++
		}
	}

	if checkingCharacter > 0 {
		// save record of the last uniform
		weight := checkingCharacter - 96
		uniforms = append(uniforms, weightedUniform{int32(weight), characterCount})
	}

	var result []string
	for i, query := range queries {
		for _, u := range uniforms {
			if query <= (u.weight*u.count) && query%u.weight == 0 {
				result = append(result, "Yes")
				break
			}
		}
		if len(result) <= i {
			result = append(result, "No") // no match criteria
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
