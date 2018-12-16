package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Function struct {
	Namespace *Namespace

	Name              string       `xml:"name,attr"`
	Blacklist         bool         `xml:"blacklist,attr"`
	GoName            string       `xml:"goname,attr"`
	Version           string       `xml:"version,attr"`
	MovedTo           string       `xml:"moved-to,attr"`
	CIdentifier       string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	Deprecated        int          `xml:"deprecated,attr"`
	DeprecatedVersion string       `xml:"deprecated-version,attr"`
	Doc               *Doc         `xml:"doc"`
	InstanceParameter *Parameter   `xml:"parameters>instance-parameter"`
	Parameters        Parameters   `xml:"parameters>parameter"`
	ReturnValue       *ReturnValue `xml:"return-value"`
	Throws            int          `xml:"throws,attr"`
	Introspectable    string       `xml:"introspectable,attr"`

	receiver   *Record
	ctorRecord *Record

	throwableErrorType      *Type
	throwableErrorCVarName  string
	throwableErrorGoVarName string
}

func (f *Function) init(ns *Namespace, receiver *Record, namePrefix string) {
	f.Namespace = ns
	f.receiver = receiver
	if f.GoName == "" {
		f.GoName = makeExportedGoName(f.Name)
	}
	f.GoName = namePrefix + f.GoName
	f.Parameters.init(ns)
	if f.InstanceParameter != nil {
		f.InstanceParameter.init(ns)
	}
	f.initThrowableError()

	if f.ReturnValue != nil {
		f.ReturnValue.init(ns)
	}
}

func (f *Function) initThrowableError() {
	if f.Throws == 0 {
		return
	}

	typ := &Type{
		Name:  "Error",
		CType: "GError**",
	}
	typ.init(f.Namespace.get("GLib"))

	f.throwableErrorType = typ
	f.throwableErrorCVarName = "cThrowableError"
	f.throwableErrorGoVarName = "goThrowableError"
}

func (f *Function) version() string {
	return f.Version
}

func (f *Function) mergeAddenda(addenda *Function) {
	f.Blacklist = addenda.Blacklist

	if f.ReturnValue != nil && addenda.ReturnValue != nil {
		f.ReturnValue.mergeAddenda(addenda.ReturnValue)
	}

	if addenda.GoName != "" {
		f.GoName = addenda.GoName
	}
}

func (f *Function) blacklisted() (bool, string) {
	return f.Blacklist, f.CIdentifier
}

func (f *Function) supported() (supported bool, reason string) {
	if supported, reason := f.Parameters.allSupported(); !supported {
		return false, fmt.Sprintf("%s : %s", f.CIdentifier, reason)
	}

	if supported, reason := f.ReturnValue.isSupported(); !supported {
		return false, fmt.Sprintf("%s : %s", f.CIdentifier, reason)
	}

	return true, ""
}

func (f *Function) generate(g *jen.Group, version *Version) {
	if !supportedByVersion(f, version) {
		return
	}

	if f.MovedTo != "" {
		return
	}

	supported, reason := f.supported()
	if !supported {
		g.Comment(reason)
		return
	}

	if blacklisted, detail := f.blacklisted(); blacklisted {
		g.Commentf("Blacklisted : %s", detail)
		g.Line()
		return
	}

	g.Commentf("%s is a wrapper around the C function %s.", f.GoName, f.CIdentifier)

	g.
		Func().
		Do(f.generateReceiverDeclaration).
		Id(f.GoName).                                         // name
		ParamsFunc(f.Parameters.generateFunctionDeclaration). // params
		ParamsFunc(f.generateReturnDeclaration).              // returns
		BlockFunc(f.generateBody).                            // body
		Line()
}

func (f *Function) generateReceiverDeclaration(s *jen.Statement) {
	if f.receiver == nil {
		return
	}

	s.Parens(jen.
		Id("recv").
		Op("*").
		Id(f.receiver.GoName))
}

func (f *Function) generateReturnDeclaration(g *jen.Group) {
	f.ReturnValue.generateFunctionDeclaration(g)
	f.Parameters.generateOutputParamsReturnDeclaration(g)
	f.generateThrowableReturnDeclaration(g)
}

func (f *Function) generateBody(g *jen.Group) {
	f.generateCParameterVars(g)
	f.generateCall(g)

	f.generateGoReturnVars(g)
	f.generateOutputParamsGoVars(g)
	f.generateReturn(g)
}

