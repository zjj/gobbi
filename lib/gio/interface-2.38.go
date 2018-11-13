// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_file_make_directory_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous directory creation, started with
// g_file_make_directory_async().
/*

C function

g_file_make_directory_finish
*/
func (recv *File) MakeDirectoryFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_make_directory_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_measure_disk_usage : unsupported parameter progress_callback : no type generator for FileMeasureProgressCallback (GFileMeasureProgressCallback) for param progress_callback

// Unsupported : g_file_measure_disk_usage_async : unsupported parameter progress_callback : no type generator for FileMeasureProgressCallback (GFileMeasureProgressCallback) for param progress_callback

// Collects the results from an earlier call to
// g_file_measure_disk_usage_async().  See g_file_measure_disk_usage() for
// more information.
/*

C function

g_file_measure_disk_usage_finish
*/
func (recv *File) MeasureDiskUsageFinish(result *AsyncResult) (bool, uint64, uint64, uint64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_disk_usage C.guint64

	var c_num_dirs C.guint64

	var c_num_files C.guint64

	var cThrowableError *C.GError

	retC := C.g_file_measure_disk_usage_finish((*C.GFile)(recv.native), c_result, &c_disk_usage, &c_num_dirs, &c_num_files, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	diskUsage := (uint64)(c_disk_usage)

	numDirs := (uint64)(c_num_dirs)

	numFiles := (uint64)(c_num_files)

	return retGo, diskUsage, numDirs, numFiles, goThrowableError
}

// Unsupported : g_file_trash_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file trashing operation, started with
// g_file_trash_async().
/*

C function

g_file_trash_finish
*/
func (recv *File) TrashFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_trash_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_icon_serialize : return type : Blacklisted record : GVariant
