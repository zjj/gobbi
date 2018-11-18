package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

type Constant struct {
	Namespace *Namespace

	Name      string `xml:"name,attr"`
	Blacklist bool   `xml:"blacklist,attr"`
	Value     string `xml:"value,attr"`
	Version   string `xml:"version,attr"`
	CType     string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Doc       *Doc   `xml:"doc"`
	Type      *Type  `xml:"type"`

	goTypeName string
}

func (c *Constant) init(ns *Namespace) {
	c.Namespace = ns

	switch c.Type.Name {
	case "gboolean":
		c.goTypeName = "bool"
	case "gint":
		c.goTypeName = "int"
	case "utf8":
		c.goTypeName = "string"
	}
}

func (c *Constant) version() string {
	return c.Version
}

func (c *Constant) mergeAddenda(addenda *Constant) {
	c.Blacklist = addenda.Blacklist
}

func (c *Constant) blacklisted() (bool, string) {
	return c.Blacklist, c.Name
}

func (c *Constant) supported() (supported bool, reason string) {
	switch c.Type.Name {
	case "gboolean", "gint", "utf8":
		return true, ""
	default:
		return false, fmt.Sprintf("type %s for %s", c.Type.Name, c.Name)
	}
}

func (c *Constant) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(c, version) {
		return
	}

	if c.Type.Name == "gboolean" {
		g.
			Const().
			Id(c.Name).
			Id(c.goTypeName).
			Op("=").
			Lit(c.Value == "true").
			Commentf("C.%s", c.CType)

		return
	}

	g.
		Const().
		Id(c.Name).
		Id(c.goTypeName).
		Op("=").
		Qual("C", c.CType)
}

func (c *Constant) generateDocs(file *DocFile) {
	doc := ""
	if c.Doc != nil {
		doc = c.Doc.Text
		// escape any '#'s
		doc = strings.Replace(doc, "\n#", "\n&num;", -1)
	}

	file.writeLinef("## `%s`", c.Name)
	file.writeLine("")

	file.writeLine(doc)
	file.writeLine("")

	file.writeLinef("C - `%s`", c.CType)
	file.writeLine("")
}
