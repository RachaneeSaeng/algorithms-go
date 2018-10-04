package main

import (
	"testing"

	"../../linkedlist/lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Add_Empty(t *testing.T) {
	q := queue{}
	q.add(10)

	newNode := lnklist.New(10, nil)
	expectedResult := queue{first: newNode, last: newNode}

	require.Equal(t, expectedResult, q)
}

func Test_Add_Existing(t *testing.T) {
	q := queue{}
	q.add(10)
	q.add(20)

	expectedResult := queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	require.Equal(t, expectedResult, q)
}

func Test_Remove_Empty(t *testing.T) {
	q := queue{}
	_, err := q.remove()

	require.Error(t, err)
}

func Test_Remove_Existing(t *testing.T) {
	q := queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result, err := q.remove()

	expectedq := queue{first: lnklist.New(20, nil), last: lnklist.New(20, nil)}

	require.NoError(t, err)
	require.Equal(t, 10, result)
	require.Equal(t, expectedq, q)
}

func Test_Peek_Empty(t *testing.T) {
	q := queue{}
	_, err := q.peek()

	require.Error(t, err)
}

func Test_Peek_Existing(t *testing.T) {
	q := queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result, err := q.peek()

	expectedq := queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}

	require.NoError(t, err)
	require.Equal(t, 10, result)
	require.Equal(t, expectedq, q)
}

func Test_IsEmpty_Empty(t *testing.T) {
	q := queue{}
	result := q.isEmpty()

	require.True(t, result)
}

func Test_IsEmpty_Existing(t *testing.T) {
	q := queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result := q.isEmpty()

	require.False(t, result)
}
