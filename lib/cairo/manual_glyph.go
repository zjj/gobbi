// +build cairo_1.4 cairo_1.6 cairo cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

type Glyph struct {
	Index int
	X     float64
	Y     float64
}

func (g Glyph) toC() C.cairo_glyph_t {
	return C.cairo_glyph_t{
		index: C.ulong(g.Index),
		x:     C.double(g.X),
		y:     C.double(g.Y),
	}
}

func glyphsToC(glyphs []Glyph) (C.int, []C.cairo_glyph_t) {
	numGlyphs := len(glyphs)
	c_numGlyphs := C.int(numGlyphs)

	c_glyphs := make([]C.cairo_glyph_t, numGlyphs, numGlyphs)
	for i, glyph := range glyphs {
		c_glyphs[i] = glyph.toC()
	}

	return c_numGlyphs, c_glyphs
}
