package generate

import (
	"github.com/pekim/jennifer/jen"
)

type Method struct {
	*Function
}

func (m *Method) init(ns *Namespace, record *Record) {
	m.Function.init(ns, record, "")

	if record.Version != "" && m.Version == "" {
		m.Version = record.Version
	}

}

func (m *Method) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(m, version) {
		return
	}

	supported, reason := m.supported()
	if !supported {
		g.Commentf("Unsupported : %s", reason)
		g.Line()
		return
	}

	if blacklisted, detail := m.blacklisted(); blacklisted {
		g.Commentf("Blacklisted : %s", detail)
		g.Line()
		return
	}

	m.Function.generate(g, version)
}
