package lexer

import (
	"testing"

	"github.com/pished/exalang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	LINK 800
	GRAB 200
	ADDI F F X
	MULI X F X
	SUBI X F F
	LINK 800
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LINK, "LINK"},
		{token.INT, "800"},
		{token.GRAB, "GRAB"},
		{token.INT, "200"},
		{token.ADDI, "ADDI"},
		{token.F, "F"},
		{token.F, "F"},
		{token.X, "X"},
		{token.MULI, "MULI"},
		{token.X, "X"},
		{token.F, "F"},
		{token.X, "X"},
		{token.SUBI, "SUBI"},
		{token.X, "X"},
		{token.F, "F"},
		{token.F, "F"},
		{token.LINK, "LINK"},
		{token.INT, "800"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
