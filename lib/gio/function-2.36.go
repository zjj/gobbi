// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	"unsafe"
)

// DbusAddressEscapeValue is a wrapper around the C function g_dbus_address_escape_value.
func DbusAddressEscapeValue(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_address_escape_value(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// NetworkingInit is a wrapper around the C function g_networking_init.
func NetworkingInit() {
	C.g_networking_init()

	return
}
