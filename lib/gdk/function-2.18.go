// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Gets the window that @window is embedded in.
/*

C function

gdk_offscreen_window_get_embedder
*/
func OffscreenWindowGetEmbedder(window *Window) *Window {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_offscreen_window_get_embedder(c_window)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets @window to be embedded in @embedder.
//
// To fully embed an offscreen window, in addition to calling this
// function, it is also necessary to handle the #GdkWindow::pick-embedded-child
// signal on the @embedder and the #GdkWindow::to-embedder and
// #GdkWindow::from-embedder signals on @window.
/*

C function

gdk_offscreen_window_set_embedder
*/
func OffscreenWindowSetEmbedder(window *Window, embedder *Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_embedder := (*C.GdkWindow)(C.NULL)
	if embedder != nil {
		c_embedder = (*C.GdkWindow)(embedder.ToC())
	}

	C.gdk_offscreen_window_set_embedder(c_window, c_embedder)

	return
}
