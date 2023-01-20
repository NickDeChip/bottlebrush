package ast

import "github.com/NickDeChip/bottle-brush/pkg/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}
func (sl *StringLiteral) Line() int {
	return sl.Token.Line
}
func (sl *StringLiteral) Col() int {
	return sl.Token.Col
}
