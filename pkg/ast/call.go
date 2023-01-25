package ast

import (
	"bytes"
	"strings"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))

	return out.String()
}

func (ce *CallExpression) Line() int {
	return ce.Token.Line
}
func (ce *CallExpression) Col() int {
	return ce.Token.Col
}
