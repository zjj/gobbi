// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// ChecksumTypeGetLength is a wrapper around the C function g_checksum_type_get_length.
func ChecksumTypeGetLength(checksumType ChecksumType) int64 {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_type_get_length(c_checksum_type)
	retGo := (int64)(retC)

	return retGo
}

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_clear_handle_id : unsupported parameter clear_func : no type generator for ClearHandleFunc (GClearHandleFunc) for param clear_func

// Unsupported : g_clear_pointer : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_compute_checksum_for_data : unsupported parameter data : no type generator for guint8 () for array param data

// ComputeChecksumForString is a wrapper around the C function g_compute_checksum_for_string.
func ComputeChecksumForString(checksumType ChecksumType, str string, length int64) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(length)

	retC := C.g_compute_checksum_for_string(c_checksum_type, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	retC := C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_environ_getenv : unsupported parameter envp :

// Unsupported : g_environ_setenv : unsupported parameter envp :

// Unsupported : g_environ_unsetenv : unsupported parameter envp :

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_set_contents : unsupported parameter contents : no type generator for guint8 () for array param contents

// FormatSizeForDisplay is a wrapper around the C function g_format_size_for_display.
func FormatSizeForDisplay(size uint64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_locale_variants : no return type

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

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

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

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

// TestBug is a wrapper around the C function g_test_bug.
func TestBug(bugUriSnippet string) {
	c_bug_uri_snippet := C.CString(bugUriSnippet)
	defer C.free(unsafe.Pointer(c_bug_uri_snippet))

	C.g_test_bug(c_bug_uri_snippet)

	return
}

// TestBugBase is a wrapper around the C function g_test_bug_base.
func TestBugBase(uriPattern string) {
	c_uri_pattern := C.CString(uriPattern)
	defer C.free(unsafe.Pointer(c_uri_pattern))

	C.g_test_bug_base(c_uri_pattern)

	return
}

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	retC := C.g_test_create_suite(c_suite_name)
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() *TestSuite {
	retC := C.g_test_get_root()
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_init : unsupported parameter argv : in string with indirection level of 3

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// TestQueueFree is a wrapper around the C function g_test_queue_free.
func TestQueueFree(gfreePointer uintptr) {
	c_gfree_pointer := (C.gpointer)(gfreePointer)

	C.g_test_queue_free(c_gfree_pointer)

	return
}

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo := (float64)(retC)

	return retGo
}

// TestRandDoubleRange is a wrapper around the C function g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo := (float64)(retC)

	return retGo
}

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo := (int32)(retC)

	return retGo
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() int32 {
	retC := C.g_test_run()
	retGo := (int32)(retC)

	return retGo
}

// TestRunSuite is a wrapper around the C function g_test_run_suite.
func TestRunSuite(suite *TestSuite) int32 {
	c_suite := (*C.GTestSuite)(suite.ToC())

	retC := C.g_test_run_suite(c_suite)
	retGo := (int32)(retC)

	return retGo
}

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerStart is a wrapper around the C function g_test_timer_start.
func TestTimerStart() {
	C.g_test_timer_start()

	return
}

// TestTrapFork is a wrapper around the C function g_test_trap_fork.
func TestTrapFork(usecTimeout uint64, testTrapFlags TestTrapFlags) bool {
	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_trap_flags := (C.GTestTrapFlags)(testTrapFlags)

	retC := C.g_test_trap_fork(c_usec_timeout, c_test_trap_flags)
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapHasPassed is a wrapper around the C function g_test_trap_has_passed.
func TestTrapHasPassed() bool {
	retC := C.g_test_trap_has_passed()
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapReachedTimeout is a wrapper around the C function g_test_trap_reached_timeout.
func TestTrapReachedTimeout() bool {
	retC := C.g_test_trap_reached_timeout()
	retGo := retC == C.TRUE

	return retGo
}

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

// UriEscapeString is a wrapper around the C function g_uri_escape_string.
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) string {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_uri_escape_string(c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_uri_list_extract_uris : no return type

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_escaped_string_end := C.CString(escapedStringEnd)
	defer C.free(unsafe.Pointer(c_escaped_string_end))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_segment(c_escaped_string, c_escaped_string_end, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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
