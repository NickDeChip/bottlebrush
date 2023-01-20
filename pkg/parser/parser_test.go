package parser_test

import (
	//"fmt"
	"fmt"
	"testing"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/lexer"
	"github.com/NickDeChip/bottle-brush/pkg/parser"
	"github.com/stretchr/testify/assert"
)

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
