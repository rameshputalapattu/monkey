package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and Literals
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func LookupIdent(ident string) TokenType {
	var keywords = map[string]TokenType{
		"fn":  FUNCTION,
		"let": LET,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
