// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
/*

	void dbusobject_interfaceAddedHandler(GObject *, GDBusInterface *, gpointer);

	static gulong DBusObject_signal_connect_interface_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "interface-added", G_CALLBACK(dbusobject_interfaceAddedHandler), data);
	}

*/
/*

	void dbusobject_interfaceRemovedHandler(GObject *, GDBusInterface *, gpointer);

	static gulong DBusObject_signal_connect_interface_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "interface-removed", G_CALLBACK(dbusobject_interfaceRemovedHandler), data);
	}

*/
/*

	void dbusobjectmanager_interfaceAddedHandler(GObject *, GDBusObject *, GDBusInterface *, gpointer);

	static gulong DBusObjectManager_signal_connect_interface_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "interface-added", G_CALLBACK(dbusobjectmanager_interfaceAddedHandler), data);
	}

*/
/*

	void dbusobjectmanager_interfaceRemovedHandler(GObject *, GDBusObject *, GDBusInterface *, gpointer);

	static gulong DBusObjectManager_signal_connect_interface_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "interface-removed", G_CALLBACK(dbusobjectmanager_interfaceRemovedHandler), data);
	}

*/
/*

	void dbusobjectmanager_objectAddedHandler(GObject *, GDBusObject *, gpointer);

	static gulong DBusObjectManager_signal_connect_object_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "object-added", G_CALLBACK(dbusobjectmanager_objectAddedHandler), data);
	}

*/
/*

	void dbusobjectmanager_objectRemovedHandler(GObject *, GDBusObject *, gpointer);

	static gulong DBusObjectManager_signal_connect_object_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "object-removed", G_CALLBACK(dbusobjectmanager_objectRemovedHandler), data);
	}

*/
import "C"

// Unsupported : g_action_change_state : unsupported parameter value : Blacklisted record : GVariant

// The #GDBusInterface type is the base type for D-Bus interfaces both
// on the service side (see #GDBusInterfaceSkeleton) and client side
// (see #GDBusProxy).
/*

C type

GDBusInterface
*/
type DBusInterface struct {
	native *C.GDBusInterface
}

func DBusInterfaceNewFromC(u unsafe.Pointer) *DBusInterface {
	c := (*C.GDBusInterface)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterface{native: c}

	return g
}

