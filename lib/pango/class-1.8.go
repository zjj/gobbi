// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_context_list_families : unsupported parameter families : no param type

// Blacklisted : PangoEngine

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : no param type

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Blacklisted : PangoFontsetSimple

// GetFontDescription is a wrapper around the C function pango_layout_get_font_description.
func (recv *Layout) GetFontDescription() *FontDescription {
	retC := C.pango_layout_get_font_description((*C.PangoLayout)(recv.native))
	retGo := FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : no param type

// Unsupported : pango_layout_get_log_attrs_readonly : unsupported parameter n_attrs : no type generator for gint, gint*

// Unsupported : pango_layout_get_pixel_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_get_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_index_to_line_x : unsupported parameter line : no type generator for gint, int*

// Unsupported : pango_layout_move_cursor_visually : unsupported parameter new_index : no type generator for gint, int*

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*

// Renderer is a wrapper around the C record PangoRenderer.
type Renderer struct {
	native *C.PangoRenderer
	// Private : parent_instance
	// Private : underline
	// Private : strikethrough
	// Private : active_count
	// matrix : record
	// Private : priv
}

func RendererNewFromC(u unsafe.Pointer) *Renderer {
	c := (*C.PangoRenderer)(u)
	if c == nil {
		return nil
	}

	g := &Renderer{native: c}

	return g
}

func (recv *Renderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Activate is a wrapper around the C function pango_renderer_activate.
func (recv *Renderer) Activate() {
	C.pango_renderer_activate((*C.PangoRenderer)(recv.native))

	return
}

// Deactivate is a wrapper around the C function pango_renderer_deactivate.
func (recv *Renderer) Deactivate() {
	C.pango_renderer_deactivate((*C.PangoRenderer)(recv.native))

	return
}

// DrawErrorUnderline is a wrapper around the C function pango_renderer_draw_error_underline.
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_error_underline((*C.PangoRenderer)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// DrawGlyph is a wrapper around the C function pango_renderer_draw_glyph.
func (recv *Renderer) DrawGlyph(font *Font, glyph Glyph, x float64, y float64) {
	c_font := (*C.PangoFont)(font.ToC())

	c_glyph := (C.PangoGlyph)(glyph)

	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_renderer_draw_glyph((*C.PangoRenderer)(recv.native), c_font, c_glyph, c_x, c_y)

	return
}

// DrawGlyphs is a wrapper around the C function pango_renderer_draw_glyphs.
func (recv *Renderer) DrawGlyphs(font *Font, glyphs *GlyphString, x int32, y int32) {
	c_font := (*C.PangoFont)(font.ToC())

	c_glyphs := (*C.PangoGlyphString)(glyphs.ToC())

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_glyphs((*C.PangoRenderer)(recv.native), c_font, c_glyphs, c_x, c_y)

	return
}

// DrawLayout is a wrapper around the C function pango_renderer_draw_layout.
func (recv *Renderer) DrawLayout(layout *Layout, x int32, y int32) {
	c_layout := (*C.PangoLayout)(layout.ToC())

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout((*C.PangoRenderer)(recv.native), c_layout, c_x, c_y)

	return
}

// DrawLayoutLine is a wrapper around the C function pango_renderer_draw_layout_line.
func (recv *Renderer) DrawLayoutLine(line *LayoutLine, x int32, y int32) {
	c_line := (*C.PangoLayoutLine)(line.ToC())

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout_line((*C.PangoRenderer)(recv.native), c_line, c_x, c_y)

	return
}

// DrawRectangle is a wrapper around the C function pango_renderer_draw_rectangle.
func (recv *Renderer) DrawRectangle(part RenderPart, x int32, y int32, width int32, height int32) {
	c_part := (C.PangoRenderPart)(part)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_rectangle((*C.PangoRenderer)(recv.native), c_part, c_x, c_y, c_width, c_height)

	return
}

// DrawTrapezoid is a wrapper around the C function pango_renderer_draw_trapezoid.
func (recv *Renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	c_part := (C.PangoRenderPart)(part)

	c_y1_ := (C.double)(y1)

	c_x11 := (C.double)(x11)

	c_x21 := (C.double)(x21)

	c_y2 := (C.double)(y2)

	c_x12 := (C.double)(x12)

	c_x22 := (C.double)(x22)

	C.pango_renderer_draw_trapezoid((*C.PangoRenderer)(recv.native), c_part, c_y1_, c_x11, c_x21, c_y2, c_x12, c_x22)

	return
}

// GetColor is a wrapper around the C function pango_renderer_get_color.
func (recv *Renderer) GetColor(part RenderPart) *Color {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_color((*C.PangoRenderer)(recv.native), c_part)
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMatrix is a wrapper around the C function pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	retC := C.pango_renderer_get_matrix((*C.PangoRenderer)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PartChanged is a wrapper around the C function pango_renderer_part_changed.
func (recv *Renderer) PartChanged(part RenderPart) {
	c_part := (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed((*C.PangoRenderer)(recv.native), c_part)

	return
}

// SetColor is a wrapper around the C function pango_renderer_set_color.
func (recv *Renderer) SetColor(part RenderPart, color *Color) {
	c_part := (C.PangoRenderPart)(part)

	c_color := (*C.PangoColor)(color.ToC())

	C.pango_renderer_set_color((*C.PangoRenderer)(recv.native), c_part, c_color)

	return
}

// SetMatrix is a wrapper around the C function pango_renderer_set_matrix.
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(matrix.ToC())

	C.pango_renderer_set_matrix((*C.PangoRenderer)(recv.native), c_matrix)

	return
}