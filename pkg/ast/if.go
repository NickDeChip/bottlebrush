package ast

import (
	"bytes"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type IfExspression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (ie *IfExspression) expressionNode() {}

func (ie *IfExspression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExspression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	return out.String()
}

func (ie *IfExspression) Line() int {
	return ie.Token.Line
}

func (ie *IfExspression) Col() int {
	return ie.Token.Col
}
