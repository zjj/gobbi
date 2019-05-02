// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void celleditable_editingDoneHandler(GObject *, gpointer);

	static gulong CellEditable_signal_connect_editing_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-done", G_CALLBACK(celleditable_editingDoneHandler), data);
	}

*/
/*

	void celleditable_removeWidgetHandler(GObject *, gpointer);

	static gulong CellEditable_signal_connect_remove_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-widget", G_CALLBACK(celleditable_removeWidgetHandler), data);
	}

*/
/*

	void editable_changedHandler(GObject *, gpointer);

	static gulong Editable_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(editable_changedHandler), data);
	}

*/
/*

	void editable_deleteTextHandler(GObject *, gint, gint, gpointer);

	static gulong Editable_signal_connect_delete_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "delete-text", G_CALLBACK(editable_deleteTextHandler), data);
	}

*/
/*

	void filechooser_currentFolderChangedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_current_folder_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "current-folder-changed", G_CALLBACK(filechooser_currentFolderChangedHandler), data);
	}

*/
/*

	void filechooser_fileActivatedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_file_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-activated", G_CALLBACK(filechooser_fileActivatedHandler), data);
	}

*/
/*

	void filechooser_selectionChangedHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", G_CALLBACK(filechooser_selectionChangedHandler), data);
	}

*/
/*

	void filechooser_updatePreviewHandler(GObject *, gpointer);

	static gulong FileChooser_signal_connect_update_preview(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update-preview", G_CALLBACK(filechooser_updatePreviewHandler), data);
	}

*/
/*

	void fontchooser_fontActivatedHandler(GObject *, gchar*, gpointer);

	static gulong FontChooser_signal_connect_font_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "font-activated", G_CALLBACK(fontchooser_fontActivatedHandler), data);
	}

*/
/*

	void printoperationpreview_gotPageSizeHandler(GObject *, GtkPrintContext *, GtkPageSetup *, gpointer);

	static gulong PrintOperationPreview_signal_connect_got_page_size(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-page-size", G_CALLBACK(printoperationpreview_gotPageSizeHandler), data);
	}

*/
/*

	void printoperationpreview_readyHandler(GObject *, GtkPrintContext *, gpointer);

	static gulong PrintOperationPreview_signal_connect_ready(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "ready", G_CALLBACK(printoperationpreview_readyHandler), data);
	}

*/
/*

	void treemodel_rowChangedHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-changed", G_CALLBACK(treemodel_rowChangedHandler), data);
	}

*/
/*

	void treemodel_rowDeletedHandler(GObject *, GtkTreePath *, gpointer);

	static gulong TreeModel_signal_connect_row_deleted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-deleted", G_CALLBACK(treemodel_rowDeletedHandler), data);
	}

*/
/*

	void treemodel_rowHasChildToggledHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_has_child_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-has-child-toggled", G_CALLBACK(treemodel_rowHasChildToggledHandler), data);
	}

*/
/*

	void treemodel_rowInsertedHandler(GObject *, GtkTreePath *, GtkTreeIter *, gpointer);

	static gulong TreeModel_signal_connect_row_inserted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-inserted", G_CALLBACK(treemodel_rowInsertedHandler), data);
	}

*/
/*

	void treesortable_sortColumnChangedHandler(GObject *, gpointer);

	static gulong TreeSortable_signal_connect_sort_column_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "sort-column-changed", G_CALLBACK(treesortable_sortColumnChangedHandler), data);
	}

*/
import "C"

