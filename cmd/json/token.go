package main

type (
	TokenType string

	Token struct {
		Type    TokenType
		Literal string
	}
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literal types
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	// Boolean and null
	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"

	// Delimiters
	COMMA = ","
	COLON = ":"

	// Braces and brackets
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
)

var keywords = map[string]TokenType{
	"false": FALSE,
	"null":  NULL,
	"true":  TRUE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return ILLEGAL
}
