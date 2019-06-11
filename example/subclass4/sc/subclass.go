package sc

/*
#cgo pkg-config: gtk+-3.0

#include <gtk/gtk.h>
#include <stdlib.h>

void drawing_area_class_init2(GtkDrawingAreaClass *g_class, gpointer class_data);
*/
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"math"
	"unsafe"
)

type DrawingAreaVirtualDraw2 interface {
	Draw(cr *cairo.Context) bool
}

type DrawingAreaDerivedClass2 struct {
	name  string
	gtype C.GType
}

type DrawingAreaDerived2 struct {
	class  *DrawingAreaDerivedClass2
	native *C.GtkDrawingArea
}

func DrawingAreaDerive2(name string) *DrawingAreaDerivedClass2 {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkDrawingAreaClass
	typeInfo.instance_size = C.sizeof_GtkDrawingArea
	typeInfo.class_init = C.GClassInitFunc(C.drawing_area_class_init2)

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_DRAWING_AREA, cTypeName, &typeInfo, 0)

	class := &DrawingAreaDerivedClass2{
		name:  name,
		gtype: gtype,
	}

	return class
}

func (c *DrawingAreaDerivedClass2) New(virtualFunctions interface{}) *DrawingAreaDerived2 {
	f, ok := virtualFunctions.(DrawingAreaVirtualDraw2)
	fmt.Println("draw func :", ok)
	f.Draw(nil)

	native := (*C.GtkDrawingArea)(C.g_object_newv(c.gtype, 0, nil))

	instance := &DrawingAreaDerived2{
		class:  c,
		native: native,
	}

	return instance
}

// DrawingArea upcasts to *DrawingArea
func (recv *DrawingAreaDerived2) DrawingArea() *gtk.DrawingArea {
	return gtk.DrawingAreaNewFromC(unsafe.Pointer(recv.native))

}

//export DrawingAreaDraw2
func DrawingAreaDraw2(daC *C.GtkDrawingArea, contextC *C.cairo_t) C.gboolean {
	da := gtk.DrawingAreaNewFromC(unsafe.Pointer(daC))

	widget := da.Widget()
	cr := cairo.ContextNewFromC(unsafe.Pointer(contextC))

	// find the dimensions that the widget's  been allocated, and
	// therefore the size of the area to draw in
	alloc := widget.GetAllocation()
	height := float64(alloc.Height)
	width := float64(alloc.Width)

	// render background first
	gtk.RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	// an arc that describes a circle to the path
	cr.Arc(width/2.0, height/2.0, math.Min(width, height)/2.0, 0, 2*math.Pi)

	styleContext := widget.GetStyleContext()
	colour := styleContext.GetColor(styleContext.GetState())

	// use the widget's context's colour for the source pattern
	gdk.CairoSetSourceRgba(cr, colour)

	// fill the path (that describes a circle)
	cr.Fill()

	return C.FALSE
}
