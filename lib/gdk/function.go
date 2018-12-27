// This is a generated file - DO NOT EDIT

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// AddOptionEntriesLibgtkOnly is a wrapper around the C function gdk_add_option_entries_libgtk_only.
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.gdk_add_option_entries_libgtk_only(c_group)

	return
}

// Unsupported : gdk_atom_intern : return type : Blacklisted record : GdkAtom

// Beep is a wrapper around the C function gdk_beep.
func Beep() {
	C.gdk_beep()

	return
}

// Blacklisted : gdk_cairo_get_clip_rectangle

// CairoRegionCreateFromSurface is a wrapper around the C function gdk_cairo_region_create_from_surface.
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gdk_cairo_region_create_from_surface(c_surface)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_color_parse

// DragAbort is a wrapper around the C function gdk_drag_abort.
func DragAbort(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_abort(c_context, c_time_)

	return
}

// DragBegin is a wrapper around the C function gdk_drag_begin.
func DragBegin(window *Window, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin(c_window, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragBeginForDevice is a wrapper around the C function gdk_drag_begin_for_device.
func DragBeginForDevice(window *Window, device *Device, targets *glib.List) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.gdk_drag_begin_for_device(c_window, c_device, c_targets)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragDrop is a wrapper around the C function gdk_drag_drop.
func DragDrop(context *DragContext, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_drag_drop(c_context, c_time_)

	return
}

// Unsupported : gdk_drag_get_selection : return type : Blacklisted record : GdkAtom

// DragMotion is a wrapper around the C function gdk_drag_motion.
func DragMotion(context *DragContext, destWindow *Window, protocol DragProtocol, xRoot int32, yRoot int32, suggestedAction DragAction, possibleActions DragAction, time uint32) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_dest_window := (*C.GdkWindow)(C.NULL)
	if destWindow != nil {
		c_dest_window = (*C.GdkWindow)(destWindow.ToC())
	}

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	c_suggested_action := (C.GdkDragAction)(suggestedAction)

	c_possible_actions := (C.GdkDragAction)(possibleActions)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_drag_motion(c_context, c_dest_window, c_protocol, c_x_root, c_y_root, c_suggested_action, c_possible_actions, c_time_)
	retGo := retC == C.TRUE

	return retGo
}

// DragStatus is a wrapper around the C function gdk_drag_status.
func DragStatus(context *DragContext, action DragAction, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_action := (C.GdkDragAction)(action)

	c_time_ := (C.guint32)(time)

	C.gdk_drag_status(c_context, c_action, c_time_)

	return
}

// DropFinish is a wrapper around the C function gdk_drop_finish.
func DropFinish(context *DragContext, success bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_finish(c_context, c_success, c_time_)

	return
}

// DropReply is a wrapper around the C function gdk_drop_reply.
func DropReply(context *DragContext, accepted bool, time uint32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_accepted :=
		boolToGboolean(accepted)

	c_time_ := (C.guint32)(time)

	C.gdk_drop_reply(c_context, c_accepted, c_time_)

	return
}

// ErrorTrapPop is a wrapper around the C function gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	retC := C.gdk_error_trap_pop()
	retGo := (int32)(retC)

	return retGo
}

// ErrorTrapPush is a wrapper around the C function gdk_error_trap_push.
func ErrorTrapPush() {
	C.gdk_error_trap_push()

	return
}

// Unsupported : gdk_event_get : no return generator

// Unsupported : gdk_event_handler_set : unsupported parameter func : no type generator for EventFunc (GdkEventFunc) for param func

// Unsupported : gdk_event_peek : no return generator

