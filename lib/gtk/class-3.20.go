// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void PlacesSidebar_mountHandler();

	static gulong PlacesSidebar_signal_connect_mount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount", PlacesSidebar_mountHandler, data);
	}

*/
/*

	void PlacesSidebar_showOtherLocationsWithFlagsHandler();

	static gulong PlacesSidebar_signal_connect_show_other_locations_with_flags(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-other-locations-with-flags", PlacesSidebar_showOtherLocationsWithFlagsHandler, data);
	}

*/
/*

	void PlacesSidebar_unmountHandler();

	static gulong PlacesSidebar_signal_connect_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmount", PlacesSidebar_unmountHandler, data);
	}

*/
/*

	void ShortcutsSection_changeCurrentPageHandler();

	static gulong ShortcutsSection_signal_connect_change_current_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-current-page", ShortcutsSection_changeCurrentPageHandler, data);
	}

*/
/*

	void ShortcutsWindow_closeHandler();

	static gulong ShortcutsWindow_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", ShortcutsWindow_closeHandler, data);
	}

*/
/*

	void ShortcutsWindow_searchHandler();

	static gulong ShortcutsWindow_signal_connect_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search", ShortcutsWindow_searchHandler, data);
	}

*/
import "C"

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// GetHelpOverlay is a wrapper around the C function gtk_application_window_get_help_overlay.
func (recv *ApplicationWindow) GetHelpOverlay() *ShortcutsWindow {
	retC := C.gtk_application_window_get_help_overlay((*C.GtkApplicationWindow)(recv.native))
	retGo := ShortcutsWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetHelpOverlay is a wrapper around the C function gtk_application_window_set_help_overlay.
func (recv *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	c_help_overlay := (*C.GtkShortcutsWindow)(helpOverlay.ToC())

	C.gtk_application_window_set_help_overlay((*C.GtkApplicationWindow)(recv.native), c_help_overlay)

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// FileChooserNative is a wrapper around the C record GtkFileChooserNative.
type FileChooserNative struct {
	native *C.GtkFileChooserNative
}

func FileChooserNativeNewFromC(u unsafe.Pointer) *FileChooserNative {
	c := (*C.GtkFileChooserNative)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNative{native: c}

	return g
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NativeDialog upcasts to *NativeDialog
func (recv *FileChooserNative) NativeDialog() *NativeDialog {
	return NativeDialogNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileChooserNative) Object() *gobject.Object {
	return recv.NativeDialog().Object()
}

// CastToWidget down casts any arbitary Object to FileChooserNative.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserNative.
func CastToFileChooserNative(object *gobject.Object) *FileChooserNative {
	return FileChooserNativeNewFromC(object.ToC())
}

// FileChooserNativeNew is a wrapper around the C function gtk_file_chooser_native_new.
func FileChooserNativeNew(title string, parent *Window, action FileChooserAction, acceptLabel string, cancelLabel string) *FileChooserNative {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(parent.ToC())

	c_action := (C.GtkFileChooserAction)(action)

	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	retC := C.gtk_file_chooser_native_new(c_title, c_parent, c_action, c_accept_label, c_cancel_label)
	retGo := FileChooserNativeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_get_accept_label.
func (recv *FileChooserNative) GetAcceptLabel() string {
	retC := C.gtk_file_chooser_native_get_accept_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCancelLabel is a wrapper around the C function gtk_file_chooser_native_get_cancel_label.
func (recv *FileChooserNative) GetCancelLabel() string {
	retC := C.gtk_file_chooser_native_get_cancel_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_set_accept_label.
func (recv *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	C.gtk_file_chooser_native_set_accept_label((*C.GtkFileChooserNative)(recv.native), c_accept_label)

	return
}

// SetCancelLabel is a wrapper around the C function gtk_file_chooser_native_set_cancel_label.
func (recv *FileChooserNative) SetCancelLabel(cancelLabel string) {
	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	C.gtk_file_chooser_native_set_cancel_label((*C.GtkFileChooserNative)(recv.native), c_cancel_label)

	return
}

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// NativeDialog is a wrapper around the C record GtkNativeDialog.
type NativeDialog struct {
	native *C.GtkNativeDialog
	// parent_instance : record
}

func NativeDialogNewFromC(u unsafe.Pointer) *NativeDialog {
	c := (*C.GtkNativeDialog)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialog{native: c}

	return g
}

func (recv *NativeDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *NativeDialog) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to NativeDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a NativeDialog.
func CastToNativeDialog(object *gobject.Object) *NativeDialog {
	return NativeDialogNewFromC(object.ToC())
}

// Destroy is a wrapper around the C function gtk_native_dialog_destroy.
func (recv *NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy((*C.GtkNativeDialog)(recv.native))

	return
}

// GetModal is a wrapper around the C function gtk_native_dialog_get_modal.
func (recv *NativeDialog) GetModal() bool {
	retC := C.gtk_native_dialog_get_modal((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_native_dialog_get_title.
func (recv *NativeDialog) GetTitle() string {
	retC := C.gtk_native_dialog_get_title((*C.GtkNativeDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTransientFor is a wrapper around the C function gtk_native_dialog_get_transient_for.
func (recv *NativeDialog) GetTransientFor() *Window {
	retC := C.gtk_native_dialog_get_transient_for((*C.GtkNativeDialog)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisible is a wrapper around the C function gtk_native_dialog_get_visible.
func (recv *NativeDialog) GetVisible() bool {
	retC := C.gtk_native_dialog_get_visible((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hide is a wrapper around the C function gtk_native_dialog_hide.
func (recv *NativeDialog) Hide() {
	C.gtk_native_dialog_hide((*C.GtkNativeDialog)(recv.native))

	return
}

// Run is a wrapper around the C function gtk_native_dialog_run.
func (recv *NativeDialog) Run() int32 {
	retC := C.gtk_native_dialog_run((*C.GtkNativeDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetModal is a wrapper around the C function gtk_native_dialog_set_modal.
func (recv *NativeDialog) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_native_dialog_set_modal((*C.GtkNativeDialog)(recv.native), c_modal)

	return
}

// SetTitle is a wrapper around the C function gtk_native_dialog_set_title.
func (recv *NativeDialog) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_native_dialog_set_title((*C.GtkNativeDialog)(recv.native), c_title)

	return
}

// SetTransientFor is a wrapper around the C function gtk_native_dialog_set_transient_for.
func (recv *NativeDialog) SetTransientFor(parent *Window) {
	c_parent := (*C.GtkWindow)(parent.ToC())

	C.gtk_native_dialog_set_transient_for((*C.GtkNativeDialog)(recv.native), c_parent)

	return
}

// Show is a wrapper around the C function gtk_native_dialog_show.
func (recv *NativeDialog) Show() {
	C.gtk_native_dialog_show((*C.GtkNativeDialog)(recv.native))

	return
}

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// PadController is a wrapper around the C record GtkPadController.
type PadController struct {
	native *C.GtkPadController
}

func PadControllerNewFromC(u unsafe.Pointer) *PadController {
	c := (*C.GtkPadController)(u)
	if c == nil {
		return nil
	}

	g := &PadController{native: c}

	return g
}

func (recv *PadController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventController upcasts to *EventController
func (recv *PadController) EventController() *EventController {
	return EventControllerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PadController) Object() *gobject.Object {
	return recv.EventController().Object()
}

// CastToWidget down casts any arbitary Object to PadController.
// Exercise care, as this is a potentially dangerous function if the Object is not a PadController.
func CastToPadController(object *gobject.Object) *PadController {
	return PadControllerNewFromC(object.ToC())
}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

var signalPlacesSidebarMountId int
var signalPlacesSidebarMountMap = make(map[int]PlacesSidebarSignalMountCallback)
var signalPlacesSidebarMountLock sync.Mutex

// PlacesSidebarSignalMountCallback is a callback function for a 'mount' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalMountCallback func(mountOperation *gio.MountOperation)

/*
ConnectMount connects the callback to the 'mount' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectMount to remove it.
*/
func (recv *PlacesSidebar) ConnectMount(callback PlacesSidebarSignalMountCallback) int {
	signalPlacesSidebarMountLock.Lock()
	defer signalPlacesSidebarMountLock.Unlock()

	signalPlacesSidebarMountId++
	signalPlacesSidebarMountMap[signalPlacesSidebarMountId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.PlacesSidebar_signal_connect_mount(instance, C.gpointer(uintptr(signalPlacesSidebarMountId)))
	return int(retC)
}

/*
DisconnectMount disconnects a callback from the 'mount' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectMount.
*/
func (recv *PlacesSidebar) DisconnectMount(connectionID int) {
	_, exists := signalPlacesSidebarMountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalPlacesSidebarMountMap, connectionID)
}

//export PlacesSidebar_mountHandler
func PlacesSidebar_mountHandler() {
	fmt.Println("cb")
}

// Unsupported signal : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal : unsupported parameter selected_item : no type generator for Gio.File,

var signalPlacesSidebarShowOtherLocationsWithFlagsId int
var signalPlacesSidebarShowOtherLocationsWithFlagsMap = make(map[int]PlacesSidebarSignalShowOtherLocationsWithFlagsCallback)
var signalPlacesSidebarShowOtherLocationsWithFlagsLock sync.Mutex

// PlacesSidebarSignalShowOtherLocationsWithFlagsCallback is a callback function for a 'show-other-locations-with-flags' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowOtherLocationsWithFlagsCallback func(openFlags PlacesOpenFlags)

/*
ConnectShowOtherLocationsWithFlags connects the callback to the 'show-other-locations-with-flags' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowOtherLocationsWithFlags to remove it.
*/
func (recv *PlacesSidebar) ConnectShowOtherLocationsWithFlags(callback PlacesSidebarSignalShowOtherLocationsWithFlagsCallback) int {
	signalPlacesSidebarShowOtherLocationsWithFlagsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsWithFlagsLock.Unlock()

	signalPlacesSidebarShowOtherLocationsWithFlagsId++
	signalPlacesSidebarShowOtherLocationsWithFlagsMap[signalPlacesSidebarShowOtherLocationsWithFlagsId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.PlacesSidebar_signal_connect_show_other_locations_with_flags(instance, C.gpointer(uintptr(signalPlacesSidebarShowOtherLocationsWithFlagsId)))
	return int(retC)
}

/*
DisconnectShowOtherLocationsWithFlags disconnects a callback from the 'show-other-locations-with-flags' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowOtherLocationsWithFlags.
*/
func (recv *PlacesSidebar) DisconnectShowOtherLocationsWithFlags(connectionID int) {
	_, exists := signalPlacesSidebarShowOtherLocationsWithFlagsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalPlacesSidebarShowOtherLocationsWithFlagsMap, connectionID)
}

//export PlacesSidebar_showOtherLocationsWithFlagsHandler
func PlacesSidebar_showOtherLocationsWithFlagsHandler() {
	fmt.Println("cb")
}

var signalPlacesSidebarUnmountId int
var signalPlacesSidebarUnmountMap = make(map[int]PlacesSidebarSignalUnmountCallback)
var signalPlacesSidebarUnmountLock sync.Mutex

// PlacesSidebarSignalUnmountCallback is a callback function for a 'unmount' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalUnmountCallback func(mountOperation *gio.MountOperation)

/*
ConnectUnmount connects the callback to the 'unmount' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectUnmount to remove it.
*/
func (recv *PlacesSidebar) ConnectUnmount(callback PlacesSidebarSignalUnmountCallback) int {
	signalPlacesSidebarUnmountLock.Lock()
	defer signalPlacesSidebarUnmountLock.Unlock()

	signalPlacesSidebarUnmountId++
	signalPlacesSidebarUnmountMap[signalPlacesSidebarUnmountId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.PlacesSidebar_signal_connect_unmount(instance, C.gpointer(uintptr(signalPlacesSidebarUnmountId)))
	return int(retC)
}

/*
DisconnectUnmount disconnects a callback from the 'unmount' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectUnmount.
*/
func (recv *PlacesSidebar) DisconnectUnmount(connectionID int) {
	_, exists := signalPlacesSidebarUnmountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalPlacesSidebarUnmountMap, connectionID)
}

//export PlacesSidebar_unmountHandler
func PlacesSidebar_unmountHandler() {
	fmt.Println("cb")
}

// GetConstrainTo is a wrapper around the C function gtk_popover_get_constrain_to.
func (recv *Popover) GetConstrainTo() PopoverConstraint {
	retC := C.gtk_popover_get_constrain_to((*C.GtkPopover)(recv.native))
	retGo := (PopoverConstraint)(retC)

	return retGo
}

// SetConstrainTo is a wrapper around the C function gtk_popover_set_constrain_to.
func (recv *Popover) SetConstrainTo(constraint PopoverConstraint) {
	c_constraint := (C.GtkPopoverConstraint)(constraint)

	C.gtk_popover_set_constrain_to((*C.GtkPopover)(recv.native), c_constraint)

	return
}

// Unsupported signal : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// ResetProperty is a wrapper around the C function gtk_settings_reset_property.
func (recv *Settings) ResetProperty(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_settings_reset_property((*C.GtkSettings)(recv.native), c_name)

	return
}

// ShortcutLabel is a wrapper around the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native *C.GtkShortcutLabel
}

func ShortcutLabelNewFromC(u unsafe.Pointer) *ShortcutLabel {
	c := (*C.GtkShortcutLabel)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabel{native: c}

	return g
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutLabel) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutLabel) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutLabel) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutLabel) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutLabel) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutLabel.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutLabel.
func CastToShortcutLabel(object *gobject.Object) *ShortcutLabel {
	return ShortcutLabelNewFromC(object.ToC())
}

// ShortcutsGroup is a wrapper around the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native *C.GtkShortcutsGroup
}

func ShortcutsGroupNewFromC(u unsafe.Pointer) *ShortcutsGroup {
	c := (*C.GtkShortcutsGroup)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroup{native: c}

	return g
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsGroup) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsGroup) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsGroup) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsGroup) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsGroup) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsGroup.
func CastToShortcutsGroup(object *gobject.Object) *ShortcutsGroup {
	return ShortcutsGroupNewFromC(object.ToC())
}

// ShortcutsSection is a wrapper around the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native *C.GtkShortcutsSection
}

func ShortcutsSectionNewFromC(u unsafe.Pointer) *ShortcutsSection {
	c := (*C.GtkShortcutsSection)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSection{native: c}

	return g
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsSection) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsSection) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsSection) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsSection) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsSection) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsSection.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsSection.
func CastToShortcutsSection(object *gobject.Object) *ShortcutsSection {
	return ShortcutsSectionNewFromC(object.ToC())
}

var signalShortcutsSectionChangeCurrentPageId int
var signalShortcutsSectionChangeCurrentPageMap = make(map[int]ShortcutsSectionSignalChangeCurrentPageCallback)
var signalShortcutsSectionChangeCurrentPageLock sync.Mutex

// ShortcutsSectionSignalChangeCurrentPageCallback is a callback function for a 'change-current-page' signal emitted from a ShortcutsSection.
type ShortcutsSectionSignalChangeCurrentPageCallback func(object int32) bool

/*
ConnectChangeCurrentPage connects the callback to the 'change-current-page' signal for the ShortcutsSection.

The returned value represents the connection, and may be passed to DisconnectChangeCurrentPage to remove it.
*/
func (recv *ShortcutsSection) ConnectChangeCurrentPage(callback ShortcutsSectionSignalChangeCurrentPageCallback) int {
	signalShortcutsSectionChangeCurrentPageLock.Lock()
	defer signalShortcutsSectionChangeCurrentPageLock.Unlock()

	signalShortcutsSectionChangeCurrentPageId++
	signalShortcutsSectionChangeCurrentPageMap[signalShortcutsSectionChangeCurrentPageId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ShortcutsSection_signal_connect_change_current_page(instance, C.gpointer(uintptr(signalShortcutsSectionChangeCurrentPageId)))
	return int(retC)
}

/*
DisconnectChangeCurrentPage disconnects a callback from the 'change-current-page' signal for the ShortcutsSection.

The connectionID should be a value returned from a call to ConnectChangeCurrentPage.
*/
func (recv *ShortcutsSection) DisconnectChangeCurrentPage(connectionID int) {
	_, exists := signalShortcutsSectionChangeCurrentPageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalShortcutsSectionChangeCurrentPageMap, connectionID)
}

//export ShortcutsSection_changeCurrentPageHandler
func ShortcutsSection_changeCurrentPageHandler() {
	fmt.Println("cb")
}

// ShortcutsShortcut is a wrapper around the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native *C.GtkShortcutsShortcut
}

func ShortcutsShortcutNewFromC(u unsafe.Pointer) *ShortcutsShortcut {
	c := (*C.GtkShortcutsShortcut)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcut{native: c}

	return g
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsShortcut) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsShortcut) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsShortcut) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsShortcut) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsShortcut) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsShortcut.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsShortcut.
func CastToShortcutsShortcut(object *gobject.Object) *ShortcutsShortcut {
	return ShortcutsShortcutNewFromC(object.ToC())
}

