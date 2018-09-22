// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_str_is_ascii : no return generator

// Unsupported : g_str_match_string : unsupported parameter accept_alternates : no type generator for gboolean, gboolean

// StrToAscii is a wrapper around the C function g_str_to_ascii.
func StrToAscii(str string, fromLocale string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_from_locale := C.CString(fromLocale)
	defer C.free(unsafe.Pointer(c_from_locale))

	retC := C.g_str_to_ascii(c_str, c_from_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : no param type

// VariantParseErrorPrintContext is a wrapper around the C function g_variant_parse_error_print_context.
func VariantParseErrorPrintContext(error *Error, sourceStr string) string {
	c_error := error.toC()

	c_source_str := C.CString(sourceStr)
	defer C.free(unsafe.Pointer(c_source_str))

	retC := C.g_variant_parse_error_print_context(c_error, c_source_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
