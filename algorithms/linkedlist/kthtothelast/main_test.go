package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthToTheLast_Zero(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 0)

	require.Nil(t, result)
}

func Test_kthToTheLast_SingleNode(t *testing.T) {
	first := &node{data: 1}

	result := kthToTheLast(first, 1)

	require.Equal(t, first, result)
}

func Test_kthToTheLast_TwoNode(t *testing.T) {
	second := &node{data: 2}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 1)

	require.Equal(t, second, result)
}

func Test_kthToTheLast_Normal(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 6)

	require.Equal(t, third, result)
}

func Test_kthToTheLast_ExceedLength(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 9)

	require.Nil(t, result)
}

func Test_kthToTheLast_FirstNode(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 8)

	require.Equal(t, first, result)
}

func Test_kthToTheLast_LastNode(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := kthToTheLast(first, 1)

	require.Equal(t, eighth, result)
}



func Test_nthToTheLast_Zero(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 0)

	require.Nil(t, result)
}

func Test_nthToTheLast_SingleNode(t *testing.T) {
	first := &node{data: 1}

	result := nthToTheLast(first, 1)

	require.Equal(t, first, result)
}

func Test_nthToTheLast_TwoNode(t *testing.T) {
	second := &node{data: 2}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 1)

	require.Equal(t, second, result)
}

func Test_nthToTheLast_Normal(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 6)

	require.Equal(t, third, result)
}

func Test_nthToTheLast_ExceedLength(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 9)

	require.Nil(t, result)
}

func Test_nthToTheLast_FirstNode(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 8)

	require.Equal(t, first, result)
}

func Test_nthToTheLast_LastNode(t *testing.T) {
	eighth := &node{data: 8}
	seventh := &node{data: 7, next: eighth}
	sixth := &node{data: 6, next: seventh}
	fifth := &node{data: 5, next: sixth}
	fourth := &node{data: 4, next: fifth}
	third := &node{data: 3, next: fourth}
	second := &node{data: 2, next: third}
	first := &node{data: 1, next: second}

	result := nthToTheLast(first, 1)

	require.Equal(t, eighth, result)
}
