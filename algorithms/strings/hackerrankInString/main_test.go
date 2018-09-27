package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	result := hackerrankInString("")
	require.Equal(t, "NO", result)
}

func Test_TooShort(t *testing.T) {
	result := hackerrankInString("hackerran")
	require.Equal(t, "NO", result)
}

func Test_Yes_Short(t *testing.T) {
	result := hackerrankInString("hereiamstackerrank")
	require.Equal(t, "YES", result)
}

func Test_Yes_Long(t *testing.T) {
	result := hackerrankInString("hereiamstackerrankkdfjdkferijfkajfdfjdjvkfjkjkjfdkjaerieriujfdkjfdlkjeriueakjfdf")
	require.Equal(t, "YES", result)
}

func Test_No_Short(t *testing.T) {
	result := hackerrankInString("hackerworld")
	require.Equal(t, "NO", result)
}

func Test_No_Long(t *testing.T) {
	result := hackerrankInString("rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt")
	require.Equal(t, "NO", result)
}
