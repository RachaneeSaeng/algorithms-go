package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1(t *testing.T) {
	result := oneEditAway("","")
	require.Equal(t, true, result)
}

func Test_2(t *testing.T) {
	result := oneEditAway("pale","pale")
	require.Equal(t, true, result)
}

func Test_3(t *testing.T) {
	result := oneEditAway("pales","pale")
	require.Equal(t, true, result)
}

func Test_4(t *testing.T) {
	result := oneEditAway("pale","bale")
	require.Equal(t, true, result)
}

func Test_5(t *testing.T) {
	result := oneEditAway("pale","bae")
	require.Equal(t, false, result)
}

func Test_6(t *testing.T) {
	result := oneEditAway("pale","pale")
	require.Equal(t, true, result)
}