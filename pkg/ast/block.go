package ast

import (
	"bytes"

	"github.com/NickDeChip/bottle-brush/pkg/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (bs *BlockStatement) Line() int {
	return bs.Token.Line
}
func (bs *BlockStatement) Col() int {
	return bs.Token.Col
}