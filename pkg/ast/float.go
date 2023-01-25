package ast

import "github.com/NickDeChip/bottlebrush/pkg/token"

type FloatLiteral struct {
	Token token.Token
	Value float32
}

func (il *FloatLiteral) expressionNode() {}

func (il *FloatLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *FloatLiteral) String() string {
	return il.Token.Literal
}

func (il *FloatLiteral) Line() int {
	return il.Token.Line
}

func (il *FloatLiteral) Col() int {
	return il.Token.Col
}
