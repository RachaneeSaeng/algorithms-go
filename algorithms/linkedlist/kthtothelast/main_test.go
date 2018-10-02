package main

import (
	"testing"

	"../lnklist"

	"github.com/stretchr/testify/require"
)

func Test_kthToTheLast_Zero(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 0)

	require.Nil(t, result)
}

func Test_kthToTheLast_SingleNode(t *testing.T) {
	first := &lnklist.Node{Data: 1}

	result := kthToTheLast(first, 1)

	require.Equal(t, first, result)
}

func Test_kthToTheLast_TwoNode(t *testing.T) {
	second := &lnklist.Node{Data: 2}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 1)

	require.Equal(t, second, result)
}

func Test_kthToTheLast_Normal(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 6)

	require.Equal(t, third, result)
}

func Test_kthToTheLast_ExceedLength(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 9)

	require.Nil(t, result)
}

func Test_kthToTheLast_FirstNode(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 8)

	require.Equal(t, first, result)
}

func Test_kthToTheLast_LastNode(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := kthToTheLast(first, 1)

	require.Equal(t, eighth, result)
}

func Test_nthToTheLast_Zero(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 0)

	require.Nil(t, result)
}

func Test_nthToTheLast_SingleNode(t *testing.T) {
	first := &lnklist.Node{Data: 1}

	result := nthToTheLast(first, 1)

	require.Equal(t, first, result)
}

func Test_nthToTheLast_TwoNode(t *testing.T) {
	second := &lnklist.Node{Data: 2}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 1)

	require.Equal(t, second, result)
}

func Test_nthToTheLast_Normal(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 6)

	require.Equal(t, third, result)
}

func Test_nthToTheLast_ExceedLength(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 9)

	require.Nil(t, result)
}

func Test_nthToTheLast_FirstNode(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 8)

	require.Equal(t, first, result)
}

func Test_nthToTheLast_LastNode(t *testing.T) {
	eighth := &lnklist.Node{Data: 8}
	seventh := &lnklist.Node{Data: 7, Next: eighth}
	sixth := &lnklist.Node{Data: 6, Next: seventh}
	fifth := &lnklist.Node{Data: 5, Next: sixth}
	fourth := &lnklist.Node{Data: 4, Next: fifth}
	third := &lnklist.Node{Data: 3, Next: fourth}
	second := &lnklist.Node{Data: 2, Next: third}
	first := &lnklist.Node{Data: 1, Next: second}

	result := nthToTheLast(first, 1)

	require.Equal(t, eighth, result)
}