// ShortcutsWindow is a wrapper around the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native *C.GtkShortcutsWindow
	// window : record
}

func ShortcutsWindowNewFromC(u unsafe.Pointer) *ShortcutsWindow {
	c := (*C.GtkShortcutsWindow)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindow{native: c}

	return g
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *ShortcutsWindow) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ShortcutsWindow) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *ShortcutsWindow) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsWindow) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsWindow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsWindow) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsWindow.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsWindow.
func CastToShortcutsWindow(object *gobject.Object) *ShortcutsWindow {
	return ShortcutsWindowNewFromC(object.ToC())
}

var signalShortcutsWindowCloseId int
var signalShortcutsWindowCloseMap = make(map[int]ShortcutsWindowSignalCloseCallback)
var signalShortcutsWindowCloseLock sync.Mutex

// ShortcutsWindowSignalCloseCallback is a callback function for a 'close' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalCloseCallback func()

/*
ConnectClose connects the callback to the 'close' signal for the ShortcutsWindow.

The returned value represents the connection, and may be passed to DisconnectClose to remove it.
*/
func (recv *ShortcutsWindow) ConnectClose(callback ShortcutsWindowSignalCloseCallback) int {
	signalShortcutsWindowCloseLock.Lock()
	defer signalShortcutsWindowCloseLock.Unlock()

	signalShortcutsWindowCloseId++
	signalShortcutsWindowCloseMap[signalShortcutsWindowCloseId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ShortcutsWindow_signal_connect_close(instance, C.gpointer(uintptr(signalShortcutsWindowCloseId)))
	return int(retC)
}

/*
DisconnectClose disconnects a callback from the 'close' signal for the ShortcutsWindow.

The connectionID should be a value returned from a call to ConnectClose.
*/
func (recv *ShortcutsWindow) DisconnectClose(connectionID int) {
	_, exists := signalShortcutsWindowCloseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalShortcutsWindowCloseMap, connectionID)
}

