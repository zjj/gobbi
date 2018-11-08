// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_assertion_message_cmpnum : unsupported parameter numtype : no type generator for gchar (char) for param numtype

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_base64_decode : no return type

// Unsupported : g_base64_decode_inplace : unsupported parameter text : no type generator for guint8 () for array param text

// Unsupported : g_base64_decode_step : unsupported parameter in : no type generator for guint8 () for array param in

// Unsupported : g_base64_encode : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_base64_encode_close : unsupported parameter out : output array param out

// Unsupported : g_base64_encode_step : unsupported parameter in : no type generator for guint8 () for array param in

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_filename_valist : unsupported parameter args : no type generator for va_list (va_list*) for param args

// Unsupported : g_build_filenamev : unsupported parameter args :

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Unsupported : g_build_pathv : unsupported parameter args :

// Unsupported : g_byte_array_new : no return type

// Unsupported : g_byte_array_new_take : no return type

// ByteArrayUnref is a wrapper around the C function g_byte_array_unref.
func ByteArrayUnref(array []uint8) {
	c_array := &array[0]

	C.g_byte_array_unref((*C.GByteArray)(unsafe.Pointer(c_array)))

	return
}

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_clear_handle_id : unsupported parameter clear_func : no type generator for ClearHandleFunc (GClearHandleFunc) for param clear_func

// Unsupported : g_clear_pointer : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_compute_checksum_for_data : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_convert : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_convert_with_fallback : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_convert_with_iconv : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_datalist_clear : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_foreach : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_dup_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_replace_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_init : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_set_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_unset_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_dataset_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Unsupported : g_dataset_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// DoubleEqual is a wrapper around the C function g_double_equal.
func DoubleEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_double_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// DoubleHash is a wrapper around the C function g_double_hash.
func DoubleHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_double_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_environ_getenv : unsupported parameter envp :

// Unsupported : g_environ_setenv : unsupported parameter envp :

// Unsupported : g_environ_unsetenv : unsupported parameter envp :

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_set_contents : unsupported parameter contents : no type generator for guint8 () for array param contents

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_locale_variants : no return type

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// HostnameIsAsciiEncoded is a wrapper around the C function g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ascii_encoded(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsIpAddress is a wrapper around the C function g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ip_address(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsNonAscii is a wrapper around the C function g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_non_ascii(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_ascii(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_unicode(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Int64Equal is a wrapper around the C function g_int64_equal.
func Int64Equal(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_int64_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Int64Hash is a wrapper around the C function g_int64_hash.
func Int64Hash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int64_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_io_add_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_add_watch_full : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_create_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_listenv : no return type

// Unsupported : g_locale_from_utf8 : no return type

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no type generator for guint8 () for array param opsysstring

// Unsupported : g_log : unsupported parameter ... : varargs

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_set_handler_full : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_log_set_writer_func : unsupported parameter func : no type generator for LogWriterFunc (GLogWriterFunc) for param func

// Unsupported : g_log_structured : unsupported parameter ... : varargs

// Unsupported : g_log_structured_array : unsupported parameter fields :

// Unsupported : g_log_structured_standard : unsupported parameter ... : varargs

// Unsupported : g_log_variant : unsupported parameter fields : Blacklisted record : GVariant

// Unsupported : g_log_writer_default : unsupported parameter fields :

// Unsupported : g_log_writer_format_fields : unsupported parameter fields :

// Unsupported : g_log_writer_journald : unsupported parameter fields :

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields :

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	retC := C.g_main_context_get_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// MkstempFull is a wrapper around the C function g_mkstemp_full.
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_flags := (C.gint)(flags)

	c_mode := (C.gint)(mode)

	retC := C.g_mkstemp_full(c_tmpl, c_flags, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_prefix_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_propagate_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_ptr_array_find_with_equal_func : unsupported parameter equal_func : no type generator for EqualFunc (GEqualFunc) for param equal_func

// Unsupported : g_qsort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// Unsupported : g_regex_escape_string : unsupported parameter string :

// Unsupported : g_regex_split_simple : no return type

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()

	return
}

// Unsupported : g_set_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_set_error_literal : unsupported parameter err : record with indirection level of 2

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_shell_parse_argv : unsupported parameter argvp : output array param argvp

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// Unsupported : g_spawn_async : unsupported parameter argv :

// Unsupported : g_spawn_async_with_pipes : unsupported parameter argv :

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

// Unsupported : g_spawn_sync : unsupported parameter argv :

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : output array param ascii_alternates

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strsplit : no return type

// Unsupported : g_strsplit_set : no return type

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Unsupported : g_test_add_data_func_full : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// Unsupported : g_test_add_vtable : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// Unsupported : g_test_init : unsupported parameter argv : in string with indirection level of 3

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_ucs4_to_utf16 : no return generator

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_uri_list_extract_uris : no return type

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no type generator for guint16 (const gunichar2*) for param str

// Unsupported : g_utf8_to_utf16 : no return generator

// Unsupported : g_utf8_validate : unsupported parameter str : no type generator for guint8 () for array param str

// Unsupported : g_variant_parse : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_type_checked_ : return type : Blacklisted record : GVariantType

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// Unsupported : g_vfprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsprintf : unsupported parameter args : no type generator for va_list (va_list) for param args
