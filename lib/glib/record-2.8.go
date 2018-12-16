// This is a generated file - DO NOT EDIT
// +build glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// MappedFileNew is a wrapper around the C function g_mapped_file_new.
func MappedFileNew(filename string, writable bool) (*MappedFile, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new(c_filename, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_mapped_file_free.
func (recv *MappedFile) Free() {
	C.g_mapped_file_free((*C.GMappedFile)(recv.native))

	return
}

// GetContents is a wrapper around the C function g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() string {
	retC := C.g_mapped_file_get_contents((*C.GMappedFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function g_mapped_file_get_length.
func (recv *MappedFile) GetLength() uint64 {
	retC := C.g_mapped_file_get_length((*C.GMappedFile)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}
