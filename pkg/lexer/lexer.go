package lexer

import (
	"unicode"

	"github.com/NickDeChip/bottlebrush/pkg/token"
)

type Lexer struct {
	input     []rune
	index     int
	readIndex int
	line      int
	col       int
	ru        rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
		line:  1,
	}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	l.skipComments()

	switch l.ru {
	case '=':
		if l.peekRune() == '=' {
			tok = l.newTokenDouble(token.EQ)
		} else {
			tok = l.newToken(token.ASSIGN)
		}
	case ':':
		if l.peekRune() == '=' {
			tok = l.newTokenDouble(token.VAR)
		} else if l.peekRune() == ':' {
			tok = l.newTokenDouble(token.CONST)
		} else {
			tok = l.newToken(token.ILLEGAL)
		}
	case '!':
		if l.peekRune() == '=' {
			tok = l.newTokenDouble(token.NOTEQ)
		} else {
			tok = l.newToken(token.BANG)
		}
	case '>':
		if l.peekRune() == '=' {
			tok = l.newTokenDouble(token.GTEQ)
		} else {
			tok = l.newToken(token.GT)
		}
	case '<':
		if l.peekRune() == '=' {
			tok = l.newTokenDouble(token.LTEQ)
		} else {
			tok = l.newToken(token.LTEQ)
		}
	case '"':
		tok.Type = token.STRING
		tok.Col = l.col
		tok.Literal = l.readString()
		tok.Line = l.line
	case '\n':
		tok = l.newToken(token.NL)
		l.line++
	case '(':
		tok = l.newToken(token.LPAREN)
	case ')':
		tok = l.newToken(token.RPAREN)
	case ',':
		tok = l.newToken(token.COMMA)
	case '+':
		tok = l.newToken(token.ADD)
	case '-':
		tok = l.newToken(token.SUB)
	case '*':
		tok = l.newToken(token.TIMES)
	case '/':
		tok = l.newToken(token.DIV)
	case '%':
		tok = l.newToken(token.MOD)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
		tok.Col = l.col
	default:
		if isLetter(l.ru) {
			tok.Col = l.col
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ru) {
			tok.Col = l.col
			tok.Literal, tok.Type = l.readNumber()
			tok.Line = l.line
			return tok
		}

		tok = token.NewR(token.ILLEGAL, l.ru, l.line, l.col)
	}

	l.readRune()
	return tok
}

func (l *Lexer) readRune() {
	if l.readIndex >= len(l.input) {
		l.ru = 0
	} else {
		l.ru = l.input[l.readIndex]
	}
	l.index = l.readIndex
	l.readIndex++
	l.col++
}

func (l *Lexer) readIdentifier() string {
	pos := l.index
	for isLetter(l.ru) {
		l.readRune()
	}
	return string(l.input[pos:l.index])
}

func (l *Lexer) readNumber() (string, token.Type) {
	pos := l.index
	for isDigit(l.ru) {
		l.readRune()
	}
	if isDecimal(l.ru) {
		l.readRune()
		for isDigit(l.ru) {
			l.readRune()
		}
		return string(l.input[pos:l.index]), token.FLOAT
	}

	return string(l.input[pos:l.index]), token.INT
}

func (l *Lexer) readString() string {
	pos := l.index + 1
	for {
		l.readRune()
		if l.ru == '"' || l.ru == 0 {
			break
		}
	}
	return string(l.input[pos:l.index])
}

func (l *Lexer) peekRune() rune {
	if l.readIndex >= len(l.input) {
		return 0
	}
	return l.input[l.readIndex]
}

func (l *Lexer) skipWhitespace() {
	for l.ru == ' ' || l.ru == '\t' || l.ru == '\r' {
		l.readRune()
	}
}

func (l *Lexer) skipComments() {
	for l.ru == '/' && l.peekRune() == '/' {
		for l.ru != '\n' {
			l.readRune()
		}
		l.skipWhitespace()
	}
}

func (l *Lexer) newToken(tokType token.Type) token.Token {
	return token.NewR(tokType, l.ru, l.line, l.col)
}

func (l *Lexer) newTokenDouble(tokType token.Type) token.Token {
	col := l.col
	ch := l.ru
	l.readRune()
	literal := string(ch) + string(l.ru)
	return token.New(tokType, literal, l.line, col)
}

func isDigit(ru rune) bool {
	return ru >= '0' && ru <= '9'
}

func isDecimal(ru rune) bool {
	return ru == '.'
}

func isLetter(ru rune) bool {
	return ru == '_' || unicode.IsLetter(ru)
}