//export ShortcutsWindow_closeHandler
func ShortcutsWindow_closeHandler() {
	fmt.Println("cb")
}

var signalShortcutsWindowSearchId int
var signalShortcutsWindowSearchMap = make(map[int]ShortcutsWindowSignalSearchCallback)
var signalShortcutsWindowSearchLock sync.Mutex

// ShortcutsWindowSignalSearchCallback is a callback function for a 'search' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalSearchCallback func()

/*
ConnectSearch connects the callback to the 'search' signal for the ShortcutsWindow.

The returned value represents the connection, and may be passed to DisconnectSearch to remove it.
*/
func (recv *ShortcutsWindow) ConnectSearch(callback ShortcutsWindowSignalSearchCallback) int {
	signalShortcutsWindowSearchLock.Lock()
	defer signalShortcutsWindowSearchLock.Unlock()

	signalShortcutsWindowSearchId++
	signalShortcutsWindowSearchMap[signalShortcutsWindowSearchId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.ShortcutsWindow_signal_connect_search(instance, C.gpointer(uintptr(signalShortcutsWindowSearchId)))
	return int(retC)
}

/*
DisconnectSearch disconnects a callback from the 'search' signal for the ShortcutsWindow.

The connectionID should be a value returned from a call to ConnectSearch.
*/
func (recv *ShortcutsWindow) DisconnectSearch(connectionID int) {
	_, exists := signalShortcutsWindowSearchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalShortcutsWindowSearchMap, connectionID)
}

