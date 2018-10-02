package main

import (
	"testing"

	"../lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Empty(t *testing.T) {
	result := findIntersaction(nil, nil)
	require.Nil(t, result)
}

func Test_OneNil(t *testing.T) {
	result := findIntersaction(nil, lnklist.New(12, nil))
	require.Nil(t, result)
}

func Test_NotIntersect(t *testing.T) {
	result := findIntersaction(lnklist.New(1, nil), lnklist.New(1, nil))
	require.Nil(t, result)
}

func Test_NotIntersect_OneLonger(t *testing.T) {
	result := findIntersaction(lnklist.New(2, nil), lnklist.New(1, lnklist.New(2, nil)))
	require.Nil(t, result)
}

func Test_Intersect(t *testing.T) {
	nodeA := lnklist.New(9, nil)
	nodeB := lnklist.New(5, nodeA)
	l1 := lnklist.New(3, lnklist.New(4, lnklist.New(7, nodeB)))
	l2 := lnklist.New(3, lnklist.New(4, lnklist.New(7, nodeB)))

	result := findIntersaction(l1, l2)

	require.Equal(t, &nodeB, &result)
	require.Equal(t, nodeB, result)
}

func Test_Intersect_OneLonger(t *testing.T) {
	nodeA := lnklist.New(9, nil)
	nodeB := lnklist.New(5, nodeA)
	l1 := lnklist.New(3, lnklist.New(4, lnklist.New(7, nodeB)))
	l2 := lnklist.New(1, lnklist.New(1, lnklist.New(3, lnklist.New(4, lnklist.New(4, nodeB)))))

	result := findIntersaction(l1, l2)

	require.Equal(t, &nodeB, &result)
	require.Equal(t, nodeB, result)
}
