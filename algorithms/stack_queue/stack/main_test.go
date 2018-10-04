package main

import (
	"testing"

	"../../linkedlist/lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Push_Empty(t *testing.T) {
	stk := stack{}
	stk.push(10)

	expectedResult := stack{top: lnklist.New(10, nil)}
	require.Equal(t, expectedResult, stk)
}

func Test_Push_Existing(t *testing.T) {
	stk := stack{top: lnklist.New(10, nil)}
	stk.push(20)

	expectedResult := stack{top: lnklist.New(20, lnklist.New(10, nil))}
	require.Equal(t, expectedResult, stk)
}

func Test_Pop_Empty(t *testing.T) {
	stk := stack{}
	_, err := stk.pop()

	require.Error(t, err)
}

func Test_Pop_Existing(t *testing.T) {
	stk := stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result, err := stk.pop()

	expectedStk := stack{top: lnklist.New(10, nil)}

	require.NoError(t, err)
	require.Equal(t, 20, result)
	require.Equal(t, expectedStk, stk)
}

func Test_Peek_Empty(t *testing.T) {
	stk := stack{}
	_, err := stk.peek()

	require.Error(t, err)
}

func Test_Peek_Existing(t *testing.T) {
	stk := stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result, err := stk.peek()

	expectedStk := stack{top: lnklist.New(20, lnklist.New(10, nil))}

	require.NoError(t, err)
	require.Equal(t, 20, result)
	require.Equal(t, expectedStk, stk)
}

func Test_IsEmpty_Empty(t *testing.T) {
	stk := stack{}
	result := stk.isEmpty()

	require.True(t, result)
}

func Test_IsEmpty_Existing(t *testing.T) {
	stk := stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result := stk.isEmpty()

	require.False(t, result)
}
