// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
)

// DragSetIconGicon is a wrapper around the C function gtk_drag_set_icon_gicon.
func DragSetIconGicon(context *gdk.DragContext, icon *gio.Icon, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_icon := (*C.GIcon)(icon.ToC())

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_gicon(c_context, c_icon, c_hot_x, c_hot_y)

	return
}

// RenderIcon is a wrapper around the C function gtk_render_icon.
func RenderIcon(context *StyleContext, cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x float64, y float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	C.gtk_render_icon(c_context, c_cr, c_pixbuf, c_x, c_y)

	return
}
