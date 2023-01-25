package ast

import "github.com/NickDeChip/bottlebrush/pkg/token"

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) Line() int {
	return il.Token.Line
}

func (il *IntegerLiteral) Col() int {
	return il.Token.Col
}
