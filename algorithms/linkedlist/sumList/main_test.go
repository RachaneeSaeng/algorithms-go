package main

import (
	"testing"

	"../lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Empty(t *testing.T) {
	result := sumList(nil, nil)
	require.Nil(t, result)
}

func Test_L1Nil(t *testing.T) {
	expectedResult := lnklist.New(5, nil)

	result := sumList(nil, lnklist.New(5, nil))

	require.Equal(t, expectedResult, result)
}

func Test_L2Nil(t *testing.T) {
	expectedResult := lnklist.New(5, nil)

	result := sumList(lnklist.New(5, nil), nil)

	require.Equal(t, expectedResult, result)
}

func Test_AllZero(t *testing.T) {
	l1 := lnklist.New(0, lnklist.New(0, nil))
	l2 := lnklist.New(0, lnklist.New(0, nil))

	expectedResult := lnklist.New(0, lnklist.New(0, nil))

	result := sumList(l1, l2)

	require.Equal(t, expectedResult, result)
}

func Test_WitoutCarry(t *testing.T) {
	l1 := lnklist.New(2, lnklist.New(5, nil))
	l2 := lnklist.New(3, lnklist.New(7, lnklist.New(4, nil)))

	expectedResult := lnklist.New(3, lnklist.New(9, lnklist.New(9, nil)))

	result := sumList(l1, l2)

	require.Equal(t, expectedResult, result)
}

func Test_WithCarry(t *testing.T) {
	l1 := lnklist.New(2, lnklist.New(5, nil))
	l2 := lnklist.New(3, lnklist.New(7, lnklist.New(5, nil)))

	expectedResult := lnklist.New(4, lnklist.New(0, lnklist.New(0, nil)))

	result := sumList(l1, l2)

	require.Equal(t, expectedResult, result)
}

func Test_Example(t *testing.T) {
	l1 := lnklist.New(6, lnklist.New(1, lnklist.New(7, nil)))
	l2 := lnklist.New(2, lnklist.New(9, lnklist.New(5, nil)))

	expectedResult := lnklist.New(9, lnklist.New(1, lnklist.New(2, nil)))

	result := sumList(l1, l2)

	require.Equal(t, expectedResult, result)
}
