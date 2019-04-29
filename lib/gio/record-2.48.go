// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// DatagramBasedInterface is a wrapper around the C record GDatagramBasedInterface.
type DatagramBasedInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for receive_messages
	// no type for send_messages
	// no type for create_source
	// no type for condition_check
	// no type for condition_wait
}

// DtlsClientConnectionInterface is a wrapper around the C record GDtlsClientConnectionInterface.
type DtlsClientConnectionInterface struct {
	native unsafe.Pointer
	// g_iface : record
}

// DtlsConnectionInterface is a wrapper around the C record GDtlsConnectionInterface.
type DtlsConnectionInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// no type for shutdown
	// no type for shutdown_async
	// no type for shutdown_finish
}

// DtlsServerConnectionInterface is a wrapper around the C record GDtlsServerConnectionInterface.
type DtlsServerConnectionInterface struct {
	native unsafe.Pointer
	// g_iface : record
}

// InputMessage is a wrapper around the C record GInputMessage.
type InputMessage struct {
	native unsafe.Pointer
	// address : record
	// no type for vectors
	NumVectors    uint32
	BytesReceived uint64
	Flags         int32
	// no type for control_messages
	// num_control_messages : guint* with indirection level of 1
}
