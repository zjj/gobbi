// This is a generated file - DO NOT EDIT
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_attr_iterator_get_font : unsupported parameter language : record with indirection level of 2

// Unsupported : pango_attr_iterator_range : unsupported parameter start : no type generator for gint, gint*

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc, PangoAttrFilterFunc

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : no param type

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// Blacklisted : PangoFontMetrics

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// Unsupported : pango_glyph_item_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs : no param type

// Unsupported : pango_glyph_string_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_string_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_glyph_string_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Blacklisted : PangoIncludedModule

// Unsupported : pango_language_get_scripts : unsupported parameter num_scripts : no type generator for gint, int*

// IncludesScript is a wrapper around the C function pango_language_includes_script.
func (recv *Language) IncludesScript(script Script) bool {
	c_script := (C.PangoScript)(script)

	retC := C.pango_language_includes_script((*C.PangoLanguage)(recv.native), c_script)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : PangoLayoutIter

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : no param type

// Unsupported : pango_layout_line_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_layout_line_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Blacklisted : PangoScriptForLang

// Free is a wrapper around the C function pango_script_iter_free.
func (recv *ScriptIter) Free() {
	C.pango_script_iter_free((*C.PangoScriptIter)(recv.native))

	return
}

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Next is a wrapper around the C function pango_script_iter_next.
func (recv *ScriptIter) Next() bool {
	retC := C.pango_script_iter_next((*C.PangoScriptIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2
