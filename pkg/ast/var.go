package ast

import (
	"bytes"

	"github.com/NickDeChip/bottle-brush/pkg/token"
)

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
	Mut   bool
}

func (vs *VarStatement) statementNode() {}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.Name.String() + " ")
	out.WriteString(vs.TokenLiteral() + " ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	return out.String()
}

func (vs *VarStatement) Line() int {
	return vs.Name.Token.Line
}
func (vs *VarStatement) Col() int {
	return vs.Name.Token.Line
}
