// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

/*

	static void _g_variant_builder_add_parsed(GVariantBuilder* builder, const gchar* format) {
		return g_variant_builder_add_parsed(builder, format);
    }
*/
import "C"

// Equals compares this DateTime with another DateTime, and returns true if they represent the same GObject.
func (recv *DateTime) Equals(other *DateTime) bool {
	return other.ToC() == recv.ToC()
}

// g_date_time_compare : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_equal : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_hash : unsupported parameter datetime : no type generator for gpointer (gconstpointer) for param datetime
// Add is a wrapper around the C function g_date_time_add.
func (recv *DateTime) Add(timespan TimeSpan) *DateTime {
	c_timespan := (C.GTimeSpan)(timespan)

	retC := C.g_date_time_add((*C.GDateTime)(recv.native), c_timespan)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddDays is a wrapper around the C function g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) *DateTime {
	c_days := (C.gint)(days)

	retC := C.g_date_time_add_days((*C.GDateTime)(recv.native), c_days)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFull is a wrapper around the C function g_date_time_add_full.
func (recv *DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) *DateTime {
	c_years := (C.gint)(years)

	c_months := (C.gint)(months)

	c_days := (C.gint)(days)

	c_hours := (C.gint)(hours)

	c_minutes := (C.gint)(minutes)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_full((*C.GDateTime)(recv.native), c_years, c_months, c_days, c_hours, c_minutes, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddHours is a wrapper around the C function g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) *DateTime {
	c_hours := (C.gint)(hours)

	retC := C.g_date_time_add_hours((*C.GDateTime)(recv.native), c_hours)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMinutes is a wrapper around the C function g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	c_minutes := (C.gint)(minutes)

	retC := C.g_date_time_add_minutes((*C.GDateTime)(recv.native), c_minutes)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMonths is a wrapper around the C function g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) *DateTime {
	c_months := (C.gint)(months)

	retC := C.g_date_time_add_months((*C.GDateTime)(recv.native), c_months)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddSeconds is a wrapper around the C function g_date_time_add_seconds.
func (recv *DateTime) AddSeconds(seconds float64) *DateTime {
	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_seconds((*C.GDateTime)(recv.native), c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWeeks is a wrapper around the C function g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	c_weeks := (C.gint)(weeks)

	retC := C.g_date_time_add_weeks((*C.GDateTime)(recv.native), c_weeks)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddYears is a wrapper around the C function g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) *DateTime {
	c_years := (C.gint)(years)

	retC := C.g_date_time_add_years((*C.GDateTime)(recv.native), c_years)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Difference is a wrapper around the C function g_date_time_difference.
func (recv *DateTime) Difference(begin *DateTime) TimeSpan {
	c_begin := (*C.GDateTime)(C.NULL)
	if begin != nil {
		c_begin = (*C.GDateTime)(begin.ToC())
	}

	retC := C.g_date_time_difference((*C.GDateTime)(recv.native), c_begin)
	retGo := (TimeSpan)(retC)

	return retGo
}

// Format is a wrapper around the C function g_date_time_format.
func (recv *DateTime) Format(format string) string {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	retC := C.g_date_time_format((*C.GDateTime)(recv.native), c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDayOfMonth is a wrapper around the C function g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() int32 {
	retC := C.g_date_time_get_day_of_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfWeek is a wrapper around the C function g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() int32 {
	retC := C.g_date_time_get_day_of_week((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHour is a wrapper around the C function g_date_time_get_hour.
func (recv *DateTime) GetHour() int32 {
	retC := C.g_date_time_get_hour((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMicrosecond is a wrapper around the C function g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() int32 {
	retC := C.g_date_time_get_microsecond((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinute is a wrapper around the C function g_date_time_get_minute.
func (recv *DateTime) GetMinute() int32 {
	retC := C.g_date_time_get_minute((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	retC := C.g_date_time_get_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSecond is a wrapper around the C function g_date_time_get_second.
func (recv *DateTime) GetSecond() int32 {
	retC := C.g_date_time_get_second((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSeconds is a wrapper around the C function g_date_time_get_seconds.
func (recv *DateTime) GetSeconds() float64 {
	retC := C.g_date_time_get_seconds((*C.GDateTime)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetTimezoneAbbreviation is a wrapper around the C function g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() string {
	retC := C.g_date_time_get_timezone_abbreviation((*C.GDateTime)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUtcOffset is a wrapper around the C function g_date_time_get_utc_offset.
func (recv *DateTime) GetUtcOffset() TimeSpan {
	retC := C.g_date_time_get_utc_offset((*C.GDateTime)(recv.native))
	retGo := (TimeSpan)(retC)

	return retGo
}

// GetWeekNumberingYear is a wrapper around the C function g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() int32 {
	retC := C.g_date_time_get_week_numbering_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWeekOfYear is a wrapper around the C function g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() int32 {
	retC := C.g_date_time_get_week_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	retC := C.g_date_time_get_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYmd is a wrapper around the C function g_date_time_get_ymd.
func (recv *DateTime) GetYmd() (int32, int32, int32) {
	var c_year C.gint

	var c_month C.gint

	var c_day C.gint

	C.g_date_time_get_ymd((*C.GDateTime)(recv.native), &c_year, &c_month, &c_day)

	year := (int32)(c_year)

	month := (int32)(c_month)

	day := (int32)(c_day)

	return year, month, day
}

// IsDaylightSavings is a wrapper around the C function g_date_time_is_daylight_savings.
func (recv *DateTime) IsDaylightSavings() bool {
	retC := C.g_date_time_is_daylight_savings((*C.GDateTime)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_date_time_ref.
func (recv *DateTime) Ref() *DateTime {
	retC := C.g_date_time_ref((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToLocal is a wrapper around the C function g_date_time_to_local.
func (recv *DateTime) ToLocal() *DateTime {
	retC := C.g_date_time_to_local((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToTimeval is a wrapper around the C function g_date_time_to_timeval.
func (recv *DateTime) ToTimeval(tv *TimeVal) bool {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_to_timeval((*C.GDateTime)(recv.native), c_tv)
	retGo := retC == C.TRUE

	return retGo
}

// ToTimezone is a wrapper around the C function g_date_time_to_timezone.
func (recv *DateTime) ToTimezone(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_to_timezone((*C.GDateTime)(recv.native), c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToUnix is a wrapper around the C function g_date_time_to_unix.
func (recv *DateTime) ToUnix() int64 {
	retC := C.g_date_time_to_unix((*C.GDateTime)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// ToUtc is a wrapper around the C function g_date_time_to_utc.
func (recv *DateTime) ToUtc() *DateTime {
	retC := C.g_date_time_to_utc((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_date_time_unref.
func (recv *DateTime) Unref() {
	C.g_date_time_unref((*C.GDateTime)(recv.native))

	return
}

// GetInt64 is a wrapper around the C function g_key_file_get_int64.
func (recv *KeyFile) GetInt64(groupName string, key string) (int64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetUint64 is a wrapper around the C function g_key_file_get_uint64.
func (recv *KeyFile) GetUint64(groupName string, key string) (uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (uint64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetInt64 is a wrapper around the C function g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	C.g_key_file_set_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetUint64 is a wrapper around the C function g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	C.g_key_file_set_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SourceSetNameById is a wrapper around the C function g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) {
	c_tag := (C.guint)(tag)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name_by_id(c_tag, c_name)

	return
}

// GetName is a wrapper around the C function g_source_get_name.
func (recv *Source) GetName() string {
	retC := C.g_source_get_name((*C.GSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetName is a wrapper around the C function g_source_set_name.
func (recv *Source) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name((*C.GSource)(recv.native), c_name)

	return
}

// Equals compares this TimeZone with another TimeZone, and returns true if they represent the same GObject.
func (recv *TimeZone) Equals(other *TimeZone) bool {
	return other.ToC() == recv.ToC()
}

// AdjustTime is a wrapper around the C function g_time_zone_adjust_time.
func (recv *TimeZone) AdjustTime(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_adjust_time((*C.GTimeZone)(recv.native), c_type, &c_time_)
	retGo := (int32)(retC)

	return retGo
}

// FindInterval is a wrapper around the C function g_time_zone_find_interval.
func (recv *TimeZone) FindInterval(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_find_interval((*C.GTimeZone)(recv.native), c_type, c_time_)
	retGo := (int32)(retC)

	return retGo
}

// GetAbbreviation is a wrapper around the C function g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_abbreviation((*C.GTimeZone)(recv.native), c_interval)
	retGo := C.GoString(retC)

	return retGo
}

// GetOffset is a wrapper around the C function g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) int32 {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_offset((*C.GTimeZone)(recv.native), c_interval)
	retGo := (int32)(retC)

	return retGo
}

// IsDst is a wrapper around the C function g_time_zone_is_dst.
func (recv *TimeZone) IsDst(interval int32) bool {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_is_dst((*C.GTimeZone)(recv.native), c_interval)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_time_zone_ref.
func (recv *TimeZone) Ref() *TimeZone {
	retC := C.g_time_zone_ref((*C.GTimeZone)(recv.native))
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_time_zone_unref.
func (recv *TimeZone) Unref() {
	C.g_time_zone_unref((*C.GTimeZone)(recv.native))

	return
}

// Compare is a wrapper around the C function g_variant_compare.
func (recv *Variant) Compare(two *Variant) int32 {
	c_two := (C.gconstpointer)(C.NULL)
	if two != nil {
		c_two = (C.gconstpointer)(two.ToC())
	}

	retC := C.g_variant_compare((C.gconstpointer)(recv.native), c_two)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_variant_dup_bytestring : array return type :

// DupBytestringArray is a wrapper around the C function g_variant_dup_bytestring_array.
func (recv *Variant) DupBytestringArray() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_bytestring_array((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// Unsupported : g_variant_get_bytestring : array return type :

// GetBytestringArray is a wrapper around the C function g_variant_get_bytestring_array.
func (recv *Variant) GetBytestringArray() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_bytestring_array((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}

// IsFloating is a wrapper around the C function g_variant_is_floating.
func (recv *Variant) IsFloating() bool {
	retC := C.g_variant_is_floating((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddParsed is a wrapper around the C function g_variant_builder_add_parsed.
func (recv *VariantBuilder) AddParsed(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_variant_builder_add_parsed((*C.GVariantBuilder)(recv.native), c_format)

	return
}
