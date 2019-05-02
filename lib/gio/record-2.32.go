// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// Equals compares this ActionMapInterface with another ActionMapInterface, and returns true if they represent the same GObject.
func (recv *ActionMapInterface) Equals(other *ActionMapInterface) bool {
	return other.ToC() == recv.ToC()
}

// ToString is a wrapper around the C function g_file_attribute_matcher_to_string.
func (recv *FileAttributeMatcher) ToString() string {
	retC := C.g_file_attribute_matcher_to_string((*C.GFileAttributeMatcher)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Equals compares this NetworkMonitorInterface with another NetworkMonitorInterface, and returns true if they represent the same GObject.
func (recv *NetworkMonitorInterface) Equals(other *NetworkMonitorInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this RemoteActionGroupInterface with another RemoteActionGroupInterface, and returns true if they represent the same GObject.
func (recv *RemoteActionGroupInterface) Equals(other *RemoteActionGroupInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Resource with another Resource, and returns true if they represent the same GObject.
func (recv *Resource) Equals(other *Resource) bool {
	return other.ToC() == recv.ToC()
}

// ResourceLoad is a wrapper around the C function g_resource_load.
func ResourceLoad(filename string) (*Resource, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_resource_load(c_filename, &cThrowableError)
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EnumerateChildren is a wrapper around the C function g_resource_enumerate_children.
func (recv *Resource) EnumerateChildren(path string, lookupFlags ResourceLookupFlags) ([]string, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resource_enumerate_children((*C.GResource)(recv.native), c_path, c_lookup_flags, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetInfo is a wrapper around the C function g_resource_get_info.
func (recv *Resource) GetInfo(path string, lookupFlags ResourceLookupFlags) (bool, uint64, uint32, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var c_size C.gsize

	var c_flags C.guint32

	var cThrowableError *C.GError

	retC := C.g_resource_get_info((*C.GResource)(recv.native), c_path, c_lookup_flags, &c_size, &c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	size := (uint64)(c_size)

	flags := (uint32)(c_flags)

	return retGo, size, flags, goError
}

// LookupData is a wrapper around the C function g_resource_lookup_data.
func (recv *Resource) LookupData(path string, lookupFlags ResourceLookupFlags) (*glib.Bytes, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resource_lookup_data((*C.GResource)(recv.native), c_path, c_lookup_flags, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// OpenStream is a wrapper around the C function g_resource_open_stream.
func (recv *Resource) OpenStream(path string, lookupFlags ResourceLookupFlags) (*InputStream, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resource_open_stream((*C.GResource)(recv.native), c_path, c_lookup_flags, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Ref is a wrapper around the C function g_resource_ref.
func (recv *Resource) Ref() *Resource {
	retC := C.g_resource_ref((*C.GResource)(recv.native))
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_resource_unref.
func (recv *Resource) Unref() {
	C.g_resource_unref((*C.GResource)(recv.native))

	return
}

// Equals compares this SettingsSchema with another SettingsSchema, and returns true if they represent the same GObject.
func (recv *SettingsSchema) Equals(other *SettingsSchema) bool {
	return other.ToC() == recv.ToC()
}

// GetId is a wrapper around the C function g_settings_schema_get_id.
func (recv *SettingsSchema) GetId() string {
	retC := C.g_settings_schema_get_id((*C.GSettingsSchema)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPath is a wrapper around the C function g_settings_schema_get_path.
func (recv *SettingsSchema) GetPath() string {
	retC := C.g_settings_schema_get_path((*C.GSettingsSchema)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Ref is a wrapper around the C function g_settings_schema_ref.
func (recv *SettingsSchema) Ref() *SettingsSchema {
	retC := C.g_settings_schema_ref((*C.GSettingsSchema)(recv.native))
	retGo := SettingsSchemaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_settings_schema_unref.
func (recv *SettingsSchema) Unref() {
	C.g_settings_schema_unref((*C.GSettingsSchema)(recv.native))

	return
}

// Equals compares this SettingsSchemaSource with another SettingsSchemaSource, and returns true if they represent the same GObject.
func (recv *SettingsSchemaSource) Equals(other *SettingsSchemaSource) bool {
	return other.ToC() == recv.ToC()
}

// SettingsSchemaSourceGetDefault is a wrapper around the C function g_settings_schema_source_get_default.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_get_default()
	var retGo (*SettingsSchemaSource)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Lookup is a wrapper around the C function g_settings_schema_source_lookup.
func (recv *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_recursive :=
		boolToGboolean(recursive)

	retC := C.g_settings_schema_source_lookup((*C.GSettingsSchemaSource)(recv.native), c_schema_id, c_recursive)
	var retGo (*SettingsSchema)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SettingsSchemaNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Ref is a wrapper around the C function g_settings_schema_source_ref.
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_ref((*C.GSettingsSchemaSource)(recv.native))
	retGo := SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_settings_schema_source_unref.
func (recv *SettingsSchemaSource) Unref() {
	C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(recv.native))

	return
}

// Fini is a wrapper around the C function g_static_resource_fini.
func (recv *StaticResource) Fini() {
	C.g_static_resource_fini((*C.GStaticResource)(recv.native))

	return
}

// GetResource is a wrapper around the C function g_static_resource_get_resource.
func (recv *StaticResource) GetResource() *Resource {
	retC := C.g_static_resource_get_resource((*C.GStaticResource)(recv.native))
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_static_resource_init.
func (recv *StaticResource) Init() {
	C.g_static_resource_init((*C.GStaticResource)(recv.native))

	return
}

// GetOptions is a wrapper around the C function g_unix_mount_point_get_options.
func (recv *UnixMountPoint) GetOptions() string {
	retC := C.g_unix_mount_point_get_options((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
