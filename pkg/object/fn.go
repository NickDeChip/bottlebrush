package object

import (
	"bytes"
	"strings"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
)

type Fn struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Fn) Type() Type {
	return FUNTION
}

func (f *Fn) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") start\n")
	out.WriteString(f.Body.String())
	out.WriteString("\nend")

	return out.String()
}
