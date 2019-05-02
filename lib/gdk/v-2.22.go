// Code generated - DO NOT EDIT.
// +build gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetCursorType is a wrapper around the C function gdk_cursor_get_cursor_type.
func (recv *Cursor) GetCursorType() CursorType {
	retC := C.gdk_cursor_get_cursor_type((*C.GdkCursor)(recv.native))
	retGo := (CursorType)(retC)

	return retGo
}

// IsClosed is a wrapper around the C function gdk_display_is_closed.
func (recv *Display) IsClosed() bool {
	retC := C.gdk_display_is_closed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetActions is a wrapper around the C function gdk_drag_context_get_actions.
func (recv *DragContext) GetActions() DragAction {
	retC := C.gdk_drag_context_get_actions((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// GetSelectedAction is a wrapper around the C function gdk_drag_context_get_selected_action.
func (recv *DragContext) GetSelectedAction() DragAction {
	retC := C.gdk_drag_context_get_selected_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// GetSourceWindow is a wrapper around the C function gdk_drag_context_get_source_window.
func (recv *DragContext) GetSourceWindow() *Window {
	retC := C.gdk_drag_context_get_source_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSuggestedAction is a wrapper around the C function gdk_drag_context_get_suggested_action.
func (recv *DragContext) GetSuggestedAction() DragAction {
	retC := C.gdk_drag_context_get_suggested_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// ListTargets is a wrapper around the C function gdk_drag_context_list_targets.
func (recv *DragContext) ListTargets() *glib.List {
	retC := C.gdk_drag_context_list_targets((*C.GdkDragContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBitsPerRgb is a wrapper around the C function gdk_visual_get_bits_per_rgb.
func (recv *Visual) GetBitsPerRgb() int32 {
	retC := C.gdk_visual_get_bits_per_rgb((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetBluePixelDetails is a wrapper around the C function gdk_visual_get_blue_pixel_details.
func (recv *Visual) GetBluePixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_blue_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// GetByteOrder is a wrapper around the C function gdk_visual_get_byte_order.
func (recv *Visual) GetByteOrder() ByteOrder {
	retC := C.gdk_visual_get_byte_order((*C.GdkVisual)(recv.native))
	retGo := (ByteOrder)(retC)

	return retGo
}

// GetColormapSize is a wrapper around the C function gdk_visual_get_colormap_size.
func (recv *Visual) GetColormapSize() int32 {
	retC := C.gdk_visual_get_colormap_size((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDepth is a wrapper around the C function gdk_visual_get_depth.
func (recv *Visual) GetDepth() int32 {
	retC := C.gdk_visual_get_depth((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetGreenPixelDetails is a wrapper around the C function gdk_visual_get_green_pixel_details.
func (recv *Visual) GetGreenPixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_green_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// GetRedPixelDetails is a wrapper around the C function gdk_visual_get_red_pixel_details.
func (recv *Visual) GetRedPixelDetails() (uint32, int32, int32) {
	var c_mask C.guint32

	var c_shift C.gint

	var c_precision C.gint

	C.gdk_visual_get_red_pixel_details((*C.GdkVisual)(recv.native), &c_mask, &c_shift, &c_precision)

	mask := (uint32)(c_mask)

	shift := (int32)(c_shift)

	precision := (int32)(c_precision)

	return mask, shift, precision
}

// GetVisualType is a wrapper around the C function gdk_visual_get_visual_type.
func (recv *Visual) GetVisualType() VisualType {
	retC := C.gdk_visual_get_visual_type((*C.GdkVisual)(recv.native))
	retGo := (VisualType)(retC)

	return retGo
}

// CoordsFromParent is a wrapper around the C function gdk_window_coords_from_parent.
func (recv *Window) CoordsFromParent(parentX float64, parentY float64) (float64, float64) {
	c_parent_x := (C.gdouble)(parentX)

	c_parent_y := (C.gdouble)(parentY)

	var c_x C.gdouble

	var c_y C.gdouble

	C.gdk_window_coords_from_parent((*C.GdkWindow)(recv.native), c_parent_x, c_parent_y, &c_x, &c_y)

	x := (float64)(c_x)

	y := (float64)(c_y)

	return x, y
}

// CoordsToParent is a wrapper around the C function gdk_window_coords_to_parent.
func (recv *Window) CoordsToParent(x float64, y float64) (float64, float64) {
	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	var c_parent_x C.gdouble

	var c_parent_y C.gdouble

	C.gdk_window_coords_to_parent((*C.GdkWindow)(recv.native), c_x, c_y, &c_parent_x, &c_parent_y)

	parentX := (float64)(c_parent_x)

	parentY := (float64)(c_parent_y)

	return parentX, parentY
}

// CreateSimilarSurface is a wrapper around the C function gdk_window_create_similar_surface.
func (recv *Window) CreateSimilarSurface(content cairo.Content, width int32, height int32) *cairo.Surface {
	c_content := (C.cairo_content_t)(content)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_window_create_similar_surface((*C.GdkWindow)(recv.native), c_content, c_width, c_height)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAcceptFocus is a wrapper around the C function gdk_window_get_accept_focus.
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gdk_window_get_accept_focus((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBackgroundPattern is a wrapper around the C function gdk_window_get_background_pattern.
func (recv *Window) GetBackgroundPattern() *cairo.Pattern {
	retC := C.gdk_window_get_background_pattern((*C.GdkWindow)(recv.native))
	var retGo (*cairo.Pattern)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.PatternNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetComposited is a wrapper around the C function gdk_window_get_composited.
func (recv *Window) GetComposited() bool {
	retC := C.gdk_window_get_composited((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEffectiveParent is a wrapper around the C function gdk_window_get_effective_parent.
func (recv *Window) GetEffectiveParent() *Window {
	retC := C.gdk_window_get_effective_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEffectiveToplevel is a wrapper around the C function gdk_window_get_effective_toplevel.
func (recv *Window) GetEffectiveToplevel() *Window {
	retC := C.gdk_window_get_effective_toplevel((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusOnMap is a wrapper around the C function gdk_window_get_focus_on_map.
func (recv *Window) GetFocusOnMap() bool {
	retC := C.gdk_window_get_focus_on_map((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModalHint is a wrapper around the C function gdk_window_get_modal_hint.
func (recv *Window) GetModalHint() bool {
	retC := C.gdk_window_get_modal_hint((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasNative is a wrapper around the C function gdk_window_has_native.
func (recv *Window) HasNative() bool {
	retC := C.gdk_window_has_native((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsInputOnly is a wrapper around the C function gdk_window_is_input_only.
func (recv *Window) IsInputOnly() bool {
	retC := C.gdk_window_is_input_only((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsShaped is a wrapper around the C function gdk_window_is_shaped.
func (recv *Window) IsShaped() bool {
	retC := C.gdk_window_is_shaped((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
