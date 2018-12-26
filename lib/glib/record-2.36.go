// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Ref is a wrapper around the C function g_markup_parse_context_ref.
func (recv *MarkupParseContext) Ref() *MarkupParseContext {
	retC := C.g_markup_parse_context_ref((*C.GMarkupParseContext)(recv.native))
	retGo := MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_markup_parse_context_unref.
func (recv *MarkupParseContext) Unref() {
	C.g_markup_parse_context_unref((*C.GMarkupParseContext)(recv.native))

	return
}

// AddUnixFd is a wrapper around the C function g_source_add_unix_fd.
func (recv *Source) AddUnixFd(fd int32, events IOCondition) uintptr {
	c_fd := (C.gint)(fd)

	c_events := (C.GIOCondition)(events)

	retC := C.g_source_add_unix_fd((*C.GSource)(recv.native), c_fd, c_events)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// ModifyUnixFd is a wrapper around the C function g_source_modify_unix_fd.
func (recv *Source) ModifyUnixFd(tag uintptr, newEvents IOCondition) {
	c_tag := (C.gpointer)(tag)

	c_new_events := (C.GIOCondition)(newEvents)

	C.g_source_modify_unix_fd((*C.GSource)(recv.native), c_tag, c_new_events)

	return
}

// QueryUnixFd is a wrapper around the C function g_source_query_unix_fd.
func (recv *Source) QueryUnixFd(tag uintptr) IOCondition {
	c_tag := (C.gpointer)(tag)

	retC := C.g_source_query_unix_fd((*C.GSource)(recv.native), c_tag)
	retGo := (IOCondition)(retC)

	return retGo
}

// RemoveUnixFd is a wrapper around the C function g_source_remove_unix_fd.
func (recv *Source) RemoveUnixFd(tag uintptr) {
	c_tag := (C.gpointer)(tag)

	C.g_source_remove_unix_fd((*C.GSource)(recv.native), c_tag)

	return
}

// SetReadyTime is a wrapper around the C function g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) {
	c_ready_time := (C.gint64)(readyTime)

	C.g_source_set_ready_time((*C.GSource)(recv.native), c_ready_time)

	return
}

// VariantNewFromBytes is a wrapper around the C function g_variant_new_from_bytes.
func VariantNewFromBytes(type_ *VariantType, bytes *Bytes, trusted bool) *Variant {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	c_trusted :=
		boolToGboolean(trusted)

	retC := C.g_variant_new_from_bytes(c_type, c_bytes, c_trusted)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDataAsBytes is a wrapper around the C function g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() *Bytes {
	retC := C.g_variant_get_data_as_bytes((*C.GVariant)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}
