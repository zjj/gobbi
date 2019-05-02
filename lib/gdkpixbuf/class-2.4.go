// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	"unsafe"
)

// PixbufGetFileInfo is a wrapper around the C function gdk_pixbuf_get_file_info.
func PixbufGetFileInfo(filename string) (*PixbufFormat, int32, int32) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var c_width C.gint

	var c_height C.gint

	retC := C.gdk_pixbuf_get_file_info(c_filename, &c_width, &c_height)
	var retGo (*PixbufFormat)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufFormatNewFromC(unsafe.Pointer(retC))
	}

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func
