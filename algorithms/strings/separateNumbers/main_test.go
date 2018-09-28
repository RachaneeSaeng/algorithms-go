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

func Test_LongDigit_1(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("429496729542949672964294967297")
	})

	require.Equal(t, "YES 4294967295\n", output)
}

func Test_LongDigit_2(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("11111111111111111111111111111112")
	})

	require.Equal(t, "YES 1111111111111111\n", output)
}

func Test_LongDigit_3(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("11111111111111111111111111111111")
	})

	require.Equal(t, "NO\n", output)
}

func Test_LongDigit_4(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("26272829303132333435363738384041")
	})

	require.Equal(t, "NO\n", output)
}

func Test_LongDigit_5(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("32333435363738394041424344454647")
	})

	require.Equal(t, "YES 32\n", output)
}

func Test_LongDigit_6(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("32333435363738394041424344354647")
	})

	require.Equal(t, "NO\n", output)
}

func Test_LongDigit_7(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("888990919293949596979899100101")
	})

	require.Equal(t, "YES 88\n", output)
}

func Test_LongDigit_8(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("888990919293949596979898100101")
	})

	require.Equal(t, "NO\n", output)
}

func Test_LongDigit_9(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("16171819202122232425262728293031")
	})

	require.Equal(t, "YES 16\n", output)
}

func Test_LongDigit_10(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("16171819202122232425262718293031")
	})

	require.Equal(t, "NO\n", output)
}

func Test_ChangeToNextLevel_1(t *testing.T) {
	output := capturer.CaptureStdout(func() {
		separateNumbers("9101112131415161718192021222324")
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
