package token

type Type string

type Token struct {
	Type    Type
	Literal string
	Line    int
	Col     int
}

func NewR(TokenType Type, ru rune, line int, col int) Token {
	return Token{
		Type:    TokenType,
		Literal: string(ru),
		Line:    line,
		Col:     col,
	}
}

func New(TokenType Type, literal string, line int, col int) Token {
	return Token{
		Type:    TokenType,
		Literal: literal,
		Line:    line,
		Col:     col,
	}
}

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	IDENT Type = "IDENT"

	// Literals
	INT    Type = "INT"
	FLOAT  Type = "FLOAT"
	STRING Type = "STRING"

	// Operators
	ADD   Type = "+"
	SUB   Type = "-"
	TIMES Type = "*"
	DIV   Type = "/"
	MOD   Type = "%"

	ASSIGN Type = "="

	VAR   Type = ":="
	CONST Type = "::"

	// Delimiters
	NL     Type = "NL"
	LPAREN Type = "("
	RPAREN Type = ")"
	COMMA  Type = ","

	// Keywords
	START  Type = "START"
	END    Type = "END"
	TRUE   Type = "TRUE"
	FALSE  Type = "FALSE"
	FN     Type = "FN"
	RETURN Type = "RETURN"
	FOR    Type = "FOR"
	IN     Type = "IN"
	IF     Type = "IF"
)

var keywords = map[string]Type{
	"start":  START,
	"end":    END,
	"true":   TRUE,
	"false":  FALSE,
	"fn":     FN,
	"return": RETURN,
	"for":    FOR,
	"in":     IN,
	"if":     IF,
}

func LookUpIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
