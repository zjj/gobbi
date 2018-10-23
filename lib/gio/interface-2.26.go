// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Proxy is a wrapper around the C record GProxy.
type Proxy struct {
	native *C.GProxy
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	c := (*C.GProxy)(u)
	if c == nil {
		return nil
	}

	g := &Proxy{native: c}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Connect is a wrapper around the C function g_proxy_connect.
func (recv *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable) (*IOStream, error) {
	c_connection := (*C.GIOStream)(connection.ToC())

	c_proxy_address := (*C.GProxyAddress)(proxyAddress.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_proxy_connect((*C.GProxy)(recv.native), c_connection, c_proxy_address, c_cancellable, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_proxy_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_proxy_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SupportsHostname is a wrapper around the C function g_proxy_supports_hostname.
func (recv *Proxy) SupportsHostname() bool {
	retC := C.g_proxy_supports_hostname((*C.GProxy)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ProxyResolver is a wrapper around the C record GProxyResolver.
type ProxyResolver struct {
	native *C.GProxyResolver
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	c := (*C.GProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolver{native: c}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IsSupported is a wrapper around the C function g_proxy_resolver_is_supported.
func (recv *ProxyResolver) IsSupported() bool {
	retC := C.g_proxy_resolver_is_supported((*C.GProxyResolver)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_proxy_resolver_lookup : no return type

// Unsupported : g_proxy_resolver_lookup_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_proxy_resolver_lookup_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ProxyEnumerate is a wrapper around the C function g_socket_connectable_proxy_enumerate.
func (recv *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_proxy_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}