// This is a generated file - DO NOT EDIT
// +build glib_2.2 glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Retrieves the name of the currently open element.
//
// If called from the start_element or end_element handlers this will
// give the element_name as passed to those functions. For the parent
// elements, see g_markup_parse_context_get_element_stack().
/*

C function : g_markup_parse_context_get_element
*/
func (recv *MarkupParseContext) GetElement() string {
	retC := C.g_markup_parse_context_get_element((*C.GMarkupParseContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