// EventsPending is a wrapper around the C function gdk_events_pending.
func EventsPending() bool {
	retC := C.gdk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// Flush is a wrapper around the C function gdk_flush.
func Flush() {
	C.gdk_flush()

	return
}

// GetDefaultRootWindow is a wrapper around the C function gdk_get_default_root_window.
func GetDefaultRootWindow() *Window {
	retC := C.gdk_get_default_root_window()
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_get_display.
func GetDisplay() string {
	retC := C.gdk_get_display()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetProgramClass is a wrapper around the C function gdk_get_program_class.
func GetProgramClass() string {
	retC := C.gdk_get_program_class()
	retGo := C.GoString(retC)

	return retGo
}

// GetShowEvents is a wrapper around the C function gdk_get_show_events.
func GetShowEvents() bool {
	retC := C.gdk_get_show_events()
	retGo := retC == C.TRUE

	return retGo
}

// Init is a wrapper around the C function gdk_init.
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// InitCheck is a wrapper around the C function gdk_init_check.
func InitCheck(args []string) (bool, []string) {
	cArgc, cArgv := argsIn(args)

	retC := C.gdk_init_check(&cArgc, &cArgv)
	retGo := retC == C.TRUE

	args = argsOut(cArgc, cArgv)

	return retGo, args
}

// KeyboardGrab is a wrapper around the C function gdk_keyboard_grab.
func KeyboardGrab(window *Window, ownerEvents bool, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_time_ := (C.guint32)(time)

	retC := C.gdk_keyboard_grab(c_window, c_owner_events, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// KeyboardUngrab is a wrapper around the C function gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_keyboard_ungrab(c_time_)

	return
}

// KeyvalConvertCase is a wrapper around the C function gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	c_symbol := (C.guint)(symbol)

	var c_lower C.guint

	var c_upper C.guint

	C.gdk_keyval_convert_case(c_symbol, &c_lower, &c_upper)

	lower := (uint32)(c_lower)

	upper := (uint32)(c_upper)

	return lower, upper
}

// KeyvalFromName is a wrapper around the C function gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	c_keyval_name := C.CString(keyvalName)
	defer C.free(unsafe.Pointer(c_keyval_name))

	retC := C.gdk_keyval_from_name(c_keyval_name)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalIsLower is a wrapper around the C function gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_lower(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalIsUpper is a wrapper around the C function gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint32) bool {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_is_upper(c_keyval)
	retGo := retC == C.TRUE

	return retGo
}

// KeyvalName is a wrapper around the C function gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_name(c_keyval)
	retGo := C.GoString(retC)

	return retGo
}

// KeyvalToLower is a wrapper around the C function gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_lower(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUnicode is a wrapper around the C function gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_unicode(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// KeyvalToUpper is a wrapper around the C function gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	c_keyval := (C.guint)(keyval)

	retC := C.gdk_keyval_to_upper(c_keyval)
	retGo := (uint32)(retC)

	return retGo
}

// ListVisuals is a wrapper around the C function gdk_list_visuals.
func ListVisuals() *glib.List {
	retC := C.gdk_list_visuals()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OffscreenWindowGetSurface is a wrapper around the C function gdk_offscreen_window_get_surface.
func OffscreenWindowGetSurface(window *Window) *cairo.Surface {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_offscreen_window_get_surface(c_window)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PangoContextGet is a wrapper around the C function gdk_pango_context_get.
func PangoContextGet() *pango.Context {
	retC := C.gdk_pango_context_get()
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoLayoutGetClipRegion is a wrapper around the C function gdk_pango_layout_get_clip_region.
func PangoLayoutGetClipRegion(layout *pango.Layout, xOrigin int32, yOrigin int32, indexRanges int32, nRanges int32) *cairo.Region {
	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges := (C.gint)(indexRanges)

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_get_clip_region(c_layout, c_x_origin, c_y_origin, &c_index_ranges, c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PangoLayoutLineGetClipRegion is a wrapper around the C function gdk_pango_layout_line_get_clip_region.
func PangoLayoutLineGetClipRegion(line *pango.LayoutLine, xOrigin int32, yOrigin int32, indexRanges []int32, nRanges int32) *cairo.Region {
	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x_origin := (C.gint)(xOrigin)

	c_y_origin := (C.gint)(yOrigin)

	c_index_ranges := &indexRanges[0]

	c_n_ranges := (C.gint)(nRanges)

	retC := C.gdk_pango_layout_line_get_clip_region(c_line, c_x_origin, c_y_origin, (*C.gint)(unsafe.Pointer(c_index_ranges)), c_n_ranges)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PixbufGetFromSurface is a wrapper around the C function gdk_pixbuf_get_from_surface.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_surface(c_surface, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PixbufGetFromWindow is a wrapper around the C function gdk_pixbuf_get_from_window.
func PixbufGetFromWindow(window *Window, srcX int32, srcY int32, width int32, height int32) *gdkpixbuf.Pixbuf {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	retC := C.gdk_pixbuf_get_from_window(c_window, c_src_x, c_src_y, c_width, c_height)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PointerGrab is a wrapper around the C function gdk_pointer_grab.
func PointerGrab(window *Window, ownerEvents bool, eventMask EventMask, confineTo *Window, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_confine_to := (*C.GdkWindow)(C.NULL)
	if confineTo != nil {
		c_confine_to = (*C.GdkWindow)(confineTo.ToC())
	}

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gdk_pointer_grab(c_window, c_owner_events, c_event_mask, c_confine_to, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// PointerIsGrabbed is a wrapper around the C function gdk_pointer_is_grabbed.
func PointerIsGrabbed() bool {
	retC := C.gdk_pointer_is_grabbed()
	retGo := retC == C.TRUE

	return retGo
}

// PointerUngrab is a wrapper around the C function gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_pointer_ungrab(c_time_)

	return
}

// PreParseLibgtkOnly is a wrapper around the C function gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()

	return
}

// Unsupported : gdk_property_change : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_delete : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_property_get : unsupported parameter property : Blacklisted record : GdkAtom

// Unsupported : gdk_query_depths : unsupported parameter depths : output array param depths

// Unsupported : gdk_query_visual_types : unsupported parameter visual_types : output array param visual_types

// Unsupported : gdk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_get : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_selection_property_get : unsupported parameter data : guchar** with indirection level of 2

// Unsupported : gdk_selection_send_notify : unsupported parameter selection : Blacklisted record : GdkAtom

// SetDoubleClickTime is a wrapper around the C function gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_set_double_click_time(c_msec)

	return
}

// SetProgramClass is a wrapper around the C function gdk_set_program_class.
func SetProgramClass(programClass string) {
	c_program_class := C.CString(programClass)
	defer C.free(unsafe.Pointer(c_program_class))

	C.gdk_set_program_class(c_program_class)

	return
}

// SetShowEvents is a wrapper around the C function gdk_set_show_events.
func SetShowEvents(showEvents bool) {
	c_show_events :=
		boolToGboolean(showEvents)

	C.gdk_set_show_events(c_show_events)

	return
}

// SettingGet is a wrapper around the C function gdk_setting_get.
func SettingGet(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.gdk_setting_get(c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gdk_synthesize_window_state

// ThreadsEnter is a wrapper around the C function gdk_threads_enter.
func ThreadsEnter() {
	C.gdk_threads_enter()

	return
}

// ThreadsInit is a wrapper around the C function gdk_threads_init.
func ThreadsInit() {
	C.gdk_threads_init()

	return
}

// ThreadsLeave is a wrapper around the C function gdk_threads_leave.
func ThreadsLeave() {
	C.gdk_threads_leave()

	return
}

// UnicodeToKeyval is a wrapper around the C function gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	c_wc := (C.guint32)(wc)

	retC := C.gdk_unicode_to_keyval(c_wc)
	retGo := (uint32)(retC)

	return retGo
}

// Utf8ToStringTarget is a wrapper around the C function gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gdk_utf8_to_string_target(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
