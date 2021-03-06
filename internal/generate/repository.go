package generate

import (
	"encoding/xml"
)

type Repository struct {
	XMLName   xml.Name   `xml:"repository"`
	Version   string     `xml:"version,attr"`
	CIncludes CIncludes  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Includes  []Include  `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Packages  []Package  `xml:"package"`
	Namespace *Namespace `xml:"namespace"`
}

func (r *Repository) Init() {
	r.Namespace.init(r)
}

func (r *Repository) MergeAddenda(addenda *Repository) {
	r.CIncludes = append(r.CIncludes, addenda.CIncludes...)
	r.Includes = append(r.Includes, addenda.Includes...)
	r.Namespace.mergeAddenda(addenda.Namespace)
}

func (r *Repository) Generate(addToTotal func(n int), oneDone func()) {
	r.Namespace.generate(addToTotal, oneDone)
}

type Package struct {
	Name string `xml:"name,attr"`
}
