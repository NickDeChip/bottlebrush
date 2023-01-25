package ast

import (
	"bytes"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type AssignmentStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *AssignmentStatement) statementNode() {}

func (vs *AssignmentStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *AssignmentStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.Name.String() + " ")
	out.WriteString(vs.TokenLiteral() + " ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	return out.String()
}

func (vs *AssignmentStatement) Line() int {
	return vs.Name.Token.Line
}
func (vs *AssignmentStatement) Col() int {
	return vs.Name.Token.Line
}
