// This is a generated file - DO NOT EDIT

package cairo

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

/*

C type

cairo_t
*/
type Context struct {
	native *C.cairo_t
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.cairo_t)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_device_t
*/
type Device struct {
	native *C.cairo_device_t
}

func DeviceNewFromC(u unsafe.Pointer) *Device {
	c := (*C.cairo_device_t)(u)
	if c == nil {
		return nil
	}

	g := &Device{native: c}

	return g
}

func (recv *Device) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_surface_t
*/
type Surface struct {
	native *C.cairo_surface_t
}

func SurfaceNewFromC(u unsafe.Pointer) *Surface {
	c := (*C.cairo_surface_t)(u)
	if c == nil {
		return nil
	}

	g := &Surface{native: c}

	return g
}

func (recv *Surface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_matrix_t
*/
type Matrix struct {
	native *C.cairo_matrix_t
}

func MatrixNewFromC(u unsafe.Pointer) *Matrix {
	c := (*C.cairo_matrix_t)(u)
	if c == nil {
		return nil
	}

	g := &Matrix{native: c}

	return g
}

func (recv *Matrix) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_pattern_t
*/
type Pattern struct {
	native *C.cairo_pattern_t
}

func PatternNewFromC(u unsafe.Pointer) *Pattern {
	c := (*C.cairo_pattern_t)(u)
	if c == nil {
		return nil
	}

	g := &Pattern{native: c}

	return g
}

func (recv *Pattern) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_region_t
*/
type Region struct {
	native *C.cairo_region_t
}

func RegionNewFromC(u unsafe.Pointer) *Region {
	c := (*C.cairo_region_t)(u)
	if c == nil {
		return nil
	}

	g := &Region{native: c}

	return g
}

func (recv *Region) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_font_options_t
*/
type FontOptions struct {
	native *C.cairo_font_options_t
}

func FontOptionsNewFromC(u unsafe.Pointer) *FontOptions {
	c := (*C.cairo_font_options_t)(u)
	if c == nil {
		return nil
	}

	g := &FontOptions{native: c}

	return g
}

func (recv *FontOptions) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_font_face_t
*/
type FontFace struct {
	native *C.cairo_font_face_t
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	c := (*C.cairo_font_face_t)(u)
	if c == nil {
		return nil
	}

	g := &FontFace{native: c}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_scaled_font_t
*/
type ScaledFont struct {
	native *C.cairo_scaled_font_t
}

func ScaledFontNewFromC(u unsafe.Pointer) *ScaledFont {
	c := (*C.cairo_scaled_font_t)(u)
	if c == nil {
		return nil
	}

	g := &ScaledFont{native: c}

	return g
}

func (recv *ScaledFont) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_path_t
*/
type Path struct {
	native *C.cairo_path_t
}

func PathNewFromC(u unsafe.Pointer) *Path {
	c := (*C.cairo_path_t)(u)
	if c == nil {
		return nil
	}

	g := &Path{native: c}

	return g
}

func (recv *Path) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

cairo_rectangle_int_t
*/
type RectangleInt struct {
	native *C.cairo_rectangle_int_t
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleIntNewFromC(u unsafe.Pointer) *RectangleInt {
	c := (*C.cairo_rectangle_int_t)(u)
	if c == nil {
		return nil
	}

	g := &RectangleInt{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *RectangleInt) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}
