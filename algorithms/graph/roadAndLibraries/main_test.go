package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildConnectedCitites_NotConnected(t *testing.T) {
	cities := [][]int32{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	expectedResult := map[int32][]int32{
		1: []int32{2},
		2: []int32{1},
		3: []int32{4},
		4: []int32{3},
		5: []int32{6},
		6: []int32{5},
	}

	result := buildConnectedCitites(cities)
	require.Equal(t, expectedResult, result)
}

func Test_buildConnectedCitites_2Groups(t *testing.T) {
	cities := [][]int32{
		{1, 2},
		{1, 3},
		{4, 5},
	}
	expectedResult := map[int32][]int32{
		1: []int32{2, 3},
		2: []int32{1},
		3: []int32{1},
		4: []int32{5},
		5: []int32{4},
	}

	result := buildConnectedCitites(cities)
	require.Equal(t, expectedResult, result)
}

func Test_DFS_NotConnected(t *testing.T) {
	connectedCitites = map[int32][]int32{
		1: []int32{2},
		2: []int32{1},
		3: []int32{4},
		4: []int32{3},
		5: []int32{6},
		6: []int32{5},
	}
	visited = make([]bool, 7)
	expectedVisited := []bool{false, true, true, false, false, false, false}

	result := dfs(2)
	require.Equal(t, int32(1), result)
	require.Equal(t, expectedVisited, visited)
}

func Test_DFS_2Group(t *testing.T) {
	connectedCitites = map[int32][]int32{
		1: []int32{2, 3, 4},
		2: []int32{1},
		3: []int32{1},
		4: []int32{1, 5},
		5: []int32{4},
	}
	visited = make([]bool, 6)

	expectedVisited := []bool{false, true, true, true, true, true}
	result := dfs(1)
	require.Equal(t, int32(4), result)
	require.Equal(t, expectedVisited, visited)

	expectedVisited = []bool{false, true, true, true, true, true}
	result = dfs(4)
	require.Equal(t, int32(0), result)
	require.Equal(t, expectedVisited, visited)
}

func Test_roadsAndLibraries_Example1(t *testing.T) {
	cities := [][]int32{
		{1, 2},
		{3, 1},
		{2, 3},
	}
	result := roadsAndLibraries(3, 2, 1, cities)
	require.Equal(t, int64(4), result)
}

func Test_roadsAndLibraries_Example2(t *testing.T) {
	cities := [][]int32{
		{1, 2},
		{3, 1},
		{2, 3},
		{3, 4},
	}
	result := roadsAndLibraries(4, 2, 1, cities)
	require.Equal(t, int64(5), result)
}

func Test_roadsAndLibraries_Example3(t *testing.T) {
	cities := [][]int32{
		{1, 2},
		{3, 1},
		{2, 3},
		{4, 5},
	}
	result := roadsAndLibraries(5, 2, 1, cities)
	require.Equal(t, int64(7), result)
}

func Test_roadsAndLibraries_Example4(t *testing.T) {
	cities := [][]int32{
		{1, 3},
		{3, 4},
		{2, 4},
		{1, 2},
		{2, 3},
		{5, 6},
	}
	result := roadsAndLibraries(6, 2, 5, cities)
	require.Equal(t, int64(12), result)
}

func Test_roadsAndLibraries_Example5(t *testing.T) {
	cities := [][]int32{
		{8, 2},
		{2, 9},
	}
	result := roadsAndLibraries(9, 91, 84, cities)
	require.Equal(t, int64(805), result)
}

func Test_roadsAndLibraries_Example6(t *testing.T) {
	cities := [][]int32{
		{2, 1},
		{5, 3},
		{5, 1},
		{3, 4},
		{3, 1},
		{5, 4},
		{4, 1},
		{5, 2},
		{4, 2},
	}
	result := roadsAndLibraries(5, 92, 23, cities)
	require.Equal(t, int64(184), result)
}

func Test_roadsAndLibraries_Example7(t *testing.T) {
	cities := [][]int32{
		{6, 4},
		{3, 2},
		{7, 1},
	}
	result := roadsAndLibraries(8, 10, 55, cities)
	require.Equal(t, int64(80), result)
}

func Test_roadsAndLibraries_Example8(t *testing.T) {
	cities := [][]int32{}
	result := roadsAndLibraries(1, 5, 3, cities)
	require.Equal(t, int64(5), result)
}

func Test_roadsAndLibraries_Example9(t *testing.T) {
	cities := [][]int32{}
	result := roadsAndLibraries(2, 102, 1, cities)
	require.Equal(t, int64(204), result)
}
