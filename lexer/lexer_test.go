package lexer

import (
	"testing"

	"github.com/rameshputalapattu/monkey/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := `=+(){}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected = %q, actual = %q\n",
				i,
				tt.expectedType,
				tok.Type,
			)

		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected = %s, actual = %s\n",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestNextTokenAdvanced(t *testing.T) {
	input := `let five_12 = 512;
	let five =5;
	let ten = 10;

	let add = fn(x,y) {
		x + y
	}
	
	let result = add(five,ten)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five_12"},
		{token.ASSIGN, "="},
		{token.INT, "512"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected = %q, actual = %q\n",
				i,
				tt.expectedType,
				tok.Type,
			)

		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected = %s, actual = %s\n",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}

}
