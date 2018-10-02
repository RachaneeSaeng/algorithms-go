package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NodeLenght(t *testing.T) {
	eighth := &Node{Data: 8}
	seventh := &Node{Data: 7, Next: eighth}
	sixth := &Node{Data: 6, Next: seventh}
	fifth := &Node{Data: 5, Next: sixth}
	fourth := &Node{Data: 4, Next: fifth}
	third := &Node{Data: 3, Next: fourth}
	second := &Node{Data: 2, Next: third}
	first := &Node{Data: 1, Next: second}

	result := first.Length()
	require.Equal(t, 8, result)

	result = fifth.Length()
	require.Equal(t, 4, result)

	result = eighth.Length()
	require.Equal(t, 1, result)
}
