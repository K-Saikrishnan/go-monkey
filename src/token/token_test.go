package token_test

import (
	"testing"

	"github.com/K-Saikrishnan/go-monkey/src/token"
)

func TestLookupIndent(t *testing.T) {
	tests := []struct {
		input    string
		expected token.Type
	}{
		{"fn", token.FUNCTION},
		{"let", token.LET},
		{"true", token.TRUE},
		{"false", token.FALSE},
		{"if", token.IF},
		{"else", token.ELSE},
		{"return", token.RETURN},
		{"foobar", token.IDENT},
		{"", token.IDENT},
	}

	for _, tt := range tests {
		tok := token.LookupIdent(tt.input)
		if tok != tt.expected {
			t.Fatalf("token.LookupIdent(%q) expected-%q, got-%q", tt.input, tok, tt.expected)
		}
	}
}
