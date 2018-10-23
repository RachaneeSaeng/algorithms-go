package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumXor_1(t *testing.T) {
	result := sumXor(5)
	require.Equal(t, int64(2), result)
}

func Test_sumXor_2(t *testing.T) {
	result := sumXor(10)
	require.Equal(t, int64(4), result)
}

func Test_sumXor_3(t *testing.T) {
	result := sumXor(0)
	require.Equal(t, int64(1), result)
}

func Test_sumXor_4(t *testing.T) {
	result := sumXor(0)
	require.Equal(t, int64(1), result)
}
