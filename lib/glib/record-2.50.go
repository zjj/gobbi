// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// LogField is a wrapper around the C record GLogField.
type LogField struct {
	native unsafe.Pointer
	Key    string
	// value : no type generator for gpointer, gconstpointer
	Length int64
}
