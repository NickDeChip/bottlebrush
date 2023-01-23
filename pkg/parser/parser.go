package parser

import (
	"fmt"
	"strconv"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/lexer"
	"github.com/NickDeChip/bottle-brush/pkg/token"
)

const (
	_ int = iota
	LOWEST
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedence = map[token.Type]int{
	token.ADD:    SUM,
	token.SUB:    SUM,
	token.DIV:    PRODUCT,
	token.TIMES:  PRODUCT,
	token.MOD:    PRODUCT,
	token.LPAREN: CALL,
}

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
	p.registerPrefix(token.SUB, p.parsePrefixExpression)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpressoin)
	p.registerPrefix(token.TRUE, p.parseBool)
	p.registerPrefix(token.FALSE, p.parseBool)
	p.registerPrefix(token.FN, p.parseFunctionLiteral)

	p.infixParseFns = make(map[token.Type]infixParseFn)
	p.registerInfix(token.ADD, p.parseInfixExpression)
	p.registerInfix(token.SUB, p.parseInfixExpression)
	p.registerInfix(token.DIV, p.parseInfixExpression)
	p.registerInfix(token.TIMES, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

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
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IDENT:
		if p.isAssignment() {
			return p.parseAssignment()
		} else if p.isDeclaration() {
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
	if p.curTokenIs(token.NL) {
		return nil
	}
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.NL) && precdence < p.peekPrecedence() {
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

func (p *Parser) peekPrecedence() int {
	if p, ok := precedence[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseAssignment() *ast.AssignmentStatement {
	stmt := &ast.AssignmentStatement{}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	p.nextToken()

	stmt.Token = p.curToken

	p.nextToken()

	stmt.Value = p.parserExpression(LOWEST)

	if p.peekTokenIs(token.NL) {
		p.nextToken()
	}
	return stmt
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

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{
		Token: p.curToken,
	}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 32)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer; line=%d; col=%d", p.curToken.Literal, p.curToken.Line, p.curToken.Col)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = int32(val)

	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{
		Token: p.curToken,
	}

	val, err := strconv.ParseFloat(p.curToken.Literal, 32)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float; line=%d; col=%d", p.curToken.Literal, p.curToken.Line, p.curToken.Col)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = float32(val)

	return lit
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{
		Token: p.curToken,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.START) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.END) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parserExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parserExpression(LOWEST))
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseGroupedExpressoin() ast.Expression {
	p.nextToken()

	exp := p.parserExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parserExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expresssion := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecdence()
	p.nextToken()
	expresssion.Right = p.parserExpression(precedence)

	return expresssion
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {

	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	if p.peekTokenIs(token.END) {
		stmt.ReturnValue = nil
		return stmt
	}

	p.nextToken()

	if p.curTokenIs(token.NL) {
		stmt.ReturnValue = nil
		return stmt
	}

	exp := p.parserExpression(LOWEST)
	stmt.ReturnValue = exp

	return stmt
}

func (p *Parser) parseBool() ast.Expression {
	return &ast.Bool{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) curPrecdence() int {
	if p, ok := precedence[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
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

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function %s found; line=%d; col=%d", t, p.curToken.Line, p.curToken.Col)
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

func (p *Parser) isAssignment() bool {
	return p.peekTokenIs(token.ASSIGN)
}
