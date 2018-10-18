package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type centralCity struct {
	cityNum int
	count   int32
}

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(numberOfCities int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if c_lib <= c_road {
		return (int64)(c_lib * numberOfCities)
	}
	var lib int32
	var road int32
	hadLib := make(map[int32]bool)
	for len(cities) > 0 {
		centralCity := mostCentralCity(numberOfCities, cities, &hadLib)
		if centralCity > 0 {
			lib++
			road += getRoad(&cities, centralCity, &hadLib)
			removCities(&cities, hadLib)
		}
	}
	// city with no road
	lib += numberOfCities - int32(len(hadLib))
	return (int64)((lib * c_lib) + (road * c_road))
}

func mostCentralCity(numberOfCities int32, cities [][]int32, hadLib *map[int32]bool) int32 {
	cityCount := make([]int32, numberOfCities)
	for _, c := range cities {
		cityCount[c[0]-1]++
		cityCount[c[1]-1]++
	}
	var mostCentralCity *centralCity
	for index, c := range cityCount {
		if mostCentralCity == nil {
			mostCentralCity = &centralCity{cityNum: index + 1, count: c}
		} else {
			if _, exist := (*hadLib)[(int32)(index+1)]; mostCentralCity.count < c && !exist {
				mostCentralCity = &centralCity{cityNum: index + 1, count: c}
				(*hadLib)[(int32)(index+1)] = true
			}
		}
	}
	if mostCentralCity != nil && mostCentralCity.count > 0 {
		return (int32)(mostCentralCity.cityNum)
	}
	return 0
}

func getRoad(cities *[][]int32, cityNum int32, hadLib *map[int32]bool) int32 {
	var road int32
	for index := 0; index < len(*cities); index++ {
		city := (*cities)[index]
		_, firstExist := (*hadLib)[city[0]]
		_, secondExist := (*hadLib)[city[1]]
		if (city[0] == cityNum || city[1] == cityNum) && !(firstExist && secondExist) {
			road++
			(*hadLib)[city[0]] = true
			(*hadLib)[city[1]] = true
		}
	}
	return road
}

func removCities(cities *[][]int32, hadLib map[int32]bool) {
	for index := 0; index < len(*cities); index++ {
		city := (*cities)[index]
		_, firstExist := hadLib[city[0]]
		_, secondExist := hadLib[city[1]]
		if firstExist && secondExist {
			if index+1 < len(*cities) {
				*cities = append((*cities)[:index], (*cities)[index+1:]...)
			} else { // last element
				*cities = (*cities)[:index]
			}
			index--
		}
	}
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
