package ast

import "github.com/NickDeChip/bottlebrush/pkg/token"

type BreakStatement struct {
	Token token.Token
}

func (bs *BreakStatement) statementNode() {}

func (bs *BreakStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BreakStatement) String() string {
	return bs.Token.Literal
}

func (bs *BreakStatement) Line() int {
	return bs.Token.Line
}

func (bs *BreakStatement) Col() int {
	return bs.Token.Col
}
