package ast

type Node interface {
	TokenLiteral() string
	String() string
	Line() int
	Col() int
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
