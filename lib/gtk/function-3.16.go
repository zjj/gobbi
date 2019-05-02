// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdk "github.com/pekim/gobbi/lib/gdk"
)

// DragCancel is a wrapper around the C function gtk_drag_cancel.
func DragCancel(context *gdk.DragContext) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	C.gtk_drag_cancel(c_context)

	return
}
