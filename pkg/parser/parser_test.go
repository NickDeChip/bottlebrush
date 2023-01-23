package parser_test

import (
	"fmt"
	"testing"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/lexer"
	"github.com/NickDeChip/bottle-brush/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func Test_simpleString(t *testing.T) {
	l := lexer.New(`"hello"`)
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt := prog.Statements[0].(*ast.ExpressionStatement)
	literal := stmt.Expression.(*ast.StringLiteral)

	assert.Equal(t, "hello", literal.Value)
}

func Test_simpleInt(t *testing.T) {
	l := lexer.New("33")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt := prog.Statements[0].(*ast.ExpressionStatement)
	literal := stmt.Expression.(*ast.IntegerLiteral)

	assert.Equal(t, int32(33), literal.Value)
}

func Test_simpleFloat(t *testing.T) {
	l := lexer.New("33.3")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt := prog.Statements[0].(*ast.ExpressionStatement)
	literal := stmt.Expression.(*ast.FloatLiteral)

	assert.Equal(t, float32(33.3), literal.Value)
}

func Test_SimpleDeclaration(t *testing.T) {
	l := lexer.New("pie := \"hello\"")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt, ok := prog.Statements[0].(*ast.VarStatement)
	assert.True(t, ok, fmt.Sprintf("expect statement to be *VarStatement, got %T", stmt))
	assert.Equal(t, ":=", stmt.TokenLiteral())
	assert.Equal(t, "pie", stmt.Name.Value)
	assert.Equal(t, "hello", stmt.Value.TokenLiteral())
	assert.True(t, stmt.Mut)
}

func Test_SimpleDeclarationConst(t *testing.T) {
	l := lexer.New("pie :: \"hello\"")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt, ok := prog.Statements[0].(*ast.VarStatement)
	assert.True(t, ok, fmt.Sprintf("expect statement to be *VarStatement, got %T", stmt))
	assert.Equal(t, "::", stmt.TokenLiteral())
	assert.Equal(t, "pie", stmt.Name.Value)
	assert.Equal(t, "hello", stmt.Value.TokenLiteral())
	assert.False(t, stmt.Mut)
}

func Test_SimpleAssignment(t *testing.T) {
	l := lexer.New("pie = \"hello\"")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt, ok := prog.Statements[0].(*ast.AssignmentStatement)
	assert.True(t, ok, fmt.Sprintf("expect statement to be *VarStatement, got %T", stmt))
	assert.Equal(t, "=", stmt.TokenLiteral())
	assert.Equal(t, "pie", stmt.Name.Value)
	assert.Equal(t, "hello", stmt.Value.TokenLiteral())
}

func Test_simpleFunctionCall(t *testing.T) {
	l := lexer.New("say(10)")
	p := parser.New(l)

	prog := p.ParseProgram()

	assert.Empty(t, p.Errors())

	assert.Len(t, prog.Statements, 1)

	stmt := prog.Statements[0].(*ast.ExpressionStatement)
	literal := stmt.Expression.(*ast.CallExpression)

	assert.Equal(t, "say", literal.Function.String())
	assert.Equal(t, "10", literal.Arguments[0].String())
}
