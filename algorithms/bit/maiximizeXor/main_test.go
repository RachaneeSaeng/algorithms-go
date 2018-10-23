package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetBit(t *testing.T) {
	result := getBit(2, 0)
	require.Equal(t, int32(0), result)

	result = getBit(2, 1)
	require.Equal(t, int32(1), result)

	result = getBit(2, 2)
	require.Equal(t, int32(0), result)

	result = getBit(2, 5)
	require.Equal(t, int32(0), result)
}

func Test_SetBit(t *testing.T) {
	result := setBit(12, 0, 1)
	require.Equal(t, int32(13), result)

	result = setBit(12, 1, 1)
	require.Equal(t, int32(14), result)

	result = setBit(12, 2, 0)
	require.Equal(t, int32(8), result)

	result = setBit(12, 3, 1)
	require.Equal(t, int32(12), result)

	result = setBit(12, 5, 1)
	require.Equal(t, int32(44), result)
}

func Test_maximizingXor_1(t *testing.T) {
	result := maximizingXor(1, 2)
	require.Equal(t, int32(3), result)
}

func Test_maximizingXor_2(t *testing.T) {
	result := maximizingXor(10, 15)
	require.Equal(t, int32(7), result)
}

func Test_maximizingXor_3(t *testing.T) {
	result := maximizingXor(11, 100)
	require.Equal(t, int32(127), result)
}
