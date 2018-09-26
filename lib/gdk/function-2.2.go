// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter dest_window : record with indirection level of 2

// GetDisplayArgName is a wrapper around the C function gdk_get_display_arg_name.
func GetDisplayArgName() string {
	retC := C.gdk_get_display_arg_name()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_notify_startup_complete : no return generator

// PangoContextGetForScreen is a wrapper around the C function gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	c_screen := (*C.GdkScreen)(screen.ToC())

	retC := C.gdk_pango_context_get_for_screen(c_screen)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_parse_args : unsupported parameter argc : no type generator for gint, gint*

// SelectionOwnerGetForDisplay is a wrapper around the C function gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection *Atom) *Window {
	c_display := (*C.GdkDisplay)(display.ToC())

	c_selection := (C.GdkAtom)(selection.ToC())

	retC := C.gdk_selection_owner_get_for_display(c_display, c_selection)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SelectionOwnerSetForDisplay is a wrapper around the C function gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	c_display := (*C.GdkDisplay)(display.ToC())

	c_owner := (*C.GdkWindow)(owner.ToC())

	c_selection := (C.GdkAtom)(selection.ToC())

	c_time_ := (C.guint32)(time)

	c_send_event :=
		boolToGboolean(sendEvent)

	retC := C.gdk_selection_owner_set_for_display(c_display, c_owner, c_selection, c_time_, c_send_event)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_selection_send_notify_for_display : no return generator

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter text : no param type
