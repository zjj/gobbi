// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

/*

	static void _g_menu_item_set_action_and_target(GMenuItem* menu_item, const gchar* action, const gchar* format_string) {
		return g_menu_item_set_action_and_target(menu_item, action, format_string);
    }
*/
/*

	static void _g_menu_item_set_attribute(GMenuItem* menu_item, const gchar* attribute, const gchar* format_string) {
		return g_menu_item_set_attribute(menu_item, attribute, format_string);
    }
*/
/*

	void menumodel_itemsChangedHandler(GObject *, gint, gint, gint, gpointer);

	static gulong MenuModel_signal_connect_items_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "items-changed", G_CALLBACK(menumodel_itemsChangedHandler), data);
	}

*/
/*

	static gboolean _g_menu_model_get_item_attribute(GMenuModel* model, gint item_index, const gchar* attribute, const gchar* format_string) {
		return g_menu_model_get_item_attribute(model, item_index, attribute, format_string);
    }
*/
/*

	void socketclient_eventHandler(GObject *, GSocketClientEvent, GSocketConnectable *, GIOStream *, gpointer);

	static gulong SocketClient_signal_connect_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "event", G_CALLBACK(socketclient_eventHandler), data);
	}

*/
import "C"

// GetEnvironment is a wrapper around the C function g_app_launch_context_get_environment.
func (recv *AppLaunchContext) GetEnvironment() []string {
	retC := C.g_app_launch_context_get_environment((*C.GAppLaunchContext)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// Setenv is a wrapper around the C function g_app_launch_context_setenv.
func (recv *AppLaunchContext) Setenv(variable string, value string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_app_launch_context_setenv((*C.GAppLaunchContext)(recv.native), c_variable, c_value)

	return
}

// Unsetenv is a wrapper around the C function g_app_launch_context_unsetenv.
func (recv *AppLaunchContext) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_app_launch_context_unsetenv((*C.GAppLaunchContext)(recv.native), c_variable)

	return
}

// ApplicationGetDefault is a wrapper around the C function g_application_get_default.
func ApplicationGetDefault() *Application {
	retC := C.g_application_get_default()
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Quit is a wrapper around the C function g_application_quit.
func (recv *Application) Quit() {
	C.g_application_quit((*C.GApplication)(recv.native))

	return
}

// SetDefault is a wrapper around the C function g_application_set_default.
func (recv *Application) SetDefault() {
	C.g_application_set_default((*C.GApplication)(recv.native))

	return
}

// DBusActionGroupGet is a wrapper around the C function g_dbus_action_group_get.
func DBusActionGroupGet(connection *DBusConnection, busName string, objectPath string) *DBusActionGroup {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_bus_name := C.CString(busName)
	defer C.free(unsafe.Pointer(c_bus_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_action_group_get(c_connection, c_bus_name, c_object_path)
	retGo := DBusActionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExportActionGroup is a wrapper around the C function g_dbus_connection_export_action_group.
func (recv *DBusConnection) ExportActionGroup(objectPath string, actionGroup *ActionGroup) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_action_group := (*C.GActionGroup)(actionGroup.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_export_action_group((*C.GDBusConnection)(recv.native), c_object_path, c_action_group, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ExportMenuModel is a wrapper around the C function g_dbus_connection_export_menu_model.
func (recv *DBusConnection) ExportMenuModel(objectPath string, menu *MenuModel) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_menu := (*C.GMenuModel)(C.NULL)
	if menu != nil {
		c_menu = (*C.GMenuModel)(menu.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_export_menu_model((*C.GDBusConnection)(recv.native), c_object_path, c_menu, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnexportActionGroup is a wrapper around the C function g_dbus_connection_unexport_action_group.
func (recv *DBusConnection) UnexportActionGroup(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_action_group((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// UnexportMenuModel is a wrapper around the C function g_dbus_connection_unexport_menu_model.
func (recv *DBusConnection) UnexportMenuModel(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_menu_model((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// GetConnections is a wrapper around the C function g_dbus_interface_skeleton_get_connections.
func (recv *DBusInterfaceSkeleton) GetConnections() *glib.List {
	retC := C.g_dbus_interface_skeleton_get_connections((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasConnection is a wrapper around the C function g_dbus_interface_skeleton_has_connection.
func (recv *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	retC := C.g_dbus_interface_skeleton_has_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)
	retGo := retC == C.TRUE

	return retGo
}

// UnexportFromConnection is a wrapper around the C function g_dbus_interface_skeleton_unexport_from_connection.
func (recv *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	C.g_dbus_interface_skeleton_unexport_from_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)

	return
}

// DBusMenuModelGet is a wrapper around the C function g_dbus_menu_model_get.
func DBusMenuModelGet(connection *DBusConnection, busName string, objectPath string) *DBusMenuModel {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_bus_name := C.CString(busName)
	defer C.free(unsafe.Pointer(c_bus_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_menu_model_get(c_connection, c_bus_name, c_object_path)
	retGo := DBusMenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetKeywords is a wrapper around the C function g_desktop_app_info_get_keywords.
func (recv *DesktopAppInfo) GetKeywords() []string {
	retC := C.g_desktop_app_info_get_keywords((*C.GDesktopAppInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// Equals compares this InetAddressMask with another InetAddressMask, and returns true if they represent the same GObject.
func (recv *InetAddressMask) Equals(other *InetAddressMask) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *InetAddressMask) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to InetAddressMask.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetAddressMask.
func CastToInetAddressMask(object *gobject.Object) *InetAddressMask {
	return InetAddressMaskNewFromC(object.ToC())
}

// Equal is a wrapper around the C function g_inet_address_mask_equal.
func (recv *InetAddressMask) Equal(mask2 *InetAddressMask) bool {
	c_mask2 := (*C.GInetAddressMask)(C.NULL)
	if mask2 != nil {
		c_mask2 = (*C.GInetAddressMask)(mask2.ToC())
	}

	retC := C.g_inet_address_mask_equal((*C.GInetAddressMask)(recv.native), c_mask2)
	retGo := retC == C.TRUE

	return retGo
}

// GetAddress is a wrapper around the C function g_inet_address_mask_get_address.
func (recv *InetAddressMask) GetAddress() *InetAddress {
	retC := C.g_inet_address_mask_get_address((*C.GInetAddressMask)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function g_inet_address_mask_get_family.
func (recv *InetAddressMask) GetFamily() SocketFamily {
	retC := C.g_inet_address_mask_get_family((*C.GInetAddressMask)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLength is a wrapper around the C function g_inet_address_mask_get_length.
func (recv *InetAddressMask) GetLength() uint32 {
	retC := C.g_inet_address_mask_get_length((*C.GInetAddressMask)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Matches is a wrapper around the C function g_inet_address_mask_matches.
func (recv *InetAddressMask) Matches(address *InetAddress) bool {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	retC := C.g_inet_address_mask_matches((*C.GInetAddressMask)(recv.native), c_address)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function g_inet_address_mask_to_string.
func (recv *InetAddressMask) ToString() string {
	retC := C.g_inet_address_mask_to_string((*C.GInetAddressMask)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetFlowinfo is a wrapper around the C function g_inet_socket_address_get_flowinfo.
func (recv *InetSocketAddress) GetFlowinfo() uint32 {
	retC := C.g_inet_socket_address_get_flowinfo((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetScopeId is a wrapper around the C function g_inet_socket_address_get_scope_id.
func (recv *InetSocketAddress) GetScopeId() uint32 {
	retC := C.g_inet_socket_address_get_scope_id((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Equals compares this Menu with another Menu, and returns true if they represent the same GObject.
func (recv *Menu) Equals(other *Menu) bool {
	return other.ToC() == recv.ToC()
}

// MenuModel upcasts to *MenuModel
func (recv *Menu) MenuModel() *MenuModel {
	return MenuModelNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Menu) Object() *gobject.Object {
	return recv.MenuModel().Object()
}

// CastToWidget down casts any arbitrary Object to Menu.
// Exercise care, as this is a potentially dangerous function if the Object is not a Menu.
func CastToMenu(object *gobject.Object) *Menu {
	return MenuNewFromC(object.ToC())
}

// Append is a wrapper around the C function g_menu_append.
func (recv *Menu) Append(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_append((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// AppendItem is a wrapper around the C function g_menu_append_item.
func (recv *Menu) AppendItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_append_item((*C.GMenu)(recv.native), c_item)

	return
}

// AppendSection is a wrapper around the C function g_menu_append_section.
func (recv *Menu) AppendSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_append_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// AppendSubmenu is a wrapper around the C function g_menu_append_submenu.
func (recv *Menu) AppendSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_append_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Freeze is a wrapper around the C function g_menu_freeze.
func (recv *Menu) Freeze() {
	C.g_menu_freeze((*C.GMenu)(recv.native))

	return
}

// Insert is a wrapper around the C function g_menu_insert.
func (recv *Menu) Insert(position int32, label string, detailedAction string) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_insert((*C.GMenu)(recv.native), c_position, c_label, c_detailed_action)

	return
}

// InsertItem is a wrapper around the C function g_menu_insert_item.
func (recv *Menu) InsertItem(position int32, item *MenuItem) {
	c_position := (C.gint)(position)

	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_insert_item((*C.GMenu)(recv.native), c_position, c_item)

	return
}

// InsertSection is a wrapper around the C function g_menu_insert_section.
func (recv *Menu) InsertSection(position int32, label string, section *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_insert_section((*C.GMenu)(recv.native), c_position, c_label, c_section)

	return
}

// InsertSubmenu is a wrapper around the C function g_menu_insert_submenu.
func (recv *Menu) InsertSubmenu(position int32, label string, submenu *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_insert_submenu((*C.GMenu)(recv.native), c_position, c_label, c_submenu)

	return
}

// Prepend is a wrapper around the C function g_menu_prepend.
func (recv *Menu) Prepend(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_prepend((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// PrependItem is a wrapper around the C function g_menu_prepend_item.
func (recv *Menu) PrependItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_prepend_item((*C.GMenu)(recv.native), c_item)

	return
}

// PrependSection is a wrapper around the C function g_menu_prepend_section.
func (recv *Menu) PrependSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_prepend_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// PrependSubmenu is a wrapper around the C function g_menu_prepend_submenu.
func (recv *Menu) PrependSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_prepend_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Remove is a wrapper around the C function g_menu_remove.
func (recv *Menu) Remove(position int32) {
	c_position := (C.gint)(position)

	C.g_menu_remove((*C.GMenu)(recv.native), c_position)

	return
}

// Equals compares this MenuAttributeIter with another MenuAttributeIter, and returns true if they represent the same GObject.
func (recv *MenuAttributeIter) Equals(other *MenuAttributeIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuAttributeIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuAttributeIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuAttributeIter.
func CastToMenuAttributeIter(object *gobject.Object) *MenuAttributeIter {
	return MenuAttributeIterNewFromC(object.ToC())
}

// GetName is a wrapper around the C function g_menu_attribute_iter_get_name.
func (recv *MenuAttributeIter) GetName() string {
	retC := C.g_menu_attribute_iter_get_name((*C.GMenuAttributeIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNext is a wrapper around the C function g_menu_attribute_iter_get_next.
func (recv *MenuAttributeIter) GetNext() (bool, string, *glib.Variant) {
	var c_out_name *C.gchar

	var c_value *C.GVariant

	retC := C.g_menu_attribute_iter_get_next((*C.GMenuAttributeIter)(recv.native), &c_out_name, &c_value)
	retGo := retC == C.TRUE

	outName := C.GoString(c_out_name)

	value := glib.VariantNewFromC(unsafe.Pointer(c_value))

	return retGo, outName, value
}

// GetValue is a wrapper around the C function g_menu_attribute_iter_get_value.
func (recv *MenuAttributeIter) GetValue() *glib.Variant {
	retC := C.g_menu_attribute_iter_get_value((*C.GMenuAttributeIter)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Next is a wrapper around the C function g_menu_attribute_iter_next.
func (recv *MenuAttributeIter) Next() bool {
	retC := C.g_menu_attribute_iter_next((*C.GMenuAttributeIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Equals compares this MenuItem with another MenuItem, and returns true if they represent the same GObject.
func (recv *MenuItem) Equals(other *MenuItem) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuItem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuItem.
func CastToMenuItem(object *gobject.Object) *MenuItem {
	return MenuItemNewFromC(object.ToC())
}

// SetActionAndTarget is a wrapper around the C function g_menu_item_set_action_and_target.
func (recv *MenuItem) SetActionAndTarget(action string, formatString string, args ...interface{}) {
	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_menu_item_set_action_and_target((*C.GMenuItem)(recv.native), c_action, c_format_string)

	return
}

// SetActionAndTargetValue is a wrapper around the C function g_menu_item_set_action_and_target_value.
func (recv *MenuItem) SetActionAndTargetValue(action string, targetValue *glib.Variant) {
	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	c_target_value := (*C.GVariant)(C.NULL)
	if targetValue != nil {
		c_target_value = (*C.GVariant)(targetValue.ToC())
	}

	C.g_menu_item_set_action_and_target_value((*C.GMenuItem)(recv.native), c_action, c_target_value)

	return
}

// SetAttribute is a wrapper around the C function g_menu_item_set_attribute.
func (recv *MenuItem) SetAttribute(attribute string, formatString string, args ...interface{}) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_menu_item_set_attribute((*C.GMenuItem)(recv.native), c_attribute, c_format_string)

	return
}

// SetAttributeValue is a wrapper around the C function g_menu_item_set_attribute_value.
func (recv *MenuItem) SetAttributeValue(attribute string, value *glib.Variant) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_menu_item_set_attribute_value((*C.GMenuItem)(recv.native), c_attribute, c_value)

	return
}

// SetDetailedAction is a wrapper around the C function g_menu_item_set_detailed_action.
func (recv *MenuItem) SetDetailedAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_item_set_detailed_action((*C.GMenuItem)(recv.native), c_detailed_action)

	return
}

// SetLabel is a wrapper around the C function g_menu_item_set_label.
func (recv *MenuItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.g_menu_item_set_label((*C.GMenuItem)(recv.native), c_label)

	return
}

// SetLink is a wrapper around the C function g_menu_item_set_link.
func (recv *MenuItem) SetLink(link string, model *MenuModel) {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	C.g_menu_item_set_link((*C.GMenuItem)(recv.native), c_link, c_model)

	return
}

// SetSection is a wrapper around the C function g_menu_item_set_section.
func (recv *MenuItem) SetSection(section *MenuModel) {
	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_item_set_section((*C.GMenuItem)(recv.native), c_section)

	return
}

// SetSubmenu is a wrapper around the C function g_menu_item_set_submenu.
func (recv *MenuItem) SetSubmenu(submenu *MenuModel) {
	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_item_set_submenu((*C.GMenuItem)(recv.native), c_submenu)

	return
}

// Equals compares this MenuLinkIter with another MenuLinkIter, and returns true if they represent the same GObject.
func (recv *MenuLinkIter) Equals(other *MenuLinkIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuLinkIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuLinkIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuLinkIter.
func CastToMenuLinkIter(object *gobject.Object) *MenuLinkIter {
	return MenuLinkIterNewFromC(object.ToC())
}

// GetName is a wrapper around the C function g_menu_link_iter_get_name.
func (recv *MenuLinkIter) GetName() string {
	retC := C.g_menu_link_iter_get_name((*C.GMenuLinkIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNext is a wrapper around the C function g_menu_link_iter_get_next.
func (recv *MenuLinkIter) GetNext() (bool, string, *MenuModel) {
	var c_out_link *C.gchar

	var c_value *C.GMenuModel

	retC := C.g_menu_link_iter_get_next((*C.GMenuLinkIter)(recv.native), &c_out_link, &c_value)
	retGo := retC == C.TRUE

	outLink := C.GoString(c_out_link)

	value := MenuModelNewFromC(unsafe.Pointer(c_value))

	return retGo, outLink, value
}

// GetValue is a wrapper around the C function g_menu_link_iter_get_value.
func (recv *MenuLinkIter) GetValue() *MenuModel {
	retC := C.g_menu_link_iter_get_value((*C.GMenuLinkIter)(recv.native))
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Next is a wrapper around the C function g_menu_link_iter_next.
func (recv *MenuLinkIter) Next() bool {
	retC := C.g_menu_link_iter_next((*C.GMenuLinkIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Equals compares this MenuModel with another MenuModel, and returns true if they represent the same GObject.
func (recv *MenuModel) Equals(other *MenuModel) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MenuModel) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MenuModel.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuModel.
func CastToMenuModel(object *gobject.Object) *MenuModel {
	return MenuModelNewFromC(object.ToC())
}

type signalMenuModelItemsChangedDetail struct {
	callback  MenuModelSignalItemsChangedCallback
	handlerID C.gulong
}

var signalMenuModelItemsChangedId int
var signalMenuModelItemsChangedMap = make(map[int]signalMenuModelItemsChangedDetail)
var signalMenuModelItemsChangedLock sync.RWMutex

// MenuModelSignalItemsChangedCallback is a callback function for a 'items-changed' signal emitted from a MenuModel.
type MenuModelSignalItemsChangedCallback func(position int32, removed int32, added int32)

/*
ConnectItemsChanged connects the callback to the 'items-changed' signal for the MenuModel.

The returned value represents the connection, and may be passed to DisconnectItemsChanged to remove it.
*/
func (recv *MenuModel) ConnectItemsChanged(callback MenuModelSignalItemsChangedCallback) int {
	signalMenuModelItemsChangedLock.Lock()
	defer signalMenuModelItemsChangedLock.Unlock()

	signalMenuModelItemsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.MenuModel_signal_connect_items_changed(instance, C.gpointer(uintptr(signalMenuModelItemsChangedId)))

	detail := signalMenuModelItemsChangedDetail{callback, handlerID}
	signalMenuModelItemsChangedMap[signalMenuModelItemsChangedId] = detail

	return signalMenuModelItemsChangedId
}

/*
DisconnectItemsChanged disconnects a callback from the 'items-changed' signal for the MenuModel.

The connectionID should be a value returned from a call to ConnectItemsChanged.
*/
func (recv *MenuModel) DisconnectItemsChanged(connectionID int) {
	signalMenuModelItemsChangedLock.Lock()
	defer signalMenuModelItemsChangedLock.Unlock()

	detail, exists := signalMenuModelItemsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuModelItemsChangedMap, connectionID)
}

//export menumodel_itemsChangedHandler
func menumodel_itemsChangedHandler(_ *C.GObject, c_position C.gint, c_removed C.gint, c_added C.gint, data C.gpointer) {
	signalMenuModelItemsChangedLock.RLock()
	defer signalMenuModelItemsChangedLock.RUnlock()

	position := int32(c_position)

	removed := int32(c_removed)

	added := int32(c_added)

	index := int(uintptr(data))
	callback := signalMenuModelItemsChangedMap[index].callback
	callback(position, removed, added)
}

// GetItemAttribute is a wrapper around the C function g_menu_model_get_item_attribute.
func (recv *MenuModel) GetItemAttribute(itemIndex int32, attribute string, formatString string, args ...interface{}) bool {
	c_item_index := (C.gint)(itemIndex)

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_menu_model_get_item_attribute((*C.GMenuModel)(recv.native), c_item_index, c_attribute, c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// GetItemAttributeValue is a wrapper around the C function g_menu_model_get_item_attribute_value.
func (recv *MenuModel) GetItemAttributeValue(itemIndex int32, attribute string, expectedType *glib.VariantType) *glib.Variant {
	c_item_index := (C.gint)(itemIndex)

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_expected_type := (*C.GVariantType)(C.NULL)
	if expectedType != nil {
		c_expected_type = (*C.GVariantType)(expectedType.ToC())
	}

	retC := C.g_menu_model_get_item_attribute_value((*C.GMenuModel)(recv.native), c_item_index, c_attribute, c_expected_type)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetItemLink is a wrapper around the C function g_menu_model_get_item_link.
func (recv *MenuModel) GetItemLink(itemIndex int32, link string) *MenuModel {
	c_item_index := (C.gint)(itemIndex)

	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_model_get_item_link((*C.GMenuModel)(recv.native), c_item_index, c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNItems is a wrapper around the C function g_menu_model_get_n_items.
func (recv *MenuModel) GetNItems() int32 {
	retC := C.g_menu_model_get_n_items((*C.GMenuModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsMutable is a wrapper around the C function g_menu_model_is_mutable.
func (recv *MenuModel) IsMutable() bool {
	retC := C.g_menu_model_is_mutable((*C.GMenuModel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ItemsChanged is a wrapper around the C function g_menu_model_items_changed.
func (recv *MenuModel) ItemsChanged(position int32, removed int32, added int32) {
	c_position := (C.gint)(position)

	c_removed := (C.gint)(removed)

	c_added := (C.gint)(added)

	C.g_menu_model_items_changed((*C.GMenuModel)(recv.native), c_position, c_removed, c_added)

	return
}

// IterateItemAttributes is a wrapper around the C function g_menu_model_iterate_item_attributes.
func (recv *MenuModel) IterateItemAttributes(itemIndex int32) *MenuAttributeIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_attributes((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuAttributeIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterateItemLinks is a wrapper around the C function g_menu_model_iterate_item_links.
func (recv *MenuModel) IterateItemLinks(itemIndex int32) *MenuLinkIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_links((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuLinkIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateAction is a wrapper around the C function g_settings_create_action.
func (recv *Settings) CreateAction(key string) *Action {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_create_action((*C.GSettings)(recv.native), c_key)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetCheckCancellable is a wrapper around the C function g_simple_async_result_set_check_cancellable.
func (recv *SimpleAsyncResult) SetCheckCancellable(checkCancellable *Cancellable) {
	c_check_cancellable := (*C.GCancellable)(C.NULL)
	if checkCancellable != nil {
		c_check_cancellable = (*C.GCancellable)(checkCancellable.ToC())
	}

	C.g_simple_async_result_set_check_cancellable((*C.GSimpleAsyncResult)(recv.native), c_check_cancellable)

	return
}

// ConditionTimedWait is a wrapper around the C function g_socket_condition_timed_wait.
func (recv *Socket) ConditionTimedWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_condition_timed_wait((*C.GSocket)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAvailableBytes is a wrapper around the C function g_socket_get_available_bytes.
func (recv *Socket) GetAvailableBytes() int64 {
	retC := C.g_socket_get_available_bytes((*C.GSocket)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetBroadcast is a wrapper around the C function g_socket_get_broadcast.
func (recv *Socket) GetBroadcast() bool {
	retC := C.g_socket_get_broadcast((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMulticastLoopback is a wrapper around the C function g_socket_get_multicast_loopback.
func (recv *Socket) GetMulticastLoopback() bool {
	retC := C.g_socket_get_multicast_loopback((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMulticastTtl is a wrapper around the C function g_socket_get_multicast_ttl.
func (recv *Socket) GetMulticastTtl() uint32 {
	retC := C.g_socket_get_multicast_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTtl is a wrapper around the C function g_socket_get_ttl.
func (recv *Socket) GetTtl() uint32 {
	retC := C.g_socket_get_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// JoinMulticastGroup is a wrapper around the C function g_socket_join_multicast_group.
func (recv *Socket) JoinMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LeaveMulticastGroup is a wrapper around the C function g_socket_leave_multicast_group.
func (recv *Socket) LeaveMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBroadcast is a wrapper around the C function g_socket_set_broadcast.
func (recv *Socket) SetBroadcast(broadcast bool) {
	c_broadcast :=
		boolToGboolean(broadcast)

	C.g_socket_set_broadcast((*C.GSocket)(recv.native), c_broadcast)

	return
}

// SetMulticastLoopback is a wrapper around the C function g_socket_set_multicast_loopback.
func (recv *Socket) SetMulticastLoopback(loopback bool) {
	c_loopback :=
		boolToGboolean(loopback)

	C.g_socket_set_multicast_loopback((*C.GSocket)(recv.native), c_loopback)

	return
}

// SetMulticastTtl is a wrapper around the C function g_socket_set_multicast_ttl.
func (recv *Socket) SetMulticastTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_multicast_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

// SetTtl is a wrapper around the C function g_socket_set_ttl.
func (recv *Socket) SetTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

type signalSocketClientEventDetail struct {
	callback  SocketClientSignalEventCallback
	handlerID C.gulong
}

var signalSocketClientEventId int
var signalSocketClientEventMap = make(map[int]signalSocketClientEventDetail)
var signalSocketClientEventLock sync.RWMutex

// SocketClientSignalEventCallback is a callback function for a 'event' signal emitted from a SocketClient.
type SocketClientSignalEventCallback func(event SocketClientEvent, connectable *SocketConnectable, connection *IOStream)

/*
ConnectEvent connects the callback to the 'event' signal for the SocketClient.

The returned value represents the connection, and may be passed to DisconnectEvent to remove it.
*/
func (recv *SocketClient) ConnectEvent(callback SocketClientSignalEventCallback) int {
	signalSocketClientEventLock.Lock()
	defer signalSocketClientEventLock.Unlock()

	signalSocketClientEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.SocketClient_signal_connect_event(instance, C.gpointer(uintptr(signalSocketClientEventId)))

	detail := signalSocketClientEventDetail{callback, handlerID}
	signalSocketClientEventMap[signalSocketClientEventId] = detail

	return signalSocketClientEventId
}

/*
DisconnectEvent disconnects a callback from the 'event' signal for the SocketClient.

The connectionID should be a value returned from a call to ConnectEvent.
*/
func (recv *SocketClient) DisconnectEvent(connectionID int) {
	signalSocketClientEventLock.Lock()
	defer signalSocketClientEventLock.Unlock()

	detail, exists := signalSocketClientEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketClientEventMap, connectionID)
}

//export socketclient_eventHandler
func socketclient_eventHandler(_ *C.GObject, c_event C.GSocketClientEvent, c_connectable *C.GSocketConnectable, c_connection *C.GIOStream, data C.gpointer) {
	signalSocketClientEventLock.RLock()
	defer signalSocketClientEventLock.RUnlock()

	event := SocketClientEvent(c_event)

	connectable := SocketConnectableNewFromC(unsafe.Pointer(c_connectable))

	connection := IOStreamNewFromC(unsafe.Pointer(c_connection))

	index := int(uintptr(data))
	callback := signalSocketClientEventMap[index].callback
	callback(event, connectable, connection)
}

// Connect is a wrapper around the C function g_socket_connection_connect.
func (recv *SocketConnection) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_connection_connect((*C.GSocketConnection)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectFinish is a wrapper around the C function g_socket_connection_connect_finish.
func (recv *SocketConnection) ConnectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_connection_connect_finish((*C.GSocketConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IsConnected is a wrapper around the C function g_socket_connection_is_connected.
func (recv *SocketConnection) IsConnected() bool {
	retC := C.g_socket_connection_is_connected((*C.GSocketConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_unix_connection_receive_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReceiveCredentialsFinish is a wrapper around the C function g_unix_connection_receive_credentials_finish.
func (recv *UnixConnection) ReceiveCredentialsFinish(result *AsyncResult) (*Credentials, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_credentials_finish((*C.GUnixConnection)(recv.native), c_result, &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_unix_connection_send_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SendCredentialsFinish is a wrapper around the C function g_unix_connection_send_credentials_finish.
func (recv *UnixConnection) SendCredentialsFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_credentials_finish((*C.GUnixConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
