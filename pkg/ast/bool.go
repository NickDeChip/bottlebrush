package ast

import "github.com/NickDeChip/bottle-brush/pkg/token"

type Bool struct {
	Token token.Token
	Value bool
}

func (il *Bool) expressionNode() {}

func (il *Bool) TokenLiteral() string {
	return il.Token.Literal
}

func (il *Bool) String() string {
	return il.Token.Literal
}

func (il *Bool) Line() int {
	return il.Token.Line
}

func (il *Bool) Col() int {
	return il.Token.Col
}
