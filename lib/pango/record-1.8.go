// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Determines whether the size of the font is in points (not absolute) or device units (absolute).
// See pango_font_description_set_size() and pango_font_description_set_absolute_size().
/*

C function

pango_font_description_get_size_is_absolute
*/
func (recv *FontDescription) GetSizeIsAbsolute() bool {
	retC := C.pango_font_description_get_size_is_absolute((*C.PangoFontDescription)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the size field of a font description, in device units. This is mutually
// exclusive with pango_font_description_set_size() which sets the font size
// in points.
/*

C function

pango_font_description_set_absolute_size
*/
func (recv *FontDescription) SetAbsoluteSize(size float64) {
	c_size := (C.double)(size)

	C.pango_font_description_set_absolute_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// Class structure for #PangoRenderer.
/*

C type

PangoRendererClass
*/
type RendererClass struct {
	native *C.PangoRendererClass
	// Private : parent_class
	// no type for draw_glyphs
	// no type for draw_rectangle
	// no type for draw_error_underline
	// no type for draw_shape
	// no type for draw_trapezoid
	// no type for draw_glyph
	// no type for part_changed
	// no type for begin
	// no type for end
	// no type for prepare_run
	// no type for draw_glyph_item
	// no type for _pango_reserved2
	// no type for _pango_reserved3
	// no type for _pango_reserved4
}

func RendererClassNewFromC(u unsafe.Pointer) *RendererClass {
	c := (*C.PangoRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &RendererClass{native: c}

	return g
}

func (recv *RendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
