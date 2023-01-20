package parser

import (
	"fmt"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/lexer"
	"github.com/NickDeChip/bottle-brush/pkg/token"
)

const (
	_ int = iota
	LOWEST
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.registerPrefix(token.STRING, p.parseStringLiteral)

	p.infixParseFns = make(map[token.Type]infixParseFn)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.IDENT:
		// TODO: Handle Assighnemt
		if p.isDeclaration() {
			return p.parseDeclaration()
		}
		return p.parseExpressionStatement()

	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parserExpression(LOWEST)

	if p.peekTokenIs(token.NL) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parserExpression(precdence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.NL) {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)

	}
	return leftExp
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseDeclaration() *ast.VarStatement {
	stmt := &ast.VarStatement{}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	p.nextToken()

	stmt.Token = p.curToken
	if p.curToken.Type == token.VAR {
		stmt.Mut = true
	}

	p.nextToken()

	stmt.Value = p.parserExpression(LOWEST)
	// TODO: FUNCTION LITERAL

	if p.peekTokenIs(token.NL) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("exspected next token to be %s, got %s instead; line=%d; col=%d", t, p.peekToken.Type, p.peekToken.Line, p.peekToken.Col)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerInfix(Type token.Type, fn infixParseFn) {
	p.infixParseFns[Type] = fn
}

func (p *Parser) registerPrefix(Type token.Type, fn prefixParseFn) {
	p.prefixParseFns[Type] = fn
}

func (p *Parser) isDeclaration() bool {
	return p.peekTokenIs(token.VAR) || p.peekTokenIs(token.CONST)
}
