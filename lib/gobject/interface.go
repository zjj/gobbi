// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypePlugin is a wrapper around the C record GTypePlugin.
type TypePlugin struct {
	native *C.GTypePlugin
}

func TypePluginNewFromC(u unsafe.Pointer) *TypePlugin {
	c := (*C.GTypePlugin)(u)
	if c == nil {
		return nil
	}

	g := &TypePlugin{native: c}

	return g
}

func (recv *TypePlugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Calls the @complete_interface_info function from the
// #GTypePluginClass of @plugin. There should be no need to use this
// function outside of the GObject type system itself.
/*

C function : g_type_plugin_complete_interface_info
*/
func (recv *TypePlugin) CompleteInterfaceInfo(instanceType Type, interfaceType Type, info *InterfaceInfo) {
	c_instance_type := (C.GType)(instanceType)

	c_interface_type := (C.GType)(interfaceType)

	c_info := (*C.GInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GInterfaceInfo)(info.ToC())
	}

	C.g_type_plugin_complete_interface_info((*C.GTypePlugin)(recv.native), c_instance_type, c_interface_type, c_info)

	return
}

// Calls the @complete_type_info function from the #GTypePluginClass of @plugin.
// There should be no need to use this function outside of the GObject
// type system itself.
/*

C function : g_type_plugin_complete_type_info
*/
func (recv *TypePlugin) CompleteTypeInfo(gType Type, info *TypeInfo, valueTable *TypeValueTable) {
	c_g_type := (C.GType)(gType)

	c_info := (*C.GTypeInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GTypeInfo)(info.ToC())
	}

	c_value_table := (*C.GTypeValueTable)(C.NULL)
	if valueTable != nil {
		c_value_table = (*C.GTypeValueTable)(valueTable.ToC())
	}

	C.g_type_plugin_complete_type_info((*C.GTypePlugin)(recv.native), c_g_type, c_info, c_value_table)

	return
}

// Calls the @unuse_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
/*

C function : g_type_plugin_unuse
*/
func (recv *TypePlugin) Unuse() {
	C.g_type_plugin_unuse((*C.GTypePlugin)(recv.native))

	return
}

// Calls the @use_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
/*

C function : g_type_plugin_use
*/
func (recv *TypePlugin) Use() {
	C.g_type_plugin_use((*C.GTypePlugin)(recv.native))

	return
}
