// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GLArea is a wrapper around the C record GtkGLArea.
type GLArea struct {
	native *C.GtkGLArea
	// Private : parent_instance
}

func GLAreaNewFromC(u unsafe.Pointer) *GLArea {
	c := (*C.GtkGLArea)(u)
	if c == nil {
		return nil
	}

	g := &GLArea{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GLArea) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GLArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GLAreaNew is a wrapper around the C function gtk_gl_area_new.
func GLAreaNew() *GLArea {
	retC := C.gtk_gl_area_new()
	retGo := GLAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ModelButtonNew is a wrapper around the C function gtk_model_button_new.
func ModelButtonNew() *ModelButton {
	retC := C.gtk_model_button_new()
	retGo := ModelButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopoverMenuNew is a wrapper around the C function gtk_popover_menu_new.
func PopoverMenuNew() *PopoverMenu {
	retC := C.gtk_popover_menu_new()
	retGo := PopoverMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StackSidebarNew is a wrapper around the C function gtk_stack_sidebar_new.
func StackSidebarNew() *StackSidebar {
	retC := C.gtk_stack_sidebar_new()
	retGo := StackSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GLAreaClass is a wrapper around the C record GtkGLAreaClass.
type GLAreaClass struct {
	native *C.GtkGLAreaClass
	// Private : parent_class
	// no type for render
	// no type for resize
	// no type for create_context
	// Private : _padding
}

func GLAreaClassNewFromC(u unsafe.Pointer) *GLAreaClass {
	c := (*C.GtkGLAreaClass)(u)
	if c == nil {
		return nil
	}

	g := &GLAreaClass{native: c}

	return g
}

func (recv *GLAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PaperSizeNewFromIpp is a wrapper around the C function gtk_paper_size_new_from_ipp.
func PaperSizeNewFromIpp(ippName string, width float64, height float64) *PaperSize {
	c_ipp_name := C.CString(ippName)
	defer C.free(unsafe.Pointer(c_ipp_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	retC := C.gtk_paper_size_new_from_ipp(c_ipp_name, c_width, c_height)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}
