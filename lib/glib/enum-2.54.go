// This is a generated file - DO NOT EDIT
// +build glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

type NumberParserError C.GNumberParserError

const (
	NUMBER_PARSER_ERROR_INVALID       NumberParserError = 0
	NUMBER_PARSER_ERROR_OUT_OF_BOUNDS NumberParserError = 1
)

// Blacklisted : GTestResult
