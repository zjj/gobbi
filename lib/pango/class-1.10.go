// This is a generated file - DO NOT EDIT
// +build pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontMap is a wrapper around the C function pango_font_get_font_map.
func (recv *Font) GetFontMap() *FontMap {
	retC := C.pango_font_get_font_map((*C.PangoFont)(recv.native))
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}
