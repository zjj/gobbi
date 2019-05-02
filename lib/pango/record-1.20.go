// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	"C"
	"unsafe"
)

// Init is a wrapper around the C function pango_attribute_init.
func (recv *Attribute) Init(klass *AttrClass) {
	c_klass := (*C.PangoAttrClass)(C.NULL)
	if klass != nil {
		c_klass = (*C.PangoAttrClass)(klass.ToC())
	}

	C.pango_attribute_init((*C.PangoAttribute)(recv.native), c_klass)

	return
}

// Copy is a wrapper around the C function pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	retC := C.pango_glyph_item_copy((*C.PangoGlyphItem)(recv.native))
	var retGo (*GlyphItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GlyphItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
