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

// AppLaunchContext is a wrapper around the C record GAppLaunchContext.
type AppLaunchContext struct {
	native *C.GAppLaunchContext
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	c := (*C.GAppLaunchContext)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContext{native: c}

	return g
}

func (recv *AppLaunchContext) toC() *C.GAppLaunchContext {

	return recv.native
}

// Unsupported : g_app_launch_context_new : no return generator

// Unsupported : g_app_launch_context_get_display : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Unsupported : g_app_launch_context_get_environment : no return type

// Unsupported : g_app_launch_context_get_startup_notify_id : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Unsupported : g_app_launch_context_launch_failed : no return generator

// Unsupported : g_app_launch_context_setenv : no return generator

// Unsupported : g_app_launch_context_unsetenv : no return generator

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native *C.GApplicationCommandLine
	// Private : parent_instance
	// Private : priv
}

func ApplicationCommandLineNewFromC(u unsafe.Pointer) *ApplicationCommandLine {
	c := (*C.GApplicationCommandLine)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLine{native: c}

	return g
}

func (recv *ApplicationCommandLine) toC() *C.GApplicationCommandLine {

	return recv.native
}

// Unsupported : g_application_command_line_create_file_for_arg : no return generator

// Unsupported : g_application_command_line_get_arguments : unsupported parameter argc : no type generator for gint, int*

// Unsupported : g_application_command_line_get_environ : no return type

// Unsupported : g_application_command_line_get_is_remote : no return generator

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// Unsupported : g_application_command_line_get_stdin : no return generator

// Unsupported : g_application_command_line_print : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_printerr : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_set_exit_status : no return generator

// BufferedInputStream is a wrapper around the C record GBufferedInputStream.
type BufferedInputStream struct {
	native *C.GBufferedInputStream
	// parent_instance : no type generator for FilterInputStream, GFilterInputStream
	// Private : priv
}

