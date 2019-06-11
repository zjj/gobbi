#include <gtk/gtk.h>
#include "_cgo_export.h"

gboolean drawing_area_vf_draw2(GtkWidget *widget, cairo_t *cr) {
	return DrawingAreaDraw2(widget, cr);
}

void drawing_area_class_init2(GtkDrawingAreaClass *g_class, gpointer class_data) {
	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;

	widget_class->draw = drawing_area_vf_draw2;
}
