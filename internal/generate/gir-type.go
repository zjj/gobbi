package generate

import (
	"errors"
	"fmt"
	"github.com/dave/jennifer/jen"
	"strconv"
)

type Type struct {
	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	namespace *Namespace
}

func (t *Type) init(ns *Namespace) {
	t.namespace = ns
}

var jenGoTypes = map[string]*jen.Statement{
	// signed
	"gint8":  jen.Int8(),
	"gshort": jen.Int16(),
	"gint16": jen.Int16(),
	"int":    jen.Int32(),
	"gint":   jen.Int32(),
	"gint32": jen.Int32(),
	"glong":  jen.Int64(),
	"gint64": jen.Int64(),

	// unsigned
	"guchar":  jen.Uint8(),
	"guint8":  jen.Uint8(),
	"gushort": jen.Uint16(),
	"guint16": jen.Uint16(),
	"guint":   jen.Uint32(),
	"guint32": jen.Uint32(),
	"gulong":  jen.Uint64(),
	"guint64": jen.Uint64(),

	"gsize": jen.Uintptr(),
	"utf8":  jen.String(),
}

func (t *Type) jenGoType() (*jen.Statement, error) {
	if t == nil {
		return nil, errors.New("missing Type")
	}
	if t.CType == "" {
		return nil, errors.New("missing Type.Name")
	}

	if jenType, found := jenGoTypes[t.Name]; found {
		return jenType, nil
	}

	goType, ok := t.namespace.jenGoTypeForTypeName(t.Name)
	if ok {
		return goType, nil
	}

	return nil, fmt.Errorf("No Go type for '%s'\n", t.Name)
}

func (t *Type) jenValue(stringValue string) (*jen.Statement, error) {
	if t == nil {
		return nil, errors.New("missing Type")
	}

	switch t.Name {
	case "gchar", "gint8", "gshort", "gint16", "int", "gint", "gint32", "glong", "gint64":
		return t.jenValueInt(stringValue)
	case "guchar", "guint8", "gushort", "guint16", "guint", "guint32", "gulong", "guint64":
		return t.jenValueUint(stringValue)
	case "gdouble":
		value, err := strconv.ParseFloat(stringValue, 64)
		return jen.Lit(value), err
	case "gboolean":
		return jen.Lit(stringValue == "true"), nil
	case "utf8":
		return jen.Lit(stringValue), nil
	}

	return nil, fmt.Errorf("Cannot generate literal value for '%s'\n", t.Name)
}

func (t *Type) jenValueInt(stringValue string) (*jen.Statement, error) {
	value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return nil, err
	}

	switch t.Name {
	case "gchar", "gint8":
		return jen.Lit(int8(value)), nil
	case "gshort", "gint16":
		return jen.Lit(int16(value)), nil
	case "int", "gint", "gint32":
		return jen.Lit(int32(value)), nil
	case "glong", "gint64":
		return jen.Lit(int64(value)), nil
	}

	return nil, fmt.Errorf("Unknown type'%s'\n", t.Name)
}

func (t *Type) jenValueUint(stringValue string) (*jen.Statement, error) {
	value, err := strconv.ParseUint(stringValue, 10, 64)
	if err != nil {
		return nil, err
	}

	switch t.Name {
	case "guchar", "guint8":
		return jen.Lit(uint8(value)), nil
	case "gushort", "guint16":
		return jen.Lit(uint16(value)), nil
	case "guint", "guint32":
		return jen.Lit(uint32(value)), nil
	case "gulong", "guint64":
		return jen.Lit(uint64(value)), nil
	}

	return nil, fmt.Errorf("Unknown type'%s'\n", t.Name)
}

func (t *Type) supportedAsReturnValue() bool {
	if t == nil || t.Name == "none" {
		// return type is void
		return true
	}

	if _, ok := returnValueExtractFunctionNames[t.Name]; ok {
		return true
	}

	if t.Name == "utf8" {
		return true
	}

	return false
}

var returnValueExtractFunctionNames = map[string]string{
	"gchar":  "Int8",
	"gint8":  "Int8",
	"gshort": "Int16",
	"gint16": "Int16",
	"int":    "Int32",
	"gint":   "Int32",
	"gint32": "Int32",
	"glong":  "Int64",
	"gint64": "Int64",

	"guchar":  "Uint8",
	"guint8":  "Uint8",
	"gushort": "Uint16",
	"guint16": "Uint16",
	"guint":   "Uint32",
	"guint32": "Uint32",
	"gulong":  "Uint64",
	"guint64": "Uint64",

	"utf8": "String",
}

func (t *Type) returnValueExtractFunctionName() string {
	return returnValueExtractFunctionNames[t.Name]
}
