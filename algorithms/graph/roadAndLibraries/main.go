package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var connectedCitites map[int32][]int32
var visited []bool

func dfs(n int32) int32 {
	var adjNodes int32 = 0
	visited[n] = true
	if adjCities, exist := connectedCitites[n]; exist {
		for _, adjCity := range adjCities {
			if !visited[adjCity] {
				adjNodes++
				adjNodes += dfs(adjCity)
			}
		}
	}
	return adjNodes
}

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(numberOfCities int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if c_lib <= c_road {
		return int64(c_lib) * int64(numberOfCities)
	}
	// build graph
	var libs int32
	var roads int32
	connectedCitites = buildConnectedCitites(cities)
	visited = make([]bool, numberOfCities+1)
	for city := int32(1); city <= numberOfCities; city++ {
		if !visited[city] {
			libs++
			roads += dfs(city)
		}
	}

	return (int64(libs) * int64(c_lib)) + (int64(roads) * int64(c_road))
}

func buildConnectedCitites(cities [][]int32) map[int32][]int32 {
	connectedCitites := make(map[int32][]int32)
	for _, c := range cities {
		connectedCitites[c[0]] = append(connectedCitites[c[0]], c[1])
		connectedCitites[c[1]] = append(connectedCitites[c[1]], c[0])
	}
	return connectedCitites
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
		nmC_libC_road := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nmC_libC_road[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nmC_libC_road[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(nmC_libC_road[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(nmC_libC_road[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(readLine(reader), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

		fmt.Fprintf(writer, "%d\n", result)
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
