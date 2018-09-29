// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_key : no type generator for guint, guint*

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_key : no type generator for guint, guint*

// BindingEntryAddSignalFromString is a wrapper around the C function gtk_binding_entry_add_signal_from_string.
func BindingEntryAddSignalFromString(bindingSet *BindingSet, signalDesc string) glib.TokenType {
	c_binding_set := (*C.GtkBindingSet)(bindingSet.ToC())

	c_signal_desc := C.CString(signalDesc)
	defer C.free(unsafe.Pointer(c_signal_desc))

	retC := C.gtk_binding_entry_add_signal_from_string(c_binding_set, c_signal_desc)
	retGo := (glib.TokenType)(retC)

	return retGo
}

// CairoShouldDrawWindow is a wrapper around the C function gtk_cairo_should_draw_window.
func CairoShouldDrawWindow(cr *cairo.Context, window *gdk.Window) bool {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_cairo_should_draw_window(c_cr, c_window)
	retGo := retC == C.TRUE

	return retGo
}

// CairoTransformToWindow is a wrapper around the C function gtk_cairo_transform_to_window.
func CairoTransformToWindow(cr *cairo.Context, widget *Widget, window *gdk.Window) {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_window := (*C.GdkWindow)(window.ToC())

	C.gtk_cairo_transform_to_window(c_cr, c_widget, c_window)

	return
}

// DeviceGrabAdd is a wrapper around the C function gtk_device_grab_add.
func DeviceGrabAdd(widget *Widget, device *gdk.Device, blockOthers bool) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_device := (*C.GdkDevice)(device.ToC())

	c_block_others :=
		boolToGboolean(blockOthers)

	C.gtk_device_grab_add(c_widget, c_device, c_block_others)

	return
}

// DeviceGrabRemove is a wrapper around the C function gtk_device_grab_remove.
func DeviceGrabRemove(widget *Widget, device *gdk.Device) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_device := (*C.GdkDevice)(device.ToC())

	C.gtk_device_grab_remove(c_widget, c_device)

	return
}

// Unsupported : gtk_drag_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// GetBinaryAge is a wrapper around the C function gtk_get_binary_age.
func GetBinaryAge() uint32 {
	retC := C.gtk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetInterfaceAge is a wrapper around the C function gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	retC := C.gtk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetMajorVersion is a wrapper around the C function gtk_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.gtk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function gtk_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.gtk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function gtk_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.gtk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_lookup_for_settings : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint, GtkIconSize

// Unsupported : gtk_init : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_init_check : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_init_with_args : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc, GtkKeySnoopFunc

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_parse_args : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc, GtkPageSetupDoneFunc

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_rc_get_default_files : no return type

// Unsupported : gtk_rc_get_style_by_paths : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames : no param type

// RenderActivity is a wrapper around the C function gtk_render_activity.
func RenderActivity(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_activity(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderArrow is a wrapper around the C function gtk_render_arrow.
func RenderArrow(context *StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_angle := (C.gdouble)(angle)

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_size := (C.gdouble)(size)

	C.gtk_render_arrow(c_context, c_cr, c_angle, c_x, c_y, c_size)

	return
}

// Unsupported : gtk_render_background_get_clip : unsupported parameter out_clip : Blacklisted record : GdkRectangle

// RenderCheck is a wrapper around the C function gtk_render_check.
func RenderCheck(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_check(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderExpander is a wrapper around the C function gtk_render_expander.
func RenderExpander(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_expander(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderExtension is a wrapper around the C function gtk_render_extension.
func RenderExtension(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	C.gtk_render_extension(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side)

	return
}

// RenderFocus is a wrapper around the C function gtk_render_focus.
func RenderFocus(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_focus(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderFrame is a wrapper around the C function gtk_render_frame.
func RenderFrame(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_frame(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderFrameGap is a wrapper around the C function gtk_render_frame_gap.
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_xy0_gap := (C.gdouble)(xy0Gap)

	c_xy1_gap := (C.gdouble)(xy1Gap)

	C.gtk_render_frame_gap(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side, c_xy0_gap, c_xy1_gap)

	return
}

// RenderHandle is a wrapper around the C function gtk_render_handle.
func RenderHandle(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_handle(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// RenderLayout is a wrapper around the C function gtk_render_layout.
func RenderLayout(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_layout := (*C.PangoLayout)(layout.ToC())

	C.gtk_render_layout(c_context, c_cr, c_x, c_y, c_layout)

	return
}

// RenderLine is a wrapper around the C function gtk_render_line.
func RenderLine(context *StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	C.gtk_render_line(c_context, c_cr, c_x0, c_y0, c_x1, c_y1)

	return
}

// RenderOption is a wrapper around the C function gtk_render_option.
func RenderOption(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_option(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderSlider is a wrapper around the C function gtk_render_slider.
func RenderSlider(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_render_slider(c_context, c_cr, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// Unsupported : gtk_rgb_to_hsv : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs

// Unsupported : gtk_stock_add : unsupported parameter items : no param type

// Unsupported : gtk_stock_add_static : unsupported parameter items : no param type

// Unsupported : gtk_stock_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// Unsupported : gtk_target_table_free : unsupported parameter targets : no param type

// Unsupported : gtk_target_table_new_from_list : unsupported parameter n_targets : no type generator for gint, gint*

// Unsupported : gtk_targets_include_image : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_uri : unsupported parameter targets : no param type

// Unsupported : gtk_test_create_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_find_sibling : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_find_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_init : unsupported parameter argcp : no type generator for gint, int*

// Unsupported : gtk_test_list_all_types : unsupported parameter n_types : no type generator for guint, guint*

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_row_reference_reordered : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*