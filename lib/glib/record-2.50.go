// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// LogField is a wrapper around the C record GLogField.
type LogField struct {
	native *C.GLogField
	Key    string
	Value  uintptr
	Length int64
}

func logFieldNewFromC(c *C.GLogField) *LogField {
	if c == nil {
		return nil
	}

	g := &LogField{
		Key:    C.GoString(c.key),
		Length: (int64)(c.length),
		Value:  (uintptr)(c.value),
		native: c,
	}

	return g
}

func (recv *LogField) toC() *C.GLogField {
	// TODO marshall fields to native

	return recv.native
}
