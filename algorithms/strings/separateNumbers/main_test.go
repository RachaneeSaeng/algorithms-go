package main

import (
	"testing"

	capturer "github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func Test_EmptyString(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("")
	})

	require.Equal(t, "NO\n", output)
}

func Test_SingleCharacter(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("5")
	})

	require.Equal(t, "NO\n", output)
}

func Test_SingleDigit(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("5678")
	})

	require.Equal(t, "YES 5\n", output)
}

func Test_MultiDigit(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("939495")
	})

	require.Equal(t, "YES 93\n", output)
}

func Test_ChangeToNextLevel_1(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("91011")
	})

	require.Equal(t, "YES 9\n", output)
}

func Test_ChangeToNextLevel_2(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("99100101")
	})

	require.Equal(t, "YES 99\n", output)
}

func Test_LeadingZero(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("010203")
	})

	require.Equal(t, "NO\n", output)
}

func Test_NoSequenc_1(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("13")
	})

	require.Equal(t, "NO\n", output)
}

func Test_NoSequenc_2(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("101103")
	})

	require.Equal(t, "NO\n", output)
}

func Test_NoSequenc_3(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("93100")
	})

	require.Equal(t, "NO\n", output)
}
