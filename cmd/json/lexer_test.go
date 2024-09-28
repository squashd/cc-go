package main

import "testing"

func TestNextToken(t *testing.T) {
	input := `
{
  "name": "John",
  "balance": .1,
  "is_admin": false,
  "is_member": true,
  "timezone": null,
  "friends": [
    "Alice",
    "Bob",
    "Charlie"
  ]
}
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LBRACE, "{"},
		{STRING, "name"},
		{COLON, ":"},
		{STRING, "John"},
		{COMMA, ","},
		{STRING, "balance"},
		{COLON, ":"},
		{FLOAT, ".1"},
		{COMMA, ","},
		{STRING, "is_admin"},
		{COLON, ":"},
		{FALSE, "false"},
		{COMMA, ","},
		{STRING, "is_member"},
		{COLON, ":"},
		{TRUE, "true"},
		{COMMA, ","},
		{STRING, "timezone"},
		{COLON, ":"},
		{NULL, "null"},
		{COMMA, ","},
		{STRING, "friends"},
		{COLON, ":"},
		{LBRACKET, "["},
		{STRING, "Alice"},
		{COMMA, ","},
		{STRING, "Bob"},
		{COMMA, ","},
		{STRING, "Charlie"},
		{RBRACKET, "]"},
		{RBRACE, "}"},
		{EOF, ""},
	}

	lexer := NewLexer(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
