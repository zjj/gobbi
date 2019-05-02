// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	atk "github.com/pekim/gobbi/lib/atk"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// Equals compares this CssSection with another CssSection, and returns true if they represent the same GObject.
func (recv *CssSection) Equals(other *CssSection) bool {
	return other.ToC() == recv.ToC()
}

// GetEndLine is a wrapper around the C function gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	retC := C.gtk_css_section_get_end_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetEndPosition is a wrapper around the C function gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	retC := C.gtk_css_section_get_end_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetFile is a wrapper around the C function gtk_css_section_get_file.
func (recv *CssSection) GetFile() *gio.File {
	retC := C.gtk_css_section_get_file((*C.GtkCssSection)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParent is a wrapper around the C function gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	retC := C.gtk_css_section_get_parent((*C.GtkCssSection)(recv.native))
	var retGo (*CssSection)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CssSectionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSectionType is a wrapper around the C function gtk_css_section_get_section_type.
func (recv *CssSection) GetSectionType() CssSectionType {
	retC := C.gtk_css_section_get_section_type((*C.GtkCssSection)(recv.native))
	retGo := (CssSectionType)(retC)

	return retGo
}

// GetStartLine is a wrapper around the C function gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	retC := C.gtk_css_section_get_start_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetStartPosition is a wrapper around the C function gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	retC := C.gtk_css_section_get_start_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Ref is a wrapper around the C function gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	retC := C.gtk_css_section_ref((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_css_section_unref.
func (recv *CssSection) Unref() {
	C.gtk_css_section_unref((*C.GtkCssSection)(recv.native))

	return
}

// Assign is a wrapper around the C function gtk_text_iter_assign.
func (recv *TextIter) Assign(other *TextIter) {
	c_other := (*C.GtkTextIter)(C.NULL)
	if other != nil {
		c_other = (*C.GtkTextIter)(other.ToC())
	}

	C.gtk_text_iter_assign((*C.GtkTextIter)(recv.native), c_other)

	return
}

// SetAccessibleRole is a wrapper around the C function gtk_widget_class_set_accessible_role.
func (recv *WidgetClass) SetAccessibleRole(role atk.Role) {
	c_role := (C.AtkRole)(role)

	C.gtk_widget_class_set_accessible_role((*C.GtkWidgetClass)(recv.native), c_role)

	return
}

// SetAccessibleType is a wrapper around the C function gtk_widget_class_set_accessible_type.
func (recv *WidgetClass) SetAccessibleType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_widget_class_set_accessible_type((*C.GtkWidgetClass)(recv.native), c_type)

	return
}

// AppendForWidget is a wrapper around the C function gtk_widget_path_append_for_widget.
func (recv *WidgetPath) AppendForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_widget_path_append_for_widget((*C.GtkWidgetPath)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

// AppendWithSiblings is a wrapper around the C function gtk_widget_path_append_with_siblings.
func (recv *WidgetPath) AppendWithSiblings(siblings *WidgetPath, siblingIndex uint32) int32 {
	c_siblings := (*C.GtkWidgetPath)(C.NULL)
	if siblings != nil {
		c_siblings = (*C.GtkWidgetPath)(siblings.ToC())
	}

	c_sibling_index := (C.guint)(siblingIndex)

	retC := C.gtk_widget_path_append_with_siblings((*C.GtkWidgetPath)(recv.native), c_siblings, c_sibling_index)
	retGo := (int32)(retC)

	return retGo
}

// Ref is a wrapper around the C function gtk_widget_path_ref.
func (recv *WidgetPath) Ref() *WidgetPath {
	retC := C.gtk_widget_path_ref((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToString is a wrapper around the C function gtk_widget_path_to_string.
func (recv *WidgetPath) ToString() string {
	retC := C.gtk_widget_path_to_string((*C.GtkWidgetPath)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	C.gtk_widget_path_unref((*C.GtkWidgetPath)(recv.native))

	return
}
