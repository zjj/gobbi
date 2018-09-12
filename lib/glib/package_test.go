package glib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, AsciiType(16), ASCII_GRAPH)
}

func TestFunctionCallWithIntegers(t *testing.T) {
	assert.Equal(t, int32(4), AsciiDigitValue('4'))
}

func TestFunctionCallWithStringsNotFreeReturnedString(t *testing.T) {
	assert.Equal(t, "edcba", Strreverse("abcde"))
}

func TestFunctionCallWithReturnedStringToFree(t *testing.T) {
	assert.Equal(t, "abcde", InternString("abcde"))
}
