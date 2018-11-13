// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Gets whether hidden files and folders are displayed in the file selector.
// See gtk_file_chooser_set_show_hidden().
/*

C function

gtk_file_chooser_get_show_hidden
*/
func (recv *FileChooser) GetShowHidden() bool {
	retC := C.gtk_file_chooser_get_show_hidden((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether hidden files and folders are displayed in the file selector.
/*

C function

gtk_file_chooser_set_show_hidden
*/
func (recv *FileChooser) SetShowHidden(showHidden bool) {
	c_show_hidden :=
		boolToGboolean(showHidden)

	C.gtk_file_chooser_set_show_hidden((*C.GtkFileChooser)(recv.native), c_show_hidden)

	return
}
