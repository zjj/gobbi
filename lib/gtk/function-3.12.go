// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "C"

// GetLocaleDirection is a wrapper around the C function gtk_get_locale_direction.
func GetLocaleDirection() TextDirection {
	retC := C.gtk_get_locale_direction()
	retGo := (TextDirection)(retC)

	return retGo
}
