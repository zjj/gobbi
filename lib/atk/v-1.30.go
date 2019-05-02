// Code generated - DO NOT EDIT.
// +build atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetId is a wrapper around the C function atk_plug_get_id.
func (recv *Plug) GetId() string {
	retC := C.atk_plug_get_id((*C.AtkPlug)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Embed is a wrapper around the C function atk_socket_embed.
func (recv *Socket) Embed(plugId string) {
	c_plug_id := C.CString(plugId)
	defer C.free(unsafe.Pointer(c_plug_id))

	C.atk_socket_embed((*C.AtkSocket)(recv.native), c_plug_id)

	return
}

// IsOccupied is a wrapper around the C function atk_socket_is_occupied.
func (recv *Socket) IsOccupied() bool {
	retC := C.atk_socket_is_occupied((*C.AtkSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
