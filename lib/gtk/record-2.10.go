// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// PaperSizeGetDefault is a wrapper around the C function gtk_paper_size_get_default.
func PaperSizeGetDefault() string {
	retC := C.gtk_paper_size_get_default()
	retGo := C.GoString(retC)

	return retGo
}

// Copy is a wrapper around the C function gtk_paper_size_copy.
func (recv *PaperSize) Copy() *PaperSize {
	retC := C.gtk_paper_size_copy((*C.GtkPaperSize)(recv.native))
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_paper_size_free.
func (recv *PaperSize) Free() {
	C.gtk_paper_size_free((*C.GtkPaperSize)(recv.native))

	return
}

// GetDefaultBottomMargin is a wrapper around the C function gtk_paper_size_get_default_bottom_margin.
func (recv *PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_bottom_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultLeftMargin is a wrapper around the C function gtk_paper_size_get_default_left_margin.
func (recv *PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_left_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultRightMargin is a wrapper around the C function gtk_paper_size_get_default_right_margin.
func (recv *PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_right_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDefaultTopMargin is a wrapper around the C function gtk_paper_size_get_default_top_margin.
func (recv *PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_top_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_paper_size_get_display_name.
func (recv *PaperSize) GetDisplayName() string {
	retC := C.gtk_paper_size_get_display_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHeight is a wrapper around the C function gtk_paper_size_get_height.
func (recv *PaperSize) GetHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_height((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_paper_size_get_name.
func (recv *PaperSize) GetName() string {
	retC := C.gtk_paper_size_get_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPpdName is a wrapper around the C function gtk_paper_size_get_ppd_name.
func (recv *PaperSize) GetPpdName() string {
	retC := C.gtk_paper_size_get_ppd_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWidth is a wrapper around the C function gtk_paper_size_get_width.
func (recv *PaperSize) GetWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_width((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// IsEqual is a wrapper around the C function gtk_paper_size_is_equal.
func (recv *PaperSize) IsEqual(size2 *PaperSize) bool {
	c_size2 := (*C.GtkPaperSize)(C.NULL)
	if size2 != nil {
		c_size2 = (*C.GtkPaperSize)(size2.ToC())
	}

	retC := C.gtk_paper_size_is_equal((*C.GtkPaperSize)(recv.native), c_size2)
	retGo := retC == C.TRUE

	return retGo
}

// SetSize is a wrapper around the C function gtk_paper_size_set_size.
func (recv *PaperSize) SetSize(width float64, height float64, unit Unit) {
	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_paper_size_set_size((*C.GtkPaperSize)(recv.native), c_width, c_height, c_unit)

	return
}

// Equals compares this RecentInfo with another RecentInfo, and returns true if they represent the same GObject.
func (recv *RecentInfo) Equals(other *RecentInfo) bool {
	return other.ToC() == recv.ToC()
}

// CreateAppInfo is a wrapper around the C function gtk_recent_info_create_app_info.
func (recv *RecentInfo) CreateAppInfo(appName string) (*gio.AppInfo, error) {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	var cThrowableError *C.GError

	retC := C.gtk_recent_info_create_app_info((*C.GtkRecentInfo)(recv.native), c_app_name, &cThrowableError)
	var retGo (*gio.AppInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.AppInfoNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Exists is a wrapper around the C function gtk_recent_info_exists.
func (recv *RecentInfo) Exists() bool {
	retC := C.gtk_recent_info_exists((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetAdded is a wrapper around the C function gtk_recent_info_get_added.
func (recv *RecentInfo) GetAdded() int64 {
	retC := C.gtk_recent_info_get_added((*C.GtkRecentInfo)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetAge is a wrapper around the C function gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	retC := C.gtk_recent_info_get_age((*C.GtkRecentInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetApplicationInfo is a wrapper around the C function gtk_recent_info_get_application_info.
func (recv *RecentInfo) GetApplicationInfo(appName string) (bool, string, uint32, int64) {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	var c_app_exec *C.gchar

	var c_count C.guint

	var c_time_ C.time_t

	retC := C.gtk_recent_info_get_application_info((*C.GtkRecentInfo)(recv.native), c_app_name, &c_app_exec, &c_count, &c_time_)
	retGo := retC == C.TRUE

	appExec := C.GoString(c_app_exec)

	count := (uint32)(c_count)

	time := (int64)(c_time_)

	return retGo, appExec, count, time
}

// GetApplications is a wrapper around the C function gtk_recent_info_get_applications.
func (recv *RecentInfo) GetApplications() ([]string, uint64) {
	var c_length C.gsize

	retC := C.gtk_recent_info_get_applications((*C.GtkRecentInfo)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetDescription is a wrapper around the C function gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	retC := C.gtk_recent_info_get_description((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	retC := C.gtk_recent_info_get_display_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetGroups is a wrapper around the C function gtk_recent_info_get_groups.
func (recv *RecentInfo) GetGroups() ([]string, uint64) {
	var c_length C.gsize

	retC := C.gtk_recent_info_get_groups((*C.GtkRecentInfo)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetIcon is a wrapper around the C function gtk_recent_info_get_icon.
func (recv *RecentInfo) GetIcon(size int32) *gdkpixbuf.Pixbuf {
	c_size := (C.gint)(size)

	retC := C.gtk_recent_info_get_icon((*C.GtkRecentInfo)(recv.native), c_size)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMimeType is a wrapper around the C function gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	retC := C.gtk_recent_info_get_mime_type((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetModified is a wrapper around the C function gtk_recent_info_get_modified.
func (recv *RecentInfo) GetModified() int64 {
	retC := C.gtk_recent_info_get_modified((*C.GtkRecentInfo)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetPrivateHint is a wrapper around the C function gtk_recent_info_get_private_hint.
func (recv *RecentInfo) GetPrivateHint() bool {
	retC := C.gtk_recent_info_get_private_hint((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShortName is a wrapper around the C function gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	retC := C.gtk_recent_info_get_short_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	retC := C.gtk_recent_info_get_uri((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUriDisplay is a wrapper around the C function gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	retC := C.gtk_recent_info_get_uri_display((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetVisited is a wrapper around the C function gtk_recent_info_get_visited.
func (recv *RecentInfo) GetVisited() int64 {
	retC := C.gtk_recent_info_get_visited((*C.GtkRecentInfo)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// HasApplication is a wrapper around the C function gtk_recent_info_has_application.
func (recv *RecentInfo) HasApplication(appName string) bool {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	retC := C.gtk_recent_info_has_application((*C.GtkRecentInfo)(recv.native), c_app_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasGroup is a wrapper around the C function gtk_recent_info_has_group.
func (recv *RecentInfo) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.gtk_recent_info_has_group((*C.GtkRecentInfo)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// IsLocal is a wrapper around the C function gtk_recent_info_is_local.
func (recv *RecentInfo) IsLocal() bool {
	retC := C.gtk_recent_info_is_local((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// LastApplication is a wrapper around the C function gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	retC := C.gtk_recent_info_last_application((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Match is a wrapper around the C function gtk_recent_info_match.
func (recv *RecentInfo) Match(infoB *RecentInfo) bool {
	c_info_b := (*C.GtkRecentInfo)(C.NULL)
	if infoB != nil {
		c_info_b = (*C.GtkRecentInfo)(infoB.ToC())
	}

	retC := C.gtk_recent_info_match((*C.GtkRecentInfo)(recv.native), c_info_b)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	retC := C.gtk_recent_info_ref((*C.GtkRecentInfo)(recv.native))
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	C.gtk_recent_info_unref((*C.GtkRecentInfo)(recv.native))

	return
}

// Equals compares this RecentManagerClass with another RecentManagerClass, and returns true if they represent the same GObject.
func (recv *RecentManagerClass) Equals(other *RecentManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// TargetsIncludeRichText is a wrapper around the C function gtk_selection_data_targets_include_rich_text.
func (recv *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) bool {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	retC := C.gtk_selection_data_targets_include_rich_text((*C.GtkSelectionData)(recv.native), c_buffer)
	retGo := retC == C.TRUE

	return retGo
}

// TargetsIncludeUri is a wrapper around the C function gtk_selection_data_targets_include_uri.
func (recv *SelectionData) TargetsIncludeUri() bool {
	retC := C.gtk_selection_data_targets_include_uri((*C.GtkSelectionData)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddRichTextTargets is a wrapper around the C function gtk_target_list_add_rich_text_targets.
func (recv *TargetList) AddRichTextTargets(info uint32, deserializable bool, buffer *TextBuffer) {
	c_info := (C.guint)(info)

	c_deserializable :=
		boolToGboolean(deserializable)

	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	C.gtk_target_list_add_rich_text_targets((*C.GtkTargetList)(recv.native), c_info, c_deserializable, c_buffer)

	return
}
