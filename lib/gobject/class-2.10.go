// This is a generated file - DO NOT EDIT
// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import (
	"C"
	"unsafe"
)

// ForceFloating is a wrapper around the C function g_object_force_floating.
func (recv *Object) ForceFloating() {
	C.g_object_force_floating((*C.GObject)(recv.native))

	return
}

// IsFloating is a wrapper around the C function g_object_is_floating.
func (recv *Object) IsFloating() bool {
	retC := C.g_object_is_floating((C.gpointer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RefSink is a wrapper around the C function g_object_ref_sink.
func (recv *Object) RefSink() Object {
	retC := C.g_object_ref_sink((C.gpointer)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefSink is a wrapper around the C function g_param_spec_ref_sink.
func (recv *ParamSpec) RefSink() *ParamSpec {
	retC := C.g_param_spec_ref_sink((*C.GParamSpec)(recv.native))
	retGo := ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equals compares this ParamSpecGType with another ParamSpecGType, and returns true if they represent the same GObject.
func (recv *ParamSpecGType) Equals(other *ParamSpecGType) bool {
	return other.ToC() == recv.ToC()
}

// ParamSpec upcasts to *ParamSpec
func (recv *ParamSpecGType) ParamSpec() *ParamSpec {
	return ParamSpecNewFromC(unsafe.Pointer(recv.native))
}
