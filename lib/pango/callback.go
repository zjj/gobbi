// This is a generated file - DO NOT EDIT

package pango

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
/*

	gboolean callback_attrfilterfuncHandler(GObject *, PangoAttribute*, gpointer, gpointer);

*/
import "C"

// Unsupported : callback AttrDataCopyFunc : no return generator

var callbackAttrfilterfuncId int
var callbackAttrfilterfuncMap = make(map[int]AttrfilterfuncCallback)
var callbackAttrfilterfuncLock sync.RWMutex

// AttrfilterfuncCallback is a callback function for a 'AttrFilterFunc' callback.
type AttrfilterfuncCallback func(attribute *Attribute) bool

//export callback_attrfilterfuncHandler
func callback_attrfilterfuncHandler(_ *C.GObject, c_attribute *C.PangoAttribute, data C.gpointer) C.gboolean {
	callbackAttrfilterfuncLock.RLock()
	defer callbackAttrfilterfuncLock.RUnlock()

	attribute := AttributeNewFromC(unsafe.Pointer(c_attribute))

	index := int(uintptr(data))
	callback := callbackAttrfilterfuncMap[index].callback
	retGo := callback(attribute, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}
