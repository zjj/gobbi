// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Unsupported : g_io_extension_get_type : no return generator

// Unsupported : g_io_extension_point_get_required_type : no return generator

// Unsupported : g_io_extension_point_set_required_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_resource_enumerate_children : no return type

// Unsupported : g_resource_get_info : unsupported parameter size : no type generator for gsize, gsize*

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// GetKey is a wrapper around the C function g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_get_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasKey is a wrapper around the C function g_settings_schema_has_key.
func (recv *SettingsSchema) HasKey(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_has_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_schema_list_children : no return type

// Unsupported : g_settings_schema_list_keys : no return type

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Ref is a wrapper around the C function g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	retC := C.g_settings_schema_key_ref((*C.GSettingsSchemaKey)(recv.native))
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(recv.native))

	return
}

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : no param type

// Unsupported : g_unix_mount_point_guess_icon : no return generator

// Unsupported : g_unix_mount_point_guess_symbolic_icon : no return generator