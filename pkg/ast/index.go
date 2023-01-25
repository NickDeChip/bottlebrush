package ast

import (
	"bytes"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type IndexExspression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExspression) expressionNode() {}

func (ie *IndexExspression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExspression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("]")

	return out.String()
}

func (ie *IndexExspression) Line() int {
	return ie.Token.Line
}

func (ie *IndexExspression) Col() int {
	return ie.Token.Col
}
