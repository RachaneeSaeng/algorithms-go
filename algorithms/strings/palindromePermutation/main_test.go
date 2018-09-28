package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkOnlyOneBitSet_Zero(t *testing.T) {
	result := checkOnlyOneBitSet(0)
	require.Equal(t, true, result)
}

func Test_checkOnlyOneBitSet_True(t *testing.T) {
	result := checkOnlyOneBitSet(8)
	require.Equal(t, true, result)
}

func Test_checkOnlyOneBitSet_False(t *testing.T) {
	result := checkOnlyOneBitSet(9)
	require.Equal(t, false, result)
}

func Test_toggleBit_On(t *testing.T) {
	result := toggleBit(9, 2) // 00001001 ^ 00000100 = 00001101
	require.Equal(t, int32(13), result)
}

func Test_toggleBit_Off(t *testing.T) {
	result := toggleBit(9, 3) // 00001001 ^ 00001000 = 00000001
	require.Equal(t, int32(1), result)
}

func Test_EmptyString(t *testing.T) {
	result := isPermutationOfPalindrome("")
	require.Equal(t, false, result)
}

func Test_WithWhiteSpace(t *testing.T) {
	result := isPermutationOfPalindrome("tact coa")
	require.Equal(t, true, result)
}

func Test_NoSingleChar(t *testing.T) {
	result := isPermutationOfPalindrome("abab")
	require.Equal(t, true, result)
}

func Test_TooMuchSingleChar(t *testing.T) {
	result := isPermutationOfPalindrome("ababde")
	require.Equal(t, false, result)
}
