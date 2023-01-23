package ast

import (
	"bytes"

	"github.com/NickDeChip/bottle-brush/pkg/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	return out.String()
}

func (rs *ReturnStatement) Line() int {
	return rs.Token.Line
}

func (rs *ReturnStatement) Col() int {
	return rs.Token.Col
}