//export ShortcutsWindow_searchHandler
func ShortcutsWindow_searchHandler() {
	fmt.Println("cb")
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// ToString is a wrapper around the C function gtk_style_context_to_string.
func (recv *StyleContext) ToString(flags StyleContextPrintFlags) string {
	c_flags := (C.GtkStyleContextPrintFlags)(flags)

	retC := C.gtk_style_context_to_string((*C.GtkStyleContext)(recv.native), c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Changed is a wrapper around the C function gtk_text_tag_changed.
func (recv *TextTag) Changed(sizeChanged bool) {
	c_size_changed :=
		boolToGboolean(sizeChanged)

	C.gtk_text_tag_changed((*C.GtkTextTag)(recv.native), c_size_changed)

	return
}

// ResetCursorBlink is a wrapper around the C function gtk_text_view_reset_cursor_blink.
func (recv *TextView) ResetCursorBlink() {
	C.gtk_text_view_reset_cursor_blink((*C.GtkTextView)(recv.native))

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported signal : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetFocusOnClick is a wrapper around the C function gtk_widget_get_focus_on_click.
func (recv *Widget) GetFocusOnClick() bool {
	retC := C.gtk_widget_get_focus_on_click((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// QueueAllocate is a wrapper around the C function gtk_widget_queue_allocate.
func (recv *Widget) QueueAllocate() {
	C.gtk_widget_queue_allocate((*C.GtkWidget)(recv.native))

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_widget_set_focus_on_click.
func (recv *Widget) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_widget_set_focus_on_click((*C.GtkWidget)(recv.native), c_focus_on_click)

	return
}
