package lexer_test

import (
	"testing"

	"github.com/NickDeChip/bottlebrush/pkg/lexer"
	"github.com/NickDeChip/bottlebrush/pkg/token"
	"github.com/stretchr/testify/assert"
)

func Test_SimpleIntAssign(t *testing.T) {
	lex := lexer.New("apple := 10")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.INT, "10", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 12), lex.NextToken())
}

func Test_SimpleFuncCall(t *testing.T) {
	lex := lexer.New(`say("Hello, World!")`)

	assert.Equal(t, token.New(token.IDENT, "say", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.LPAREN, "(", 1, 4), lex.NextToken())
	assert.Equal(t, token.New(token.STRING, "Hello, World!", 1, 5), lex.NextToken())
	assert.Equal(t, token.New(token.RPAREN, ")", 1, 20), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 21), lex.NextToken())
}

func Test_SimpleFloatAssign(t *testing.T) {
	lex := lexer.New("apple := 10.22")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.FLOAT, "10.22", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 15), lex.NextToken())
}

func Test_ComplexFloatAssign(t *testing.T) {
	lex := lexer.New("apple := 10.")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.FLOAT, "10.", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 13), lex.NextToken())
}

func Test_SimpleStringAssign(t *testing.T) {
	lex := lexer.New("apple := \"hello\" ")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.STRING, "hello", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 18), lex.NextToken())
}

func Test_ComplexStringAssign(t *testing.T) {
	lex := lexer.New("apple := \"hello")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.STRING, "hello", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 17), lex.NextToken())
}
func Test_MultiLinwStringAssign(t *testing.T) {
	lex := lexer.New("apple := \"hello\nworld\"")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.STRING, "hello\nworld", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 1, 23), lex.NextToken())
}

func Test_NewLine(t *testing.T) {
	lex := lexer.New("apple := 0\npie := 1\n")

	assert.Equal(t, token.New(token.IDENT, "apple", 1, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 1, 7), lex.NextToken())
	assert.Equal(t, token.New(token.INT, "0", 1, 10), lex.NextToken())
	assert.Equal(t, token.New(token.NL, "\n", 1, 11), lex.NextToken())
	assert.Equal(t, token.New(token.IDENT, "pie", 2, 1), lex.NextToken())
	assert.Equal(t, token.New(token.VAR, ":=", 2, 5), lex.NextToken())
	assert.Equal(t, token.New(token.INT, "1", 2, 8), lex.NextToken())
	assert.Equal(t, token.New(token.NL, "\n", 2, 9), lex.NextToken())
	assert.Equal(t, token.New(token.EOF, "", 3, 1), lex.NextToken())
}