func (f *Function) generateCParameterVars(g *jen.Group) {
	f.Parameters.generateCVars(g)
	f.generateThrowableErrorCVar(g)
}

func (f *Function) generateCall(g *jen.Group) *jen.Statement {
	return g.
		Do(func(s *jen.Statement) {
			if f.ReturnValue.Type.Name != "none" {
				s.
					Id("retC").
					Op(":=")
			}
		}).
		Qual("C", f.CIdentifier).
		CallFunc(func(g *jen.Group) {
			// Assumption that receiver is always first agumment.
			// If turns out not to be true, will need to look at
			// using <InstanceParameter> position.
			f.generateReceiverArgument(g)

			f.Parameters.generateCallArguments(g)
			f.generateThrowableCallArgument(g)
		})
}

func (f *Function) generateReceiverArgument(g *jen.Group) {
	if f.receiver == nil {
		return
	}

	typ := f.InstanceParameter.Type

	g.
		Parens(jen.
			Op(strings.Repeat("*", typ.indirectLevel)).
			Qual("C", typ.cTypeName)).
		Parens(jen.
			Id("recv").
			Op(".").
			Id("native"))
}

func (f *Function) generateOutputParamsGoVars(g *jen.Group) {
	f.Parameters.generateOutputParamsGoVars(g)
}

func (f *Function) generateGoReturnVars(g *jen.Group) {
	f.generateReturnGoVar(g)
	f.generateThrowableReturnGoVar(g)
}

func (f *Function) generateReturn(g *jen.Group) {
	g.ReturnFunc(func(g *jen.Group) {
		if f.ReturnValue.Type.Name != "none" {
			g.Id("retGo")
		}

		f.Parameters.generateOutputParamsReturns(g)
		f.generateThrowableReturn(g)
	})
}

func (f *Function) generateReturnGoVar(g *jen.Group) {
	if f.ReturnValue.Type.Name != "none" {
		f.ReturnValue.generateCToGo(g, "retC", "retGo")

		r := f.ctorRecord
		if f.ctorRecord != nil &&
			r.isDerivedFrom("gobject", "Object") &&
			!r.isDerivedFrom("gobject", "InitiallyUnowned") {

			g.Line()

			// A call to a ...NewFromC function will have incremented the ref
			// count (to 2). So decrement it back to 1.
			g.
				If(jen.Id("retC").Op("!=").Nil()).
				Block(jen.
					Qual("C", "g_object_unref").
					Call(jen.
						Parens(jen.Qual("C", "gpointer")).
						Parens(jen.Id("retC"))))
		}
	}
	g.Line()
}

func (f *Function) generateThrowableReturnDeclaration(g *jen.Group) {
	if f.Throws == 0 {
		return
	}

	g.Id("error")
}

func (f *Function) generateThrowableReturn(g *jen.Group) {
	if f.Throws == 0 {
		return
	}

	g.Id("goError")
}

func (f *Function) generateThrowableReturnGoVar(g *jen.Group) {
	if f.Throws == 0 {
		return
	}

	// The returned error should be of type error, not just a type that implements error.
	// So make this explicit, and initially nil.
	//
	// see https://golang.org/doc/faq#nil_error
	g.Var().Id("goError").Id("error").Op("=").Nil()

	g.
		If(jen.Id(f.throwableErrorCVarName).Op("!=").Id("nil")).
		BlockFunc(func(g *jen.Group) {
			pkg := ""
			if f.Namespace.Name != "GLib" {
				pkg = f.Namespace.get("GLib").fullGoPackageName
			}
			// Create the error as an instance of glib.Error.
			f.throwableErrorType.generator.generateReturnCToGo(g, false,
				f.throwableErrorCVarName, f.throwableErrorGoVarName,
				pkg, "", false)

			g.Id("goError").Op("=").Id(f.throwableErrorGoVarName)

			// Free the error.
			g.Line()
			g.
				Qual("C", "g_error_free").
				Call(jen.Id(f.throwableErrorCVarName))
		})

	g.Line()
}

func (f *Function) generateThrowableErrorCVar(g *jen.Group) {
	if f.Throws == 0 {
		return
	}

	f.throwableErrorType.generator.generateParamOutCVar(g, f.throwableErrorCVarName)
	g.Line()
}

func (f *Function) generateThrowableCallArgument(g *jen.Group) {
	if f.Throws == 0 {
		return
	}

	f.throwableErrorType.generator.generateParamOutCallArgument(g, f.throwableErrorCVarName)
}
