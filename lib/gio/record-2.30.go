// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// DBusInterfaceIface is a wrapper around the C record GDBusInterfaceIface.
type DBusInterfaceIface struct {
	native *C.GDBusInterfaceIface
	// parent_iface : record
	// no type for get_info
	// no type for get_object
	// no type for set_object
	// no type for dup_object
}

func DBusInterfaceIfaceNewFromC(u unsafe.Pointer) *DBusInterfaceIface {
	c := (*C.GDBusInterfaceIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceIface{native: c}

	return g
}

func (recv *DBusInterfaceIface) toC() *C.GDBusInterfaceIface {

	return recv.native
}

// DBusInterfaceSkeletonClass is a wrapper around the C record GDBusInterfaceSkeletonClass.
type DBusInterfaceSkeletonClass struct {
	native *C.GDBusInterfaceSkeletonClass
	// parent_class : record
	// no type for get_info
	// no type for get_vtable
	// no type for get_properties
	// no type for flush
	// Private : vfunc_padding
	// no type for g_authorize_method
	// Private : signal_padding
}

func DBusInterfaceSkeletonClassNewFromC(u unsafe.Pointer) *DBusInterfaceSkeletonClass {
	c := (*C.GDBusInterfaceSkeletonClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeletonClass{native: c}

	return g
}

func (recv *DBusInterfaceSkeletonClass) toC() *C.GDBusInterfaceSkeletonClass {

	return recv.native
}

// DBusObjectIface is a wrapper around the C record GDBusObjectIface.
type DBusObjectIface struct {
	native *C.GDBusObjectIface
	// parent_iface : record
	// no type for get_object_path
	// no type for get_interfaces
	// no type for get_interface
	// no type for interface_added
	// no type for interface_removed
}

func DBusObjectIfaceNewFromC(u unsafe.Pointer) *DBusObjectIface {
	c := (*C.GDBusObjectIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectIface{native: c}

	return g
}

func (recv *DBusObjectIface) toC() *C.GDBusObjectIface {

	return recv.native
}

// DBusObjectManagerClientClass is a wrapper around the C record GDBusObjectManagerClientClass.
type DBusObjectManagerClientClass struct {
	native *C.GDBusObjectManagerClientClass
	// parent_class : record
	// no type for interface_proxy_signal
	// no type for interface_proxy_properties_changed
	// Private : padding
}

func DBusObjectManagerClientClassNewFromC(u unsafe.Pointer) *DBusObjectManagerClientClass {
	c := (*C.GDBusObjectManagerClientClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClientClass{native: c}

	return g
}

func (recv *DBusObjectManagerClientClass) toC() *C.GDBusObjectManagerClientClass {

	return recv.native
}

// DBusObjectManagerIface is a wrapper around the C record GDBusObjectManagerIface.
type DBusObjectManagerIface struct {
	native *C.GDBusObjectManagerIface
	// parent_iface : record
	// no type for get_object_path
	// no type for get_objects
	// no type for get_object
	// no type for get_interface
	// no type for object_added
	// no type for object_removed
	// no type for interface_added
	// no type for interface_removed
}

func DBusObjectManagerIfaceNewFromC(u unsafe.Pointer) *DBusObjectManagerIface {
	c := (*C.GDBusObjectManagerIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerIface{native: c}

	return g
}

func (recv *DBusObjectManagerIface) toC() *C.GDBusObjectManagerIface {

	return recv.native
}

// DBusObjectManagerServerClass is a wrapper around the C record GDBusObjectManagerServerClass.
type DBusObjectManagerServerClass struct {
	native *C.GDBusObjectManagerServerClass
	// parent_class : record
	// Private : padding
}

func DBusObjectManagerServerClassNewFromC(u unsafe.Pointer) *DBusObjectManagerServerClass {
	c := (*C.GDBusObjectManagerServerClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServerClass{native: c}

	return g
}

func (recv *DBusObjectManagerServerClass) toC() *C.GDBusObjectManagerServerClass {

	return recv.native
}

// DBusObjectProxyClass is a wrapper around the C record GDBusObjectProxyClass.
type DBusObjectProxyClass struct {
	native *C.GDBusObjectProxyClass
	// parent_class : record
	// Private : padding
}

func DBusObjectProxyClassNewFromC(u unsafe.Pointer) *DBusObjectProxyClass {
	c := (*C.GDBusObjectProxyClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxyClass{native: c}

	return g
}

func (recv *DBusObjectProxyClass) toC() *C.GDBusObjectProxyClass {

	return recv.native
}

// DBusObjectSkeletonClass is a wrapper around the C record GDBusObjectSkeletonClass.
type DBusObjectSkeletonClass struct {
	native *C.GDBusObjectSkeletonClass
	// parent_class : record
	// no type for authorize_method
	// Private : padding
}

func DBusObjectSkeletonClassNewFromC(u unsafe.Pointer) *DBusObjectSkeletonClass {
	c := (*C.GDBusObjectSkeletonClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeletonClass{native: c}

	return g
}

func (recv *DBusObjectSkeletonClass) toC() *C.GDBusObjectSkeletonClass {

	return recv.native
}

// IOModuleScope is a wrapper around the C record GIOModuleScope.
type IOModuleScope struct {
	native *C.GIOModuleScope
}

func IOModuleScopeNewFromC(u unsafe.Pointer) *IOModuleScope {
	c := (*C.GIOModuleScope)(u)
	if c == nil {
		return nil
	}

	g := &IOModuleScope{native: c}

	return g
}

func (recv *IOModuleScope) toC() *C.GIOModuleScope {

	return recv.native
}

// Unsupported : g_io_module_scope_block : no return generator

// Unsupported : g_io_module_scope_free : no return generator

// TlsDatabaseClass is a wrapper around the C record GTlsDatabaseClass.
type TlsDatabaseClass struct {
	native *C.GTlsDatabaseClass
	// parent_class : record
	// no type for verify_chain
	// no type for verify_chain_async
	// no type for verify_chain_finish
	// no type for create_certificate_handle
	// no type for lookup_certificate_for_handle
	// no type for lookup_certificate_for_handle_async
	// no type for lookup_certificate_for_handle_finish
	// no type for lookup_certificate_issuer
	// no type for lookup_certificate_issuer_async
	// no type for lookup_certificate_issuer_finish
	// no type for lookup_certificates_issued_by
	// no type for lookup_certificates_issued_by_async
	// no type for lookup_certificates_issued_by_finish
	// Private : padding
}

func TlsDatabaseClassNewFromC(u unsafe.Pointer) *TlsDatabaseClass {
	c := (*C.GTlsDatabaseClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabaseClass{native: c}

	return g
}

func (recv *TlsDatabaseClass) toC() *C.GTlsDatabaseClass {

	return recv.native
}

// TlsInteractionClass is a wrapper around the C record GTlsInteractionClass.
type TlsInteractionClass struct {
	native *C.GTlsInteractionClass
	// Private : parent_class
	// no type for ask_password
	// no type for ask_password_async
	// no type for ask_password_finish
	// no type for request_certificate
	// no type for request_certificate_async
	// no type for request_certificate_finish
	// Private : padding
}

func TlsInteractionClassNewFromC(u unsafe.Pointer) *TlsInteractionClass {
	c := (*C.GTlsInteractionClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteractionClass{native: c}

	return g
}

func (recv *TlsInteractionClass) toC() *C.GTlsInteractionClass {

	return recv.native
}