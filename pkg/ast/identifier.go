package ast

import (
	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) Line() int {
	return i.Token.Line
}

func (i *Identifier) Col() int {
	return i.Token.Col
}