func (recv *DBusInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets D-Bus introspection information for the D-Bus interface
// implemented by @interface_.
/*

C function

g_dbus_interface_get_info
*/
func (recv *DBusInterface) GetInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_get_info((*C.GDBusInterface)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GDBusObject that @interface_ belongs to, if any.
//
// It is not safe to use the returned object if @interface_ or
// the returned object is being used from other threads. See
// g_dbus_interface_dup_object() for a thread-safe alternative.
/*

C function

g_dbus_interface_get_object
*/
func (recv *DBusInterface) GetObject() *DBusObject {
	retC := C.g_dbus_interface_get_object((*C.GDBusInterface)(recv.native))
	retGo := DBusObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the #GDBusObject for @interface_ to @object.
//
// Note that @interface_ will hold a weak reference to @object.
/*

C function

g_dbus_interface_set_object
*/
func (recv *DBusInterface) SetObject(object *DBusObject) {
	c_object := (*C.GDBusObject)(object.ToC())

	C.g_dbus_interface_set_object((*C.GDBusInterface)(recv.native), c_object)

	return
}

type signalDBusObjectInterfaceAddedDetail struct {
	callback  DBusObjectSignalInterfaceAddedCallback
	handlerID C.gulong
}

var signalDBusObjectInterfaceAddedId int
var signalDBusObjectInterfaceAddedMap = make(map[int]signalDBusObjectInterfaceAddedDetail)
var signalDBusObjectInterfaceAddedLock sync.Mutex

// DBusObjectSignalInterfaceAddedCallback is a callback function for a 'interface-added' signal emitted from a DBusObject.
type DBusObjectSignalInterfaceAddedCallback func(interface_ *DBusInterface)

/*
ConnectInterfaceAdded connects the callback to the 'interface-added' signal for the DBusObject.

The returned value represents the connection, and may be passed to DisconnectInterfaceAdded to remove it.
*/
func (recv *DBusObject) ConnectInterfaceAdded(callback DBusObjectSignalInterfaceAddedCallback) int {
	signalDBusObjectInterfaceAddedLock.Lock()
	defer signalDBusObjectInterfaceAddedLock.Unlock()

	signalDBusObjectInterfaceAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObject_signal_connect_interface_added(instance, C.gpointer(uintptr(signalDBusObjectInterfaceAddedId)))

	detail := signalDBusObjectInterfaceAddedDetail{callback, handlerID}
	signalDBusObjectInterfaceAddedMap[signalDBusObjectInterfaceAddedId] = detail

	return signalDBusObjectInterfaceAddedId
}

/*
DisconnectInterfaceAdded disconnects a callback from the 'interface-added' signal for the DBusObject.

The connectionID should be a value returned from a call to ConnectInterfaceAdded.
*/
func (recv *DBusObject) DisconnectInterfaceAdded(connectionID int) {
	signalDBusObjectInterfaceAddedLock.Lock()
	defer signalDBusObjectInterfaceAddedLock.Unlock()

	detail, exists := signalDBusObjectInterfaceAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectInterfaceAddedMap, connectionID)
}

//export dbusobject_interfaceAddedHandler
func dbusobject_interfaceAddedHandler(_ *C.GObject, c_interface *C.GDBusInterface, data C.gpointer) {
	interface_ := DBusInterfaceNewFromC(unsafe.Pointer(c_interface))

	index := int(uintptr(data))
	callback := signalDBusObjectInterfaceAddedMap[index].callback
	callback(interface_)
}

type signalDBusObjectInterfaceRemovedDetail struct {
	callback  DBusObjectSignalInterfaceRemovedCallback
	handlerID C.gulong
}

var signalDBusObjectInterfaceRemovedId int
var signalDBusObjectInterfaceRemovedMap = make(map[int]signalDBusObjectInterfaceRemovedDetail)
var signalDBusObjectInterfaceRemovedLock sync.Mutex

// DBusObjectSignalInterfaceRemovedCallback is a callback function for a 'interface-removed' signal emitted from a DBusObject.
type DBusObjectSignalInterfaceRemovedCallback func(interface_ *DBusInterface)

/*
ConnectInterfaceRemoved connects the callback to the 'interface-removed' signal for the DBusObject.

The returned value represents the connection, and may be passed to DisconnectInterfaceRemoved to remove it.
*/
func (recv *DBusObject) ConnectInterfaceRemoved(callback DBusObjectSignalInterfaceRemovedCallback) int {
	signalDBusObjectInterfaceRemovedLock.Lock()
	defer signalDBusObjectInterfaceRemovedLock.Unlock()

	signalDBusObjectInterfaceRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObject_signal_connect_interface_removed(instance, C.gpointer(uintptr(signalDBusObjectInterfaceRemovedId)))

	detail := signalDBusObjectInterfaceRemovedDetail{callback, handlerID}
	signalDBusObjectInterfaceRemovedMap[signalDBusObjectInterfaceRemovedId] = detail

	return signalDBusObjectInterfaceRemovedId
}

/*
DisconnectInterfaceRemoved disconnects a callback from the 'interface-removed' signal for the DBusObject.

The connectionID should be a value returned from a call to ConnectInterfaceRemoved.
*/
func (recv *DBusObject) DisconnectInterfaceRemoved(connectionID int) {
	signalDBusObjectInterfaceRemovedLock.Lock()
	defer signalDBusObjectInterfaceRemovedLock.Unlock()

	detail, exists := signalDBusObjectInterfaceRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectInterfaceRemovedMap, connectionID)
}

//export dbusobject_interfaceRemovedHandler
func dbusobject_interfaceRemovedHandler(_ *C.GObject, c_interface *C.GDBusInterface, data C.gpointer) {
	interface_ := DBusInterfaceNewFromC(unsafe.Pointer(c_interface))

	index := int(uintptr(data))
	callback := signalDBusObjectInterfaceRemovedMap[index].callback
	callback(interface_)
}

// Gets the D-Bus interface with name @interface_name associated with
// @object, if any.
/*

C function

g_dbus_object_get_interface
*/
func (recv *DBusObject) GetInterface(interfaceName string) *DBusInterface {
	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	retC := C.g_dbus_object_get_interface((*C.GDBusObject)(recv.native), c_interface_name)
	retGo := DBusInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the D-Bus interfaces associated with @object.
/*

C function

g_dbus_object_get_interfaces
*/
func (recv *DBusObject) GetInterfaces() *glib.List {
	retC := C.g_dbus_object_get_interfaces((*C.GDBusObject)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the object path for @object.
/*

C function

g_dbus_object_get_object_path
*/
func (recv *DBusObject) GetObjectPath() string {
	retC := C.g_dbus_object_get_object_path((*C.GDBusObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

type signalDBusObjectManagerInterfaceAddedDetail struct {
	callback  DBusObjectManagerSignalInterfaceAddedCallback
	handlerID C.gulong
}

var signalDBusObjectManagerInterfaceAddedId int
var signalDBusObjectManagerInterfaceAddedMap = make(map[int]signalDBusObjectManagerInterfaceAddedDetail)
var signalDBusObjectManagerInterfaceAddedLock sync.Mutex

// DBusObjectManagerSignalInterfaceAddedCallback is a callback function for a 'interface-added' signal emitted from a DBusObjectManager.
type DBusObjectManagerSignalInterfaceAddedCallback func(object *DBusObject, interface_ *DBusInterface)

/*
ConnectInterfaceAdded connects the callback to the 'interface-added' signal for the DBusObjectManager.

The returned value represents the connection, and may be passed to DisconnectInterfaceAdded to remove it.
*/
func (recv *DBusObjectManager) ConnectInterfaceAdded(callback DBusObjectManagerSignalInterfaceAddedCallback) int {
	signalDBusObjectManagerInterfaceAddedLock.Lock()
	defer signalDBusObjectManagerInterfaceAddedLock.Unlock()

	signalDBusObjectManagerInterfaceAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectManager_signal_connect_interface_added(instance, C.gpointer(uintptr(signalDBusObjectManagerInterfaceAddedId)))

	detail := signalDBusObjectManagerInterfaceAddedDetail{callback, handlerID}
	signalDBusObjectManagerInterfaceAddedMap[signalDBusObjectManagerInterfaceAddedId] = detail

	return signalDBusObjectManagerInterfaceAddedId
}

/*
DisconnectInterfaceAdded disconnects a callback from the 'interface-added' signal for the DBusObjectManager.

The connectionID should be a value returned from a call to ConnectInterfaceAdded.
*/
func (recv *DBusObjectManager) DisconnectInterfaceAdded(connectionID int) {
	signalDBusObjectManagerInterfaceAddedLock.Lock()
	defer signalDBusObjectManagerInterfaceAddedLock.Unlock()

	detail, exists := signalDBusObjectManagerInterfaceAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectManagerInterfaceAddedMap, connectionID)
}

//export dbusobjectmanager_interfaceAddedHandler
func dbusobjectmanager_interfaceAddedHandler(_ *C.GObject, c_object *C.GDBusObject, c_interface *C.GDBusInterface, data C.gpointer) {
	object := DBusObjectNewFromC(unsafe.Pointer(c_object))

	interface_ := DBusInterfaceNewFromC(unsafe.Pointer(c_interface))

	index := int(uintptr(data))
	callback := signalDBusObjectManagerInterfaceAddedMap[index].callback
	callback(object, interface_)
}

type signalDBusObjectManagerInterfaceRemovedDetail struct {
	callback  DBusObjectManagerSignalInterfaceRemovedCallback
	handlerID C.gulong
}

var signalDBusObjectManagerInterfaceRemovedId int
var signalDBusObjectManagerInterfaceRemovedMap = make(map[int]signalDBusObjectManagerInterfaceRemovedDetail)
var signalDBusObjectManagerInterfaceRemovedLock sync.Mutex

// DBusObjectManagerSignalInterfaceRemovedCallback is a callback function for a 'interface-removed' signal emitted from a DBusObjectManager.
type DBusObjectManagerSignalInterfaceRemovedCallback func(object *DBusObject, interface_ *DBusInterface)

/*
ConnectInterfaceRemoved connects the callback to the 'interface-removed' signal for the DBusObjectManager.

The returned value represents the connection, and may be passed to DisconnectInterfaceRemoved to remove it.
*/
func (recv *DBusObjectManager) ConnectInterfaceRemoved(callback DBusObjectManagerSignalInterfaceRemovedCallback) int {
	signalDBusObjectManagerInterfaceRemovedLock.Lock()
	defer signalDBusObjectManagerInterfaceRemovedLock.Unlock()

	signalDBusObjectManagerInterfaceRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectManager_signal_connect_interface_removed(instance, C.gpointer(uintptr(signalDBusObjectManagerInterfaceRemovedId)))

	detail := signalDBusObjectManagerInterfaceRemovedDetail{callback, handlerID}
	signalDBusObjectManagerInterfaceRemovedMap[signalDBusObjectManagerInterfaceRemovedId] = detail

	return signalDBusObjectManagerInterfaceRemovedId
}

/*
DisconnectInterfaceRemoved disconnects a callback from the 'interface-removed' signal for the DBusObjectManager.

The connectionID should be a value returned from a call to ConnectInterfaceRemoved.
*/
func (recv *DBusObjectManager) DisconnectInterfaceRemoved(connectionID int) {
	signalDBusObjectManagerInterfaceRemovedLock.Lock()
	defer signalDBusObjectManagerInterfaceRemovedLock.Unlock()

	detail, exists := signalDBusObjectManagerInterfaceRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectManagerInterfaceRemovedMap, connectionID)
}

//export dbusobjectmanager_interfaceRemovedHandler
func dbusobjectmanager_interfaceRemovedHandler(_ *C.GObject, c_object *C.GDBusObject, c_interface *C.GDBusInterface, data C.gpointer) {
	object := DBusObjectNewFromC(unsafe.Pointer(c_object))

	interface_ := DBusInterfaceNewFromC(unsafe.Pointer(c_interface))

	index := int(uintptr(data))
	callback := signalDBusObjectManagerInterfaceRemovedMap[index].callback
	callback(object, interface_)
}

type signalDBusObjectManagerObjectAddedDetail struct {
	callback  DBusObjectManagerSignalObjectAddedCallback
	handlerID C.gulong
}

var signalDBusObjectManagerObjectAddedId int
var signalDBusObjectManagerObjectAddedMap = make(map[int]signalDBusObjectManagerObjectAddedDetail)
var signalDBusObjectManagerObjectAddedLock sync.Mutex

// DBusObjectManagerSignalObjectAddedCallback is a callback function for a 'object-added' signal emitted from a DBusObjectManager.
type DBusObjectManagerSignalObjectAddedCallback func(object *DBusObject)

/*
ConnectObjectAdded connects the callback to the 'object-added' signal for the DBusObjectManager.

The returned value represents the connection, and may be passed to DisconnectObjectAdded to remove it.
*/
func (recv *DBusObjectManager) ConnectObjectAdded(callback DBusObjectManagerSignalObjectAddedCallback) int {
	signalDBusObjectManagerObjectAddedLock.Lock()
	defer signalDBusObjectManagerObjectAddedLock.Unlock()

	signalDBusObjectManagerObjectAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectManager_signal_connect_object_added(instance, C.gpointer(uintptr(signalDBusObjectManagerObjectAddedId)))

	detail := signalDBusObjectManagerObjectAddedDetail{callback, handlerID}
	signalDBusObjectManagerObjectAddedMap[signalDBusObjectManagerObjectAddedId] = detail

	return signalDBusObjectManagerObjectAddedId
}

/*
DisconnectObjectAdded disconnects a callback from the 'object-added' signal for the DBusObjectManager.

The connectionID should be a value returned from a call to ConnectObjectAdded.
*/
func (recv *DBusObjectManager) DisconnectObjectAdded(connectionID int) {
	signalDBusObjectManagerObjectAddedLock.Lock()
	defer signalDBusObjectManagerObjectAddedLock.Unlock()

	detail, exists := signalDBusObjectManagerObjectAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectManagerObjectAddedMap, connectionID)
}

//export dbusobjectmanager_objectAddedHandler
func dbusobjectmanager_objectAddedHandler(_ *C.GObject, c_object *C.GDBusObject, data C.gpointer) {
	object := DBusObjectNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalDBusObjectManagerObjectAddedMap[index].callback
	callback(object)
}

type signalDBusObjectManagerObjectRemovedDetail struct {
	callback  DBusObjectManagerSignalObjectRemovedCallback
	handlerID C.gulong
}

var signalDBusObjectManagerObjectRemovedId int
var signalDBusObjectManagerObjectRemovedMap = make(map[int]signalDBusObjectManagerObjectRemovedDetail)
var signalDBusObjectManagerObjectRemovedLock sync.Mutex

// DBusObjectManagerSignalObjectRemovedCallback is a callback function for a 'object-removed' signal emitted from a DBusObjectManager.
type DBusObjectManagerSignalObjectRemovedCallback func(object *DBusObject)

/*
ConnectObjectRemoved connects the callback to the 'object-removed' signal for the DBusObjectManager.

The returned value represents the connection, and may be passed to DisconnectObjectRemoved to remove it.
*/
func (recv *DBusObjectManager) ConnectObjectRemoved(callback DBusObjectManagerSignalObjectRemovedCallback) int {
	signalDBusObjectManagerObjectRemovedLock.Lock()
	defer signalDBusObjectManagerObjectRemovedLock.Unlock()

	signalDBusObjectManagerObjectRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusObjectManager_signal_connect_object_removed(instance, C.gpointer(uintptr(signalDBusObjectManagerObjectRemovedId)))

	detail := signalDBusObjectManagerObjectRemovedDetail{callback, handlerID}
	signalDBusObjectManagerObjectRemovedMap[signalDBusObjectManagerObjectRemovedId] = detail

	return signalDBusObjectManagerObjectRemovedId
}

/*
DisconnectObjectRemoved disconnects a callback from the 'object-removed' signal for the DBusObjectManager.

The connectionID should be a value returned from a call to ConnectObjectRemoved.
*/
func (recv *DBusObjectManager) DisconnectObjectRemoved(connectionID int) {
	signalDBusObjectManagerObjectRemovedLock.Lock()
	defer signalDBusObjectManagerObjectRemovedLock.Unlock()

	detail, exists := signalDBusObjectManagerObjectRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusObjectManagerObjectRemovedMap, connectionID)
}

//export dbusobjectmanager_objectRemovedHandler
func dbusobjectmanager_objectRemovedHandler(_ *C.GObject, c_object *C.GDBusObject, data C.gpointer) {
	object := DBusObjectNewFromC(unsafe.Pointer(c_object))

	index := int(uintptr(data))
	callback := signalDBusObjectManagerObjectRemovedMap[index].callback
	callback(object)
}

// Gets the interface proxy for @interface_name at @object_path, if
// any.
/*

C function

g_dbus_object_manager_get_interface
*/
func (recv *DBusObjectManager) GetInterface(objectPath string, interfaceName string) *DBusInterface {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	retC := C.g_dbus_object_manager_get_interface((*C.GDBusObjectManager)(recv.native), c_object_path, c_interface_name)
	retGo := DBusInterfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GDBusObjectProxy at @object_path, if any.
/*

C function

g_dbus_object_manager_get_object
*/
func (recv *DBusObjectManager) GetObject(objectPath string) *DBusObject {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_manager_get_object((*C.GDBusObjectManager)(recv.native), c_object_path)
	retGo := DBusObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the object path that @manager is for.
/*

C function

g_dbus_object_manager_get_object_path
*/
func (recv *DBusObjectManager) GetObjectPath() string {
	retC := C.g_dbus_object_manager_get_object_path((*C.GDBusObjectManager)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets all #GDBusObject objects known to @manager.
/*

C function

g_dbus_object_manager_get_objects
*/
func (recv *DBusObjectManager) GetObjects() *glib.List {
	retC := C.g_dbus_object_manager_get_objects((*C.GDBusObjectManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the default #GTlsDatabase used to verify TLS connections.
/*

C function

g_tls_backend_get_default_database
*/
func (recv *TlsBackend) GetDefaultDatabase() *TlsDatabase {
	retC := C.g_tls_backend_get_default_database((*C.GTlsBackend)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GType of @backend's #GTlsFileDatabase implementation.
/*

C function

g_tls_backend_get_file_database_type
*/
func (recv *TlsBackend) GetFileDatabaseType() gobject.Type {
	retC := C.g_tls_backend_get_file_database_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// #GTlsFileDatabase is implemented by #GTlsDatabase objects which load
// their certificate information from a file. It is an interface which
// TLS library specific subtypes implement.
/*

C type

GTlsFileDatabase
*/
type TlsFileDatabase struct {
	native *C.GTlsFileDatabase
}

func TlsFileDatabaseNewFromC(u unsafe.Pointer) *TlsFileDatabase {
	c := (*C.GTlsFileDatabase)(u)
	if c == nil {
		return nil
	}

	g := &TlsFileDatabase{native: c}

	return g
}

func (recv *TlsFileDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
