// This is a generated file - DO NOT EDIT
// +build atk_1.4 atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// IsSelectedLink is a wrapper around the C function atk_hyperlink_is_selected_link.
func (recv *Hyperlink) IsSelectedLink() bool {
	retC := C.atk_hyperlink_is_selected_link((*C.AtkHyperlink)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : atk_object_connect_property_change_handler : unsupported parameter handler : no type generator for PropertyChangeHandler, AtkPropertyChangeHandler*

// Unsupported : atk_object_factory_get_accessible_type : no return generator

// Unsupported : atk_registry_get_factory : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_get_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_registry_set_factory_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : atk_relation_new : unsupported parameter targets : no param type

// Unsupported : atk_relation_get_target : no return type

// Unsupported : atk_state_set_add_states : unsupported parameter types : no param type

// Unsupported : atk_state_set_contains_states : unsupported parameter types : no param type