// Equals compares this Actionable with another Actionable, and returns true if they represent the same GObject.
func (recv *Actionable) Equals(other *Actionable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Activatable with another Activatable, and returns true if they represent the same GObject.
func (recv *Activatable) Equals(other *Activatable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this AppChooser with another AppChooser, and returns true if they represent the same GObject.
func (recv *AppChooser) Equals(other *AppChooser) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Buildable with another Buildable, and returns true if they represent the same GObject.
func (recv *Buildable) Equals(other *Buildable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this CellAccessibleParent with another CellAccessibleParent, and returns true if they represent the same GObject.
func (recv *CellAccessibleParent) Equals(other *CellAccessibleParent) bool {
	return other.ToC() == recv.ToC()
}

// Activate is a wrapper around the C function gtk_cell_accessible_parent_activate.
func (recv *CellAccessibleParent) Activate(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_activate((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

// Edit is a wrapper around the C function gtk_cell_accessible_parent_edit.
func (recv *CellAccessibleParent) Edit(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_edit((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

// ExpandCollapse is a wrapper around the C function gtk_cell_accessible_parent_expand_collapse.
func (recv *CellAccessibleParent) ExpandCollapse(cell *CellAccessible) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	C.gtk_cell_accessible_parent_expand_collapse((*C.GtkCellAccessibleParent)(recv.native), c_cell)

	return
}

// GetCellArea is a wrapper around the C function gtk_cell_accessible_parent_get_cell_area.
func (recv *CellAccessibleParent) GetCellArea(cell *CellAccessible, cellRect *gdk.Rectangle) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	c_cell_rect := (*C.GdkRectangle)(C.NULL)
	if cellRect != nil {
		c_cell_rect = (*C.GdkRectangle)(cellRect.ToC())
	}

	C.gtk_cell_accessible_parent_get_cell_area((*C.GtkCellAccessibleParent)(recv.native), c_cell, c_cell_rect)

	return
}

// GetCellExtents is a wrapper around the C function gtk_cell_accessible_parent_get_cell_extents.
func (recv *CellAccessibleParent) GetCellExtents(cell *CellAccessible, x int32, y int32, width int32, height int32, coordType atk.CoordType) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_coord_type := (C.AtkCoordType)(coordType)

	C.gtk_cell_accessible_parent_get_cell_extents((*C.GtkCellAccessibleParent)(recv.native), c_cell, &c_x, &c_y, &c_width, &c_height, c_coord_type)

	return
}

// GetChildIndex is a wrapper around the C function gtk_cell_accessible_parent_get_child_index.
func (recv *CellAccessibleParent) GetChildIndex(cell *CellAccessible) int32 {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_get_child_index((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := (int32)(retC)

	return retGo
}

// GetRendererState is a wrapper around the C function gtk_cell_accessible_parent_get_renderer_state.
func (recv *CellAccessibleParent) GetRendererState(cell *CellAccessible) CellRendererState {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_get_renderer_state((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := (CellRendererState)(retC)

	return retGo
}

// GrabFocus is a wrapper around the C function gtk_cell_accessible_parent_grab_focus.
func (recv *CellAccessibleParent) GrabFocus(cell *CellAccessible) bool {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	retC := C.gtk_cell_accessible_parent_grab_focus((*C.GtkCellAccessibleParent)(recv.native), c_cell)
	retGo := retC == C.TRUE

	return retGo
}

// UpdateRelationset is a wrapper around the C function gtk_cell_accessible_parent_update_relationset.
func (recv *CellAccessibleParent) UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet) {
	c_cell := (*C.GtkCellAccessible)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellAccessible)(cell.ToC())
	}

	c_relationset := (*C.AtkRelationSet)(C.NULL)
	if relationset != nil {
		c_relationset = (*C.AtkRelationSet)(relationset.ToC())
	}

	C.gtk_cell_accessible_parent_update_relationset((*C.GtkCellAccessibleParent)(recv.native), c_cell, c_relationset)

	return
}

// Equals compares this CellEditable with another CellEditable, and returns true if they represent the same GObject.
func (recv *CellEditable) Equals(other *CellEditable) bool {
	return other.ToC() == recv.ToC()
}

type signalCellEditableEditingDoneDetail struct {
	callback  CellEditableSignalEditingDoneCallback
	handlerID C.gulong
}

var signalCellEditableEditingDoneId int
var signalCellEditableEditingDoneMap = make(map[int]signalCellEditableEditingDoneDetail)
var signalCellEditableEditingDoneLock sync.RWMutex

// CellEditableSignalEditingDoneCallback is a callback function for a 'editing-done' signal emitted from a CellEditable.
type CellEditableSignalEditingDoneCallback func()

/*
ConnectEditingDone connects the callback to the 'editing-done' signal for the CellEditable.

The returned value represents the connection, and may be passed to DisconnectEditingDone to remove it.
*/
func (recv *CellEditable) ConnectEditingDone(callback CellEditableSignalEditingDoneCallback) int {
	signalCellEditableEditingDoneLock.Lock()
	defer signalCellEditableEditingDoneLock.Unlock()

	signalCellEditableEditingDoneId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellEditable_signal_connect_editing_done(instance, C.gpointer(uintptr(signalCellEditableEditingDoneId)))

	detail := signalCellEditableEditingDoneDetail{callback, handlerID}
	signalCellEditableEditingDoneMap[signalCellEditableEditingDoneId] = detail

	return signalCellEditableEditingDoneId
}

/*
DisconnectEditingDone disconnects a callback from the 'editing-done' signal for the CellEditable.

The connectionID should be a value returned from a call to ConnectEditingDone.
*/
func (recv *CellEditable) DisconnectEditingDone(connectionID int) {
	signalCellEditableEditingDoneLock.Lock()
	defer signalCellEditableEditingDoneLock.Unlock()

	detail, exists := signalCellEditableEditingDoneMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellEditableEditingDoneMap, connectionID)
}

//export celleditable_editingDoneHandler
func celleditable_editingDoneHandler(_ *C.GObject, data C.gpointer) {
	signalCellEditableEditingDoneLock.RLock()
	defer signalCellEditableEditingDoneLock.RUnlock()

	index := int(uintptr(data))
	callback := signalCellEditableEditingDoneMap[index].callback
	callback()
}

type signalCellEditableRemoveWidgetDetail struct {
	callback  CellEditableSignalRemoveWidgetCallback
	handlerID C.gulong
}

var signalCellEditableRemoveWidgetId int
var signalCellEditableRemoveWidgetMap = make(map[int]signalCellEditableRemoveWidgetDetail)
var signalCellEditableRemoveWidgetLock sync.RWMutex

// CellEditableSignalRemoveWidgetCallback is a callback function for a 'remove-widget' signal emitted from a CellEditable.
type CellEditableSignalRemoveWidgetCallback func()

/*
ConnectRemoveWidget connects the callback to the 'remove-widget' signal for the CellEditable.

The returned value represents the connection, and may be passed to DisconnectRemoveWidget to remove it.
*/
func (recv *CellEditable) ConnectRemoveWidget(callback CellEditableSignalRemoveWidgetCallback) int {
	signalCellEditableRemoveWidgetLock.Lock()
	defer signalCellEditableRemoveWidgetLock.Unlock()

	signalCellEditableRemoveWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellEditable_signal_connect_remove_widget(instance, C.gpointer(uintptr(signalCellEditableRemoveWidgetId)))

	detail := signalCellEditableRemoveWidgetDetail{callback, handlerID}
	signalCellEditableRemoveWidgetMap[signalCellEditableRemoveWidgetId] = detail

	return signalCellEditableRemoveWidgetId
}

/*
DisconnectRemoveWidget disconnects a callback from the 'remove-widget' signal for the CellEditable.

The connectionID should be a value returned from a call to ConnectRemoveWidget.
*/
func (recv *CellEditable) DisconnectRemoveWidget(connectionID int) {
	signalCellEditableRemoveWidgetLock.Lock()
	defer signalCellEditableRemoveWidgetLock.Unlock()

	detail, exists := signalCellEditableRemoveWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellEditableRemoveWidgetMap, connectionID)
}

//export celleditable_removeWidgetHandler
func celleditable_removeWidgetHandler(_ *C.GObject, data C.gpointer) {
	signalCellEditableRemoveWidgetLock.RLock()
	defer signalCellEditableRemoveWidgetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalCellEditableRemoveWidgetMap[index].callback
	callback()
}

// EditingDone is a wrapper around the C function gtk_cell_editable_editing_done.
func (recv *CellEditable) EditingDone() {
	C.gtk_cell_editable_editing_done((*C.GtkCellEditable)(recv.native))

	return
}

// RemoveWidget is a wrapper around the C function gtk_cell_editable_remove_widget.
func (recv *CellEditable) RemoveWidget() {
	C.gtk_cell_editable_remove_widget((*C.GtkCellEditable)(recv.native))

	return
}

// Unsupported : gtk_cell_editable_start_editing : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Equals compares this CellLayout with another CellLayout, and returns true if they represent the same GObject.
func (recv *CellLayout) Equals(other *CellLayout) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ColorChooser with another ColorChooser, and returns true if they represent the same GObject.
func (recv *ColorChooser) Equals(other *ColorChooser) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Editable with another Editable, and returns true if they represent the same GObject.
func (recv *Editable) Equals(other *Editable) bool {
	return other.ToC() == recv.ToC()
}

type signalEditableChangedDetail struct {
	callback  EditableSignalChangedCallback
	handlerID C.gulong
}

var signalEditableChangedId int
var signalEditableChangedMap = make(map[int]signalEditableChangedDetail)
var signalEditableChangedLock sync.RWMutex

// EditableSignalChangedCallback is a callback function for a 'changed' signal emitted from a Editable.
type EditableSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Editable.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Editable) ConnectChanged(callback EditableSignalChangedCallback) int {
	signalEditableChangedLock.Lock()
	defer signalEditableChangedLock.Unlock()

	signalEditableChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Editable_signal_connect_changed(instance, C.gpointer(uintptr(signalEditableChangedId)))

	detail := signalEditableChangedDetail{callback, handlerID}
	signalEditableChangedMap[signalEditableChangedId] = detail

	return signalEditableChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Editable.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Editable) DisconnectChanged(connectionID int) {
	signalEditableChangedLock.Lock()
	defer signalEditableChangedLock.Unlock()

	detail, exists := signalEditableChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEditableChangedMap, connectionID)
}

//export editable_changedHandler
func editable_changedHandler(_ *C.GObject, data C.gpointer) {
	signalEditableChangedLock.RLock()
	defer signalEditableChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalEditableChangedMap[index].callback
	callback()
}

type signalEditableDeleteTextDetail struct {
	callback  EditableSignalDeleteTextCallback
	handlerID C.gulong
}

var signalEditableDeleteTextId int
var signalEditableDeleteTextMap = make(map[int]signalEditableDeleteTextDetail)
var signalEditableDeleteTextLock sync.RWMutex

// EditableSignalDeleteTextCallback is a callback function for a 'delete-text' signal emitted from a Editable.
type EditableSignalDeleteTextCallback func(startPos int32, endPos int32)

/*
ConnectDeleteText connects the callback to the 'delete-text' signal for the Editable.

The returned value represents the connection, and may be passed to DisconnectDeleteText to remove it.
*/
func (recv *Editable) ConnectDeleteText(callback EditableSignalDeleteTextCallback) int {
	signalEditableDeleteTextLock.Lock()
	defer signalEditableDeleteTextLock.Unlock()

	signalEditableDeleteTextId++
	instance := C.gpointer(recv.native)
	handlerID := C.Editable_signal_connect_delete_text(instance, C.gpointer(uintptr(signalEditableDeleteTextId)))

	detail := signalEditableDeleteTextDetail{callback, handlerID}
	signalEditableDeleteTextMap[signalEditableDeleteTextId] = detail

	return signalEditableDeleteTextId
}

/*
DisconnectDeleteText disconnects a callback from the 'delete-text' signal for the Editable.

The connectionID should be a value returned from a call to ConnectDeleteText.
*/
func (recv *Editable) DisconnectDeleteText(connectionID int) {
	signalEditableDeleteTextLock.Lock()
	defer signalEditableDeleteTextLock.Unlock()

	detail, exists := signalEditableDeleteTextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEditableDeleteTextMap, connectionID)
}

//export editable_deleteTextHandler
func editable_deleteTextHandler(_ *C.GObject, c_start_pos C.gint, c_end_pos C.gint, data C.gpointer) {
	signalEditableDeleteTextLock.RLock()
	defer signalEditableDeleteTextLock.RUnlock()

	startPos := int32(c_start_pos)

	endPos := int32(c_end_pos)

	index := int(uintptr(data))
	callback := signalEditableDeleteTextMap[index].callback
	callback(startPos, endPos)
}

// Unsupported signal 'insert-text' for Editable : param position : gpointer

// CopyClipboard is a wrapper around the C function gtk_editable_copy_clipboard.
func (recv *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard((*C.GtkEditable)(recv.native))

	return
}

// CutClipboard is a wrapper around the C function gtk_editable_cut_clipboard.
func (recv *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard((*C.GtkEditable)(recv.native))

	return
}

// DeleteSelection is a wrapper around the C function gtk_editable_delete_selection.
func (recv *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection((*C.GtkEditable)(recv.native))

	return
}

// DeleteText is a wrapper around the C function gtk_editable_delete_text.
func (recv *Editable) DeleteText(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.gtk_editable_delete_text((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)

	return
}

// GetChars is a wrapper around the C function gtk_editable_get_chars.
func (recv *Editable) GetChars(startPos int32, endPos int32) string {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	retC := C.gtk_editable_get_chars((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetEditable is a wrapper around the C function gtk_editable_get_editable.
func (recv *Editable) GetEditable() bool {
	retC := C.gtk_editable_get_editable((*C.GtkEditable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPosition is a wrapper around the C function gtk_editable_get_position.
func (recv *Editable) GetPosition() int32 {
	retC := C.gtk_editable_get_position((*C.GtkEditable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSelectionBounds is a wrapper around the C function gtk_editable_get_selection_bounds.
func (recv *Editable) GetSelectionBounds() (bool, int32, int32) {
	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.gtk_editable_get_selection_bounds((*C.GtkEditable)(recv.native), &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// InsertText is a wrapper around the C function gtk_editable_insert_text.
func (recv *Editable) InsertText(newText string, newTextLength int32, position int32) {
	c_new_text := C.CString(newText)
	defer C.free(unsafe.Pointer(c_new_text))

	c_new_text_length := (C.gint)(newTextLength)

	c_position := (C.gint)(position)

	C.gtk_editable_insert_text((*C.GtkEditable)(recv.native), c_new_text, c_new_text_length, &c_position)

	return
}

// PasteClipboard is a wrapper around the C function gtk_editable_paste_clipboard.
func (recv *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard((*C.GtkEditable)(recv.native))

	return
}

// SelectRegion is a wrapper around the C function gtk_editable_select_region.
func (recv *Editable) SelectRegion(startPos int32, endPos int32) {
	c_start_pos := (C.gint)(startPos)

	c_end_pos := (C.gint)(endPos)

	C.gtk_editable_select_region((*C.GtkEditable)(recv.native), c_start_pos, c_end_pos)

	return
}

// SetEditable is a wrapper around the C function gtk_editable_set_editable.
func (recv *Editable) SetEditable(isEditable bool) {
	c_is_editable :=
		boolToGboolean(isEditable)

	C.gtk_editable_set_editable((*C.GtkEditable)(recv.native), c_is_editable)

	return
}

// SetPosition is a wrapper around the C function gtk_editable_set_position.
func (recv *Editable) SetPosition(position int32) {
	c_position := (C.gint)(position)

	C.gtk_editable_set_position((*C.GtkEditable)(recv.native), c_position)

	return
}

// Equals compares this FileChooser with another FileChooser, and returns true if they represent the same GObject.
func (recv *FileChooser) Equals(other *FileChooser) bool {
	return other.ToC() == recv.ToC()
}

type signalFileChooserCurrentFolderChangedDetail struct {
	callback  FileChooserSignalCurrentFolderChangedCallback
	handlerID C.gulong
}

var signalFileChooserCurrentFolderChangedId int
var signalFileChooserCurrentFolderChangedMap = make(map[int]signalFileChooserCurrentFolderChangedDetail)
var signalFileChooserCurrentFolderChangedLock sync.RWMutex

// FileChooserSignalCurrentFolderChangedCallback is a callback function for a 'current-folder-changed' signal emitted from a FileChooser.
type FileChooserSignalCurrentFolderChangedCallback func()

/*
ConnectCurrentFolderChanged connects the callback to the 'current-folder-changed' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectCurrentFolderChanged to remove it.
*/
func (recv *FileChooser) ConnectCurrentFolderChanged(callback FileChooserSignalCurrentFolderChangedCallback) int {
	signalFileChooserCurrentFolderChangedLock.Lock()
	defer signalFileChooserCurrentFolderChangedLock.Unlock()

	signalFileChooserCurrentFolderChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_current_folder_changed(instance, C.gpointer(uintptr(signalFileChooserCurrentFolderChangedId)))

	detail := signalFileChooserCurrentFolderChangedDetail{callback, handlerID}
	signalFileChooserCurrentFolderChangedMap[signalFileChooserCurrentFolderChangedId] = detail

	return signalFileChooserCurrentFolderChangedId
}

/*
DisconnectCurrentFolderChanged disconnects a callback from the 'current-folder-changed' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectCurrentFolderChanged.
*/
func (recv *FileChooser) DisconnectCurrentFolderChanged(connectionID int) {
	signalFileChooserCurrentFolderChangedLock.Lock()
	defer signalFileChooserCurrentFolderChangedLock.Unlock()

	detail, exists := signalFileChooserCurrentFolderChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserCurrentFolderChangedMap, connectionID)
}

//export filechooser_currentFolderChangedHandler
func filechooser_currentFolderChangedHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserCurrentFolderChangedLock.RLock()
	defer signalFileChooserCurrentFolderChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserCurrentFolderChangedMap[index].callback
	callback()
}

type signalFileChooserFileActivatedDetail struct {
	callback  FileChooserSignalFileActivatedCallback
	handlerID C.gulong
}

var signalFileChooserFileActivatedId int
var signalFileChooserFileActivatedMap = make(map[int]signalFileChooserFileActivatedDetail)
var signalFileChooserFileActivatedLock sync.RWMutex

// FileChooserSignalFileActivatedCallback is a callback function for a 'file-activated' signal emitted from a FileChooser.
type FileChooserSignalFileActivatedCallback func()

/*
ConnectFileActivated connects the callback to the 'file-activated' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectFileActivated to remove it.
*/
func (recv *FileChooser) ConnectFileActivated(callback FileChooserSignalFileActivatedCallback) int {
	signalFileChooserFileActivatedLock.Lock()
	defer signalFileChooserFileActivatedLock.Unlock()

	signalFileChooserFileActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_file_activated(instance, C.gpointer(uintptr(signalFileChooserFileActivatedId)))

	detail := signalFileChooserFileActivatedDetail{callback, handlerID}
	signalFileChooserFileActivatedMap[signalFileChooserFileActivatedId] = detail

	return signalFileChooserFileActivatedId
}

/*
DisconnectFileActivated disconnects a callback from the 'file-activated' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectFileActivated.
*/
func (recv *FileChooser) DisconnectFileActivated(connectionID int) {
	signalFileChooserFileActivatedLock.Lock()
	defer signalFileChooserFileActivatedLock.Unlock()

	detail, exists := signalFileChooserFileActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserFileActivatedMap, connectionID)
}

//export filechooser_fileActivatedHandler
func filechooser_fileActivatedHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserFileActivatedLock.RLock()
	defer signalFileChooserFileActivatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserFileActivatedMap[index].callback
	callback()
}

type signalFileChooserSelectionChangedDetail struct {
	callback  FileChooserSignalSelectionChangedCallback
	handlerID C.gulong
}

var signalFileChooserSelectionChangedId int
var signalFileChooserSelectionChangedMap = make(map[int]signalFileChooserSelectionChangedDetail)
var signalFileChooserSelectionChangedLock sync.RWMutex

// FileChooserSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a FileChooser.
type FileChooserSignalSelectionChangedCallback func()

/*
ConnectSelectionChanged connects the callback to the 'selection-changed' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectSelectionChanged to remove it.
*/
func (recv *FileChooser) ConnectSelectionChanged(callback FileChooserSignalSelectionChangedCallback) int {
	signalFileChooserSelectionChangedLock.Lock()
	defer signalFileChooserSelectionChangedLock.Unlock()

	signalFileChooserSelectionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_selection_changed(instance, C.gpointer(uintptr(signalFileChooserSelectionChangedId)))

	detail := signalFileChooserSelectionChangedDetail{callback, handlerID}
	signalFileChooserSelectionChangedMap[signalFileChooserSelectionChangedId] = detail

	return signalFileChooserSelectionChangedId
}

/*
DisconnectSelectionChanged disconnects a callback from the 'selection-changed' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectSelectionChanged.
*/
func (recv *FileChooser) DisconnectSelectionChanged(connectionID int) {
	signalFileChooserSelectionChangedLock.Lock()
	defer signalFileChooserSelectionChangedLock.Unlock()

	detail, exists := signalFileChooserSelectionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserSelectionChangedMap, connectionID)
}

//export filechooser_selectionChangedHandler
func filechooser_selectionChangedHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserSelectionChangedLock.RLock()
	defer signalFileChooserSelectionChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserSelectionChangedMap[index].callback
	callback()
}

type signalFileChooserUpdatePreviewDetail struct {
	callback  FileChooserSignalUpdatePreviewCallback
	handlerID C.gulong
}

var signalFileChooserUpdatePreviewId int
var signalFileChooserUpdatePreviewMap = make(map[int]signalFileChooserUpdatePreviewDetail)
var signalFileChooserUpdatePreviewLock sync.RWMutex

// FileChooserSignalUpdatePreviewCallback is a callback function for a 'update-preview' signal emitted from a FileChooser.
type FileChooserSignalUpdatePreviewCallback func()

/*
ConnectUpdatePreview connects the callback to the 'update-preview' signal for the FileChooser.

The returned value represents the connection, and may be passed to DisconnectUpdatePreview to remove it.
*/
func (recv *FileChooser) ConnectUpdatePreview(callback FileChooserSignalUpdatePreviewCallback) int {
	signalFileChooserUpdatePreviewLock.Lock()
	defer signalFileChooserUpdatePreviewLock.Unlock()

	signalFileChooserUpdatePreviewId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooser_signal_connect_update_preview(instance, C.gpointer(uintptr(signalFileChooserUpdatePreviewId)))

	detail := signalFileChooserUpdatePreviewDetail{callback, handlerID}
	signalFileChooserUpdatePreviewMap[signalFileChooserUpdatePreviewId] = detail

	return signalFileChooserUpdatePreviewId
}

/*
DisconnectUpdatePreview disconnects a callback from the 'update-preview' signal for the FileChooser.

The connectionID should be a value returned from a call to ConnectUpdatePreview.
*/
func (recv *FileChooser) DisconnectUpdatePreview(connectionID int) {
	signalFileChooserUpdatePreviewLock.Lock()
	defer signalFileChooserUpdatePreviewLock.Unlock()

	detail, exists := signalFileChooserUpdatePreviewMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserUpdatePreviewMap, connectionID)
}

//export filechooser_updatePreviewHandler
func filechooser_updatePreviewHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserUpdatePreviewLock.RLock()
	defer signalFileChooserUpdatePreviewLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserUpdatePreviewMap[index].callback
	callback()
}

// GetUsePreviewLabel is a wrapper around the C function gtk_file_chooser_get_use_preview_label.
func (recv *FileChooser) GetUsePreviewLabel() bool {
	retC := C.gtk_file_chooser_get_use_preview_label((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Equals compares this FontChooser with another FontChooser, and returns true if they represent the same GObject.
func (recv *FontChooser) Equals(other *FontChooser) bool {
	return other.ToC() == recv.ToC()
}

type signalFontChooserFontActivatedDetail struct {
	callback  FontChooserSignalFontActivatedCallback
	handlerID C.gulong
}

var signalFontChooserFontActivatedId int
var signalFontChooserFontActivatedMap = make(map[int]signalFontChooserFontActivatedDetail)
var signalFontChooserFontActivatedLock sync.RWMutex

// FontChooserSignalFontActivatedCallback is a callback function for a 'font-activated' signal emitted from a FontChooser.
type FontChooserSignalFontActivatedCallback func(fontname string)

/*
ConnectFontActivated connects the callback to the 'font-activated' signal for the FontChooser.

The returned value represents the connection, and may be passed to DisconnectFontActivated to remove it.
*/
func (recv *FontChooser) ConnectFontActivated(callback FontChooserSignalFontActivatedCallback) int {
	signalFontChooserFontActivatedLock.Lock()
	defer signalFontChooserFontActivatedLock.Unlock()

	signalFontChooserFontActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FontChooser_signal_connect_font_activated(instance, C.gpointer(uintptr(signalFontChooserFontActivatedId)))

	detail := signalFontChooserFontActivatedDetail{callback, handlerID}
	signalFontChooserFontActivatedMap[signalFontChooserFontActivatedId] = detail

	return signalFontChooserFontActivatedId
}

/*
DisconnectFontActivated disconnects a callback from the 'font-activated' signal for the FontChooser.

The connectionID should be a value returned from a call to ConnectFontActivated.
*/
func (recv *FontChooser) DisconnectFontActivated(connectionID int) {
	signalFontChooserFontActivatedLock.Lock()
	defer signalFontChooserFontActivatedLock.Unlock()

	detail, exists := signalFontChooserFontActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFontChooserFontActivatedMap, connectionID)
}

//export fontchooser_fontActivatedHandler
func fontchooser_fontActivatedHandler(_ *C.GObject, c_fontname *C.gchar, data C.gpointer) {
	signalFontChooserFontActivatedLock.RLock()
	defer signalFontChooserFontActivatedLock.RUnlock()

	fontname := C.GoString(c_fontname)

	index := int(uintptr(data))
	callback := signalFontChooserFontActivatedMap[index].callback
	callback(fontname)
}

// Equals compares this Orientable with another Orientable, and returns true if they represent the same GObject.
func (recv *Orientable) Equals(other *Orientable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PrintOperationPreview with another PrintOperationPreview, and returns true if they represent the same GObject.
func (recv *PrintOperationPreview) Equals(other *PrintOperationPreview) bool {
	return other.ToC() == recv.ToC()
}

type signalPrintOperationPreviewGotPageSizeDetail struct {
	callback  PrintOperationPreviewSignalGotPageSizeCallback
	handlerID C.gulong
}

var signalPrintOperationPreviewGotPageSizeId int
var signalPrintOperationPreviewGotPageSizeMap = make(map[int]signalPrintOperationPreviewGotPageSizeDetail)
var signalPrintOperationPreviewGotPageSizeLock sync.RWMutex

// PrintOperationPreviewSignalGotPageSizeCallback is a callback function for a 'got-page-size' signal emitted from a PrintOperationPreview.
type PrintOperationPreviewSignalGotPageSizeCallback func(context *PrintContext, pageSetup *PageSetup)

/*
ConnectGotPageSize connects the callback to the 'got-page-size' signal for the PrintOperationPreview.

The returned value represents the connection, and may be passed to DisconnectGotPageSize to remove it.
*/
func (recv *PrintOperationPreview) ConnectGotPageSize(callback PrintOperationPreviewSignalGotPageSizeCallback) int {
	signalPrintOperationPreviewGotPageSizeLock.Lock()
	defer signalPrintOperationPreviewGotPageSizeLock.Unlock()

	signalPrintOperationPreviewGotPageSizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperationPreview_signal_connect_got_page_size(instance, C.gpointer(uintptr(signalPrintOperationPreviewGotPageSizeId)))

	detail := signalPrintOperationPreviewGotPageSizeDetail{callback, handlerID}
	signalPrintOperationPreviewGotPageSizeMap[signalPrintOperationPreviewGotPageSizeId] = detail

	return signalPrintOperationPreviewGotPageSizeId
}

/*
DisconnectGotPageSize disconnects a callback from the 'got-page-size' signal for the PrintOperationPreview.

The connectionID should be a value returned from a call to ConnectGotPageSize.
*/
func (recv *PrintOperationPreview) DisconnectGotPageSize(connectionID int) {
	signalPrintOperationPreviewGotPageSizeLock.Lock()
	defer signalPrintOperationPreviewGotPageSizeLock.Unlock()

	detail, exists := signalPrintOperationPreviewGotPageSizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPreviewGotPageSizeMap, connectionID)
}

//export printoperationpreview_gotPageSizeHandler
func printoperationpreview_gotPageSizeHandler(_ *C.GObject, c_context *C.GtkPrintContext, c_page_setup *C.GtkPageSetup, data C.gpointer) {
	signalPrintOperationPreviewGotPageSizeLock.RLock()
	defer signalPrintOperationPreviewGotPageSizeLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	pageSetup := PageSetupNewFromC(unsafe.Pointer(c_page_setup))

	index := int(uintptr(data))
	callback := signalPrintOperationPreviewGotPageSizeMap[index].callback
	callback(context, pageSetup)
}

type signalPrintOperationPreviewReadyDetail struct {
	callback  PrintOperationPreviewSignalReadyCallback
	handlerID C.gulong
}

var signalPrintOperationPreviewReadyId int
var signalPrintOperationPreviewReadyMap = make(map[int]signalPrintOperationPreviewReadyDetail)
var signalPrintOperationPreviewReadyLock sync.RWMutex

// PrintOperationPreviewSignalReadyCallback is a callback function for a 'ready' signal emitted from a PrintOperationPreview.
type PrintOperationPreviewSignalReadyCallback func(context *PrintContext)

/*
ConnectReady connects the callback to the 'ready' signal for the PrintOperationPreview.

The returned value represents the connection, and may be passed to DisconnectReady to remove it.
*/
func (recv *PrintOperationPreview) ConnectReady(callback PrintOperationPreviewSignalReadyCallback) int {
	signalPrintOperationPreviewReadyLock.Lock()
	defer signalPrintOperationPreviewReadyLock.Unlock()

	signalPrintOperationPreviewReadyId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperationPreview_signal_connect_ready(instance, C.gpointer(uintptr(signalPrintOperationPreviewReadyId)))

	detail := signalPrintOperationPreviewReadyDetail{callback, handlerID}
	signalPrintOperationPreviewReadyMap[signalPrintOperationPreviewReadyId] = detail

	return signalPrintOperationPreviewReadyId
}

/*
DisconnectReady disconnects a callback from the 'ready' signal for the PrintOperationPreview.

The connectionID should be a value returned from a call to ConnectReady.
*/
func (recv *PrintOperationPreview) DisconnectReady(connectionID int) {
	signalPrintOperationPreviewReadyLock.Lock()
	defer signalPrintOperationPreviewReadyLock.Unlock()

	detail, exists := signalPrintOperationPreviewReadyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationPreviewReadyMap, connectionID)
}

//export printoperationpreview_readyHandler
func printoperationpreview_readyHandler(_ *C.GObject, c_context *C.GtkPrintContext, data C.gpointer) {
	signalPrintOperationPreviewReadyLock.RLock()
	defer signalPrintOperationPreviewReadyLock.RUnlock()

	context := PrintContextNewFromC(unsafe.Pointer(c_context))

	index := int(uintptr(data))
	callback := signalPrintOperationPreviewReadyMap[index].callback
	callback(context)
}

// Equals compares this RecentChooser with another RecentChooser, and returns true if they represent the same GObject.
func (recv *RecentChooser) Equals(other *RecentChooser) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this Scrollable with another Scrollable, and returns true if they represent the same GObject.
func (recv *Scrollable) Equals(other *Scrollable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this StyleProvider with another StyleProvider, and returns true if they represent the same GObject.
func (recv *StyleProvider) Equals(other *StyleProvider) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ToolShell with another ToolShell, and returns true if they represent the same GObject.
func (recv *ToolShell) Equals(other *ToolShell) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TreeDragDest with another TreeDragDest, and returns true if they represent the same GObject.
func (recv *TreeDragDest) Equals(other *TreeDragDest) bool {
	return other.ToC() == recv.ToC()
}

// DragDataReceived is a wrapper around the C function gtk_tree_drag_dest_drag_data_received.
func (recv *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	c_dest := (*C.GtkTreePath)(C.NULL)
	if dest != nil {
		c_dest = (*C.GtkTreePath)(dest.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_dest_drag_data_received((*C.GtkTreeDragDest)(recv.native), c_dest, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

// RowDropPossible is a wrapper around the C function gtk_tree_drag_dest_row_drop_possible.
func (recv *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	c_dest_path := (*C.GtkTreePath)(C.NULL)
	if destPath != nil {
		c_dest_path = (*C.GtkTreePath)(destPath.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_dest_row_drop_possible((*C.GtkTreeDragDest)(recv.native), c_dest_path, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

// Equals compares this TreeDragSource with another TreeDragSource, and returns true if they represent the same GObject.
func (recv *TreeDragSource) Equals(other *TreeDragSource) bool {
	return other.ToC() == recv.ToC()
}

// DragDataDelete is a wrapper around the C function gtk_tree_drag_source_drag_data_delete.
func (recv *TreeDragSource) DragDataDelete(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_drag_source_drag_data_delete((*C.GtkTreeDragSource)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// DragDataGet is a wrapper around the C function gtk_tree_drag_source_drag_data_get.
func (recv *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_selection_data := (*C.GtkSelectionData)(C.NULL)
	if selectionData != nil {
		c_selection_data = (*C.GtkSelectionData)(selectionData.ToC())
	}

	retC := C.gtk_tree_drag_source_drag_data_get((*C.GtkTreeDragSource)(recv.native), c_path, c_selection_data)
	retGo := retC == C.TRUE

	return retGo
}

// RowDraggable is a wrapper around the C function gtk_tree_drag_source_row_draggable.
func (recv *TreeDragSource) RowDraggable(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_drag_source_row_draggable((*C.GtkTreeDragSource)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// Equals compares this TreeModel with another TreeModel, and returns true if they represent the same GObject.
func (recv *TreeModel) Equals(other *TreeModel) bool {
	return other.ToC() == recv.ToC()
}

type signalTreeModelRowChangedDetail struct {
	callback  TreeModelSignalRowChangedCallback
	handlerID C.gulong
}

var signalTreeModelRowChangedId int
var signalTreeModelRowChangedMap = make(map[int]signalTreeModelRowChangedDetail)
var signalTreeModelRowChangedLock sync.RWMutex

// TreeModelSignalRowChangedCallback is a callback function for a 'row-changed' signal emitted from a TreeModel.
type TreeModelSignalRowChangedCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowChanged connects the callback to the 'row-changed' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowChanged to remove it.
*/
func (recv *TreeModel) ConnectRowChanged(callback TreeModelSignalRowChangedCallback) int {
	signalTreeModelRowChangedLock.Lock()
	defer signalTreeModelRowChangedLock.Unlock()

	signalTreeModelRowChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_changed(instance, C.gpointer(uintptr(signalTreeModelRowChangedId)))

	detail := signalTreeModelRowChangedDetail{callback, handlerID}
	signalTreeModelRowChangedMap[signalTreeModelRowChangedId] = detail

	return signalTreeModelRowChangedId
}

/*
DisconnectRowChanged disconnects a callback from the 'row-changed' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowChanged.
*/
func (recv *TreeModel) DisconnectRowChanged(connectionID int) {
	signalTreeModelRowChangedLock.Lock()
	defer signalTreeModelRowChangedLock.Unlock()

	detail, exists := signalTreeModelRowChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowChangedMap, connectionID)
}

//export treemodel_rowChangedHandler
func treemodel_rowChangedHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	signalTreeModelRowChangedLock.RLock()
	defer signalTreeModelRowChangedLock.RUnlock()

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowChangedMap[index].callback
	callback(path, iter)
}

type signalTreeModelRowDeletedDetail struct {
	callback  TreeModelSignalRowDeletedCallback
	handlerID C.gulong
}

var signalTreeModelRowDeletedId int
var signalTreeModelRowDeletedMap = make(map[int]signalTreeModelRowDeletedDetail)
var signalTreeModelRowDeletedLock sync.RWMutex

// TreeModelSignalRowDeletedCallback is a callback function for a 'row-deleted' signal emitted from a TreeModel.
type TreeModelSignalRowDeletedCallback func(path *TreePath)

/*
ConnectRowDeleted connects the callback to the 'row-deleted' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowDeleted to remove it.
*/
func (recv *TreeModel) ConnectRowDeleted(callback TreeModelSignalRowDeletedCallback) int {
	signalTreeModelRowDeletedLock.Lock()
	defer signalTreeModelRowDeletedLock.Unlock()

	signalTreeModelRowDeletedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_deleted(instance, C.gpointer(uintptr(signalTreeModelRowDeletedId)))

	detail := signalTreeModelRowDeletedDetail{callback, handlerID}
	signalTreeModelRowDeletedMap[signalTreeModelRowDeletedId] = detail

	return signalTreeModelRowDeletedId
}

/*
DisconnectRowDeleted disconnects a callback from the 'row-deleted' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowDeleted.
*/
func (recv *TreeModel) DisconnectRowDeleted(connectionID int) {
	signalTreeModelRowDeletedLock.Lock()
	defer signalTreeModelRowDeletedLock.Unlock()

	detail, exists := signalTreeModelRowDeletedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowDeletedMap, connectionID)
}

//export treemodel_rowDeletedHandler
func treemodel_rowDeletedHandler(_ *C.GObject, c_path *C.GtkTreePath, data C.gpointer) {
	signalTreeModelRowDeletedLock.RLock()
	defer signalTreeModelRowDeletedLock.RUnlock()

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	index := int(uintptr(data))
	callback := signalTreeModelRowDeletedMap[index].callback
	callback(path)
}

type signalTreeModelRowHasChildToggledDetail struct {
	callback  TreeModelSignalRowHasChildToggledCallback
	handlerID C.gulong
}

var signalTreeModelRowHasChildToggledId int
var signalTreeModelRowHasChildToggledMap = make(map[int]signalTreeModelRowHasChildToggledDetail)
var signalTreeModelRowHasChildToggledLock sync.RWMutex

// TreeModelSignalRowHasChildToggledCallback is a callback function for a 'row-has-child-toggled' signal emitted from a TreeModel.
type TreeModelSignalRowHasChildToggledCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowHasChildToggled connects the callback to the 'row-has-child-toggled' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowHasChildToggled to remove it.
*/
func (recv *TreeModel) ConnectRowHasChildToggled(callback TreeModelSignalRowHasChildToggledCallback) int {
	signalTreeModelRowHasChildToggledLock.Lock()
	defer signalTreeModelRowHasChildToggledLock.Unlock()

	signalTreeModelRowHasChildToggledId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_has_child_toggled(instance, C.gpointer(uintptr(signalTreeModelRowHasChildToggledId)))

	detail := signalTreeModelRowHasChildToggledDetail{callback, handlerID}
	signalTreeModelRowHasChildToggledMap[signalTreeModelRowHasChildToggledId] = detail

	return signalTreeModelRowHasChildToggledId
}

/*
DisconnectRowHasChildToggled disconnects a callback from the 'row-has-child-toggled' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowHasChildToggled.
*/
func (recv *TreeModel) DisconnectRowHasChildToggled(connectionID int) {
	signalTreeModelRowHasChildToggledLock.Lock()
	defer signalTreeModelRowHasChildToggledLock.Unlock()

	detail, exists := signalTreeModelRowHasChildToggledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowHasChildToggledMap, connectionID)
}

//export treemodel_rowHasChildToggledHandler
func treemodel_rowHasChildToggledHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	signalTreeModelRowHasChildToggledLock.RLock()
	defer signalTreeModelRowHasChildToggledLock.RUnlock()

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowHasChildToggledMap[index].callback
	callback(path, iter)
}

type signalTreeModelRowInsertedDetail struct {
	callback  TreeModelSignalRowInsertedCallback
	handlerID C.gulong
}

var signalTreeModelRowInsertedId int
var signalTreeModelRowInsertedMap = make(map[int]signalTreeModelRowInsertedDetail)
var signalTreeModelRowInsertedLock sync.RWMutex

// TreeModelSignalRowInsertedCallback is a callback function for a 'row-inserted' signal emitted from a TreeModel.
type TreeModelSignalRowInsertedCallback func(path *TreePath, iter *TreeIter)

/*
ConnectRowInserted connects the callback to the 'row-inserted' signal for the TreeModel.

The returned value represents the connection, and may be passed to DisconnectRowInserted to remove it.
*/
func (recv *TreeModel) ConnectRowInserted(callback TreeModelSignalRowInsertedCallback) int {
	signalTreeModelRowInsertedLock.Lock()
	defer signalTreeModelRowInsertedLock.Unlock()

	signalTreeModelRowInsertedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeModel_signal_connect_row_inserted(instance, C.gpointer(uintptr(signalTreeModelRowInsertedId)))

	detail := signalTreeModelRowInsertedDetail{callback, handlerID}
	signalTreeModelRowInsertedMap[signalTreeModelRowInsertedId] = detail

	return signalTreeModelRowInsertedId
}

/*
DisconnectRowInserted disconnects a callback from the 'row-inserted' signal for the TreeModel.

The connectionID should be a value returned from a call to ConnectRowInserted.
*/
func (recv *TreeModel) DisconnectRowInserted(connectionID int) {
	signalTreeModelRowInsertedLock.Lock()
	defer signalTreeModelRowInsertedLock.Unlock()

	detail, exists := signalTreeModelRowInsertedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeModelRowInsertedMap, connectionID)
}

//export treemodel_rowInsertedHandler
func treemodel_rowInsertedHandler(_ *C.GObject, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, data C.gpointer) {
	signalTreeModelRowInsertedLock.RLock()
	defer signalTreeModelRowInsertedLock.RUnlock()

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalTreeModelRowInsertedMap[index].callback
	callback(path, iter)
}

// Unsupported signal 'rows-reordered' for TreeModel : unsupported parameter new_order : no type generator for gpointer, gpointer

// Unsupported : gtk_tree_model_foreach : unsupported parameter func : no type generator for TreeModelForeachFunc (GtkTreeModelForeachFunc) for param func

// Unsupported : gtk_tree_model_get : unsupported parameter ... : varargs

// GetColumnType is a wrapper around the C function gtk_tree_model_get_column_type.
func (recv *TreeModel) GetColumnType(index int32) gobject.Type {
	c_index_ := (C.gint)(index)

	retC := C.gtk_tree_model_get_column_type((*C.GtkTreeModel)(recv.native), c_index_)
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function gtk_tree_model_get_flags.
func (recv *TreeModel) GetFlags() TreeModelFlags {
	retC := C.gtk_tree_model_get_flags((*C.GtkTreeModel)(recv.native))
	retGo := (TreeModelFlags)(retC)

	return retGo
}

// GetIter is a wrapper around the C function gtk_tree_model_get_iter.
func (recv *TreeModel) GetIter(path *TreePath) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_tree_model_get_iter((*C.GtkTreeModel)(recv.native), &c_iter, c_path)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// GetIterFirst is a wrapper around the C function gtk_tree_model_get_iter_first.
func (recv *TreeModel) GetIterFirst() (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	retC := C.gtk_tree_model_get_iter_first((*C.GtkTreeModel)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// GetIterFromString is a wrapper around the C function gtk_tree_model_get_iter_from_string.
func (recv *TreeModel) GetIterFromString(pathString string) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_path_string := C.CString(pathString)
	defer C.free(unsafe.Pointer(c_path_string))

	retC := C.gtk_tree_model_get_iter_from_string((*C.GtkTreeModel)(recv.native), &c_iter, c_path_string)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// GetNColumns is a wrapper around the C function gtk_tree_model_get_n_columns.
func (recv *TreeModel) GetNColumns() int32 {
	retC := C.gtk_tree_model_get_n_columns((*C.GtkTreeModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPath is a wrapper around the C function gtk_tree_model_get_path.
func (recv *TreeModel) GetPath(iter *TreeIter) *TreePath {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_get_path((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// GetValue is a wrapper around the C function gtk_tree_model_get_value.
func (recv *TreeModel) GetValue(iter *TreeIter, column int32) *gobject.Value {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_column := (C.gint)(column)

	var c_value C.GValue

	C.gtk_tree_model_get_value((*C.GtkTreeModel)(recv.native), c_iter, c_column, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// IterChildren is a wrapper around the C function gtk_tree_model_iter_children.
func (recv *TreeModel) IterChildren(parent *TreeIter) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	retC := C.gtk_tree_model_iter_children((*C.GtkTreeModel)(recv.native), &c_iter, c_parent)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// IterHasChild is a wrapper around the C function gtk_tree_model_iter_has_child.
func (recv *TreeModel) IterHasChild(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_has_child((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// IterNChildren is a wrapper around the C function gtk_tree_model_iter_n_children.
func (recv *TreeModel) IterNChildren(iter *TreeIter) int32 {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_n_children((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := (int32)(retC)

	return retGo
}

// IterNext is a wrapper around the C function gtk_tree_model_iter_next.
func (recv *TreeModel) IterNext(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_next((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// IterNthChild is a wrapper around the C function gtk_tree_model_iter_nth_child.
func (recv *TreeModel) IterNthChild(parent *TreeIter, n int32) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	c_n := (C.gint)(n)

	retC := C.gtk_tree_model_iter_nth_child((*C.GtkTreeModel)(recv.native), &c_iter, c_parent, c_n)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// IterParent is a wrapper around the C function gtk_tree_model_iter_parent.
func (recv *TreeModel) IterParent(child *TreeIter) (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	c_child := (*C.GtkTreeIter)(C.NULL)
	if child != nil {
		c_child = (*C.GtkTreeIter)(child.ToC())
	}

	retC := C.gtk_tree_model_iter_parent((*C.GtkTreeModel)(recv.native), &c_iter, c_child)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// RefNode is a wrapper around the C function gtk_tree_model_ref_node.
func (recv *TreeModel) RefNode(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_ref_node((*C.GtkTreeModel)(recv.native), c_iter)

	return
}

// RowChanged is a wrapper around the C function gtk_tree_model_row_changed.
func (recv *TreeModel) RowChanged(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_changed((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// RowDeleted is a wrapper around the C function gtk_tree_model_row_deleted.
func (recv *TreeModel) RowDeleted(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_model_row_deleted((*C.GtkTreeModel)(recv.native), c_path)

	return
}

// RowHasChildToggled is a wrapper around the C function gtk_tree_model_row_has_child_toggled.
func (recv *TreeModel) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_has_child_toggled((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// RowInserted is a wrapper around the C function gtk_tree_model_row_inserted.
func (recv *TreeModel) RowInserted(path *TreePath, iter *TreeIter) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_row_inserted((*C.GtkTreeModel)(recv.native), c_path, c_iter)

	return
}

// RowsReordered is a wrapper around the C function gtk_tree_model_rows_reordered.
func (recv *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder int32) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order := (C.gint)(newOrder)

	C.gtk_tree_model_rows_reordered((*C.GtkTreeModel)(recv.native), c_path, c_iter, &c_new_order)

	return
}

// SortNewWithModel is a wrapper around the C function gtk_tree_model_sort_new_with_model.
func (recv *TreeModel) SortNewWithModel() *TreeModel {
	retC := C.gtk_tree_model_sort_new_with_model((*C.GtkTreeModel)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnrefNode is a wrapper around the C function gtk_tree_model_unref_node.
func (recv *TreeModel) UnrefNode(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_tree_model_unref_node((*C.GtkTreeModel)(recv.native), c_iter)

	return
}

// Equals compares this TreeSortable with another TreeSortable, and returns true if they represent the same GObject.
func (recv *TreeSortable) Equals(other *TreeSortable) bool {
	return other.ToC() == recv.ToC()
}

type signalTreeSortableSortColumnChangedDetail struct {
	callback  TreeSortableSignalSortColumnChangedCallback
	handlerID C.gulong
}

var signalTreeSortableSortColumnChangedId int
var signalTreeSortableSortColumnChangedMap = make(map[int]signalTreeSortableSortColumnChangedDetail)
var signalTreeSortableSortColumnChangedLock sync.RWMutex

// TreeSortableSignalSortColumnChangedCallback is a callback function for a 'sort-column-changed' signal emitted from a TreeSortable.
type TreeSortableSignalSortColumnChangedCallback func()

/*
ConnectSortColumnChanged connects the callback to the 'sort-column-changed' signal for the TreeSortable.

The returned value represents the connection, and may be passed to DisconnectSortColumnChanged to remove it.
*/
func (recv *TreeSortable) ConnectSortColumnChanged(callback TreeSortableSignalSortColumnChangedCallback) int {
	signalTreeSortableSortColumnChangedLock.Lock()
	defer signalTreeSortableSortColumnChangedLock.Unlock()

	signalTreeSortableSortColumnChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TreeSortable_signal_connect_sort_column_changed(instance, C.gpointer(uintptr(signalTreeSortableSortColumnChangedId)))

	detail := signalTreeSortableSortColumnChangedDetail{callback, handlerID}
	signalTreeSortableSortColumnChangedMap[signalTreeSortableSortColumnChangedId] = detail

	return signalTreeSortableSortColumnChangedId
}

/*
DisconnectSortColumnChanged disconnects a callback from the 'sort-column-changed' signal for the TreeSortable.

The connectionID should be a value returned from a call to ConnectSortColumnChanged.
*/
func (recv *TreeSortable) DisconnectSortColumnChanged(connectionID int) {
	signalTreeSortableSortColumnChangedLock.Lock()
	defer signalTreeSortableSortColumnChangedLock.Unlock()

	detail, exists := signalTreeSortableSortColumnChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTreeSortableSortColumnChangedMap, connectionID)
}

//export treesortable_sortColumnChangedHandler
func treesortable_sortColumnChangedHandler(_ *C.GObject, data C.gpointer) {
	signalTreeSortableSortColumnChangedLock.RLock()
	defer signalTreeSortableSortColumnChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalTreeSortableSortColumnChangedMap[index].callback
	callback()
}

// Unsupported : gtk_tree_sortable_get_sort_column_id : unsupported parameter order : GtkSortType* with indirection level of 1

// HasDefaultSortFunc is a wrapper around the C function gtk_tree_sortable_has_default_sort_func.
func (recv *TreeSortable) HasDefaultSortFunc() bool {
	retC := C.gtk_tree_sortable_has_default_sort_func((*C.GtkTreeSortable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_sortable_set_default_sort_func : unsupported parameter sort_func : no type generator for TreeIterCompareFunc (GtkTreeIterCompareFunc) for param sort_func

// SetSortColumnId is a wrapper around the C function gtk_tree_sortable_set_sort_column_id.
func (recv *TreeSortable) SetSortColumnId(sortColumnId int32, order SortType) {
	c_sort_column_id := (C.gint)(sortColumnId)

	c_order := (C.GtkSortType)(order)

	C.gtk_tree_sortable_set_sort_column_id((*C.GtkTreeSortable)(recv.native), c_sort_column_id, c_order)

	return
}

// Unsupported : gtk_tree_sortable_set_sort_func : unsupported parameter sort_func : no type generator for TreeIterCompareFunc (GtkTreeIterCompareFunc) for param sort_func

// SortColumnChanged is a wrapper around the C function gtk_tree_sortable_sort_column_changed.
func (recv *TreeSortable) SortColumnChanged() {
	C.gtk_tree_sortable_sort_column_changed((*C.GtkTreeSortable)(recv.native))

	return
}
