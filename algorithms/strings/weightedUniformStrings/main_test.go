package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	result := weightedUniformStrings("", []int32{1, 2, 3, 5})
	expectedResult := []string{"No", "No", "No", "No"}
	require.Equal(t, expectedResult, result)
}

func Test_EmptyQuery(t *testing.T) {
	result := weightedUniformStrings("abccddde", []int32{})
	expectedResult := []string{}
	require.Equal(t, expectedResult, result)
}

func Test_MultiQuery_1(t *testing.T) {
	result := weightedUniformStrings("abccddde", []int32{1, 3, 12, 5, 9, 10})
	expectedResult := []string{"Yes", "Yes", "Yes", "Yes", "No", "No"}
	require.Equal(t, expectedResult, result)
}

func Test_MultiQuery_2(t *testing.T) {
	result := weightedUniformStrings("aaabbbbcccddd", []int32{9, 7, 8, 12, 5})
	expectedResult := []string{"Yes", "No", "Yes", "Yes", "No"}
	require.Equal(t, expectedResult, result)
}
