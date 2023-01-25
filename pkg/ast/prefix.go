package ast

import (
	"bytes"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (ie *PrefixExpression) expressionNode() {}

func (ie *PrefixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Operator)
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *PrefixExpression) Line() int {
	return ie.Token.Line
}

func (ie *PrefixExpression) Col() int {
	return ie.Token.Col
}
