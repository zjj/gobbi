// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrLetterSpacingNew is a wrapper around the C function pango_attr_letter_spacing_new.
func AttrLetterSpacingNew(letterSpacing int32) *Attribute {
	c_letter_spacing := (C.int)(letterSpacing)

	retC := C.pango_attr_letter_spacing_new(c_letter_spacing)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_break : unsupported parameter attrs : no param type

// Blacklisted : pango_config_key_get

// Blacklisted : pango_config_key_get_system

// Blacklisted : pango_default_break

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Unsupported : pango_find_paragraph_boundary : unsupported parameter paragraph_delimiter_index : no type generator for gint, gint*

// Blacklisted : pango_get_lib_subdirectory

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs : no param type

// Unsupported : pango_get_mirror_char : unsupported parameter mirrored_ch : no type generator for gunichar, gunichar*

// Blacklisted : pango_get_sysconf_subdirectory

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

// Unsupported : pango_lookup_aliases : unsupported parameter families : no param type

// Unsupported : pango_markup_parser_finish : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Unsupported : pango_parse_enum : unsupported parameter type : no type generator for GType, GType

// Unsupported : pango_parse_markup : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_quantize_line_geometry : unsupported parameter thickness : no type generator for gint, int*

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer, FILE*

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type