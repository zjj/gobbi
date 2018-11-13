// This is a generated file - DO NOT EDIT
// +build gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns a textual specification of @color in the hexadecimal
// form “\#rrrrggggbbbb” where “r”, “g” and “b” are hex digits
// representing the red, green and blue components respectively.
//
// The returned string can be parsed by gdk_color_parse().
/*

C function : gdk_color_to_string
*/
func (recv *Color) ToString() string {
	retC := C.gdk_color_to_string((*C.GdkColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
