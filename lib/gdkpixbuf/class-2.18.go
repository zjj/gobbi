// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Gets whether @animation should loop indefinitely when it reaches the end.
/*

C function : gdk_pixbuf_simple_anim_get_loop
*/
func (recv *PixbufSimpleAnim) GetLoop() bool {
	retC := C.gdk_pixbuf_simple_anim_get_loop((*C.GdkPixbufSimpleAnim)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether @animation should loop indefinitely when it reaches the end.
/*

C function : gdk_pixbuf_simple_anim_set_loop
*/
func (recv *PixbufSimpleAnim) SetLoop(loop bool) {
	c_loop :=
		boolToGboolean(loop)

	C.gdk_pixbuf_simple_anim_set_loop((*C.GdkPixbufSimpleAnim)(recv.native), c_loop)

	return
}
