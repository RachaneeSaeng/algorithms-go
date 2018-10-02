package main

import (
	"testing"

	"../lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Empty(t *testing.T) {
	result := findLoopBeginning(nil)
	require.Nil(t, result)
}

func Test_SingleItem(t *testing.T) {
	result := findLoopBeginning(lnklist.New(3, nil))
	require.Nil(t, result)
}

func Test_NoLoop(t *testing.T) {
	fifth := lnklist.New(5, nil)
	fourth := lnklist.New(4, fifth)
	third := lnklist.New(4, fourth)
	second := lnklist.New(2, third)
	first := lnklist.New(1, second)

	result := findLoopBeginning(first)
	require.Nil(t, result)
}

func Test_Loop(t *testing.T) {
	fifth := lnklist.New(5, nil)
	fourth := lnklist.New(4, fifth)
	third := lnklist.New(3, fourth)
	second := lnklist.New(2, third)
	first := lnklist.New(1, second)

	fifth.Next = third

	result := findLoopBeginning(first)

	require.Equal(t, third, result)
}
