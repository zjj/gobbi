package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type TypeGeneratorInteger struct {
	TypeGeneratorPanic
	typ *Type
}

func TypeGeneratorIntegerNew(typ *Type) *TypeGeneratorInteger {
	return &TypeGeneratorInteger{
		typ: typ,
	}
}

func (t *TypeGeneratorInteger) isSupportedAsParam(direction string) (supported bool, reason string) {
	if t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsArrayParam(direction string) (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsArrayParamC(direction string) (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorInteger) isSupportedAsParamC() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsField() (supported bool, reason string) {
	if t.typ.indirectLevel > 0 {
		return false, fmt.Sprintf("%s with indirection level of %d",
			t.typ.CType, t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsArrayReturnValue() (supported bool, reason string) {
	return false, ""
}

func (t *TypeGeneratorInteger) isSupportedAsReturnValue() (supported bool, reason string) {
	if strings.HasPrefix(t.typ.Name, "G") || t.typ.Name == "Quark" {
		return false, ""
	}

	if strings.Contains(t.typ.Name, ".") {
		return false, ""
	}

	return true, ""
}

func (t *TypeGeneratorInteger) isSupportedAsReturnCValue() (supported bool, reason string) {
	return true, ""
}

func (t *TypeGeneratorInteger) generateDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInteger) generateArrayDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Index().
		Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInteger) generateArrayDeclarationC(g *jen.Group, cVarName string) {
}

func (t *TypeGeneratorInteger) generateDeclarationC(g *jen.Group, cVarName string) {
	g.Id(cVarName).Qual("C", t.typ.CType)
}

func (t *TypeGeneratorInteger) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.
		Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("&")
			}
		}).
		Id(cVarName)
}

func (t *TypeGeneratorInteger) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorInteger) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Parens(jen.Qual("C", t.typ.cTypeName)).
		Parens(jen.Id(goVarName))
}

func (t *TypeGeneratorInteger) generateParamGoVar(g *jen.Group, goVarName string, cVarName string, pkg string) {
	g.
		Id(goVarName).
		Op(":=").
		Id(integerCTypeMap[t.typ.CType]).
		Call(jen.Id(cVarName))

}

func (t *TypeGeneratorInteger) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorInteger) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Do(t.typ.qname.generate)
}

func (t *TypeGeneratorInteger) generateReturnFunctionDeclarationCtype(g *jen.Group) {
	g.Qual("C", t.typ.CType)
}

func (t *TypeGeneratorInteger) generateReturnCToGo(g *jen.Group, isParam bool,
	cVarName string, goVarName string, pkg string, transferOwnership string, nullable bool) {
	g.
		Id(goVarName).
		Op(":=").
		Parens(jen.Do(t.typ.qname.generate)).
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.Name == "gpointer" {
				s.
					Qual("unsafe", "Pointer").
					CallFunc(func(g *jen.Group) {
						if t.typ.indirectLevel == 1 {
							g.
								Op("&").
								Id(cVarName)
						} else {
							g.Id(cVarName)
						}
					})
			} else {
				s.Id(cVarName)
			}
		}))
}

func (t *TypeGeneratorInteger) generateCToGo(pkg string, cVarReference *jen.Statement) *jen.Statement {
	return jen.
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("*")
			}
			t.typ.qname.generate(s)
		})).
		Parens(jen.Do(func(s *jen.Statement) {
			if t.typ.indirectLevel == 1 {
				s.Op("&")
			}
			s.Add(cVarReference)
		}))
}

func (t *TypeGeneratorInteger) generateGoToC(g *jen.Group, goVarReference *jen.Statement) {
	g.
		Parens(jen.Qual("C", t.typ.cTypeName)).
		Parens(goVarReference)
}

func (t *TypeGeneratorInteger) generateCallReturnType() string {
	switch t.typ.CType {
	case "void":
		return "RT_VOID"
	case "int", "gint", "gint32":
		return "RT_INT"
	case "gint64":
		return "RT_LONG"
	case "guint", "guint32":
		return "RT_UINT"
	case "gdouble", "double":
		return "RT_DOUBLE"
	}

	panic(fmt.Sprintf("Return type not supported : %s", t.typ.CType))
}
