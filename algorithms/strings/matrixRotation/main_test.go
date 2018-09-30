package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0(t *testing.T) {
	matrix := [][]int32{}
	result := rotateMatrix(matrix, 180)
	expectedResult := [][]int32{}
	require.Equal(t, expectedResult, result)
}

func Test_2_180(t *testing.T) {
	matrix := [][]int32{
		{0, 1},
		{10, 11},
	}
	result := rotateMatrix(matrix, 180)
	expectedResult := [][]int32{
		{11, 10},
		{1, 0},
	}
	require.Equal(t, expectedResult, result)
}

func Test_3_90(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2},
		{10, 11, 12},
		{20, 21, 22},
	}
	result := rotateMatrix(matrix, 90)
	expectedResult := [][]int32{
		{20, 10, 0},
		{21, 11, 1},
		{22, 12, 2},
	}
	require.Equal(t, expectedResult, result)
}

func Test_3_270(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2},
		{10, 11, 12},
		{20, 21, 22},
	}
	result := rotateMatrix(matrix, 270)
	expectedResult := [][]int32{
		{2, 12, 22},
		{1, 11, 21},
		{0, 10, 20},
	}
	require.Equal(t, expectedResult, result)
}

func Test_4_90(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3},
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}
	result := rotateMatrix(matrix, 90)
	expectedResult := [][]int32{
		{30, 20, 10, 0},
		{31, 21, 11, 1},
		{32, 22, 12, 2},
		{33, 23, 13, 3},
	}
	require.Equal(t, expectedResult, result)
}

func Test_4_270(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3},
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}
	result := rotateMatrix(matrix, 270)
	expectedResult := [][]int32{
		{3, 13, 23, 33},
		{2, 12, 22, 32},
		{1, 11, 21, 31},
		{0, 10, 20, 30},
	}
	require.Equal(t, expectedResult, result)
}

func Test_4_360(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3},
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}
	result := rotateMatrix(matrix, 360)
	expectedResult := [][]int32{
		{0, 1, 2, 3},
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}
	require.Equal(t, expectedResult, result)
}

func Test_4_450(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3},
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}
	result := rotateMatrix(matrix, 450)
	expectedResult := [][]int32{
		{30, 20, 10, 0},
		{31, 21, 11, 1},
		{32, 22, 12, 2},
		{33, 23, 13, 3},
	}
	require.Equal(t, expectedResult, result)
}

func Test_5_90(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34},
		{40, 41, 42, 43, 44},
	}
	result := rotateMatrix(matrix, 90)
	expectedResult := [][]int32{
		{40, 30, 20, 10, 0},
		{41, 31, 21, 11, 1},
		{42, 32, 22, 12, 2},
		{43, 33, 23, 13, 3},
		{44, 34, 24, 14, 4},
	}
	require.Equal(t, expectedResult, result)
}

func Test_5_180(t *testing.T) {
	matrix := [][]int32{
		{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34},
		{40, 41, 42, 43, 44},
	}
	result := rotateMatrix(matrix, 180)
	expectedResult := [][]int32{
		{44, 43, 42, 41, 40},
		{34, 33, 32, 31, 30},
		{24, 23, 22, 21, 20},
		{14, 13, 12, 11, 10},
		{4, 3, 2, 1, 0},
	}
	require.Equal(t, expectedResult, result)
}
