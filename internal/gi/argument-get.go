package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func (a Argument) Int8() int8 {
	return (int8)(*(*C.gint8)(unsafe.Pointer(&a)))
}

func (a Argument) Int16() int16 {
	return (int16)(*(*C.gint16)(unsafe.Pointer(&a)))
}

func (a Argument) Int32() int32 {
	return (int32)(*(*C.gint)(unsafe.Pointer(&a)))
}

func (a Argument) Int64() int64 {
	return (int64)(*(*C.gint)(unsafe.Pointer(&a)))
}

func (a Argument) Uint8() uint8 {
	return (uint8)(*(*C.guint)(unsafe.Pointer(&a)))
}

func (a Argument) Uint16() uint16 {
	return (uint16)(*(*C.guint)(unsafe.Pointer(&a)))
}

func (a Argument) Uint32() uint32 {
	return (uint32)(*(*C.guint)(unsafe.Pointer(&a)))
}

func (a Argument) Uint64() uint64 {
	return (uint64)(*(*C.guint)(unsafe.Pointer(&a)))
}

func (a Argument) String(transferOwnership bool) string {
	cString := *(**C.gchar)(unsafe.Pointer(&a))
	if transferOwnership {
		defer C.free(unsafe.Pointer(cString))
	}

	goString := C.GoString(cString)
	return goString
}