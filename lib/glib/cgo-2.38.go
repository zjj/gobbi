// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static GVariant* _g_variant_new_printf(const gchar* format_string) {
		return g_variant_new_printf(format_string);
    }
*/
import "C"

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// VariantNewPrintf is a wrapper around the C function g_variant_new_printf.
func VariantNewPrintf(formatString string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_new_printf(c_format_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewTakeString is a wrapper around the C function g_variant_new_take_string.
func VariantNewTakeString(string_ string) *Variant {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_new_take_string(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}
