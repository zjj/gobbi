// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	"unsafe"
)

// Copy is a wrapper around the C function gdk_pixbuf_format_copy.
func (recv *PixbufFormat) Copy() *PixbufFormat {
	retC := C.gdk_pixbuf_format_copy((*C.GdkPixbufFormat)(recv.native))
	retGo := PixbufFormatNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gdk_pixbuf_format_free.
func (recv *PixbufFormat) Free() {
	C.gdk_pixbuf_format_free((*C.GdkPixbufFormat)(recv.native))

	return
}
