package ast

import (
	"bytes"
	"strings"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type ListLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (ll *ListLiteral) expressionNode() {}

func (ll *ListLiteral) TokenLiteral() string {
	return ll.Token.Literal
}

func (ll *ListLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range ll.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ll *ListLiteral) Line() int {
	return ll.Token.Line
}

func (ll *ListLiteral) Col() int {
	return ll.Token.Col
}
