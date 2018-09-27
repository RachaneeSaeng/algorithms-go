package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	result := pangrams("")
	require.Equal(t, "not pangram", result)
}

func Test_Pangram(t *testing.T) {
	result := pangrams("We promptly judged antique ivory buckles for the next prize")
	require.Equal(t, "pangram", result)
}

func Test_NotPangram(t *testing.T) {
	result := pangrams("We promptly judged antique ivory buckles for the prize")
	require.Equal(t, "not pangram", result)
}
