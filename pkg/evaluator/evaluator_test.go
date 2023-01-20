package evaluator_test

import (
	"testing"

	"github.com/NickDeChip/bottle-brush/pkg/evaluator"
	"github.com/NickDeChip/bottle-brush/pkg/lexer"
	"github.com/NickDeChip/bottle-brush/pkg/object"
	"github.com/NickDeChip/bottle-brush/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func Test_EvalStringLiteral(t *testing.T) {
	input := "\"Hello, World\""

	evaluated := testEval(input)
	str := evaluated.(*object.String)
	assert.Equal(t, "Hello, World", str.Value)
}

func Test_EvalVarStatement(t *testing.T) {
	input := "pear := \"an apple\""
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()
	env := object.NewEnvironment()

	evaluator.Eval(prog, env)

	val, ok := env.Get("pear")
	assert.True(t, ok)
	assert.Equal(t, "an apple", val.Inspect())
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(prog, env)
}
