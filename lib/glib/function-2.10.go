// This is a generated file - DO NOT EDIT
// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// InternStaticString is a wrapper around the C function g_intern_static_string.
func InternStaticString(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_static_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// InternString is a wrapper around the C function g_intern_string.
func InternString(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_slice_alloc : no return generator

// Unsupported : g_slice_alloc0 : no return generator

// Unsupported : g_slice_free1 : unsupported parameter mem_block : no type generator for gpointer (gpointer) for param mem_block

// Unsupported : g_slice_free_chain_with_offset : unsupported parameter mem_chain : no type generator for gpointer (gpointer) for param mem_chain
