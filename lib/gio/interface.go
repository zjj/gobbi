// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
/*

	void drive_changedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(drive_changedHandler), data);
	}

*/
/*

	void drive_disconnectedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_disconnected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnected", G_CALLBACK(drive_disconnectedHandler), data);
	}

*/
/*

	void drive_ejectButtonHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "eject-button", G_CALLBACK(drive_ejectButtonHandler), data);
	}

*/
/*

	void mount_changedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(mount_changedHandler), data);
	}

*/
/*

	void mount_unmountedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_unmounted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmounted", G_CALLBACK(mount_unmountedHandler), data);
	}

*/
/*

	void volume_changedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(volume_changedHandler), data);
	}

*/
/*

	void volume_removedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "removed", G_CALLBACK(volume_removedHandler), data);
	}

*/
import "C"

// #GAction represents a single named action.
//
// The main interface to an action is that it can be activated with
// g_action_activate().  This results in the 'activate' signal being
// emitted.  An activation has a #GVariant parameter (which may be
// %NULL).  The correct type for the parameter is determined by a static
// parameter type (which is given at construction time).
//
// An action may optionally have a state, in which case the state may be
// set with g_action_change_state().  This call takes a #GVariant.  The
// correct type for the state is determined by a static state type
// (which is given at construction time).
//
// The state may have a hint associated with it, specifying its valid
// range.
//
// #GAction is merely the interface to the concept of an action, as
// described above.  Various implementations of actions exist, including
// #GSimpleAction.
//
// In all cases, the implementing class is responsible for storing the
// name of the action, the parameter type, the enabled state, the
// optional state type and the state and emitting the appropriate
// signals when these change.  The implementor is responsible for filtering
// calls to g_action_activate() and g_action_change_state() for type
// safety and for the state being enabled.
//
// Probably the only useful thing to do with a #GAction is to put it
// inside of a #GSimpleActionGroup.
/*

C record/class : GAction
*/
type Action struct {
	native *C.GAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GActionGroup represents a group of actions. Actions can be used to
// expose functionality in a structured way, either from one part of a
// program to another, or to the outside world. Action groups are often
// used together with a #GMenuModel that provides additional
// representation data for displaying the actions to the user, e.g. in
// a menu.
//
// The main way to interact with the actions in a GActionGroup is to
// activate them with g_action_group_activate_action(). Activating an
// action may require a #GVariant parameter. The required type of the
// parameter can be inquired with g_action_group_get_action_parameter_type().
// Actions may be disabled, see g_action_group_get_action_enabled().
// Activating a disabled action has no effect.
//
// Actions may optionally have a state in the form of a #GVariant. The
// current state of an action can be inquired with
// g_action_group_get_action_state(). Activating a stateful action may
// change its state, but it is also possible to set the state by calling
// g_action_group_change_action_state().
//
// As typical example, consider a text editing application which has an
// option to change the current font to 'bold'. A good way to represent
// this would be a stateful action, with a boolean state. Activating the
// action would toggle the state.
//
// Each action in the group has a unique name (which is a string).  All
// method calls, except g_action_group_list_actions() take the name of
// an action as an argument.
//
// The #GActionGroup API is meant to be the 'public' API to the action
// group.  The calls here are exactly the interaction that 'external
// forces' (eg: UI, incoming D-Bus messages, etc.) are supposed to have
// with actions.  'Internal' APIs (ie: ones meant only to be accessed by
// the action group implementation) are found on subclasses.  This is
// why you will find - for example - g_action_group_get_action_enabled()
// but not an equivalent set() call.
//
// Signals are emitted on the action group in response to state changes
// on individual actions.
//
// Implementations of #GActionGroup should provide implementations for
// the virtual functions g_action_group_list_actions() and
// g_action_group_query_action().  The other virtual functions should
// not be implemented - their "wrappers" are actually implemented with
// calls to g_action_group_query_action().
/*

C record/class : GActionGroup
*/
type ActionGroup struct {
	native *C.GActionGroup
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The GActionMap interface is implemented by #GActionGroup
// implementations that operate by containing a number of
// named #GAction instances, such as #GSimpleActionGroup.
//
// One useful application of this interface is to map the
// names of actions from various action groups to unique,
// prefixed names (e.g. by prepending "app." or "win.").
// This is the motivation for the 'Map' part of the interface
// name.
/*

C record/class : GActionMap
*/
type ActionMap struct {
	native *C.GActionMap
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	c := (*C.GActionMap)(u)
	if c == nil {
		return nil
	}

	g := &ActionMap{native: c}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GAppInfo and #GAppLaunchContext are used for describing and launching
// applications installed on the system.
//
// As of GLib 2.20, URIs will always be converted to POSIX paths
// (using g_file_get_path()) when using g_app_info_launch() even if
// the application requested an URI and not a POSIX path. For example
// for an desktop-file based application with Exec key `totem
// %U` and a single URI, `sftp://foo/file.avi`, then
// `/home/user/.gvfs/sftp on foo/file.avi` will be passed. This will
// only work if a set of suitable GIO extensions (such as gvfs 2.26
// compiled with FUSE support), is available and operational; if this
// is not the case, the URI will be passed unmodified to the application.
// Some URIs, such as `mailto:`, of course cannot be mapped to a POSIX
// path (in gvfs there's no FUSE mount for it); such URIs will be
// passed unmodified to the application.
//
// Specifically for gvfs 2.26 and later, the POSIX URI will be mapped
// back to the GIO URI in the #GFile constructors (since gvfs
// implements the #GVfs extension point). As such, if the application
// needs to examine the URI, it needs to use g_file_get_uri() or
// similar on #GFile. In other words, an application cannot assume
// that the URI passed to e.g. g_file_new_for_commandline_arg() is
// equal to the result of g_file_get_uri(). The following snippet
// illustrates this:
//
// |[
// GFile *f;
// char *uri;
//
// file = g_file_new_for_commandline_arg (uri_from_commandline);
//
// uri = g_file_get_uri (file);
// strcmp (uri, uri_from_commandline) == 0;
// g_free (uri);
//
// if (g_file_has_uri_scheme (file, "cdda"))
// {
// do something special with uri
// }
// g_object_unref (file);
// ]|
//
// This code will work when both `cdda://sr0/Track 1.wav` and
// `/home/user/.gvfs/cdda on sr0/Track 1.wav` is passed to the
// application. It should be noted that it's generally not safe
// for applications to rely on the format of a particular URIs.
// Different launcher applications (e.g. file managers) may have
// different ideas of what a given URI means.
/*

C record/class : GAppInfo
*/
type AppInfo struct {
	native *C.GAppInfo
}

func AppInfoNewFromC(u unsafe.Pointer) *AppInfo {
	c := (*C.GAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &AppInfo{native: c}

	return g
}

func (recv *AppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Adds a content type to the application information to indicate the
// application is capable of opening files with the given content type.
/*

C function : g_app_info_add_supports_type
*/
func (recv *AppInfo) AddSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_add_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if a supported content type can be removed from an application.
/*

C function : g_app_info_can_remove_supports_type
*/
func (recv *AppInfo) CanRemoveSupportsType() bool {
	retC := C.g_app_info_can_remove_supports_type((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Creates a duplicate of a #GAppInfo.
/*

C function : g_app_info_dup
*/
func (recv *AppInfo) Dup() *AppInfo {
	retC := C.g_app_info_dup((*C.GAppInfo)(recv.native))
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if two #GAppInfos are equal.
//
// Note that the check <emphasis>may not</emphasis> compare each individual
// field, and only does an identity check. In case detecting changes in the
// contents is needed, program code must additionally compare relevant fields.
/*

C function : g_app_info_equal
*/
func (recv *AppInfo) Equal(appinfo2 *AppInfo) bool {
	c_appinfo2 := (*C.GAppInfo)(appinfo2.ToC())

	retC := C.g_app_info_equal((*C.GAppInfo)(recv.native), c_appinfo2)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a human-readable description of an installed application.
/*

C function : g_app_info_get_description
*/
func (recv *AppInfo) GetDescription() string {
	retC := C.g_app_info_get_description((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the executable's name for the installed application.
/*

C function : g_app_info_get_executable
*/
func (recv *AppInfo) GetExecutable() string {
	retC := C.g_app_info_get_executable((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the icon for the application.
/*

C function : g_app_info_get_icon
*/
func (recv *AppInfo) GetIcon() *Icon {
	retC := C.g_app_info_get_icon((*C.GAppInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the ID of an application. An id is a string that
// identifies the application. The exact format of the id is
// platform dependent. For instance, on Unix this is the
// desktop file id from the xdg menu specification.
//
// Note that the returned ID may be %NULL, depending on how
// the @appinfo has been constructed.
/*

C function : g_app_info_get_id
*/
func (recv *AppInfo) GetId() string {
	retC := C.g_app_info_get_id((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the installed name of the application.
/*

C function : g_app_info_get_name
*/
func (recv *AppInfo) GetName() string {
	retC := C.g_app_info_get_name((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Launches the application. Passes @files to the launched application
// as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on).
// On error, @error will be set accordingly.
//
// To launch the application without arguments pass a %NULL @files list.
//
// Note that even if the launch is successful the application launched
// can fail to start if it runs into problems during startup. There is
// no way to detect this.
//
// Some URIs can be changed when passed through a GFile (for instance
// unsupported URIs with strange formats like mailto:), so if you have
// a textual URI you want to pass in as argument, consider using
// g_app_info_launch_uris() instead.
//
// The launched application inherits the environment of the launching
// process, but it can be modified with g_app_launch_context_setenv()
// and g_app_launch_context_unsetenv().
//
// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE`
// environment variable with the path of the launched desktop file and
// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`,
// should it be inherited by further processes. The `DISPLAY` and
// `DESKTOP_STARTUP_ID` environment variables are also set, based
// on information provided in @context.
/*

C function : g_app_info_launch
*/
func (recv *AppInfo) Launch(files *glib.List, context *AppLaunchContext) (bool, error) {
	c_files := (*C.GList)(C.NULL)
	if files != nil {
		c_files = (*C.GList)(files.ToC())
	}

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch((*C.GAppInfo)(recv.native), c_files, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Launches the application. This passes the @uris to the launched application
// as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on).
// On error, @error will be set accordingly.
//
// To launch the application without arguments pass a %NULL @uris list.
//
// Note that even if the launch is successful the application launched
// can fail to start if it runs into problems during startup. There is
// no way to detect this.
/*

C function : g_app_info_launch_uris
*/
func (recv *AppInfo) LaunchUris(uris *glib.List, context *AppLaunchContext) (bool, error) {
	c_uris := (*C.GList)(C.NULL)
	if uris != nil {
		c_uris = (*C.GList)(uris.ToC())
	}

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_uris((*C.GAppInfo)(recv.native), c_uris, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Removes a supported type from an application, if possible.
/*

C function : g_app_info_remove_supports_type
*/
func (recv *AppInfo) RemoveSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_remove_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the application as the default handler for the given file extension.
/*

C function : g_app_info_set_as_default_for_extension
*/
func (recv *AppInfo) SetAsDefaultForExtension(extension string) (bool, error) {
	c_extension := C.CString(extension)
	defer C.free(unsafe.Pointer(c_extension))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_extension((*C.GAppInfo)(recv.native), c_extension, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the application as the default handler for a given type.
/*

C function : g_app_info_set_as_default_for_type
*/
func (recv *AppInfo) SetAsDefaultForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the application as the last used application for a given type.
// This will make the application appear as first in the list returned
// by g_app_info_get_recommended_for_type(), regardless of the default
// application for that content type.
/*

C function : g_app_info_set_as_last_used_for_type
*/
func (recv *AppInfo) SetAsLastUsedForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_last_used_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if the application info should be shown in menus that
// list available applications.
/*

C function : g_app_info_should_show
*/
func (recv *AppInfo) ShouldShow() bool {
	retC := C.g_app_info_should_show((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the application accepts files as arguments.
/*

C function : g_app_info_supports_files
*/
func (recv *AppInfo) SupportsFiles() bool {
	retC := C.g_app_info_supports_files((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the application supports reading files and directories from URIs.
/*

C function : g_app_info_supports_uris
*/
func (recv *AppInfo) SupportsUris() bool {
	retC := C.g_app_info_supports_uris((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Provides a base class for implementing asynchronous function results.
//
// Asynchronous operations are broken up into two separate operations
// which are chained together by a #GAsyncReadyCallback. To begin
// an asynchronous operation, provide a #GAsyncReadyCallback to the
// asynchronous function. This callback will be triggered when the
// operation has completed, and must be run in a later iteration of
// the [thread-default main context][g-main-context-push-thread-default]
// from where the operation was initiated. It will be passed a
// #GAsyncResult instance filled with the details of the operation's
// success or failure, the object the asynchronous function was
// started for and any error codes returned. The asynchronous callback
// function is then expected to call the corresponding "_finish()"
// function, passing the object the function was called for, the
// #GAsyncResult instance, and (optionally) an @error to grab any
// error conditions that may have occurred.
//
// The "_finish()" function for an operation takes the generic result
// (of type #GAsyncResult) and returns the specific result that the
// operation in question yields (e.g. a #GFileEnumerator for a
// "enumerate children" operation). If the result or error status of the
// operation is not needed, there is no need to call the "_finish()"
// function; GIO will take care of cleaning up the result and error
// information after the #GAsyncReadyCallback returns. You can pass
// %NULL for the #GAsyncReadyCallback if you don't need to take any
// action at all after the operation completes. Applications may also
// take a reference to the #GAsyncResult and call "_finish()" later;
// however, the "_finish()" function may be called at most once.
//
// Example of a typical asynchronous operation flow:
// |[<!-- language="C" -->
// void _theoretical_frobnitz_async (Theoretical         *t,
// GCancellable        *c,
// GAsyncReadyCallback  cb,
// gpointer             u);
//
// gboolean _theoretical_frobnitz_finish (Theoretical   *t,
// GAsyncResult  *res,
// GError       **e);
//
// static void
// frobnitz_result_func (GObject      *source_object,
// GAsyncResult *res,
// gpointer      user_data)
// {
// gboolean success = FALSE;
//
// success = _theoretical_frobnitz_finish (source_object, res, NULL);
//
// if (success)
// g_printf ("Hurray!\n");
// else
// g_printf ("Uh oh!\n");
//
// ...
//
// }
//
// int main (int argc, void *argv[])
// {
// ...
//
// _theoretical_frobnitz_async (theoretical_data,
// NULL,
// frobnitz_result_func,
// NULL);
//
// ...
// }
// ]|
//
// The callback for an asynchronous operation is called only once, and is
// always called, even in the case of a cancelled operation. On cancellation
// the result is a %G_IO_ERROR_CANCELLED error.
//
// ## I/O Priority # {#io-priority}
//
// Many I/O-related asynchronous operations have a priority parameter,
// which is used in certain cases to determine the order in which
// operations are executed. They are not used to determine system-wide
// I/O scheduling. Priorities are integers, with lower numbers indicating
// higher priority. It is recommended to choose priorities between
// %G_PRIORITY_LOW and %G_PRIORITY_HIGH, with %G_PRIORITY_DEFAULT
// as a default.
/*

C record/class : GAsyncResult
*/
type AsyncResult struct {
	native *C.GAsyncResult
}

func AsyncResultNewFromC(u unsafe.Pointer) *AsyncResult {
	c := (*C.GAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResult{native: c}

	return g
}

func (recv *AsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets the source object from a #GAsyncResult.
/*

C function : g_async_result_get_source_object
*/
func (recv *AsyncResult) GetSourceObject() *gobject.Object {
	retC := C.g_async_result_get_source_object((*C.GAsyncResult)(recv.native))
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the user data from a #GAsyncResult.
/*

C function : g_async_result_get_user_data
*/
func (recv *AsyncResult) GetUserData() uintptr {
	retC := C.g_async_result_get_user_data((*C.GAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// The #GDBusObject type is the base type for D-Bus objects on both
// the service side (see #GDBusObjectSkeleton) and the client side
// (see #GDBusObjectProxy). It is essentially just a container of
// interfaces.
/*

C record/class : GDBusObject
*/
type DBusObject struct {
	native *C.GDBusObject
}

func DBusObjectNewFromC(u unsafe.Pointer) *DBusObject {
	c := (*C.GDBusObject)(u)
	if c == nil {
		return nil
	}

	g := &DBusObject{native: c}

	return g
}

func (recv *DBusObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GDBusObjectManager type is the base type for service- and
// client-side implementations of the standardized
// [org.freedesktop.DBus.ObjectManager](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See #GDBusObjectManagerClient for the client-side implementation
// and #GDBusObjectManagerServer for the service-side implementation.
/*

C record/class : GDBusObjectManager
*/
type DBusObjectManager struct {
	native *C.GDBusObjectManager
}

func DBusObjectManagerNewFromC(u unsafe.Pointer) *DBusObjectManager {
	c := (*C.GDBusObjectManager)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManager{native: c}

	return g
}

func (recv *DBusObjectManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GDesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
/*

C record/class : GDesktopAppInfoLookup
*/
type DesktopAppInfoLookup struct {
	native *C.GDesktopAppInfoLookup
}

func DesktopAppInfoLookupNewFromC(u unsafe.Pointer) *DesktopAppInfoLookup {
	c := (*C.GDesktopAppInfoLookup)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookup{native: c}

	return g
}

func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets the default application for launching applications
// using this URI scheme for a particular GDesktopAppInfoLookup
// implementation.
//
// The GDesktopAppInfoLookup interface and this function is used
// to implement g_app_info_get_default_for_uri_scheme() backends
// in a GIO module. There is no reason for applications to use it
// directly. Applications should use g_app_info_get_default_for_uri_scheme().
/*

C function : g_desktop_app_info_lookup_get_default_for_uri_scheme
*/
func (recv *DesktopAppInfoLookup) GetDefaultForUriScheme(uriScheme string) *AppInfo {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_desktop_app_info_lookup_get_default_for_uri_scheme((*C.GDesktopAppInfoLookup)(recv.native), c_uri_scheme)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// #GDrive - this represent a piece of hardware connected to the machine.
// It's generally only created for removable hardware or hardware with
// removable media.
//
// #GDrive is a container class for #GVolume objects that stem from
// the same piece of media. As such, #GDrive abstracts a drive with
// (or without) removable media and provides operations for querying
// whether media is available, determining whether media change is
// automatically detected and ejecting the media.
//
// If the #GDrive reports that media isn't automatically detected, one
// can poll for media; typically one should not do this periodically
// as a poll for media operation is potententially expensive and may
// spin up the drive creating noise.
//
// #GDrive supports starting and stopping drives with authentication
// support for the former. This can be used to support a diverse set
// of use cases including connecting/disconnecting iSCSI devices,
// powering down external disk enclosures and starting/stopping
// multi-disk devices such as RAID devices. Note that the actual
// semantics and side-effects of starting/stopping a #GDrive may vary
// according to implementation. To choose the correct verbs in e.g. a
// file manager, use g_drive_get_start_stop_type().
//
// For porting from GnomeVFS note that there is no equivalent of
// #GDrive in that API.
/*

C record/class : GDrive
*/
type Drive struct {
	native *C.GDrive
}

func DriveNewFromC(u unsafe.Pointer) *Drive {
	c := (*C.GDrive)(u)
	if c == nil {
		return nil
	}

	g := &Drive{native: c}

	return g
}

func (recv *Drive) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalDriveChangedDetail struct {
	callback  DriveSignalChangedCallback
	handlerID C.gulong
}

var signalDriveChangedId int
var signalDriveChangedMap = make(map[int]signalDriveChangedDetail)
var signalDriveChangedLock sync.Mutex

// DriveSignalChangedCallback is a callback function for a 'changed' signal emitted from a Drive.
type DriveSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Drive) ConnectChanged(callback DriveSignalChangedCallback) int {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	signalDriveChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_changed(instance, C.gpointer(uintptr(signalDriveChangedId)))

	detail := signalDriveChangedDetail{callback, handlerID}
	signalDriveChangedMap[signalDriveChangedId] = detail

	return signalDriveChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Drive.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Drive) DisconnectChanged(connectionID int) {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	detail, exists := signalDriveChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveChangedMap, connectionID)
}

//export drive_changedHandler
func drive_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDriveChangedMap[index].callback
	callback()
}

type signalDriveDisconnectedDetail struct {
	callback  DriveSignalDisconnectedCallback
	handlerID C.gulong
}

var signalDriveDisconnectedId int
var signalDriveDisconnectedMap = make(map[int]signalDriveDisconnectedDetail)
var signalDriveDisconnectedLock sync.Mutex

// DriveSignalDisconnectedCallback is a callback function for a 'disconnected' signal emitted from a Drive.
type DriveSignalDisconnectedCallback func()

/*
ConnectDisconnected connects the callback to the 'disconnected' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectDisconnected to remove it.
*/
func (recv *Drive) ConnectDisconnected(callback DriveSignalDisconnectedCallback) int {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	signalDriveDisconnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_disconnected(instance, C.gpointer(uintptr(signalDriveDisconnectedId)))

	detail := signalDriveDisconnectedDetail{callback, handlerID}
	signalDriveDisconnectedMap[signalDriveDisconnectedId] = detail

	return signalDriveDisconnectedId
}

/*
DisconnectDisconnected disconnects a callback from the 'disconnected' signal for the Drive.

The connectionID should be a value returned from a call to ConnectDisconnected.
*/
func (recv *Drive) DisconnectDisconnected(connectionID int) {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	detail, exists := signalDriveDisconnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveDisconnectedMap, connectionID)
}

//export drive_disconnectedHandler
func drive_disconnectedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDriveDisconnectedMap[index].callback
	callback()
}

type signalDriveEjectButtonDetail struct {
	callback  DriveSignalEjectButtonCallback
	handlerID C.gulong
}

var signalDriveEjectButtonId int
var signalDriveEjectButtonMap = make(map[int]signalDriveEjectButtonDetail)
var signalDriveEjectButtonLock sync.Mutex

// DriveSignalEjectButtonCallback is a callback function for a 'eject-button' signal emitted from a Drive.
type DriveSignalEjectButtonCallback func()

/*
ConnectEjectButton connects the callback to the 'eject-button' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectEjectButton to remove it.
*/
func (recv *Drive) ConnectEjectButton(callback DriveSignalEjectButtonCallback) int {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	signalDriveEjectButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_eject_button(instance, C.gpointer(uintptr(signalDriveEjectButtonId)))

	detail := signalDriveEjectButtonDetail{callback, handlerID}
	signalDriveEjectButtonMap[signalDriveEjectButtonId] = detail

	return signalDriveEjectButtonId
}

/*
DisconnectEjectButton disconnects a callback from the 'eject-button' signal for the Drive.

The connectionID should be a value returned from a call to ConnectEjectButton.
*/
func (recv *Drive) DisconnectEjectButton(connectionID int) {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	detail, exists := signalDriveEjectButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveEjectButtonMap, connectionID)
}

//export drive_ejectButtonHandler
func drive_ejectButtonHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDriveEjectButtonMap[index].callback
	callback()
}

// Checks if a drive can be ejected.
/*

C function : g_drive_can_eject
*/
func (recv *Drive) CanEject() bool {
	retC := C.g_drive_can_eject((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a drive can be polled for media changes.
/*

C function : g_drive_can_poll_for_media
*/
func (recv *Drive) CanPollForMedia() bool {
	retC := C.g_drive_can_poll_for_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a drive.
/*

C function : g_drive_eject_finish
*/
func (recv *Drive) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_drive_enumerate_identifiers : no return type

// Gets the icon for @drive.
/*

C function : g_drive_get_icon
*/
func (recv *Drive) GetIcon() *Icon {
	retC := C.g_drive_get_icon((*C.GDrive)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the identifier of the given kind for @drive.
/*

C function : g_drive_get_identifier
*/
func (recv *Drive) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_drive_get_identifier((*C.GDrive)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of @drive.
/*

C function : g_drive_get_name
*/
func (recv *Drive) GetName() string {
	retC := C.g_drive_get_name((*C.GDrive)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Get a list of mountable volumes for @drive.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
/*

C function : g_drive_get_volumes
*/
func (recv *Drive) GetVolumes() *glib.List {
	retC := C.g_drive_get_volumes((*C.GDrive)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if the @drive has media. Note that the OS may not be polling
// the drive for media changes; see g_drive_is_media_check_automatic()
// for more details.
/*

C function : g_drive_has_media
*/
func (recv *Drive) HasMedia() bool {
	retC := C.g_drive_has_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Check if @drive has any mountable volumes.
/*

C function : g_drive_has_volumes
*/
func (recv *Drive) HasVolumes() bool {
	retC := C.g_drive_has_volumes((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @drive is capabable of automatically detecting media changes.
/*

C function : g_drive_is_media_check_automatic
*/
func (recv *Drive) IsMediaCheckAutomatic() bool {
	retC := C.g_drive_is_media_check_automatic((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the @drive supports removable media.
/*

C function : g_drive_is_media_removable
*/
func (recv *Drive) IsMediaRemovable() bool {
	retC := C.g_drive_is_media_removable((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_poll_for_media : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_drive_poll_for_media() on a drive.
/*

C function : g_drive_poll_for_media_finish
*/
func (recv *Drive) PollForMediaFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_poll_for_media_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GFile is a high level abstraction for manipulating files on a
// virtual file system. #GFiles are lightweight, immutable objects
// that do no I/O upon creation. It is necessary to understand that
// #GFile objects do not represent files, merely an identifier for a
// file. All file content I/O is implemented as streaming operations
// (see #GInputStream and #GOutputStream).
//
// To construct a #GFile, you can use:
// - g_file_new_for_path() if you have a path.
// - g_file_new_for_uri() if you have a URI.
// - g_file_new_for_commandline_arg() for a command line argument.
// - g_file_new_tmp() to create a temporary file from a template.
// - g_file_parse_name() from a UTF-8 string gotten from g_file_get_parse_name().
// - g_file_new_build_filename() to create a file from path elements.
//
// One way to think of a #GFile is as an abstraction of a pathname. For
// normal files the system pathname is what is stored internally, but as
// #GFiles are extensible it could also be something else that corresponds
// to a pathname in a userspace implementation of a filesystem.
//
// #GFiles make up hierarchies of directories and files that correspond to
// the files on a filesystem. You can move through the file system with
// #GFile using g_file_get_parent() to get an identifier for the parent
// directory, g_file_get_child() to get a child within a directory,
// g_file_resolve_relative_path() to resolve a relative path between two
// #GFiles. There can be multiple hierarchies, so you may not end up at
// the same root if you repeatedly call g_file_get_parent() on two different
// files.
//
// All #GFiles have a basename (get with g_file_get_basename()). These names
// are byte strings that are used to identify the file on the filesystem
// (relative to its parent directory) and there is no guarantees that they
// have any particular charset encoding or even make any sense at all. If
// you want to use filenames in a user interface you should use the display
// name that you can get by requesting the
// %G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME attribute with g_file_query_info().
// This is guaranteed to be in UTF-8 and can be used in a user interface.
// But always store the real basename or the #GFile to use to actually
// access the file, because there is no way to go from a display name to
// the actual name.
//
// Using #GFile as an identifier has the same weaknesses as using a path
// in that there may be multiple aliases for the same file. For instance,
// hard or soft links may cause two different #GFiles to refer to the same
// file. Other possible causes for aliases are: case insensitive filesystems,
// short and long names on FAT/NTFS, or bind mounts in Linux. If you want to
// check if two #GFiles point to the same file you can query for the
// %G_FILE_ATTRIBUTE_ID_FILE attribute. Note that #GFile does some trivial
// canonicalization of pathnames passed in, so that trivial differences in
// the path string used at creation (duplicated slashes, slash at end of
// path, "." or ".." path segments, etc) does not create different #GFiles.
//
// Many #GFile operations have both synchronous and asynchronous versions
// to suit your application. Asynchronous versions of synchronous functions
// simply have _async() appended to their function names. The asynchronous
// I/O functions call a #GAsyncReadyCallback which is then used to finalize
// the operation, producing a GAsyncResult which is then passed to the
// function's matching _finish() operation.
//
// It is highly recommended to use asynchronous calls when running within a
// shared main loop, such as in the main thread of an application. This avoids
// I/O operations blocking other sources on the main loop from being dispatched.
// Synchronous I/O operations should be performed from worker threads. See the
// [introduction to asynchronous programming section][async-programming] for
// more.
//
// Some #GFile operations almost always take a noticeable amount of time, and
// so do not have synchronous analogs. Notable cases include:
// - g_file_mount_mountable() to mount a mountable file.
// - g_file_unmount_mountable_with_operation() to unmount a mountable file.
// - g_file_eject_mountable_with_operation() to eject a mountable file.
//
// ## Entity Tags # {#gfile-etag}
//
// One notable feature of #GFiles are entity tags, or "etags" for
// short. Entity tags are somewhat like a more abstract version of the
// traditional mtime, and can be used to quickly determine if the file
// has been modified from the version on the file system. See the
// HTTP 1.1
// [specification](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html)
// for HTTP Etag headers, which are a very similar concept.
/*

C record/class : GFile
*/
type File struct {
	native *C.GFile
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets an output stream for appending data to the file.
// If the file doesn't already exist it is created.
//
// By default files created are generally readable by everyone,
// but if you pass #G_FILE_CREATE_PRIVATE in @flags the file
// will be made readable only to the current user, to the level that
// is supported on the target filesystem.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// Some file systems don't allow all file names, and may return an
// %G_IO_ERROR_INVALID_FILENAME error. If the file is a directory the
// %G_IO_ERROR_IS_DIRECTORY error will be returned. Other errors are
// possible too, and depend on what kind of filesystem the file is on.
/*

C function : g_file_append_to
*/
func (recv *File) AppendTo(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_append_to((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_append_to_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file append operation started with
// g_file_append_to_async().
/*

C function : g_file_append_to_finish
*/
func (recv *File) AppendToFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_append_to_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_copy : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Unsupported : g_file_copy_async : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Copies the file attributes from @source to @destination.
//
// Normally only a subset of the file attributes are copied,
// those that are copies in a normal file copy operation
// (which for instance does not include e.g. owner). However
// if #G_FILE_COPY_ALL_METADATA is specified in @flags, then
// all the metadata that is possible to copy is copied. This
// is useful when implementing move by copy + delete source.
/*

C function : g_file_copy_attributes
*/
func (recv *File) CopyAttributes(destination *File, flags FileCopyFlags, cancellable *Cancellable) (bool, error) {
	c_destination := (*C.GFile)(destination.ToC())

	c_flags := (C.GFileCopyFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_copy_attributes((*C.GFile)(recv.native), c_destination, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes copying the file started with g_file_copy_async().
/*

C function : g_file_copy_finish
*/
func (recv *File) CopyFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_copy_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new file and returns an output stream for writing to it.
// The file must not already exist.
//
// By default files created are generally readable by everyone,
// but if you pass #G_FILE_CREATE_PRIVATE in @flags the file
// will be made readable only to the current user, to the level
// that is supported on the target filesystem.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If a file or directory with this name already exists the
// %G_IO_ERROR_EXISTS error will be returned. Some file systems don't
// allow all file names, and may return an %G_IO_ERROR_INVALID_FILENAME
// error, and if the name is to long %G_IO_ERROR_FILENAME_TOO_LONG will
// be returned. Other errors are possible too, and depend on what kind
// of filesystem the file is on.
/*

C function : g_file_create
*/
func (recv *File) Create(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_create((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_create_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file create operation started with
// g_file_create_async().
/*

C function : g_file_create_finish
*/
func (recv *File) CreateFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Deletes a file. If the @file is a directory, it will only be
// deleted if it is empty. This has the same semantics as g_unlink().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_delete
*/
func (recv *File) Delete(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_delete((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Duplicates a #GFile handle. This operation does not duplicate
// the actual file or directory represented by the #GFile; see
// g_file_copy() if attempting to copy a file.
//
// This call does no blocking I/O.
/*

C function : g_file_dup
*/
func (recv *File) Dup() *File {
	retC := C.g_file_dup((*C.GFile)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_eject_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous eject operation started by
// g_file_eject_mountable().
/*

C function : g_file_eject_mountable_finish
*/
func (recv *File) EjectMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the requested information about the files in a directory.
// The result is a #GFileEnumerator object that will give out
// #GFileInfo objects for all the files in the directory.
//
// The @attributes value is a string that specifies the file
// attributes that should be gathered. It is not an error if
// it's not possible to read a particular requested attribute
// from a file - it just won't be set. @attributes should
// be a comma-separated list of attributes or attribute wildcards.
// The wildcard "*" means all attributes, and a wildcard like
// "standard::*" means all attributes in the standard namespace.
// An example attribute query be "standard::*,owner::user".
// The standard attributes are available as defines, like
// #G_FILE_ATTRIBUTE_STANDARD_NAME.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If the file does not exist, the %G_IO_ERROR_NOT_FOUND error will
// be returned. If the file is not a directory, the %G_IO_ERROR_NOT_DIRECTORY
// error will be returned. Other errors are possible too.
/*

C function : g_file_enumerate_children
*/
func (recv *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileEnumerator, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_enumerate_children_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async enumerate children operation.
// See g_file_enumerate_children_async().
/*

C function : g_file_enumerate_children_finish
*/
func (recv *File) EnumerateChildrenFinish(res *AsyncResult) (*FileEnumerator, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if the two given #GFiles refer to the same file.
//
// Note that two #GFiles that differ can still refer to the same
// file on the filesystem due to various forms of filename
// aliasing.
//
// This call does no blocking I/O.
/*

C function : g_file_equal
*/
func (recv *File) Equal(file2 *File) bool {
	c_file2 := (*C.GFile)(file2.ToC())

	retC := C.g_file_equal((*C.GFile)(recv.native), c_file2)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a #GMount for the #GFile.
//
// If the #GFileIface for @file does not have a mount (e.g.
// possibly a remote share), @error will be set to %G_IO_ERROR_NOT_FOUND
// and %NULL will be returned.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_find_enclosing_mount
*/
func (recv *File) FindEnclosingMount(cancellable *Cancellable) (*Mount, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_find_enclosing_mount_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous find mount request.
// See g_file_find_enclosing_mount_async().
/*

C function : g_file_find_enclosing_mount_finish
*/
func (recv *File) FindEnclosingMountFinish(res *AsyncResult) (*Mount, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the base name (the last component of the path) for a given #GFile.
//
// If called for the top level of a system (such as the filesystem root
// or a uri like sftp://host/) it will return a single directory separator
// (and on Windows, possibly a drive letter).
//
// The base name is a byte string (not UTF-8). It has no defined encoding
// or rules other than it may not contain zero bytes.  If you want to use
// filenames in a user interface you should use the display name that you
// can get by requesting the %G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME
// attribute with g_file_query_info().
//
// This call does no blocking I/O.
/*

C function : g_file_get_basename
*/
func (recv *File) GetBasename() string {
	retC := C.g_file_get_basename((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets a child of @file with basename equal to @name.
//
// Note that the file with that specific name might not exist, but
// you can still have a #GFile that points to it. You can use this
// for instance to create that file.
//
// This call does no blocking I/O.
/*

C function : g_file_get_child
*/
func (recv *File) GetChild(name string) *File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_file_get_child((*C.GFile)(recv.native), c_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the child of @file for a given @display_name (i.e. a UTF-8
// version of the name). If this function fails, it returns %NULL
// and @error will be set. This is very useful when constructing a
// #GFile for a new file and the user entered the filename in the
// user interface, for instance when you select a directory and
// type a filename in the file selector.
//
// This call does no blocking I/O.
/*

C function : g_file_get_child_for_display_name
*/
func (recv *File) GetChildForDisplayName(displayName string) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	var cThrowableError *C.GError

	retC := C.g_file_get_child_for_display_name((*C.GFile)(recv.native), c_display_name, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the parent directory for the @file.
// If the @file represents the root directory of the
// file system, then %NULL will be returned.
//
// This call does no blocking I/O.
/*

C function : g_file_get_parent
*/
func (recv *File) GetParent() *File {
	retC := C.g_file_get_parent((*C.GFile)(recv.native))
	var retGo (*File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the parse name of the @file.
// A parse name is a UTF-8 string that describes the
// file such that one can get the #GFile back using
// g_file_parse_name().
//
// This is generally used to show the #GFile as a nice
// full-pathname kind of string in a user interface,
// like in a location entry.
//
// For local files with names that can safely be converted
// to UTF-8 the pathname is used, otherwise the IRI is used
// (a form of URI that allows UTF-8 characters unescaped).
//
// This call does no blocking I/O.
/*

C function : g_file_get_parse_name
*/
func (recv *File) GetParseName() string {
	retC := C.g_file_get_parse_name((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the local pathname for #GFile, if one exists. If non-%NULL, this is
// guaranteed to be an absolute, canonical path. It might contain symlinks.
//
// This call does no blocking I/O.
/*

C function : g_file_get_path
*/
func (recv *File) GetPath() string {
	retC := C.g_file_get_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the path for @descendant relative to @parent.
//
// This call does no blocking I/O.
/*

C function : g_file_get_relative_path
*/
func (recv *File) GetRelativePath(descendant *File) string {
	c_descendant := (*C.GFile)(descendant.ToC())

	retC := C.g_file_get_relative_path((*C.GFile)(recv.native), c_descendant)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the URI for the @file.
//
// This call does no blocking I/O.
/*

C function : g_file_get_uri
*/
func (recv *File) GetUri() string {
	retC := C.g_file_get_uri((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the URI scheme for a #GFile.
// RFC 3986 decodes the scheme as:
// |[
// URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]
// ]|
// Common schemes include "file", "http", "ftp", etc.
//
// This call does no blocking I/O.
/*

C function : g_file_get_uri_scheme
*/
func (recv *File) GetUriScheme() string {
	retC := C.g_file_get_uri_scheme((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Checks whether @file has the prefix specified by @prefix.
//
// In other words, if the names of initial elements of @file's
// pathname match @prefix. Only full pathname elements are matched,
// so a path like /foo is not considered a prefix of /foobar, only
// of /foo/bar.
//
// A #GFile is not a prefix of itself. If you want to check for
// equality, use g_file_equal().
//
// This call does no I/O, as it works purely on names. As such it can
// sometimes return %FALSE even if @file is inside a @prefix (from a
// filesystem point of view), because the prefix of @file is an alias
// of @prefix.
/*

C function : g_file_has_prefix
*/
func (recv *File) HasPrefix(prefix *File) bool {
	c_prefix := (*C.GFile)(prefix.ToC())

	retC := C.g_file_has_prefix((*C.GFile)(recv.native), c_prefix)
	retGo := retC == C.TRUE

	return retGo
}

// Checks to see if a #GFile has a given URI scheme.
//
// This call does no blocking I/O.
/*

C function : g_file_has_uri_scheme
*/
func (recv *File) HasUriScheme(uriScheme string) bool {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_file_has_uri_scheme((*C.GFile)(recv.native), c_uri_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// Creates a hash value for a #GFile.
//
// This call does no blocking I/O.
/*

C function : g_file_hash
*/
func (recv *File) Hash() uint32 {
	retC := C.g_file_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Checks to see if a file is native to the platform.
//
// A native file s one expressed in the platform-native filename format,
// e.g. "C:\Windows" or "/usr/bin/". This does not mean the file is local,
// as it might be on a locally mounted remote filesystem.
//
// On some systems non-native files may be available using the native
// filesystem via a userspace filesystem (FUSE), in these cases this call
// will return %FALSE, but g_file_get_path() will still return a native path.
//
// This call does no blocking I/O.
/*

C function : g_file_is_native
*/
func (recv *File) IsNative() bool {
	retC := C.g_file_is_native((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_load_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_load_contents_finish : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_partial_contents_async : unsupported parameter read_more_callback : no type generator for FileReadMoreCallback (GFileReadMoreCallback) for param read_more_callback

// Unsupported : g_file_load_partial_contents_finish : unsupported parameter contents : output array param contents

// Creates a directory. Note that this will only create a child directory
// of the immediate parent directory of the path or URI given by the #GFile.
// To recursively create directories, see g_file_make_directory_with_parents().
// This function will fail if the parent directory does not exist, setting
// @error to %G_IO_ERROR_NOT_FOUND. If the file system doesn't support
// creating directories, this function will fail, setting @error to
// %G_IO_ERROR_NOT_SUPPORTED.
//
// For a local #GFile the newly created directory will have the default
// (current) ownership and permissions of the current process.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_make_directory
*/
func (recv *File) MakeDirectory(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_directory((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a symbolic link named @file which contains the string
// @symlink_value.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_make_symbolic_link
*/
func (recv *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable) (bool, error) {
	c_symlink_value := C.CString(symlinkValue)
	defer C.free(unsafe.Pointer(c_symlink_value))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_symbolic_link((*C.GFile)(recv.native), c_symlink_value, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Obtains a directory monitor for the given file.
// This may fail if directory monitoring is not supported.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
//
// It does not make sense for @flags to contain
// %G_FILE_MONITOR_WATCH_HARD_LINKS, since hard links can not be made to
// directories.  It is not possible to monitor all the files in a
// directory for changes made via hard links; if you want to do this then
// you must register individual watches with g_file_monitor().
/*

C function : g_file_monitor_directory
*/
func (recv *File) MonitorDirectory(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor_directory((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Obtains a file monitor for the given file. If no file notification
// mechanism exists, then regular polling of the file is used.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
//
// If @flags contains %G_FILE_MONITOR_WATCH_HARD_LINKS then the monitor
// will also attempt to report changes made to the file via another
// filename (ie, a hard link). Without this flag, you can only rely on
// changes made through the filename contained in @file to be
// reported. Using this flag may result in an increase in resource
// usage, and may not have any effect depending on the #GFileMonitor
// backend and/or filesystem type.
/*

C function : g_file_monitor_file
*/
func (recv *File) MonitorFile(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor_file((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_mount_enclosing_volume : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a mount operation started by g_file_mount_enclosing_volume().
/*

C function : g_file_mount_enclosing_volume_finish
*/
func (recv *File) MountEnclosingVolumeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_enclosing_volume_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_mount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a mount operation. See g_file_mount_mountable() for details.
//
// Finish an asynchronous mount operation that was started
// with g_file_mount_mountable().
/*

C function : g_file_mount_mountable_finish
*/
func (recv *File) MountMountableFinish(result *AsyncResult) (*File, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_move : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Returns the #GAppInfo that is registered as the default
// application to handle the file specified by @file.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_query_default_handler
*/
func (recv *File) QueryDefaultHandler(cancellable *Cancellable) (*AppInfo, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_default_handler((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Utility function to check if a particular file exists. This is
// implemented using g_file_query_info() and as such does blocking I/O.
//
// Note that in many cases it is [racy to first check for file existence](https://en.wikipedia.org/wiki/Time_of_check_to_time_of_use)
// and then execute something based on the outcome of that, because the
// file might have been created or removed in between the operations. The
// general approach to handling that is to not check, but just do the
// operation and handle the errors as they come.
//
// As an example of race-free checking, take the case of reading a file,
// and if it doesn't exist, creating it. There are two racy versions: read
// it, and on error create it; and: check if it exists, if not create it.
// These can both result in two processes creating the file (with perhaps
// a partially written file as the result). The correct approach is to
// always try to create the file with g_file_create() which will either
// atomically create the file or fail with a %G_IO_ERROR_EXISTS error.
//
// However, in many cases an existence check is useful in a user interface,
// for instance to make a menu item sensitive/insensitive, so that you don't
// have to fool users that something is possible and then just show an error
// dialog. If you do this, you should make sure to also handle the errors
// that can happen due to races when you execute the operation.
/*

C function : g_file_query_exists
*/
func (recv *File) QueryExists(cancellable *Cancellable) bool {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_file_query_exists((*C.GFile)(recv.native), c_cancellable)
	retGo := retC == C.TRUE

	return retGo
}

// Similar to g_file_query_info(), but obtains information
// about the filesystem the @file is on, rather than the file itself.
// For instance the amount of space available and the type of
// the filesystem.
//
// The @attributes value is a string that specifies the attributes
// that should be gathered. It is not an error if it's not possible
// to read a particular requested attribute from a file - it just
// won't be set. @attributes should be a comma-separated list of
// attributes or attribute wildcards. The wildcard "*" means all
// attributes, and a wildcard like "filesystem::*" means all attributes
// in the filesystem namespace. The standard namespace for filesystem
// attributes is "filesystem". Common attributes of interest are
// #G_FILE_ATTRIBUTE_FILESYSTEM_SIZE (the total size of the filesystem
// in bytes), #G_FILE_ATTRIBUTE_FILESYSTEM_FREE (number of bytes available),
// and #G_FILE_ATTRIBUTE_FILESYSTEM_TYPE (type of the filesystem).
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If the file does not exist, the %G_IO_ERROR_NOT_FOUND error will
// be returned. Other errors are possible too, and depend on what
// kind of filesystem the file is on.
/*

C function : g_file_query_filesystem_info
*/
func (recv *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info((*C.GFile)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_query_filesystem_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous filesystem info query.
// See g_file_query_filesystem_info_async().
/*

C function : g_file_query_filesystem_info_finish
*/
func (recv *File) QueryFilesystemInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the requested information about specified @file.
// The result is a #GFileInfo object that contains key-value
// attributes (such as the type or size of the file).
//
// The @attributes value is a string that specifies the file
// attributes that should be gathered. It is not an error if
// it's not possible to read a particular requested attribute
// from a file - it just won't be set. @attributes should be a
// comma-separated list of attributes or attribute wildcards.
// The wildcard "*" means all attributes, and a wildcard like
// "standard::*" means all attributes in the standard namespace.
// An example attribute query be "standard::*,owner::user".
// The standard attributes are available as defines, like
// #G_FILE_ATTRIBUTE_STANDARD_NAME.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// For symlinks, normally the information about the target of the
// symlink is returned, rather than information about the symlink
// itself. However if you pass #G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS
// in @flags the information about the symlink itself will be returned.
// Also, for symlinks that point to non-existing files the information
// about the symlink itself will be returned.
//
// If the file does not exist, the %G_IO_ERROR_NOT_FOUND error will be
// returned. Other errors are possible too, and depend on what kind of
// filesystem the file is on.
/*

C function : g_file_query_info
*/
func (recv *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_info((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file info query.
// See g_file_query_info_async().
/*

C function : g_file_query_info_finish
*/
func (recv *File) QueryInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Obtain the list of settable attributes for the file.
//
// Returns the type and full attribute name of all the attributes
// that can be set on this file. This doesn't mean setting it will
// always succeed though, you might get an access failure, or some
// specific file may not support a specific attribute.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_query_settable_attributes
*/
func (recv *File) QuerySettableAttributes(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_settable_attributes((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Obtain the list of attribute namespaces where new attributes
// can be created by a user. An example of this is extended
// attributes (in the "xattr" namespace).
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_query_writable_namespaces
*/
func (recv *File) QueryWritableNamespaces(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_writable_namespaces((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Opens a file for reading. The result is a #GFileInputStream that
// can be used to read the contents of the file.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
//
// If the file does not exist, the %G_IO_ERROR_NOT_FOUND error will be
// returned. If the file is a directory, the %G_IO_ERROR_IS_DIRECTORY
// error will be returned. Other errors are possible too, and depend
// on what kind of filesystem the file is on.
/*

C function : g_file_read
*/
func (recv *File) Read(cancellable *Cancellable) (*FileInputStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_read((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous file read operation started with
// g_file_read_async().
/*

C function : g_file_read_finish
*/
func (recv *File) ReadFinish(res *AsyncResult) (*FileInputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_read_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns an output stream for overwriting the file, possibly
// creating a backup copy of the file first. If the file doesn't exist,
// it will be created.
//
// This will try to replace the file in the safest way possible so
// that any errors during the writing will not affect an already
// existing copy of the file. For instance, for local files it
// may write to a temporary file and then atomically rename over
// the destination when the stream is closed.
//
// By default files created are generally readable by everyone,
// but if you pass #G_FILE_CREATE_PRIVATE in @flags the file
// will be made readable only to the current user, to the level that
// is supported on the target filesystem.
//
// If @cancellable is not %NULL, then the operation can be cancelled
// by triggering the cancellable object from another thread. If the
// operation was cancelled, the error %G_IO_ERROR_CANCELLED will be
// returned.
//
// If you pass in a non-%NULL @etag value and @file already exists, then
// this value is compared to the current entity tag of the file, and if
// they differ an %G_IO_ERROR_WRONG_ETAG error is returned. This
// generally means that the file has been changed since you last read
// it. You can get the new etag from g_file_output_stream_get_etag()
// after you've finished writing and closed the #GFileOutputStream. When
// you load a new file you can use g_file_input_stream_query_info() to
// get the etag of the file.
//
// If @make_backup is %TRUE, this function will attempt to make a
// backup of the current file before overwriting it. If this fails
// a %G_IO_ERROR_CANT_CREATE_BACKUP error will be returned. If you
// want to replace anyway, try again with @make_backup set to %FALSE.
//
// If the file is a directory the %G_IO_ERROR_IS_DIRECTORY error will
// be returned, and if the file is some other form of non-regular file
// then a %G_IO_ERROR_NOT_REGULAR_FILE error will be returned. Some
// file systems don't allow all file names, and may return an
// %G_IO_ERROR_INVALID_FILENAME error, and if the name is to long
// %G_IO_ERROR_FILENAME_TOO_LONG will be returned. Other errors are
// possible too, and depend on what kind of filesystem the file is on.
/*

C function : g_file_replace
*/
func (recv *File) Replace(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace((*C.GFile)(recv.native), c_etag, c_make_backup, c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_replace_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Replaces the contents of @file with @contents of @length bytes.
//
// If @etag is specified (not %NULL), any existing file must have that etag,
// or the error %G_IO_ERROR_WRONG_ETAG will be returned.
//
// If @make_backup is %TRUE, this function will attempt to make a backup
// of @file. Internally, it uses g_file_replace(), so will try to replace the
// file contents in the safest way possible. For example, atomic renames are
// used when replacing local files’ contents.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
//
// The returned @new_etag can be used to verify that the file hasn't
// changed the next time it is saved over.
/*

C function : g_file_replace_contents
*/
func (recv *File) ReplaceContents(contents []uint8, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (bool, string, error) {
	c_contents := &contents[0]

	c_length := (C.gsize)(len(contents))

	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	var c_new_etag *C.char

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents((*C.GFile)(recv.native), (*C.char)(unsafe.Pointer(c_contents)), c_length, c_etag, c_make_backup, c_flags, &c_new_etag, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goThrowableError
}

// Unsupported : g_file_replace_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous replace of the given @file. See
// g_file_replace_contents_async(). Sets @new_etag to the new entity
// tag for the document, if present.
/*

C function : g_file_replace_contents_finish
*/
func (recv *File) ReplaceContentsFinish(res *AsyncResult) (bool, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_new_etag *C.char

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents_finish((*C.GFile)(recv.native), c_res, &c_new_etag, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goThrowableError
}

// Finishes an asynchronous file replace operation started with
// g_file_replace_async().
/*

C function : g_file_replace_finish
*/
func (recv *File) ReplaceFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Resolves a relative path for @file to an absolute path.
//
// This call does no blocking I/O.
/*

C function : g_file_resolve_relative_path
*/
func (recv *File) ResolveRelativePath(relativePath string) *File {
	c_relative_path := C.CString(relativePath)
	defer C.free(unsafe.Pointer(c_relative_path))

	retC := C.g_file_resolve_relative_path((*C.GFile)(recv.native), c_relative_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets an attribute in the file with attribute name @attribute to @value.
//
// Some attributes can be unset by setting @type to
// %G_FILE_ATTRIBUTE_TYPE_INVALID and @value_p to %NULL.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute
*/
func (recv *File) SetAttribute(attribute string, type_ FileAttributeType, valueP uintptr, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_type := (C.GFileAttributeType)(type_)

	c_value_p := (C.gpointer)(valueP)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute((*C.GFile)(recv.native), c_attribute, c_type, c_value_p, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_BYTE_STRING to @value.
// If @attribute is of a different type, this operation will fail,
// returning %FALSE.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_byte_string
*/
func (recv *File) SetAttributeByteString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_byte_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_INT32 to @value.
// If @attribute is of a different type, this operation will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_int32
*/
func (recv *File) SetAttributeInt32(attribute string, value int32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_INT64 to @value.
// If @attribute is of a different type, this operation will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_int64
*/
func (recv *File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_STRING to @value.
// If @attribute is of a different type, this operation will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_string
*/
func (recv *File) SetAttributeString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_UINT32 to @value.
// If @attribute is of a different type, this operation will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_uint32
*/
func (recv *File) SetAttributeUint32(attribute string, value uint32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @attribute of type %G_FILE_ATTRIBUTE_TYPE_UINT64 to @value.
// If @attribute is of a different type, this operation will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attribute_uint64
*/
func (recv *File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_set_attributes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes setting an attribute started in g_file_set_attributes_async().
/*

C function : g_file_set_attributes_finish
*/
func (recv *File) SetAttributesFinish(result *AsyncResult) (bool, *FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_info *C.GFileInfo

	var cThrowableError *C.GError

	retC := C.g_file_set_attributes_finish((*C.GFile)(recv.native), c_result, &c_info, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	info := FileInfoNewFromC(unsafe.Pointer(c_info))

	return retGo, info, goThrowableError
}

// Tries to set all attributes in the #GFileInfo on the target
// values, not stopping on the first error.
//
// If there is any error during this operation then @error will
// be set to the first error. Error on particular fields are flagged
// by setting the "status" field in the attribute value to
// %G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING, which means you can
// also detect further errors.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_attributes_from_info
*/
func (recv *File) SetAttributesFromInfo(info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_info := (*C.GFileInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GFileInfo)(info.ToC())
	}

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attributes_from_info((*C.GFile)(recv.native), c_info, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Renames @file to the specified display name.
//
// The display name is converted from UTF-8 to the correct encoding
// for the target filesystem if possible and the @file is renamed to this.
//
// If you want to implement a rename operation in the user interface the
// edit name (#G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME) should be used as the
// initial value in the rename widget, and then the result after editing
// should be passed to g_file_set_display_name().
//
// On success the resulting converted filename is returned.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_set_display_name
*/
func (recv *File) SetDisplayName(displayName string, cancellable *Cancellable) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name((*C.GFile)(recv.native), c_display_name, c_cancellable, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_set_display_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes setting a display name started with
// g_file_set_display_name_async().
/*

C function : g_file_set_display_name_finish
*/
func (recv *File) SetDisplayNameFinish(res *AsyncResult) (*File, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sends @file to the "Trashcan", if possible. This is similar to
// deleting it, but the user can recover it before emptying the trashcan.
// Not all file systems support trashing, so this call can return the
// %G_IO_ERROR_NOT_SUPPORTED error.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_file_trash
*/
func (recv *File) Trash(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_trash((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_unmount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an unmount operation, see g_file_unmount_mountable() for details.
//
// Finish an asynchronous unmount operation that was started
// with g_file_unmount_mountable().
/*

C function : g_file_unmount_mountable_finish
*/
func (recv *File) UnmountMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GFileDescriptorBased is implemented by streams (implementations of
// #GInputStream or #GOutputStream) that are based on file descriptors.
//
// Note that `<gio/gfiledescriptorbased.h>` belongs to the UNIX-specific
// GIO interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
/*

C record/class : GFileDescriptorBased
*/
type FileDescriptorBased struct {
	native *C.GFileDescriptorBased
}

func FileDescriptorBasedNewFromC(u unsafe.Pointer) *FileDescriptorBased {
	c := (*C.GFileDescriptorBased)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBased{native: c}

	return g
}

func (recv *FileDescriptorBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GIcon is a very minimal interface for icons. It provides functions
// for checking the equality of two icons, hashing of icons and
// serializing an icon to and from strings.
//
// #GIcon does not provide the actual pixmap for the icon as this is out
// of GIO's scope, however implementations of #GIcon may contain the name
// of an icon (see #GThemedIcon), or the path to an icon (see #GLoadableIcon).
//
// To obtain a hash of a #GIcon, see g_icon_hash().
//
// To check if two #GIcons are equal, see g_icon_equal().
//
// For serializing a #GIcon, use g_icon_serialize() and
// g_icon_deserialize().
//
// If you want to consume #GIcon (for example, in a toolkit) you must
// be prepared to handle at least the three following cases:
// #GLoadableIcon, #GThemedIcon and #GEmblemedIcon.  It may also make
// sense to have fast-paths for other cases (like handling #GdkPixbuf
// directly, for example) but all compliant #GIcon implementations
// outside of GIO must implement #GLoadableIcon.
//
// If your application or library provides one or more #GIcon
// implementations you need to ensure that your new implementation also
// implements #GLoadableIcon.  Additionally, you must provide an
// implementation of g_icon_serialize() that gives a result that is
// understood by g_icon_deserialize(), yielding one of the built-in icon
// types.
/*

C record/class : GIcon
*/
type Icon struct {
	native *C.GIcon
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	c := (*C.GIcon)(u)
	if c == nil {
		return nil
	}

	g := &Icon{native: c}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Checks if two icons are equal.
/*

C function : g_icon_equal
*/
func (recv *Icon) Equal(icon2 *Icon) bool {
	c_icon2 := (*C.GIcon)(icon2.ToC())

	retC := C.g_icon_equal((*C.GIcon)(recv.native), c_icon2)
	retGo := retC == C.TRUE

	return retGo
}

// #GListModel is an interface that represents a mutable list of
// #GObjects. Its main intention is as a model for various widgets in
// user interfaces, such as list views, but it can also be used as a
// convenient method of returning lists of data, with support for
// updates.
//
// Each object in the list may also report changes in itself via some
// mechanism (normally the #GObject::notify signal).  Taken together
// with the #GListModel::items-changed signal, this provides for a list
// that can change its membership, and in which the members can change
// their individual properties.
//
// A good example would be the list of visible wireless network access
// points, where each access point can report dynamic properties such as
// signal strength.
//
// It is important to note that the #GListModel itself does not report
// changes to the individual items.  It only reports changes to the list
// membership.  If you want to observe changes to the objects themselves
// then you need to connect signals to the objects that you are
// interested in.
//
// All items in a #GListModel are of (or derived from) the same type.
// g_list_model_get_item_type() returns that type.  The type may be an
// interface, in which case all objects in the list must implement it.
//
// The semantics are close to that of an array:
// g_list_model_get_n_items() returns the number of items in the list and
// g_list_model_get_item() returns an item at a (0-based) position. In
// order to allow implementations to calculate the list length lazily,
// you can also iterate over items: starting from 0, repeatedly call
// g_list_model_get_item() until it returns %NULL.
//
// An implementation may create objects lazily, but must take care to
// return the same object for a given position until all references to
// it are gone.
//
// On the other side, a consumer is expected only to hold references on
// objects that are currently "user visible", in order to faciliate the
// maximum level of laziness in the implementation of the list and to
// reduce the required number of signal connections at a given time.
//
// This interface is intended only to be used from a single thread.  The
// thread in which it is appropriate to use it depends on the particular
// implementation, but typically it will be from the thread that owns
// the [thread-default main context][g-main-context-push-thread-default]
// in effect at the time that the model was created.
/*

C record/class : GListModel
*/
type ListModel struct {
	native *C.GListModel
}

func ListModelNewFromC(u unsafe.Pointer) *ListModel {
	c := (*C.GListModel)(u)
	if c == nil {
		return nil
	}

	g := &ListModel{native: c}

	return g
}

func (recv *ListModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Extends the #GIcon interface and adds the ability to
// load icons from streams.
/*

C record/class : GLoadableIcon
*/
type LoadableIcon struct {
	native *C.GLoadableIcon
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	c := (*C.GLoadableIcon)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIcon{native: c}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
/*

C function : g_loadable_icon_load
*/
func (recv *LoadableIcon) Load(size int32, cancellable *Cancellable) (*InputStream, string, error) {
	c_size := (C.int)(size)

	var c_type *C.char

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load((*C.GLoadableIcon)(recv.native), c_size, &c_type, c_cancellable, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goThrowableError
}

// Unsupported : g_loadable_icon_load_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous icon load started in g_loadable_icon_load_async().
/*

C function : g_loadable_icon_load_finish
*/
func (recv *LoadableIcon) LoadFinish(res *AsyncResult) (*InputStream, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_type *C.char

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load_finish((*C.GLoadableIcon)(recv.native), c_res, &c_type, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goThrowableError
}

// The #GMount interface represents user-visible mounts. Note, when
// porting from GnomeVFS, #GMount is the moral equivalent of #GnomeVFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in
// quotes because it's not the same as a unix mount, it might be a gvfs
// mount, but you can still access the files on it if you use GIO. Might or
// might not be related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For
// more information about asynchronous operations, see #GAsyncResult
// and #GTask. To unmount a #GMount instance, first call
// g_mount_unmount_with_operation() with (at least) the #GMount instance and a
// #GAsyncReadyCallback.  The callback will be fired when the
// operation has resolved (either with success or failure), and a
// #GAsyncResult structure will be passed to the callback.  That
// callback should then call g_mount_unmount_with_operation_finish() with the #GMount
// and the #GAsyncResult data to see if the operation was completed
// successfully.  If an @error is present when g_mount_unmount_with_operation_finish()
// is called, then it will be filled with any error information.
/*

C record/class : GMount
*/
type Mount struct {
	native *C.GMount
}

func MountNewFromC(u unsafe.Pointer) *Mount {
	c := (*C.GMount)(u)
	if c == nil {
		return nil
	}

	g := &Mount{native: c}

	return g
}

func (recv *Mount) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalMountChangedDetail struct {
	callback  MountSignalChangedCallback
	handlerID C.gulong
}

var signalMountChangedId int
var signalMountChangedMap = make(map[int]signalMountChangedDetail)
var signalMountChangedLock sync.Mutex

// MountSignalChangedCallback is a callback function for a 'changed' signal emitted from a Mount.
type MountSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Mount) ConnectChanged(callback MountSignalChangedCallback) int {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	signalMountChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_changed(instance, C.gpointer(uintptr(signalMountChangedId)))

	detail := signalMountChangedDetail{callback, handlerID}
	signalMountChangedMap[signalMountChangedId] = detail

	return signalMountChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Mount.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Mount) DisconnectChanged(connectionID int) {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	detail, exists := signalMountChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountChangedMap, connectionID)
}

//export mount_changedHandler
func mount_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalMountChangedMap[index].callback
	callback()
}

type signalMountUnmountedDetail struct {
	callback  MountSignalUnmountedCallback
	handlerID C.gulong
}

var signalMountUnmountedId int
var signalMountUnmountedMap = make(map[int]signalMountUnmountedDetail)
var signalMountUnmountedLock sync.Mutex

// MountSignalUnmountedCallback is a callback function for a 'unmounted' signal emitted from a Mount.
type MountSignalUnmountedCallback func()

/*
ConnectUnmounted connects the callback to the 'unmounted' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectUnmounted to remove it.
*/
func (recv *Mount) ConnectUnmounted(callback MountSignalUnmountedCallback) int {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	signalMountUnmountedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_unmounted(instance, C.gpointer(uintptr(signalMountUnmountedId)))

	detail := signalMountUnmountedDetail{callback, handlerID}
	signalMountUnmountedMap[signalMountUnmountedId] = detail

	return signalMountUnmountedId
}

/*
DisconnectUnmounted disconnects a callback from the 'unmounted' signal for the Mount.

The connectionID should be a value returned from a call to ConnectUnmounted.
*/
func (recv *Mount) DisconnectUnmounted(connectionID int) {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	detail, exists := signalMountUnmountedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountUnmountedMap, connectionID)
}

//export mount_unmountedHandler
func mount_unmountedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalMountUnmountedMap[index].callback
	callback()
}

// Checks if @mount can be ejected.
/*

C function : g_mount_can_eject
*/
func (recv *Mount) CanEject() bool {
	retC := C.g_mount_can_eject((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @mount can be unmounted.
/*

C function : g_mount_can_unmount
*/
func (recv *Mount) CanUnmount() bool {
	retC := C.g_mount_can_unmount((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_mount_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_mount_eject_finish
*/
func (recv *Mount) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the default location of @mount. The default location of the given
// @mount is a path that reflects the main entry point for the user (e.g.
// the home directory, or the root of the volume).
/*

C function : g_mount_get_default_location
*/
func (recv *Mount) GetDefaultLocation() *File {
	retC := C.g_mount_get_default_location((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the drive for the @mount.
//
// This is a convenience method for getting the #GVolume and then
// using that object to get the #GDrive.
/*

C function : g_mount_get_drive
*/
func (recv *Mount) GetDrive() *Drive {
	retC := C.g_mount_get_drive((*C.GMount)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the icon for @mount.
/*

C function : g_mount_get_icon
*/
func (recv *Mount) GetIcon() *Icon {
	retC := C.g_mount_get_icon((*C.GMount)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of @mount.
/*

C function : g_mount_get_name
*/
func (recv *Mount) GetName() string {
	retC := C.g_mount_get_name((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the root directory on @mount.
/*

C function : g_mount_get_root
*/
func (recv *Mount) GetRoot() *File {
	retC := C.g_mount_get_root((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the UUID for the @mount. The reference is typically based on
// the file system UUID for the mount in question and should be
// considered an opaque string. Returns %NULL if there is no UUID
// available.
/*

C function : g_mount_get_uuid
*/
func (recv *Mount) GetUuid() string {
	retC := C.g_mount_get_uuid((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the volume for the @mount.
/*

C function : g_mount_get_volume
*/
func (recv *Mount) GetVolume() *Volume {
	retC := C.g_mount_get_volume((*C.GMount)(recv.native))
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_mount_remount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes remounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_mount_remount_finish
*/
func (recv *Mount) RemountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_remount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_mount_unmount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes unmounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_mount_unmount_finish
*/
func (recv *Mount) UnmountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// The GRemoteActionGroup interface is implemented by #GActionGroup
// instances that either transmit action invocations to other processes
// or receive action invocations in the local process from other
// processes.
//
// The interface has `_full` variants of the two
// methods on #GActionGroup used to activate actions:
// g_action_group_activate_action() and
// g_action_group_change_action_state(). These variants allow a
// "platform data" #GVariant to be specified: a dictionary providing
// context for the action invocation (for example: timestamps, startup
// notification IDs, etc).
//
// #GDBusActionGroup implements #GRemoteActionGroup.  This provides a
// mechanism to send platform data for action invocations over D-Bus.
//
// Additionally, g_dbus_connection_export_action_group() will check if
// the exported #GActionGroup implements #GRemoteActionGroup and use the
// `_full` variants of the calls if available.  This
// provides a mechanism by which to receive platform data for action
// invocations that arrive by way of D-Bus.
/*

C record/class : GRemoteActionGroup
*/
type RemoteActionGroup struct {
	native *C.GRemoteActionGroup
}

func RemoteActionGroupNewFromC(u unsafe.Pointer) *RemoteActionGroup {
	c := (*C.GRemoteActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroup{native: c}

	return g
}

func (recv *RemoteActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GSeekable is implemented by streams (implementations of
// #GInputStream or #GOutputStream) that support seeking.
//
// Seekable streams largely fall into two categories: resizable and
// fixed-size.
//
// #GSeekable on fixed-sized streams is approximately the same as POSIX
// lseek() on a block device (for example: attmepting to seek past the
// end of the device is an error).  Fixed streams typically cannot be
// truncated.
//
// #GSeekable on resizable streams is approximately the same as POSIX
// lseek() on a normal file.  Seeking past the end and writing data will
// usually cause the stream to resize by introducing zero bytes.
/*

C record/class : GSeekable
*/
type Seekable struct {
	native *C.GSeekable
}

func SeekableNewFromC(u unsafe.Pointer) *Seekable {
	c := (*C.GSeekable)(u)
	if c == nil {
		return nil
	}

	g := &Seekable{native: c}

	return g
}

func (recv *Seekable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Tests if the stream supports the #GSeekableIface.
/*

C function : g_seekable_can_seek
*/
func (recv *Seekable) CanSeek() bool {
	retC := C.g_seekable_can_seek((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
/*

C function : g_seekable_can_truncate
*/
func (recv *Seekable) CanTruncate() bool {
	retC := C.g_seekable_can_truncate((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Seeks in the stream by the given @offset, modified by @type.
//
// Attempting to seek past the end of the stream will have different
// results depending on if the stream is fixed-sized or resizable.  If
// the stream is resizable then seeking past the end and then writing
// will result in zeros filling the empty space.  Seeking past the end
// of a resizable stream and reading will result in EOF.  Seeking past
// the end of a fixed-sized stream will fail.
//
// Any operation that would result in a negative offset will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_seekable_seek
*/
func (recv *Seekable) Seek(offset uint64, type_ glib.SeekType, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_type := (C.GSeekType)(type_)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_seekable_seek((*C.GSeekable)(recv.native), c_offset, c_type, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tells the current position within the stream.
/*

C function : g_seekable_tell
*/
func (recv *Seekable) Tell() uint64 {
	retC := C.g_seekable_tell((*C.GSeekable)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Sets the length of the stream to @offset. If the stream was previously
// larger than @offset, the extra data is discarded. If the stream was
// previouly shorter than @offset, it is extended with NUL ('\0') bytes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
/*

C function : g_seekable_truncate
*/
func (recv *Seekable) Truncate(offset uint64, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_seekable_truncate((*C.GSeekable)(recv.native), c_offset, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Objects that describe one or more potential socket endpoints
// implement #GSocketConnectable. Callers can then use
// g_socket_connectable_enumerate() to get a #GSocketAddressEnumerator
// to try out each socket address in turn until one succeeds, as shown
// in the sample code below.
//
// |[<!-- language="C" -->
// MyConnectionType *
// connect_to_host (const char    *hostname,
// guint16        port,
// GCancellable  *cancellable,
// GError       **error)
// {
// MyConnection *conn = NULL;
// GSocketConnectable *addr;
// GSocketAddressEnumerator *enumerator;
// GSocketAddress *sockaddr;
// GError *conn_error = NULL;
//
// addr = g_network_address_new (hostname, port);
// enumerator = g_socket_connectable_enumerate (addr);
// g_object_unref (addr);
//
// Try each sockaddr until we succeed. Record the first connection error,
// but not any further ones (since they'll probably be basically the same
// as the first).
// while (!conn && (sockaddr = g_socket_address_enumerator_next (enumerator, cancellable, error))
// {
// conn = connect_to_sockaddr (sockaddr, conn_error ? NULL : &conn_error);
// g_object_unref (sockaddr);
// }
// g_object_unref (enumerator);
//
// if (conn)
// {
// if (conn_error)
// {
// We couldn't connect to the first address, but we succeeded
// in connecting to a later address.
// g_error_free (conn_error);
// }
// return conn;
// }
// else if (error)
// {
/// Either initial lookup failed, or else the caller cancelled us.
// if (conn_error)
// g_error_free (conn_error);
// return NULL;
// }
// else
// {
// g_error_propagate (error, conn_error);
// return NULL;
// }
// }
// ]|
/*

C record/class : GSocketConnectable
*/
type SocketConnectable struct {
	native *C.GSocketConnectable
}

func SocketConnectableNewFromC(u unsafe.Pointer) *SocketConnectable {
	c := (*C.GSocketConnectable)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectable{native: c}

	return g
}

func (recv *SocketConnectable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GVolume interface represents user-visible objects that can be
// mounted. Note, when porting from GnomeVFS, #GVolume is the moral
// equivalent of #GnomeVFSDrive.
//
// Mounting a #GVolume instance is an asynchronous operation. For more
// information about asynchronous operations, see #GAsyncResult and
// #GTask. To mount a #GVolume, first call g_volume_mount() with (at
// least) the #GVolume instance, optionally a #GMountOperation object
// and a #GAsyncReadyCallback.
//
// Typically, one will only want to pass %NULL for the
// #GMountOperation if automounting all volumes when a desktop session
// starts since it's not desirable to put up a lot of dialogs asking
// for credentials.
//
// The callback will be fired when the operation has resolved (either
// with success or failure), and a #GAsyncReady structure will be
// passed to the callback.  That callback should then call
// g_volume_mount_finish() with the #GVolume instance and the
// #GAsyncReady data to see if the operation was completed
// successfully.  If an @error is present when g_volume_mount_finish()
// is called, then it will be filled with any error information.
//
// ## Volume Identifiers # {#volume-identifier}
//
// It is sometimes necessary to directly access the underlying
// operating system object behind a volume (e.g. for passing a volume
// to an application via the commandline). For this purpose, GIO
// allows to obtain an 'identifier' for the volume. There can be
// different kinds of identifiers, such as Hal UDIs, filesystem labels,
// traditional Unix devices (e.g. `/dev/sda2`), UUIDs. GIO uses predefined
// strings as names for the different kinds of identifiers:
// #G_VOLUME_IDENTIFIER_KIND_HAL_UDI, #G_VOLUME_IDENTIFIER_KIND_LABEL, etc.
// Use g_volume_get_identifier() to obtain an identifier for a volume.
//
//
// Note that #G_VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available
// when the gvfs hal volume monitor is in use. Other volume monitors
// will generally be able to provide the #G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE
// identifier, which can be used to obtain a hal device by means of
// libhal_manager_find_device_string_match().
/*

C record/class : GVolume
*/
type Volume struct {
	native *C.GVolume
}

func VolumeNewFromC(u unsafe.Pointer) *Volume {
	c := (*C.GVolume)(u)
	if c == nil {
		return nil
	}

	g := &Volume{native: c}

	return g
}

func (recv *Volume) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

type signalVolumeChangedDetail struct {
	callback  VolumeSignalChangedCallback
	handlerID C.gulong
}

var signalVolumeChangedId int
var signalVolumeChangedMap = make(map[int]signalVolumeChangedDetail)
var signalVolumeChangedLock sync.Mutex

// VolumeSignalChangedCallback is a callback function for a 'changed' signal emitted from a Volume.
type VolumeSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Volume) ConnectChanged(callback VolumeSignalChangedCallback) int {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	signalVolumeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_changed(instance, C.gpointer(uintptr(signalVolumeChangedId)))

	detail := signalVolumeChangedDetail{callback, handlerID}
	signalVolumeChangedMap[signalVolumeChangedId] = detail

	return signalVolumeChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Volume) DisconnectChanged(connectionID int) {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	detail, exists := signalVolumeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeChangedMap, connectionID)
}

//export volume_changedHandler
func volume_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalVolumeChangedMap[index].callback
	callback()
}

type signalVolumeRemovedDetail struct {
	callback  VolumeSignalRemovedCallback
	handlerID C.gulong
}

var signalVolumeRemovedId int
var signalVolumeRemovedMap = make(map[int]signalVolumeRemovedDetail)
var signalVolumeRemovedLock sync.Mutex

// VolumeSignalRemovedCallback is a callback function for a 'removed' signal emitted from a Volume.
type VolumeSignalRemovedCallback func()

/*
ConnectRemoved connects the callback to the 'removed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectRemoved to remove it.
*/
func (recv *Volume) ConnectRemoved(callback VolumeSignalRemovedCallback) int {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	signalVolumeRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_removed(instance, C.gpointer(uintptr(signalVolumeRemovedId)))

	detail := signalVolumeRemovedDetail{callback, handlerID}
	signalVolumeRemovedMap[signalVolumeRemovedId] = detail

	return signalVolumeRemovedId
}

/*
DisconnectRemoved disconnects a callback from the 'removed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectRemoved.
*/
func (recv *Volume) DisconnectRemoved(connectionID int) {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	detail, exists := signalVolumeRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeRemovedMap, connectionID)
}

//export volume_removedHandler
func volume_removedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalVolumeRemovedMap[index].callback
	callback()
}

// Checks if a volume can be ejected.
/*

C function : g_volume_can_eject
*/
func (recv *Volume) CanEject() bool {
	retC := C.g_volume_can_eject((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a volume can be mounted.
/*

C function : g_volume_can_mount
*/
func (recv *Volume) CanMount() bool {
	retC := C.g_volume_can_mount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_volume_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes ejecting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
/*

C function : g_volume_eject_finish
*/
func (recv *Volume) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_volume_enumerate_identifiers : no return type

// Gets the drive for the @volume.
/*

C function : g_volume_get_drive
*/
func (recv *Volume) GetDrive() *Drive {
	retC := C.g_volume_get_drive((*C.GVolume)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the icon for @volume.
/*

C function : g_volume_get_icon
*/
func (recv *Volume) GetIcon() *Icon {
	retC := C.g_volume_get_icon((*C.GVolume)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the identifier of the given kind for @volume.
// See the [introduction][volume-identifier] for more
// information about volume identifiers.
/*

C function : g_volume_get_identifier
*/
func (recv *Volume) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_volume_get_identifier((*C.GVolume)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the mount for the @volume.
/*

C function : g_volume_get_mount
*/
func (recv *Volume) GetMount() *Mount {
	retC := C.g_volume_get_mount((*C.GVolume)(recv.native))
	retGo := MountNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the name of @volume.
/*

C function : g_volume_get_name
*/
func (recv *Volume) GetName() string {
	retC := C.g_volume_get_name((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the UUID for the @volume. The reference is typically based on
// the file system UUID for the volume in question and should be
// considered an opaque string. Returns %NULL if there is no UUID
// available.
/*

C function : g_volume_get_uuid
*/
func (recv *Volume) GetUuid() string {
	retC := C.g_volume_get_uuid((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_mount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes mounting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
//
// If the mount operation succeeded, g_volume_get_mount() on @volume
// is guaranteed to return the mount right after calling this
// function; there's no need to listen for the 'mount-added' signal on
// #GVolumeMonitor.
/*

C function : g_volume_mount_finish
*/
func (recv *Volume) MountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_mount_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns whether the volume should be automatically mounted.
/*

C function : g_volume_should_automount
*/
func (recv *Volume) ShouldAutomount() bool {
	retC := C.g_volume_should_automount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
