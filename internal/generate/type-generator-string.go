package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type TypeGeneratorString struct {
	typ *Type
}

func TypeGeneratorStringNew(typ *Type) *TypeGeneratorString {
	return &TypeGeneratorString{
		typ: typ,
	}
}

func (t *TypeGeneratorString) isSupportedAsParam(direction string) (supported bool, reason string) {
	if direction != "out" && t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("in string with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorString) isSupportedAsReturnValue() (supported bool, reason string) {
	if t.typ.indirectLevel > 1 {
		return false, fmt.Sprintf("string with indirection level of %d",
			t.typ.indirectLevel)
	}

	return true, ""
}

func (t *TypeGeneratorString) generateParamFunctionDeclaration(g *jen.Group, goVarName string) {
	g.
		Id(goVarName).
		Id(t.typ.goType)
}

func (t *TypeGeneratorString) generateParamCallArgument(g *jen.Group, cVarName string) {
	g.Id(cVarName)
}

func (t *TypeGeneratorString) generateParamOutCallArgument(g *jen.Group, cVarName string) {
	g.
		Op("&").
		Id(cVarName)
}

func (t *TypeGeneratorString) generateParamCVar(g *jen.Group, cVarName string, goVarName string, transferOwnership string) {
	g.
		Id(cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(goVarName))

	if transferOwnership != "none" {
		// ownership is transferred (to the library) so we should not
		// free the string memory
		return
	}

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(cVarName)))
}

func (t *TypeGeneratorString) generateParamOutCVar(g *jen.Group, cVarName string) {
	g.
		Var().
		Id(cVarName).
		Op(strings.Repeat("*", t.typ.indirectLevel-1)).
		Qual("C", t.typ.cTypeName)
}

func (t *TypeGeneratorString) generateReturnFunctionDeclaration(g *jen.Group) {
	g.Id(t.typ.goType)
}

func (t *TypeGeneratorString) generateReturnCToGo(g *jen.Group, cVarName string, transferOwnership string) {
	g.
		Qual("C", "GoString").
		Call(jen.Id(cVarName))

	if transferOwnership != "none" {
		// the library will be responsible for the memory
		return
	}

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(cVarName)))
}
