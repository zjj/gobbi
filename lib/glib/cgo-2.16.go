// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_compute_checksum_for_data : return type :

// Unsupported : g_compute_checksum_for_string : return type :

// Unsupported : g_dpgettext : return type :

// Unsupported : g_format_size_for_display : return type :

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// Unsupported : g_test_create_suite : return type :

// Unsupported : g_test_get_root : return type :

// Unsupported : g_test_init : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Unsupported : g_test_queue_free : unsupported parameter gfree_pointer : no type generator for gpointer (gpointer) for param gfree_pointer

// Unsupported : g_test_trap_fork : return type :

// Unsupported : g_test_trap_has_passed : return type :

// Unsupported : g_test_trap_reached_timeout : return type :

// Unsupported : g_uri_escape_string : return type :

// Unsupported : g_uri_parse_scheme : return type :

// Unsupported : g_uri_unescape_segment : return type :

// Unsupported : g_uri_unescape_string : return type :

// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func ChecksumNewFromC(u unsafe.Pointer) *Checksum {
	c := (*C.GChecksum)(u)
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}