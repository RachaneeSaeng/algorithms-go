package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maximizingXor function below.
func maximizingXor(l int32, r int32) int32 {
	a := l
	b := r
	// when r <= 1000, then 10 binary digit is enough
	for index := 9; index >= 0; index-- {
		lCurrentBit := getBit(l, index)
		rCurrentBit := getBit(r, index)

		var tempA int32
		var tempB int32
		if lCurrentBit == rCurrentBit && lCurrentBit == 0 { // result for this position = 0
			// try to see if we can make result to be 1
			tempA = setBit(a, index, 1)
			tempB = setBit(b, index, 1)
		} else if lCurrentBit == rCurrentBit && lCurrentBit == 1 { // result for this position = 0
			// try to see if we can make result to be 1
			tempA = setBit(a, index, 0)
			tempB = setBit(b, index, 0)
		}

		if tempA != 0 && l <= tempA && tempA <= b {
			a = tempA
			continue
		}
		if tempB != 0 && a <= tempB && tempB <= r {
			b = tempB
			continue
		}
	}
	return a ^ b
}

// get gata of bit from
func getBit(n int32, i int) int32 {
	result := n & (1 << uint(i))
	if result != 0 {
		return 1
	} else {
		return 0
	}
}

func setBit(n int32, i int, value int) int32 {
	clearer := ^(1 << uint(i))
	n = n & int32(clearer) // clear the i bit to be zero
	return n | int32(value<<uint(i))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	lTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	l := int32(lTemp)

	rTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	r := int32(rTemp)

	result := maximizingXor(l, r)

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
