package main

import (
	"testing"

	"../lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Partition_Empty(t *testing.T) {
	result := isPalindrome(nil)
	require.True(t, result)
}

func Test_Partition_Single(t *testing.T) {
	result := isPalindrome(lnklist.New(10, nil))
	require.True(t, result)
}

func Test_Partition_Palindrome(t *testing.T) {
	result := isPalindrome(lnklist.New(10, lnklist.New(2, lnklist.New(5, lnklist.New(2, lnklist.New(10, nil))))))
	require.True(t, result)
}

func Test_Partition_NotPalindrome(t *testing.T) {
	result := isPalindrome(lnklist.New(10, lnklist.New(2, lnklist.New(5, lnklist.New(2, lnklist.New(1, nil))))))
	require.False(t, result)
}
