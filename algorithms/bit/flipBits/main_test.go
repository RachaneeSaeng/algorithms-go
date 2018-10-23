package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetBit_1(t *testing.T) {
	result := flippingBits(2147483647)
	require.Equal(t, int64(2147483648), result)
}

func Test_GetBit_2(t *testing.T) {
	result := flippingBits(1)
	require.Equal(t, int64(4294967294), result)
}

func Test_GetBit_3(t *testing.T) {
	result := flippingBits(0)
	require.Equal(t, int64(4294967295), result)
}
