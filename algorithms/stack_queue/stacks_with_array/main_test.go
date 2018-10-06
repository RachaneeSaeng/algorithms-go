package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Push_Invalid(t *testing.T) {
	stks := new(5)
	result := stks.push(6, 2)

	require.Error(t, result)
}

func Test_Push_FirstPush(t *testing.T) {
	stks := new(5)
	result := stks.push(2, 2)
	require.NoError(t, result)

	expectedResult := []int{0, 2, 0, 0, 0}

	require.Equal(t, expectedResult, stks.arr)
}

func Test_Push_Existing_1(t *testing.T) {
	stks := new(5)
	result := stks.push(3, 2)
	require.NoError(t, result)

	result = stks.push(3, 5)
	require.NoError(t, result)

	result = stks.push(1, 7)
	require.NoError(t, result)

	result = stks.push(1, 3)
	require.NoError(t, result)

	expectedResult := []int{7, 0, 2, 0, 0, 3, 0, 5}

	require.Equal(t, expectedResult, stks.arr)
}

func Test_Pop_Invalid(t *testing.T) {
	stks := new(5)
	_, err := stks.pop(8)

	require.Error(t, err)
}

func Test_Pop_Empty(t *testing.T) {
	stks := new(5)
	_, err := stks.pop(1)

	require.Error(t, err)
}

func Test_Pop_InvalidIndex(t *testing.T) {
	stks := new(5)
	stks.push(1, 45)
	stks.tops[0] = 9
	_, err := stks.pop(1)

	require.Error(t, err)
}

func Test_Pop_Normal(t *testing.T) {
	stks := new(5)
	stks.push(2, 99)
	stks.push(2, 88)
	stks.push(2, 77)
	stks.push(1, 9)
	stks.push(1, 8)
	stks.push(1, 7)

	result, err := stks.pop(2)
	require.NoError(t, err)
	require.Equal(t, 77, result)

	result, err = stks.pop(2)
	require.NoError(t, err)
	require.Equal(t, 88, result)

	result, err = stks.pop(1)
	require.NoError(t, err)
	require.Equal(t, 7, result)
}

func Test_Pop_UntilEmpty(t *testing.T) {
	stks := new(5)
	stks.push(2, 99)
	stks.push(2, 88)
	stks.push(2, 77)

	result, err := stks.pop(2)
	require.NoError(t, err)
	require.Equal(t, 77, result)

	result, err = stks.pop(2)
	require.NoError(t, err)
	require.Equal(t, 88, result)

	result, err = stks.pop(2)
	require.NoError(t, err)
	require.Equal(t, 99, result)

	_, err = stks.pop(2)
	require.Error(t, err)
}
