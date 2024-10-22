
package lexer

import (
	"testing"

	"example/user/interpreter/token"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`


	// Parameterized tests
	tests := []struct{

		expectedType token.TokenType
		expectedLiteral string

	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// Initializing the lexer
	l := New(input)

	// Run tests
	for i, tt := range tests {
		tok := l.NextToken()

		// verification of token type
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		// verification of token literal
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}


}