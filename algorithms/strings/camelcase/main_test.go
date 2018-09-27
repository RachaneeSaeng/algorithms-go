package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	result := camelcase("      ")
	require.Equal(t, int32(0), result)
}

func Test_InvalidFormat(t *testing.T) {
	result := camelcase("IamInValid")
	require.Equal(t, int32(0), result)
}

func Test_SpecialCharacter(t *testing.T) {
	result := camelcase("iamLove$")
	require.Equal(t, int32(0), result)
}

func Test_OneWords(t *testing.T) {
	result := camelcase("dfhdferikjfd")
	require.Equal(t, int32(1), result)
}

func Test_ManyWords(t *testing.T) {
	result := camelcase("iamValidCaseT")
	require.Equal(t, int32(4), result)
}