func BufferedInputStreamNewFromC(u unsafe.Pointer) *BufferedInputStream {
	c := (*C.GBufferedInputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStream{native: c}

	return g
}

func (recv *BufferedInputStream) toC() *C.GBufferedInputStream {

	return recv.native
}

// Unsupported : g_buffered_input_stream_new : unsupported parameter base_stream : no type generator for InputStream, GInputStream*

// Unsupported : g_buffered_input_stream_new_sized : unsupported parameter base_stream : no type generator for InputStream, GInputStream*

// Unsupported : g_buffered_input_stream_fill : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_buffered_input_stream_fill_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_buffered_input_stream_fill_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetAvailable is a wrapper around the C function g_buffered_input_stream_get_available.
func (recv *BufferedInputStream) GetAvailable() uint64 {
	retC := C.g_buffered_input_stream_get_available((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetBufferSize is a wrapper around the C function g_buffered_input_stream_get_buffer_size.
func (recv *BufferedInputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_input_stream_get_buffer_size((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_buffered_input_stream_peek : unsupported parameter buffer : no param type

// Unsupported : g_buffered_input_stream_peek_buffer : unsupported parameter count : no type generator for gsize, gsize*

// Unsupported : g_buffered_input_stream_read_byte : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_buffered_input_stream_set_buffer_size : no return generator

// BufferedOutputStream is a wrapper around the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native *C.GBufferedOutputStream
	// parent_instance : no type generator for FilterOutputStream, GFilterOutputStream
	// priv : record
}

func BufferedOutputStreamNewFromC(u unsafe.Pointer) *BufferedOutputStream {
	c := (*C.GBufferedOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStream{native: c}

	return g
}

func (recv *BufferedOutputStream) toC() *C.GBufferedOutputStream {

	return recv.native
}

// Unsupported : g_buffered_output_stream_new : unsupported parameter base_stream : no type generator for OutputStream, GOutputStream*

// Unsupported : g_buffered_output_stream_new_sized : unsupported parameter base_stream : no type generator for OutputStream, GOutputStream*

// Unsupported : g_buffered_output_stream_get_auto_grow : no return generator

// GetBufferSize is a wrapper around the C function g_buffered_output_stream_get_buffer_size.
func (recv *BufferedOutputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_output_stream_get_buffer_size((*C.GBufferedOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_buffered_output_stream_set_auto_grow : unsupported parameter auto_grow : no type generator for gboolean, gboolean

// Unsupported : g_buffered_output_stream_set_buffer_size : no return generator

// BytesIcon is a wrapper around the C record GBytesIcon.
type BytesIcon struct {
	native *C.GBytesIcon
}

func BytesIconNewFromC(u unsafe.Pointer) *BytesIcon {
	c := (*C.GBytesIcon)(u)
	if c == nil {
		return nil
	}

	g := &BytesIcon{native: c}

	return g
}

func (recv *BytesIcon) toC() *C.GBytesIcon {

	return recv.native
}

// Unsupported : g_bytes_icon_new : no return generator

// Cancellable is a wrapper around the C record GCancellable.
type Cancellable struct {
	native *C.GCancellable
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func CancellableNewFromC(u unsafe.Pointer) *Cancellable {
	c := (*C.GCancellable)(u)
	if c == nil {
		return nil
	}

	g := &Cancellable{native: c}

	return g
}

func (recv *Cancellable) toC() *C.GCancellable {

	return recv.native
}

// Unsupported : g_cancellable_new : no return generator

// Unsupported : g_cancellable_cancel : no return generator

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback, GCallback

// Unsupported : g_cancellable_disconnect : no return generator

// GetFd is a wrapper around the C function g_cancellable_get_fd.
func (recv *Cancellable) GetFd() int32 {
	retC := C.g_cancellable_get_fd((*C.GCancellable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_cancellable_is_cancelled : no return generator

// Unsupported : g_cancellable_make_pollfd : no return generator

// Unsupported : g_cancellable_pop_current : no return generator

// Unsupported : g_cancellable_push_current : no return generator

// Unsupported : g_cancellable_release_fd : no return generator

// Unsupported : g_cancellable_reset : no return generator

// Unsupported : g_cancellable_set_error_if_cancelled : no return generator

// CharsetConverter is a wrapper around the C record GCharsetConverter.
type CharsetConverter struct {
	native *C.GCharsetConverter
}

func CharsetConverterNewFromC(u unsafe.Pointer) *CharsetConverter {
	c := (*C.GCharsetConverter)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverter{native: c}

	return g
}

func (recv *CharsetConverter) toC() *C.GCharsetConverter {

	return recv.native
}

// Unsupported : g_charset_converter_new : no return generator

// Unsupported : g_charset_converter_get_use_fallback : no return generator

// Unsupported : g_charset_converter_set_use_fallback : unsupported parameter use_fallback : no type generator for gboolean, gboolean

// ConverterInputStream is a wrapper around the C record GConverterInputStream.
type ConverterInputStream struct {
	native *C.GConverterInputStream
	// parent_instance : no type generator for FilterInputStream, GFilterInputStream
	// Private : priv
}

func ConverterInputStreamNewFromC(u unsafe.Pointer) *ConverterInputStream {
	c := (*C.GConverterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStream{native: c}

	return g
}

func (recv *ConverterInputStream) toC() *C.GConverterInputStream {

	return recv.native
}

// Unsupported : g_converter_input_stream_new : unsupported parameter base_stream : no type generator for InputStream, GInputStream*

// Unsupported : g_converter_input_stream_get_converter : no return generator

// ConverterOutputStream is a wrapper around the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native *C.GConverterOutputStream
	// parent_instance : no type generator for FilterOutputStream, GFilterOutputStream
	// Private : priv
}

func ConverterOutputStreamNewFromC(u unsafe.Pointer) *ConverterOutputStream {
	c := (*C.GConverterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStream{native: c}

	return g
}

func (recv *ConverterOutputStream) toC() *C.GConverterOutputStream {

	return recv.native
}

// Unsupported : g_converter_output_stream_new : unsupported parameter base_stream : no type generator for OutputStream, GOutputStream*

// Unsupported : g_converter_output_stream_get_converter : no return generator

// DBusActionGroup is a wrapper around the C record GDBusActionGroup.
type DBusActionGroup struct {
	native *C.GDBusActionGroup
}

func DBusActionGroupNewFromC(u unsafe.Pointer) *DBusActionGroup {
	c := (*C.GDBusActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &DBusActionGroup{native: c}

	return g
}

func (recv *DBusActionGroup) toC() *C.GDBusActionGroup {

	return recv.native
}

// DBusMenuModel is a wrapper around the C record GDBusMenuModel.
type DBusMenuModel struct {
	native *C.GDBusMenuModel
}

func DBusMenuModelNewFromC(u unsafe.Pointer) *DBusMenuModel {
	c := (*C.GDBusMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &DBusMenuModel{native: c}

	return g
}

func (recv *DBusMenuModel) toC() *C.GDBusMenuModel {

	return recv.native
}

// DataInputStream is a wrapper around the C record GDataInputStream.
type DataInputStream struct {
	native *C.GDataInputStream
	// parent_instance : no type generator for BufferedInputStream, GBufferedInputStream
	// Private : priv
}

func DataInputStreamNewFromC(u unsafe.Pointer) *DataInputStream {
	c := (*C.GDataInputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStream{native: c}

	return g
}

func (recv *DataInputStream) toC() *C.GDataInputStream {

	return recv.native
}

// Unsupported : g_data_input_stream_new : unsupported parameter base_stream : no type generator for InputStream, GInputStream*

// GetByteOrder is a wrapper around the C function g_data_input_stream_get_byte_order.
func (recv *DataInputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_input_stream_get_byte_order((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// GetNewlineType is a wrapper around the C function g_data_input_stream_get_newline_type.
func (recv *DataInputStream) GetNewlineType() DataStreamNewlineType {
	retC := C.g_data_input_stream_get_newline_type((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamNewlineType)(retC)

	return retGo
}

// Unsupported : g_data_input_stream_read_byte : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_int16 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_int32 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_int64 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_line : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_line_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_line_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_line_finish_utf8 : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_line_utf8 : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_uint16 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_uint32 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_uint64 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_until : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_until_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_until_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_upto : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_input_stream_read_upto_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_set_byte_order : no return generator

// Unsupported : g_data_input_stream_set_newline_type : no return generator

// DataOutputStream is a wrapper around the C record GDataOutputStream.
type DataOutputStream struct {
	native *C.GDataOutputStream
	// parent_instance : no type generator for FilterOutputStream, GFilterOutputStream
	// Private : priv
}

func DataOutputStreamNewFromC(u unsafe.Pointer) *DataOutputStream {
	c := (*C.GDataOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStream{native: c}

	return g
}

func (recv *DataOutputStream) toC() *C.GDataOutputStream {

	return recv.native
}

// Unsupported : g_data_output_stream_new : unsupported parameter base_stream : no type generator for OutputStream, GOutputStream*

// GetByteOrder is a wrapper around the C function g_data_output_stream_get_byte_order.
func (recv *DataOutputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_output_stream_get_byte_order((*C.GDataOutputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// Unsupported : g_data_output_stream_put_byte : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_int16 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_int32 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_int64 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_string : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_uint16 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_uint32 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_put_uint64 : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_data_output_stream_set_byte_order : no return generator

// DesktopAppInfo is a wrapper around the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native *C.GDesktopAppInfo
}

func DesktopAppInfoNewFromC(u unsafe.Pointer) *DesktopAppInfo {
	c := (*C.GDesktopAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfo{native: c}

	return g
}

func (recv *DesktopAppInfo) toC() *C.GDesktopAppInfo {

	return recv.native
}

// Unsupported : g_desktop_app_info_new : no return generator

// Unsupported : g_desktop_app_info_new_from_filename : no return generator

// Unsupported : g_desktop_app_info_new_from_keyfile : no return generator

// Unsupported : g_desktop_app_info_get_boolean : no return generator

// GetCategories is a wrapper around the C function g_desktop_app_info_get_categories.
func (recv *DesktopAppInfo) GetCategories() string {
	retC := C.g_desktop_app_info_get_categories((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetGenericName is a wrapper around the C function g_desktop_app_info_get_generic_name.
func (recv *DesktopAppInfo) GetGenericName() string {
	retC := C.g_desktop_app_info_get_generic_name((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_desktop_app_info_get_is_hidden : no return generator

// Unsupported : g_desktop_app_info_get_keywords : no return type

// Unsupported : g_desktop_app_info_get_nodisplay : no return generator

// Unsupported : g_desktop_app_info_get_show_in : no return generator

// Unsupported : g_desktop_app_info_has_key : no return generator

// Unsupported : g_desktop_app_info_launch_action : unsupported parameter launch_context : no type generator for AppLaunchContext, GAppLaunchContext*

// Unsupported : g_desktop_app_info_launch_uris_as_manager : unsupported parameter launch_context : no type generator for AppLaunchContext, GAppLaunchContext*

// Unsupported : g_desktop_app_info_list_actions : no return type

// Emblem is a wrapper around the C record GEmblem.
type Emblem struct {
	native *C.GEmblem
}

func EmblemNewFromC(u unsafe.Pointer) *Emblem {
	c := (*C.GEmblem)(u)
	if c == nil {
		return nil
	}

	g := &Emblem{native: c}

	return g
}

func (recv *Emblem) toC() *C.GEmblem {

	return recv.native
}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_get_icon : no return generator

// EmblemedIcon is a wrapper around the C record GEmblemedIcon.
type EmblemedIcon struct {
	native *C.GEmblemedIcon
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func EmblemedIconNewFromC(u unsafe.Pointer) *EmblemedIcon {
	c := (*C.GEmblemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIcon{native: c}

	return g
}

func (recv *EmblemedIcon) toC() *C.GEmblemedIcon {

	return recv.native
}

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_add_emblem : unsupported parameter emblem : no type generator for Emblem, GEmblem*

// Unsupported : g_emblemed_icon_clear_emblems : no return generator

// Unsupported : g_emblemed_icon_get_icon : no return generator

// FileEnumerator is a wrapper around the C record GFileEnumerator.
type FileEnumerator struct {
	native *C.GFileEnumerator
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func FileEnumeratorNewFromC(u unsafe.Pointer) *FileEnumerator {
	c := (*C.GFileEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumerator{native: c}

	return g
}

func (recv *FileEnumerator) toC() *C.GFileEnumerator {

	return recv.native
}

// Unsupported : g_file_enumerator_close : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_enumerator_close_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_enumerator_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_enumerator_get_child : unsupported parameter info : no type generator for FileInfo, GFileInfo*

// Unsupported : g_file_enumerator_get_container : no return generator

// Unsupported : g_file_enumerator_has_pending : no return generator

// Unsupported : g_file_enumerator_is_closed : no return generator

// Unsupported : g_file_enumerator_iterate : unsupported parameter out_info : no type generator for FileInfo, GFileInfo**

// Unsupported : g_file_enumerator_next_file : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_enumerator_next_files_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_enumerator_next_files_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_enumerator_set_pending : unsupported parameter pending : no type generator for gboolean, gboolean

// FileIOStream is a wrapper around the C record GFileIOStream.
type FileIOStream struct {
	native *C.GFileIOStream
	// parent_instance : no type generator for IOStream, GIOStream
	// Private : priv
}

func FileIOStreamNewFromC(u unsafe.Pointer) *FileIOStream {
	c := (*C.GFileIOStream)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStream{native: c}

	return g
}

func (recv *FileIOStream) toC() *C.GFileIOStream {

	return recv.native
}

// Unsupported : g_file_io_stream_query_info : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_io_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// FileIcon is a wrapper around the C record GFileIcon.
type FileIcon struct {
	native *C.GFileIcon
}

func FileIconNewFromC(u unsafe.Pointer) *FileIcon {
	c := (*C.GFileIcon)(u)
	if c == nil {
		return nil
	}

	g := &FileIcon{native: c}

	return g
}

func (recv *FileIcon) toC() *C.GFileIcon {

	return recv.native
}

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_file_icon_get_file : no return generator

// FileInfo is a wrapper around the C record GFileInfo.
type FileInfo struct {
	native *C.GFileInfo
}

func FileInfoNewFromC(u unsafe.Pointer) *FileInfo {
	c := (*C.GFileInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileInfo{native: c}

	return g
}

func (recv *FileInfo) toC() *C.GFileInfo {

	return recv.native
}

// Unsupported : g_file_info_new : no return generator

// Unsupported : g_file_info_clear_status : no return generator

// Unsupported : g_file_info_copy_into : unsupported parameter dest_info : no type generator for FileInfo, GFileInfo*

// Unsupported : g_file_info_dup : no return generator

// GetAttributeAsString is a wrapper around the C function g_file_info_get_attribute_as_string.
func (recv *FileInfo) GetAttributeAsString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_as_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_info_get_attribute_boolean : no return generator

// GetAttributeByteString is a wrapper around the C function g_file_info_get_attribute_byte_string.
func (recv *FileInfo) GetAttributeByteString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_byte_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_file_info_get_attribute_data : unsupported parameter type : GFileAttributeType* with indirection level of 1

// GetAttributeInt32 is a wrapper around the C function g_file_info_get_attribute_int32.
func (recv *FileInfo) GetAttributeInt32(attribute string) int32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int32)(retC)

	return retGo
}

// GetAttributeInt64 is a wrapper around the C function g_file_info_get_attribute_int64.
func (recv *FileInfo) GetAttributeInt64(attribute string) int64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int64)(retC)

	return retGo
}

// Unsupported : g_file_info_get_attribute_object : no return generator

// GetAttributeStatus is a wrapper around the C function g_file_info_get_attribute_status.
func (recv *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_status((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeStatus)(retC)

	return retGo
}

// GetAttributeString is a wrapper around the C function g_file_info_get_attribute_string.
func (recv *FileInfo) GetAttributeString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_file_info_get_attribute_stringv : no return type

// GetAttributeType is a wrapper around the C function g_file_info_get_attribute_type.
func (recv *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_type((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeType)(retC)

	return retGo
}

// GetAttributeUint32 is a wrapper around the C function g_file_info_get_attribute_uint32.
func (recv *FileInfo) GetAttributeUint32(attribute string) uint32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint32)(retC)

	return retGo
}

// GetAttributeUint64 is a wrapper around the C function g_file_info_get_attribute_uint64.
func (recv *FileInfo) GetAttributeUint64(attribute string) uint64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint64)(retC)

	return retGo
}

// GetContentType is a wrapper around the C function g_file_info_get_content_type.
func (recv *FileInfo) GetContentType() string {
	retC := C.g_file_info_get_content_type((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function g_file_info_get_display_name.
func (recv *FileInfo) GetDisplayName() string {
	retC := C.g_file_info_get_display_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEditName is a wrapper around the C function g_file_info_get_edit_name.
func (recv *FileInfo) GetEditName() string {
	retC := C.g_file_info_get_edit_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEtag is a wrapper around the C function g_file_info_get_etag.
func (recv *FileInfo) GetEtag() string {
	retC := C.g_file_info_get_etag((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFileType is a wrapper around the C function g_file_info_get_file_type.
func (recv *FileInfo) GetFileType() FileType {
	retC := C.g_file_info_get_file_type((*C.GFileInfo)(recv.native))
	retGo := (FileType)(retC)

	return retGo
}

// Unsupported : g_file_info_get_icon : no return generator

// Unsupported : g_file_info_get_is_backup : no return generator

// Unsupported : g_file_info_get_is_hidden : no return generator

// Unsupported : g_file_info_get_is_symlink : no return generator

// Unsupported : g_file_info_get_modification_time : no return generator

// GetName is a wrapper around the C function g_file_info_get_name.
func (recv *FileInfo) GetName() string {
	retC := C.g_file_info_get_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSize is a wrapper around the C function g_file_info_get_size.
func (recv *FileInfo) GetSize() uint64 {
	retC := C.g_file_info_get_size((*C.GFileInfo)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetSortOrder is a wrapper around the C function g_file_info_get_sort_order.
func (recv *FileInfo) GetSortOrder() int32 {
	retC := C.g_file_info_get_sort_order((*C.GFileInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_file_info_get_symbolic_icon : no return generator

// GetSymlinkTarget is a wrapper around the C function g_file_info_get_symlink_target.
func (recv *FileInfo) GetSymlinkTarget() string {
	retC := C.g_file_info_get_symlink_target((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_file_info_has_attribute : no return generator

// Unsupported : g_file_info_has_namespace : no return generator

// Unsupported : g_file_info_list_attributes : no return type

// Unsupported : g_file_info_remove_attribute : no return generator

// Unsupported : g_file_info_set_attribute : no return generator

// Unsupported : g_file_info_set_attribute_boolean : unsupported parameter attr_value : no type generator for gboolean, gboolean

// Unsupported : g_file_info_set_attribute_byte_string : no return generator

// Unsupported : g_file_info_set_attribute_int32 : no return generator

// Unsupported : g_file_info_set_attribute_int64 : no return generator

// Unsupported : g_file_info_set_attribute_mask : no return generator

// Unsupported : g_file_info_set_attribute_object : unsupported parameter attr_value : no type generator for GObject.Object, GObject*

// Unsupported : g_file_info_set_attribute_status : no return generator

// Unsupported : g_file_info_set_attribute_string : no return generator

// Unsupported : g_file_info_set_attribute_stringv : unsupported parameter attr_value : no param type

// Unsupported : g_file_info_set_attribute_uint32 : no return generator

// Unsupported : g_file_info_set_attribute_uint64 : no return generator

// Unsupported : g_file_info_set_content_type : no return generator

// Unsupported : g_file_info_set_display_name : no return generator

// Unsupported : g_file_info_set_edit_name : no return generator

// Unsupported : g_file_info_set_file_type : no return generator

// Unsupported : g_file_info_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_info_set_is_hidden : unsupported parameter is_hidden : no type generator for gboolean, gboolean

// Unsupported : g_file_info_set_is_symlink : unsupported parameter is_symlink : no type generator for gboolean, gboolean

// Unsupported : g_file_info_set_modification_time : no return generator

// Unsupported : g_file_info_set_name : no return generator

// Unsupported : g_file_info_set_size : no return generator

// Unsupported : g_file_info_set_sort_order : no return generator

// Unsupported : g_file_info_set_symbolic_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_info_set_symlink_target : no return generator

// Unsupported : g_file_info_unset_attribute_mask : no return generator

// FileInputStream is a wrapper around the C record GFileInputStream.
type FileInputStream struct {
	native *C.GFileInputStream
	// parent_instance : no type generator for InputStream, GInputStream
	// Private : priv
}

func FileInputStreamNewFromC(u unsafe.Pointer) *FileInputStream {
	c := (*C.GFileInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStream{native: c}

	return g
}

func (recv *FileInputStream) toC() *C.GFileInputStream {

	return recv.native
}

// Unsupported : g_file_input_stream_query_info : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_input_stream_query_info_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_input_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// FileMonitor is a wrapper around the C record GFileMonitor.
type FileMonitor struct {
	native *C.GFileMonitor
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func FileMonitorNewFromC(u unsafe.Pointer) *FileMonitor {
	c := (*C.GFileMonitor)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitor{native: c}

	return g
}

func (recv *FileMonitor) toC() *C.GFileMonitor {

	return recv.native
}

// Unsupported : g_file_monitor_cancel : no return generator

// Unsupported : g_file_monitor_emit_event : unsupported parameter child : no type generator for File, GFile*

// Unsupported : g_file_monitor_is_cancelled : no return generator

// Unsupported : g_file_monitor_set_rate_limit : no return generator

// FileOutputStream is a wrapper around the C record GFileOutputStream.
type FileOutputStream struct {
	native *C.GFileOutputStream
	// parent_instance : no type generator for OutputStream, GOutputStream
	// Private : priv
}

func FileOutputStreamNewFromC(u unsafe.Pointer) *FileOutputStream {
	c := (*C.GFileOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStream{native: c}

	return g
}

func (recv *FileOutputStream) toC() *C.GFileOutputStream {

	return recv.native
}

// GetEtag is a wrapper around the C function g_file_output_stream_get_etag.
func (recv *FileOutputStream) GetEtag() string {
	retC := C.g_file_output_stream_get_etag((*C.GFileOutputStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_output_stream_query_info : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_output_stream_query_info_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_file_output_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// FilenameCompleter is a wrapper around the C record GFilenameCompleter.
type FilenameCompleter struct {
	native *C.GFilenameCompleter
}

func FilenameCompleterNewFromC(u unsafe.Pointer) *FilenameCompleter {
	c := (*C.GFilenameCompleter)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleter{native: c}

	return g
}

func (recv *FilenameCompleter) toC() *C.GFilenameCompleter {

	return recv.native
}

// Unsupported : g_filename_completer_new : no return generator

// GetCompletionSuffix is a wrapper around the C function g_filename_completer_get_completion_suffix.
func (recv *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	c_initial_text := C.CString(initialText)
	defer C.free(unsafe.Pointer(c_initial_text))

	retC := C.g_filename_completer_get_completion_suffix((*C.GFilenameCompleter)(recv.native), c_initial_text)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_filename_completer_get_completions : no return type

// Unsupported : g_filename_completer_set_dirs_only : unsupported parameter dirs_only : no type generator for gboolean, gboolean

// FilterInputStream is a wrapper around the C record GFilterInputStream.
type FilterInputStream struct {
	native *C.GFilterInputStream
	// parent_instance : no type generator for InputStream, GInputStream
	// base_stream : no type generator for InputStream, GInputStream*
}

func FilterInputStreamNewFromC(u unsafe.Pointer) *FilterInputStream {
	c := (*C.GFilterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStream{native: c}

	return g
}

func (recv *FilterInputStream) toC() *C.GFilterInputStream {

	return recv.native
}

// Unsupported : g_filter_input_stream_get_base_stream : no return generator

// Unsupported : g_filter_input_stream_get_close_base_stream : no return generator

// Unsupported : g_filter_input_stream_set_close_base_stream : unsupported parameter close_base : no type generator for gboolean, gboolean

// FilterOutputStream is a wrapper around the C record GFilterOutputStream.
type FilterOutputStream struct {
	native *C.GFilterOutputStream
	// parent_instance : no type generator for OutputStream, GOutputStream
	// base_stream : no type generator for OutputStream, GOutputStream*
}

func FilterOutputStreamNewFromC(u unsafe.Pointer) *FilterOutputStream {
	c := (*C.GFilterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStream{native: c}

	return g
}

func (recv *FilterOutputStream) toC() *C.GFilterOutputStream {

	return recv.native
}

// Unsupported : g_filter_output_stream_get_base_stream : no return generator

// Unsupported : g_filter_output_stream_get_close_base_stream : no return generator

// Unsupported : g_filter_output_stream_set_close_base_stream : unsupported parameter close_base : no type generator for gboolean, gboolean

// IOModule is a wrapper around the C record GIOModule.
type IOModule struct {
	native *C.GIOModule
}

func IOModuleNewFromC(u unsafe.Pointer) *IOModule {
	c := (*C.GIOModule)(u)
	if c == nil {
		return nil
	}

	g := &IOModule{native: c}

	return g
}

func (recv *IOModule) toC() *C.GIOModule {

	return recv.native
}

// Unsupported : g_io_module_new : no return generator

// Unsupported : g_io_module_load : no return generator

// Unsupported : g_io_module_unload : no return generator

// IOStream is a wrapper around the C record GIOStream.
type IOStream struct {
	native *C.GIOStream
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func IOStreamNewFromC(u unsafe.Pointer) *IOStream {
	c := (*C.GIOStream)(u)
	if c == nil {
		return nil
	}

	g := &IOStream{native: c}

	return g
}

func (recv *IOStream) toC() *C.GIOStream {

	return recv.native
}

// Unsupported : g_io_stream_clear_pending : no return generator

// Unsupported : g_io_stream_close : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_io_stream_close_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_io_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_io_stream_get_input_stream : no return generator

// Unsupported : g_io_stream_get_output_stream : no return generator

// Unsupported : g_io_stream_has_pending : no return generator

// Unsupported : g_io_stream_is_closed : no return generator

// Unsupported : g_io_stream_set_pending : no return generator

// Unsupported : g_io_stream_splice_async : unsupported parameter stream2 : no type generator for IOStream, GIOStream*

// InetAddress is a wrapper around the C record GInetAddress.
type InetAddress struct {
	native *C.GInetAddress
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func InetAddressNewFromC(u unsafe.Pointer) *InetAddress {
	c := (*C.GInetAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetAddress{native: c}

	return g
}

func (recv *InetAddress) toC() *C.GInetAddress {

	return recv.native
}

// Unsupported : g_inet_address_new_any : no return generator

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_inet_address_new_from_string : no return generator

// Unsupported : g_inet_address_new_loopback : no return generator

// Unsupported : g_inet_address_equal : unsupported parameter other_address : no type generator for InetAddress, GInetAddress*

// Unsupported : g_inet_address_get_is_any : no return generator

// Unsupported : g_inet_address_get_is_link_local : no return generator

// Unsupported : g_inet_address_get_is_loopback : no return generator

// Unsupported : g_inet_address_get_is_mc_global : no return generator

// Unsupported : g_inet_address_get_is_mc_link_local : no return generator

// Unsupported : g_inet_address_get_is_mc_node_local : no return generator

// Unsupported : g_inet_address_get_is_mc_org_local : no return generator

// Unsupported : g_inet_address_get_is_mc_site_local : no return generator

// Unsupported : g_inet_address_get_is_multicast : no return generator

// Unsupported : g_inet_address_get_is_site_local : no return generator

// Unsupported : g_inet_address_to_bytes : no return generator

// InetSocketAddress is a wrapper around the C record GInetSocketAddress.
type InetSocketAddress struct {
	native *C.GInetSocketAddress
	// parent_instance : no type generator for SocketAddress, GSocketAddress
	// Private : priv
}

func InetSocketAddressNewFromC(u unsafe.Pointer) *InetSocketAddress {
	c := (*C.GInetSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddress{native: c}

	return g
}

func (recv *InetSocketAddress) toC() *C.GInetSocketAddress {

	return recv.native
}

// Unsupported : g_inet_socket_address_new : unsupported parameter address : no type generator for InetAddress, GInetAddress*

// Unsupported : g_inet_socket_address_new_from_string : no return generator

// Unsupported : g_inet_socket_address_get_address : no return generator

// InputStream is a wrapper around the C record GInputStream.
type InputStream struct {
	native *C.GInputStream
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func InputStreamNewFromC(u unsafe.Pointer) *InputStream {
	c := (*C.GInputStream)(u)
	if c == nil {
		return nil
	}

	g := &InputStream{native: c}

	return g
}

func (recv *InputStream) toC() *C.GInputStream {

	return recv.native
}

// Unsupported : g_input_stream_clear_pending : no return generator

// Unsupported : g_input_stream_close : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_close_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_has_pending : no return generator

// Unsupported : g_input_stream_is_closed : no return generator

// Unsupported : g_input_stream_read : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_bytes : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_read_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_set_pending : no return generator

// Unsupported : g_input_stream_skip : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_skip_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_input_stream_skip_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ListStore is a wrapper around the C record GListStore.
type ListStore struct {
	native *C.GListStore
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	return g
}

func (recv *ListStore) toC() *C.GListStore {

	return recv.native
}

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_list_store_append : no return generator

// Unsupported : g_list_store_insert : no return generator

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_remove : no return generator

// Unsupported : g_list_store_remove_all : no return generator

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_splice : unsupported parameter additions : no param type

// MemoryInputStream is a wrapper around the C record GMemoryInputStream.
type MemoryInputStream struct {
	native *C.GMemoryInputStream
	// parent_instance : no type generator for InputStream, GInputStream
	// Private : priv
}

func MemoryInputStreamNewFromC(u unsafe.Pointer) *MemoryInputStream {
	c := (*C.GMemoryInputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStream{native: c}

	return g
}

func (recv *MemoryInputStream) toC() *C.GMemoryInputStream {

	return recv.native
}

// Unsupported : g_memory_input_stream_new : no return generator

// Unsupported : g_memory_input_stream_new_from_bytes : no return generator

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_input_stream_add_bytes : no return generator

// Unsupported : g_memory_input_stream_add_data : unsupported parameter data : no param type

// MemoryOutputStream is a wrapper around the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native *C.GMemoryOutputStream
	// parent_instance : no type generator for OutputStream, GOutputStream
	// Private : priv
}

func MemoryOutputStreamNewFromC(u unsafe.Pointer) *MemoryOutputStream {
	c := (*C.GMemoryOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStream{native: c}

	return g
}

func (recv *MemoryOutputStream) toC() *C.GMemoryOutputStream {

	return recv.native
}

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Unsupported : g_memory_output_stream_new_resizable : no return generator

// GetData is a wrapper around the C function g_memory_output_stream_get_data.
func (recv *MemoryOutputStream) GetData() uintptr {
	retC := C.g_memory_output_stream_get_data((*C.GMemoryOutputStream)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetSize is a wrapper around the C function g_memory_output_stream_get_size.
func (recv *MemoryOutputStream) GetSize() uint64 {
	retC := C.g_memory_output_stream_get_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// MountOperation is a wrapper around the C record GMountOperation.
type MountOperation struct {
	native *C.GMountOperation
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	c := (*C.GMountOperation)(u)
	if c == nil {
		return nil
	}

	g := &MountOperation{native: c}

	return g
}

func (recv *MountOperation) toC() *C.GMountOperation {

	return recv.native
}

// Unsupported : g_mount_operation_new : no return generator

// Unsupported : g_mount_operation_get_anonymous : no return generator

// GetChoice is a wrapper around the C function g_mount_operation_get_choice.
func (recv *MountOperation) GetChoice() int32 {
	retC := C.g_mount_operation_get_choice((*C.GMountOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDomain is a wrapper around the C function g_mount_operation_get_domain.
func (recv *MountOperation) GetDomain() string {
	retC := C.g_mount_operation_get_domain((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPassword is a wrapper around the C function g_mount_operation_get_password.
func (recv *MountOperation) GetPassword() string {
	retC := C.g_mount_operation_get_password((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPasswordSave is a wrapper around the C function g_mount_operation_get_password_save.
func (recv *MountOperation) GetPasswordSave() PasswordSave {
	retC := C.g_mount_operation_get_password_save((*C.GMountOperation)(recv.native))
	retGo := (PasswordSave)(retC)

	return retGo
}

// GetUsername is a wrapper around the C function g_mount_operation_get_username.
func (recv *MountOperation) GetUsername() string {
	retC := C.g_mount_operation_get_username((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_mount_operation_reply : no return generator

// Unsupported : g_mount_operation_set_anonymous : unsupported parameter anonymous : no type generator for gboolean, gboolean

// Unsupported : g_mount_operation_set_choice : no return generator

// Unsupported : g_mount_operation_set_domain : no return generator

// Unsupported : g_mount_operation_set_password : no return generator

// Unsupported : g_mount_operation_set_password_save : no return generator

// Unsupported : g_mount_operation_set_username : no return generator

// NativeVolumeMonitor is a wrapper around the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native *C.GNativeVolumeMonitor
	// parent_instance : no type generator for VolumeMonitor, GVolumeMonitor
}

func NativeVolumeMonitorNewFromC(u unsafe.Pointer) *NativeVolumeMonitor {
	c := (*C.GNativeVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitor{native: c}

	return g
}

func (recv *NativeVolumeMonitor) toC() *C.GNativeVolumeMonitor {

	return recv.native
}

// NetworkAddress is a wrapper around the C record GNetworkAddress.
type NetworkAddress struct {
	native *C.GNetworkAddress
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func NetworkAddressNewFromC(u unsafe.Pointer) *NetworkAddress {
	c := (*C.GNetworkAddress)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddress{native: c}

	return g
}

func (recv *NetworkAddress) toC() *C.GNetworkAddress {

	return recv.native
}

// Unsupported : g_network_address_new : no return generator

// Unsupported : g_network_address_new_loopback : no return generator

// NetworkService is a wrapper around the C record GNetworkService.
type NetworkService struct {
	native *C.GNetworkService
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func NetworkServiceNewFromC(u unsafe.Pointer) *NetworkService {
	c := (*C.GNetworkService)(u)
	if c == nil {
		return nil
	}

	g := &NetworkService{native: c}

	return g
}

func (recv *NetworkService) toC() *C.GNetworkService {

	return recv.native
}

// Unsupported : g_network_service_new : no return generator

// Unsupported : g_network_service_set_scheme : no return generator

// OutputStream is a wrapper around the C record GOutputStream.
type OutputStream struct {
	native *C.GOutputStream
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func OutputStreamNewFromC(u unsafe.Pointer) *OutputStream {
	c := (*C.GOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &OutputStream{native: c}

	return g
}

func (recv *OutputStream) toC() *C.GOutputStream {

	return recv.native
}

// Unsupported : g_output_stream_clear_pending : no return generator

// Unsupported : g_output_stream_close : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_close_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_flush : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_flush_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_flush_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_has_pending : no return generator

// Unsupported : g_output_stream_is_closed : no return generator

// Unsupported : g_output_stream_is_closing : no return generator

// Unsupported : g_output_stream_printf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_output_stream_set_pending : no return generator

// Unsupported : g_output_stream_splice : unsupported parameter source : no type generator for InputStream, GInputStream*

// Unsupported : g_output_stream_splice_async : unsupported parameter source : no type generator for InputStream, GInputStream*

// Unsupported : g_output_stream_splice_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_vprintf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_output_stream_write : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all_async : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_write_async : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_bytes : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_write_bytes_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_output_stream_write_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_write_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Permission is a wrapper around the C record GPermission.
type Permission struct {
	native *C.GPermission
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func PermissionNewFromC(u unsafe.Pointer) *Permission {
	c := (*C.GPermission)(u)
	if c == nil {
		return nil
	}

	g := &Permission{native: c}

	return g
}

func (recv *Permission) toC() *C.GPermission {

	return recv.native
}

// Unsupported : g_permission_acquire : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_permission_acquire_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_permission_acquire_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_permission_get_allowed : no return generator

// Unsupported : g_permission_get_can_acquire : no return generator

// Unsupported : g_permission_get_can_release : no return generator

// Unsupported : g_permission_impl_update : unsupported parameter allowed : no type generator for gboolean, gboolean

// Unsupported : g_permission_release : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_permission_release_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_permission_release_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ProxyAddressEnumerator is a wrapper around the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native *C.GProxyAddressEnumerator
	// parent_instance : no type generator for SocketAddressEnumerator, GSocketAddressEnumerator
	// priv : record
}

func ProxyAddressEnumeratorNewFromC(u unsafe.Pointer) *ProxyAddressEnumerator {
	c := (*C.GProxyAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumerator{native: c}

	return g
}

func (recv *ProxyAddressEnumerator) toC() *C.GProxyAddressEnumerator {

	return recv.native
}

// Resolver is a wrapper around the C record GResolver.
type Resolver struct {
	native *C.GResolver
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func ResolverNewFromC(u unsafe.Pointer) *Resolver {
	c := (*C.GResolver)(u)
	if c == nil {
		return nil
	}

	g := &Resolver{native: c}

	return g
}

func (recv *Resolver) toC() *C.GResolver {

	return recv.native
}

// Unsupported : g_resolver_lookup_by_address : unsupported parameter address : no type generator for InetAddress, GInetAddress*

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter address : no type generator for InetAddress, GInetAddress*

// Unsupported : g_resolver_lookup_by_address_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_by_name : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_by_name_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_records : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_records_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_records_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_service : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_service_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_resolver_lookup_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_set_default : no return generator

// Settings is a wrapper around the C record GSettings.
type Settings struct {
	native *C.GSettings
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.GSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	return g
}

func (recv *Settings) toC() *C.GSettings {

	return recv.native
}

// Unsupported : g_settings_new : no return generator

// Unsupported : g_settings_new_full : unsupported parameter backend : no type generator for SettingsBackend, GSettingsBackend*

// Unsupported : g_settings_new_with_backend : unsupported parameter backend : no type generator for SettingsBackend, GSettingsBackend*

// Unsupported : g_settings_new_with_backend_and_path : unsupported parameter backend : no type generator for SettingsBackend, GSettingsBackend*

// Unsupported : g_settings_new_with_path : no return generator

// Unsupported : g_settings_apply : no return generator

// Unsupported : g_settings_bind : unsupported parameter flags : no type generator for SettingsBindFlags, GSettingsBindFlags

// Unsupported : g_settings_bind_with_mapping : unsupported parameter flags : no type generator for SettingsBindFlags, GSettingsBindFlags

// Unsupported : g_settings_bind_writable : unsupported parameter inverted : no type generator for gboolean, gboolean

// Unsupported : g_settings_create_action : no return generator

// Unsupported : g_settings_delay : no return generator

// Unsupported : g_settings_get : unsupported parameter ... : varargs

// Unsupported : g_settings_get_boolean : no return generator

// Unsupported : g_settings_get_child : no return generator

// Unsupported : g_settings_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_has_unapplied : no return generator

// Unsupported : g_settings_get_mapped : unsupported parameter mapping : no type generator for SettingsGetMapping, GSettingsGetMapping

// Unsupported : g_settings_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_strv : no return type

// Unsupported : g_settings_get_user_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_is_writable : no return generator

// Unsupported : g_settings_list_children : no return type

// Unsupported : g_settings_list_keys : no return type

// Unsupported : g_settings_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_reset : no return generator

// Unsupported : g_settings_revert : no return generator

// Unsupported : g_settings_set : unsupported parameter ... : varargs

// Unsupported : g_settings_set_boolean : unsupported parameter value : no type generator for gboolean, gboolean

// Unsupported : g_settings_set_double : no return generator

// Unsupported : g_settings_set_enum : no return generator

// Unsupported : g_settings_set_flags : no return generator

// Unsupported : g_settings_set_int : no return generator

// Unsupported : g_settings_set_int64 : no return generator

// Unsupported : g_settings_set_string : no return generator

// Unsupported : g_settings_set_strv : unsupported parameter value : no param type

// Unsupported : g_settings_set_uint : no return generator

// Unsupported : g_settings_set_uint64 : no return generator

// Unsupported : g_settings_set_value : unsupported parameter value : Blacklisted record : GVariant

// SettingsBackend is a wrapper around the C record GSettingsBackend.
type SettingsBackend struct {
	native *C.GSettingsBackend
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func SettingsBackendNewFromC(u unsafe.Pointer) *SettingsBackend {
	c := (*C.GSettingsBackend)(u)
	if c == nil {
		return nil
	}

	g := &SettingsBackend{native: c}

	return g
}

func (recv *SettingsBackend) toC() *C.GSettingsBackend {

	return recv.native
}

// Unsupported : g_settings_backend_changed : no return generator

// Unsupported : g_settings_backend_changed_tree : no return generator

// Unsupported : g_settings_backend_keys_changed : unsupported parameter items : no param type

// Unsupported : g_settings_backend_path_changed : no return generator

// Unsupported : g_settings_backend_path_writable_changed : no return generator

// Unsupported : g_settings_backend_writable_changed : no return generator

// SimpleAction is a wrapper around the C record GSimpleAction.
type SimpleAction struct {
	native *C.GSimpleAction
}

func SimpleActionNewFromC(u unsafe.Pointer) *SimpleAction {
	c := (*C.GSimpleAction)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAction{native: c}

	return g
}

func (recv *SimpleAction) toC() *C.GSimpleAction {

	return recv.native
}

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_set_enabled : unsupported parameter enabled : no type generator for gboolean, gboolean

// Unsupported : g_simple_action_set_state : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_set_state_hint : unsupported parameter state_hint : Blacklisted record : GVariant

// SimpleAsyncResult is a wrapper around the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native *C.GSimpleAsyncResult
}

func SimpleAsyncResultNewFromC(u unsafe.Pointer) *SimpleAsyncResult {
	c := (*C.GSimpleAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResult{native: c}

	return g
}

func (recv *SimpleAsyncResult) toC() *C.GSimpleAsyncResult {

	return recv.native
}

// Unsupported : g_simple_async_result_new : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_simple_async_result_new_error : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_simple_async_result_complete : no return generator

// Unsupported : g_simple_async_result_complete_in_idle : no return generator

// Unsupported : g_simple_async_result_get_op_res_gboolean : no return generator

// GetOpResGpointer is a wrapper around the C function g_simple_async_result_get_op_res_gpointer.
func (recv *SimpleAsyncResult) GetOpResGpointer() uintptr {
	retC := C.g_simple_async_result_get_op_res_gpointer((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// GetOpResGssize is a wrapper around the C function g_simple_async_result_get_op_res_gssize.
func (recv *SimpleAsyncResult) GetOpResGssize() int64 {
	retC := C.g_simple_async_result_get_op_res_gssize((*C.GSimpleAsyncResult)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetSourceTag is a wrapper around the C function g_simple_async_result_get_source_tag.
func (recv *SimpleAsyncResult) GetSourceTag() uintptr {
	retC := C.g_simple_async_result_get_source_tag((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_simple_async_result_propagate_error : no return generator

// Unsupported : g_simple_async_result_run_in_thread : unsupported parameter func : no type generator for SimpleAsyncThreadFunc, GSimpleAsyncThreadFunc

// Unsupported : g_simple_async_result_set_check_cancellable : unsupported parameter check_cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_simple_async_result_set_error : unsupported parameter ... : varargs

// Unsupported : g_simple_async_result_set_error_va : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_simple_async_result_set_from_error : no return generator

// Unsupported : g_simple_async_result_set_handle_cancellation : unsupported parameter handle_cancellation : no type generator for gboolean, gboolean

// Unsupported : g_simple_async_result_set_op_res_gboolean : unsupported parameter op_res : no type generator for gboolean, gboolean

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter destroy_op_res : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_simple_async_result_set_op_res_gssize : no return generator

// Unsupported : g_simple_async_result_take_error : no return generator

// SimplePermission is a wrapper around the C record GSimplePermission.
type SimplePermission struct {
	native *C.GSimplePermission
}

func SimplePermissionNewFromC(u unsafe.Pointer) *SimplePermission {
	c := (*C.GSimplePermission)(u)
	if c == nil {
		return nil
	}

	g := &SimplePermission{native: c}

	return g
}

func (recv *SimplePermission) toC() *C.GSimplePermission {

	return recv.native
}

// Unsupported : g_simple_permission_new : unsupported parameter allowed : no type generator for gboolean, gboolean

// SimpleProxyResolver is a wrapper around the C record GSimpleProxyResolver.
type SimpleProxyResolver struct {
	native *C.GSimpleProxyResolver
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func SimpleProxyResolverNewFromC(u unsafe.Pointer) *SimpleProxyResolver {
	c := (*C.GSimpleProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolver{native: c}

	return g
}

func (recv *SimpleProxyResolver) toC() *C.GSimpleProxyResolver {

	return recv.native
}

// Unsupported : g_simple_proxy_resolver_set_default_proxy : no return generator

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// Unsupported : g_simple_proxy_resolver_set_uri_proxy : no return generator

// SocketAddress is a wrapper around the C record GSocketAddress.
type SocketAddress struct {
	native *C.GSocketAddress
	// parent_instance : no type generator for GObject.Object, GObject
}

func SocketAddressNewFromC(u unsafe.Pointer) *SocketAddress {
	c := (*C.GSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddress{native: c}

	return g
}

func (recv *SocketAddress) toC() *C.GSocketAddress {

	return recv.native
}

// Unsupported : g_socket_address_new_from_native : no return generator

// Unsupported : g_socket_address_to_native : no return generator

// SocketAddressEnumerator is a wrapper around the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native *C.GSocketAddressEnumerator
	// parent_instance : no type generator for GObject.Object, GObject
}

func SocketAddressEnumeratorNewFromC(u unsafe.Pointer) *SocketAddressEnumerator {
	c := (*C.GSocketAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumerator{native: c}

	return g
}

func (recv *SocketAddressEnumerator) toC() *C.GSocketAddressEnumerator {

	return recv.native
}

// Unsupported : g_socket_address_enumerator_next : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_address_enumerator_next_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_address_enumerator_next_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SocketControlMessage is a wrapper around the C record GSocketControlMessage.
type SocketControlMessage struct {
	native *C.GSocketControlMessage
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func SocketControlMessageNewFromC(u unsafe.Pointer) *SocketControlMessage {
	c := (*C.GSocketControlMessage)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessage{native: c}

	return g
}

func (recv *SocketControlMessage) toC() *C.GSocketControlMessage {

	return recv.native
}

// Unsupported : g_socket_control_message_serialize : no return generator

// Task is a wrapper around the C record GTask.
type Task struct {
	native *C.GTask
}

func TaskNewFromC(u unsafe.Pointer) *Task {
	c := (*C.GTask)(u)
	if c == nil {
		return nil
	}

	g := &Task{native: c}

	return g
}

func (recv *Task) toC() *C.GTask {

	return recv.native
}

// Unsupported : g_task_new : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_task_get_cancellable : no return generator

// Unsupported : g_task_get_check_cancellable : no return generator

// Unsupported : g_task_get_completed : no return generator

// Unsupported : g_task_get_return_on_cancel : no return generator

// Unsupported : g_task_had_error : no return generator

// Unsupported : g_task_propagate_boolean : no return generator

// Unsupported : g_task_return_boolean : unsupported parameter result : no type generator for gboolean, gboolean

// Unsupported : g_task_return_error : no return generator

// Unsupported : g_task_return_error_if_cancelled : no return generator

// Unsupported : g_task_return_int : no return generator

// Unsupported : g_task_return_new_error : unsupported parameter ... : varargs

// Unsupported : g_task_return_pointer : unsupported parameter result_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// Unsupported : g_task_set_check_cancellable : unsupported parameter check_cancellable : no type generator for gboolean, gboolean

// Unsupported : g_task_set_priority : no return generator

// Unsupported : g_task_set_return_on_cancel : unsupported parameter return_on_cancel : no type generator for gboolean, gboolean

// Unsupported : g_task_set_source_tag : no return generator

// Unsupported : g_task_set_task_data : unsupported parameter task_data_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// TcpWrapperConnection is a wrapper around the C record GTcpWrapperConnection.
type TcpWrapperConnection struct {
	native *C.GTcpWrapperConnection
	// parent_instance : no type generator for TcpConnection, GTcpConnection
	// priv : record
}

func TcpWrapperConnectionNewFromC(u unsafe.Pointer) *TcpWrapperConnection {
	c := (*C.GTcpWrapperConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnection{native: c}

	return g
}

func (recv *TcpWrapperConnection) toC() *C.GTcpWrapperConnection {

	return recv.native
}

// Unsupported : g_tcp_wrapper_connection_new : unsupported parameter base_io_stream : no type generator for IOStream, GIOStream*

// Unsupported : g_tcp_wrapper_connection_get_base_io_stream : no return generator

// ThemedIcon is a wrapper around the C record GThemedIcon.
type ThemedIcon struct {
	native *C.GThemedIcon
}

func ThemedIconNewFromC(u unsafe.Pointer) *ThemedIcon {
	c := (*C.GThemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIcon{native: c}

	return g
}

func (recv *ThemedIcon) toC() *C.GThemedIcon {

	return recv.native
}

// Unsupported : g_themed_icon_new : no return generator

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_themed_icon_new_with_default_fallbacks : no return generator

// Unsupported : g_themed_icon_append_name : no return generator

// Unsupported : g_themed_icon_get_names : no return type

// Unsupported : g_themed_icon_prepend_name : no return generator

// UnixConnection is a wrapper around the C record GUnixConnection.
type UnixConnection struct {
	native *C.GUnixConnection
	// parent_instance : no type generator for SocketConnection, GSocketConnection
	// priv : record
}

func UnixConnectionNewFromC(u unsafe.Pointer) *UnixConnection {
	c := (*C.GUnixConnection)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnection{native: c}

	return g
}

func (recv *UnixConnection) toC() *C.GUnixConnection {

	return recv.native
}

// Unsupported : g_unix_connection_receive_credentials : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_unix_connection_receive_credentials_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_unix_connection_receive_credentials_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_unix_connection_receive_fd : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_unix_connection_send_credentials : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_unix_connection_send_credentials_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_unix_connection_send_credentials_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_unix_connection_send_fd : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// UnixFDList is a wrapper around the C record GUnixFDList.
type UnixFDList struct {
	native *C.GUnixFDList
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func UnixFDListNewFromC(u unsafe.Pointer) *UnixFDList {
	c := (*C.GUnixFDList)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDList{native: c}

	return g
}

func (recv *UnixFDList) toC() *C.GUnixFDList {

	return recv.native
}

// Unsupported : g_unix_fd_list_new : no return generator

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_fd_list_peek_fds : unsupported parameter length : no type generator for gint, gint*

// Unsupported : g_unix_fd_list_steal_fds : unsupported parameter length : no type generator for gint, gint*

// UnixFDMessage is a wrapper around the C record GUnixFDMessage.
type UnixFDMessage struct {
	native *C.GUnixFDMessage
	// parent_instance : no type generator for SocketControlMessage, GSocketControlMessage
	// priv : record
}

func UnixFDMessageNewFromC(u unsafe.Pointer) *UnixFDMessage {
	c := (*C.GUnixFDMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessage{native: c}

	return g
}

func (recv *UnixFDMessage) toC() *C.GUnixFDMessage {

	return recv.native
}

// Unsupported : g_unix_fd_message_new : no return generator

// Unsupported : g_unix_fd_message_new_with_fd_list : unsupported parameter fd_list : no type generator for UnixFDList, GUnixFDList*

// Unsupported : g_unix_fd_message_append_fd : no return generator

// Unsupported : g_unix_fd_message_get_fd_list : no return generator

// Unsupported : g_unix_fd_message_steal_fds : unsupported parameter length : no type generator for gint, gint*

// UnixInputStream is a wrapper around the C record GUnixInputStream.
type UnixInputStream struct {
	native *C.GUnixInputStream
	// parent_instance : no type generator for InputStream, GInputStream
	// Private : priv
}

func UnixInputStreamNewFromC(u unsafe.Pointer) *UnixInputStream {
	c := (*C.GUnixInputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStream{native: c}

	return g
}

func (recv *UnixInputStream) toC() *C.GUnixInputStream {

	return recv.native
}

// Unsupported : g_unix_input_stream_new : unsupported parameter close_fd : no type generator for gboolean, gboolean

// Unsupported : g_unix_input_stream_get_close_fd : no return generator

// Unsupported : g_unix_input_stream_set_close_fd : unsupported parameter close_fd : no type generator for gboolean, gboolean

// UnixMountMonitor is a wrapper around the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native *C.GUnixMountMonitor
}

func UnixMountMonitorNewFromC(u unsafe.Pointer) *UnixMountMonitor {
	c := (*C.GUnixMountMonitor)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitor{native: c}

	return g
}

func (recv *UnixMountMonitor) toC() *C.GUnixMountMonitor {

	return recv.native
}

// Unsupported : g_unix_mount_monitor_new : no return generator

// Unsupported : g_unix_mount_monitor_set_rate_limit : no return generator

// UnixOutputStream is a wrapper around the C record GUnixOutputStream.
type UnixOutputStream struct {
	native *C.GUnixOutputStream
	// parent_instance : no type generator for OutputStream, GOutputStream
	// Private : priv
}

func UnixOutputStreamNewFromC(u unsafe.Pointer) *UnixOutputStream {
	c := (*C.GUnixOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStream{native: c}

	return g
}

func (recv *UnixOutputStream) toC() *C.GUnixOutputStream {

	return recv.native
}

// Unsupported : g_unix_output_stream_new : unsupported parameter close_fd : no type generator for gboolean, gboolean

// Unsupported : g_unix_output_stream_get_close_fd : no return generator

// Unsupported : g_unix_output_stream_set_close_fd : unsupported parameter close_fd : no type generator for gboolean, gboolean

// UnixSocketAddress is a wrapper around the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native *C.GUnixSocketAddress
	// parent_instance : no type generator for SocketAddress, GSocketAddress
	// Private : priv
}

func UnixSocketAddressNewFromC(u unsafe.Pointer) *UnixSocketAddress {
	c := (*C.GUnixSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddress{native: c}

	return g
}

func (recv *UnixSocketAddress) toC() *C.GUnixSocketAddress {

	return recv.native
}

// Unsupported : g_unix_socket_address_new : no return generator

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_get_is_abstract : no return generator

// Vfs is a wrapper around the C record GVfs.
type Vfs struct {
	native *C.GVfs
	// parent_instance : no type generator for GObject.Object, GObject
}

func VfsNewFromC(u unsafe.Pointer) *Vfs {
	c := (*C.GVfs)(u)
	if c == nil {
		return nil
	}

	g := &Vfs{native: c}

	return g
}

func (recv *Vfs) toC() *C.GVfs {

	return recv.native
}

// Unsupported : g_vfs_get_file_for_path : no return generator

// Unsupported : g_vfs_get_file_for_uri : no return generator

// Unsupported : g_vfs_get_supported_uri_schemes : no return type

// Unsupported : g_vfs_is_active : no return generator

// Unsupported : g_vfs_parse_name : no return generator

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc, GVfsFileLookupFunc

// Unsupported : g_vfs_unregister_uri_scheme : no return generator

// VolumeMonitor is a wrapper around the C record GVolumeMonitor.
type VolumeMonitor struct {
	native *C.GVolumeMonitor
	// parent_instance : no type generator for GObject.Object, GObject
	// Private : priv
}

func VolumeMonitorNewFromC(u unsafe.Pointer) *VolumeMonitor {
	c := (*C.GVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitor{native: c}

	return g
}

func (recv *VolumeMonitor) toC() *C.GVolumeMonitor {

	return recv.native
}

// GetConnectedDrives is a wrapper around the C function g_volume_monitor_get_connected_drives.
func (recv *VolumeMonitor) GetConnectedDrives() *glib.List {
	retC := C.g_volume_monitor_get_connected_drives((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_monitor_get_mount_for_uuid : no return generator

// GetMounts is a wrapper around the C function g_volume_monitor_get_mounts.
func (recv *VolumeMonitor) GetMounts() *glib.List {
	retC := C.g_volume_monitor_get_mounts((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_monitor_get_volume_for_uuid : no return generator

// GetVolumes is a wrapper around the C function g_volume_monitor_get_volumes.
func (recv *VolumeMonitor) GetVolumes() *glib.List {
	retC := C.g_volume_monitor_get_volumes((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ZlibCompressor is a wrapper around the C record GZlibCompressor.
type ZlibCompressor struct {
	native *C.GZlibCompressor
}

func ZlibCompressorNewFromC(u unsafe.Pointer) *ZlibCompressor {
	c := (*C.GZlibCompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressor{native: c}

	return g
}

func (recv *ZlibCompressor) toC() *C.GZlibCompressor {

	return recv.native
}

// Unsupported : g_zlib_compressor_new : no return generator

// Unsupported : g_zlib_compressor_get_file_info : no return generator

// Unsupported : g_zlib_compressor_set_file_info : unsupported parameter file_info : no type generator for FileInfo, GFileInfo*

// ZlibDecompressor is a wrapper around the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native *C.GZlibDecompressor
}

func ZlibDecompressorNewFromC(u unsafe.Pointer) *ZlibDecompressor {
	c := (*C.GZlibDecompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressor{native: c}

	return g
}

func (recv *ZlibDecompressor) toC() *C.GZlibDecompressor {

	return recv.native
}

// Unsupported : g_zlib_decompressor_new : no return generator

// Unsupported : g_zlib_decompressor_get_file_info : no return generator