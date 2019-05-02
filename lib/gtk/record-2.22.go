// This is a generated file - DO NOT EDIT
// +build gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gio "github.com/pekim/gobbi/lib/gio"
	"unsafe"
)

// GetGicon is a wrapper around the C function gtk_recent_info_get_gicon.
func (recv *RecentInfo) GetGicon() *gio.Icon {
	retC := C.gtk_recent_info_get_gicon((*C.GtkRecentInfo)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
