// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// The virtual function table for #GListModel.
/*

C type

GListModelInterface
*/
type ListModelInterface struct {
	native *C.GListModelInterface
	// g_iface : record
	// no type for get_item_type
	// no type for get_n_items
	// no type for get_item
}

func ListModelInterfaceNewFromC(u unsafe.Pointer) *ListModelInterface {
	c := (*C.GListModelInterface)(u)
	if c == nil {
		return nil
	}

	g := &ListModelInterface{native: c}

	return g
}

func (recv *ListModelInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Structure used for scatter/gather data output when sending multiple
// messages or packets in one go. You generally pass in an array of
// #GOutputVectors and the operation will use all the buffers as if they
// were one buffer.
//
// If @address is %NULL then the message is sent to the default receiver
// (as previously set by g_socket_connect()).
/*

C type

GOutputMessage
*/
type OutputMessage struct {
	native *C.GOutputMessage
	// address : record
	// vectors : record
	NumVectors uint32
	BytesSent  uint32
	// no type for control_messages
	NumControlMessages uint32
}

func OutputMessageNewFromC(u unsafe.Pointer) *OutputMessage {
	c := (*C.GOutputMessage)(u)
	if c == nil {
		return nil
	}

	g := &OutputMessage{
		BytesSent:          (uint32)(c.bytes_sent),
		NumControlMessages: (uint32)(c.num_control_messages),
		NumVectors:         (uint32)(c.num_vectors),
		native:             c,
	}

	return g
}

func (recv *OutputMessage) ToC() unsafe.Pointer {
	recv.native.num_vectors =
		(C.guint)(recv.NumVectors)
	recv.native.bytes_sent =
		(C.guint)(recv.BytesSent)
	recv.native.num_control_messages =
		(C.guint)(recv.NumControlMessages)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_settings_schema_list_children : no return type

// Gets the name of @key.
/*

C function

g_settings_schema_key_get_name
*/
func (recv *SettingsSchemaKey) GetName() string {
	retC := C.g_settings_schema_key_get_name((*C.GSettingsSchemaKey)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
