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
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	MINUS    = "-"
	EQ       = "=="
	NOT_EQ   = "!="

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
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

func LookupIdent(ident string) TokenType {
	var keywords = map[string]TokenType{
		"fn":     FUNCTION,
		"let":    LET,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
