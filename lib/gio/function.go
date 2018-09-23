// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// Unsupported : g_app_info_create_from_commandline : unsupported parameter flags : no type generator for AppInfoCreateFlags, GAppInfoCreateFlags

// AppInfoGetAll is a wrapper around the C function g_app_info_get_all.
func AppInfoGetAll() *glib.List {
	retC := C.g_app_info_get_all()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetAllForType is a wrapper around the C function g_app_info_get_all_for_type.
func AppInfoGetAllForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_all_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_app_info_get_default_for_type : unsupported parameter must_support_uris : no type generator for gboolean, gboolean

// Unsupported : g_app_info_get_default_for_uri_scheme : no return generator

// Unsupported : g_app_info_launch_default_for_uri : unsupported parameter context : no type generator for AppLaunchContext, GAppLaunchContext*

// Unsupported : g_content_type_can_be_executable : no return generator

// Unsupported : g_content_type_equals : no return generator

// ContentTypeGetDescription is a wrapper around the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_description(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_content_type_get_icon : no return generator

// ContentTypeGetMimeType is a wrapper around the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_mime_type(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_content_type_guess : unsupported parameter data : no param type

// Unsupported : g_content_type_is_a : no return generator

// Unsupported : g_content_type_is_unknown : no return generator

// ContentTypesGetRegistered is a wrapper around the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retC := C.g_content_types_get_registered()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DbusErrorQuark is a wrapper around the C function g_dbus_error_quark.
func DbusErrorQuark() glib.Quark {
	retC := C.g_dbus_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : g_file_new_for_commandline_arg : no return generator

// Unsupported : g_file_new_for_path : no return generator

// Unsupported : g_file_new_for_uri : no return generator

// Unsupported : g_file_parse_name : no return generator

// IconHash is a wrapper around the C function g_icon_hash.
func IconHash(icon uintptr) uint32 {
	c_icon := (C.gconstpointer)(icon)

	retC := C.g_icon_hash(c_icon)
	retGo := (uint32)(retC)

	return retGo
}

// IoErrorFromErrno is a wrapper around the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	c_err_no := (C.gint)(errNo)

	retC := C.g_io_error_from_errno(c_err_no)
	retGo := (IOErrorEnum)(retC)

	return retGo
}

// IoErrorQuark is a wrapper around the C function g_io_error_quark.
func IoErrorQuark() glib.Quark {
	retC := C.g_io_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : g_io_extension_point_implement : unsupported parameter type : no type generator for GType, GType

// IoExtensionPointLookup is a wrapper around the C function g_io_extension_point_lookup.
func IoExtensionPointLookup(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_lookup(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoExtensionPointRegister is a wrapper around the C function g_io_extension_point_register.
func IoExtensionPointRegister(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_register(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoModulesLoadAllInDirectory is a wrapper around the C function g_io_modules_load_all_in_directory.
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	retC := C.g_io_modules_load_all_in_directory(c_dirname)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_io_scheduler_cancel_all_jobs : no return generator

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc, GIOSchedulerJobFunc

// Unsupported : g_keyfile_settings_backend_new : no return generator

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter object : no type generator for GObject.Object, GObject*

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter object : no type generator for GObject.Object, GObject*

// Unsupported : g_unix_is_mount_path_system_internal : no return generator

// Unsupported : g_unix_mount_at : unsupported parameter time_read : no type generator for guint64, guint64*

// UnixMountCompare is a wrapper around the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	c_mount1 := mount1.toC()

	c_mount2 := mount2.toC()

	retC := C.g_unix_mount_compare(c_mount1, c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_unix_mount_free : no return generator

// UnixMountGetDevicePath is a wrapper around the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	c_mount_entry := mountEntry.toC()

	retC := C.g_unix_mount_get_device_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetFsType is a wrapper around the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	c_mount_entry := mountEntry.toC()

	retC := C.g_unix_mount_get_fs_type(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetMountPath is a wrapper around the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	c_mount_entry := mountEntry.toC()

	retC := C.g_unix_mount_get_mount_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_unix_mount_guess_can_eject : no return generator

// Unsupported : g_unix_mount_guess_icon : no return generator

// UnixMountGuessName is a wrapper around the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	c_mount_entry := mountEntry.toC()

	retC := C.g_unix_mount_guess_name(c_mount_entry)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_mount_guess_should_display : no return generator

// Unsupported : g_unix_mount_is_readonly : no return generator

// Unsupported : g_unix_mount_is_system_internal : no return generator

// Unsupported : g_unix_mount_points_changed_since : no return generator

// Unsupported : g_unix_mount_points_get : unsupported parameter time_read : no type generator for guint64, guint64*

// Unsupported : g_unix_mounts_changed_since : no return generator

// Unsupported : g_unix_mounts_get : unsupported parameter time_read : no type generator for guint64, guint64*