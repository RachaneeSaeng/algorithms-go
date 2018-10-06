package main

import (
	"testing"

	"../../linkedlist/lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Push_Empty(t *testing.T) {
	stk := Stack{}
	stk.Push(10)

	expectedResult := Stack{top: lnklist.New(10, nil)}
	require.Equal(t, expectedResult, stk)
}

func Test_Push_Existing(t *testing.T) {
	stk := Stack{top: lnklist.New(10, nil)}
	stk.Push(20)

	expectedResult := Stack{top: lnklist.New(20, lnklist.New(10, nil))}
	require.Equal(t, expectedResult, stk)
}

func Test_Pop_Empty(t *testing.T) {
	stk := Stack{}
	_, err := stk.Pop()

	require.Error(t, err)
}

func Test_Pop_Existing(t *testing.T) {
	stk := Stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result, err := stk.Pop()

	expectedStk := Stack{top: lnklist.New(10, nil)}

	require.NoError(t, err)
	require.Equal(t, 20, result)
	require.Equal(t, expectedStk, stk)
}

func Test_Peek_Empty(t *testing.T) {
	stk := Stack{}
	_, err := stk.Peek()

	require.Error(t, err)
}

func Test_Peek_Existing(t *testing.T) {
	stk := Stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result, err := stk.Peek()

	expectedStk := Stack{top: lnklist.New(20, lnklist.New(10, nil))}

	require.NoError(t, err)
	require.Equal(t, 20, result)
	require.Equal(t, expectedStk, stk)
}

func Test_IsEmpty_Empty(t *testing.T) {
	stk := Stack{}
	result := stk.IsEmpty()

	require.True(t, result)
}

func Test_IsEmpty_Existing(t *testing.T) {
	stk := Stack{top: lnklist.New(20, lnklist.New(10, nil))}
	result := stk.IsEmpty()

	require.False(t, result)
}
