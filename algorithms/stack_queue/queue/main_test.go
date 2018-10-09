package queue

import (
	"testing"

	"../../linkedlist/lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Add_Empty(t *testing.T) {
	q := Queue{}
	q.Add(10)

	newNode := lnklist.New(10, nil)
	expectedResult := Queue{first: newNode, last: newNode}

	require.Equal(t, expectedResult, q)
}

func Test_Add_Existing(t *testing.T) {
	q := Queue{}
	q.Add(10)
	q.Add(20)

	expectedResult := Queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	require.Equal(t, expectedResult, q)
}

func Test_Remove_Empty(t *testing.T) {
	q := Queue{}
	_, err := q.Remove()

	require.Error(t, err)
}

func Test_Remove_Existing(t *testing.T) {
	q := Queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result, err := q.Remove()

	expectedq := Queue{first: lnklist.New(20, nil), last: lnklist.New(20, nil)}

	require.NoError(t, err)
	require.Equal(t, 10, result)
	require.Equal(t, expectedq, q)
}

func Test_Peek_Empty(t *testing.T) {
	q := Queue{}
	_, err := q.Peek()

	require.Error(t, err)
}

func Test_Peek_Existing(t *testing.T) {
	q := Queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result, err := q.Peek()

	expectedq := Queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}

	require.NoError(t, err)
	require.Equal(t, 10, result)
	require.Equal(t, expectedq, q)
}

func Test_IsEmpty_Empty(t *testing.T) {
	q := Queue{}
	result := q.IsEmpty()

	require.True(t, result)
}

func Test_IsEmpty_Existing(t *testing.T) {
	q := Queue{first: lnklist.New(10, lnklist.New(20, nil)), last: lnklist.New(20, nil)}
	result := q.IsEmpty()

	require.False(t, result)
}
