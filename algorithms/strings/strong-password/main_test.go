package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	result := minimumNumber(0, "")
	require.Equal(t, int32(6), result)
}

func Test_TooShortLen_MissRequired(t *testing.T) {
	result := minimumNumber(4, "abcd")
	require.Equal(t, int32(3), result)
}

func Test_TooShortLen_HaveAllRequired(t *testing.T) {
	result := minimumNumber(5, "1Ab$2")
	require.Equal(t, int32(1), result)
}

func Test_EnoughLen_MissRequired(t *testing.T) {
	result := minimumNumber(7, "abcdef$")
	require.Equal(t, int32(2), result)
}

func Test_EnoughLen_HaveAllRequired(t *testing.T) {
	result := minimumNumber(8, "Bc1$*edh")
	require.Equal(t, int32(0), result)
}
