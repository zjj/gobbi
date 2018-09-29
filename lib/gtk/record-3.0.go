// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// FindCellProperty is a wrapper around the C function gtk_cell_area_class_find_cell_property.
func (recv *CellAreaClass) FindCellProperty(propertyName string) *gobject.ParamSpec {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.gtk_cell_area_class_find_cell_property((*C.GtkCellAreaClass)(recv.native), c_property_name)
	retGo := gobject.ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InstallCellProperty is a wrapper around the C function gtk_cell_area_class_install_cell_property.
func (recv *CellAreaClass) InstallCellProperty(propertyId uint32, pspec *gobject.ParamSpec) {
	c_property_id := (C.guint)(propertyId)

	c_pspec := (*C.GParamSpec)(pspec.ToC())

	C.gtk_cell_area_class_install_cell_property((*C.GtkCellAreaClass)(recv.native), c_property_id, c_pspec)

	return
}

// Unsupported : gtk_cell_area_class_list_cell_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : gtk_cell_renderer_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_container_class_install_child_properties : unsupported parameter pspecs : no param type

// Unsupported : gtk_container_class_list_child_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : gtk_css_section_get_file : no return generator

// GradientNewLinear is a wrapper around the C function gtk_gradient_new_linear.
func GradientNewLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	retC := C.gtk_gradient_new_linear(c_x0, c_y0, c_x1, c_y1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GradientNewRadial is a wrapper around the C function gtk_gradient_new_radial.
func GradientNewRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_radius0 := (C.gdouble)(radius0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	c_radius1 := (C.gdouble)(radius1)

	retC := C.gtk_gradient_new_radial(c_x0, c_y0, c_radius0, c_x1, c_y1, c_radius1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddColorStop is a wrapper around the C function gtk_gradient_add_color_stop.
func (recv *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	c_offset := (C.gdouble)(offset)

	c_color := (*C.GtkSymbolicColor)(color.ToC())

	C.gtk_gradient_add_color_stop((*C.GtkGradient)(recv.native), c_offset, c_color)

	return
}

// Ref is a wrapper around the C function gtk_gradient_ref.
func (recv *Gradient) Ref() *Gradient {
	retC := C.gtk_gradient_ref((*C.GtkGradient)(recv.native))
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_gradient_resolve : unsupported parameter resolved_gradient : record with indirection level of 2

// Unref is a wrapper around the C function gtk_gradient_unref.
func (recv *Gradient) Unref() {
	C.gtk_gradient_unref((*C.GtkGradient)(recv.native))

	return
}

// Unsupported : gtk_icon_set_get_sizes : unsupported parameter sizes : no param type

// Unsupported : gtk_icon_set_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_set_render_icon_surface : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_source_get_size : no return generator

// Unsupported : gtk_icon_source_set_size : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_paper_size_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_recent_info_create_app_info : no return generator

// Unsupported : gtk_recent_info_get_added : no return generator

// Unsupported : gtk_recent_info_get_application_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : gtk_recent_info_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : gtk_recent_info_get_gicon : no return generator

// Unsupported : gtk_recent_info_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : gtk_recent_info_get_modified : no return generator

// Unsupported : gtk_recent_info_get_visited : no return generator

// RequisitionNew is a wrapper around the C function gtk_requisition_new.
func RequisitionNew() *Requisition {
	retC := C.gtk_requisition_new()
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_selection_data_get_data : no return type

// Unsupported : gtk_selection_data_get_data_type : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_data_with_length : unsupported parameter length : no type generator for gint, gint*

// Unsupported : gtk_selection_data_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_get_targets : unsupported parameter targets : no param type

// Unsupported : gtk_selection_data_get_uris : no return type

// Unsupported : gtk_selection_data_set : unsupported parameter type : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_data_set_uris : unsupported parameter uris : no param type

// Blacklisted : GtkStackAccessibleClass

// SymbolicColorNewAlpha is a wrapper around the C function gtk_symbolic_color_new_alpha.
func SymbolicColorNewAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(color.ToC())

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_alpha(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewLiteral is a wrapper around the C function gtk_symbolic_color_new_literal.
func SymbolicColorNewLiteral(color *gdk.RGBA) *SymbolicColor {
	c_color := (*C.GdkRGBA)(color.ToC())

	retC := C.gtk_symbolic_color_new_literal(c_color)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewMix is a wrapper around the C function gtk_symbolic_color_new_mix.
func SymbolicColorNewMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	c_color1 := (*C.GtkSymbolicColor)(color1.ToC())

	c_color2 := (*C.GtkSymbolicColor)(color2.ToC())

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_mix(c_color1, c_color2, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewName is a wrapper around the C function gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_symbolic_color_new_name(c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewShade is a wrapper around the C function gtk_symbolic_color_new_shade.
func SymbolicColorNewShade(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(color.ToC())

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_shade(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function gtk_symbolic_color_ref.
func (recv *SymbolicColor) Ref() *SymbolicColor {
	retC := C.gtk_symbolic_color_ref((*C.GtkSymbolicColor)(recv.native))
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Resolve is a wrapper around the C function gtk_symbolic_color_resolve.
func (recv *SymbolicColor) Resolve(props *StyleProperties) (bool, *gdk.RGBA) {
	c_props := (*C.GtkStyleProperties)(props.ToC())

	var c_resolved_color C.GdkRGBA

	retC := C.gtk_symbolic_color_resolve((*C.GtkSymbolicColor)(recv.native), c_props, &c_resolved_color)
	retGo := retC == C.TRUE

	resolvedColor := gdk.RGBANewFromC(unsafe.Pointer(&c_resolved_color))

	return retGo, resolvedColor
}

// Unref is a wrapper around the C function gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(recv.native))

	return
}

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_add : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_add_table : unsupported parameter targets : no param type

// Unsupported : gtk_target_list_find : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_remove : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_text_iter_backward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// Unsupported : gtk_text_iter_forward_find_char : unsupported parameter pred : no type generator for TextCharPredicate, GtkTextCharPredicate

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_path_get_indices : no return generator

// Unsupported : gtk_tree_path_get_indices_with_depth : unsupported parameter depth : no type generator for gint, gint*

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_get_model : no return generator

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_widget_class_install_style_property_parser : unsupported parameter parser : no type generator for RcPropertyParser, GtkRcPropertyParser

// Unsupported : gtk_widget_class_list_style_properties : unsupported parameter n_properties : no type generator for guint, guint*

// Unsupported : gtk_widget_class_set_accessible_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// WidgetPathNew is a wrapper around the C function gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {
	retC := C.gtk_widget_path_new()
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_path_append_type : unsupported parameter type : no type generator for GType, GType

// Copy is a wrapper around the C function gtk_widget_path_copy.
func (recv *WidgetPath) Copy() *WidgetPath {
	retC := C.gtk_widget_path_copy((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_widget_path_free.
func (recv *WidgetPath) Free() {
	C.gtk_widget_path_free((*C.GtkWidgetPath)(recv.native))

	return
}

// Unsupported : gtk_widget_path_get_object_type : no return generator

// Unsupported : gtk_widget_path_has_parent : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_path_is_type : unsupported parameter type : no type generator for GType, GType

// IterAddClass is a wrapper around the C function gtk_widget_path_iter_add_class.
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_add_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterAddRegion is a wrapper around the C function gtk_widget_path_iter_add_region.
func (recv *WidgetPath) IterAddRegion(pos int32, name string, flags RegionFlags) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_widget_path_iter_add_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name, c_flags)

	return
}

// IterClearClasses is a wrapper around the C function gtk_widget_path_iter_clear_classes.
func (recv *WidgetPath) IterClearClasses(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_classes((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// IterClearRegions is a wrapper around the C function gtk_widget_path_iter_clear_regions.
func (recv *WidgetPath) IterClearRegions(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_regions((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// Unsupported : gtk_widget_path_iter_get_object_type : no return generator

// IterHasClass is a wrapper around the C function gtk_widget_path_iter_has_class.
func (recv *WidgetPath) IterHasClass(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasName is a wrapper around the C function gtk_widget_path_iter_has_name.
func (recv *WidgetPath) IterHasName(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasQclass is a wrapper around the C function gtk_widget_path_iter_has_qclass.
func (recv *WidgetPath) IterHasQclass(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qclass((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasQname is a wrapper around the C function gtk_widget_path_iter_has_qname.
func (recv *WidgetPath) IterHasQname(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qname((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// IterListClasses is a wrapper around the C function gtk_widget_path_iter_list_classes.
func (recv *WidgetPath) IterListClasses(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_classes((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterListRegions is a wrapper around the C function gtk_widget_path_iter_list_regions.
func (recv *WidgetPath) IterListRegions(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_regions((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterRemoveClass is a wrapper around the C function gtk_widget_path_iter_remove_class.
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterRemoveRegion is a wrapper around the C function gtk_widget_path_iter_remove_region.
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterSetName is a wrapper around the C function gtk_widget_path_iter_set_name.
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Unsupported : gtk_widget_path_iter_set_object_type : unsupported parameter type : no type generator for GType, GType

// Length is a wrapper around the C function gtk_widget_path_length.
func (recv *WidgetPath) Length() int32 {
	retC := C.gtk_widget_path_length((*C.GtkWidgetPath)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_path_prepend_type : unsupported parameter type : no type generator for GType, GType