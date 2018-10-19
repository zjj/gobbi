// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// ActionBarNew is a wrapper around the C function gtk_action_bar_new.
func ActionBarNew() *ActionBar {
	retC := C.gtk_action_bar_new()
	retGo := ActionBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCenterWidget is a wrapper around the C function gtk_action_bar_get_center_widget.
func (recv *ActionBar) GetCenterWidget() *Widget {
	retC := C.gtk_action_bar_get_center_widget((*C.GtkActionBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PackEnd is a wrapper around the C function gtk_action_bar_pack_end.
func (recv *ActionBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_action_bar_pack_end((*C.GtkActionBar)(recv.native), c_child)

	return
}

// PackStart is a wrapper around the C function gtk_action_bar_pack_start.
func (recv *ActionBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_action_bar_pack_start((*C.GtkActionBar)(recv.native), c_child)

	return
}

// SetCenterWidget is a wrapper around the C function gtk_action_bar_set_center_widget.
func (recv *ActionBar) SetCenterWidget(centerWidget *Widget) {
	c_center_widget := (*C.GtkWidget)(centerWidget.ToC())

	C.gtk_action_bar_set_center_widget((*C.GtkActionBar)(recv.native), c_center_widget)

	return
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// GetCenterWidget is a wrapper around the C function gtk_box_get_center_widget.
func (recv *Box) GetCenterWidget() *Widget {
	retC := C.gtk_box_get_center_widget((*C.GtkBox)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetCenterWidget is a wrapper around the C function gtk_box_set_center_widget.
func (recv *Box) SetCenterWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_box_set_center_widget((*C.GtkBox)(recv.native), c_widget)

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetHeaderBar is a wrapper around the C function gtk_dialog_get_header_bar.
func (recv *Dialog) GetHeaderBar() *Widget {
	retC := C.gtk_dialog_get_header_bar((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMaxWidthChars is a wrapper around the C function gtk_entry_get_max_width_chars.
func (recv *Entry) GetMaxWidthChars() int32 {
	retC := C.gtk_entry_get_max_width_chars((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetMaxWidthChars is a wrapper around the C function gtk_entry_set_max_width_chars.
func (recv *Entry) SetMaxWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_set_max_width_chars((*C.GtkEntry)(recv.native), c_n_chars)

	return
}

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// FlowBoxNew is a wrapper around the C function gtk_flow_box_new.
func FlowBoxNew() *FlowBox {
	retC := C.gtk_flow_box_new()
	retGo := FlowBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_flow_box_get_activate_on_single_click.
func (recv *FlowBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_flow_box_get_activate_on_single_click((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetChildAtIndex is a wrapper around the C function gtk_flow_box_get_child_at_index.
func (recv *FlowBox) GetChildAtIndex(idx int32) *FlowBoxChild {
	c_idx := (C.gint)(idx)

	retC := C.gtk_flow_box_get_child_at_index((*C.GtkFlowBox)(recv.native), c_idx)
	retGo := FlowBoxChildNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetColumnSpacing is a wrapper around the C function gtk_flow_box_get_column_spacing.
func (recv *FlowBox) GetColumnSpacing() uint32 {
	retC := C.gtk_flow_box_get_column_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_flow_box_get_homogeneous.
func (recv *FlowBox) GetHomogeneous() bool {
	retC := C.gtk_flow_box_get_homogeneous((*C.GtkFlowBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMaxChildrenPerLine is a wrapper around the C function gtk_flow_box_get_max_children_per_line.
func (recv *FlowBox) GetMaxChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_max_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMinChildrenPerLine is a wrapper around the C function gtk_flow_box_get_min_children_per_line.
func (recv *FlowBox) GetMinChildrenPerLine() uint32 {
	retC := C.gtk_flow_box_get_min_children_per_line((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_flow_box_get_row_spacing.
func (recv *FlowBox) GetRowSpacing() uint32 {
	retC := C.gtk_flow_box_get_row_spacing((*C.GtkFlowBox)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSelectedChildren is a wrapper around the C function gtk_flow_box_get_selected_children.
func (recv *FlowBox) GetSelectedChildren() *glib.List {
	retC := C.gtk_flow_box_get_selected_children((*C.GtkFlowBox)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_flow_box_get_selection_mode.
func (recv *FlowBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_flow_box_get_selection_mode((*C.GtkFlowBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert is a wrapper around the C function gtk_flow_box_insert.
func (recv *FlowBox) Insert(widget *Widget, position int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_position := (C.gint)(position)

	C.gtk_flow_box_insert((*C.GtkFlowBox)(recv.native), c_widget, c_position)

	return
}

// InvalidateFilter is a wrapper around the C function gtk_flow_box_invalidate_filter.
func (recv *FlowBox) InvalidateFilter() {
	C.gtk_flow_box_invalidate_filter((*C.GtkFlowBox)(recv.native))

	return
}

// InvalidateSort is a wrapper around the C function gtk_flow_box_invalidate_sort.
func (recv *FlowBox) InvalidateSort() {
	C.gtk_flow_box_invalidate_sort((*C.GtkFlowBox)(recv.native))

	return
}

// SelectAll is a wrapper around the C function gtk_flow_box_select_all.
func (recv *FlowBox) SelectAll() {
	C.gtk_flow_box_select_all((*C.GtkFlowBox)(recv.native))

	return
}

// SelectChild is a wrapper around the C function gtk_flow_box_select_child.
func (recv *FlowBox) SelectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(child.ToC())

	C.gtk_flow_box_select_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc, GtkFlowBoxForeachFunc

// SetActivateOnSingleClick is a wrapper around the C function gtk_flow_box_set_activate_on_single_click.
func (recv *FlowBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_flow_box_set_activate_on_single_click((*C.GtkFlowBox)(recv.native), c_single)

	return
}

// SetColumnSpacing is a wrapper around the C function gtk_flow_box_set_column_spacing.
func (recv *FlowBox) SetColumnSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_column_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc, GtkFlowBoxFilterFunc

// SetHadjustment is a wrapper around the C function gtk_flow_box_set_hadjustment.
func (recv *FlowBox) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_flow_box_set_hadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_flow_box_set_homogeneous.
func (recv *FlowBox) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_flow_box_set_homogeneous((*C.GtkFlowBox)(recv.native), c_homogeneous)

	return
}

// SetMaxChildrenPerLine is a wrapper around the C function gtk_flow_box_set_max_children_per_line.
func (recv *FlowBox) SetMaxChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_max_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// SetMinChildrenPerLine is a wrapper around the C function gtk_flow_box_set_min_children_per_line.
func (recv *FlowBox) SetMinChildrenPerLine(nChildren uint32) {
	c_n_children := (C.guint)(nChildren)

	C.gtk_flow_box_set_min_children_per_line((*C.GtkFlowBox)(recv.native), c_n_children)

	return
}

// SetRowSpacing is a wrapper around the C function gtk_flow_box_set_row_spacing.
func (recv *FlowBox) SetRowSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_flow_box_set_row_spacing((*C.GtkFlowBox)(recv.native), c_spacing)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_flow_box_set_selection_mode.
func (recv *FlowBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_flow_box_set_selection_mode((*C.GtkFlowBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc, GtkFlowBoxSortFunc

// SetVadjustment is a wrapper around the C function gtk_flow_box_set_vadjustment.
func (recv *FlowBox) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_flow_box_set_vadjustment((*C.GtkFlowBox)(recv.native), c_adjustment)

	return
}

// UnselectAll is a wrapper around the C function gtk_flow_box_unselect_all.
func (recv *FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all((*C.GtkFlowBox)(recv.native))

	return
}

// UnselectChild is a wrapper around the C function gtk_flow_box_unselect_child.
func (recv *FlowBox) UnselectChild(child *FlowBoxChild) {
	c_child := (*C.GtkFlowBoxChild)(child.ToC())

	C.gtk_flow_box_unselect_child((*C.GtkFlowBox)(recv.native), c_child)

	return
}

// FlowBoxChildNew is a wrapper around the C function gtk_flow_box_child_new.
func FlowBoxChildNew() *FlowBoxChild {
	retC := C.gtk_flow_box_child_new()
	retGo := FlowBoxChildNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_flow_box_child_changed.
func (recv *FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed((*C.GtkFlowBoxChild)(recv.native))

	return
}

// GetIndex is a wrapper around the C function gtk_flow_box_child_get_index.
func (recv *FlowBoxChild) GetIndex() int32 {
	retC := C.gtk_flow_box_child_get_index((*C.GtkFlowBoxChild)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsSelected is a wrapper around the C function gtk_flow_box_child_is_selected.
func (recv *FlowBoxChild) IsSelected() bool {
	retC := C.gtk_flow_box_child_is_selected((*C.GtkFlowBoxChild)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDecorationLayout is a wrapper around the C function gtk_header_bar_get_decoration_layout.
func (recv *HeaderBar) GetDecorationLayout() string {
	retC := C.gtk_header_bar_get_decoration_layout((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHasSubtitle is a wrapper around the C function gtk_header_bar_get_has_subtitle.
func (recv *HeaderBar) GetHasSubtitle() bool {
	retC := C.gtk_header_bar_get_has_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDecorationLayout is a wrapper around the C function gtk_header_bar_set_decoration_layout.
func (recv *HeaderBar) SetDecorationLayout(layout string) {
	c_layout := C.CString(layout)
	defer C.free(unsafe.Pointer(c_layout))

	C.gtk_header_bar_set_decoration_layout((*C.GtkHeaderBar)(recv.native), c_layout)

	return
}

// SetHasSubtitle is a wrapper around the C function gtk_header_bar_set_has_subtitle.
func (recv *HeaderBar) SetHasSubtitle(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_has_subtitle((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// IsSymbolic is a wrapper around the C function gtk_icon_info_is_symbolic.
func (recv *IconInfo) IsSymbolic() bool {
	retC := C.gtk_icon_info_is_symbolic((*C.GtkIconInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// GetPopover is a wrapper around the C function gtk_menu_button_get_popover.
func (recv *MenuButton) GetPopover() *Popover {
	retC := C.gtk_menu_button_get_popover((*C.GtkMenuButton)(recv.native))
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUsePopover is a wrapper around the C function gtk_menu_button_get_use_popover.
func (recv *MenuButton) GetUsePopover() bool {
	retC := C.gtk_menu_button_get_use_popover((*C.GtkMenuButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPopover is a wrapper around the C function gtk_menu_button_set_popover.
func (recv *MenuButton) SetPopover(popover *Widget) {
	c_popover := (*C.GtkWidget)(popover.ToC())

	C.gtk_menu_button_set_popover((*C.GtkMenuButton)(recv.native), c_popover)

	return
}

// SetUsePopover is a wrapper around the C function gtk_menu_button_set_use_popover.
func (recv *MenuButton) SetUsePopover(usePopover bool) {
	c_use_popover :=
		boolToGboolean(usePopover)

	C.gtk_menu_button_set_use_popover((*C.GtkMenuButton)(recv.native), c_use_popover)

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported signal 'get-child-position' for Overlay : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal 'drag-action-ask' for PlacesSidebar : return value gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal 'populate-popup' for PlacesSidebar : unsupported parameter selected_item : no type generator for Gio.File,

// GetLocalOnly is a wrapper around the C function gtk_places_sidebar_get_local_only.
func (recv *PlacesSidebar) GetLocalOnly() bool {
	retC := C.gtk_places_sidebar_get_local_only((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLocalOnly is a wrapper around the C function gtk_places_sidebar_set_local_only.
func (recv *PlacesSidebar) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_places_sidebar_set_local_only((*C.GtkPlacesSidebar)(recv.native), c_local_only)

	return
}

// PopoverNew is a wrapper around the C function gtk_popover_new.
func PopoverNew(relativeTo *Widget) *Popover {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	retC := C.gtk_popover_new(c_relative_to)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopoverNewFromModel is a wrapper around the C function gtk_popover_new_from_model.
func PopoverNewFromModel(relativeTo *Widget, model *gio.MenuModel) *Popover {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	c_model := (*C.GMenuModel)(model.ToC())

	retC := C.gtk_popover_new_from_model(c_relative_to, c_model)
	retGo := PopoverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindModel is a wrapper around the C function gtk_popover_bind_model.
func (recv *Popover) BindModel(model *gio.MenuModel, actionNamespace string) {
	c_model := (*C.GMenuModel)(model.ToC())

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	C.gtk_popover_bind_model((*C.GtkPopover)(recv.native), c_model, c_action_namespace)

	return
}

// GetModal is a wrapper around the C function gtk_popover_get_modal.
func (recv *Popover) GetModal() bool {
	retC := C.gtk_popover_get_modal((*C.GtkPopover)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRelativeTo is a wrapper around the C function gtk_popover_get_relative_to.
func (recv *Popover) GetRelativeTo() *Widget {
	retC := C.gtk_popover_get_relative_to((*C.GtkPopover)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetModal is a wrapper around the C function gtk_popover_set_modal.
func (recv *Popover) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_popover_set_modal((*C.GtkPopover)(recv.native), c_modal)

	return
}

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// SetPosition is a wrapper around the C function gtk_popover_set_position.
func (recv *Popover) SetPosition(position PositionType) {
	c_position := (C.GtkPositionType)(position)

	C.gtk_popover_set_position((*C.GtkPopover)(recv.native), c_position)

	return
}

// SetRelativeTo is a wrapper around the C function gtk_popover_set_relative_to.
func (recv *Popover) SetRelativeTo(relativeTo *Widget) {
	c_relative_to := (*C.GtkWidget)(relativeTo.ToC())

	C.gtk_popover_set_relative_to((*C.GtkPopover)(recv.native), c_relative_to)

	return
}

// Unsupported signal 'preview' for PrintOperation : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported signal 'format-value' for Scale : return value utf8 :

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'input' for SpinButton : return value gint :

// GetChildByName is a wrapper around the C function gtk_stack_get_child_by_name.
func (recv *Stack) GetChildByName(name string) *Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_stack_get_child_by_name((*C.GtkStack)(recv.native), c_name)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTransitionRunning is a wrapper around the C function gtk_stack_get_transition_running.
func (recv *Stack) GetTransitionRunning() bool {
	retC := C.gtk_stack_get_transition_running((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// GetMarginEnd is a wrapper around the C function gtk_widget_get_margin_end.
func (recv *Widget) GetMarginEnd() int32 {
	retC := C.gtk_widget_get_margin_end((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginStart is a wrapper around the C function gtk_widget_get_margin_start.
func (recv *Widget) GetMarginStart() int32 {
	retC := C.gtk_widget_get_margin_start((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetMarginEnd is a wrapper around the C function gtk_widget_set_margin_end.
func (recv *Widget) SetMarginEnd(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_end((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginStart is a wrapper around the C function gtk_widget_set_margin_start.
func (recv *Widget) SetMarginStart(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_start((*C.GtkWidget)(recv.native), c_margin)

	return
}

// IsMaximized is a wrapper around the C function gtk_window_is_maximized.
func (recv *Window) IsMaximized() bool {
	retC := C.gtk_window_is_maximized((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
