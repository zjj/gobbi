// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Application is a wrapper around the C record GApplication.
type Application struct {
	native *C.GApplication
	// Private : parent_instance
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationNew is a wrapper around the C function g_application_new.
func ApplicationNew(applicationId string, flags ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.g_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function g_application_activate.
func (recv *Application) Activate() {
	C.g_application_activate((*C.GApplication)(recv.native))

	return
}

// GetApplicationId is a wrapper around the C function g_application_get_application_id.
func (recv *Application) GetApplicationId() string {
	retC := C.g_application_get_application_id((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_application_get_flags.
func (recv *Application) GetFlags() ApplicationFlags {
	retC := C.g_application_get_flags((*C.GApplication)(recv.native))
	retGo := (ApplicationFlags)(retC)

	return retGo
}

// GetInactivityTimeout is a wrapper around the C function g_application_get_inactivity_timeout.
func (recv *Application) GetInactivityTimeout() uint32 {
	retC := C.g_application_get_inactivity_timeout((*C.GApplication)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetIsRegistered is a wrapper around the C function g_application_get_is_registered.
func (recv *Application) GetIsRegistered() bool {
	retC := C.g_application_get_is_registered((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsRemote is a wrapper around the C function g_application_get_is_remote.
func (recv *Application) GetIsRemote() bool {
	retC := C.g_application_get_is_remote((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hold is a wrapper around the C function g_application_hold.
func (recv *Application) Hold() {
	C.g_application_hold((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_open : unsupported parameter files : no param type

// Register is a wrapper around the C function g_application_register.
func (recv *Application) Register(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_application_register((*C.GApplication)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Release is a wrapper around the C function g_application_release.
func (recv *Application) Release() {
	C.g_application_release((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_run : unsupported parameter argv : no param type

// Unsupported : g_application_set_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// SetApplicationId is a wrapper around the C function g_application_set_application_id.
func (recv *Application) SetApplicationId(applicationId string) {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	C.g_application_set_application_id((*C.GApplication)(recv.native), c_application_id)

	return
}

// SetFlags is a wrapper around the C function g_application_set_flags.
func (recv *Application) SetFlags(flags ApplicationFlags) {
	c_flags := (C.GApplicationFlags)(flags)

	C.g_application_set_flags((*C.GApplication)(recv.native), c_flags)

	return
}

// SetInactivityTimeout is a wrapper around the C function g_application_set_inactivity_timeout.
func (recv *Application) SetInactivityTimeout(inactivityTimeout uint32) {
	c_inactivity_timeout := (C.guint)(inactivityTimeout)

	C.g_application_set_inactivity_timeout((*C.GApplication)(recv.native), c_inactivity_timeout)

	return
}

// Unsupported : g_application_command_line_get_arguments : unsupported parameter argc : no type generator for gint, int*

// GetCwd is a wrapper around the C function g_application_command_line_get_cwd.
func (recv *ApplicationCommandLine) GetCwd() string {
	retC := C.g_application_command_line_get_cwd((*C.GApplicationCommandLine)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_get_environ : no return type

// GetExitStatus is a wrapper around the C function g_application_command_line_get_exit_status.
func (recv *ApplicationCommandLine) GetExitStatus() int32 {
	retC := C.g_application_command_line_get_exit_status((*C.GApplicationCommandLine)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIsRemote is a wrapper around the C function g_application_command_line_get_is_remote.
func (recv *ApplicationCommandLine) GetIsRemote() bool {
	retC := C.g_application_command_line_get_is_remote((*C.GApplicationCommandLine)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// Getenv is a wrapper around the C function g_application_command_line_getenv.
func (recv *ApplicationCommandLine) Getenv(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_application_command_line_getenv((*C.GApplicationCommandLine)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_print : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_printerr : unsupported parameter ... : varargs

// SetExitStatus is a wrapper around the C function g_application_command_line_set_exit_status.
func (recv *ApplicationCommandLine) SetExitStatus(exitStatus int32) {
	c_exit_status := (C.int)(exitStatus)

	C.g_application_command_line_set_exit_status((*C.GApplicationCommandLine)(recv.native), c_exit_status)

	return
}

// SourceNew is a wrapper around the C function g_cancellable_source_new.
func (recv *Cancellable) SourceNew() *glib.Source {
	retC := C.g_cancellable_source_new((*C.GCancellable)(recv.native))
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// ClearEmblems is a wrapper around the C function g_emblemed_icon_clear_emblems.
func (recv *EmblemedIcon) ClearEmblems() {
	C.g_emblemed_icon_clear_emblems((*C.GEmblemedIcon)(recv.native))

	return
}

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_io_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Unsupported : g_settings_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// SetEnabled is a wrapper around the C function g_simple_action_set_enabled.
func (recv *SimpleAction) SetEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.g_simple_action_set_enabled((*C.GSimpleAction)(recv.native), c_enabled)

	return
}

// SimpleActionGroup is a wrapper around the C record GSimpleActionGroup.
type SimpleActionGroup struct {
	native *C.GSimpleActionGroup
	// Private : parent_instance
	// Private : priv
}

func SimpleActionGroupNewFromC(u unsafe.Pointer) *SimpleActionGroup {
	c := (*C.GSimpleActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroup{native: c}

	return g
}

func (recv *SimpleActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleActionGroupNew is a wrapper around the C function g_simple_action_group_new.
func SimpleActionGroupNew() *SimpleActionGroup {
	retC := C.g_simple_action_group_new()
	retGo := SimpleActionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_simple_action_group_insert : unsupported parameter action : no type generator for Action, GAction*

// Unsupported : g_simple_action_group_lookup : no return generator

// Remove is a wrapper around the C function g_simple_action_group_remove.
func (recv *SimpleActionGroup) Remove(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_simple_action_group_remove((*C.GSimpleActionGroup)(recv.native), c_action_name)

	return
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// TakeError is a wrapper around the C function g_simple_async_result_take_error.
func (recv *SimpleAsyncResult) TakeError(error *glib.Error) {
	c_error := (*C.GError)(error.ToC())

	C.g_simple_async_result_take_error((*C.GSimpleAsyncResult)(recv.native), c_error)

	return
}

// GetTls is a wrapper around the C function g_socket_client_get_tls.
func (recv *SocketClient) GetTls() bool {
	retC := C.g_socket_client_get_tls((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTlsValidationFlags is a wrapper around the C function g_socket_client_get_tls_validation_flags.
func (recv *SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	retC := C.g_socket_client_get_tls_validation_flags((*C.GSocketClient)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// SetTls is a wrapper around the C function g_socket_client_set_tls.
func (recv *SocketClient) SetTls(tls bool) {
	c_tls :=
		boolToGboolean(tls)

	C.g_socket_client_set_tls((*C.GSocketClient)(recv.native), c_tls)

	return
}

// SetTlsValidationFlags is a wrapper around the C function g_socket_client_set_tls_validation_flags.
func (recv *SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_socket_client_set_tls_validation_flags((*C.GSocketClient)(recv.native), c_flags)

	return
}

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// TcpWrapperConnectionNew is a wrapper around the C function g_tcp_wrapper_connection_new.
func TcpWrapperConnectionNew(baseIoStream *IOStream, socket *Socket) *SocketConnection {
	c_base_io_stream := (*C.GIOStream)(baseIoStream.ToC())

	c_socket := (*C.GSocket)(socket.ToC())

	retC := C.g_tcp_wrapper_connection_new(c_base_io_stream, c_socket)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// TlsCertificate is a wrapper around the C record GTlsCertificate.
type TlsCertificate struct {
	native *C.GTlsCertificate
	// parent_instance : record
	// priv : record
}

func TlsCertificateNewFromC(u unsafe.Pointer) *TlsCertificate {
	c := (*C.GTlsCertificate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificate{native: c}

	return g
}

func (recv *TlsCertificate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsCertificateNewFromFile is a wrapper around the C function g_tls_certificate_new_from_file.
func TlsCertificateNewFromFile(file string) (*TlsCertificate, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_file(c_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// TlsCertificateNewFromFiles is a wrapper around the C function g_tls_certificate_new_from_files.
func TlsCertificateNewFromFiles(certFile string, keyFile string) (*TlsCertificate, error) {
	c_cert_file := C.CString(certFile)
	defer C.free(unsafe.Pointer(c_cert_file))

	c_key_file := C.CString(keyFile)
	defer C.free(unsafe.Pointer(c_key_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_files(c_cert_file, c_key_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// TlsCertificateNewFromPem is a wrapper around the C function g_tls_certificate_new_from_pem.
func TlsCertificateNewFromPem(data string, length int64) (*TlsCertificate, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gssize)(length)

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_pem(c_data, c_length, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetIssuer is a wrapper around the C function g_tls_certificate_get_issuer.
func (recv *TlsCertificate) GetIssuer() *TlsCertificate {
	retC := C.g_tls_certificate_get_issuer((*C.GTlsCertificate)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_tls_certificate_verify : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// TlsConnection is a wrapper around the C record GTlsConnection.
type TlsConnection struct {
	native *C.GTlsConnection
	// parent_instance : record
	// priv : record
}

func TlsConnectionNewFromC(u unsafe.Pointer) *TlsConnection {
	c := (*C.GTlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnection{native: c}

	return g
}

func (recv *TlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmitAcceptCertificate is a wrapper around the C function g_tls_connection_emit_accept_certificate.
func (recv *TlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	c_peer_cert := (*C.GTlsCertificate)(peerCert.ToC())

	c_errors := (C.GTlsCertificateFlags)(errors)

	retC := C.g_tls_connection_emit_accept_certificate((*C.GTlsConnection)(recv.native), c_peer_cert, c_errors)
	retGo := retC == C.TRUE

	return retGo
}

// GetCertificate is a wrapper around the C function g_tls_connection_get_certificate.
func (recv *TlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificate is a wrapper around the C function g_tls_connection_get_peer_certificate.
func (recv *TlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_peer_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificateErrors is a wrapper around the C function g_tls_connection_get_peer_certificate_errors.
func (recv *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_tls_connection_get_peer_certificate_errors((*C.GTlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// GetRehandshakeMode is a wrapper around the C function g_tls_connection_get_rehandshake_mode.
func (recv *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_tls_connection_get_rehandshake_mode((*C.GTlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// GetRequireCloseNotify is a wrapper around the C function g_tls_connection_get_require_close_notify.
func (recv *TlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_tls_connection_get_require_close_notify((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseSystemCertdb is a wrapper around the C function g_tls_connection_get_use_system_certdb.
func (recv *TlsConnection) GetUseSystemCertdb() bool {
	retC := C.g_tls_connection_get_use_system_certdb((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Handshake is a wrapper around the C function g_tls_connection_handshake.
func (recv *TlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_connection_handshake((*C.GTlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_connection_handshake_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SetCertificate is a wrapper around the C function g_tls_connection_set_certificate.
func (recv *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(certificate.ToC())

	C.g_tls_connection_set_certificate((*C.GTlsConnection)(recv.native), c_certificate)

	return
}

// SetRehandshakeMode is a wrapper around the C function g_tls_connection_set_rehandshake_mode.
func (recv *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_tls_connection_set_rehandshake_mode((*C.GTlsConnection)(recv.native), c_mode)

	return
}

// SetRequireCloseNotify is a wrapper around the C function g_tls_connection_set_require_close_notify.
func (recv *TlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_tls_connection_set_require_close_notify((*C.GTlsConnection)(recv.native), c_require_close_notify)

	return
}

// SetUseSystemCertdb is a wrapper around the C function g_tls_connection_set_use_system_certdb.
func (recv *TlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	c_use_system_certdb :=
		boolToGboolean(useSystemCertdb)

	C.g_tls_connection_set_use_system_certdb((*C.GTlsConnection)(recv.native), c_use_system_certdb)

	return
}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type
