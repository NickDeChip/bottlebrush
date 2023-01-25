package ast

import "github.com/NickDeChip/bottlebrush/pkg/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) Line() int {
	return es.Token.Line
}
func (es *ExpressionStatement) Col() int {
	return es.Token.Col
}
