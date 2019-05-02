// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// GetSupportedTypes is a wrapper around the C function g_app_info_get_supported_types.
func (recv *AppInfo) GetSupportedTypes() []string {
	retC := C.g_app_info_get_supported_types((*C.GAppInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// Unsupported : g_async_result_is_tagged : unsupported parameter source_tag : no type generator for gpointer (gpointer) for param source_tag

// LegacyPropagateError is a wrapper around the C function g_async_result_legacy_propagate_error.
func (recv *AsyncResult) LegacyPropagateError() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_async_result_legacy_propagate_error((*C.GAsyncResult)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSymbolicIcon is a wrapper around the C function g_drive_get_symbolic_icon.
func (recv *Drive) GetSymbolicIcon() *Icon {
	retC := C.g_drive_get_symbolic_icon((*C.GDrive)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_delete_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// DeleteFinish is a wrapper around the C function g_file_delete_finish.
func (recv *File) DeleteFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_delete_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSymbolicIcon is a wrapper around the C function g_mount_get_symbolic_icon.
func (recv *Mount) GetSymbolicIcon() *Icon {
	retC := C.g_mount_get_symbolic_icon((*C.GMount)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSymbolicIcon is a wrapper around the C function g_volume_get_symbolic_icon.
func (recv *Volume) GetSymbolicIcon() *Icon {
	retC := C.g_volume_get_symbolic_icon((*C.GVolume)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
