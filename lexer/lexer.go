package lexer

import (
	"github.com/rameshputalapattu/monkey/token"
)

type Lexer struct {
	input        string
	position     int  //current position in the input (points to current char)
	readPosition int  //current reading position in the input(after current char)
	ch           byte //the character being examined
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++

}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()

	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	var isMultiChar bool

	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		tok = newtoken(token.ASSIGN, '=')
	case '+':
		tok = newtoken(token.PLUS, '+')
	case '(':
		tok = newtoken(token.LPAREN, '(')

	case ')':
		tok = newtoken(token.RPAREN, ')')

	case '{':
		tok = newtoken(token.LBRACE, '{')

	case '}':
		tok = newtoken(token.RBRACE, '}')
	case ';':

		tok = newtoken(token.SEMICOLON, ';')
	case ',':
		tok = newtoken(token.COMMA, ',')

	default:

		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok.Type = token.LookupIdent(ident)
			tok.Literal = ident
			isMultiChar = true

		} else if isDigit(l.ch) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			isMultiChar = true

		} else {
			tok = newtoken(token.ILLEGAL, l.ch)
		}
	}
	if !isMultiChar {
		l.readChar()
	}

	return tok

}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]

}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]

}

func newtoken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
