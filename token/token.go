package token

type (
	TokenType string
	Token     struct {
		Type    TokenType
		Literal string
	}
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// idents + literal
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "ASSIGN"
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//  keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
