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

func Test_EvalIntegerLiteral(t *testing.T) {
	input := "69"

	evaluated := testEval(input)
	str := evaluated.(*object.Integer)
	assert.Equal(t, int32(69), str.Value)
}

func Test_EvalFloatLiteral(t *testing.T) {
	input := "69.69"

	evaluated := testEval(input)
	str := evaluated.(*object.Float)
	assert.Equal(t, float32(69.69), str.Value)
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
	assert.Equal(t, "an apple", val.Object.Inspect())
}

func Test_EvalVarAndIdentifier(t *testing.T) {
	input := "apple :: \"Hello, World\"\napple"

	evaluated := testEval(input)
	str := evaluated.(*object.String)
	assert.Equal(t, "Hello, World", str.Value)
}

func Test_EvalVarAlreadyVar(t *testing.T) {
	input := `apple := "goodbye, world"
  apple := "goodbye, world"`

	evaluated := testEval(input)
	err := evaluated.(*object.Error)
	assert.Equal(t, "Error: the identifier apple is already declared; line=2; col=3", err.Inspect())
}

func Test_EvalAssignment(t *testing.T) {
	input := `apple := "hello, world"
  apple = "goodbye, world"
  apple`

	evaluated := testEval(input)
	str := evaluated.(*object.String)
	assert.Equal(t, "goodbye, world", str.Value)
}

func Test_EvalAssignmentConst(t *testing.T) {
	input := `apple :: "hello, world"
  apple = "goodbye, world"`

	evaluated := testEval(input)
	err := evaluated.(*object.Error)
	assert.Equal(t, "Error: The identifier apple is a constant value and can not be chaged; line=2; col=3", err.Inspect())
}

func Test_EvalAssignmentNull(t *testing.T) {
	input := `apple = "goodbye, world"`

	evaluated := testEval(input)
	err := evaluated.(*object.Error)
	assert.Equal(t, "Error: the identifier apple is not declared; line=1; col=1", err.Inspect())
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(prog, env)
}
