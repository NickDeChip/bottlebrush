package ast

import "github.com/NickDeChip/bottlebrush/pkg/token"

type ContinueStatement struct {
	Token token.Token
}

func (cs *ContinueStatement) statementNode() {}

func (cs *ContinueStatement) TokenLiteral() string {
	return cs.Token.Literal
}

func (cs *ContinueStatement) String() string {
	return cs.Token.Literal
}

func (cs *ContinueStatement) Line() int {
	return cs.Token.Line
}

func (cs *ContinueStatement) Col() int {
	return cs.Token.Col
}
