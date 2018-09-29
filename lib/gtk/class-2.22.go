// This is a generated file - DO NOT EDIT
// +build gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetWidget is a wrapper around the C function gtk_accessible_get_widget.
func (recv *Accessible) GetWidget() *Widget {
	retC := C.gtk_accessible_get_widget((*C.GtkAccessible)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetWidget is a wrapper around the C function gtk_accessible_set_widget.
func (recv *Accessible) SetWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_accessible_set_widget((*C.GtkAccessible)(recv.native), c_widget)

	return
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Commit is a wrapper around the C function gtk_assistant_commit.
func (recv *Assistant) Commit() {
	C.gtk_assistant_commit((*C.GtkAssistant)(recv.native))

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// GetEventWindow is a wrapper around the C function gtk_button_get_event_window.
func (recv *Button) GetEventWindow() *gdk.Window {
	retC := C.gtk_button_get_event_window((*C.GtkButton)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// ImContextFilterKeypress is a wrapper around the C function gtk_entry_im_context_filter_keypress.
func (recv *Entry) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_entry_im_context_filter_keypress((*C.GtkEntry)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// ResetImContext is a wrapper around the C function gtk_entry_reset_im_context.
func (recv *Entry) ResetImContext() {
	C.gtk_entry_reset_im_context((*C.GtkEntry)(recv.native))

	return
}

// Unsupported : EntryIconAccessible : no CType

// GetLabelFill is a wrapper around the C function gtk_expander_get_label_fill.
func (recv *Expander) GetLabelFill() bool {
	retC := C.gtk_expander_get_label_fill((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLabelFill is a wrapper around the C function gtk_expander_set_label_fill.
func (recv *Expander) SetLabelFill(labelFill bool) {
	c_label_fill :=
		boolToGboolean(labelFill)

	C.gtk_expander_set_label_fill((*C.GtkExpander)(recv.native), c_label_fill)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GetFontSelection is a wrapper around the C function gtk_font_selection_dialog_get_font_selection.
func (recv *FontSelectionDialog) GetFontSelection() *Widget {
	retC := C.gtk_font_selection_dialog_get_font_selection((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetItemColumn is a wrapper around the C function gtk_icon_view_get_item_column.
func (recv *IconView) GetItemColumn(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_icon_view_get_item_column((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// GetItemRow is a wrapper around the C function gtk_icon_view_get_item_row.
func (recv *IconView) GetItemRow(path *TreePath) int32 {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_icon_view_get_item_row((*C.GtkIconView)(recv.native), c_path)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// GetMessageArea is a wrapper around the C function gtk_message_dialog_get_message_area.
func (recv *MessageDialog) GetMessageArea() *Widget {
	retC := C.gtk_message_dialog_get_message_area((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTabHborder is a wrapper around the C function gtk_notebook_get_tab_hborder.
func (recv *Notebook) GetTabHborder() uint16 {
	retC := C.gtk_notebook_get_tab_hborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetTabVborder is a wrapper around the C function gtk_notebook_get_tab_vborder.
func (recv *Notebook) GetTabVborder() uint16 {
	retC := C.gtk_notebook_get_tab_vborder((*C.GtkNotebook)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// RemoveAll is a wrapper around the C function gtk_statusbar_remove_all.
func (recv *Statusbar) RemoveAll(contextId uint32) {
	c_context_id := (C.guint)(contextId)

	C.gtk_statusbar_remove_all((*C.GtkStatusbar)(recv.native), c_context_id)

	return
}

// Unsupported : gtk_table_get_size : unsupported parameter rows : no type generator for guint, guint*

// GetHadjustment is a wrapper around the C function gtk_text_view_get_hadjustment.
func (recv *TextView) GetHadjustment() *Adjustment {
	retC := C.gtk_text_view_get_hadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_text_view_get_vadjustment.
func (recv *TextView) GetVadjustment() *Adjustment {
	retC := C.gtk_text_view_get_vadjustment((*C.GtkTextView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImContextFilterKeypress is a wrapper around the C function gtk_text_view_im_context_filter_keypress.
func (recv *TextView) ImContextFilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_text_view_im_context_filter_keypress((*C.GtkTextView)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// ResetImContext is a wrapper around the C function gtk_text_view_reset_im_context.
func (recv *TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context((*C.GtkTextView)(recv.native))

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetViewWindow is a wrapper around the C function gtk_viewport_get_view_window.
func (recv *Viewport) GetViewWindow() *gdk.Window {
	retC := C.gtk_viewport_get_view_window((*C.GtkViewport)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// GetCurrentGrab is a wrapper around the C function gtk_window_group_get_current_grab.
func (recv *WindowGroup) GetCurrentGrab() *Widget {
	retC := C.gtk_window_group_get_current_grab((*C.GtkWindowGroup)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}
