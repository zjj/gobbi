// This is a generated file - DO NOT EDIT
// +build gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "C"

// TypeEnsure is a wrapper around the C function g_type_ensure.
func TypeEnsure(type_ Type) {
	c_type := (C.GType)(type_)

	C.g_type_ensure(c_type)

	return
}
