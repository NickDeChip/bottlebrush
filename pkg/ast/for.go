package ast

import (
	"bytes"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type ForStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (fs *ForStatement) statementNode() {}

func (fs *ForStatement) TokenLiteral() string {
	return fs.Token.Literal
}

func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("for")
	out.WriteString(fs.Condition.String())
	out.WriteString(" ")
	out.WriteString(fs.Consequence.String())

	return out.String()
}

func (fs *ForStatement) Line() int {
	return fs.Token.Line
}

func (fs *ForStatement) Col() int {
	return fs.Token.Col
}
