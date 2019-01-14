// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
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

// ClipboardGetForDisplay is a wrapper around the C function gtk_clipboard_get_for_display.
func ClipboardGetForDisplay(display *gdk.Display, selection *gdk.Atom) *Clipboard {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gtk_clipboard_get_for_display(c_display, c_selection)
	retGo := ClipboardNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gtk_clipboard_get_display.
func (recv *Clipboard) GetDisplay() *gdk.Display {
	retC := C.gtk_clipboard_get_display((*C.GtkClipboard)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// gtk_color_selection_set_change_palette_with_screen_hook : unsupported parameter func : no type generator for ColorSelectionChangePaletteWithScreenFunc (GtkColorSelectionChangePaletteWithScreenFunc) for param func
// InvisibleNewForScreen is a wrapper around the C function gtk_invisible_new_for_screen.
func InvisibleNewForScreen(screen *gdk.Screen) *Invisible {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gtk_invisible_new_for_screen(c_screen)
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_invisible_get_screen.
func (recv *Invisible) GetScreen() *gdk.Screen {
	retC := C.gtk_invisible_get_screen((*C.GtkInvisible)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetScreen is a wrapper around the C function gtk_invisible_set_screen.
func (recv *Invisible) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_invisible_set_screen((*C.GtkInvisible)(recv.native), c_screen)

	return
}

// IterIsValid is a wrapper around the C function gtk_list_store_iter_is_valid.
func (recv *ListStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_list_store_iter_is_valid((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_list_store_move_after.
func (recv *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_list_store_move_after((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_list_store_move_before.
func (recv *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_list_store_move_before((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// Reorder is a wrapper around the C function gtk_list_store_reorder.
func (recv *ListStore) Reorder(newOrder []int32) {
	c_new_order_array := make([]C.gint, len(newOrder)+1, len(newOrder)+1)
	for i, item := range newOrder {
		c := (C.gint)(item)
		c_new_order_array[i] = c
	}
	c_new_order_array[len(newOrder)] = 0
	c_new_order_arrayPtr := &c_new_order_array[0]
	c_new_order := (*C.gint)(unsafe.Pointer(c_new_order_arrayPtr))

	C.gtk_list_store_reorder((*C.GtkListStore)(recv.native), c_new_order)

	return
}

// Swap is a wrapper around the C function gtk_list_store_swap.
func (recv *ListStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(C.NULL)
	if a != nil {
		c_a = (*C.GtkTreeIter)(a.ToC())
	}

	c_b := (*C.GtkTreeIter)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreeIter)(b.ToC())
	}

	C.gtk_list_store_swap((*C.GtkListStore)(recv.native), c_a, c_b)

	return
}

// SetScreen is a wrapper around the C function gtk_menu_set_screen.
func (recv *Menu) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_menu_set_screen((*C.GtkMenu)(recv.native), c_screen)

	return
}

// SelectFirst is a wrapper around the C function gtk_menu_shell_select_first.
func (recv *MenuShell) SelectFirst(searchSensitive bool) {
	c_search_sensitive :=
		boolToGboolean(searchSensitive)

	C.gtk_menu_shell_select_first((*C.GtkMenuShell)(recv.native), c_search_sensitive)

	return
}

// GetNPages is a wrapper around the C function gtk_notebook_get_n_pages.
func (recv *Notebook) GetNPages() int32 {
	retC := C.gtk_notebook_get_n_pages((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SettingsGetForScreen is a wrapper around the C function gtk_settings_get_for_screen.
func SettingsGetForScreen(screen *gdk.Screen) *Settings {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gtk_settings_get_for_screen(c_screen)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterIsValid is a wrapper around the C function gtk_tree_model_sort_iter_is_valid.
func (recv *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_sort_iter_is_valid((*C.GtkTreeModelSort)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// CountSelectedRows is a wrapper around the C function gtk_tree_selection_count_selected_rows.
func (recv *TreeSelection) CountSelectedRows() int32 {
	retC := C.gtk_tree_selection_count_selected_rows((*C.GtkTreeSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSelectedRows is a wrapper around the C function gtk_tree_selection_get_selected_rows.
func (recv *TreeSelection) GetSelectedRows() (*glib.List, *TreeModel) {
	var c_model *C.GtkTreeModel

	retC := C.gtk_tree_selection_get_selected_rows((*C.GtkTreeSelection)(recv.native), &c_model)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	return retGo, model
}

// UnselectRange is a wrapper around the C function gtk_tree_selection_unselect_range.
func (recv *TreeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	c_start_path := (*C.GtkTreePath)(C.NULL)
	if startPath != nil {
		c_start_path = (*C.GtkTreePath)(startPath.ToC())
	}

	c_end_path := (*C.GtkTreePath)(C.NULL)
	if endPath != nil {
		c_end_path = (*C.GtkTreePath)(endPath.ToC())
	}

	C.gtk_tree_selection_unselect_range((*C.GtkTreeSelection)(recv.native), c_start_path, c_end_path)

	return
}

// IterIsValid is a wrapper around the C function gtk_tree_store_iter_is_valid.
func (recv *TreeStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_store_iter_is_valid((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_tree_store_move_after.
func (recv *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_tree_store_move_after((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_tree_store_move_before.
func (recv *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_tree_store_move_before((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// Reorder is a wrapper around the C function gtk_tree_store_reorder.
func (recv *TreeStore) Reorder(parent *TreeIter, newOrder []int32) {
	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	c_new_order_array := make([]C.gint, len(newOrder)+1, len(newOrder)+1)
	for i, item := range newOrder {
		c := (C.gint)(item)
		c_new_order_array[i] = c
	}
	c_new_order_array[len(newOrder)] = 0
	c_new_order_arrayPtr := &c_new_order_array[0]
	c_new_order := (*C.gint)(unsafe.Pointer(c_new_order_arrayPtr))

	C.gtk_tree_store_reorder((*C.GtkTreeStore)(recv.native), c_parent, c_new_order)

	return
}

// Swap is a wrapper around the C function gtk_tree_store_swap.
func (recv *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(C.NULL)
	if a != nil {
		c_a = (*C.GtkTreeIter)(a.ToC())
	}

	c_b := (*C.GtkTreeIter)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreeIter)(b.ToC())
	}

	C.gtk_tree_store_swap((*C.GtkTreeStore)(recv.native), c_a, c_b)

	return
}

// ExpandToPath is a wrapper around the C function gtk_tree_view_expand_to_path.
func (recv *TreeView) ExpandToPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_view_expand_to_path((*C.GtkTreeView)(recv.native), c_path)

	return
}

// SetCursorOnCell is a wrapper around the C function gtk_tree_view_set_cursor_on_cell.
func (recv *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_focus_column := (*C.GtkTreeViewColumn)(C.NULL)
	if focusColumn != nil {
		c_focus_column = (*C.GtkTreeViewColumn)(focusColumn.ToC())
	}

	c_focus_cell := (*C.GtkCellRenderer)(C.NULL)
	if focusCell != nil {
		c_focus_cell = (*C.GtkCellRenderer)(focusCell.ToC())
	}

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_tree_view_set_cursor_on_cell((*C.GtkTreeView)(recv.native), c_path, c_focus_column, c_focus_cell, c_start_editing)

	return
}

// FocusCell is a wrapper around the C function gtk_tree_view_column_focus_cell.
func (recv *TreeViewColumn) FocusCell(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_tree_view_column_focus_cell((*C.GtkTreeViewColumn)(recv.native), c_cell)

	return
}

// GetClipboard is a wrapper around the C function gtk_widget_get_clipboard.
func (recv *Widget) GetClipboard(selection *gdk.Atom) *Clipboard {
	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gtk_widget_get_clipboard((*C.GtkWidget)(recv.native), c_selection)
	retGo := ClipboardNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gtk_widget_get_display.
func (recv *Widget) GetDisplay() *gdk.Display {
	retC := C.gtk_widget_get_display((*C.GtkWidget)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRootWindow is a wrapper around the C function gtk_widget_get_root_window.
func (recv *Widget) GetRootWindow() *gdk.Window {
	retC := C.gtk_widget_get_root_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_widget_get_screen.
func (recv *Widget) GetScreen() *gdk.Screen {
	retC := C.gtk_widget_get_screen((*C.GtkWidget)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasScreen is a wrapper around the C function gtk_widget_has_screen.
func (recv *Widget) HasScreen() bool {
	retC := C.gtk_widget_has_screen((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WindowSetAutoStartupNotification is a wrapper around the C function gtk_window_set_auto_startup_notification.
func WindowSetAutoStartupNotification(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_auto_startup_notification(c_setting)

	return
}

// WindowSetDefaultIconFromFile is a wrapper around the C function gtk_window_set_default_icon_from_file.
func WindowSetDefaultIconFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_window_set_default_icon_from_file(c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Fullscreen is a wrapper around the C function gtk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gtk_window_fullscreen((*C.GtkWindow)(recv.native))

	return
}

// GetScreen is a wrapper around the C function gtk_window_get_screen.
func (recv *Window) GetScreen() *gdk.Screen {
	retC := C.gtk_window_get_screen((*C.GtkWindow)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSkipPagerHint is a wrapper around the C function gtk_window_get_skip_pager_hint.
func (recv *Window) GetSkipPagerHint() bool {
	retC := C.gtk_window_get_skip_pager_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSkipTaskbarHint is a wrapper around the C function gtk_window_get_skip_taskbar_hint.
func (recv *Window) GetSkipTaskbarHint() bool {
	retC := C.gtk_window_get_skip_taskbar_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIconFromFile is a wrapper around the C function gtk_window_set_icon_from_file.
func (recv *Window) SetIconFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_window_set_icon_from_file((*C.GtkWindow)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetScreen is a wrapper around the C function gtk_window_set_screen.
func (recv *Window) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_window_set_screen((*C.GtkWindow)(recv.native), c_screen)

	return
}

// SetSkipPagerHint is a wrapper around the C function gtk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_pager_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetSkipTaskbarHint is a wrapper around the C function gtk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_taskbar_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Unfullscreen is a wrapper around the C function gtk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gtk_window_unfullscreen((*C.GtkWindow)(recv.native))

	return
}
