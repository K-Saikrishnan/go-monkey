package lexer

import (
	"unicode"

	"github.com/K-Saikrishnan/go-monkey/src/token"
)

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readPos int  // current reading position in input (after current char)
	ch      byte // current char under examination
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()

	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPos >= len(lex.input) {
		lex.ch = 0 // ASCII code for "NUL"
	} else {
		lex.ch = lex.input[lex.readPos]
	}
	lex.pos = lex.readPos
	lex.readPos++
}

func (lex *Lexer) peekChar() byte {
	if lex.readPos >= len(lex.input) {
		return 0
	}

	return lex.input[lex.readPos]
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)

	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)

	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)

	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)

	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if unicode.IsDigit(rune(lex.ch)) {
			tok = token.Token{
				Type:    token.INT,
				Literal: lex.readNumber(),
			}

			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	lex.readChar()

	return tok
}

func (lex *Lexer) readIdentifier() string {
	pos := lex.pos
	for unicode.IsLetter(rune(lex.ch)) {
		lex.readChar()
	}

	return lex.input[pos:lex.pos]
}

func (lex *Lexer) readNumber() string {
	position := lex.pos
	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.pos]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
