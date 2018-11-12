package generate

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

type Enumeration struct {
	Namespace *Namespace

	Name         string  `xml:"name,attr"`
	Blacklist    bool    `xml:"blacklist,attr"`
	GoTypeName   string  `xml:"goname,attr"` // used in addenda files
	Version      string  `xml:"version,attr"`
	CType        string  `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GlibTypeName string  `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GlibGetType  string  `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	Doc          *Doc    `xml:"doc"`
	Members      Members `xml:"member"`

	goTypeName string
}

func (e *Enumeration) init(ns *Namespace) {
	e.Namespace = ns

	e.goTypeName = e.Name
	if e.GoTypeName != "" {
		e.goTypeName = e.GoTypeName
	}

	for _, member := range e.Members {
		member.Namespace = ns
	}
}

func (e *Enumeration) version() string {
	return e.Version
}

func (e *Enumeration) mergeAddenda(addenda *Enumeration) {
	e.Blacklist = addenda.Blacklist
	e.GoTypeName = addenda.GoTypeName
	if addenda.CType != "" {
		e.CType = addenda.CType
	}
	if addenda.Version != "" {
		e.Version = addenda.Version
	}
	e.Members.mergeAddenda(addenda.Members)
}

func (e *Enumeration) blacklisted() (bool, string) {
	return e.Blacklist, e.CType
}

func (e *Enumeration) supported() (supported bool, reason string) {
	return true, ""
}

func (e *Enumeration) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(e, version) {
		return
	}

	// define the type
	g.
		Type().
		Id(e.goTypeName).
		Qual("C", e.CType)

	// define members
	memberNamePrefix := e.Namespace.CIdentifierPrefixes + "_"
	g.Const().DefsFunc(func(g *jen.Group) {
		for _, member := range e.Members {
			member.generate(g, memberNamePrefix, e.goTypeName)
		}
	})
}

func (e *Enumeration) generateDocs(file *DocFile) {
	doc := ""
	if e.Doc != nil {
		doc = e.Doc.Text
		// escape any '#'s
		doc = strings.Replace(doc, "\n#", "\n&num;", -1)
	}

	file.writeLinef("## `%s`", e.Name)
	file.writeLine("")

	file.writeLine(doc)
	file.writeLine("")

	for _, member := range e.Members {
		member.generateDocs(file)
	}
	file.writeLine("")

	file.writeLinef("C - `%s`", e.CType)
	file.writeLine("")
}
